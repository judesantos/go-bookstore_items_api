package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/judesantos/go-bookstore_items_api/domain/items"
	"github.com/judesantos/go-bookstore_items_api/services"
	"github.com/judesantos/go-bookstore_items_api/utils/http_utils"
	"github.com/judesantos/go-bookstore_oauth/oauth"
	"github.com/judesantos/go-bookstore_utils/rest_errors"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

//
// Create - Create index controller
//
func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {

	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.JsonErrorResponse(w, err)
		return
	}

	fmt.Println("Headers:", r.Header)
	userId := oauth.GetUserId(r)
	if userId == 0 {
		rerr := rest_errors.UnauthorizedError(
			"Invalid request params - user id required")
		http_utils.JsonErrorResponse(w, rerr)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rerr := rest_errors.BadRequestError("invalid request body")
		http_utils.JsonErrorResponse(w, rerr)
		return
	}

	defer r.Body.Close()

	var item items.Item

	if err := json.Unmarshal(reqBody, &item); err != nil {
		rerr := rest_errors.BadRequestError("invalid json body")
		http_utils.JsonErrorResponse(w, rerr)
		return
	}

	item.Seller = userId

	result, rerr := services.ItemsService.Create(item)
	if rerr != nil {
		http_utils.JsonErrorResponse(w, rerr)
		return
	}

	http_utils.JsonSuccessResponse(w, http.StatusCreated, result)
}

//
// Get - Get index controller
//
func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
