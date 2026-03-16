package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec" 
	"path/filepath"
	"time"
)

const imageDir = "./images/"

func Capture(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02T15-04-05")
	filename := currentTime + ".jpg"
	fullPath := filepath.Join(imageDir, filename)

	cmd := exec.Command("rpicam-still", "--tuning-file", "/usr/share/libcamera/ipa/rpi/vc4/imx219_noir.json", "--width", "1024", "--height", "768", "--vflip", "--hflip", "-o", fullPath)
	output, err := cmd.CombinedOutput()

	if err != nil {
		http.Error(w, fmt.Sprintf("Command failed: %s\n%s", err, output), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filename)
}

func ListFiles(w http.ResponseWriter, r *http.Request) {
    err := os.MkdirAll(imageDir, 0755)
    if err != nil {
        http.Error(w, "Unable to create directory: "+err.Error(), http.StatusInternalServerError)
        return
    }

    entries, err := os.ReadDir(imageDir)
    if err != nil {
        http.Error(w, "Unable to read files: "+err.Error(), http.StatusInternalServerError)
        return
    }

    var fileNames []string
    for _, entry := range entries {
        if !entry.IsDir() {
            fileNames = append(fileNames, entry.Name())
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(fileNames)
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	filename := r.PathValue("filename")

	safePath := filepath.Join(imageDir, filepath.Base(filename))

	if _, err := os.Stat(safePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, safePath)
}
