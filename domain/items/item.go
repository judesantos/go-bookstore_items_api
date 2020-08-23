package items

import (
	"github.com/judesantos/go-bookstore_items_api/domain/queries"
	"github.com/judesantos/go-bookstore_utils/rest_errors"
)

type IItem interface {
	Save() rest_errors.IRestError
	Get() rest_errors.IRestError
	Search(query queries.EsQuery) ([]Item, rest_errors.IRestError)
}

type Item struct {
	Id          string
	Seller      int64
	Title       string
	Description Description
	Pictures    []Picture
	Video       string
	Price       float32
	Available   int
	Sold        int
	Status      string
}

type Description struct {
	PlainText string `json:"plain_text"`
	Html      string
}

type Picture struct {
	Id  int64
	Url string
}
