// To create an executable, it looks like the name of the package needs to be main?
package main

// This is a comment
// To compile, you can simply use: go build
// In such case, the name of the executable corresponds to the name of the folder that contains the GO code.
// If you plan to have multiple executables, then you can instead do: go build <name of go file>.
// For instance, go build simple.go (in such cases, the name of the executable has the same name of the file you compiled).

import (
	"fmt"
)

// Define a struct named SimpleAsset
type SimpleAsset struct {
}

// The starting point for any Go program is the main function.
func main() {
	fmt.Println("Hello World!")
	test()
	// See how you can create a new instance of the SimpleAsset struct
	asset := new(SimpleAsset)
	asset.test()

	//switch cases
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	str := "ricardo"
	fmt.Print("Write ", str, " as ")
	switch str {
	case "ricardo":
		fmt.Println("r")
	case "jen":
		fmt.Println("j")
	case "katy":
		fmt.Println("k")
	}

}

func (t *SimpleAsset) test() {
	fmt.Println("In test() as a function of SimpleAsset...")
}

func test() {
	fmt.Println("In test()")
}
