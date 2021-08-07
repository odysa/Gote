package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/odysa/Gote/config"
	"github.com/odysa/Gote/model"
	"github.com/odysa/Gote/router"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var configName = pflag.StringP("config", "c", "", "server config file path")

func main() {
	pflag.Parse()

	if err := config.InitConfig(*configName); err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("mode"))

	app := gin.New()

	model.DB.Init()
	defer model.DB.Close()

	middleware := []gin.HandlerFunc{}

	// load router
	router.Load(app, middleware...)

	healthCheck()

	port := viper.GetInt("port")

	log.Printf("Server launched on port %d", port)

	log.Printf(http.ListenAndServe(fmt.Sprintf(":%d", port), app).Error())
}

// check server's health
func healthCheck() {
	go func() {
		if err := ping(); err != nil {
			log.Fatal("failed to launch server")
		}
	}()
}

func ping() error {
	for i := 0; i < 10; i++ {
		url := fmt.Sprintf("http://localhost:%d/hc/health", viper.GetInt("port"))
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Print("Waiting for server response")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to server")
}
