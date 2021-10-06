1. move prerequisite all folders inside nfs shared directory & follow deployments for first 7 folders.

2. after 7th step, do following:
    enter in shell of peer0 of org1 cli
    - ./scripts/createAppChannel.sh
    for all possible orgs,
    - peer channel join -b ./channel-artifacts/mychannel.block

3. create packages on nfs server
    - cd export/chaincode/basic/packaging
    - tar cfz code.tar.gz connection.json 
    - tar cfz basic-org1.tgz code.tar.gz metadata.json
    - rm -rf code.tar.gz
    - nano connection.json & update address for peer 2
    - tar cfz code.tar.gz connection.json
    - tar cfz basic-org2.tgz code.tar.gz metadata.json

4. enter shell of cli peer0 of org1
    - cd /opt/gopath/src/github.com/chaincode/basic/packaging
    - peer lifecycle chaincode install basic-org1.tgz (you can run - peer lifecycle chaincode queryinstalled & get package id)
    - do same on shell of cli peer0 of org2 for basic-org2.tgz

5. Create docker image for own chaincode 8.chaincode directory.

6. Update image & env from deployment yaml for 9.cc-deploy directory

7. Run approveformyorg, commitreadiness & commit commands for both orgs

8. Create docker image for node sdk api, and providing same image deploy api service.

9. Run script/ccp.sh from nfs server
