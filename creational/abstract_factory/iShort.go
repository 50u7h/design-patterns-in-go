package main

type iShort interface {
	setBrand(brand string)
	setSize(size string)
	getBrand() string
	getSize() string
}

type short struct {
	brand string
	size  string
}

func (s *short) setBrand(brand string) {
	s.brand = brand
}

func (s *short) getBrand() string {
	return s.brand
}

func (s *short) setSize(size string) {
	s.size = size
}

func (s *short) getSize() string {
	return s.size
}
