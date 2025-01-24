package globi

import (
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestGlob(t *testing.T) {
	testDir := t.TempDir()

	files := []string{
		"File.txt",
		"UPPER.TXT",
		"lower.txt",
		"mixed.TXT",
		"other.dat",
		"test.go",
		"prefix_file.txt",
		"file_suffix.txt",
		"Readme.txt",
	}

	for _, f := range files {
		if err := os.WriteFile(filepath.Join(testDir, f), []byte("test"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	tests := []struct {
		pattern string
		want    []string
		wantErr bool
	}{
		{
			pattern: filepath.Join(testDir, "*.txt"),
			want: []string{
				filepath.Join(testDir, "File.txt"),
				filepath.Join(testDir, "UPPER.TXT"),
				filepath.Join(testDir, "lower.txt"),
				filepath.Join(testDir, "mixed.TXT"),
				filepath.Join(testDir, "prefix_file.txt"),
				filepath.Join(testDir, "file_suffix.txt"),
				filepath.Join(testDir, "Readme.txt"),
			},
		},
		{
			pattern: filepath.Join(testDir, "FILE.TXT"),
			want:    []string{filepath.Join(testDir, "File.txt")},
		},
		{
			pattern: filepath.Join(testDir, "*_file.txt"),
			want:    []string{filepath.Join(testDir, "prefix_file.txt")},
		},
		{
			pattern: filepath.Join(testDir, "file_*.txt"),
			want:    []string{filepath.Join(testDir, "file_suffix.txt")},
		},
		{
			pattern: filepath.Join(testDir, "nonexistent*"),
			want:    nil,
		},
		{
			pattern: filepath.Join(testDir, "RE*"),
			want:    []string{filepath.Join(testDir, "Readme.txt")},
		},
		{
			pattern: filepath.Join(testDir, "*"),
			want: []string{
				filepath.Join(testDir, "File.txt"),
				filepath.Join(testDir, "UPPER.TXT"),
				filepath.Join(testDir, "lower.txt"),
				filepath.Join(testDir, "mixed.TXT"),
				filepath.Join(testDir, "other.dat"),
				filepath.Join(testDir, "test.go"),
				filepath.Join(testDir, "prefix_file.txt"),
				filepath.Join(testDir, "file_suffix.txt"),
				filepath.Join(testDir, "Readme.txt"),
			},
		},
	}

	for _, tt := range tests {
		got, err := Glob(tt.pattern)
		if (err != nil) != tt.wantErr {
			t.Errorf("pattern %q: got error %v, wantErr %v", tt.pattern, err, tt.wantErr)
			continue
		}
		if err != nil {
			continue
		}
		sort.Strings(got)
		sort.Strings(tt.want)
		if !equalStringSlices(got, tt.want) {
			t.Errorf("pattern %q:\ngot:  %q\nwant: %q", tt.pattern, got, tt.want)
		}
	}
}
