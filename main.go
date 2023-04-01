package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/judwhite/go-svc"
	"github.com/m3rashid-org/hmis-go-server/redis"
	"github.com/m3rashid-org/hmis-go-server/routers"
)
const portNumber = ":8080"

type program struct{}

func (p *program) Init(env svc.Environment) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}
	redis.ConnectRedis()
	return nil
}

func (p *program) Stop() error {
	fmt.Println("\nServer stoping . . .")
	time.Sleep(time.Duration(1) * time.Second)
	return nil
}

func (p *program) Start() error {
	routers.InitRouter(os.Interrupt)
	fmt.Println("Server Started 🏄 on :", os.Getenv("HTTP_PORT"))
	return nil
}

func main() {
	prg := &program{}
	if err := svc.Run(prg, os.Interrupt); err != nil {
		log.Fatal(err)
	}
}
