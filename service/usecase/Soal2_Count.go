package usecase

import (
	"errors"
	"fmt"
	"logkar-test/service"
	"logkar-test/service/model"
	"logkar-test/service/model/constant"
)

type soal2Usecase struct {
}

func NewSoal2Case() service.ISoal2Usecase {
	return soal2Usecase{}
}

func (s soal2Usecase) CountSilinderArea(args model.SilinderArgs) (res string, err error) {
	var pi = constant.Phi

	if args.JariJari <= 0 {
		err = errors.New("Jari-Jari is Invalid")
		return res, err
	}

	if args.Tinggi <= 0 {
		err = errors.New("Tinggi is Invalid")
		return res, err
	}

	luas := 2 * pi * args.JariJari * (args.JariJari + args.Tinggi)
	res = fmt.Sprintf("%.2f", luas)

	return res, err
}

func (s soal2Usecase) CountCubeArea(args model.CubeArgs) (res string, err error) {
	if args.Sisi <= 0 {
		err = errors.New("Sisi Invalid")
		return res, err
	}

	luas := 6 * (args.Sisi * args.Sisi)
	res = fmt.Sprintf("%.2f", luas)

	return res, err
}

func (s soal2Usecase) CountBlockArea(args model.BalokArgs) (res string, err error) {
	if args.Tinggi <= 0 {
		err = errors.New("Tinggi is Invalid")
		return res, err
	}

	if args.Lebar <= 0 {
		err = errors.New("Lebar is Invalid")
		return res, err
	}

	if args.Panjang <= 0 {
		err = errors.New("Panjang is Invalid ")
		return res, err
	}

	luas := 2 * (args.Panjang*args.Lebar + args.Panjang*args.Tinggi + args.Lebar*args.Tinggi)
	res = fmt.Sprintf("%.2f", luas)

	return res, err
}

func (s soal2Usecase) CountKelilingSilinder(args model.SilinderArgs) (res string, err error) {
	if args.JariJari <= 0 {
		err = errors.New("Jari-Jari is Invalid")
		return res, err
	}

	if args.Tinggi <= 0 {
		err = errors.New("Tinggi is Invalid")
		return res, err
	}

	keliling := 2 * ((2 * args.JariJari) + args.Tinggi)
	res = fmt.Sprintf("%.2f", keliling)

	return res, err
}

func (s soal2Usecase) CountKelilingKubus(args model.CubeArgs) (res string, err error) {
	if args.Sisi <= 0 {
		err = errors.New("Sisi is Invalid")
		return res, err
	}

	keliling := 12 * args.Sisi
	res = fmt.Sprintf("%.2f", keliling)

	return res, err
}

func (s soal2Usecase) CountKelilingBalok(args model.BalokArgs) (res string, err error) {
	if args.Tinggi <= 0 {
		err = errors.New("Tinggi is Invalid")
		return res, err
	}

	if args.Lebar <= 0 {
		err = errors.New("Lebar is Invalid")
		return res, err
	}

	if args.Panjang <= 0 {
		err = errors.New("Panjang is Invalid")
		return res, err
	}

	keliling := 4 * (args.Panjang + args.Lebar + args.Tinggi)
	res = fmt.Sprintf("%.2f", keliling)

	return res, err
}

func (s soal2Usecase) CountVolumeSilinder(args model.SilinderArgs) (res string, err error) {
	var pi = constant.Phi

	if args.JariJari <= 0 {
		err = errors.New("Jari-Jari is Invalid")
		return res, err
	}

	if args.Tinggi <= 0 {
		err = errors.New("Tinggi is Invalid")
		return res, err
	}

	volume := pi * args.JariJari * args.JariJari * args.Tinggi
	res = fmt.Sprintf("%.2f", volume)

	return res, err
}

func (s soal2Usecase) CountVolumeKubus(args model.CubeArgs) (res string, err error) {
	if args.Sisi <= 0 {
		err = errors.New("Sisi Invalid")
		return res, err
	}

	volume := args.Sisi * args.Sisi * args.Sisi
	res = fmt.Sprintf("%.2f", volume)

	return res, err
}

func (s soal2Usecase) CountVolumeBalok(args model.BalokArgs) (res string, err error) {
	if args.Tinggi <= 0 {
		err = errors.New("Tinggi is Invalid")
		return res, err
	}

	if args.Lebar <= 0 {
		err = errors.New("Lebar is Invalid")
		return res, err
	}

	if args.Panjang <= 0 {
		err = errors.New("Panjang is Invalid")
		return res, err
	}

	volume := args.Panjang * args.Tinggi * args.Lebar
	res = fmt.Sprintf("%.2f", volume)

	return res, err
}
