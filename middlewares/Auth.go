package middlewares

import (
	"net/http"

	"github.com/mishrabhi0123/screen-time/database"
	"github.com/mishrabhi0123/screen-time/utils"
)

type AuthMiddleware struct {
	next func(http.ResponseWriter, *http.Request)
}

func (m *AuthMiddleware) Run(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-auth-token")
	if token == "" {
		res := utils.CreateResponse(false, "Please provide the authentication token", 400, nil)
		utils.SendResponse(w, res)
		return
	}

	claims, err := utils.ParseToken(token)

	if err != nil {
		res := utils.CreateResponse(false, "Unauthorized action", 401, nil)
		utils.SendResponse(w, res)
		return
	}

	user := database.GetUser(claims.Email)
	device := database.GetDevice(claims.UserId, claims.DeviceName)

	if user.Id == 0 || device.Id == 0 {
		res := utils.CreateResponse(false, "Unidentified user or device", 404, nil)
		utils.SendResponse(w, res)
		return
	}

	m.next(w, r)
}

func WithAuth(handlerFunction func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	wrappedFunction := &AuthMiddleware{handlerFunction}
	return wrappedFunction.Run
}
