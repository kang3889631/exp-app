package websrv

import (
	"net/http"
	"srv/core/config"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

// Init ...
func Init(httpsMux http.Handler) *errgroup.Group {
	// Set mode
	setMode()

	// Support running multiple web server (HTTP / HTTPS)
	eg := new(errgroup.Group)

	httpsSrv := &http.Server{
		Addr:    config.Data.WebBindAddrHTTPPort(),
		Handler: httpsMux,
		// Good practice: enforce timeouts for servers you create!
		ReadTimeout:  time.Duration(60) * time.Second,
		WriteTimeout: time.Duration(60) * time.Second,
		// in MB
		MaxHeaderBytes: 1 << 20,
	}

	eg.Go(func() error {
		var err error

		// Start listing and server requests
		if err = httpsSrv.ListenAndServe(); err != nil {
			return err
		}

		return nil
	})

	return eg
}

func setMode() {
	if config.Data.IsSysModeDev() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
