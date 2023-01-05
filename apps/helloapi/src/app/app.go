// This file is part of Epixel CRM Software

package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"

	"github.com/epixelnan/libepicrm/apiparts"
)

const SERVICE = "epicmod-hello-helloapi"

type App struct {
}

func NewApp() *App {
	return &App{}
}

func NewAppHandler() http.Handler {
	myapp := NewApp()

	spec, err := GetSwagger()
	if err != nil {
		epicrm_apiparts.LogSpecErrorAndQuit(err)
	}

	epicrm_apiparts.HandleSpec(spec)
	
	r := chi.NewRouter()
	
	epicrm_apiparts.AddCommonMiddlewares(r, SERVICE)
	r.Use(middleware.OapiRequestValidatorWithOptions(spec, nil))
	
	return HandlerFromMux(myapp, r)
}
