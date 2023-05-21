package logext

import (
	"path/filepath"
	"runtime"
)

func normalizeOutput(fileName string) string {
	_, f, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(f), "../logs", fileName)
}
