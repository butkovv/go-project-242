package code

import (
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
