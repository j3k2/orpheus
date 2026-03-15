package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"os/exec"
	"time"
)

func CaptureHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	cmd := exec.Command("rpicam-still", "--tuning-file", "/usr/share/libcamera/ipa/rpi/vc4/imx219_noir.json", "--width", "1024", "--height", "768", "--vflip", "--hflip", "-o", "./static/" + currentTime + ".jpg")
	output, err := cmd.CombinedOutput()

    if err != nil {
        http.Error(w, fmt.Sprintf("Command failed: %s\n%s", err, output), http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
	// json.NewEncoder(w).Encode(currentTime)
}