// package main provides math solution
package main

import (
	"fmt"

	mypackage "github.com/AbrarMirje/mypackage/mypackagethree"
)

func main() {
	fmt.Println(mypackage.Sum(10, 20, 30))
	fmt.Println(mypackage.Sum(6, 8))
	fmt.Println(mypackage.Sum(3, 8, 9, 2))

}
