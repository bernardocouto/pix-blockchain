package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type QrCodeEstatico struct {
	ChaveEnderecamento string  `json:"chave_enderecamento"`
	Valor              float64 `json:"valor"`
}

type QueryResult struct {
	Key    string `json:"key"`
	Record *QrCodeEstatico
}

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) ChangeQrEstatico(ctx contractapi.TransactionContextInterface, qrCodeEstaticoNumber string, newValor float64) error {
	qrCodeEstatico, err := s.QueryQrCodeEstatico(ctx, qrCodeEstaticoNumber)
	if err != nil {
		return err
	}
	qrCodeEstatico.Valor = newValor
	qrCodeEstaticoAsBytes, _ := json.Marshal(qrCodeEstatico)
	return ctx.GetStub().PutState(qrCodeEstaticoNumber, qrCodeEstaticoAsBytes)
}

func (s *SmartContract) CreateQrCodeEstatico(ctx contractapi.TransactionContextInterface, qrCodeEstaticoNumber string, chaveEnderecamento string, valor float64) error {
	qrCodeEstatico := QrCodeEstatico{
		ChaveEnderecamento: chaveEnderecamento,
		Valor:              valor,
	}
	qrCodeEstaticoAsBytes, _ := json.Marshal(qrCodeEstatico)
	return ctx.GetStub().PutState(qrCodeEstaticoNumber, qrCodeEstaticoAsBytes)
}

func (s *SmartContract) QueryQrCodeEstatico(ctx contractapi.TransactionContextInterface, qrCodeEstaticoNumber string) (*QrCodeEstatico, error) {
	qrCodeEstaticoAsBytes, err := ctx.GetStub().GetState(qrCodeEstaticoNumber)
	if err != nil {
		return nil, fmt.Errorf("Falha ao realizar leitura. %s", err.Error())
	}
	if qrCodeEstaticoAsBytes == nil {
		return nil, fmt.Errorf("%s não existe.", qrCodeEstaticoNumber)
	}
	qrCodeEstatico := new(QrCodeEstatico)
	_ = json.Unmarshal(qrCodeEstaticoAsBytes, qrCodeEstatico)
	return qrCodeEstatico, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Erro ao criar QR Code Estático: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Erro ao iniciar QR Code Estático: %s", err.Error())
	}
}
