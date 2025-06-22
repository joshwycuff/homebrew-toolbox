package score

import (
	"slices"
)

const baseScore = 1

type finger string

const none finger = "none"
const leftPinky finger = "leftPinky"
const leftRing finger = "leftRing"
const leftMiddle finger = "leftMiddle"
const leftIndex finger = "leftIndex"
const rightIndex finger = "rightIndex"
const rightMiddle finger = "rightMiddle"
const rightRing finger = "rightRing"
const rightPinky finger = "rightPinky"
const thumb finger = "thumb"

var char2Finger = map[rune]finger{
	'`': leftPinky, '~': leftPinky,
	'1': leftPinky, 'q': leftPinky, 'a': leftPinky, 'z': leftPinky,
	'!': leftPinky, 'Q': leftPinky, 'A': leftPinky, 'Z': leftPinky,
	'2': leftRing, 'w': leftRing, 's': leftRing, 'x': leftRing,
	'@': leftRing, 'W': leftRing, 'S': leftRing, 'X': leftRing,
	'3': leftMiddle, 'e': leftMiddle, 'd': leftMiddle, 'c': leftMiddle,
	'#': leftMiddle, 'E': leftMiddle, 'D': leftMiddle, 'C': leftMiddle,
	'4': leftIndex, 'r': leftIndex, 'f': leftIndex, 'v': leftIndex,
	'$': leftIndex, 'R': leftIndex, 'F': leftIndex, 'V': leftIndex,
	'5': leftIndex, 't': leftIndex, 'g': leftIndex, 'b': leftIndex,
	'%': leftIndex, 'T': leftIndex, 'G': leftIndex, 'B': leftIndex,
	'6': rightIndex, 'y': rightIndex, 'h': rightIndex, 'n': rightIndex,
	'^': rightIndex, 'Y': rightIndex, 'H': rightIndex, 'N': rightIndex,
	'7': rightIndex, 'u': rightIndex, 'j': rightIndex, 'm': rightIndex,
	'&': rightIndex, 'U': rightIndex, 'J': rightIndex, 'M': rightIndex,
	'8': rightMiddle, 'i': rightMiddle, 'k': rightMiddle, ',': rightMiddle,
	'*': rightMiddle, 'I': rightMiddle, 'K': rightMiddle, '<': rightMiddle,
	'9': rightRing, 'o': rightRing, 'l': rightRing, '.': rightRing,
	'(': rightRing, 'O': rightRing, 'L': rightRing, '>': rightRing,
	'0': rightPinky, 'p': rightPinky, ';': rightPinky, '/': rightPinky,
	')': rightPinky, 'P': rightPinky, ':': rightPinky, '?': rightPinky,
	'-': rightPinky, '[': rightPinky, '\'': rightPinky,
	'_': rightPinky, '{': rightPinky, '"': rightPinky,
	'=': rightPinky, ']': rightPinky,
	'+': rightPinky, '}': rightPinky,
	'|': rightPinky, '\\': rightPinky,
	' ': thumb,
}

var char2Row = map[rune]int{
	'`': 4, '1': 4, '2': 4, '3': 4, '4': 4, '5': 4, '6': 4, '7': 4, '8': 4, '9': 4, '0': 4, '-': 4, '=': 4,
	'~': 4, '!': 4, '@': 4, '#': 4, '$': 4, '%': 4, '^': 4, '&': 4, '*': 4, '(': 4, ')': 4, '_': 4, '+': 4,
	'q': 3, 'w': 3, 'e': 3, 'r': 3, 't': 3, 'y': 3, 'u': 3, 'i': 3, 'o': 3, 'p': 3, '[': 3, ']': 3, '\\': 3,
	'Q': 3, 'W': 3, 'E': 3, 'R': 3, 'T': 3, 'Y': 3, 'U': 3, 'I': 3, 'O': 3, 'P': 3, '{': 3, '}': 3, '|': 3,
	'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2, ';': 2, '\'': 2,
	'A': 2, 'S': 2, 'D': 2, 'F': 2, 'G': 2, 'H': 2, 'J': 2, 'K': 2, 'L': 2, ':': 2, '"': 2,
	'z': 1, 'x': 1, 'c': 1, 'v': 1, 'b': 1, 'n': 1, 'm': 1, ',': 1, '.': 1, '/': 1,
	'Z': 1, 'X': 1, 'C': 1, 'V': 1, 'B': 1, 'N': 1, 'M': 1, '<': 1, '>': 1, '?': 1,
	' ': 0,
}

var shiftChars = []rune{
	'~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+',
	'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P', '{', '}', '|',
	'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', ':', '"',
	'Z', 'X', 'C', 'V', 'B', 'N', 'M', '<', '>', '?',
}

var fingerCoefficients = map[finger]int{
	leftPinky: 2, leftRing: 1, leftMiddle: 1, leftIndex: 1, rightIndex: 1, rightMiddle: 1, rightRing: 1, rightPinky: 2, thumb: 1,
}

var rowCoefficients = map[int]int{
	4: 3, 3: 1, 2: 1, 1: 1, 0: 1,
}

var charModifiers = map[rune]int{
	'a': -1, 'A': -1, ';': -1, ':': -1,
	'6': 1, '^': 1,
	'`': 2, '~': 2,
	'=': 1, '+': 1,
	'\\': 2, '|': 2,
}

func Score(s string) int {
	score := 0
	previousRune := 'ԗ'
	previousFinger := none
	for _, r := range s {
		f := char2Finger[r]
		if r == previousRune {
			score += 1 + runeScore(r)
		} else if f == previousFinger {
			score += 2 * runeScore(r)
		} else {
			score += runeScore(r)
		}
		previousRune = r
		previousFinger = f
	}
	return score
}

func runeScore(r rune) int {
	finger := char2Finger[r]
	row := char2Row[r]
	score := (fingerCoefficients[finger] + charModifiers[r]) * rowCoefficients[row]
	if slices.Contains(shiftChars, r) {
		score += 2
	}
	if score == 0 {
		return 100
	}
	return score
}
