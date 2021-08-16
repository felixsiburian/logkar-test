package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"logkar-test/service"
	"logkar-test/service/model"
	"net/http"
)

type Soal1Handler struct {
	e         *echo.Echo
	soal1Case service.ISoal1Usecase
}

func NewSoal1Handler(
	e *echo.Echo,
	soal1Case service.ISoal1Usecase,
) *Soal1Handler {
	return &Soal1Handler{
		e:         e,
		soal1Case: soal1Case,
	}
}

func (ox *Soal1Handler) ValidationNIK(ec echo.Context) error {
	var form model.ModelSoal1
	err := json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	res, err := ox.soal1Case.ValidationNik(form.NIK)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": err.Error(),
			"en": err.Error(),
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":     "Berhasil",
		"en":     "Success",
		"status": res,
	})
}
