package main

import ( 
	
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"math/rand"
	"strconv"
	"time"
)

var count int = 0

var _MAIN_LOGGER = shim.NewLogger("SmartContractMain")


// SimpleChaincode example simple Chaincode implementation
type SmartContract struct {
}


func (sc *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}


func (sc *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	_MAIN_LOGGER.Info("Inside invoke method")
	function, _ := stub.GetFunctionAndParameters()
	_MAIN_LOGGER.Info("invoke is running " + function)


	if function == "createGoldBar" {
		res := sc.CreateGoldBar(stub)
		return shim.Success(MarshalToBytes(res))
	} else if function == "getGoldReserve" {
		obj, _ := sc.GetReserveGold(stub)
		return shim.Success(MarshalToBytes(obj))
	} else if function == "purchaseGold" {
		pc,_ := sc.PurchaseGold(stub)
		return shim.Success(MarshalToBytes(pc))
	} else if function == "giftGold" {
		pc,_ := sc.GiftGold(stub)
		return shim.Success(MarshalToBytes(pc))
	} else if function == "sellGold" {
		pc,_ := sc.SellGold(stub)
		return shim.Success(MarshalToBytes(pc))
	} else if function == "getUserData" {
		pc,_ := sc.GetUserData(stub)
		return shim.Success(MarshalToBytes(pc))
	} else if function == "getReserveGold" {
		pc,_ := sc.GetReserveGold(stub)
		return shim.Success(MarshalToBytes(pc))
	} else if function == "getGoldBars" {
		pc,_ := sc.GetGoldBars(stub)
		return shim.Success(MarshalToBytes(pc))
	} else {
		fmt.Println("invoke did not find func: " + function) 
	}
	
	
	
	return shim.Error("Received unknown function invocation")

}


