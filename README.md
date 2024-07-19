# A Go client for ICCU API

https://api.iccu.sbn.it/devportal/apis

_Work in progress, do not use._

Example usage:

```
package main

import (
	"fmt"

	"github.com/atomotic/iccu/client"
	"github.com/atomotic/iccu/nomi"
)

func main() {

	client.New("$key", "$secret")

	names := nomi.Search("alei* crowley")

	for _, name := range names.Docs {
		fmt.Printf("%s - %s - %s\n", name.ID, name.Bid(), name.Name())
	}

}
```

```
IT\ICCU\CFIV\025223 - http://id.sbn.it/bid/CFIV025223 - Crowley, Aleister
```
