package factory

import "fmt"

type TomatoSoup struct {
}

func (soup *TomatoSoup) Type() string {
	return "Tomato Soup"
}
func(soup *TomatoSoup) Prepare() {
	fmt.Printf("Preparing %s\n", soup.Type())
}

type EgusiSoup struct {
}
func (soup *EgusiSoup) Type() string {
	return "Egusi Soup"
}
func(soup *EgusiSoup) Prepare() {
	fmt.Printf("Preparing %s\n", soup.Type())
}

type EfoSoup struct {
}
func (soup *EfoSoup) Type() string {
	return "Efo Soup"
}
func(soup *EfoSoup) Prepare() {
	fmt.Printf("Preparing %s\n", soup.Type())
}