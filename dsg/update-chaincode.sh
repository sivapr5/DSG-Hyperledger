#!/bin/bash


CC_PKG_NAME="dsg3" # you have to change this on every upgradation
version="3"  # you have to change this on every upgradation
label="dsg3"  # you have to change this on every upgradation
channel="mychannel"
sequence="3"   # you have to change this on every upgradation
name="dsg"  # do not change this as at start of network we gave this
#PACKAGE_ID=""

CC_SRC_PATH=github.com/hyperledger/fabric-samples/chaincode/dsg/go
CORE_PEER_MSPCONFIGPATH_ORG1=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
CORE_PEER_MSPCONFIGPATH_ORG2=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
ORDERER_TLS_ROOTCERT_FILE= /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

echo "----------------------  STEP 1 ---------------------------- \n"
echo "Packaging of chaincode"
docker exec cli peer lifecycle chaincode package ${CC_PKG_NAME}.tar.gz --path  "$CC_SRC_PATH"  --label ${label}

echo "--------------------------------  STEP 2 ----------------------------------------------- \n"
echo "===== install chaincode package  to peer0.org1 ====="
echo  ${CC_PKG_NAME}
echo  ${CORE_PEER_MSPCONFIGPATH_ORG1}
docker exec cli peer lifecycle chaincode install ${CC_PKG_NAME}.tar.gz
echo "==== install chaincode package to peer1.org1 ======"
docker exec -e CORE_PEER_MSPCONFIGPATH=${CORE_PEER_MSPCONFIGPATH_ORG1} -e CORE_PEER_ADDRESS=peer1.org1.example.com:8051 -e CORE_PEER_LOCALMSPID="Org1MSP" -e CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt cli peer lifecycle chaincode install ${CC_PKG_NAME}.tar.gz
echo "==== install chaincode package  to peer0.org2 ===="
docker exec -e CORE_PEER_MSPCONFIGPATH=${CORE_PEER_MSPCONFIGPATH_ORG2} -e CORE_PEER_ADDRESS=peer0.org2.example.com:9051 -e CORE_PEER_LOCALMSPID="Org2MSP" -e CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt cli peer lifecycle chaincode install ${CC_PKG_NAME}.tar.gz
echo "==== install chaincode package  to peer1.org2 ===="
docker exec -e CORE_PEER_MSPCONFIGPATH=${CORE_PEER_MSPCONFIGPATH_ORG2} -e CORE_PEER_ADDRESS=peer1.org2.example.com:10051 -e CORE_PEER_LOCALMSPID="Org2MSP" -e CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt cli peer lifecycle chaincode install ${CC_PKG_NAME}.tar.gz

echo "---------------------------------  STEP 3 ----------------------------------"
echo "Determining package ID for smart contract on peer0.org1.example.com"
docker exec cli peer lifecycle chaincode queryinstalled >&log.txt
 cat log.txt
PACKAGE_ID=$(sed -n "/${label}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
echo PackageID is ${PACKAGE_ID}
  
	

echo "----------------------------------  STEP 4 --------------------------"
echo "========= Approve of chaincode for Org 1 ==================="
docker exec cli peer lifecycle chaincode approveformyorg --tls --cafile ${ORDERER_TLS_ROOTCERT_FILE}  --channelID ${channel} --name ${name} --version ${version} --sequence ${sequence} --signature-policy "OR('Org1MSP.member','Org2MSP.member')"  --waitForEvent --package-id ${PACKAGE_ID}

echo "=============Approve of chaincode for Org 2============"
docker exec -e CORE_PEER_MSPCONFIGPATH=${CORE_PEER_MSPCONFIGPATH_ORG2} -e CORE_PEER_ADDRESS=peer0.org2.example.com:9051 -e CORE_PEER_LOCALMSPID="Org2MSP" -e CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt cli peer lifecycle chaincode approveformyorg --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --channelID ${channel} --name ${name} --version ${version} --sequence ${sequence} --signature-policy "OR('Org1MSP.member','Org2MSP.member')"  --waitForEvent --package-id ${PACKAGE_ID}

echo "-------------------------  STEP 5 -----------------------"
echo "Commit the New Chaincode Definition"
docker exec cli peer lifecycle chaincode commit -o orderer.example.com:7050 --tls --cafile ${ORDERER_TLS_ROOTCERT_FILE} --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt --channelID ${channel} --name ${name} --version ${version}  --signature-policy "OR('Org1MSP.member','Org2MSP.member')" --sequence ${sequence}

echo "Test the newly commited chaincode definition"
docker exec cli peer lifecycle chaincode querycommitted -o orderer.example.com:7050 --tls --cafile ${ORDERER_TLS_ROOTCERT_FILE} --channelID ${channel} --name ${name} --output json
