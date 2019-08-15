package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	Templates *template.Template
	GlobalErr error
)

func HandleError(err error, em, sm string) {
	if err != nil {
		log.Println(em)
		log.Fatalln(err)
	} else if sm != "" {
		log.Println(sm)
	}
}

func init() {
	Templates, GlobalErr = template.ParseGlob("templates/*")
	HandleError(GlobalErr, "Error parsing glob templates", "Parsed glob templates")

	err := godotenv.Load()
	HandleError(err, "Err loading env", "")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)

	log.Fatalln(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
}

func Home(w http.ResponseWriter, r *http.Request) {

	err := Templates.ExecuteTemplate(w, "home.html", nil)
	HandleError(err, "Err exec home.html", "")

}
