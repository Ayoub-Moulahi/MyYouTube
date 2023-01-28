package main

import (
	"fmt"
	"github.com/Ayoub-Moulahi/MyYouTube/controllers"
	"github.com/Ayoub-Moulahi/MyYouTube/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	dialect := "postgres"
	connInfo := "postgresql://ayoub:myPwd@localhost:5432/youtube?sslmode=disable"
	//ctx := context.Background()
	services, err := models.NewService(dialect, connInfo)
	if err != nil {
		fmt.Println(err.Error())
	}

	stc := controllers.NewStaticController()
	uc := controllers.NewUserController(services)
	r := mux.NewRouter()
	//static page route
	r.Handle("/", stc.HomePage).Methods("GET")
	r.Handle("/contact", stc.ContactPage).Methods("GET")
	r.Handle("/login", uc.LogIN).Methods("GET")
	r.Handle("/signin", uc.SignIN).Methods("GET")
	r.HandleFunc("/signin", uc.SingIN_POST).Methods("POST")

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
