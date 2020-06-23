package main


const (

	SUCCESS_RESPONSE                 = `{"status": 200}`
	GOLDBAR_CREATE_SUCCESS           = `{"status": 200, "message":"New GoldBar created successfully"}`
	GOLDBAR_UPDATED_SUCCESS          = `{"status": 200, "message":"GoldBar details updated successfully"}`
	
)


const (

	GOLDBAR_CREATE_FAILED    = `{"status" : 402, "message":"Error in creating New GoldBar"}`
	GOLDBAR_ALREADY_EXIST = `{"status": 403, "message":"GoldBar already exist"}`
	FAILED_UPDATE_GOLDBAR   = `{"status":402, "message": "Error in Updating GoldBar"}`
	GOLDBAR_NOT_FOUND   = `{"status":404, "message": "No GoldBar found "}`

	GOLDBAR_DOESNT_EXIST = `{"status" : 404, "message":"GoldBar ID doesn't exist"}`

	ERROR_RETRIEVE_GOLDBAR = `{"status" : 401, "message" : "Error in retrieving GoldBar with GoldBarID<UID>"}`

)
