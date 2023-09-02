package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalesb16/tkoifood/internal/handlers"
)

type API struct {
	server   *http.Server
	sHandler handlers.StoreHandler
	oHandler handlers.OrderHandler
}

func New(sHandler handlers.StoreHandler, oHandler handlers.OrderHandler) *API {
	a := &API{
		sHandler: sHandler,
		oHandler: oHandler,
	}

	a.server = &http.Server{
		Addr:    ":8080",
		Handler: a.router(),
	}

	return a
}

func (a *API) router() *gin.Engine {
	r := gin.Default()

	sGroup := r.Group("/store")
	oGroup := r.Group("/order")

	sGroup.GET("/:store_id", a.sHandler.Get)
	sGroup.POST("/", a.sHandler.Create)

	oGroup.POST("/", a.oHandler.Create)

	return r
}

func (a *API) Run() error {
	return a.server.ListenAndServe()
}
