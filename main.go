package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/exp/slices"
)

var tpl = template.Must(template.ParseFiles("./view/main.gohtml"))

const (
	UPLOAD_FOLDER = "/tmp"
)

var ALLOWED_EXTENSIONS = []string{".doc", ".docx"}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9091"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandeler)
	http.ListenAndServe((":" + port), mux)
}

func indexHandeler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tpl.Execute(w, nil)
	case "POST":
		r.ParseForm()
		file, handler, err := r.FormFile("docFile")
		if err != nil {
			fmt.Println("Error retrieving the file")
			fmt.Println(err)
			return
		}
		defer file.Close()

		if slices.Contains(ALLOWED_EXTENSIONS, filepath.Ext(handler.Filename)) {
			tempFile, err := ioutil.TempFile("./files", "*.docx")
			if err != nil {
				fmt.Println(err)
			}
			defer tempFile.Close()
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}
			tempFile.Write(fileBytes)

			convertBytes, _ := ioutil.ReadFile(tempFile.Name())
			responseBytes := convertFile(convertBytes)
			w.Header().Add("Content-Type", "application/octet-stream")
			w.Header().Add("Content-Disposition", "attachment; filename=docFile.docx")
			w.Write(responseBytes)
			return
		} else {
			tpl.Execute(w, nil)
		}
	}
}

func convertFile(bytes []byte) []byte {
	return []byte{}
}
