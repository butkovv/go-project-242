package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "pathsize",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)
			size, err := GetSize(path)
			if err != nil {
				return err
			}
			fmt.Printf("%dB	%s\n", size, path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		slog.Error(err.Error())
	}
}

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
