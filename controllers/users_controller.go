package controllers

import (
	"fmt"
	"net/http"

	"github.com/judesantos/go-bookstore_items_api/domain/items"
	"github.com/judesantos/go-bookstore_items_api/services"
	"github.com/judesantos/go-bookstore_oauth/oauth"
)

var (
	UsersController usersControllerInterface = &usersController{}
)

type usersControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type usersController struct{}

func (c *usersController) Create(w http.ResponseWriter, r *http.Request) {

	if err := oauth.AuthenticateRequest(r); err != nil {
		return
	}

	item := items.Item{
		Seller: oauth.GetUserId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		return
	}

	fmt.Println(result)
}

func (c *usersController) Get(w http.ResponseWriter, r *http.Request) {

}
