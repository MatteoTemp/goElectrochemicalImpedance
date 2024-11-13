package main

import (
	"goEl/lexer"
)

//cmath "math/cmplx"

func main() {

	cirtcuit_descriprion := "R(CR)"

	cd := lexer.Init(cirtcuit_descriprion)

	if err1 := cd.Tokenize(); err1 != nil {
		panic("Error reading tokens, ")
	}

	if err2 := cd.MatchBrakets(); err2 != nil {
		panic("matching error")
	}

	cd.Display()

}
