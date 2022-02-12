package middleware

import (
	"context"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/sakuraapp/shared/pkg/model"
	"github.com/sakuraapp/shared/pkg/resource"
	"net/http"
)

const UserIdCtxKey = "user_id"

func sendUnauthorized(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, resource.ErrUnauthorized)
}

func UserIdFromContext(ctx context.Context) model.UserId {
	userId, _ := ctx.Value(UserIdCtxKey).(model.UserId)

	return userId
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqCtx := r.Context()
		token, _, err := jwtauth.FromContext(reqCtx)

		if err != nil {
			sendUnauthorized(w, r)
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			sendUnauthorized(w, r)
			return
		}

		rawId, _ := token.Get("id")
		floatId, _ := rawId.(float64)
		id := model.UserId(floatId)

		if id == 0 {
			sendUnauthorized(w, r)
			return
		}

		ctx := context.WithValue(reqCtx, UserIdCtxKey, id)
		r = r.WithContext(ctx)

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}