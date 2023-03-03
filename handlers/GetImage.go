package handlers

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gorilla/mux"
)

type GetImage struct {
	StoragePath string
}

func (i GetImage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slug, found := mux.Vars(r)["slug"]
	if !found || slug == "" {
		http.Error(w, "No image ID provided", http.StatusBadRequest)
		return
	}

	mimeType := mime.TypeByExtension(filepath.Ext(slug))
	imgPath := path.Join(i.StoragePath, slug)
	img, err := os.Open(imgPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Image is not found with error %s", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", mimeType)
	if _, err = io.Copy(w, img); err != nil {
		http.Error(w, "fail to copy file", http.StatusInternalServerError)
		return
	}
}
