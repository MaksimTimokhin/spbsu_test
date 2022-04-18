package main

import (
	"log"
	"spbsu_test/onchain/chaincode"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	cc, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating chaincode: %v", err)
	}

	if err := cc.Start(); err != nil {
		log.Panicf("Error starting chaincode: %v", err)
	}
}
