package code

import (
	"fmt"
	"os"
	"strings"
)

func GetPathSize(path string, recursive, humanReadable, inclHidden bool) (string, error) {
	s, err := getSize(path, inclHidden, recursive)
	if err != nil {
		return formatSize(0, humanReadable), err
	}
	return formatSize(s, humanReadable), nil
}

func getSize(path string, inclHidden bool, recursive bool) (int, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	size := 0
	isHidden := len(info.Name()) > 1 && strings.HasPrefix(info.Name(), ".")
	if info.IsDir() && (!isHidden || inclHidden) {
		files, err := os.ReadDir(path)
		if err != nil {
			return 0, err
		}
		for _, f := range files {
			if f.IsDir() && recursive {
				p := path
				if strings.HasSuffix(p, "/") {
					p += f.Name()
				} else {
					p = p + "/" + f.Name()
				}
				s, err := getSize(p, inclHidden, recursive)
				if err != nil {
					return 0, err
				}
				size += s
			} else if !f.IsDir() {
				info, err := f.Info()
				if err != nil {
					return 0, err
				}
				isHidden := strings.HasPrefix(info.Name(), ".")
				if !isHidden || inclHidden {
					size += int(info.Size())
				}
			}
		}
	} else {
		if !isHidden || inclHidden {
			size = int(info.Size())
		}
	}
	return size, nil
}

func formatSize(size int, human bool) string {
	const (
		B  = 1
		KB = 1 << 10
		MB = 1 << 20
		GB = 1 << 30
		TB = 1 << 40
		PB = 1 << 50
		EB = 1 << 60
	)
	unit := "B"
	if size < 0 {
		size = 0
	}
	if !human || size < KB {
		return fmt.Sprintf("%d%s", size, unit)
	}
	formattedSize := float64(size)
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
