package controllers

import (
	"log"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter,req *http.Request){
	userId,err := (strconv.ParseInt(req.URL.Query().Get("user_id"),10,64))
	if err != nil {
		// just return the bad request to the client
		return
	}
//	user, err : =  
	log.Printf("about to process user_id %v", userId)
}