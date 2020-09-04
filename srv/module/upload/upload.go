package upload

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"srv/core/util"

	"github.com/gin-gonic/gin"
)

// Upload ...
type Upload struct {
	Data []UploadItem
}

// UploadItem ...
type UploadItem struct {
	Name string
	Msg  string
}

// SrvUpload ...
func SrvUpload(c *gin.Context) {
	var err error

	//
	data := map[string]interface{}{
		"status": true,
	}

	// TODO: Check the user isLogin(sessionID string)
	// Get sessionID from cookie
	var sessionID string

	if sessionID, err = c.Cookie("sessionID"); err != nil {
		log.Printf("Error: %s", err.Error())
		data["status"] = false
		c.SecureJSON(http.StatusOK, data)
		return
	}

	log.Printf("Here sessionID: %s", sessionID)

	// single file
	var file *multipart.FileHeader

	if file, err = c.FormFile("fileKey"); err != nil {
		log.Printf("Error: %s", err.Error())
		data["status"] = false
		c.SecureJSON(http.StatusOK, data)
		return
	}

	log.Printf("Filename: %s", file.Filename)

	// Never ever trust user's input (filename), so we hashed the filename
	filenameHashEncoded := util.HashString([]byte(file.Filename))

	log.Printf("filenameHashEncoded: %s", filenameHashEncoded)

	// Upload the file to specific dst.
	dst := fmt.Sprintf(
		"/home/shiyin/goproj/exp-app/upload/%s",
		filenameHashEncoded,
	)

	if err = c.SaveUploadedFile(file, dst); err != nil {
		log.Printf("Error: %s", err.Error())
		data["status"] = false
		c.SecureJSON(http.StatusOK, data)
		return
	}

	// Read the file content
	var content []byte

	if content, err = ioutil.ReadFile(dst); err != nil {
		log.Printf("Error: %s", err.Error())
		data["status"] = false
		c.SecureJSON(http.StatusOK, data)
		return
	}

	log.Printf("Here content: %s", content)

	// Conver the JSON string to an object
	upload := Upload{}

	if err = json.Unmarshal(content, &upload); err != nil {
		log.Printf("JSON Error: %s", err.Error())
		data["status"] = false
		c.SecureJSON(http.StatusOK, data)
		return
	}

	// Insert items to database
	if err = InsertUploadItems(upload); err != nil {
		log.Printf("JSON Error: %s", err.Error())
		data["status"] = false
		c.SecureJSON(http.StatusOK, data)
		return
	}

	//
	c.SecureJSON(http.StatusOK, data)
}

// InsertUploadItems ...
func InsertUploadItems(upload Upload) error {
	// Prepare statement

	// Loop through the upload item, Exec
	for k, v := range upload.Data {
		log.Printf("item %d: %#v", k, v)
	}

	return nil
}
