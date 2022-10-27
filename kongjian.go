package kongjian

import "sync"

type KongJianInter[D any, R any] interface {
	InitFeatures(features []*Feature[D, R]) error
	SearchInter[D]
}

var _ KongJianInter[[]byte, int64] = (*KongJian)(nil)
var _ SearchInter[[]byte] = (*KongJian)(nil)

type KongJian struct {
	sync.RWMutex
	featureDatas []*SearchData[[]byte, int64]
}

func (k *KongJian) InitFeatures(features []*Feature[[]byte, int64]) error {
	if features == nil {
		return nil
	}
	k.featureDatas = func() (res []*SearchData[[]byte, int64]) {
		for i := range features {
			res = append(res, &SearchData[[]byte, int64]{
				Feature: features[i],
			})
		}
		return
	}()

	return nil
}

func (k *KongJian) SearchAdd(data []byte) error {
	k.Lock()
	defer k.Unlock()
	for i := range k.featureDatas {
		err := k.featureDatas[i].SearchAdd(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k *KongJian) SearchExist(data []byte) (bool, error) {
	k.RLock()
	defer k.RUnlock()
	for i := range k.featureDatas {
		ok, err := k.featureDatas[i].SearchExist(data)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return true, nil
}

func (k *KongJian) SearchDelete(data []byte) error {
	k.Lock()
	defer k.Unlock()
	for i := range k.featureDatas {
		err := k.featureDatas[i].SearchDelete(data)
		if err != nil {
			return err
		}
	}
	return nil
}
