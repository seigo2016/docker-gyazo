package main

import (
	"net/http"
	"html/template"
	"log"
	"io/ioutil"
        "path"
	"os"
)

type Path struct {
	Link string
}

func getImages() ([]Path){
	dir, _ := os.Getwd()
	log.Println(dir + "/data/")
	files, _ := ioutil.ReadDir(dir + "/data/")
	var imgs []Path
	log.Println(files)
	for _, file := range files {
		if !file.IsDir() && path.Ext(file.Name())==".png" {
			link := pathToLink(file.Name())
			imgs = append(imgs, Path{Link: link})
		}
	}
	log.Println(imgs)
	return imgs
}

func pathToLink(path string) (string){
	lb := "https://showgyazo.seigo2016.com/data/"
	return lb + path
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("view.tmpl")
	imgs := getImages()
	err = tmpl.ExecuteTemplate(w, "view.tmpl", imgs)
	if err != nil {
	    log.Printf("[ERROR] execute template error: %s", err)
	    return
        }
	return
}

func main() {
	//cgi.Serve(http.HandlerFunc(handler))
	http.HandleFunc("/", viewHandler)
        err := http.ListenAndServe(":80", nil)
	if err != nil {
	    log.Fatal("ListenAndServe: ", err)
	}
}

//func handler(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case http.MethodGet:
//		viewHandler(w, r)
//	default:
//		w.WriteHeader(http.StatusMethodNotAllowed)
//	}
//}

