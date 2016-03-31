package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

var staticAssets = assetMap()

func handleStatic(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path

	if file[0] == '/' {
		file = file[1:]
	}

	if len(file) == 0 {
		file = "index.html"
	}

	bs, ok := staticAssets[file]
	if !ok {
		http.NotFound(w, r)
		return
	}

	mtype := mimeTypeForFile(file)
	if len(mtype) != 0 {
		w.Header().Set("Content-Type", mtype)
	}
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
	} else {
		// ungzip if browser didn't indicate it understands gzip
		var gr *gzip.Reader
		gr, _ = gzip.NewReader(bytes.NewReader(bs))
		bs, _ = ioutil.ReadAll(gr)
		gr.Close()
	}
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(bs)))
	w.Header().Set("Cache-Control", "public")

	w.Write(bs)
}

func mimeTypeForFile(file string) string {
	switch path.Ext(file) {
	case ".html":
		return "text/html"
	case ".js":
		return "application/javascript"
	default:
		return "application/octet-stream"
	}
}
