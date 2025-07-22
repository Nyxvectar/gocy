/**
 * Author:  Nyxvectar Yan
 * Repo:    siracusan
 * Created: 07/22/2025
 */

package block

type makeSound interface {
	speaker()
}

type Dog struct {
	name string
}

type Phone struct {
	model string
}

func (d Dog) speaker() {
	println(d.name, "woofed")
}

func (p Phone) speaker() {
	println(p.model, "sounded")
}

func LetItGo(s makeSound) {
	s.speaker()
}

func InterfaceTrain() {
	var p = Phone{model: "Gentoo To Go"}
	var d = Dog{name: "玄离"}
	LetItGo(p)
	LetItGo(d)
}
