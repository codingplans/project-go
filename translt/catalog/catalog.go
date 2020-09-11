package catalog

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var Lans map[string]*message.Printer

// init
func Instance() map[string]*message.Printer {
	Lans = make(map[string]*message.Printer, 0)
	initEn(language.Make("en"))
	initZhHans(language.Make("zh-Hans"))
	return Lans
}

// initEn will init en support.
func initEn(tag language.Tag) {
	_ = message.SetString(tag, "%s has %d cat.", "%s has %d c7777777at.")
	message.SetString(tag, "%s has %d cats.", "%s has %d ca7777777ts.")
	message.SetString(tag, "%s have %d apple.", "%s have %d app7777777le.")
	message.SetString(tag, "%s have %d apples.", "%s have %d appl7777777es.")
	message.SetString(tag, "%s have an apple.", "%s have an app7777777le.")
	message.SetString(tag, "%s have two apples.", "%s have two appl7777777es.")
	message.SetString(tag, "hello %s!", "hello 7777777%s!")
	message.SetString(tag, "hello world111!", "hello world1777777711!")
	Lans[tag.String()] = message.NewPrinter(tag)
}

// initZhHans will init zh-Hans support.
func initZhHans(tag language.Tag) {
	message.SetString(tag, "%s has %d cat.", "%s has %d 88888888cat.")
	message.SetString(tag, "%s has %d cats.", "%s has %d 88888888cats.")
	message.SetString(tag, "%s have %d apple.", "%s have %d 88888888apple.")
	message.SetString(tag, "%s have %d apples.", "%s have %d 88888888apples.")
	message.SetString(tag, "%s have an apple.", "%s have an 88888888apple.")
	message.SetString(tag, "%s have two apples.", "%s have two 88888888apples.")
	message.SetString(tag, "hello %s!", "hello %88888888s!")
	message.SetString(tag, "hello world111!", "hello w8888orld111!")
	Lans[tag.String()] = message.NewPrinter(tag)
}
