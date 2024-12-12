package main

import "context"

type Boo struct {
	Bar int
	Foo int
}

func (b *Boo) Get(ctx context.Context) string {
	return "Boo"
}

func (b *Boo) Set(ctx context.Context, str string) string {
	return "Boo"
}

type Boo2 struct {
	Bar int
	Foo int
	fzz []int
}

func (b *Boo2) Get(ctx context.Context) string {
	return "Boo2"
}

func (b *Boo2) Set(ctx context.Context, str string) string {
	return "Boo2"
}

type Api interface {
	Get(ctx context.Context) string
	Set(ctx context.Context, str string) string
}

func Dess() {
	B1 := new(Boo)
	B2 := new(Boo2)
	apiObj1 := Api(B1)
	apiObj1.Get(context.Background())
	apiObj1 = B2
	apiObj1.Get(context.Background())

}
