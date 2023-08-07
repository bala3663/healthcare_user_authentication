package routes

import (
	controller "Final-Project-gin_helthcare/controllers"
	"Final-Project-gin_helthcare/middleware"

	"github.com/gin-gonic/gin"
)

func Patient_info_routes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())

	incomingRoutes.GET("/Get_patients1", controller.Get_patients1())

	incomingRoutes.POST("/Post_patient1", controller.Post_patient1())

	incomingRoutes.PUT("/Update_patientinfo1", controller.Update_patientinfo1())
	incomingRoutes.DELETE("/Delete_Patient1", controller.Delete_Patient1())
	incomingRoutes.GET("/Search_patient_by_bill_id1", controller.Search_patient_by_bill_id1())
}
