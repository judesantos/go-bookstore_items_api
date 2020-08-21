package services

// items_service.go

import (
	"github.com/judesantos/go-bookstore_items_api/domain/items"
	"github.com/judesantos/go-bookstore_utils/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.IRestError)
	Get(string) (*items.Item, rest_errors.IRestError)
}

type itemsService struct{}

//
// Create - Create items
//
func (s *itemsService) Create(
	item items.Item,
) (*items.Item, rest_errors.IRestError) {

	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil
}

//
// Get - Get items
//
func (s *itemsService) Get(string) (*items.Item, rest_errors.IRestError) {
	return nil, nil
}
