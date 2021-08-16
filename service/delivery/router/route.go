package router

import (
	"github.com/labstack/echo"
	"logkar-test/service"
	"logkar-test/service/delivery/handler"
)

func NewRouter(
	e *echo.Echo,
	soal1Case service.ISoal1Usecase,
	soal2Case service.ISoal2Usecase,
	soal3Case service.ISoal3Usecase,
) {
	s1 := handler.NewSoal1Handler(e, soal1Case)
	s2 := handler.NewSoal2Handler(e, soal2Case)
	s3 := handler.NewSoal3Handler(e, soal3Case)

	r := e.Group("v1/logkar")
	// no1
	r.POST("/validation/nik", s1.ValidationNIK)

	// no2 Luas
	r.POST("/count/silinder/area", s2.CountSilinderArea)
	r.POST("/count/cube/area", s2.CountCubeArea)
	r.POST("/count/balok/area", s2.CountBalokArea)

	//	no2 keliling
	r.POST("/count/silinder/keliling", s2.CountKelilingSilinder)
	r.POST("/count/cube/keliling", s2.CountKelilingKubus)
	r.POST("/count/balok/keliling", s2.CountKelilingBalok)

	//	no2 Volume
	r.POST("/count/silinder/volume", s2.CountVolumeSilinder)
	r.POST("/count/cube/volume", s2.CountVolumeKubus)
	r.POST("/count/balok/volume", s2.CountVolumeBalok)

	// no3 simple HTTP Client
	r.GET("/get/movies/omdb", s3.GetMoviesOMDB)
}
