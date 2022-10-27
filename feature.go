package kongjian

type Feature[D any, R any] struct {
	Latest  EvolveInter[D, R]
	Evolves []EvolveInter[D, D]
	Store   StoreInter[R]
}
