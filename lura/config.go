package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func readConfig() []string {
	dir := "lura/configs/routers/"

	var configFile []string
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".json" {
			configFile = append(configFile, path)
		}
		return nil

	})
	if err != nil {
		fmt.Println("Error readConfig", err)
	}
	return configFile
}
