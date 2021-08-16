package usecase

import (
	"errors"
	"logkar-test/service"
	"logkar-test/service/tools"
)

type soal1Usecase struct {
}

func NewSoal1Usecase() service.ISoal1Usecase {
	return soal1Usecase{}
}

func (s soal1Usecase) ValidationNik(nik string) (res string, err error) {
	if len(nik) < 16 {
		err = errors.New("Invalid NIK")
		return res, err
	}

	if !tools.CheckContainsString(nik) {
		err = errors.New("Invalid NIK")
		return res, err
	}

	res = "NIK is Valid"

	return res, err
}
