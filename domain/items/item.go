package items

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
