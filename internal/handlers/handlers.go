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

	bytes, err := os.ReadFile("../index.html")
	if err != nil {
		http.Error(res, fmt.Sprintf("An error occurred while opening html form"), http.StatusInternalServerError)
		return
	}

	res.Write(bytes)
}

func HandlerUpload(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(10 << 20)

	//get the file
	file, _, err := req.FormFile("myFile")
	if err != nil {
		http.Error(w, "An error occurred while getting file from html form", http.StatusInternalServerError)
		return
	}

	// read the file
	text, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "An error occurred while reading data from the file", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	// convert data to morse or backwards
	data, err := service.Determine(string(text))

	// get name for the file
	nameForFile := time.Now().Format("20060102_1504") + ".txt"

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
	w.Write([]byte(data))

}
