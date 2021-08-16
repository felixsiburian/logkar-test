package model

type (
	SilinderArgs struct {
		JariJari float64 `json:"jariJari"`
		Tinggi   float64 `json:"tinggi"`
	}

	BalokArgs struct {
		Panjang float64 `json:"panjang"`
		Lebar   float64 `json:"lebar"`
		Tinggi  float64 `json:"tinggi"`
	}

	CubeArgs struct {
		Sisi float64 `json:"sisi"`
	}
)
