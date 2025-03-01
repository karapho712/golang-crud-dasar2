package app

import (
	"crud-dasar-go-2/controller"
	"crud-dasar-go-2/exception"
	"crud-dasar-go-2/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(kamarController controller.KamarController, barangController controller.BarangController, authController controller.AuthController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/auth/login", authController.Login)
	router.POST("/api/auth/register", authController.Register)

	router.GET("/api/kamars", middleware.JWTMiddleware(kamarController.FindAll))
	router.GET("/api/kamars/:kamarId", middleware.JWTMiddleware(kamarController.FindById))
	router.POST("/api/kamars", middleware.JWTMiddleware(kamarController.Create))
	router.PUT("/api/kamars/:kamarId", middleware.JWTMiddleware(kamarController.Update))
	router.DELETE("/api/kamars/:kamarId", middleware.JWTMiddleware(kamarController.Delete))

	router.GET("/api/barangs", middleware.JWTMiddleware(barangController.FindAll))
	router.GET("/api/barangs/:barangId", middleware.JWTMiddleware(barangController.FindById))
	router.POST("/api/barangs", middleware.JWTMiddleware(barangController.Create))
	router.PUT("/api/barangs/:barangId", middleware.JWTMiddleware(barangController.Update))
	router.DELETE("/api/barangs/:barangId", middleware.JWTMiddleware(barangController.Delete))

	router.PanicHandler = exception.ErrorHandler

	return router
}