func (sc *SmartContract) CreateGoldBar(stub shim.ChaincodeStubInterface) string {
	_MAIN_LOGGER.Info("Inside CreateGoldBar Method")

	defer _MAIN_LOGGER.Info("Exit CreateGoldBar Method")

	_, args := stub.GetFunctionAndParameters()
	if len(args) < 1 {
		_MAIN_LOGGER.Info("AddBarData" + ": Incorrect number of arguments.")
		return "Incorrect number of arguments"
	}

	var commonInput GoldBarData
	//var commonInputArr []GoldBarData

	//fmt.Println("Printing input ::", args[0])

	err := json.Unmarshal([]byte(args[0]), &commonInput)

	if err != nil {
		_MAIN_LOGGER.Error("Error while unmarshal Bar data", err)
		//utils.LOGGER.Error("CreateContract: Error during json.Unmarshal: ", err)
		return "Error while unmarshal Bar data"
	}

	key1 := strconv.Itoa(rand.Intn(10000))
	key2 := strconv.Itoa(rand.Intn(1000))

	commonInput.BarID = "VS"+key1+key2+strconv.Itoa(count)
	//commonInput.TotalGoldAmount = commonInput.TotalGoldAmount + commonInput.BarWeightInGms

	_MAIN_LOGGER.Info("Print commonInput.BarID ::", commonInput.BarID)

	commonInputAsBytes1, erx1 := json.MarshalIndent(commonInput, "", "\r")
	_ = erx1
	errPutState1 := stub.PutState(commonInput.BarID, commonInputAsBytes1)

	if errPutState1 != nil {
		_MAIN_LOGGER.Info("GoldBarData" + ": Error while putting transition data on Ledger.")
		return "Error while putting transition data on Ledger"
	}

	// ==============================================
	// Fetch total Gold and Bar Count from Inventory

	var goldReserveBar GoldReserve

	reserveData, getErr := stub.GetState(commonInput.BarRefiner)
	if reserveData == nil {
		goldReserveBar.TotalGoldAmount = commonInput.BarWeightInGms
		goldReserveBar.BarCount = goldReserveBar.BarCount + 1
		goldReserveBar.BarRefiner = commonInput.BarRefiner
	} else {
		err11 := json.Unmarshal(reserveData, &goldReserveBar)
		if err11 != nil {
			_MAIN_LOGGER.Info("Error while unnarshalling goldreserve data!!")
			return "Error while unnarshalling goldreserve data!!"
		}

		goldReserveBar.TotalGoldAmount = goldReserveBar.TotalGoldAmount + commonInput.BarWeightInGms
		goldReserveBar.BarCount = goldReserveBar.BarCount + 1
		goldReserveBar.BarRefiner = commonInput.BarRefiner

	}

	if getErr != nil {
		_MAIN_LOGGER.Info("Can't find any gold reserve data in ledger!!")
	}

	//commonInputArr = append(commonInputArr, commonInput)

	goldRerserveAsBytes1, _ := json.MarshalIndent(goldReserveBar, "", "\r")
	//_ = erx1
	errPutState2 := stub.PutState(goldReserveBar.BarRefiner, goldRerserveAsBytes1)

	if errPutState2 != nil {
		_MAIN_LOGGER.Info("GoldBarReserveData" + ": Error while putting transition data on Ledger.")
		return "Error while putting transition data on Ledger"
	}

	_MAIN_LOGGER.Info("Print GoldBarReserve Data ::::", goldReserveBar)	


	// ==================================
	// Storing bars in a array
	var allGoldBarObj AllGoldBarDetails
	var allGoldBarArray AllGoldBarDetails
 
	allGoldBarArray.DocType = "BAR"+"-"+commonInput.BarRefiner
	allGoldBarObj.DocType = "BAR"+"-"+commonInput.BarRefiner

	allBarData, err := stub.GetState(allGoldBarArray.DocType)
	if allBarData == nil {
		_MAIN_LOGGER.Info("No data found for given key ::", allGoldBarArray.DocType)

		allGoldBarObj.GoldBarArray = append(allGoldBarObj.GoldBarArray ,commonInput)

		allGoldBarDetailsAsBytes1, _ := json.MarshalIndent(allGoldBarObj.GoldBarArray, "", "\r")
	
		errPutState3 := stub.PutState(allGoldBarArray.DocType, allGoldBarDetailsAsBytes1)

		if errPutState3 != nil {
			_MAIN_LOGGER.Info("GoldBarReserveData" + ": Error while putting transition data on Ledger.")
			return "Error while putting transition data on Ledger"
		}

		_MAIN_LOGGER.Info("Print GoldBarArray when adding for first time :::::::::::::::::::", allGoldBarObj.GoldBarArray)

	} else {
		err0 := json.Unmarshal(allBarData, &allGoldBarArray.GoldBarArray)
		if err0 != nil {
			_MAIN_LOGGER.Errorf("Error while unmarshalling data!!")
		}
		fmt.Println("===========================================================")
		fmt.Println("===========================================================")
		_MAIN_LOGGER.Info("Print unmarshalled datA ::", allGoldBarArray.GoldBarArray)
		fmt.Println("===========================================================")
		fmt.Println("===========================================================")

		arr := allGoldBarArray.GoldBarArray
		allGoldBarArray.GoldBarArray = append(arr, commonInput)

		allGoldBarDetailsAsBytes2, _ := json.MarshalIndent(allGoldBarArray.GoldBarArray, "", "\r")
	
		errPutState3 := stub.PutState(allGoldBarArray.DocType, allGoldBarDetailsAsBytes2)

		if errPutState3 != nil {
			_MAIN_LOGGER.Info("GoldBarReserveData" + ": Error while putting transition data on Ledger.")
			return "Error while putting transition data on Ledger"
		}

		_MAIN_LOGGER.Info("Print GoldBarArray when adding  :::::::::::::::::::", allGoldBarArray.GoldBarArray)
	}

//	err = json.Unmarshal(allBarData, &allGoldBarArray)

/*	if err != nil {
		_MAIN_LOGGER.Errorf("Error in parsing goldbar Obj", err)
		//return nullObj, "Error in Unmarshalling data"
	}*/
	

	// allGoldBarObj.GoldBarArray = append(allGoldBarObj.GoldBarArray ,commonInput)


	// ==================================

	count = count + 1
	return commonInput.BarID
}



