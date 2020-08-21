package elasticsearch

import (
	"context"
	"time"

	"github.com/judesantos/go-bookstore_utils/logger"
	"github.com/olivere/elastic"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
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
		elastic.SetURL("http://127.0.0.1:9200"), elastic.SetHealthcheckInterval(10*time.Second),
		//elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC: ", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
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
	doc interface{},
) (*elastic.IndexResponse, error) {

	ctx := context.Background()

	res, err := c.client.Index().Index("Items").BodyJson(doc).Do(ctx)

	if err != nil {
		logger.Error("index document error", err)
		return nil, err
	}

	return res, nil
}
