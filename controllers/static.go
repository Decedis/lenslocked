package controllers

import (
	"net/http"

	"github.com/decedis/lenslocked/views"
)

// where our logic for handlling rendering logic lives
func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
