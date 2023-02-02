package request

import (
	"context"
	"github.com/Ayoub-Moulahi/MyYouTube/models"
)

// AddUserToContext puts the request user into the current context.
func AddUserToContext(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, "user", user)
}

// GetUserFromContext returns the request user from the context.
// A nil value is returned if there are no user in the
// current context.
func GetUserFromContext(ctx context.Context) *models.User {
	val := ctx.Value("user")
	if x, ok := val.(*models.User); ok {
		return x
	}
	return nil
}

//package context
//
//import (
//	"context"
//
//	"github.com/Ayoub-Moulahi/MyYouTube/models"
//)
//
//const (
//	userKey privateKey = "user"
//)
//
//type privateKey string
//
//func WithUser(ctx context.Context, user *models.User) context.Context {
//	return context.WithValue(ctx, "user", user)
//}
//
//func User(ctx context.Context) *models.User {
//	if temp := ctx.Value("user"); temp != nil {
//		if user, ok := temp.(*models.User); ok {
//			return user
//		}
//	}
//	return nil
//}
