package resources

import "embed"

//go:embed template/*
var TemplateFiles embed.FS

func GetTemplateFile(path string) ([]byte, error) {
	return TemplateFiles.ReadFile(path)
}
