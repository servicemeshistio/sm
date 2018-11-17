package backend

import "sm/source/restapi/operations"
import "github.com/go-openapi/runtime/middleware"

//PostPatientDataHandler is the controller logic for the POST api
func PostPatientDataHandler() operations.PostPatientDataHandlerFunc {
	return operations.PostPatientDataHandlerFunc(func(params operations.PostPatientDataParams) middleware.Responder {

		resp := operations.NewPostPatientDataOK()
		return resp
	})

}
