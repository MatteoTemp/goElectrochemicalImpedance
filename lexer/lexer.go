package lexer

import (
	"errors"
	"fmt"
)

// Token codice
const (
	Resitor = iota + 1
	Capacitor
	Inductor
	Warbourg
	CPE

	Left
	Right
	Series

	END
)

type CDC_Code struct {
	Text           string
	tokens         []int
	match_position []int
}

func Init(input string) CDC_Code {

	return CDC_Code{
		Text: input,
	}
}

func (circuit *CDC_Code) Tokenize() error {

	for _, symbol := range circuit.Text {
		circuit.match_position = append(circuit.match_position, 0)
		switch symbol {
		case 'R':
			circuit.tokens = append(circuit.tokens, Resitor)
		case 'C':
			circuit.tokens = append(circuit.tokens, Capacitor)
		case 'L':
			circuit.tokens = append(circuit.tokens, Inductor)
		case 'W':
			circuit.tokens = append(circuit.tokens, Warbourg)
		case 'Q':
			circuit.tokens = append(circuit.tokens, CPE)
		case '-':
			circuit.tokens = append(circuit.tokens, Series)
		case '(':
			circuit.tokens = append(circuit.tokens, Left)
		case ')':
			circuit.tokens = append(circuit.tokens, Right)
		default:
			return errors.New("Invalid symbol: " + string(symbol) + " found in the input")
		}
	}

	if len(circuit.tokens) != len(circuit.match_position) {
		return errors.New("Generic error")
	}

	return nil
}

func (circuit *CDC_Code) MatchBrakets() error {

	/*Ma quanto è bella sta funzione */

	for i := 0; i < len(circuit.Text)-1; i++ {
		var count int = 0
		if circuit.Text[i] == '(' {
			count = 1
			for j := i + 1; j < len(circuit.Text); j++ {
				switch circuit.Text[j] {
				case '(':
					count += 1
				case ')':
					count += -1
				default:
					count += 0
				}

				if count == 0 {
					circuit.match_position[i] = j
					break
				}

			}

			if count != 0 {
				return errors.New("Matching Error")
			}

		}

	}

	return nil
}

func (circuit *CDC_Code) Display() {
	fmt.Println("Testo inserito: ", circuit.Text)

	for i, tok := range circuit.tokens {

		symbol := ' '
		switch tok {
		case Resitor:
			symbol = 'R'
		case Capacitor:
			symbol = 'C'
		case Inductor:
			symbol = 'L'
		case Warbourg:
			symbol = 'W'
		case CPE:
			symbol = 'Q'
		case Series:
			symbol = '-'
		case Left:
			symbol = '('
		case Right:
			symbol = ')'
		}

		fmt.Println(i, "\t=>\t", string(symbol), "\t", circuit.match_position[i])
	}
}
