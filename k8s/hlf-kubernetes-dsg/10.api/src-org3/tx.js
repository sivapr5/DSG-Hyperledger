const { getCCP } = require("./buildCCP");
const { Wallets, Gateway } = require('fabric-network');
const path = require("path");
const walletPath = path.join(__dirname, "wallet-org3");
const {buildWallet} =require('./AppUtils')


/*
{
    org:Org1MSP,
    channelName:"mychannel",
    chaincodeName:"basic",
    userId:"aditya"
    data:{
        id:"asset1",
        color:"red",
        size:5,
        appraisedValue:200,
        owner:"TOM"
    }
}

*/
exports.createAsset = async (request) => {
    let org = request.org;
    let num = Number(org.match(/\d/g).join(""));
    const ccp = getCCP(num);

    const wallet = await buildWallet(Wallets, walletPath);

    const gateway = new Gateway();

    await gateway.connect(ccp, {
        wallet,
        identity: request.userId,
        discovery: { enabled: true, asLocalhost: false } // using asLocalhost as this gateway is using a fabric network deployed locally
    });

    // Build a network instance based on the channel where the smart contract is deployed
    const network = await gateway.getNetwork(request.channelName);

    // Get the contract from the network.
    const contract = network.getContract(request.chaincodeName);
    let data=request.data;
    let result = await contract.submitTransaction('CreateAsset',data.ID,data.color,data.size,data.owner,data.appraisedValue);
    return (result);
}


/*
{
    org:Org1MSP,
    channelName:"mychannel",
    chaincodeName:"basic",
    userId:"aditya"
    data:{
        id:"asset1",
        color:"red",
        size:5,
        appraisedValue:200,
        owner:"TOM"
    }
}

*/
exports.updateAsset=async (request) => {
    let org = request.org;
    let num = Number(org.match(/\d/g).join(""));
    const ccp = getCCP(num);

    const wallet = await buildWallet(Wallets, walletPath);

    const gateway = new Gateway();

    await gateway.connect(ccp, {
        wallet,
        identity: request.userId,
        discovery: { enabled: true, asLocalhost: false } // using asLocalhost as this gateway is using a fabric network deployed locally
    });

    // Build a network instance based on the channel where the smart contract is deployed
    const network = await gateway.getNetwork(request.channelName);

    // Get the contract from the network.
    const contract = network.getContract(request.chaincodeName);
    let data=request.data;
    let result = await contract.submitTransaction('UpdateAsset',data.ID,data.color,data.size,data.owner,data.appraisedValue);
    return (result);
}



/*
{
    org:Org1MSP,
    channelName:"mychannel",
    chaincodeName:"basic",
    userId:"aditya"
    data:{
        id:"asset1",
        color:"red",
        size:5,
        appraisedValue:200,
        owner:"TOM"
    }
}
*/
exports.deleteAsset=async (request) => {
    let org = request.org;
    let num = Number(org.match(/\d/g).join(""));
    const ccp = getCCP(num);

    const wallet = await buildWallet(Wallets, walletPath);

    const gateway = new Gateway();

    await gateway.connect(ccp, {
        wallet,
        identity: request.userId,
        discovery: { enabled: true, asLocalhost: false } // using asLocalhost as this gateway is using a fabric network deployed locally
    });

    // Build a network instance based on the channel where the smart contract is deployed
    const network = await gateway.getNetwork(request.channelName);

    // Get the contract from the network.
    const contract = network.getContract(request.chaincodeName);
    let data=request.data;
    let result = await contract.submitTransaction('DeleteAsset',data.id);
    return (result);
}


/*
{
    org:Org1MSP,
    channelName:"mychannel",
    chaincodeName:"basic",
    userId:"aditya"
    data:{
        id:"asset1",
        newOwner:"TOM"
    }
}
*/
exports.TransferAsset=async (request) => {
    let org = request.org;
    let num = Number(org.match(/\d/g).join(""));
    const ccp = getCCP(num);

    const wallet = await buildWallet(Wallets, walletPath);

    const gateway = new Gateway();

    await gateway.connect(ccp, {
        wallet,
        identity: request.userId,
        discovery: { enabled: true, asLocalhost: false } // using asLocalhost as this gateway is using a fabric network deployed locally
    });

    // Build a network instance based on the channel where the smart contract is deployed
    const network = await gateway.getNetwork(request.channelName);

    // Get the contract from the network.
    const contract = network.getContract(request.chaincodeName);
    let data=request.data;
    let result = await contract.submitTransaction('TransferAsset',data.id,data.newOwner);
    return JSON.parse(result);
}

/*exports.InvokeNew=async (request,funcName,args,res)=>{
    try {
       
        let org = request.org;
    let num = Number(org.match(/\d/g).join(""));
    const ccp = getCCP(num);

    const wallet = await buildWallet(Wallets, walletPath);

    const gateway = new Gateway();

    await gateway.connect(ccp, {
        wallet,
        identity: request.userId,
        discovery: { enabled: true, asLocalhost: false } // using asLocalhost as this gateway is using a fabric network deployed locally
    });

    // Build a network instance based on the channel where the smart contract is deployed
    const network = await gateway.getNetwork(request.channelName);

    // Get the contract from the network.
    const contract = network.getContract(request.chaincodeName);
    // let data=request.data;       
    
    if(args.length == 6 ){
       
        await contract.submitTransaction(funcName,args[0],args[1],args[2],args[3],args[4],args[5]);
    }
    else if(args.length == 11){
     console.log("11 paramss")
        await contract.submitTransaction(funcName,args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7],args[8],args[9],args[10]);   
    }
    else if(args.length == 10){
     
        await contract.submitTransaction(funcName,args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7],args[8],args[9]);   
    }
    else if(args.length == 9){
     
        await contract.submitTransaction(funcName,args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7],args[8]);   
    }
    else if(args.length == 8){
     
        await contract.submitTransaction(funcName,args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7]);   
    } 
    else if(args.length == 7){
     
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
}*/

