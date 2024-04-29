package golocates

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-the-way/golocal"
)

var (
	logger    = log.New(os.Stdout, "[golocates] ", log.LstdFlags|log.Lshortfile)
	locateMap = map[string]map[string]string{}
)

func Load(fs embed.FS, name string) {
	dirs, _ := fs.ReadDir(name)
	for _, dir := range dirs {
		if !dir.IsDir() {
			dirName := dir.Name()
			if buf, err := fs.ReadFile(name + "/" + dirName); err != nil {
				fmt.Println(fmt.Sprintf("load locate %s err: %v", dirName, err))
			} else {
				var m map[string]string
				if err = json.Unmarshal(buf, &m); err != nil {
					logger.Println(fmt.Sprintf("load locate %s err: %v", dirName, err))
				} else {
					idx := strings.LastIndex(dirName, ".")
					loc := dirName[:idx]
					locateMap[loc] = m
					logger.Println("loaded locate `" + loc + "` file `" + dirName + "`")
				}
			}
		}
	}
}

func Get(key string, args ...any) (value string) {
	lang := golocal.Get()
	mm, ok := locateMap[lang]
	if !ok {
		return
	}
	return fmt.Sprintf(mm[key], args...)
}

func GetError(key string, args ...any) (err error) {
	return errors.New(Get(key, args...))
}
