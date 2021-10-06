/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';
const { Gateway, Wallets } = require('fabric-network');
const FabricCAServices = require('fabric-ca-client');
const fs = require('fs');
const path = require('path');
const { v4: uuidv4 } = require('uuid');
const { registerUser, userExist } = require("./registerUser");
const {QueryAllNew,QueryNew,InvokeNew,QueryAllNew1,QueryNew1} =require('./tx')
//const ccpPath = path.resolve(__dirname, '..', '..', 'first-network', 'connection-org1.json');
const ccpPath = path.resolve(__dirname, 'connection-profile/connection-org1.json');
const chaincodeName = "basic";
const channelName = "mychannel"
const request1={
        "org": "Org1MSP",
        "channelName": channelName,
        "chaincodeName": chaincodeName,
        "userId": 'dsg-user'
    }
    const type ={
        "buy":"1","sell":"2","send":"3","trade":"4"
    }

exports.registerUser = async function(req,res,next){
try {
    console.log("Inside registeruser function")
//    let org = req.body.org;
//    let userId = req.body.userId;
    let org = request1.org
   let userId = request1.userId
   console.log("userId :")
    let result = await registerUser({ OrgMSP: org, userId: userId });
    res.send(result);

} catch (error) {
    res.status(500).send(error)
}
}
exports.welcome = async function(req,res,next){
    return res.send({"msg":"Welcome to Dsg"})
}
exports.createBar = async function (req, res, next) {
	var BarLocation         = req.body.barLocation;
	var BarSerialNumber     = req.body.barSerialNumber;
	var Purity  	        = req.body.purity;
	var BarRefiner   	    = req.body.barRefiner;
	var BarHallmarkVerfied  = req.body.barHallmarkVerfied;
    var BarWeightInGms		= req.body.barWeightInGms;
    var DateCreated = req.body.dateCreated;
    let BarId               = uuidv4()
    console.log("Bar id uuid :",BarId)
    var args = [BarId,BarLocation,BarSerialNumber,Purity,BarRefiner,BarHallmarkVerfied,BarWeightInGms,DateCreated];
    
   // Invoke("CreateBar", args, res);
   InvokeNew(request1,"CreateBar", args, res);
   
}
exports.createBarWithToken = async function (req, res, next) {
    console.log("function createBarWithToken body:",req.body)
    var BarLocation         = req.body.barLocation;
	var BarSerialNumber     = req.body.barSerialNumber;
	var Purity  	        = req.body.purity;
	var BarRefiner   	    = req.body.barRefiner;
	var BarHallmarkVerfied  = req.body.barHallmarkVerfied;
    var BarWeightInGms		= req.body.barWeightInGms;
    var DateCreated = req.body.dateCreated;
    let BarId               = req.body.token;
    console.log("Bar id uuid :",BarId)
    var args = [BarId,BarLocation,BarSerialNumber,Purity,BarRefiner,BarHallmarkVerfied,BarWeightInGms,DateCreated];
    
   // Invoke("CreateBar", args, res);
   InvokeNew(request1,"CreateBar", args, res);
   
}
exports.queryBar = async function (req, res, next) {
    var BarSerialNumber = req.query.barSerialNumber;
    var args = [BarSerialNumber];
   // Query("QueryBar", args, res);
   QueryNew(request1,"QueryBar", args, res);
}
/*exports.getBarList = async function (req, res, next) {
    let token = req.query.token ? req.query.token : null;
    let BarId = req.query.barId ? req.query.barId : null;
    let args =[token]
  //  QueryAll("GetBarList",agrs,res);
  try{
 const getBars= await QueryAllNew1(request1,"GetBarList", args, res);
 console.log("console getBars ::",getBars.Result)
 let finalData 
 if(token){
    console.log("console in if ::")
 finalData=getBars.Result.filter(item=>{
    return item.Record.barSerialNumber === token
    })
}
 else{
    console.log("console in else ::")
    finalData=getBars.Result.filter(item=>{
        return item.Key === BarId
    })
}
 console.log("getBarlist ::",finalData)
 return res.send({Result:finalData})
  }catch(err){
    console.log("getBarlist err::",err)
return res.send({status:500,error:err})
  }
}*/
exports.getListData = async function(req, res, next){
    let Key = req.query.token ? req.query.token : null;
    let Type = req.query.type ? req.query.type : null;
    let finalData
    let args=[Key,Type]
    console.log("token Key:",Key, " type:",Type)
    try{
    if(Type == 0){
        console.log("in if type =0:",Type)
            const getBars= await QueryAllNew1(request1,"GetBarList",[], res);
            finalData=getBars.Result.filter(item=>{
                return item.Key === Key
                });
    }
    else if(Type == 1){
        console.log("in if type =1:",Type, "args :",args)
        const getBuys= await QueryAllNew1(request1,"GetTransactionBuyList",args, res);
        console.log("result --",getBuys)
        if(getBuys.Result){
        finalData=getBuys.Result.filter(item=>{
            return item.Key === Key
            });
        }else{
            finalData=[]
        }
    }
    else if(Type == 2){
        console.log("in if type =2:",Type)
        const getSells= await QueryAllNew1(request1,"GetTransactionSellList",args,res);
        console.log("in getSells",getSells)
        if(getSells.Result){
        finalData=getSells.Result.filter(item=>{
            return item.Key === Key
            });
        }
            else
            finalData =[]

    }
    else if(Type == 3){
        console.log("in if type =3:",Type)
        const getSends= await QueryAllNew1(request1,"GetTransactionSendList", args, res);
        if(getSends.Result)
        finalData=getSends.Result.filter(item=>{
            return item.Key === Key
        });
        else
        finalData =[]
    }
    else if(Type == 4){
        console.log("in if type =4:",Type)
        const getTrade= await QueryAllNew1(request1,"GetTransactionTradeList", args, res);
        finalData=getTrade.Result.filter(item=>{
            return item.Key === token
            });
    }
    console.log("getBarlist finalData ::",finalData)
    return res.send({Result:finalData })
}catch(err){
    console.log("getBarlist err::",err)
return res.send({status:500,error:err})
}
}
exports.getBar = async function (req, res, next) {
    var BarSerialNumber = req.query.barSerialNumber;
    var args = [BarSerialNumber];

    QueryNew(request1,"GetBar",args,res)
} 
exports.createBuy = async function (req, res, next) {		                  
	var OrderId            =  req.body.orderId;           
    var Amount             =  req.body.amount;        
    var AmountWithFees     =  req.body.amountWithFees;        
    var Stage              =  req.body.stage;         
	var PaymentStatus      =  req.body.paymentStatus;              
	var EstimatedGrams     =  req.body.estimatedGrams;              
    var UserId             =  req.body.userId;       
    var DateCreated = req.body.dateCreated;
    let BuyId               = uuidv4();
    let Type               = type.buy
    let BarIdList               = req.body.barIdList.toString();
    console.log("----req.body.barIdList :",req.body.barIdList,"req.body.barIdList.toString()--",req.body.barIdList.toString()) 
        console.log("BuyId id uuid :",BuyId)
    var args = [BuyId,OrderId,Amount,AmountWithFees,Stage,PaymentStatus,EstimatedGrams,UserId,DateCreated,Type,BarIdList];

    InvokeNew(request1,"CreateBuy", args, res);
}
exports.createBuyWithToken = async function (req, res, next) {		                  
	var OrderId            =  req.body.orderId;           
    var Amount             =  req.body.amount;        
    var AmountWithFees     =  req.body.amountWithFees;        
    var Stage              =  req.body.stage;         
	var PaymentStatus      =  req.body.paymentStatus;              
	var EstimatedGrams     =  req.body.estimatedGrams;              
    var UserId             =  req.body.userId;       
    var DateCreated = req.body.dateCreated;
    let BuyId               = req.body.token; 
    let Type               = type.buy
        console.log("BuyId id uuid :",BuyId)
        let BarIdList               = req.body.barIdList.toString(); 
        console.log("BuyId BarIdList :",BarIdList)
    var args = [BuyId,OrderId,Amount,AmountWithFees,Stage,PaymentStatus,EstimatedGrams,UserId,DateCreated,Type,BarIdList];

    InvokeNew(request1,"CreateBuy", args, res);
}
exports.getBuyList = async function (req, res, next) {
    QueryAllNew(request1,"GetBuyList",res)
}
exports.queryBuy = async function (req, res, next) {
    var OrderId = req.query.orderId;
    var args = [OrderId];

    QueryNew(request1,"QueryBuy", args, res);
}
exports.getBuy = async function (req, res, next) {
    var OrderId = req.query.orderId;
    var args = [OrderId];
try{
   let data= await QueryNew1(request1,"GetBuy",args,res)
    for(let i=0;i<data.Result.length;i++){
        data.Result[i].Record.barIdList.split(',');
    }
   console.log("--->>",data)
   return res.send(data)
}catch(err){
    return res.send(err)
}
} 
exports.createSell = async function (req, res, next) {		                  
	var OrderId            =  req.body.orderId;           
    var Grams             =  req.body.grams;                      
	var EstimatedAmount     =  req.body.estimatedAmount;              
    var UserId             =  req.body.userId;       
    var DateCreated = req.body.dateCreated;
    let SellId               = uuidv4()
    let Type               = type.sell
    // let BarIdList               = req.body.barIdList; 
    console.log("BuyId id uuid :",SellId)
    let BarIdList               = req.body.barIdList.toString(); 
    console.log("----req.body.barIdList :",req.body.barIdList) 
    var args = [SellId,OrderId,Grams,EstimatedAmount,UserId,DateCreated,Type,BarIdList];

    InvokeNew(request1,"CreateSell", args, res);
}
exports.createSellWithToken = async function (req, res, next) {		                  
	var OrderId            =  req.body.orderId;           
    var Grams             =  req.body.grams;                      
	var EstimatedAmount     =  req.body.estimatedAmount;              
    var UserId             =  req.body.userId;       
    var DateCreated = req.body.dateCreated;
    let SellId               = req.body.token;
    let Type               = type.sell
    console.log("BuyId id uuid :",SellId)
    let BarIdList               = req.body.barIdList.toString(); 
    console.log("--",BarIdList)
    var args = [SellId,OrderId,Grams,EstimatedAmount,UserId,DateCreated,Type,BarIdList];

    InvokeNew(request1,"CreateSell", args, res);
}
exports.getSellList = async function (req, res, next) {
    QueryAllNew(request1,"GetSellList",res)
}
exports.querySell = async function (req, res, next) {
    var OrderId = req.query.orderId;
    var args = [OrderId];

    QueryNew(request1,"QuerySell", args, res);
}
exports.getSell = async function (req, res, next) {
    var OrderId = req.query.orderId;
    var args = [OrderId];

    //QueryNew(request1,"GetSell",args,res)
    try{
        let data= await QueryNew1(request1,"GetSell",args,res)
         for(let i=0;i<data.Result.length;i++){
             data.Result[i].Record.barIdList.split(',');
         }
        console.log("--->>",data)
        return res.send(data)
     }catch(err){
         return res.send(err)
     }
} 
exports.createSend = async function (req, res, next) {		                  
	var OrderId            =  req.body.orderId;           
    var Grams              =  req.body.grams;                      
	var SenderUserId       =  req.body.senderUserId;              
    var ReceiverUserId     =  req.body.receiverUserId;       
    var DateCreated = req.body.dateCreated;
let SendId= uuidv4()
let Type               = type.send
let BarIdList               = req.body.barIdList.toString(); 
console.log("sendId uuid :",SendId)
console.log("send BarIdList :",BarIdList," ::",req.body.barIdList)
    var args = [SendId,OrderId,Grams,SenderUserId,ReceiverUserId,DateCreated,Type,BarIdList];

    InvokeNew(request1,"CreateSend", args, res);
}
exports.createSendWithToken = async function (req, res, next) {		                  
	var OrderId            =  req.body.orderId;           
    var Grams              =  req.body.grams;                      
	var SenderUserId       =  req.body.senderUserId;              
    var ReceiverUserId     =  req.body.receiverUserId;       
    var DateCreated = req.body.dateCreated;
let SendId= req.body.token
let Type               = type.send

let BarIdList               = req.body.barIdList.toString(); 
console.log("sendId uuid :",SendId)
console.log("send BarIdList :",BarIdList," ::",req.body.barIdList)
    var args = [SendId,OrderId,Grams,SenderUserId,ReceiverUserId,DateCreated,Type,BarIdList];

    InvokeNew(request1,"CreateSend", args, res);
}
exports.getSendList = async function (req, res, next) {
    QueryAllNew(request1,"GetSendList",res)
}
exports.querySend = async function (req, res, next) {
    var OrderId = req.query.orderId;
    var args = [OrderId];

    QueryNew(request1,"QuerySend", args, res);
}
exports.getSend = async function (req, res, next) {
    var OrderId = req.query.orderId;
    var args = [OrderId];

    //QueryNew(request1,"GetSend",args,res)
    try{
        let data= await QueryNew1(request1,"GetSend",args,res)
         for(let i=0;i<data.Result.length;i++){
             data.Result[i].Record.barIdList.split(',');
         }
        console.log("--->>",data)
        return res.send(data)
     }catch(err){
         return res.send(err)
     }
} 
exports.createTrade = async function (req, res, next) {		                  
	var OrderId            =  req.body.orderId;           
    var Grams              =  req.body.grams;                      
	var UserId       =  req.body.userId;                   
    let TradeId =uuidv4()
    let Type               = type.trade
console.log("trade id uuid ",TradeId)
    var args = [TradeId,OrderId,Grams,UserId,Type];

    InvokeNew(request1,"CreateTrade", args, res);
}
exports.getTradeList = async function (req, res, next) {
    QueryAllNew(request1,"GetTradeList",res)
}
exports.queryTrade = async function (req, res, next) {
    var OrderId = req.query.orderId;
    var args = [OrderId];

    QueryNew(request1,"QueryTrade", args, res);
}
exports.getTrade = async function (req, res, next) {
    var OrderId = req.query.orderId;
    var args = [OrderId];

    QueryNew(request1,"GetTrade",args,res)
}
async function Invoke(funcName,args,res){
    try {
        // load the network configuration
      //  const ccpPath = path.resolve(__dirname, '..', '..', 'first-network', 'connection-org1.json');
        let ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get('user1');
        if (!identity) {
            console.log(`An identity for the user user1 does not exist in the wallet`);
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork(channelName);

        // Get the contract from the network.
        const contract = network.getContract(chaincodeName);
        if(args.length == 6 ){
       
        await contract.submitTransaction(funcName,args[0],args[1],args[2],args[3],args[4],args[5]);
    }else if(args.length == 7){
     
        await contract.submitTransaction(funcName,args[0],args[1],args[2],args[3],args[4],args[5],args[6]);   
    } else if(args.length == 5){
     
        await contract.submitTransaction(funcName,args[0],args[1],args[2],args[3],args[4]);   
    }else if(args.length == 4){
     
        await contract.submitTransaction(funcName,args[0],args[1],args[2],args[3]);   
    }else if(args.length == 3){
     
        await contract.submitTransaction(funcName,args[0],args[1],args[2]);   
    }else if(args.length == 2){
     
        await contract.submitTransaction(funcName,args[0],args[1]);   
    }else if(args.length == 1){
     
        await contract.submitTransaction(funcName,args[0]);   
    }   
        console.log({message:'Success'});
        res.send({message:'Success'});

        // Disconnect from the gateway.
        await gateway.disconnect();

    } catch (error) {
        console.error(`failure: ${error}`);
        res.send(`failure: ${error}`);

    }
}
async function QueryAll(funcName,res){
    try {
        // load the network configuration
        //const ccpPath = path.resolve(__dirname, '..', '..', 'first-network', 'connection-org1.json');
        const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get('user1');
        if (!identity) {
            console.log(`An identity for the user user1does not exist in the wallet`);
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork(channelName);

        // Get the contract from the network.
        const contract = network.getContract(chaincodeName);

        // Evaluate the specified transaction.
        // queryCar transaction - requires 1 argument, ex: ('queryCar', 'CAR4')
        // queryAllCars transaction - requires no arguments, ex: ('queryAllCars')
    
        const result = await contract.evaluateTransaction(funcName);
let p = JSON.parse(result)
        console.log({Result:p});
        res.send({Result:p});

    } catch (error) {
        console.error(`failure: ${error}`);
        res.send(`failure: ${error}`);

    }
}
async function Query(funcName,args,res){
    try {
        // load the network configuration
        //const ccpPath = path.resolve(__dirname, '..', '..', 'first-network', 'connection-org1.json');
        const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get('user1');
        if (!identity) {
            console.log(`An identity for the user user1does not exist in the wallet`);
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork(channelName);

        // Get the contract from the network.
        const contract = network.getContract(chaincodeName);

        // Evaluate the specified transaction.
        // queryCar transaction - requires 1 argument, ex: ('queryCar', 'CAR4')
        // queryAllCars transaction - requires no arguments, ex: ('queryAllCars')
        if(args.length == 1 ){
       
        const result = await contract.evaluateTransaction(funcName,args[0]);
        let p = JSON.parse(result)
        console.log({Result:p});
        res.send({Result:p});
        }else if(args.length == 6 ){
       
            const result = await contract.evaluateTransaction(funcName,args[0],args[1],args[2],args[3],args[4],args[5]);
            let p = JSON.parse(result)
            console.log({Result:p});
            res.send({Result:p});
        }else if(args.length == 4 ){
       
            const result = await contract.evaluateTransaction(funcName,args[0],args[1],args[2],args[3]);
            let p = JSON.parse(result)
            console.log({Result:p});
            res.send({Result:p});
        }else if(args.length == 5 ){
       
            const result = await contract.evaluateTransaction(funcName,args[0],args[1],args[2],args[3],args[4]);
            let p = JSON.parse(result)
            console.log({Result:p});
            res.send({Result:p});
        }else if(args.length == 3 ){
       
            const result = await contract.evaluateTransaction(funcName,args[0],args[1],args[2]);
            let p = JSON.parse(result)
            console.log({Result:p});
            res.send({Result:p});
        }else if(args.length == 2 ){
       
            const result = await contract.evaluateTransaction(funcName,args[0],args[1]);
            let p = JSON.parse(result)
            console.log({Result:p});
            res.send({Result:p});
        }
    } catch (error) {
        console.error(`failure: ${error}`);
        res.send(`failure: ${error}`);

    }
}