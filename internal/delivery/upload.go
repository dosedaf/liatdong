package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	_ "time"

	"github.com/google/uuid"
)

const uploadPath = "./uploads" // Directory to store files

type UploadResponse struct {
	Success bool   `json:"success"`
	Link    string `json:"link"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file from request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate a unique file ID
	fileID := uuid.New().String()
	filePath := filepath.Join(uploadPath, fileID)

	// Save file to disk
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	// Respond with the shareable link
	link := fmt.Sprintf("http://localhost:8080/files/%s", fileID)
	response := UploadResponse{Success: true, Link: link}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