func (sc *SmartContract) PurchaseGold(stub shim.ChaincodeStubInterface) (DWR, string) {
	_MAIN_LOGGER.Info("Inside PurchaseGold Method")

	defer _MAIN_LOGGER.Info("Exit PurchaseGold Method")

	//var commonInput GoldBarData

	var purchaseGoldBar PurchaseGoldBar
	var goldReserve GoldReserve
	var purchaseConfirmation PurchaseConfirmation
	var dwr DWR
	var userWallet UserWallet

	_, args := stub.GetFunctionAndParameters()
	if len(args) < 1 {
		_MAIN_LOGGER.Info("Incorrect number of arguments")
		return dwr, "Incorrect number of arguments!!"
	}

	/*if len(args) < 1 {
		_MAIN_LOGGER.Error("GetGoldBarata" + ": Incorrect number of arguments.")
	}*/

	purchaseGoldErr := json.Unmarshal([]byte(args[0]), &purchaseGoldBar)
	if purchaseGoldErr != nil {
		_MAIN_LOGGER.Errorf("Error in unmarshalling  purchaseGoldBar Obj %+v", purchaseGoldErr)
	}

	// Get TotalGoldAmount reserved in inventory
	key := "Valacumbi Sussie"
	totalGoldData, err := stub.GetState(key)

	err = json.Unmarshal(totalGoldData, &goldReserve)

	if err != nil {
		_MAIN_LOGGER.Errorf("Error in unmarshalling  totalGoldData Obj %+v", err)
	}

	_MAIN_LOGGER.Info("Print GoldReserveData ::", goldReserve)

			//var purchaseConfirmation PurchaseConfirmation
			

	if goldReserve.TotalGoldAmount >= purchaseGoldBar.WeightInGms {
		
		// Fetch Gold Bars reserved in inventory
		var allGoldBarObj AllGoldBarDetails
		allBarObj := allGoldBarObj.GoldBarArray
		key1 := "BAR-Valacumbi Sussie"

		allBarData, _ := stub.GetState(key1)
		allGoldBarErr := json.Unmarshal(allBarData, &allBarObj)
			if allGoldBarErr != nil {
				_MAIN_LOGGER.Errorf("Error in unmarshalling  AllGoldBar Obj %+v", allGoldBarErr)
			}

			
			
			for i,_ := range allBarObj {
				_MAIN_LOGGER.Info("Inside For Loop for loopling allbarObj")

				if allBarObj[i].BarWeightInGms >= purchaseGoldBar.WeightInGms {
					
					// Check if user is purchasing whole bar
					if allBarObj[i].BarWeightInGms == purchaseGoldBar.WeightInGms {
						goldReserve.BarCount = goldReserve.BarCount - 1
					}
					_MAIN_LOGGER.Info("When we have enough gold in a bar to continue......")
					//var purchaseConfirmation PurchaseConfirmation
					//var dwr DWR

					purchaseConfirmation.GoldBarID = allBarObj[i].BarID
					purchaseConfirmation.GoldBarWeightAssigned = purchaseGoldBar.WeightInGms
					purchaseConfirmation.OrderID = purchaseGoldBar.OrderID
					purchaseConfirmation.UserID = purchaseGoldBar.UserID

					allBarObj[i].BarWeightInGms = allBarObj[i].BarWeightInGms - purchaseGoldBar.WeightInGms
					_MAIN_LOGGER.Info("Print Gold Amount left for the bar",allBarObj[i].BarID,allBarObj[i].BarWeightInGms)

					_MAIN_LOGGER.Info("**********************************************")
					_MAIN_LOGGER.Info("**********************************************")
					_MAIN_LOGGER.Info("Print Availble GoldBar details ::::::::::", allBarObj)
					_MAIN_LOGGER.Info("**********************************************")
					_MAIN_LOGGER.Info("**********************************************")

					now := time.Now().UTC()
					fmt.Println("Current Time :::::", now.String())
					//fmt.Println(now.String())
					dwr.DWRID  =  strconv.Itoa(rand.Intn(100)) + strconv.Itoa(rand.Intn(1000)) + strconv.Itoa(rand.Intn(10000)) + strconv.Itoa(count)
					dwr.DateIssued = now.String()
					dwr.UserID = purchaseGoldBar.UserID
					dwr.WeightInGms = purchaseGoldBar.WeightInGms
					dwr.Purity = allBarObj[i].BarPurity            
					dwr.BarLocation = allBarObj[i].BarLocation
					dwr.HalmarkVerified = allBarObj[i].BarHallmarkVerified
					dwr.GoldBarNumber = allBarObj[i].BarSerialNumber   
					dwr.GoldBarID = allBarObj[i].BarID
					dwr.OrderID = purchaseGoldBar.OrderID

					// ***********************************************************
					// Update Gold amount in User Wallet
					userWallet.UserID = purchaseGoldBar.UserID
					userWallet.GoldBarID = allBarObj[i].BarID
					userWallet.GoldOwnsInGms = purchaseGoldBar.WeightInGms
					userWallet.BarPurity = allBarObj[i].BarPurity

					// Updating total weiggold amount in GoldReserve
					goldReserve.TotalGoldAmount = goldReserve.TotalGoldAmount - purchaseGoldBar.WeightInGms

					reserveAsBytes, _ := json.MarshalIndent(goldReserve, "", "\r")
	
					errPutState5 := stub.PutState(goldReserve.BarRefiner, reserveAsBytes)

					if errPutState5 != nil {
						_MAIN_LOGGER.Errorf("Update goldReserve Data :: " + ": Error while putting transition data on Ledger.", errPutState5)
						return dwr, "Error while putting transition data on Ledger"
					}
					
					_MAIN_LOGGER.Info("Print UserWallet Data ::::", userWallet)
					// ---------------------------------------------------------
					userWalletAsBytes, _ := json.MarshalIndent(userWallet, "", "\r")
	
					errPutState4 := stub.PutState(userWallet.UserID, userWalletAsBytes)

					if errPutState4 != nil {
						_MAIN_LOGGER.Errorf("Update UserWallet Data :: " + ": Error while putting transition data on Ledger.", errPutState4)
						return dwr, "Error while putting transition data on Ledger"
					}
					
					_MAIN_LOGGER.Info("Print UserWallet Data ::::", userWallet)

					// ***********************************************************

					_MAIN_LOGGER.Info("**********************************************")
					_MAIN_LOGGER.Info("**********************************************")
					_MAIN_LOGGER.Info("Print DWR Data ::::", dwr)
					_MAIN_LOGGER.Info("**********************************************")
					_MAIN_LOGGER.Info("**********************************************")

				//	_MAIN_LOGGER.Info("**********************************************")
				//	_MAIN_LOGGER.Info("Print PurchaseConfirmation Data ::::", purchaseConfirmation)
				//	_MAIN_LOGGER.Info("**********************************************")

					dwrDetailsAsBytes2, _ := json.MarshalIndent(dwr, "", "\r")
	
					errPutState3 := stub.PutState(dwr.OrderID, dwrDetailsAsBytes2)

					if errPutState3 != nil {
						_MAIN_LOGGER.Info("GoldBarReserveData" + ": Error while putting transition data on Ledger.")
						return dwr, "Error while putting transition data on Ledger"
					}

					allGoldBarDetailsAsBytes11, _ := json.MarshalIndent(allBarObj, "", "\r")
	
					errPutState7 := stub.PutState(key1, allGoldBarDetailsAsBytes11)

					if errPutState7 != nil {
						_MAIN_LOGGER.Errorf("GoldBarAllData" + ": Error while putting transition data on Ledger.", errPutState7)
						return dwr, "Error while putting transition data on Ledger"
		}
			break
				}
				
			}
		}

	return dwr, ""

}


