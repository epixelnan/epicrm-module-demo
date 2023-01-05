// This file is part of Epixel CRM Software

package app

import (
	"net/http"
	"github.com/epixelnan/libepicrm/apiparts"
)

func (app App) GetHello(w http.ResponseWriter, r *http.Request) {
	epicrm_apiparts.SendJsonResponse(w, "world")
}
