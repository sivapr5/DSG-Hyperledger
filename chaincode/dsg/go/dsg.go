package main

 import (
	 "encoding/json"
	 "fmt"
	//  "byte"

	 "github.com/hyperledger/fabric-contract-api-go/contractapi"
	 "github.com/google/uuid"
 )
 
 // SmartContract provides functions for managing a car
 type SmartContract struct {
	 contractapi.Contract
 }
 type QueryBar struct {
	Key    string `json:"Key"`
	Record *Bar

}
type Bar struct {
	BarID  	                string   `json:"id"`
	BarLocation       		string   `json:"barLocation"`
	BarSerialNumber 	    string    `json:"barSerialNumber"`
	Purity                  float32  `json:"purity"`
	BarRefiner   	        string   `json:"barRefiner"`
	BarHallmarkVerfied		bool     `json:"barHallmarkVerfied"`
	BarWeightInGms		    float32  `json:"barWeightInGms"` 
}
type QueryBuy struct {
	Key    string `json:"Key"`
	Record *Buy

}
type Buy struct {
	DSGId                   string   `json:"id"`
	OrderId                 string   `json:"orderId"`
    Amount                  string  `json:"amount"`
    AmountWithFees          string  `json:"amountWithFees"`
    Stage                   string     `json:"stage"`
	PaymentStatus           string     `json:"paymentStatus"`
	EstimatedGrams          string  `json:"estimatedGrams"`
    UserId                  string    `json:"userId"`
    
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
	fmt.Printf("Hello\n")
	return nil
 }
func (s *SmartContract) CreateBar(ctx contractapi.TransactionContextInterface,BarLocation string,BarSerialNumber string,Purity float32,BarRefiner string,BarHallmarkVerfied bool,BarWeightInGms float32) error {


	fmt.Printf("Adding Bar to the ledger ...\n")
	// if len(args) != 8 {
	// 	return fmt.Errorf("InvalidArgumentError: Incorrect number of arguments. Expecting 8")
    // }

    //Prepare key for the new Org
	uid, err := GetUId()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	id := "Bar-" + uid
	fmt.Printf("Validating Bar data\n")
	//Validate the Org data
	var bar = Bar{BarID: id,			   
					BarLocation: BarLocation,
					BarSerialNumber: BarSerialNumber,
					Purity: Purity,
					BarRefiner: BarRefiner,
					BarHallmarkVerfied: BarHallmarkVerfied,
					BarWeightInGms: BarWeightInGms,

					}

	//Encrypt and Marshal Org data in order to put in world state
	barAsBytes, _ := json.Marshal(bar)

	return ctx.GetStub().PutState(id, barAsBytes)
}
// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) GetBar(ctx contractapi.TransactionContextInterface,BarSerialNumber string) ([]QueryBar, error) {
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
		if bar.BarSerialNumber == BarSerialNumber{

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
		queryResult := QueryBar{Key: queryResponse.Key,Record: bar}
		results = append(results, queryResult)
	}
	return results, nil
}
func (s *SmartContract) CreateBuy(ctx contractapi.TransactionContextInterface,OrderId string,Amount string,AmountWithFees string,Stage string,PaymentStatus string,EstimatedGrams string,UserId string) error {


	fmt.Printf("Adding Buy to the ledger ...\n")
	// if len(args) != 8 {
	// 	return fmt.Errorf("InvalidArgumentError: Incorrect number of arguments. Expecting 8")
    // }

    //Prepare key for the new Org
	uid, err := GetUId()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	id := "DSG-" + uid
	fmt.Printf("Validating Buy data\n")
	//Validate the Org data
	var buy = Buy{DSGId: id,			                  
		            OrderId: OrderId,       		   	              
					Amount: Amount,                
					AmountWithFees: AmountWithFees, 
					Stage: Stage, 
					PaymentStatus: PaymentStatus, 
					EstimatedGrams: EstimatedGrams,
					UserId: UserId,      
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
func (s *SmartContract) GetBuy(ctx contractapi.TransactionContextInterface,OrderId string) ([]QueryBuy, error) {
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
		if buy.OrderId == OrderId{

		queryResult := QueryBuy{Key: queryResponse.Key, Record: buy}
		result = append(result, queryResult)
	}
	}
	return result, nil
}
func (s *SmartContract) GetBuyList(ctx contractapi.TransactionContextInterface,OrderId string) ([]QueryBuy, error) {
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
		if buy.OrderId == OrderId{
		queryResult := QueryBuy{Key: queryResponse.Key,Record: buy}
		results = append(results, queryResult)
		}
	}
	return results, nil
}
func main() {
 
	 chaincode, err := contractapi.NewChaincode(new(SmartContract))
 
	 if err != nil {
		 fmt.Printf("Error create  chaincode: %s", err.Error())
		 return
	 }
 
	 if err := chaincode.Start(); err != nil {
		 fmt.Printf("Error starting  chaincode: %s", err.Error())
	 }
 }
 
