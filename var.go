package golocates

import "github.com/go-the-way/golocal"

func GetLang() string     { return golocal.Get() }
func SetLang(lang string) { golocal.Set(lang) }
