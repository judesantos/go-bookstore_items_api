package elasticsearch

// es_client.go

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/judesantos/go-bookstore_utils/logger"
	"github.com/olivere/elastic"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
}

type esClient struct {
	client *elastic.Client
}

//
// Init - initialize elastic client
//
func Init() {

	var err error

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC: ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		//elastic.SetHeaders(http.Header{
		//	"X-User-Id": []string{"..."},
		//}),
	)

	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

// setClient
func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

//
// Index
//
func (c *esClient) Index(
	index string,
	docType string,
	doc interface{},
) (*elastic.IndexResponse, error) {

	ctx := context.Background()

	res, err := c.client.Index().Index(index).Type(docType).BodyJson(doc).Do(ctx)

	if err != nil {
		logger.Error("index document error", err)
		return nil, err
	}

	return res, nil
}

//
// Get
//
func (c *esClient) Get(
	index string,
	docType string,
	id string,
) (*elastic.GetResult, error) {

	ctx := context.Background()

	result, err := c.client.Get().
		Index(index).
		Type(docType).
		Id(id).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("id %d not found", id), err)
		return nil, err
	}

	return result, nil
}

//
// search
//
func (c *esClient) Search(
	index string,
	query elastic.Query,
) (*elastic.SearchResult, error) {

	ctx := context.Background()
	result, err := c.client.Search(index).
		Query(query).
		RestTotalHitsAsInt(true).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("search failed for index %s", index), err)
		return nil, err
	}

	return result, nil
}
