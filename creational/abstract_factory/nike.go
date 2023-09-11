package main

type nike struct {
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			brand: "nike",
			size:  40,
		},
	}
}

func (n *nike) makeShort() iShort {
	return &nikeShort{
		short: short{
			brand: "nike",
			size:  "S",
		},
	}
}
