package config

import "path/filepath"

func toSystemPath(path string) string {
	return filepath.FromSlash(path)
}

func fromSystemPath(path string) string {
	return filepath.ToSlash(path)
}
