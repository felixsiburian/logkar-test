package main

import (
	"fmt"
	"github.com/labstack/echo"
	"logkar-test/service/config"
	"logkar-test/service/delivery/router"
	"logkar-test/service/usecase"
	"os"
)

func main() {
	app := config.Config{}
	config.CatchEror(app.InitEnv())

	Start()
}

func Start() {
	e := echo.New()

	soal1Case := usecase.NewSoal1Usecase()
	soal2Case := usecase.NewSoal2Case()
	soal3Case := usecase.NewSoal3Usecase()

	router.NewRouter(e, soal1Case, soal2Case, soal3Case)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s%v", os.Getenv("APP_HOST"), ":", os.Getenv("APP_PORT"))))
}
