package code

import (
	"testing"
)

func TestPathSizeFile(t *testing.T) {
	path := "testdata/dir/file"
	want := 34256
	inclHidden := false
	recursive := false
	size, err := GetPathSize(path, inclHidden, recursive)
	if want != size || err != nil {
		t.Errorf(`GetPathSize(%q, %t, %t) = %d, %v, want match for %d, nil`, path, inclHidden, recursive, size, err, want)
	}
}

func TestPathSizeDirectory(t *testing.T) {
	path := "testdata/dir"
	want := 205536
	inclHidden := false
	recursive := false
	size, err := GetPathSize(path, inclHidden, recursive)
	if want != size || err != nil {
		t.Errorf(`GetPathSize(%q, %t, %t) = %d, %v, want match for %d, nil`, path, inclHidden, recursive, size, err, want)
	}
}

func TestPathSizeNonExistent(t *testing.T) {
	path := "testdata/dir/file3"
	inclHidden := false
	recursive := false
	size, err := GetPathSize(path, inclHidden, recursive)
	if err == nil {
		t.Errorf(`GetPathSize(%q, %t, %t) = %d, %v, want match for 0, err`, path, inclHidden, recursive, size, err)
	}
}

func TestPathSizeHiddenFileInclHidden(t *testing.T) {
	path := "testdata/dir/.file3"
	want := 102768
	inclHidden := true
	recursive := false
	size, err := GetPathSize(path, inclHidden, recursive)
	if want != size || err != nil {
		t.Errorf(`GetPathSize(%q, %t, %t) = %d, %v, want match for %d, nil`, path, inclHidden, recursive, size, err, want)
	}
}

func TestPathSizeHiddenFileExclHidden(t *testing.T) {
	path := "testdata/dir/.file3"
	want := 0
	inclHidden := false
	recursive := false
	size, err := GetPathSize(path, inclHidden, recursive)
	if want != size || err != nil {
		t.Errorf(`GetPathSize(%q, %t, %t) = %d, %v, want match for %d, nil`, path, inclHidden, recursive, size, err, want)
	}
}

func TestPathSizeDirectoryInclHidden(t *testing.T) {
	path := "testdata/dir/"
	want := 308304
	inclHidden := true
	recursive := false
	size, err := GetPathSize(path, inclHidden, recursive)
	if want != size || err != nil {
		t.Errorf(`GetPathSize(%q, %t, %t) = %d, %v, want match for %d, nil`, path, inclHidden, recursive, size, err, want)
	}
}

func TestPathSizeDirectoryInclHiddenRecursive(t *testing.T) {
	path := "testdata/dir"
	want := 359476
	inclHidden := true
	recursive := true
	size, err := GetPathSize(path, inclHidden, recursive)
	if want != size || err != nil {
		t.Errorf(`GetPathSize(%q, %t, %t) = %d, %v, want match for %d, nil`, path, inclHidden, recursive, size, err, want)
	}
}

func TestFormatSizeBytesHuman(t *testing.T) {
	size := 100
	want := "100.0B"
	human := true
	formattedSize := FormatSize(size, human)
	if want != formattedSize {
		t.Errorf(`FormatSize(%d, %t) = %s, want match for %s`, size, human, formattedSize, want)
	}
}

func TestFormatSizeBytesRaw(t *testing.T) {
	size := 200
	want := "200B"
	human := false
	formattedSize := FormatSize(size, human)
	if want != formattedSize {
		t.Errorf(`FormatSize(%d, %t) = %s, want match for %s`, size, human, formattedSize, want)
	}
}

func TestFormatSizeExabytesHuman(t *testing.T) {
	size := 1152921505706846976
	want := "1.0EB"
	human := true
	formattedSize := FormatSize(size, human)
	if want != formattedSize {
		t.Errorf(`FormatSize(%d, %t) = %s, want match for %s`, size, human, formattedSize, want)
	}
}

func TestFormatSizeExabytesRaw(t *testing.T) {
	size := 2152921505706846976
	want := "2152921505706846976B"
	human := false
	formattedSize := FormatSize(size, human)
	if want != formattedSize {
		t.Errorf(`FormatSize(%d, %t) = %s, want match for %s`, size, human, formattedSize, want)
	}
}

func TestFormatSizeZeroHuman(t *testing.T) {
	size := 0
	want := "0.0B"
	human := true
	formattedSize := FormatSize(size, human)
	if want != formattedSize {
		t.Errorf(`FormatSize(%d, %t) = %s, want match for %s`, size, human, formattedSize, want)
	}
}

func TestFormatSizeNegativeRaw(t *testing.T) {
	size := -100
	want := "0B"
	human := false
	formattedSize := FormatSize(size, human)
	if want != formattedSize {
		t.Errorf(`FormatSize(%d, %t) = %s, want match for %s`, size, human, formattedSize, want)
	}
}
