package items

import (
	"github.com/judesantos/go-bookstore_items_api/clients/elasticsearch"
	"github.com/judesantos/go-bookstore_utils/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() rest_errors.IRestError {

	res, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.InternalServerError("failed to save document", err)
	}

	i.Id = res.Id

	return nil
}
