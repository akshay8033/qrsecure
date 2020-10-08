package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers qrsecure-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/qrsecure/Product", createProductHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/qrsecure/Product", listProductHandler(cliCtx, "qrsecure")).Methods("GET")
		r.HandleFunc("/qrsecure/Product/{key}", getProductHandler(cliCtx, "qrsecure")).Methods("GET")
		r.HandleFunc("/qrsecure/Product", setProductHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/qrsecure/Product", deleteProductHandler(cliCtx)).Methods("DELETE")

		
}
