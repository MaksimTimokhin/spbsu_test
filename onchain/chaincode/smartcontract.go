package chaincode

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) AddRecord(ctx contractapi.TransactionContextInterface, id, record string) error {
	return ctx.GetStub().PutState(id, []byte(record))
}

func (s *SmartContract) GetAllRecords(ctx contractapi.TransactionContextInterface) ([]string, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []string
	for resultsIterator.HasNext() {
		rsp, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		records = append(records, string(rsp.Value))
	}
	return records, nil
}
