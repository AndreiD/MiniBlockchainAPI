package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/AndreiDD/tokenominatorapi/configs"
	"gitlab.com/AndreiDD/tokenominatorapi/utils"
)

var router *gin.Engine

func main() {

	Config := configs.Load()

	if (Config.GetString("environment")) == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// allow all origins
	router.Use(cors.Default())

	// initialize all the routes
	InitializeRoutes()

	fmt.Println("")
	fmt.Println("████████╗ ██████╗ ██╗  ██╗███████╗███╗   ██╗ ██████╗ ███╗   ███╗██╗███╗   ██╗ █████╗ ████████╗ ██████╗ ██████╗     ")
	fmt.Println("╚══██╔══╝██╔═══██╗██║ ██╔╝██╔════╝████╗  ██║██╔═══██╗████╗ ████║██║████╗  ██║██╔══██╗╚══██╔══╝██╔═══██╗██╔══██╗    ")
	fmt.Println("   ██║   ██║   ██║█████╔╝ █████╗  ██╔██╗ ██║██║   ██║██╔████╔██║██║██╔██╗ ██║███████║   ██║   ██║   ██║██████╔╝    ")
	fmt.Println("   ██║   ██║   ██║██╔═██╗ ██╔══╝  ██║╚██╗██║██║   ██║██║╚██╔╝██║██║██║╚██╗██║██╔══██║   ██║   ██║   ██║██╔══██╗    ")
	fmt.Println("   ██║   ╚██████╔╝██║  ██╗███████╗██║ ╚████║╚██████╔╝██║ ╚═╝ ██║██║██║ ╚████║██║  ██║   ██║   ╚██████╔╝██║  ██║    ")
	fmt.Println("   ╚═╝    ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝    ")
	fmt.Println("")
	fmt.Println("Version: 0.1")
	fmt.Println("http://" + Config.GetString("hostname") + ":" + strconv.Itoa(Config.GetInt("port")) + "/api/v1/")
	fmt.Println("")

	server := &http.Server{
		Addr:           Config.GetString("hostname") + ":" + strconv.Itoa(Config.GetInt("port")),
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.SetKeepAlivesEnabled(true)

	// Init the ETH Client
	if err := utils.InitEthClient(Config.GetString("node_address")); err != nil {
		log.Fatalf("cannot connect to the ETH provider. Please check if you have geth running: %s", err)
	}

	// Init the Rinkeby Client
	if err := utils.InitRinkeByClient(); err != nil {
		log.Fatalf("cannot connect to the ETH provider. Please check if you have geth running: %s", err)
	}

	// Serve'em
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server Initiated")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting. Bye!")

}
