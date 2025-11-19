package code

import (
	"fmt"
	"os"
)

func GetSize(path string) (int, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	size := 0
	if info.IsDir() {
		files, err := os.ReadDir(path)
		if err != nil {
			return 0, err
		}
		for _, f := range files {
			if !f.IsDir() {
				info, err := f.Info()
				if err != nil {
					return 0, err
				}
				size += int(info.Size())
			}
		}
	} else {
		size = int(info.Size())
	}
	return size, nil
}

func FormatSize(size int, human bool) string {
	if size < 0 {
		size = 0
	}
	if !human {
		return fmt.Sprintf("%dB", size)
	}
	const (
		B  = 1
		KB = 1 << 10 // 2^10
		MB = 1 << 20 // 2^20
		GB = 1 << 30 // 2^30
		TB = 1 << 40 // 2^40
		PB = 1 << 50 // 2^50
		EB = 1 << 60 // 2^60
	)
	formattedSize := float64(size)
	unit := "B"
	switch {
	case size > EB:
		formattedSize = float64(size) / EB
		unit = "EB"
	case size > PB:
		formattedSize = float64(size) / PB
		unit = "PB"
	case size > TB:
		formattedSize = float64(size) / TB
		unit = "TB"
	case size > GB:
		formattedSize = float64(size) / GB
		unit = "GB"
	case size > MB:
		formattedSize = float64(size) / MB
		unit = "MB"
	case size > KB:
		formattedSize = float64(size) / KB
		unit = "KB"
	}
	return fmt.Sprintf("%.1f%s", formattedSize, unit)
}
