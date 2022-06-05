package delivery

import (
	"github.com/julienschmidt/httprouter"
)

// API defines the methods available for API delivery
type API interface {
	RegisterRoutes(router *httprouter.Router)
}