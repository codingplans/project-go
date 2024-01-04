package cli

import "testgo/modgo/gen/pkg"

type Client struct {
}
type Client2 struct {
}

func NewMysql() (*Client, error) {
	return &Client{}, nil
}

func NewMysql2() *Client2 {
	return &Client2{}
}

func (p *cityRepos) Add() {
	// TODO implement me
	panic("implement me")
}

func (p *cityRepos) Del() {

	// TODO implement me
	panic("implement me")
}

// NewCityRepo .
func NewCityRepo() pkg.PkgInterface {
	return &cityRepos{}
}

type cityRepos struct {
}
