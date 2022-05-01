package chapter01

import "fmt"

type OptFunc func(r *Udon)

func NewUdon3(opts ...OptFunc) *Udon {
	r := &Udon{}
	for _, opt := range opts {
		opt(r)
	}

	return r
}

func OptMen(p Portion) OptFunc {
	return func(r *Udon) {
		r.men = p
	}
}

func OptAburaage() OptFunc {
	return func(r *Udon) {
		r.aburaage = true
	}
}

func OptEbiten(n uint) OptFunc {
	return func(r *Udon) {
		r.ebiten = n
	}

}

func functionalOptionExample() {
	oomoriKitsune := NewUdon3(OptMen(Large), OptAburaage())
	tokuseiUdon := NewUdon3(OptAburaage(), OptEbiten(3))

	fmt.Println(*oomoriKitsune)
	fmt.Println(*tokuseiUdon)
}
