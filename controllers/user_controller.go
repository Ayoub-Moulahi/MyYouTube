package controllers

import (
	"context"
	"fmt"
	"github.com/Ayoub-Moulahi/MyYouTube/models"
	"github.com/Ayoub-Moulahi/MyYouTube/views"
	"github.com/gorilla/schema"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

type UserController struct {
	SignIN *views.View
	LogIN  *views.View
	us     *models.Services
}

var decoder = schema.NewDecoder()
var ctx = context.Background()

// User holds the parsed fields of the  sign in form
type User struct {
	Name      string `schema:"Name,required"`
	Email     string `schema:"email,required"`
	Birthdate string `schema:"date,required"`
	Password  string `schema:"password,required"`
}

// NewUserController initialize a User
func NewUserController(us *models.Services) *UserController {
	sign_in, _ := views.NewView("layout", "views/users/signin.gohtml")
	log_in, _ := views.NewView("layout", "views/users/login.gohtml")

	return &UserController{
		sign_in,
		log_in,
		us,
	}
}

func (uc *UserController) SingIN_POST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
	}
	var u User
	err = decoder.Decode(&u, r.PostForm)
	if err != nil {
		fmt.Println(err.Error())
	}
	date, _ := time.Parse("2006-01-02", u.Birthdate)
	tmp, err := uc.us.UserInter.CreateUser(ctx, models.User{

		Username:  u.Name,
		Email:     u.Email,
		Birthdate: date,
		Password:  u.Password,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(tmp)
}
