package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
	//  "byte"

	"github.com/google/uuid"
	// "github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a car
var invokeMSPID ="Org3MSP" 

type SmartContract struct {
	contractapi.Contract
}
type serverConfig struct {
	CCID    string
	Address string
}
type QueryBar struct {
	Key    string `json:"Key"`
	Record *Bar
}
type Bar struct {
	BarID              string `json:"id"`
	BarLocation        string `json:"barLocation"`
	BarSerialNumber    string `json:"barSerialNumber"`
	Purity             string `json:"purity"`
	BarRefiner         string `json:"barRefiner"`
	BarHallmarkVerfied string `json:"barHallmarkVerfied"`
	BarWeightInGms     string `json:"barWeightInGms"`
	DateCreated        string `json:"dateCreated"`
}

type QueryBuy struct {
	Key    string `json:"Key"`
	Record *Buy
}
type Buy struct {
	DSGId          string `json:"id"`
	OrderId        string `json:"orderId"`
	Amount         string `json:"amount"`
	AmountWithFees string `json:"amountWithFees"`
	Stage          string `json:"stage"`
	PaymentStatus  string `json:"paymentStatus"`
	EstimatedGrams string `json:"estimatedGrams"`
	UserId         string `json:"userId"`
	DateCreated    string `json:"dateCreated"`
	Type           string  `json:"type"`
	BarIdList      string   `json:"barIdList"`
}
type QuerySell struct {
	Key    string `json:"Key"`
	Record *Sell
}
type Sell struct {
	DSGId           string `json:"id"`
	OrderId         string `json:"orderId"`
	Grams           string `json:"grams"`
	EstimatedAmount string `json:"estimatedamount"`
	UserId          string `json:"userId"`
	DateCreated     string `json:"dateCreated"`
	Type           string  `json:"type"`
	BarIdList      string   `json:"barIdList"`
}
type QuerySend struct {
	Key    string `json:"Key"`
	Record *Send
}
type Send struct {
	DSGId          string `json:"id"`
	OrderId        string `json:"orderId"`
	Grams          string `json:"grams"`
	SenderUserId   string `json:"senderUserId"`
	ReceiverUserId string `json:"receiverUserId"`
	DateCreated    string `json:"dateCreated"`
	Type           string  `json:"type"`
	BarIdList      string  `json:"barIdList"`
}
type QueryTrade struct {
	Key    string `json:"Key"`
	Record *Trade
}
type Trade struct {
	DSGId   string `json:"id"`
	OrderId string `json:"orderId"`
	Grams   string `json:"grams"`
	UserId  string `json:"userId"`
	Type    string  `json:"type"`
}

func GetUId() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return id.String(), err
}

// InitLedger adds a base set of cars to the ledger
func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {
	fmt.Printf("Hello \n")
	fmt.Printf("chaincod version 1.6 \n")
	return nil
}

func (s *SmartContract) CheckForOrg(ctx contractapi.TransactionContextInterface) (bool, error) {
	fmt.Printf("CheckForOrg function")	
	var response bool
	fmt.Printf("response variable : %v \n",response)
	// Get the MSP ID of submitting client identity
	clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	
	fmt.Printf("clientMSPID %s \n",clientMSPID)
	if err != nil {
		return false, fmt.Errorf("failed to get verified MSPID: %v", err)
	}
	if clientMSPID == invokeMSPID{
		fmt.Printf("in if response variable : %v \n",response)
        response = true
    } else {
		fmt.Printf("in else response variable : %v \n",response)
		response = false
    }
	return response,nil
}

