package routers

import "github.com/gorilla/mux"

func RegisteredRouters(router *mux.Router) {
	UserRouter(router)
	WalletRouter(router)
}