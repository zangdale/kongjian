package kongjian

import (
	"encoding/base64"
)

type EvolveInter[D any, R any] interface {
	Evolve(data D) (result R)
}

var _ EvolveInter[any, any] = (*EvolveFunc[any, any])(nil)

type EvolveFunc[D any, R any] func(data D) (result R)

func (f EvolveFunc[D, R]) Evolve(data D) (result R) {
	return f(data)
}

var EvolveData EvolveFunc[[]byte, []byte] = func(data []byte) (result []byte) {
	s := base64.StdEncoding.EncodeToString(data)
	return []byte(s)
}

var EvolveInt64 EvolveFunc[[]byte, int64] = func(data []byte) (result int64) {
	return int64(len(data))
}
