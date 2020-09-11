package main

import (
	"fmt"
	"github.com/mylukin/easy-i18n/i18n"
	"golang.org/x/text/message"
	"testgo/translt/catalog"

	// _ "github.com/mylukin/easy-i18n/example/catalog"
	_ "testgo/translt/catalog"
)

var p *i18n.Printer

var Lans map[string]*message.Printer

func init() {
	// a,ok:=i18n.Message["12"]
	// p = i18n.NewPrinter(language.SimplifiedChinese)
	// i18n.SetLang(language.SimplifiedChinese)
	// p = i18n.GetPrinter()

	println(555)
	Lans = catalog.Instance()
}

func main() {
	// time.Sleep(2 * time.Second)

	// p.Sprintf(`hello world111!`)
	// aa := p.Sprintf(`hello world!`)

	pp, ok := Lans["en"]
	if !ok {
		return
	}

	i18n.Printf("hello world!")
	fmt.Println(666)

	name := `Lukin`
	fmt.Printf(`hello %s!`, name)
	pp.Printf(`hello %s!`, name)

	fmt.Println()

	pp.Printf(`%s has %d cat.`, name, 1)
	fmt.Println()
	// sss()

}
