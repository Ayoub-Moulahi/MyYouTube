package controllers

import (
	"fmt"
	"github.com/Ayoub-Moulahi/MyYouTube/models"
	"github.com/Ayoub-Moulahi/MyYouTube/setting"
	"github.com/Ayoub-Moulahi/MyYouTube/views"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/twitterv2"
	"net/http"
	"time"
)

func InitCall() {
	config, err := setting.LoadConfig("../")
	if err != nil {
		fmt.Println(err.Error())
	}
	GoogleClientId := config.GoogleId
	GoogleClientSecret := config.GoogleSecret
	FacebookClientId := config.FacebookId
	FacebookClientSecret := config.FacebookSecret
	TwitterClientId := config.TwitterId
	TwitterClientSecret := config.TwitterSecret

	key := config.Sessionkey
	maxAge := config.MaxAge
	isProd := config.IsProd

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = config.SessionPath
	store.Options.HttpOnly = config.SessionHttpOnly
	store.Options.Secure = isProd
	gothic.Store = store
	goth.UseProviders(
		google.New(GoogleClientId, GoogleClientSecret, "http://localhost:8080/auth/google/callback", "email", "profile"),
		facebook.New(FacebookClientId, FacebookClientSecret, "http://localhost:8080/auth/facebook/callback", "email", "profile"),
		twitterv2.NewAuthenticate(TwitterClientId, TwitterClientSecret, "http://localhost:8080/auth/twitter/callback"),
	)

}
func (uc *UserController) BeginAuth(w http.ResponseWriter, r *http.Request) {
	InitCall()
	gothic.BeginAuthHandler(w, r)
}
func (uc *UserController) CompleteAuth(w http.ResponseWriter, r *http.Request) {
	InitCall()
	config, err := setting.LoadConfig("../")
	if err != nil {
		fmt.Println(models.ErrApp)
	}
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Println(err)
		uc.LogIN.RenderView(w, r, views.CreateAlert(err))
	}
	_, err = uc.us.UserInter.Authenticate(user.Email, config.DefaultPwd)
	if err == models.ErrNoAccount {
		_, serr := uc.us.UserInter.CreateUser(ctx, models.User{

			Username:     user.Name,
			Email:        user.Email,
			Birthdate:    time.Time{},
			Password:     config.DefaultPwd,
			RememberHash: user.AccessToken,
		})
		if serr != nil {
			fmt.Println(serr.Error())
			uc.LogIN.RenderView(w, r, views.CreateAlert(err))
			return
		}
	} else if err != nil {
		fmt.Println(err.Error())
		uc.LogIN.RenderView(w, r, views.CreateAlert(err))
		return
	}

	http.Redirect(w, r, "/index", http.StatusFound)
}
