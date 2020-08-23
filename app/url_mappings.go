package app

import (
	"net/http"

	"github.com/judesantos/go-bookstore_items_api/controllers"
)

func mapUrls() {

	// create item

	router.HandleFunc(
		"/items",
		controllers.ItemsController.Create,
	).Methods(http.MethodPost)

	// get item by id

	router.HandleFunc(
		"/items/{id}",
		controllers.ItemsController.Get,
	).Methods(http.MethodGet)

	// search item

	router.HandleFunc(
		"/items/search",
		controllers.ItemsController.Search,
	).Methods(http.MethodPost)
}
