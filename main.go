package main

// This is a comment
// To compile, you can simply use: go build
// In such case, the name of the executable corresponds to the name of the folder that contains the GO code.
// If you plan to have multiple executables, then you can instead do: go build <name of go file>.
// For instance, go build main.go (in such cases, the name of the executable has the same name of the file you compile).

import (
    "fmt"
   // "github.com/hyperledger/fabric/core/chaincode/shim"
   // "github.com/hyperledger/fabric/protos/peer"
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
}

func (t *SimpleAsset) test() {
	fmt.Println("In test() as a function of SimpleAsset...")
}

func test() {
	fmt.Println("In test()")
}

// func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {

// }