package controllers

import (
	"github.com/Ayoub-Moulahi/MyYouTube/views"
)

type StaticController struct {
	HomePage    *views.View
	ContactPage *views.View
}

func NewStaticController() *StaticController {
	home, _ := views.NewView("layout", "views/static/home.gohtml")
	contact, _ := views.NewView("layout", "views/static/contact.gohtml")
	return &StaticController{
		home,
		contact,
	}
}
