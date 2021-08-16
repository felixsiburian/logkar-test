package handler

import (
	"github.com/labstack/echo"
	"logkar-test/service"
	"net/http"
	"strconv"
)

type Soal3Handler struct {
	e         *echo.Echo
	soal3Case service.ISoal3Usecase
}

func NewSoal3Handler(
	e *echo.Echo,
	soal3Case service.ISoal3Usecase,
) *Soal3Handler {
	return &Soal3Handler{
		e:         e,
		soal3Case: soal3Case,
	}
}

func (ox *Soal3Handler) GetMoviesOMDB(ec echo.Context) error {
	page := ec.QueryParam("page")
	keyword := ec.QueryParam("keyword")
	pageInt, _ := strconv.Atoi(page)

	res, err := ox.soal3Case.GetMovieFromOmdbAPI(pageInt, keyword)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":   "Berhasil",
		"en":   "Success",
		"data": res,
	})
}
