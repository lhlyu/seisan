package api

import (
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
)

func HandlerVersion(w http.ResponseWriter, r *http.Request) {
	ss := make([]string, 0)
	filepath.WalkDir("../", func(path string, d fs.DirEntry, err error) error {
		ss = append(ss, path)
		return nil
	})
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, `{"version":"%s"}`, strings.Join(ss, ","))
}
