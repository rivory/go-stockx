## go-stockx

go-stockx is an unofficial Go client library for accessing the StockX API.

## Usage :

Example : 

Search for items
``` go
  client, _ := stockx.NewClient("User Agent")

	opts := &stockx.ProductsOptions{
		Name: "yeezy",
	}

	p, err := client.Products.Search(ctx, opts)
```


Get item details
``` go
	client, _ := stockx.NewClient("User Agent")

	p, err := client.Products.Get(ctx, "185ecb6f-2402-467c-8db4-c846bf8cdb7a")
```