package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", server)
	http.HandleFunc("/test", test)
	http.ListenAndServe(":3000", nil)
}

func server(res http.ResponseWriter, req *http.Request) {

	//Generic Header
	res.Header().Set("Server", "A Go Web Server")

	switch req.Method {
	case "GET":
		contentType := req.Header.Get("Content-type")

		fmt.Printf(contentType)

		if contentType == "text/plain" {

			res.Header().Set("Content-Type", "text/plain")
			res.WriteHeader(200)
			res.Write([]byte(contentType + " request!"))

		} else if contentType == "application/json" {

			type Profile struct {
				Name    string
				Hobbies []string
			}

			profile := Profile{"Pau", []string{"programming", "nature"}}

			js, err := json.Marshal(profile)

			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(200)
			res.Write(js)

		} else if contentType == "application/xml" {

			type Profile struct {
				Name    string
				Hobbies []string `xml:"Hobbies>Hobby"`
			}

			profile := Profile{"Pau", []string{"programming", "nature"}}

			x, err := xml.MarshalIndent(profile, "", "  ")
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/xml")
			res.WriteHeader(200)
			res.Write(x)

		} else if contentType == "image/gif" {

			// Assuming you want to serve a photo at 'images/foo.png'
			fp := path.Join("images", "cat.gif")
			http.ServeFile(res, req, fp)

			res.Header().Set("Content-Type", "image/gif")
			res.WriteHeader(200)

		} else if contentType == "text/html" {

			type Profile struct {
				Name    string
				Hobbies []string
			}

			profile := Profile{"Pau", []string{"Programming", "Nature"}}

			fp := path.Join("templates", "index.html")
			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := tmpl.Execute(res, profile); err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
			}

		} else {
			//Only Headers
			res.Header().Set("Server", "A Go Web Server")
			res.WriteHeader(200)
		}
	case "POST":

	case "PUT":

	case "PATCH":

	case "DELETE":

	default:

	}

}

func test(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Server", "A Go Web Server")
	res.Header().Set("Content-Type", "application/text")
	res.WriteHeader(200)
	res.Write([]byte("TEST REQUEST"))
}
