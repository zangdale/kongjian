package kongjian

import "testing"

func TestKongjain(t *testing.T) {
	var kj KongJianInter[[]byte, int64] = &KongJian{}
	kj.InitFeatures([]*Feature[[]byte, int64]{
		&Feature[[]byte, int64]{
			Latest: EvolveInt64,
			Evolves: []EvolveInter[[]byte, []byte]{
				EvolveData,
			},
			Store: NewStoreInt64(),
		},
		&Feature[[]byte, int64]{
			Latest: EvolveInt64,
			Evolves: []EvolveInter[[]byte, []byte]{
				EvolveData,
				EvolveData,
			},
			Store: NewStoreInt64(),
		},
	})

	t.Log(kj.SearchAdd([]byte("123")))
	t.Log(kj.SearchExist([]byte("123")))
	t.Log(kj.SearchDelete([]byte("123")))
	t.Log(kj.SearchExist([]byte("123")))
}
