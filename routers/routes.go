package routers

import (
	"github.com/gorilla/mux"
	// "github.com/megamsquare/go_gateway/handler"
)

func RegisteredRouters(router *mux.Router) {
	UserRouter(router)
	WalletRouter(router)
}