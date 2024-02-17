package helpers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func Upload(w http.ResponseWriter, file multipart.File, handler *multipart.FileHeader) {
	timestamp := time.Now().Format("20060102_150405")

	filename := fmt.Sprintf("src/uploads/%s_%s", timestamp, handler.Filename)

	out, err := os.Create(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