func (sc *SmartContract) GiftGold(stub shim.ChaincodeStubInterface) (PurchaseConfirmation, string) {
	_MAIN_LOGGER.Info("Inside GiftGold Method")

	//map[string]interface{}

	defer _MAIN_LOGGER.Info("Exit GiftGold Method")

	var sellGold SendGoldBar
	var userWallet UserWallet
	var recieverWallet UserWallet
	var sellConfirmation, nullObj PurchaseConfirmation

	_, args := stub.GetFunctionAndParameters()
	if len(args) < 1 {
		_MAIN_LOGGER.Info("Incorrect number of arguments")
		return nullObj, "Incorrect number of arguments!!"
	}
	fmt.Println("==============================")
	//fmt.Println("Print incoming input ::", args[0])

	//serr1 := json.Unmarshal([]byte(args[0]), &sellGold)

	serr1 := json.Unmarshal([]byte(args[0]), &sellGold)

	//_  = serr1
	//serr1 != nil {
		_MAIN_LOGGER.Info("Error while unmarshalling data ::::", serr1)
	//}

	fmt.Println("Print incoming input ::", sellGold)

	if sellGold.SenderUserID == "" {
		_MAIN_LOGGER.Info("UserID is a must to be passed !!!!")
	}

	userDataFetch, err := stub.GetState(sellGold.SenderUserID)

	err1 := json.Unmarshal(userDataFetch, &userWallet)

	if err1 != nil {
		_MAIN_LOGGER.Errorf("Error in parsing userwallet Obj", err)
		return nullObj,"Error in Unmarshalling data"
	}

	if userWallet.GoldOwnsInGms >= sellGold.WeightInGms {
		recieverWallet.GoldOwnsInGms = recieverWallet.GoldOwnsInGms + sellGold.WeightInGms
		userWallet.GoldOwnsInGms = userWallet.GoldOwnsInGms - sellGold.WeightInGms
		recieverWallet.GoldBarID = userWallet.GoldBarID
		recieverWallet.UserID = sellGold.RecieverUserID
		recieverWallet.BarPurity = userWallet.BarPurity
		recieverWallet.BarLocation = userWallet.BarLocation
		recieverWallet.BarSerialNumber = userWallet.BarSerialNumber
		recieverWallet.BarHallmarkVerified = userWallet.BarHallmarkVerified

					// Update Reciever Wallet data in ledger
					recieverWalletAsBytes2, _ := json.MarshalIndent(recieverWallet, "", "\r")
	
					errPutState3 := stub.PutState(recieverWallet.UserID, recieverWalletAsBytes2)

					if errPutState3 != nil {
						_MAIN_LOGGER.Info("RecieverWalletData" + ": Error while putting transition data on Ledger.")
						return nullObj, "Error while putting transition data on Ledger"
					}


					// Update Sender Wallet Data in Ledger
					senderWalletAsBytes2, _ := json.MarshalIndent(userWallet, "", "\r")
	
					errPutState4 := stub.PutState(userWallet.UserID, senderWalletAsBytes2)

					if errPutState4 != nil {
						_MAIN_LOGGER.Info("SenderWalletData" + ": Error while putting transition data on Ledger.")
						return nullObj, "Error while putting transition data on Ledger"
					}

					_MAIN_LOGGER.Info("Amount of Gold sold to Reciever ::", sellGold.WeightInGms)
					_MAIN_LOGGER.Info("Sender Data after transaction ::::", userWallet)
					_MAIN_LOGGER.Info("Reciever Data after transaction ::::", recieverWallet)

					sellConfirmation.UserID = sellGold.SenderUserID
					sellConfirmation.GoldBarID = userWallet.GoldBarID
					sellConfirmation.GoldBarWeightAssigned = sellGold.WeightInGms
					sellConfirmation.OrderID = sellGold.OrderID
					//_MAIN_LOGGER.Info("")

	} else {
		_MAIN_LOGGER.Info("Sell GoldBar :: User doesn't have enough gold to sell !!")
		return nullObj, "User doesn't have enough gold to sell !!"
	}

	/*sellGold := make(map[string]interface{})
	sgo, _ := json.Marshal(args[0])
	json.Unmarshal(sgo, &sellGold)

	for key, newValue := range sellGold {
		
	}*/
	return sellConfirmation, ""

}


