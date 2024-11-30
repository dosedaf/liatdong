package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	fileID := r.URL.Path[len("/files/"):] // Extract the file ID from the URL
	filePath := filepath.Join(uploadPath, fileID)

	file, err := os.Open(filePath)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer file.Close()

	http.ServeFile(w, r, filePath)
}
