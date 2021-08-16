package service

import "logkar-test/service/model"

type ISoal1Usecase interface {
	ValidationNik(nik string) (res string, err error)
}

type ISoal2Usecase interface {
	// Luas
	CountSilinderArea(args model.SilinderArgs) (res string, err error)
	CountBlockArea(args model.BalokArgs) (res string, err error)
	CountCubeArea(args model.CubeArgs) (res string, err error)

	//Keliling
	CountKelilingSilinder(args model.SilinderArgs) (res string, err error)
	CountKelilingKubus(args model.CubeArgs) (res string, err error)
	CountKelilingBalok(args model.BalokArgs) (res string, err error)

	//Volume
	CountVolumeSilinder(args model.SilinderArgs) (res string, err error)
	CountVolumeKubus(args model.CubeArgs) (res string, err error)
	CountVolumeBalok(args model.BalokArgs) (res string, err error)
}

type ISoal3Usecase interface {
	GetMovieFromOmdbAPI(page int, keyword string) (res model.SearchResponse, err error)
}
