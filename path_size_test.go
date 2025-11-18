package code

import (
	"testing"
)

func TestPathSizeFile(t *testing.T) {
	path := "testdata/dir/file"
	want := 34256
	size, err := GetSize(path)
	if want != size || err != nil {
		t.Errorf(`GetSize(%q) = %d, %v, want match for %d, nil`, path, size, err, want)
	}
}

func TestPathSizeDirectory(t *testing.T) {
	path := "testdata/dir"
	want := 205536
	size, err := GetSize(path)
	if want != size || err != nil {
		t.Errorf(`GetSize(%q) = %d, %v, want match for %d, nil`, path, size, err, want)
	}
}

func TestPathSizeEmptyDirectory(t *testing.T) {
	path := "testdata/dir/subdir"
	want := 0
	size, err := GetSize(path)
	if want != size || err != nil {
		t.Errorf(`GetSize(%q) = %d, %v, want match for %d, nil`, path, size, err, want)
	}
}

func TestPathSizeNonExistent(t *testing.T) {
	path := "testdata/dir/file3"
	size, err := GetSize(path)
	if err == nil {
		t.Errorf(`GetSize(%q) = %d, %v, want match for 0, err`, path, size, err)
	}
}