func (s *SmartContract) CreateBar(ctx contractapi.TransactionContextInterface, BarId string, BarLocation string, BarSerialNumber string, Purity string, BarRefiner string, BarHallmarkVerfied string, BarWeightInGms string, DateCreated string) error {

	fmt.Printf("Adding new Gold Bar to the ledger >>>>>>>>>> \n")

	exists, err := s.CheckForOrg(ctx)
	if err != nil {
	  fmt.Printf("Error in from checkorgfunction call")
		return err
	}
	if exists {
		return fmt.Errorf("Not able to add data by Org2 ")
	}

	id := "Bar-" + BarId
	fmt.Printf("Validating Bar data\n")
	//Validate the Org data
	bar := Bar{
		BarID: id,
		BarLocation:        BarLocation,
		BarSerialNumber:    BarSerialNumber,
		Purity:             Purity,
		BarRefiner:         BarRefiner,
		BarHallmarkVerfied: BarHallmarkVerfied,
		BarWeightInGms:     BarWeightInGms,
		DateCreated:        DateCreated,
	}

	barJSON, err := json.Marshal(bar)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, barJSON)
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) GetBar(ctx contractapi.TransactionContextInterface, BarSerialNumber string) ([]QueryBar, error) {
	fmt.Printf("Fetching data getbar function>>>>>>>>>>>>>>> \n")
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Bar-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QueryBar{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		bar := new(Bar)
		_ = json.Unmarshal(queryResponse.Value, bar)
		if bar.BarSerialNumber == BarSerialNumber {

			queryResult := QueryBar{Key: queryResponse.Key, Record: bar}
			result = append(result, queryResult)
		}
	}
	return result, nil
}

func (s *SmartContract) QueryBar(ctx contractapi.TransactionContextInterface, BarSerialNumber string) (*Bar, error) {
	barAsBytes, err := ctx.GetStub().GetState(BarSerialNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if barAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", BarSerialNumber)
	}

	bar := new(Bar)
	_ = json.Unmarshal(barAsBytes, bar)

	return bar, nil
}

func (s *SmartContract) GetBarHistory(ctx contractapi.TransactionContextInterface, BarSerialNumber string) ([]QueryBar, error) {
	fmt.Printf("start of  history function")
	iterator, err := ctx.GetStub().GetHistoryForKey(BarSerialNumber)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Getting history from the ledger ... %s\n",iterator)
	defer iterator.Close()
	results := []QueryBar{}
	for iterator.HasNext() {
		queryResponse, err := iterator.Next()
		if err != nil {
			return nil, err
		}
		fmt.Printf("In side loop --history function ...  %s \n",queryResponse.Value)
		bar := new(Bar)
		_ = json.Unmarshal(queryResponse.Value, bar)
		queryResult := QueryBar{Key: BarSerialNumber, Record: bar}
		results = append(results, queryResult)
	}
	fmt.Printf("details history from the ledger ...  %s \n",results)
	return results, nil
}

