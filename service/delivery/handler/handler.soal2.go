package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"logkar-test/service"
	"logkar-test/service/model"
	"net/http"
)

type Soal2Handler struct {
	e         *echo.Echo
	soal2Case service.ISoal2Usecase
}

func NewSoal2Handler(
	e *echo.Echo,
	soal2Case service.ISoal2Usecase,
) *Soal2Handler {
	return &Soal2Handler{
		e:         e,
		soal2Case: soal2Case,
	}
}

func (ox *Soal2Handler) CountSilinderArea(ec echo.Context) error {
	var form model.SilinderArgs

	err := json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	res, err := ox.soal2Case.CountSilinderArea(form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": err.Error(),
			"en": err.Error(),
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":   "Berhasil",
		"en":   "Success",
		"Luas": res,
	})
}

func (ox *Soal2Handler) CountCubeArea(ec echo.Context) error {
	var form model.CubeArgs
	err := json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	res, err := ox.soal2Case.CountCubeArea(form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": err.Error(),
			"en": err.Error(),
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":   "Berhasil",
		"en":   "Success",
		"Luas": res,
	})
}

func (ox *Soal2Handler) CountBalokArea(ec echo.Context) error {
	var form model.BalokArgs

	err := json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	res, err := ox.soal2Case.CountBlockArea(form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": err.Error(),
			"en": err.Error(),
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":   "Berhasil",
		"en":   "Success",
		"Luas": res,
	})
}

func (ox *Soal2Handler) CountKelilingSilinder(ec echo.Context) error {
	var form model.SilinderArgs

	err := json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	res, err := ox.soal2Case.CountKelilingSilinder(form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": err.Error(),
			"en": err.Error(),
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":       "Berhasil",
		"en":       "Success",
		"Keliling": res,
	})
}

func (ox *Soal2Handler) CountKelilingBalok(ec echo.Context) error {
	var form model.BalokArgs

	err := json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	res, err := ox.soal2Case.CountKelilingBalok(form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": err.Error(),
			"en": err.Error(),
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":       "Berhasil",
		"en":       "Success",
		"Keliling": res,
	})
}

func (ox *Soal2Handler) CountKelilingKubus(ec echo.Context) error {
	var form model.CubeArgs

	err := json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	res, err := ox.soal2Case.CountKelilingKubus(form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": err.Error(),
			"en": err.Error(),
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":       "Berhasil",
		"en":       "Success",
		"Keliling": res,
	})
}

func (ox *Soal2Handler) CountVolumeSilinder(ec echo.Context) error {
	var form model.SilinderArgs

	err := json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	res, err := ox.soal2Case.CountVolumeSilinder(form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": err.Error(),
			"en": err.Error(),
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":     "Berhasil",
		"en":     "Success",
		"Volume": res,
	})
}

func (ox *Soal2Handler) CountVolumeBalok(ec echo.Context) error {
	var form model.BalokArgs

	err := json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	res, err := ox.soal2Case.CountVolumeBalok(form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": err.Error(),
			"en": err.Error(),
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":     "Berhasil",
		"en":     "Success",
		"Volume": res,
	})
}

func (ox *Soal2Handler) CountVolumeKubus(ec echo.Context) error {
	var form model.CubeArgs

	err := json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	res, err := ox.soal2Case.CountVolumeKubus(form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": err.Error(),
			"en": err.Error(),
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":     "Berhasil",
		"en":     "Success",
		"Volume": res,
	})
}
