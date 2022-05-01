package chapter01

import "fmt"

type Portion int

const (
	Regular Portion = iota
	Small
	Large
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

type fluentOpt struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(p Portion) *fluentOpt {
	return &fluentOpt{
		men:      p,
		aburaage: false,
		ebiten:   1,
	}
}

func (o *fluentOpt) Aburaage() *fluentOpt {
	o.aburaage = true

	return o
}

func (o *fluentOpt) Ebiten(n uint) *fluentOpt {
	o.ebiten = n

	return o
}

func (o *fluentOpt) Order() *Udon {
	return &Udon{
		men:      o.men,
		aburaage: o.aburaage,
		ebiten:   o.ebiten,
	}
}

func builderExample() {
	oomoriKitsune := NewUdon(Large).Aburaage().Order()
	smallThreeEbiten := NewUdon(Small).Ebiten(3).Order()

	fmt.Println(*oomoriKitsune)
	fmt.Println(*smallThreeEbiten)
}
