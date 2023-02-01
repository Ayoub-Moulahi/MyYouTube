package controllers

import (
	"context"
	"fmt"
	"github.com/Ayoub-Moulahi/MyYouTube/models"
	"github.com/Ayoub-Moulahi/MyYouTube/token"
	"github.com/Ayoub-Moulahi/MyYouTube/views"
	"github.com/gorilla/schema"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

type UserController struct {
	SignIN *views.View
	LogIN  *views.View
	IndexP *views.View
	us     *models.Services
}

var ctx = context.Background()

// UserSg holds the parsed fields of the  sign in form
type UserSg struct {
	Name      string `schema:"Name,required"`
	Email     string `schema:"email,required"`
	Birthdate string `schema:"date,required"`
	Password  string `schema:"password,required"`
}

// NewUserController initialize a User
func NewUserController(us *models.Services) *UserController {
	signIn, _ := views.NewView("layout", "views/users/signin.gohtml")
	logIn, _ := views.NewView("layout", "views/users/login.gohtml")
	index, _ := views.NewView("layout", "views/users/index_page.gohtml")

	return &UserController{
		signIn,
		logIn,
		index,
		us,
	}
}

// SingIN_POST used create a user when a user sign in
func (uc *UserController) SingIN_POST(w http.ResponseWriter, r *http.Request) {
	var usg UserSg
	//TODO fix this

	err := parseForm(r, &usg)
	if err != nil {
		fmt.Println(err)
		uc.SignIN.RenderView(w, r, views.CreateAlert(err))
		return
	}
	date, _ := time.Parse("2023-01-28", usg.Birthdate)
	tmp := models.User{
		Username:  usg.Name,
		Email:     usg.Email,
		Birthdate: date,
		Password:  usg.Password,
	}
	_, err = uc.us.UserInter.CreateUser(ctx, tmp)
	if err != nil {
		fmt.Println(err.Error())
		uc.SignIN.RenderView(w, r, views.CreateAlert(err))
		return
	}
	err = uc.createSetCookies(w, &tmp)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/index", http.StatusFound)

}

// Userlg holds the parsed fields of the  login form
type Userlg struct {
	Email    string `schema:"email,required"`
	Password string `schema:"pwd,required"`
}

func (uc *UserController) LogIN_POST(w http.ResponseWriter, r *http.Request) {
	var ulg Userlg
	err := parseForm(r, &ulg)
	if err != nil {
		uc.LogIN.RenderView(w, r, views.CreateAlert(err))
	}
	tmp, err := uc.us.UserInter.Authenticate(ulg.Email, ulg.Password)
	if err != nil {
		switch err {
		case models.ErrNoAccount:
			fmt.Println(models.ErrNoAccount)
		default:
			fmt.Println(err.Error())

		}
		uc.LogIN.RenderView(w, r, views.CreateAlert(err))
		return
	}
	err = uc.createSetCookies(w, tmp)
	if err != nil {
		fmt.Println(err.Error())
		uc.LogIN.RenderView(w, r, views.CreateAlert(err))
		return
	}

	http.Redirect(w, r, "/index", http.StatusFound)

}

// createSetCookies is a helper method used to set cookies anytime a user sign in or log in
func (uc *UserController) createSetCookies(w http.ResponseWriter, u *models.User) error {
	if u.Remember == "" {
		t, err := token.GenerateToken(token.RememberTokenBytes)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		u.Remember = t

		err = uc.us.UserInter.UpdateUserRemember(ctx, u.ID, u.Remember)
		if err != nil {
			fmt.Println(err.Error())
			return err

		}
	}
	cookie := http.Cookie{
		Name:     "token",
		Value:    u.Remember,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	return nil
}

// parseForm a helper function used to parse form
func parseForm(r *http.Request, destination interface{}) error {
	var decoder = schema.NewDecoder()
	if err := r.ParseForm(); err != nil {

		return err
	}
	err := decoder.Decode(destination, r.PostForm)
	if err != nil {
		return err
	}
	return nil
}
