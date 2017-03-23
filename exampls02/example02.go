package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
  pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("remote_attestation")

const (
	tableName = "UserAccounts"
)

type UserAccountInfo struct {
	UserID            string `json:"user_id"`
	AccountID         string `json:"account_id"`
	Status            uint64 `json:"status"`
}

// RemoteDeviceAttestation implementation. This smart contract enables multiple attestors
// to perform remote attestation of device and verify that the device is running authentic
// and valid software/hardware configuration
type PotCommun struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("PotCommun Init")

  var err error

	if len(args) != 0 {
		logger.Error("Incorrect number of arguments")
    return shim.Error("Incorrect number of arguments. No arguments required for deploying this contract.")
	}

	_, err = stub.GetTable(tableName)
	if err == shim.ErrTableNotFound {
		err = stub.CreateTable(tableName, []*shim.ColumnDefinition{
			&shim.ColumnDefinition{Name: "UserID", Type: shim.ColumnDefinition_STRING, Key: true},
			&shim.ColumnDefinition{Name: "AccountID", Type: shim.ColumnDefinition_STRING, Key: true},
			&shim.ColumnDefinition{Name: "Status", Type: shim.ColumnDefinition_UINT64, Key: false},
		})
		if err != nil {
			logger.Errorf("Error creating table:%s", err.Error())
			return shim.Error("Failed creating DeviceAttestation table.")
		}
	} else {
		logger.Info("Table already exists")
	}

	logger.Info("Successfully deployed chain code")

	return shim.Success(nil)
}
