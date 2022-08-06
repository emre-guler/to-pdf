package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/exp/slices"
)

var tpl = template.Must(template.ParseFiles("./view/main.gohtml"))

const (
	UPLOAD_FOLDER = "./files"
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
			tempFile, err := ioutil.TempFile(UPLOAD_FOLDER, "*.docx")
			if err != nil {
				fmt.Println(err)
			}
			defer tempFile.Close()
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}
			tempFile.Write(fileBytes)

			convertFile((UPLOAD_FOLDER + "/" + handler.Filename))

			newFilePath := strings.Replace(UPLOAD_FOLDER+"/"+handler.Filename, ".docx", "", 1)
			pdfFile, _ := ioutil.ReadFile(newFilePath)
			w.Header().Add("Content-Type", "application/octet-stream")
			w.Write(pdfFile)
		} else {
			tpl.Execute(w, nil)
		}
	}
}

func convertFile(inputFile string) {
	exec.Command("libreoffice --headless --convert-to pdf --outdir " + UPLOAD_FOLDER + inputFile).Output()
}
