package util

import "path/filepath"

var MultiExts = map[string]struct{}{
	".bz2": {},
	".gz":  {},
	".xz":  {},
}

func SmarterExt(path string) string {
	ext := filepath.Ext(path)
	if _, ok := MultiExts[ext]; ok {
		ext = filepath.Ext(path[:len(path)-len(ext)]) + ext
	}
	return ext
}
