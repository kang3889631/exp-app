package main

import (
	"log"

	"srv/core/config"
	"srv/core/gdo"
	"srv/core/websrv"
	"srv/module/router"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/sync/errgroup"
)

func main() {
	var err error
	var eg *errgroup.Group

	defer func(e *error) {
		// TODO:
		log.Printf("Error: %s\n", (*e).Error())
	}(&err)

	// Initialize config
	if err = config.Init(); err != nil {
		return
	}

	// Initialize (Go data object) database server
	if err = gdo.Init(); err != nil {
		return
	}

	// Initialize web server
	eg = websrv.Init(
		// Define a list of routes
		router.HTTPSAll(),
	)

	// Ready to server the HTTP requests
	if err = eg.Wait(); err != nil {
		return
	}
}

//func mainNotYet() {
//	// Read flag from user arguments from command
//	srvportPtr := flag.String("srvport", ":8080", "specify a port for the server to listen on")
//
//	flag.Parse()
//
//	//
//	router := gin.Default()
//
//	// Set a lower memory limit for upload multipart forms (default is 32 MiB)
//	router.MaxMultipartMemory = 8 << 20 // 8 MiB
//
//	// Root path
//	router.GET("/", func(c *gin.Context) {
//		c.Redirect(http.StatusMovedPermanently, "/ng-app/")
//
//		//data := map[string]interface{}{
//		//	"status": true,
//		//}
//		//c.SecureJSON(http.StatusOK, data)
//		//
//	})
//
//	// Serve static files (that built by Angular and output to the ng-app)
//	router.Static("/ng-app", "../ng-app")
//
//	//
//	router.POST("/basic", profile.PostBasic)
//
//	// TODO: Login is a HTTP POST request, need to change from GET to POST
//	router.GET("/login/:name", func(c *gin.Context) {
//		name := c.Param("name")
//
//		sessionID := fmt.Sprintf("%s_%s", name, util.RandStringRunes(10))
//
//		c.SetCookie("gin_cookie", sessionID, 3600, "/", "", false, true)
//
//		//c.String(http.StatusOK, "Hello %s", name)
//		c.Redirect(http.StatusMovedPermanently, "/ng-app/")
//	})
//
//	// This handler will match /user/john but will not match /user/ or /user
//	router.GET("/user/:name", func(c *gin.Context) {
//		name := c.Param("name")
//		c.String(http.StatusOK, "Hello %s", name)
//	})
//
//	DNSDomainLookupData := map[string]string{
//		"google.com": "1.1.1.1",
//		"amazon.com": "2.1.1.1",
//	}
//	router.GET("/my-awesome-dns-service/:domain", func(c *gin.Context) {
//		domain := c.Param("domain")
//		ip := ""
//		var ok bool
//
//		//
//		data := map[string]interface{}{
//			"status": false,
//		}
//
//		if ip, ok = DNSDomainLookupData[domain]; !ok {
//			log.Printf("!ok: domain: %s; ip: %v; ok: %v", domain, ip, ok)
//			c.SecureJSON(http.StatusOK, data)
//			return
//		}
//		log.Printf("ok: domain: %s; ip: %v; ok: %v", domain, ip, ok)
//
//		data["status"] = true
//		data["ip"] = ip
//
//		c.SecureJSON(http.StatusOK, data)
//		return
//	})
//
//	// Submit data, then edit/update data on server side
//	router.POST("/user/edit/:name", func(c *gin.Context) {
//		name := c.Param("name")
//
//		age := c.PostForm("age")
//		msg := c.PostForm("msg")
//
//		data := map[string]interface{}{
//			"name": fmt.Sprintf("You input %s", name),
//			"age":  fmt.Sprintf("You input %s", age),
//			"msg":  fmt.Sprintf("You input %s", msg),
//		}
//
//		log.Printf("Data: %v", data)
//
//		c.SecureJSON(http.StatusOK, data)
//	})
//
//	router.POST("/group/edit/:name", func(c *gin.Context) {
//		var err error
//		var group model.Group
//
//		// Client will submit JSON data, then, we are trying to convert JSON
//		// string into "model.Group" (live object)
//		if err = c.ShouldBindJSON(&group); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"MY error": err.Error()})
//			return
//		}
//
//		// Data
//		data := map[string]interface{}{
//			"group": fmt.Sprintf("You input %#v", group),
//		}
//
//		// Return JSON back to client
//		c.SecureJSON(http.StatusOK, data)
//	})
//
//
//	router.Run(*srvportPtr)
//}
