package main

import (
	"cyber-space/cmd"
	"embed"
	"io/fs"
)

//go:embed ui/dist/spa/*
var f embed.FS

type assertFs struct {
	fs     embed.FS
	prefix string
}

func (a *assertFs) Open(name string) (fs.File, error) {
	return a.fs.Open(a.prefix + name)
}

func main() {
	cmd.AddStaticFS(&assertFs{
		fs:     f,
		prefix: "ui/dist/spa/",
	})
	cmd.Execute()
}
