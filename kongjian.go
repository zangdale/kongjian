package kongjian

import "sync"

type KongJianInter[D any, R any] interface {
	InitFeatures(features []*Feature[D, R]) error
	SearchAdd(data D) error
	SearchExist(data D) (bool, error)
	SearchDelete(data D) error
}

var _ SearchInter[[]byte, int64] = (*SearchDataInt64)(nil)

type SearchDataInt64 struct {
	*SearchData[[]byte, int64]
}

func (s *SearchDataInt64) VisNil(v int64) bool {
	return v == 0
}

var _ KongJianInter[[]byte, int64] = (*KongJian)(nil)

type KongJian struct {
	sync.RWMutex
	featureDatas []*SearchDataInt64
}

func (k *KongJian) InitFeatures(features []*Feature[[]byte, int64]) error {
	if features == nil {
		return nil
	}
	k.featureDatas = func() (res []*SearchDataInt64) {
		for i := range features {
			res = append(res, &SearchDataInt64{
				SearchData: &SearchData[[]byte, int64]{
					Feature: features[i],
				},
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
