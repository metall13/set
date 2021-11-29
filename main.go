package set

type SummBag struct {
	bag []int
}

func NewSummBag() SummBag {
	return SummBag{}
}

func (b *SummBag) Add(i int) SummBag {
	newBag := append(b.bag, i)
	return SummBag{newBag}
}

func (b *SummBag) Summ() int {
	s := 0
	for _, i := range b.bag {
		s += i
	}
	return s
}
