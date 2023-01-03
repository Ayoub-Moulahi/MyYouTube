package main

import (
	"github.com/Ayoub-Moulahi/MyYouTube/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	stc := controllers.NewStaticController()
	r := mux.NewRouter()
	r.Handle("/", stc.HomePage)
	r.Handle("/contact", stc.ContactPage)
	// Assets
	//serving  JS
	assetHandler := http.FileServer(http.Dir("./javascript/"))
	assetHandler = http.StripPrefix("/javascript/", assetHandler)
	r.PathPrefix("/javascript").Handler(assetHandler)

	//serving CSS
	cssHandler := http.FileServer(http.Dir("./css"))
	cssHandler = http.StripPrefix("/css/", cssHandler)
	r.PathPrefix("/css").Handler(cssHandler)

	//serving videos
	videoHandler := http.FileServer(http.Dir("./videos"))
	videoHandler = http.StripPrefix("/videos/", videoHandler)
	r.PathPrefix("/videos").Handler(videoHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
