package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalesb16/koifood/internal/handlers"
)

type API struct {
	server   *http.Server
	sHandler handlers.StoreHandler
}

func New(sHandler handlers.StoreHandler) *API {
	a := &API{
		sHandler: sHandler,
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

	sGroup.GET("/:id")
	sGroup.POST("/", a.sHandler.Create)

	oGroup.POST("/")

	return r
}

func (a *API) Run() error {
	return a.server.ListenAndServe()
}
