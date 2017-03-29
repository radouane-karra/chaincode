package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type PotCommun struct {
}

func main() {
	err := shim.Start(new(PotCommun))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *PotCommun) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

  if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *PotCommun) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "addAccount" {
		return t.addAccount(stub, args)
	} 

	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *PotCommun) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "readAccount" { //read a variable
		return t.readAccount(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// write - invoke function to write key/value pair
func (t *PotCommun) addAccount(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  var account string    // Entities
  var value int // Asset holdings
  var err error

  if len(args) != 2 {
    return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the account and value to set")
	}

  // Initialize the chaincode
	account = args[0]
	value, err = strconv.Atoi(args[1])
	if err != nil {
      return nil, errors.New("Expecting integer value for asset holding")
	}

  fmt.Printf("Account = %s, value = %d\n", account, value)

	// Write the state to the ledger
	err = stub.PutState(account, []byte(strconv.Itoa(value)))
	if err != nil {
		return shim.Error(err.Error())
	}

  return nil, nil
}

// read - query function to read key/value pair
func (t *PotCommun) readAccount(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the account number to read account")
	}

	account = args[0]
	valueAsbytes, err := stub.GetState(account)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + account + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valueAsbytes, nil
}
