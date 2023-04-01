package routers

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/m3rashid-org/hmis-go-server/controllers"
)

func InitRouter(sig ...os.Signal) {
	router := SetupRouter()
	if len(sig) == 0 {
		sig = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}
	signalChan := make(chan os.Signal, 1)
	go func() {
		router.Run(fmt.Sprintf(":%v", os.Getenv("HTTP_PORT")))
	}()
	signal.Notify(signalChan, sig...)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.ExposeHeaders = []string{"Authorization"}
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	router.GET("/ping", controllers.PingHandler)

	return router
}
