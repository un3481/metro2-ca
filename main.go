
package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"/server"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok { port := "8080" }
	
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

	env, _ := os.LookupEnv("METRO2_ENV")
	if env != "test" {
		if err := serve.ListenAndServe(); err != nil {
			return err
		}
	}
	return nil
}
