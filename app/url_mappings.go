package app

import (
	"mvc/controllers"
	//"net/http"
)

func mapUrls(){
	router.GET("/users/:user_id", controllers.GetUser)

}