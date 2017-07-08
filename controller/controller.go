package controller

import (
	"net/http"

	"spkills/model"
	"spkills/view"
)

func RootController(w http.ResponseWriter, r *http.Request) {
	res := model.RootModel()
	view.RootView(w, res)
}
