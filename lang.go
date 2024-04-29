package golocates

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-the-way/golocal"
)

var (
	//go:embed locates/*.json
	locatesFS embed.FS
	// <LANG,<KEY,VALUE>>
	langMap = map[string]map[string]string{}
)

func Load(fs embed.FS) {
	dirs, _ := locatesFS.ReadDir("locates")
	for _, dir := range dirs {
		if !dir.IsDir() {
			name := dir.Name()
			if buf, err := locatesFS.ReadFile("locates/" + name); err != nil {
				fmt.Println(fmt.Sprintf("load locate %s err: %v", name, err))
			} else {
				var m map[string]string
				if err = json.Unmarshal(buf, &m); err != nil {
					fmt.Println(fmt.Sprintf("load locate %s err: %v", name, err))
				} else {
					idx := strings.LastIndex(name, ".")
					langMap[name[:idx]] = m
					fmt.Println("loaded locate file:", name)
				}
			}
		}
	}
}

func GetError(key string, args ...any) (err error) {
	return errors.New(Get(key, args...))
}

func Get(key string, args ...any) (value string) {
	lang := golocal.Get()
	mm, ok := langMap[lang]
	if !ok {
		return
	}
	return fmt.Sprintf(mm[key], args...)
}
