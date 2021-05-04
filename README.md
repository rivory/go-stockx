## go-stockx

go-stockx is a Unoficial Go client library for accessing the StockX API.

## Usage :

Example : 

``` go
  client, _ := stockx.NewClient("User Agent")

	opts := &stockx.ProductsOptions{
		Name: "yeezy",
	}

	p, err := client.Products(opts)
```