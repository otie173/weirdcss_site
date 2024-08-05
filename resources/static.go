package resources

import "embed"

//go:embed static/*
var StaticFiles embed.FS

func GetStaticFile(path string) ([]byte, error) {
	return StaticFiles.ReadFile(path)
}
