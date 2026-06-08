package middleware

import (
	"context"
	"lotcastick-backend/internal/dto"
	"lotcastick-backend/internal/util"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler, JwtsecretKey string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			util.WriteJSON(w, http.StatusUnauthorized, dto.MessageResponse{
				Status:  false,
				Message: "Authrorization header is required",
			})
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			util.WriteJSON(w, http.StatusUnauthorized, dto.MessageResponse{
				Status:  false,
				Message: "Invalid Authorization header format",
			})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		acTknClaims, err := util.ParseAccessToken(token, JwtsecretKey)
		if err != nil {
			util.WriteJSON(w, http.StatusUnauthorized, dto.MessageResponse{
				Status:  false,
				Message: "Invalid or expired token",
			})
			return
		}

		// Attach the user ID to the request context
		ctx := context.WithValue(r.Context(), "userID", acTknClaims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
