module github.com/judesantos/go-bookstore_items_api

go 1.14

require (
	github.com/federicoleon/golang-restclient v0.0.0-20191104170228-162ed620df66 // indirect
	github.com/fortytw2/leaktest v1.3.0 // indirect
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/gorilla/mux v1.7.4
	github.com/judesantos/go-bookstore_oauth v0.0.0-20200821013207-20d534e062be
	github.com/judesantos/go-bookstore_utils v0.0.0-20200821013116-f3a26ad513fc
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/olivere/elastic v6.2.34+incompatible
	github.com/pkg/errors v0.9.1 // indirect
	go.uber.org/zap v1.15.0 // indirect
)

replace github.com/judesantos/go-bookstore_utils => /home/judesantos/dev/go/projects/bookstore/go-bookstore_utils

replace github.com/judesantos/go-bookstore_oauth => /home/judesantos/dev/go/projects/bookstore/go-bookstore_oauth
