package stockx

import "encoding/json"

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
	// ReleaseDate time.Time
	Retailprice float64
	Shoe        string
	Title       string
	URLKey      string
	// Market Market ToDo:
}

// ProductsOptions specifies the optional parameters to the
// Products method.
type ProductsOptions struct {
	Name string `url:"_search,omitempty"`
	PaginationOptions
}

// Products provides a list of products.
// It takes optionnal ProductsOptions.
func (c *Client) Products(opts *ProductsOptions) (p *Products, err error) {
	body, err := c.parser.Request(URIStockxSearch, opts)
	if err != nil {
		return nil, err
	}

	products := Products{}
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}

	return &products, nil
}
