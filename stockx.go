package stockx

import (
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

// Client is a StockX client to request various of resources on stockX website.
type Client struct {
	httpClient *http.Client
	userAgent  string
	common     service

	Products *ProductService
}

type service struct {
	client *Client
}

const (
	URIStockxSearch  = "https://stockx.com/api/browse"
	URIStockxProduct = "https://stockx.com/api/products/"
)

// PaginationOptions specifies the optional parameters for various methods
// that supports pagination.
type PaginationOptions struct {
	ResultsPerPage int `url:"resultsPerPage,omitempty"`
	Page           int `url:"page,omitempty"`
}

// NewClient creates a new StockX client.
func NewClient(userAgent string) (c *Client, err error) {
	httpClient := &http.Client{}

	c = &Client{
		httpClient: httpClient,
		userAgent:  userAgent,
	}
	c.common.client = c
	c.Products = (*ProductService)(&c.common)

	return c, nil
}

// Request creates an API request.
func (c *Client) Request(uri string, opts interface{}) (b []byte, err error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = queryOptions(opts)

	req.Header.Add("user-agent", c.userAgent)
	req.Header.Add("accept-language", "en-US")
	req.Header.Add("sec-fetch-dest", "none")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "cross-site")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// addOptions encodes options as query params.
func queryOptions(opts interface{}) string {
	qs, _ := query.Values(opts)

	return qs.Encode()
}
