package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/needl3/goreact-template/internal/core"
)

// Strict auth check and denial if not authenticated
func AuthMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie("Authorization")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		jwtToken := strings.Split(authCookie.Value, "Bearer ")[1]
		token, err := jwt.ParseWithClaims(jwtToken, &core.JwtType{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("APP_SECRET")), nil
		})
		claim, ok := token.Claims.(*core.JwtType)
		if err != nil || !token.Valid || !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		r.SetPathValue("userId", claim.AuthenticatedUserRequestValues.Id)

		next.ServeHTTP(w, r)
	})
}
