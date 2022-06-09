package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ernestngugi/go-blog/app/db"
	"github.com/ernestngugi/go-blog/app/web/router"
	"github.com/joho/godotenv"
)

const defaultPort = "3000"

func main() {

	var envFile string
	flag.StringVar(&envFile, "e", "", "path to env file")
	flag.Parse()

	if envFile != "" {
		err := godotenv.Load(envFile)
		if err != nil {
			panic(fmt.Errorf("failed to load env file =[%v]", err))
		}
	}

	dB := db.InitDB()

	defer dB.Close()

	router := router.BuildRouter(dB)

	server := &http.Server{
		Addr:    ":" + defaultPort,
		Handler: router,
	}

	fmt.Printf("server starting, listening on port: %v \n", defaultPort)

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			fmt.Println("server shut down")
		} else {
			fmt.Println("server shut down unexpectedly")
		}
	}

	done := make(chan chan struct{})

	//server shut down gracefully
	timeout := 30 * time.Second

	signalint := make(chan os.Signal, 1)
	signal.Notify(signalint, os.Interrupt, syscall.SIGTERM)

	code := 0
	select {
	case <-signalint:
		code = 1
	case <-time.After(timeout):
		code = 1
	case <-done:
		fmt.Println("shutdown completed.")
	}

	fmt.Println("server exiting")

	os.Exit(code)
}
