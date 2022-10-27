package kongjian

type SearchInter[D any] interface {
	SearchAdd(data D) error
	SearchExist(data D) (bool, error)
	SearchDelete(data D) error
}

var _ SearchInter[[]byte] = (*SearchData[[]byte, any])(nil)

type SearchData[D any, R any] struct {
	v R
	*Feature[D, R]
}

func (s *SearchData[D, R]) SearchAdd(data D) error {
	// if s.v == nil {
	s.getV(data)
	// }

	return s.Store.Add(s.v)
}

func (s *SearchData[D, R]) getV(data D) (v R) {
	for _, f := range s.Evolves {
		data = f.Evolve(data)
	}
	s.v = s.Latest.Evolve(data)
	return s.v
}

func (s *SearchData[D, R]) SearchExist(data D) (bool, error) {
	// if s.v == nil {
	s.getV(data)
	// }
	return s.Store.Exist(s.v)
}
func (s *SearchData[D, R]) SearchDelete(data D) error {
	// if s.v == nil {
	s.getV(data)
	// }
	return s.Store.Delete(s.v)
}
