package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type QrCodeDinamico struct {
	Ispb          string  `json:"ispb"`
	TipoConta     string  `json:"tipo_conta"`
	Agencia       string  `json:"agencia"`
	Conta         string  `json:"conta"`
	Link          string  `json:"link"`
	Identificador string  `json:"identificador"`
	Valor         float64 `json:"valor"`
}

type QueryResult struct {
	Key    string `json:"key"`
	Record *QrCodeDinamico
}

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) ChangeQrDinamico(ctx contractapi.TransactionContextInterface, qrCodeDinamicoNumber string, newValor float64) error {
	qrCodeDinamico, err := s.QueryQrCodeDinamico(ctx, qrCodeDinamicoNumber)
	if err != nil {
		return err
	}
	qrCodeDinamico.Valor = newValor
	qrCodeDinamicoAsBytes, _ := json.Marshal(qrCodeDinamico)
	return ctx.GetStub().PutState(qrCodeDinamicoNumber, qrCodeDinamicoAsBytes)
}

func (s *SmartContract) CreateQrCodeDinamico(ctx contractapi.TransactionContextInterface, qrCodeDinamicoNumber string, ispb string, tipoConta string, agencia string, conta string, link string, identificador string, valor float64) error {
	qrCodeDinamico := QrCodeDinamico{
		Ispb:          ispb,
		TipoConta:     tipoConta,
		Agencia:       agencia,
		Conta:         conta,
		Link:          link,
		Identificador: identificador,
		Valor:         valor,
	}
	qrCodeDinamicoAsBytes, _ := json.Marshal(qrCodeDinamico)
	return ctx.GetStub().PutState(qrCodeDinamicoNumber, qrCodeDinamicoAsBytes)
}

func (s *SmartContract) QueryQrCodeDinamico(ctx contractapi.TransactionContextInterface, qrCodeDinamicoNumber string) (*QrCodeDinamico, error) {
	qrCodeDinamicoAsBytes, err := ctx.GetStub().GetState(qrCodeDinamicoNumber)
	if err != nil {
		return nil, fmt.Errorf("Falha ao realizar leitura. %s", err.Error())
	}
	if qrCodeDinamicoAsBytes == nil {
		return nil, fmt.Errorf("%s não existe.", qrCodeDinamicoNumber)
	}
	qrCodeDinamico := new(QrCodeDinamico)
	_ = json.Unmarshal(qrCodeDinamicoAsBytes, qrCodeDinamico)
	return qrCodeDinamico, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Erro ao criar QR Code Dinâmico: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Erro ao iniciar QR Code Dinâmico: %s", err.Error())
	}
}
