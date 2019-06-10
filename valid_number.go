package leet

import "unicode"

type token int

const (
	wsToken token = iota
	digitToken
	dotToken
	expToken
	minusToken
	unknownToken
)

func toToken(r rune) token {
	if unicode.IsNumber(r) {
		return digitToken
	}

	if unicode.IsSpace(r) {
		return wsToken
	}

	if r == '.' {
		return dotToken
	}

	if r == 'e' {
		return expToken
	}

	if r == '-' || r == '+' {
		return minusToken
	}

	return unknownToken
}

type finiteState struct {
	isFinal bool
	next map[token]*finiteState
}

func (fs *finiteState) move(r rune) *finiteState {
	return fs.next[toToken(r)]
}

func createFSM() *finiteState {
	six := &finiteState{
		isFinal:true,
		next:make(map[token]*finiteState),
	}

	six.next[wsToken] = six

	five := &finiteState{
		isFinal:true,
		next:make(map[token]*finiteState),
	}

	five.next[wsToken] = six
	five.next[digitToken] = five

	fiveMinus := &finiteState{
		isFinal:false,
		next:make(map[token]*finiteState),
	}

	fiveMinus.next[digitToken] = five

	fivePrime := &finiteState{
		isFinal:false,
		next:make(map[token]*finiteState),
	}

	fivePrime.next[digitToken] = five
	fivePrime.next[minusToken] = fiveMinus

	four := &finiteState{
		isFinal:true,
		next:make(map[token]*finiteState),
	}

	four.next[expToken] = fivePrime
	four.next[digitToken] = four
	four.next[wsToken] = six

	fourPrime := &finiteState{
		isFinal:false,
		next:make(map[token]*finiteState),
	}

	fourPrime.next[digitToken] = four


	three := &finiteState{
		isFinal:true,
		next:make(map[token]*finiteState),
	}

	three.next[dotToken] = four
	three.next[digitToken] = three
	three.next[wsToken] = six
	three.next[expToken] = fivePrime

	two := &finiteState{
		isFinal:false,
		next:make(map[token]*finiteState),
	}

	two.next[digitToken] = three
	two.next[dotToken] = fourPrime

	one := &finiteState{
		isFinal:false,
		next:make(map[token]*finiteState),
	}

	one.next[wsToken] = one
	one.next[minusToken] = two
	one.next[digitToken] = three
	one.next[dotToken] = fourPrime

	return one
}

func IsNumber(str string) bool {
	fs := createFSM()

	for _, r := range str {
		fs = fs.move(r)
		if fs == nil {
			return false
		}
	}
	return fs.isFinal
}


