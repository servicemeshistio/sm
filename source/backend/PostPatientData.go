package backend

import (
	"fmt"
	"sm/source/models"
	"sm/source/restapi/operations"
	"sm/source/utils"

	"github.com/go-openapi/runtime/middleware"
)

//PostPatientDataHandler is the controller logic for the POST api
func PostPatientDataHandler() operations.PostPatientDataHandlerFunc {
	return operations.PostPatientDataHandlerFunc(func(params operations.PostPatientDataParams) middleware.Responder {

		//Db object which we initialized earlier
		server := utils.ServerInfo

		//Create the table if it doesn't exists
		if err := server.Db.AutoMigrate(&models.PatientObject{}).Error; err != nil {
			fmt.Println("Unable to create the table", err)
			resp := operations.NewPostPatientDataBadRequest()
			respModel := models.ErrorResponse{}
			respModel.Message = "Unable to crate the table"
			respModel.MessageCode = "400"
			resp.SetPayload(&respModel)
			return resp
		}

		PatientObject := params.Body
		PatientObject.ID = params.PatientID

		//Enter info inside the database
		if err := server.Db.Create(PatientObject).Error; err != nil {
			fmt.Println("Unable to write the info to database", err)
			resp := operations.NewPostPatientDataBadRequest()
			respModel := models.ErrorResponse{}
			respModel.Message = "Successfully Added record into the database"
			respModel.MessageCode = "200"
			resp.SetPayload(&respModel)
			return resp
		}

		//If everything is good then we send the success response.
		resp := operations.NewPostPatientDataOK()
		respModel := models.SuccessResponse{}
		respModel.Message = "Successfully Added record into the database"
		respModel.MessageCode = "200"
		resp.SetPayload(&respModel)
		return resp
	})

}
