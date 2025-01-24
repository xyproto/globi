package globi

import (
	"os"
	"path/filepath"
	"strings"
)

func Glob(pattern string) ([]string, error) {
	dir, base := filepath.Split(pattern)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	matches := make([]string, 0, len(entries))

	if !strings.Contains(base, "*") {
		for _, e := range entries {
			if strings.EqualFold(e.Name(), base) {
				matches = append(matches, filepath.Join(dir, e.Name()))
			}
		}
		return matches, nil
	}

	prefix, suffix, _ := strings.Cut(base, "*")
	lPrefix := strings.ToLower(prefix)
	lSuffix := strings.ToLower(suffix)

	for _, e := range entries {
		name := e.Name()
		lName := strings.ToLower(name)

		if !strings.HasPrefix(lName, lPrefix) {
			continue
		}

		if suffix == "" || strings.HasSuffix(lName, lSuffix) {
			matches = append(matches, filepath.Join(dir, name))
		}
	}

	return matches, nil
}
