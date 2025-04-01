package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandlerCore(res http.ResponseWriter, req *http.Request) {
	filePath := "../index.html"

	_, err := os.Stat(filePath)
	if err != nil {
		res.Header().Set("Content-Type", "text/plain; charset=utf-8")
		http.Error(res, fmt.Sprintf("An error occurred while accessing the file: %v", err), http.StatusInternalServerError)
		return
	}

	http.ServeFile(res, req, filePath)
}

func HandlerUpload(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(10 << 20)

	//get the file
	file, _, err := req.FormFile("myFile")
	if err != nil {
		http.Error(w, "An error occurred while getting file from html form", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	// read the file
	text, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "An error occurred while reading data from the file", http.StatusInternalServerError)
		return
	}

	// convert data to morse or backwards
	data, err := service.DetectEncoding(string(text))

	// get name for the file
	//nameForFile := time.Now().Format("20060102_1504") + ".txt"   //мой вариант был лучше
	nameForFile := time.Now().UTC().String() //на винде нельзя имя файла с ":", программа упадёт если её там запустить

	//create the file
	f, err := os.Create(nameForFile)
	if err != nil {
		http.Error(w, "An error occurred while creating local file", http.StatusInternalServerError)
	}
	defer f.Close()

	// write to the file
	_, err = fmt.Fprintln(f, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("An error occurred while writing to file: %v", err), http.StatusInternalServerError)
		return
	}

	// console check
	fmt.Println(data)

	// write data into /upload page
	_, err = w.Write([]byte(data))
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		http.Error(w, fmt.Sprintf("An error occurred while sending response to client: %v", err), http.StatusInternalServerError)
		return
	}

}
