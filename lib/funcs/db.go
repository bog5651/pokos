package funcs

import (
	"pokos/lib/types"

	"github.com/guark/guark/app"
)

func GetKKM(c app.Context) (interface{}, error) {
	return make([]types.KKM, 0), nil
}

func GetClients(c app.Context) (interface{}, error) {
	return make([]types.Client, 0), nil
}

func GetModelsKKM(c app.Context) (interface{}, error) {
	return make([]types.ModelKKM, 0), nil
}
