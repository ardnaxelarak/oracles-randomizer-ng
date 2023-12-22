package main

import (
	"embed"
	"io/fs"

	"gopkg.in/yaml.v2"
)

//go:embed hints/* logic/* romdata/*
var embedded embed.FS

func readYaml(filename string, out interface{}) error {
	b, err := embedded.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return yaml.Unmarshal(b, out)
}

func readDir(filename string) ([]fs.DirEntry, error) {
	return embedded.ReadDir(filename)
}
