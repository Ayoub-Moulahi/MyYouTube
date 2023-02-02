package middelware

import (
	"context"
	"github.com/Ayoub-Moulahi/MyYouTube/models"
	"github.com/Ayoub-Moulahi/MyYouTube/request"
	"net/http"
)

type MiddUser struct {
	models.UserInterface
}

// RequireUser used to require a user to be  logged in order to be
// able to visit a given page
func (ms *MiddUser) RequireUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//check if a request holds a cookie
		//if not redirect to login page
		cookie, _ := r.Cookie("token")
		if cookie == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		//get the user to attach it to the request
		//if no user found redirect again to the login page
		user, _ := ms.GetUserByRemember(context.Background(), cookie.Value)
		if user == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		//get the context from the request
		ctx := r.Context()
		//add user to context
		ctx = request.AddUserToContext(ctx, user)
		// change the request context with the new formed one
		r = r.WithContext(ctx) //
		next(w, r)
	})
}
