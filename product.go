package stockx

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
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
	Media       Media
	ObjectID    string
	BelowRetail bool
	Variants    []Variant `mapstructure:"omitempty"`
}

type Variant struct {
	ShoeSize string
	UUID     string
	Market   *Market `mapstructure:"omitempty"`
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

// SearchProductsOptions specifies the optional parameters to the
// Products method.
type SearchProductsOptions struct {
	Name string `url:"_search,omitempty"`
	PaginationOptions
}

// Search provides a list of products.
// It takes optionnal SearchProductsOptions.
//
// The result set is paged, iteration procedure is command with PaginationOptions
// Without opts the method will return first X items returned by stockX API.
func (s *ProductService) Search(ctx context.Context, opts *SearchProductsOptions) (p *Products, err error) {
	body, err := s.client.Request(ctx, URIStockxSearch, opts)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// GetProductOptions specifies the optional parameters to the
// Get method.
//
// Without opts the method will return basic product info and variants size.
type GetProductOptions struct {
	Includes string `url:"includes,omitempty"`
	Currency string `url:"currency,omitempty"`
}

// Get a single product.
// Allows passing product ID, product UUID, product URL key to fetch details about a singe product.
func (s *ProductService) Get(ctx context.Context, id string, opts *GetProductOptions) (p *Product, err error) {
	body, err := s.client.Request(ctx, URIStockxProduct+id, opts)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Error parsing JSON string - %s", err)
	}

	pmap := data["Product"]
	p = &Product{}
	err = mapstructure.Decode(pmap, p)
	if err != nil {
		return nil, err
	}

	childmap, ok := pmap.(map[string]interface{})["children"].(map[string]interface{})
	if ok {
		for _, child := range childmap {
			v := &Variant{}
			err = mapstructure.Decode(child, v)
			if err != nil {
				return nil, err
			}
			marketmap, ok := child.(map[string]interface{})["market"].(map[string]interface{})
			if ok {
				if len(marketmap) != 0 {
					m := &Market{}
					err = mapstructure.Decode(marketmap, m)
					if err != nil {
						return nil, err
					}
					v.Market = m
				} else {
					v.Market = nil
				}
			}
			p.Variants = append(p.Variants, *v)
		}
	}

	return p, nil
}
