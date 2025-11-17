# A Go client for ICCU API

https://api.iccu.sbn.it/devportal/apis

_Work in progress, do not use._

Example usage:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/atomotic/iccu/client"
	"github.com/atomotic/iccu/nomi"
)

func main() {
	ctx := context.Background()

	c, err := client.New(ctx, "$key", "$secret")
	if err != nil {
		log.Fatal(err)
	}

	// Search returns an iterator that handles pagination automatically
	for doc := range nomi.Search(ctx, c, "alei* crowley", nil) {
		fmt.Printf("%s - %s - %s\n", doc.ID, doc.Bid(), doc.Name())
	}
}
```

Output:
```
IT\ICCU\CFIV\025223 - http://id.sbn.it/bid/CFIV025223 - Crowley, Aleister
```

With custom page size and total count:

```go
// Get total count of results
iter, total, err := nomi.SearchWithTotal(ctx, c, "dante", nil)
if err != nil {
	log.Fatal(err)
}

count := 0
for doc := range iter {
	fmt.Printf("%s - %s\n", doc.ID, doc.Name())
	count++
}
fmt.Printf("Fetched %d out of %d total results\n", count, *total)

// Custom page size
opts := &nomi.SearchOptions{PageSize: 100}
for doc := range nomi.Search(ctx, c, "dante", opts) {
	fmt.Println(doc.Name())
}
```
