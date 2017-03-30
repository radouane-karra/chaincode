package main

import (
	"errors"
	"fmt"
  "strconv"
  "time"
  "encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	tableName = "balances"
)

type BalancesInfo struct {
	accountId            string `json:"account_id"`
	value                int64  `json:"value"`
	Time                 int64  `json:"time"`
}

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
  var err error

  if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

  _, err = stub.GetTable(tableName)

  if err == shim.ErrTableNotFound {
  		err = stub.CreateTable(tableName, []*shim.ColumnDefinition{
    			&shim.ColumnDefinition{Name: "accountId", Type: shim.ColumnDefinition_STRING, Key: true},
    			&shim.ColumnDefinition{Name: "value", Type: shim.ColumnDefinition_UINT64, Key: false},
    			&shim.ColumnDefinition{Name: "Time", Type: shim.ColumnDefinition_INT64, Key: false},
  		})
  		if err != nil {
  			fmt.Println("Error creating table:%s", err.Error())
  			return nil, errors.New("Failed creating AccountInfo table.")
  		}
	} else {
		fmt.Println("Table already exists")
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
	}else if function == "readAccounts" { //read a variable
		return t.readAccount(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// write - invoke function to write key/value pair
func (t *PotCommun) addAccount(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  var account string    // Entities
  var value int64 // Asset holdings
  var val int
  var err error

  if len(args) != 2 {
    return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the account and value to set")
	}

  // Initialize the chaincode
	account = args[0]
	val, err = strconv.Atoi(args[1])
  value = time.Now().UnixNano()
	if err != nil {
      return nil, errors.New("Expecting integer value for asset holding")
	}

  fmt.Printf("Account = %s, value = %d\n", account, value)

  now := time.Now().UnixNano()

  ok, err := stub.InsertRow(tableName, shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: account}},
			&shim.Column{Value: &shim.Column_Int64{Int64: value}},
			&shim.Column{Value: &shim.Column_Int64{Int64: now}},
		},
  })

  if !ok {
     fmt.Printf("Row for account :%s and value: %s already exists, just updating the entry", account, val)
  }

	// Write the state to the ledger
	//err = stub.PutState(account, []byte(strconv.Itoa(value)))
	if err != nil {
    return nil, err
	}

  return nil, nil
}

// read - query function to read key/value pair
func (t *PotCommun) readAccount(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var account, jsonResp string
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

// Return all attestation records
func (t *PotCommun) readAccounts(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("In attestationRecords function")
	if len(args) > 0 {
		fmt.Printf("Incorrect number of arguments")
		return nil, errors.New("Incorrect number of arguments. No arguments required")
	}

	var columns []shim.Column

	rowChannel, err := stub.GetRows(tableName, columns)
	if err != nil {
		fmt.Printf("Error in getting rows:%s", err.Error())
		return nil, errors.New("Error in fetching rows")
	}
	balancesRecords := make([]BalancesInfo, 0)
	for row := range rowChannel {
		balanceRecord := BalancesInfo{
			accountId:           row.Columns[0].GetString_(),
			value:               row.Columns[1].GetInt64(),
			Time:                row.Columns[2].GetInt64(),
		}
		balancesRecords = append(balancesRecords, balanceRecord)
	}

	payload, err := json.Marshal(balancesRecords)
	if err != nil {
		fmt.Printf("Failed marshalling payload")
		return nil, fmt.Errorf("Failed marshalling payload [%s]", err)
	}

	return payload, nil
}