func (s *SmartContract) GetBarList(ctx contractapi.TransactionContextInterface) ([]QueryBar, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Bar-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QueryBar{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		bar := new(Bar)
		_ = json.Unmarshal(queryResponse.Value, bar)
		queryResult := QueryBar{Key: queryResponse.Key, Record: bar}
		results = append(results, queryResult)
	}
	return results, nil
}
func (s *SmartContract) CreateBuy(ctx contractapi.TransactionContextInterface, BuyId string, OrderId string, Amount string, AmountWithFees string, Stage string, PaymentStatus string, EstimatedGrams string, UserId string, DateCreated string,Type string,BarIdList string) error {

	fmt.Printf("Adding Buy to the ledger ...\n")

	exists, err := s.CheckForOrg(ctx)
	if err != nil {
	  fmt.Printf("Error in from checkorgfunction call")
		return err
	}
	if exists {
		return fmt.Errorf("Not able to add data by Org2 ")
	}

	id := "DSG-" + BuyId
	fmt.Printf("Validating Buy data\n")
	//Validate the Org data
	var buy = Buy{
		DSGId: id,
		OrderId:        OrderId,
		Amount:         Amount,
		AmountWithFees: AmountWithFees,
		Stage:          Stage,
		PaymentStatus:  PaymentStatus,
		EstimatedGrams: EstimatedGrams,
		UserId:         UserId,
		DateCreated:    DateCreated,
		Type:           Type,
		BarIdList:  BarIdList,
	}

	//Encrypt and Marshal Org data in order to put in world state
	buyAsBytes, _ := json.Marshal(buy)

	return ctx.GetStub().PutState(id, buyAsBytes)

}
func (s *SmartContract) QueryBuy(ctx contractapi.TransactionContextInterface, OrderId string) (*Buy, error) {
	buyAsBytes, err := ctx.GetStub().GetState(OrderId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if buyAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", OrderId)
	}

	buy := new(Buy)
	_ = json.Unmarshal(buyAsBytes, buy)

	return buy, nil
}
func (s *SmartContract) GetBuy(ctx contractapi.TransactionContextInterface, OrderId string) ([]QueryBuy, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QueryBuy{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		buy := new(Buy)
		_ = json.Unmarshal(queryResponse.Value, buy)
		if buy.OrderId == OrderId {

			queryResult := QueryBuy{Key: queryResponse.Key, Record: buy}
			result = append(result, queryResult)
		}
	}
	return result, nil
}
func (s *SmartContract) GetBuyList(ctx contractapi.TransactionContextInterface, OrderId string) ([]QueryBuy, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QueryBuy{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		buy := new(Buy)
		_ = json.Unmarshal(queryResponse.Value, buy)
		if buy.OrderId == OrderId {
			queryResult := QueryBuy{Key: queryResponse.Key, Record: buy}
			results = append(results, queryResult)
		}
	}
	return results, nil
}
func (s *SmartContract) CreateSell(ctx contractapi.TransactionContextInterface, SellId string, OrderId string, Grams string, EstimatedAmount string, UserId string, DateCreated string,Type string,BarIdList string) error {

	fmt.Printf("Adding Sell to the ledger ...\n")

	exists, err := s.CheckForOrg(ctx)
	if err != nil {
	  fmt.Printf("Error in from checkorgfunction call")
		return err
	}
	if exists {
		return fmt.Errorf("Not able to add data by Org2 ")
	}

	id := "DSG-" + SellId
	fmt.Printf("Validating Sell data\n")
	//Validate the Org data
	var sell = Sell{
		DSGId: id,
		OrderId:         OrderId,
		Grams:           Grams,
		EstimatedAmount: EstimatedAmount,
		UserId:          UserId,
		DateCreated:     DateCreated,
		Type :          Type,
		BarIdList:   BarIdList,
	}
	//Encrypt and Marshal Org data in order to put in world state
	sellAsBytes, _ := json.Marshal(sell)

	return ctx.GetStub().PutState(id, sellAsBytes)

}
func (s *SmartContract) QuerySell(ctx contractapi.TransactionContextInterface, OrderId string) (*Sell, error) {
	sellAsBytes, err := ctx.GetStub().GetState(OrderId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if sellAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", OrderId)
	}

	sell := new(Sell)
	_ = json.Unmarshal(sellAsBytes, sell)

	return sell, nil
}
func (s *SmartContract) GetSell(ctx contractapi.TransactionContextInterface, OrderId string) ([]QuerySell, error) {
	fmt.Printf("-----------------start of function--------\n")

	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QuerySell{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		sell := new(Sell)
		_ = json.Unmarshal(queryResponse.Value, sell)
		if sell.OrderId == OrderId {

			queryResult := QuerySell{Key: queryResponse.Key, Record: sell}
			result = append(result, queryResult)
		}
	}
	fmt.Printf("Change need to validate in chaincode update\n")
	return result, nil
}
func (s *SmartContract) GetSellList(ctx contractapi.TransactionContextInterface, OrderId string) ([]QuerySell, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QuerySell{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		sell := new(Sell)
		_ = json.Unmarshal(queryResponse.Value, sell)
		if sell.OrderId == OrderId {
			queryResult := QuerySell{Key: queryResponse.Key, Record: sell}
			results = append(results, queryResult)
		}
	}
	return results, nil
}
func (s *SmartContract) CreateSend(ctx contractapi.TransactionContextInterface, SendId string, OrderId string, Grams string, SenderUserId string, ReceiverUserId string, DateCreated string, Type string,BarIdList string) error {

	fmt.Printf("Adding Send to the ledger ...\n")

	exists, err := s.CheckForOrg(ctx)
	if err != nil {
	  fmt.Printf("Error in from checkorgfunction call")
		return err
	}
	if exists {
		return fmt.Errorf("Not able to add data by Org2 ")
	}

	id := "DSG-" + SendId
	fmt.Printf("Validating Send data\n")
	//Validate the Org data
	var send = Send{
		DSGId:          id,
		OrderId:        OrderId,
		Grams:          Grams,
		SenderUserId:   SenderUserId,
		ReceiverUserId: ReceiverUserId,
		DateCreated:    DateCreated,
		Type:          Type,
		BarIdList:  BarIdList,
	}

	//Encrypt and Marshal Org data in order to put in world state
	sendAsBytes, _ := json.Marshal(send)

	return ctx.GetStub().PutState(id, sendAsBytes)

}
func (s *SmartContract) QuerySend(ctx contractapi.TransactionContextInterface, OrderId string) (*Send, error) {
	sendAsBytes, err := ctx.GetStub().GetState(OrderId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if sendAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", OrderId)
	}

	send := new(Send)
	_ = json.Unmarshal(sendAsBytes, send)

	return send, nil
}
func (s *SmartContract) GetSend(ctx contractapi.TransactionContextInterface, OrderId string) ([]QuerySend, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QuerySend{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		send := new(Send)
		_ = json.Unmarshal(queryResponse.Value, send)
		if send.OrderId == OrderId {

			queryResult := QuerySend{Key: queryResponse.Key, Record: send}
			result = append(result, queryResult)
		}
	}
	return result, nil
}
func (s *SmartContract) GetSendList(ctx contractapi.TransactionContextInterface, OrderId string) ([]QuerySend, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QuerySend{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		send := new(Send)
		_ = json.Unmarshal(queryResponse.Value, send)
		if send.OrderId == OrderId {
			queryResult := QuerySend{Key: queryResponse.Key, Record: send}
			results = append(results, queryResult)
		}
	}
	return results, nil
}
func (s *SmartContract) CreateTrade(ctx contractapi.TransactionContextInterface, TradeId string, OrderId string, Grams string, UserId string, Type string) error {

	fmt.Printf("Adding Trade to the ledger ...\n")

	exists, err := s.CheckForOrg(ctx)
	if err != nil {
	  fmt.Printf("Error in from checkorgfunction call")
		return err
	}
	if exists {
		return fmt.Errorf("Not able to add data by Org2 ")
	}

	id := "DSG-" + TradeId
	fmt.Printf("Validating Trade data\n")
	//Validate the Org data
	var trade = Trade{
		DSGId:   id,
		OrderId: OrderId,
		Grams:   Grams,
		UserId:  UserId,
		Type:    Type,
	}

	//Encrypt and Marshal Org data in order to put in world state
	tradeAsBytes, _ := json.Marshal(trade)

	return ctx.GetStub().PutState(id, tradeAsBytes)

}
func (s *SmartContract) QueryTrade(ctx contractapi.TransactionContextInterface, OrderId string) (*Trade, error) {
	tradeAsBytes, err := ctx.GetStub().GetState(OrderId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if tradeAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", OrderId)
	}

	trade := new(Trade)
	_ = json.Unmarshal(tradeAsBytes, trade)

	return trade, nil
}
func (s *SmartContract) GetTrade(ctx contractapi.TransactionContextInterface, OrderId string) ([]QueryTrade, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QueryTrade{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		trade := new(Trade)
		_ = json.Unmarshal(queryResponse.Value, trade)
		if trade.OrderId == OrderId {

			queryResult := QueryTrade{Key: queryResponse.Key, Record: trade}
			result = append(result, queryResult)
		}
	}
	return result, nil
}
func (s *SmartContract) GetTradeList(ctx contractapi.TransactionContextInterface, OrderId string) ([]QueryTrade, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QueryTrade{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		trade := new(Trade)
		_ = json.Unmarshal(queryResponse.Value, trade)
		if trade.OrderId == OrderId {
			queryResult := QueryTrade{Key: queryResponse.Key, Record: trade}
			results = append(results, queryResult)
		}
	}
	return results, nil
}
func (s *SmartContract) GetTransactionTradeList(ctx contractapi.TransactionContextInterface,Key string, Type string) ([]QueryTrade, error) {
	
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QueryTrade{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		trade := new(Trade)
		_ = json.Unmarshal(queryResponse.Value, trade)
		// if trade.OrderId == OrderId {
		// 	queryResult := QueryTrade{Key: queryResponse.Key, Record: trade}
		// 	results = append(results, queryResult)
		// }
	//	if trade.Type == Type && queryResponse.Key == Key{
			queryResult := QueryTrade{Key: queryResponse.Key, Record: trade}
			results = append(results, queryResult)
		//}
	}
	return results, nil
}
func (s *SmartContract) GetTransactionBuyList(ctx contractapi.TransactionContextInterface,Key string, Type string) ([]QueryBuy, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QueryBuy{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		buy := new(Buy)
		_ = json.Unmarshal(queryResponse.Value, buy)
		// if trade.OrderId == OrderId {
		// 	queryResult := QueryTrade{Key: queryResponse.Key, Record: trade}
		// 	results = append(results, queryResult)
		// }
		//if buy.Type == Type {
			queryResult := QueryBuy{Key: queryResponse.Key, Record: buy}
			results = append(results, queryResult)
		//}
	}
	return results, nil
}
func (s *SmartContract) GetTransactionList(ctx contractapi.TransactionContextInterface) ([]QueryBuy, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QueryBuy{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		buy := new(Buy)
		_ = json.Unmarshal(queryResponse.Value, buy)
		// if trade.OrderId == OrderId {
		// 	queryResult := QueryTrade{Key: queryResponse.Key, Record: trade}
		// 	results = append(results, queryResult)
		// }
		//if buy.Type == Type {
			queryResult := QueryBuy{Key: queryResponse.Key, Record: buy}
			results = append(results, queryResult)
		//}
	}
	return results, nil
}
func (s *SmartContract) GetTransactionSendList(ctx contractapi.TransactionContextInterface,Key string, Type string) ([]QuerySend, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QuerySend{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		send := new(Send)
		_ = json.Unmarshal(queryResponse.Value, send)
		// if trade.OrderId == OrderId {
		// 	queryResult := QueryTrade{Key: queryResponse.Key, Record: trade}
		// 	results = append(results, queryResult)
		// }
		// if send.Type == Type{
			queryResult := QuerySend{Key: queryResponse.Key, Record: send}
			results = append(results, queryResult)
		// }
	}
	return results, nil
}
func (s *SmartContract) GetTransactionSellList(ctx contractapi.TransactionContextInterface,Key string, Type string) ([]QuerySell, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QuerySell{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		sell := new(Sell)
		_ = json.Unmarshal(queryResponse.Value, sell)
		// if trade.OrderId == OrderId {
		// 	queryResult := QueryTrade{Key: queryResponse.Key, Record: trade}
		// 	results = append(results, queryResult)
		// }
		// if sell.Type == Type {
			queryResult := QuerySell{Key: queryResponse.Key, Record: sell}
			results = append(results, queryResult)
		// }
	}
	return results, nil
}

func main() {
	// See chaincode.env.example
	config := serverConfig{
		CCID:    os.Getenv("CHAINCODE_ID"),
		Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
	}

	chaincode, err := contractapi.NewChaincode(&SmartContract{})

	if err != nil {
		log.Panicf("error create asset-transfer-basic chaincode: %s", err)
	}

	server := &shim.ChaincodeServer{
		CCID:    config.CCID,
		Address: config.Address,
		CC:      chaincode,
		TLSProps: shim.TLSProperties{
			Disabled: true,
		},
	}

	if err := server.Start(); err != nil {
		log.Panicf("error starting asset-transfer-basic chaincode: %s", err)
	}
}