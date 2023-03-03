package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/sogilis/support-formation-docker/templates"
)

type PostImage struct {
	StoragePath string
}

func (i PostImage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := uuid.New().String()

	rawImage, imageFileHandler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, fmt.Sprintf("Image not found in body - %s", err), http.StatusBadRequest)
		return
	}

	ext := filepath.Ext(imageFileHandler.Filename)
	imgName := fmt.Sprintf("%s%s", id, ext)

	f, err := os.Create(path.Join(i.StoragePath, imgName))
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to create image - %s", err), http.StatusConflict)
		return
	}
	defer f.Close()

	if _, err = io.Copy(f, rawImage); err != nil {
		http.Error(w, fmt.Sprintf("Unable to copy image - %s", err), http.StatusInternalServerError)
		return
	}

	t, err := template.New("uploaded").Parse(templates.UploadedPage)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to load uploaded template - %s", err), http.StatusInternalServerError)
		return
	}
	if err = t.ExecuteTemplate(w, "uploaded", fmt.Sprintf("/i/%s", imgName)); err != nil {
		http.Error(w, fmt.Sprintf("Unable execute uploaded template %s", err), http.StatusInternalServerError)
		return
	}

}
