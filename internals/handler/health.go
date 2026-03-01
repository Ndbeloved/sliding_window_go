package handler

import (
	"net/http"

	"github.com/Ndbeloved/rate-limiter-go/pkg/response"
)

func Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
