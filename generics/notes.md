## Notes

### Difference between %v and %+v
```go
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 30}

	fmt.Printf("Using %%v: %v\n", p)
	fmt.Printf("Using %%+v: %+v\n", p)
}
```
The + modifier includes the fields names of a struct that you are printing
```
Using %v: {Alice 30}
Using %+v: {Name:Alice Age:30}
```