func (sc *SmartContract) SellGold(stub shim.ChaincodeStubInterface) (PurchaseConfirmation, string) {
	_MAIN_LOGGER.Info("Inside SellGold Method")

	//map[string]interface{}

	defer _MAIN_LOGGER.Info("Exit SellGold Method")

	var sellGold SellGoldBar
	var userWallet UserWallet
	var nullObj, purchaseConfirmation PurchaseConfirmation

	_, args := stub.GetFunctionAndParameters()

	// UnMarshal incoming sell Obj
	sellGoldErr := json.Unmarshal([]byte(args[0]), &sellGold)

	if sellGoldErr != nil {
		_MAIN_LOGGER.Errorf("Error in parsing sellGold Obj", sellGoldErr)
		return nullObj, "Error in Unmarshalling data"
	}

	fmt.Println("Print sellbar ::::::::::::::::::::", sellGold)


	// Fetch Sender User data from wallet
	userKey := sellGold.SenderUserID

	getUsrData, userErr := stub.GetState(userKey)
	if userErr != nil {
		_MAIN_LOGGER.Errorf("Error while fetching usrdata from ledger!!", userErr)
		return nullObj, "Error in fetching ledger data!!"
	}

	getUserDataErr := json.Unmarshal(getUsrData, &userWallet)
	if getUserDataErr != nil {
		_MAIN_LOGGER.Errorf("Error in parsing getUserData Obj", getUserDataErr)
		return nullObj, "Error in Unmarshalling user wallet data!!!!"
	}

	fmt.Println("Print userWallet ::::::::::::::::::::", userWallet)


	// Check if User wants to sell all of his gold amount
	if sellGold.WeightInGms == userWallet.GoldOwnsInGms {
		userWallet.GoldOwnsInGms = userWallet.GoldOwnsInGms - sellGold.WeightInGms
		userWallet.GoldBarID = ""
		userWallet.BarPurity = 0

					senderWalletAsBytes2, _ := json.MarshalIndent(userWallet, "", "\r")
	
					errPutState5 := stub.PutState(userWallet.UserID, senderWalletAsBytes2)

					if errPutState5 != nil {
						_MAIN_LOGGER.Info("UserWalletData" + ": Error while putting transition data on Ledger.")
						return nullObj, "Error while putting transition data on Ledger"
					}
	} else {
		userWallet.GoldOwnsInGms = userWallet.GoldOwnsInGms - sellGold.WeightInGms
		senderWalletAsBytes2, _ := json.MarshalIndent(userWallet, "", "\r")
	
					errPutState5 := stub.PutState(userWallet.UserID, senderWalletAsBytes2)

					if errPutState5 != nil {
						_MAIN_LOGGER.Info("UserWalletData" + ": Error while putting transition data on Ledger.")
						return nullObj, "Error while putting transition data on Ledger"
					}
	}

		// Update GoldBar details in Ledger

		var allGoldBarObj AllGoldBarDetails
	//	var newGoldBarObj AllGoldBarDetails
		var newGoldBar GoldBarData
		allBarObj := allGoldBarObj.GoldBarArray
		key3 := "BAR-Valacumbi Sussie"

		allBarData, _ := stub.GetState(key3)
		allGoldBarErr := json.Unmarshal(allBarData, &allBarObj)
			if allGoldBarErr != nil {
				_MAIN_LOGGER.Errorf("Error in unmarshalling  AllGoldBar Obj %+v", allGoldBarErr)
			}

			for i,_ := range allBarObj {
				if (allBarObj[i].BarWeightInGms + sellGold.WeightInGms) <= 1000 {
					allBarObj[i].BarWeightInGms = allBarObj[i].BarWeightInGms + sellGold.WeightInGms

					allBarObjAsBytes2, _ := json.MarshalIndent(allBarObj, "", "\r")
	
					errPutState6 := stub.PutState(key3, allBarObjAsBytes2)

					if errPutState6 != nil {
						_MAIN_LOGGER.Info("UserWalletData" + ": Error while putting transition data on Ledger.", errPutState6)
						return nullObj, "Error while putting transition data on Ledger"
					}

				} else {
					key1 := strconv.Itoa(rand.Intn(10000))
					key2 := strconv.Itoa(rand.Intn(1000))

					newGoldBar.BarID = "VS"+key1+key2+strconv.Itoa(count)
					newGoldBar.BarLocation = userWallet.BarLocation            
					newGoldBar.BarPurity   = userWallet.BarPurity                 
					newGoldBar.BarSerialNumber = userWallet.BarSerialNumber           
					newGoldBar.BarHallmarkVerified = userWallet.BarHallmarkVerified       
					newGoldBar.BarRefiner = userWallet.BarRefiner                
					newGoldBar.BarWeightInGms = sellGold.WeightInGms

					purchaseConfirmation.GoldBarID = newGoldBar.BarID
					purchaseConfirmation.GoldBarWeightAssigned = sellGold.WeightInGms
					purchaseConfirmation.UserID = sellGold.SenderUserID
					purchaseConfirmation.OrderID = sellGold.OrderID

					allBarObj = append(allBarObj, newGoldBar)

					allObjAsBytes2, _ := json.MarshalIndent(allBarObj, "", "\r")
	
					errPutState7 := stub.PutState(key3, allObjAsBytes2)

					if errPutState7 != nil {
						_MAIN_LOGGER.Info("AllBarObj" + ": Error while putting transition data on Ledger.", errPutState7)
						return nullObj, "Error while putting transition data on Ledger"
					}

					// Fetch total GoldReserve Data and update
					var goldReserve GoldReserve
					key := "Valacumbi Sussie"

					barData, getErr := stub.GetState(key)
					if getErr != nil {
					_MAIN_LOGGER.Errorf("Error in fetching data for goldReserve Obj", getErr)
						return nullObj, "Error in Unmarshalling data"
			}

					goldReserveErr := json.Unmarshal(barData, &goldReserve)

					if goldReserveErr != nil {
					_MAIN_LOGGER.Errorf("Error in parsing goldbar Obj", goldReserveErr)
						return nullObj, "Error in Unmarshalling data"
					}
					goldReserve.TotalGoldAmount = goldReserve.TotalGoldAmount + sellGold.WeightInGms
					goldReserve.BarCount = goldReserve.BarCount + 1

					goldReserveBytes2, _ := json.MarshalIndent(goldReserve, "", "\r")
	
					errPutState8 := stub.PutState(key, goldReserveBytes2)

					if errPutState8 != nil {
						_MAIN_LOGGER.Info("GoldReserve" + ": Error while putting transition data on Ledger.", errPutState8)
						return nullObj, "Error while putting transition data on Ledger"
					}
				}
				break
			}

	 


		return purchaseConfirmation, ""
}



