package app

import (
	"crud-dasar-go-2/controller"
	"crud-dasar-go-2/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(kamarController controller.KamarController, barangController controller.BarangController, authController controller.AuthController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/auth/login", authController.Login)
	router.POST("/api/auth/register", authController.Register)

	router.GET("/api/kamars", kamarController.FindAll)
	router.GET("/api/kamars/:kamarId", kamarController.FindById)
	router.POST("/api/kamars", kamarController.Create)
	router.PUT("/api/kamars/:kamarId", kamarController.Update)
	router.DELETE("/api/kamars/:kamarId", kamarController.Delete)

	router.GET("/api/barangs", barangController.FindAll)
	router.GET("/api/barangs/:barangId", barangController.FindById)
	router.POST("/api/barangs", barangController.Create)
	router.PUT("/api/barangs/:barangId", barangController.Update)
	router.DELETE("/api/barangs/:barangId", barangController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
