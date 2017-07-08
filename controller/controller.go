package controller

import (
	"net/http"

	"github.com/spkills/spkills/model"
	"github.com/spkills/spkills/view"
)

func RootController(w http.ResponseWriter, r *http.Request) {
	res := model.RootModel()
	view.RootView(w, res)
}
