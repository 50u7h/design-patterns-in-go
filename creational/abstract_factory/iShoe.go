package main

type iShoe interface {
	setBrand(brand string)
	setSize(size int)
	getBrand() string
	getSize() int
}

type shoe struct {
	brand string
	size  int
}

func (s *shoe) setBrand(brand string) {
	s.brand = brand
}

func (s *shoe) getBrand() string {
	return s.brand
}

func (s *shoe) setSize(size int) {
	s.size = size
}

func (s *shoe) getSize() int {
	return s.size
}
