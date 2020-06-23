package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"

	//_ . "github.com/SupplyChainDemo/PO"
	//"github.com/minehub/minehub-chaincode/testdata"
	"testing"
)

type SimpleChaincode1 struct {
}


func Test_SupplyChain_test(t *testing.T) {

	stub := shim.NewMockStub("mockStub", new(SmartContract))

	if stub == nil {

		t.Fatalf("MockStub creation failed")

	}

	fmt.Println("==============================<<<<<<< createGoldBar 1>>>>>>>==========================================")

	createGoldBar1 := stub.MockInvoke("createGoldBar", [][]byte{[]byte("createGoldBar"), []byte(createGoldBar1)})
	fmt.Println("Printing Test result for createGoldBar ::::", string(createGoldBar1.Payload))

	fmt.Println("==============================<<<<<<< createGoldBar 2>>>>>>>==========================================")


	createGoldBar2 := stub.MockInvoke("createGoldBar", [][]byte{[]byte("createGoldBar"), []byte(createGoldBar2)})
	fmt.Println("Printing Test result for createGoldBar ::::", string(createGoldBar2.Payload))

	//createGoldBar3 := stub.MockInvoke("createGoldBar", [][]byte{[]byte("createGoldBar"), []byte(createGoldBar3)})
	//fmt.Println("Printing Test result for createGoldBar ::::", string(createGoldBar3.Payload))


	fmt.Println("==============================<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("==============================<<<<<<< purchaseGoldBar >>>>>>>==========================================")

	purchaseGoldBar := stub.MockInvoke("createGoldBar", [][]byte{[]byte("purchaseGold"), []byte(purchaseBar)})
	fmt.Println("Printing Test result for purchaseBar ::::", string(purchaseGoldBar.Payload))


	fmt.Println("==============================<<<<<<<Get Bar Data post purchase getGoldBars >>>>>>>==========================================")
	getGoldBar11:= stub.MockInvoke("getGoldBars", [][]byte{[]byte("getGoldBars"), []byte("BAR-Valacumbi Sussie")})
	fmt.Println("Printing Test result for getGoldBar ::::", string(getGoldBar11.Payload))


	//fmt.Println("==============================<<<<<<< Get UserData  ====>  Vinod >>>>>>>==========================================")
	//getUserData:= stub.MockInvoke("getUserData", [][]byte{[]byte("getUserData"), []byte("vinod")})
	//fmt.Println("Printing Test result for getUserData ::::", string(getUserData.Payload))


	fmt.Println("==============================<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("==============================<<<<<<< giftGoldBar >>>>>>>==========================================")
	giftGold := stub.MockInvoke("giftGold", [][]byte{[]byte("giftGold"), []byte(giftBar)})
	fmt.Println("Printing Test result for purchaseBar ::::", string(giftGold.Payload))


	fmt.Println("==============================<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("==============================<<<<<<< Get UserData  ====>  Vipin >>>>>>>==========================================")
	getUserData1 := stub.MockInvoke("getUserData", [][]byte{[]byte("getUserData"), []byte("vipin")})
	fmt.Println("Printing Test result for getUserData ::::", string(getUserData1.Payload))

//	fmt.Println("==============================<<<<<<< Get UserData  ====>  Vinod >>>>>>>==========================================")
//	getUserData:= stub.MockInvoke("getUserData", [][]byte{[]byte("getUserData"), []byte("vinod")})
//	fmt.Println("Printing Test result for getUserData ::::", string(getUserData.Payload))

fmt.Println("==============================<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
fmt.Println("==============================<<<<<<< before selling :::::::: getGoldReserve >>>>>>>==========================================")
getGoldReserve1:= stub.MockInvoke("getGoldReserve", [][]byte{[]byte("getGoldReserve"), []byte("Valacumbi Sussie")})
fmt.Println("Printing Test result for getGoldBar ::::", string(getGoldReserve1.Payload))

	fmt.Println("==============================<<<<<<< sellGoldBar >>>>>>>==========================================")

	sellGold := stub.MockInvoke("sellGold", [][]byte{[]byte("sellGold"), []byte(sellBar)})
	fmt.Println("Printing Test result for purchaseBar ::::", string(sellGold.Payload))

	fmt.Println("==============================<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("==============================<<<<<<< after selling :::: getGoldReserve >>>>>>>==========================================")
	getGoldReserve:= stub.MockInvoke("getGoldReserve", [][]byte{[]byte("getGoldReserve"), []byte("Valacumbi Sussie")})
	fmt.Println("Printing Test result for getGoldBar ::::", string(getGoldReserve.Payload))

	fmt.Println("==============================<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("==============================<<<<<<< Get UserData  ====>  Vinod >>>>>>>==========================================")
	getUserData:= stub.MockInvoke("getUserData", [][]byte{[]byte("getUserData"), []byte("vinod")})
	fmt.Println("Printing Test result for getUserData ::::", string(getUserData.Payload))


	fmt.Println("==============================<<<<<<< Get UserData  ====>  Vipin >>>>>>>==========================================")
	getUserData2 := stub.MockInvoke("getUserData", [][]byte{[]byte("getUserData"), []byte("vipin")})
	fmt.Println("Printing Test result for getUserData ::::", string(getUserData2.Payload))


	fmt.Println("==============================<<<<<<< getGoldBars >>>>>>>==========================================")
	getGoldBar:= stub.MockInvoke("getGoldBars", [][]byte{[]byte("getGoldBars"), []byte("BAR-Valacumbi Sussie")})
	fmt.Println("Printing Test result for getGoldBar ::::", string(getGoldBar.Payload))


	
}



const createGoldBar1 = `{
	"barLocation"         : "Bangalore",
	"barSerialNumber"     : 1234567,
	"barPurity"           : 995.50,
	"barHallmarkVerfied"  :  true,
	"barRefiner"          : "Valacumbi Sussie",
	"barWeightInGms"      : 1000,
	"barId"               : ""
	}`

	const createGoldBar2 = `{
		"barLocation"         : "Bangalore",
		"barSerialNumber"     : 12345679,
		"barPurity"           : 995.50,
		"barHallmarkVerfied"  :  true,
		"barRefiner"          : "Valacumbi Sussie",
		"barWeightInGms"      : 595,
		"barId"               : ""    
		}`

		const createGoldBar3 = `{
			"barLocation"         : "Bangalore",
			"barSerialNumber"     : 12345610,
			"barPurity"           : 995.50,
			"barHallmarkVerfied"  :  true,
			"barRefiner"          : "Valacumbi Sussie",
			"barWeightInGms"      : 995,
			"barId"               : ""    
			}`

		const purchaseBar = `{
			"username" : "vinod",
			"createAt" : "bangalore",
			"orderId"  : "1234",
			"userId"   : "vinod",
			"orderDate"  : "20201606",
			"weightInGms" : 200,
			"symbol" : "VS",
			"paymentStatus" : "PAID",
			"kycId" : "V123456",
			"kycVerified": true,
			"kycLastVerfied" : "20201606",
			"paymentProviderId" : "Card",
			"paymentCardId"     :  "P12345",
			"paymentCardDetails" : [{
				"paymentProviderId" : "Card",
			"paymentCardId"     :  "P12345"
			}],
			"authenticatedShopperId": "111111"
		}` 

		const giftBar = `{
			"senderUserId" : "vinod",
			"recieverUserId" : "vipin",
			"createAt" : "bangalore",
			"orderId"  : "1234",
			"orderDate"  : "20201606",
			"weightInGms" : 200,
			"symbol" : "VS",
			"paymentStatus" : "PAID",
			"kycId" : "V123456",
			"kycVerified": true,
			"kycLastVerfied" : "20201606",
			"kycStatus"   :   "Finished",
			"kycProviderID" : "1111",
			"kycEnabledFlag" :  true
			}` 

		const sellBar = `{
			"senderUserId" : "vipin",
			"createAt" : "bangalore",
			"orderId"  : "1234",
			"userId"   : "vinod",
			"orderDate"  : "20201606",
			"weightInGms" : 200,
			"symbol" : "VS",
			"paymentStatus" : "PAID",
			"kycId" : "V123456",
			"kycVerified": true,
			"kycLastVerfied" : "20201606",
			"paymentProviderId" : "Card",
			"paymentCardId"     :  "P12345",
			"paymentCardDetails" : [{
				"paymentProviderId" : "Card",
			"paymentCardId"     :  "P12345"
			}],
			"authenticatedShopperId": "111111"
		}` 
