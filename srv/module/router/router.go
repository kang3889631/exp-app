package router

import (
	"net/http"

	"srv/module/OnlineOrder"
	"srv/module/upload"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// HTTPSAll ...
func HTTPSAll() http.Handler {
	// Creates a router without any middleware by default
	r := gin.New()

	HTTPSMin(r)
	HTTPSHome(r)
	HTTPSAPI(r)

	return r
}

// HTTPSMin ...
func HTTPSMin(r *gin.Engine) {
	loadHTTPSMinMiddleware(r)

	r.NoRoute(func(c *gin.Context) {
		data := map[string]interface{}{}
		c.SecureJSON(http.StatusNotFound, data)
	})
}

// HTTPSHome ...
func HTTPSHome(r *gin.Engine) {
	rg := r.Group("/")

	rg.Use()
	{
		rg.GET("/", SrvHome)
		rg.Static("/ng-app", "../ng-app")

		// Handler
		rg.GET("/trycookie", func(c *gin.Context) {
			// Process: Write a program according to the business logic and requirement
			name := "demo2"
			age := "age2"
			sessionToken := "SDF^*&^*#@^$*#@fsdfhsdjfhjvwer"

			c.SetCookie("name", name, 3600, "/", "", false, true)
			c.SetCookie("age", age, 3600, "/", "", false, true)
			c.SetCookie("sessionID", sessionToken, 3600, "/", "", false, true)
		})

	}
}

// HTTPSAPI ...
func HTTPSAPI(r *gin.Engine) {
	rg := r.Group("/api")

	rg.Use()
	{
		// Use HTTP "POST" request when modifying backend data
		rg.POST("/Upload", upload.SrvUpload)

		// Use HTTP "POST" request when modifying backend data
		rg.POST("/OnlineOrder", OnlineOrder.SrvPostOnlineOrder)

		// Use HTTP "GET" request when retrieving backend data
		rg.GET("/OnlineOrder", OnlineOrder.SrvGetOnlineOrder)
	}
}

// SrvHome ...
func SrvHome(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/ng-app")
}

func loadHTTPSMinMiddleware(r *gin.Engine) {
	r.Use(
		// Global middleware
		// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
		// By default gin.DefaultWriter = os.Stdout
		gin.Logger(),
		// Recovery middleware recovers from any panics and writes a 500 if there was one.
		gin.Recovery(),
		// gzip compression middleware
		gzip.Gzip(gzip.DefaultCompression),
		// Secure middleware
		//middleware.SecureMiddleware(),
		// CSRF token middleware
		//adapter.Wrap(middleware.CSRFMiddleware),
	)
}