exports.QueryAllNew=async (request,funcName,res)=>{
    try {
        let org = request.org;
    let num = Number(org.match(/\d/g).join(""));
    const ccp = getCCP(num);

    const wallet = await buildWallet(Wallets, walletPath);

    const gateway = new Gateway();

    await gateway.connect(ccp, {
        wallet,
        identity: request.userId,
        discovery: { enabled: true, asLocalhost: false } // using asLocalhost as this gateway is using a fabric network deployed locally
    });

    // Build a network instance based on the channel where the smart contract is deployed
    const network = await gateway.getNetwork(request.channelName);

    // Get the contract from the network.
    const contract = network.getContract(request.chaincodeName);
    // let data=request.data;
    
        const result = await contract.evaluateTransaction(funcName);
let p = JSON.parse(result)
        console.log({Result:p});
        res.send({Result:p});

    } catch (error) {
        console.error(`failure: ${error}`);
        res.send(`failure: ${error}`);

    }
}

exports.QueryAllNew1=async (request,funcName,args,res)=>{
    try {
        let org = request.org;
    let num = Number(org.match(/\d/g).join(""));
    const ccp = getCCP(num);

    const wallet = await buildWallet(Wallets, walletPath);

    const gateway = new Gateway();

    await gateway.connect(ccp, {
        wallet,
        identity: request.userId,
        discovery: { enabled: true, asLocalhost: false } // using asLocalhost as this gateway is using a fabric network deployed locally
    });

    // Build a network instance based on the channel where the smart contract is deployed
    const network = await gateway.getNetwork(request.channelName);

    // Get the contract from the network.
    const contract = network.getContract(request.chaincodeName);
    // let data=request.data;
    let result
    if(args.length == 0 ){
        result = await contract.evaluateTransaction(funcName);
    }
    if(args.length == 1 ){
        result = await contract.evaluateTransaction(funcName,args[0]);
    }else if(args.length == 2){
        console.log("in else if :",args)
         result = await contract.evaluateTransaction(funcName,args[0],args[2]);
    }
    console.log("result --from tx",result);
        //const result = await contract.evaluateTransaction(funcName);
        if(result){
       let p = JSON.parse(result)
        console.log({Result:p});
        return({Result:p});
        }else{
            return({Result:[]});
        }
    } catch (error) {
        console.error(`failure: ${error}`);
         res.send(`failure: ${error}`);
        //return

    }
}
exports.QueryNew=async (request,funcName,args,res)=>{
    try {
        let org = request.org;
        let num = Number(org.match(/\d/g).join(""));
        const ccp = getCCP(num);
    
        const wallet = await buildWallet(Wallets, walletPath);
    
        const gateway = new Gateway();
    
        await gateway.connect(ccp, {
            wallet,
            identity: request.userId,
            discovery: { enabled: true, asLocalhost: false } // using asLocalhost as this gateway is using a fabric network deployed locally
        });
    
        // Build a network instance based on the channel where the smart contract is deployed
        const network = await gateway.getNetwork(request.channelName);
    
        // Get the contract from the network.
        const contract = network.getContract(request.chaincodeName);
        // let data=request.data;
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

exports.QueryNew1=async (request,funcName,args,res)=>{
    try {
        let org = request.org;
        let num = Number(org.match(/\d/g).join(""));
        const ccp = getCCP(num);
    
        const wallet = await buildWallet(Wallets, walletPath);
    
        const gateway = new Gateway();
    
        await gateway.connect(ccp, {
            wallet,
            identity: request.userId,
            discovery: { enabled: true, asLocalhost: false } // using asLocalhost as this gateway is using a fabric network deployed locally
        });
    
        // Build a network instance based on the channel where the smart contract is deployed
        const network = await gateway.getNetwork(request.channelName);
    
        // Get the contract from the network.
        const contract = network.getContract(request.chaincodeName);
        // let data=request.data;
        if(args.length == 1 ){
       
        const result = await contract.evaluateTransaction(funcName,args[0]);
        let p = JSON.parse(result)
        console.log({Result:p});
        return({Result:p})
       // res.send({Result:p});
        }else if(args.length == 6 ){
       
            const result = await contract.evaluateTransaction(funcName,args[0],args[1],args[2],args[3],args[4],args[5]);
            let p = JSON.parse(result)
            console.log({Result:p});
            return({Result:p})
           // res.send({Result:p});
        }else if(args.length == 4 ){
       
            const result = await contract.evaluateTransaction(funcName,args[0],args[1],args[2],args[3]);
            let p = JSON.parse(result)
            console.log({Result:p});
            return({Result:p})
           // res.send({Result:p});
        }else if(args.length == 5 ){
       
            const result = await contract.evaluateTransaction(funcName,args[0],args[1],args[2],args[3],args[4]);
            let p = JSON.parse(result)
            console.log({Result:p});
            res.send({Result:p});
        }else if(args.length == 3 ){
       
            const result = await contract.evaluateTransaction(funcName,args[0],args[1],args[2]);
            let p = JSON.parse(result)
            console.log({Result:p});
            return({Result:p})
           // res.send({Result:p});
        }else if(args.length == 2 ){
       
            const result = await contract.evaluateTransaction(funcName,args[0],args[1]);
            let p = JSON.parse(result)
            console.log({Result:p});
            return({Result:p})
           // res.send({Result:p});
        }
    } catch (error) {
        console.error(`failure: ${error}`);
        res.send(`failure: ${error}`);

    }
}