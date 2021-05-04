package stockx

import (
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

type ParserItf interface {
	Request(uri string, opts interface{}) (b []byte, err error)
}

type parser struct {
	userAgent string
	client    *http.Client
}

func NewParser(userAgent string) ParserItf {
	return &parser{
		userAgent: userAgent,
		client:    &http.Client{},
	}
}

// Request creates an API request.
func (p *parser) Request(uri string, opts interface{}) (b []byte, err error) {
	req, err := http.NewRequest("GET", URIStockxSearch, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = queryOptions(opts)

	req.Header.Add("user-agent", p.userAgent)
	req.Header.Add("accept-language", "en-US")
	req.Header.Add("sec-fetch-dest", "none")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "cross-site")
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// addOptions encodes options as query params
func queryOptions(opts interface{}) string {
	qs, _ := query.Values(opts)

	return qs.Encode()
}
