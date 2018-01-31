package main

// To build/compile this chaincode, you need to first fetch the corresponding dependencies and then compile:
// go get -u --tags nopkcs11 github.com/hyperledger/fabric/core/chaincode/shim
// go build --tags nopkcs11 mycc.go
// See following for more on go get: https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies
 
import "fmt"
import "github.com/hyperledger/fabric/core/chaincode/shim"
import "github.com/hyperledger/fabric/protos/peer"
 
type SampleChaincode struct {
}

// Init method is called when the chaincode is first deployed to the network... executed by each peer that deploys its own instance of the chaincode.
func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
    fmt.Println("In Init() method...")
    args := stub.GetStringArgs()
    if len(args) != 2 {
        return shim.Error("Incorrect arguments. Expecting a key and a value!")
    }

    // Set up any variables or assets here by calling stub.PutState()
    // We store the key and the value on the ledger
    err := stub.PutState(args[0], []byte(args[1]))
    if err != nil {
        return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
    }

    return shim.Success(nil)
}

// Query method is invoked wheneber any read/get operation is performed on the blockchain state. This method does not
// run within a transactional context since it does not change the state of the blockchain. Invocations of this method are 
// not recorded in the blockchain.
// func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
//     return nil, nil
// }

// Invoke method is invoked whenever the state of the blockchain is to be modified: all create, update, delete operations should
// be encapsulated withim this method. All invocations of this method are recorded in the blockchain as transactions. 
func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    fmt.Println("In Invoke() method...")
    // Extract the function and args from the transaction proposal
    fn, args := stub.GetFunctionAndParameters()

    var result string
    var err error
    if fn == "set" {
            result, err = set(stub, args)
    } else {
            result, err = get(stub, args)
    }
    if err != nil {
            return shim.Error(err.Error())
    }

    // Return the result as success payload
    return shim.Success([]byte(result))
}

// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    fmt.Println("In set() method...")
    if len(args) != 2 {
            return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
    }

    err := stub.PutState(args[0], []byte(args[1]))
    if err != nil {
            return "", fmt.Errorf("Failed to set asset: %s", args[0])
    }
    return args[1], nil
}

// Get returns the value of the specified asset key
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    fmt.Println("In get() method...")
    if len(args) != 1 {
            return "", fmt.Errorf("Incorrect arguments. Expecting a key")
    }

    value, err := stub.GetState(args[0])
    if err != nil {
            return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
    }
    if value == nil {
            return "", fmt.Errorf("Asset not found: %s", args[0])
    }
    return string(value), nil
}
 
func main() {
    // Register the chaincode with the peer
    // Think of the chaincode as the entity (in this case, SampleChaincode) that conforms to the interface that implements Init, Query, and Invoke...
    // The implementation of these methods is required by the shim.Chaincode interface.
    err := shim.Start(new(SampleChaincode))
    if err != nil {
        fmt.Println("Could not start SampleChaincode!")
    } else {
        fmt.Println("SampleChaincode successfully started.")
    } 
}