package controllers

import (
	"mvc/services"
	"mvc/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, err := (strconv.ParseInt(c.Param("user_id"), 10, 64))
	if err != nil {

		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.Respond(c, apiErr.StatusCode, apiErr)
		return
	}
	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c,apiErr)
		// handle the err and return
		return
	}
	// return user to client
	utils.Respond(c, http.StatusOK, user)
}
