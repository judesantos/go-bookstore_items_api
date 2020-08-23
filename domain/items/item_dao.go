package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/judesantos/go-bookstore_items_api/clients/elasticsearch"
	"github.com/judesantos/go-bookstore_items_api/domain/queries"
	"github.com/judesantos/go-bookstore_utils/rest_errors"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

//
// Save
//
func (i *Item) Save() rest_errors.IRestError {

	res, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.InternalServerError("failed to save document", err)
	}

	i.Id = res.Id

	return nil
}

//
// Get
//
func (i *Item) Get() rest_errors.IRestError {

	itemId := i.Id

	res, err := elasticsearch.Client.Get(indexItems, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NotFoundError(
				fmt.Sprintf("no entry found for id %d", i.Id))
		}
		return rest_errors.InternalServerError(
			fmt.Sprintf("get id %d error", i.Id),
			err,
		)
	}

	bytes, err := res.Source.MarshalJSON()
	if err != nil {
		return rest_errors.InternalServerError(
			"error reading db response", err)
	}

	if err := json.Unmarshal(bytes, i); err != nil {
		return rest_errors.InternalServerError(
			"error reading db response", err)
	}

	i.Id = itemId

	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, rest_errors.IRestError) {

	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, rest_errors.InternalServerError("search documents error",
			errors.New("db Error"))
	}

	items := make([]Item, result.TotalHits())
	for idx, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.InternalServerError(
				"parse response error",
				errors.New("db Error"))
		}
		items[idx] = item
	}

	if 0 == len(items) {
		return nil, rest_errors.NotFoundError("no entries found")
	}

	return items, nil
}
