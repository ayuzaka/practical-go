package chapter01

import (
	"fmt"
	"time"
)

type Option struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon2(opt Option) *Udon {
	if opt.ebiten == 0 && time.Now().Hour() < 10 {
		opt.ebiten = 1
	}

	return &Udon{
		men:      opt.men,
		aburaage: opt.aburaage,
		ebiten:   opt.ebiten,
	}
}

func structExample() {
	opt := Option{
		men:      Large,
		aburaage: true,
		ebiten:   1,
	}
	udon := NewUdon2(opt)

	fmt.Println(*udon)
}
