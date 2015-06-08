package handler

import (
	"net/http"

	_render "gopkg.in/unrolled/render.v1"
)

var render *_render.Render

type Book struct {
	Title  string
	Author string
}

func Home(rw http.ResponseWriter, r *http.Request) {
	if render == nil {
		render = _render.New()
	}

	book := Book{"Hello Go", "Hemanta Sapkota"}
	render.HTML(rw, http.StatusOK, "home", book)
}
