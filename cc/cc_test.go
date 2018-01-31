package main

// Reference: https://medium.com/wearetheledger/about-smart-contract-testing-fbf7b576bb9f

import (
	"fmt"
    "testing"
	"github.com/stretchr/testify/assert"
	// "encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestXYZ (t *testing.T) {
    var a string = "Hello"
    var b string = "Hello"  
    assert.Equal(t, a, b, "The two words should be the same.")
}

func TestCreateNewUser(t *testing.T) {
	// ARRANGE
	simpleAsset := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleAsset)
	fmt.Println(mockStub)
	// testStub := TestAPIStub{data: make(map[string][]byte)}
	txId := "mockTxID"
	internalId := "00001"
	firstName := "John"

	// args := []string {
	// 	internalId,
	// 	firstName,
	// 	"Doe",
	// 	"john.doek@somewhere.com",
	// }

	args := []string {
		internalId,
		firstName,
	}

	 // ACT
	 mockStub.MockTransactionStart(txId)
	 response := simpleAsset.createUser(mockStub, args)
	 mockStub.MockTransactionEnd(txId)

	// ASSERT
	if s := response.GetStatus(); s != 200 {
	   t.Errorf("The response status is %d, instead of 200", s)
	   t.Errorf("The message: %s", response.Message)
	}

	something := mockStub.State[internalId]
	fmt.Println("something")
	fmt.Println(something)
	fmt.Println("something")

	s := string(something[:])
	fmt.Println(s)

	//TODO
	// var user User
	// var found bool
	// json.Unmarshal(mockStub.State[internalId], &user)
	// // Now you can assert properties of the user object you stored.
	// // ...
 }