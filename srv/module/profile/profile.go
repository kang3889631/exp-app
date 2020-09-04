package profile

import (
	"fmt"
	"log"
	"net/http"
	"srv/model"

	"github.com/gin-gonic/gin"
)

// PostBasic ...
func PostBasic(c *gin.Context) {
	//data := map[string]interface{}{
	//	"status": false,
	//}

	//
	var err error
	var profileBasic model.ProfileBasic

	// Client will submit JSON data, then, we are trying to convert JSON
	// string into "model.Group" (live object)
	if err = c.ShouldBindJSON(&profileBasic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"MY error": err.Error(),
		})
		return
	}

	log.Printf("User input: %#v", profileBasic)
	log.Printf("User firstname: %s", profileBasic.Firstname)
	log.Printf("User lastname: %s", profileBasic.Lastname)

	// Data
	data := map[string]interface{}{
		"profileBasic": fmt.Sprintf("You input %#v", profileBasic),
	}

	// Return JSON back to client
	c.SecureJSON(http.StatusOK, data)
}
