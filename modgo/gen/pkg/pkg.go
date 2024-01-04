package pkg

func (p Pkgs) Add() {
	// TODO implement me
	panic("implement me")
}

func (p Pkgs) Del() {
	// TODO implement me
	panic("implement me")
}

type PkgInterface interface {
	Add()
	Del()
}

type Pkgs struct {
}

func NewKafka() (PkgInterface, error) {
	return &Pkgs{}, nil
}