func (sc *SmartContract) GetUserData(stub shim.ChaincodeStubInterface) (UserWallet, string) {
	_MAIN_LOGGER.Info("Inside GetUserData Method")

	defer _MAIN_LOGGER.Info("Exit GetUserData Method")

	_, args := stub.GetFunctionAndParameters()

	key := args[0]
	fmt.Println("Pirnt key ::", key)

	var commonInput,nullObj UserWallet

	barData, err := stub.GetState(key)

	err = json.Unmarshal(barData, &commonInput)

	if err != nil {
		_MAIN_LOGGER.Errorf("Error in parsing goldbar Obj", err)
		return nullObj, "Error in Unmarshalling data"
	}

	fmt.Println("Printing GoldReserveData ::::", commonInput)
	return commonInput, ""

}

func (sc *SmartContract) GetReserveGold(stub shim.ChaincodeStubInterface) (GoldReserve, string) {
	_MAIN_LOGGER.Info("Inside GetReserveGold Method")

	defer _MAIN_LOGGER.Info("Exit GetReserveGold Method")

	_, args := stub.GetFunctionAndParameters()

	key := args[0]
	fmt.Println("Pirnt key ::", key)

	var commonInput,nullObj GoldReserve

	barData, err := stub.GetState(key)

	err = json.Unmarshal(barData, &commonInput)

	if err != nil {
		_MAIN_LOGGER.Errorf("Error in parsing goldbar Obj", err)
		return nullObj, "Error in Unmarshalling data"
	}

	fmt.Println("Printing GoldReserveData ::::", commonInput)
	return commonInput, ""

}


func (sc *SmartContract) GetGoldBars(stub shim.ChaincodeStubInterface) ([]GoldBarData, string) {
	_MAIN_LOGGER.Info("Inside GetGoldBars Method")

	defer _MAIN_LOGGER.Info("Exit GetGoldBars Method")

	_, args := stub.GetFunctionAndParameters()

	key := args[0]
	fmt.Println("Pirnt key ::", key)

	var commonInput AllGoldBarDetails
	var nullObj []GoldBarData
	gi := commonInput.GoldBarArray


	barData, err := stub.GetState(key)

	err = json.Unmarshal(barData, &gi)

	if err != nil {
		_MAIN_LOGGER.Errorf("Error in parsing AllGoldBarDetails Obj", err)
		return nullObj, "Error in Unmarshalling data"
	}

	fmt.Println("Printing AllGoldBarDetails ::::", gi)
	return gi, ""

}

func MarshalToBytes(value interface{}) []byte {
	bytes, marshallErr := json.Marshal(value)
	if marshallErr != nil {
		fmt.Println("Error in marshalling : ", marshallErr, value)
		return bytes
	}
	return bytes
}

func main(){
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}