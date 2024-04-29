package golocates

import (
	"embed"
	"testing"
)

//go:embed locate-files/*.json
var fs embed.FS

func TestLoad(t *testing.T) { Load(fs, "locate-files") }
