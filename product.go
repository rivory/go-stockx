package stockx

import (
	"context"
	"encoding/json"
)

type ProductService service

type Products struct {
	Pagination Pagination
	Products   []Product
}

type Pagination struct {
	Limit        string
	Page         int
	Total        int
	LastPage     string
	Sort         []string
	Order        []string
	CurrentPage  *string
	NextPage     *string
	PreviousPage *string
}

type Product struct {
	ID          string
	UUID        string
	Brand       string
	Category    string
	Colorway    string
	Condition   string
	Description string
	Gender      string
	Name        string
	// ReleaseDate time.Time //2017-02-11 Custom unmarshal
	Retailprice float64
	Shoe        string
	Title       string
	URLKey      string
	// Market Market ToDo:
	Media       Media
	ObjectID    string
	BelowRetail bool
}

type Variant struct {
	ShoeSize string
	UUID     string
	Market   Market `mapstructure:"omitempty"`
}

type ProductDetail struct {
	Product
	Variants []Variant
}

type Market struct {
	SkuUUID                   string
	ProductUUID               string
	LowestAsk                 int
	LowestAskSize             string
	NumberfAsks               int
	HasAsks                   int
	SalesThisPeriod           int
	SalesLastPeriod           int
	HighestBid                int
	HighestBidSize            string
	NumberOfBids              int
	HasBids                   int
	AnnualHigh                int
	AnnualLow                 int
	DeadstockRangeLow         int
	DeadstockRangeHigh        int
	Volality                  float64
	DeadstockSold             int
	AverageDeadstockPrice     int
	LastSale                  int
	LastSaleSize              string
	SalesLast72Hours          int
	AverageDeadstockPriceRank int
}

type Media struct {
	Gallery360    []string `json:"360"`
	ImageURL      string
	SmallImageURL string
	ThumbURL      string
	Has360        bool
	Gallery       []string
}

// ProductsOptions specifies the optional parameters to the
// Products method.
type ProductsOptions struct {
	Name string `url:"_search,omitempty"`
	PaginationOptions
}

// Search provides a list of products.
// It takes optionnal ProductsOptions.
func (s *ProductService) Search(ctx context.Context, opts *ProductsOptions) (p *Products, err error) {
	body, err := s.client.Request(URIStockxSearch, opts)
	if err != nil {
		return nil, err
	}

	// spew.Dump(string(body))
	// return nil, nil
	products := Products{}
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}

	return &products, nil
}

// Get a single product.
// Allows passing product ID, product UUID, product URL key to fetch details about a singe product.
func (s *ProductService) Get(ctx context.Context, id string) (p *ProductDetail, err error) {
	// body, err := s.client.Request(URIStockxProduct+id+"?includes=market&currency=EUR", nil)
	// if err != nil {
	// 	return nil, err
	// }

	return p, nil
}
