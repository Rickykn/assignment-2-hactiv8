package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rickykn/assignment-2-hactiv8.git/database"
	"github.com/Rickykn/assignment-2-hactiv8.git/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	errConnect := database.Connect()

	if errConnect != nil {
		panic(errConnect)
	}
	gin.SetMode(gin.DebugMode)

	server := http.Server{
		Addr:    ":8082",
		Handler: routers.Server(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Server Listen Error : %s ", err.Error())
		}
	}()

	signals := make(chan os.Signal)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Println("Server is Shutdown")
}
