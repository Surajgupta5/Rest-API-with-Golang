package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"restapiwithgo/models"
)

var (
	file *os.File
	err  error
)

type response struct {
	//ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func create(name string) {
	if _, err := os.Stat(name); err == nil {
		fmt.Printf("File exists\n")
		return

	}

	file, err = os.Create(name)

	if err != nil {
		panic(err)
	}
	log.Println("File Created")
	//	log.Println(file)
	//file.Close()
}
func Createfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var f models.File
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	create(f.Name)
	res := response{
		Message: "File created successfully",
	}
	json.NewEncoder(w).Encode(res)

}

func rename(name, rname string) {
	err = os.Rename(name, rname)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File Renamed")
}
func Renamefile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	

	var f models.File
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	rename(f.Name, f.Rname)
	res := response{
		Message: "File renamed successfully",
	}
	json.NewEncoder(w).Encode(res)

}

func delete(name string) {
	err = os.Remove(name)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s Deleted", name)
}

func Deletefile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	

	var f models.File
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	delete(f.Name)
	res := response{
		Message: "File deleted successfully",
	}
	json.NewEncoder(w).Encode(res)

}
func Get(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// check if it is a regular file (not dir)
		if info.Mode().IsRegular() {
			fmt.Println("file name:", info.Name())
		}
		return nil
	})
}
