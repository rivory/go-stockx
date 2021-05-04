package stockx

// Client is a StockX client to request various of resources on stockX website.
type Client struct {
	parser ParserItf
}

const (
	URIStockxSearch = "https://stockx.com/api/browse"
)

// PaginationOptions specifies the optional parameters for various methods
// that supports pagination.
type PaginationOptions struct {
	ResultsPerPage int `url:"resultsPerPage,omitempty"`
	Page           int `url:"page,omitempty"`
}

// NewClient creates a new StockX client.
func NewClient(userAgent string) (c *Client, err error) {
	return &Client{
		parser: NewParser(userAgent),
	}, nil
}
