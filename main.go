
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/un3481/sample-metro2-server/server"
)

func main() {
	port := os.Getenv("PORT") || "8080"
	fmt.Fprintf(os.Stdout, "Starting web server on port %s\n\n", port)

	timeout, _ := time.ParseDuration("30s")
	handler, _ := server.ConfigureHandlers()
	serve := &http.Server{
		Addr:              "0.0.0.0:" + port,
		Handler:           handler,
		ReadTimeout:       timeout,
		ReadHeaderTimeout: timeout,
		WriteTimeout:      timeout,
		IdleTimeout:       timeout,
	}

	if "test" != os.Getenv("ENV") {
		if err := serve.ListenAndServe(); err != nil {
			fmt.Fprintf(os.Stdout, "Could not start web server on port %s\n\n", port)
		}
	}
}
