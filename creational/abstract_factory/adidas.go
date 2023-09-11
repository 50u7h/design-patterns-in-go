package main

type adidas struct {
}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			brand: "adidas",
			size:  42,
		},
	}
}

func (a *adidas) makeShort() iShort {
	return &adidasShort{
		short: short{
			brand: "adidas",
			size:  "L",
		},
	}
}
