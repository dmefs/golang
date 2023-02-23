package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, nil
	}
	defer f.Close()

	n, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", 0, nil
	}

	// Close file, but prefer error from Copy, if any.
	if err = f.Close(); err != nil {
		return "", 0, err
	}
	return local, n, err
}
