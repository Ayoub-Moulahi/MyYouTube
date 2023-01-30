package main

import (
	"fmt"
	"github.com/Ayoub-Moulahi/MyYouTube/controllers"
	"github.com/Ayoub-Moulahi/MyYouTube/models"
	"github.com/Ayoub-Moulahi/MyYouTube/setting"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	config, err := setting.LoadConfig(".") //relative path of app.env file to main.go
	if err != nil {
		fmt.Println(err.Error())
	}
	dialect := config.Dialect
	connInfo := config.ConnInfo
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
	//user route
	r.Handle("/login", uc.LogIN).Methods("GET")
	r.Handle("/index", uc.IndexP).Methods("GET")
	r.Handle("/signin", uc.SignIN).Methods("GET")
	r.HandleFunc("/signin", uc.SingIN_POST).Methods("POST")
	r.HandleFunc("/login", uc.LogIN_POST).Methods("POST")
	//third party Oauth
	r.HandleFunc("/auth/{provider}", uc.BeginAuth)
	r.HandleFunc("/auth/{provider}/callback", uc.CompleteAuth)

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

	log.Fatal(http.ListenAndServe(config.AdressPort, r))
}
