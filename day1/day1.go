package day1

import (
	"log"
	"strconv"

	"github.com/daveseco7/advent-of-code-2025/util"
)

const regex = `\b(?P<direction>R|L)(?P<amount>\d{1,4})\b`

type direction string

const (
	l direction = "L"
	r direction = "R"
)

type instruction struct {
	value     int
	direction direction
}

func Exe2(lines []string, startingPoint int) int {
	circularCounterMax := 99
	numberOfZeros := 0
	pointer := startingPoint

	for _, line := range lines {
		op := fromLine(line)

		switch op.direction {
		case l:
			quotient := floorDiv(pointer-op.value, circularCounterMax+1)
			remainder := mod(pointer-op.value, circularCounterMax+1)

			// reset position to 0, since 100 is not a valid position.
			if remainder == circularCounterMax+1 {
				remainder = 0
			}

			numberOfZeros += util.Abs(quotient)

			if pointer == 0 && remainder > 0 {
				numberOfZeros--
			}
			if pointer > 0 && remainder == 0 {
				numberOfZeros++
			}

			pointer = remainder

		case r:
			quotient := floorDiv(pointer+op.value, circularCounterMax+1)
			remainder := mod(pointer+op.value, circularCounterMax+1)

			numberOfZeros += quotient
			pointer = remainder

		default:
			log.Fatalf("unsupported direction value provided")
		}
	}

	return numberOfZeros
}

func Exe1(lines []string, startingPoint int) int {
	circularCounterMax := 99
	numberOfZeros := 0
	pointer := startingPoint

	for _, line := range lines {
		op := fromLine(line)

		switch op.direction {
		case l:
			// go down with overflow to 100, if 0 is reached.
			pointer = (pointer - op.value + (circularCounterMax + 1)) % (circularCounterMax + 1)
		case r:
			// go up with overflow to 0, if 100 is reached.
			pointer = (pointer + op.value) % (circularCounterMax + 1)
		default:
			log.Fatalf("unsupported direction value provided")
		}

		if pointer == 0 {
			numberOfZeros++
		}
	}

	return numberOfZeros
}

func fromLine(line string) instruction {
	captureGroups := util.MustGetFromCaptureGroups(regex, line)

	if _, ok := captureGroups["amount"]; !ok {
		log.Fatalf("something went while extracting the capture groups from the regex")
	}

	if _, ok := captureGroups["direction"]; !ok {
		log.Fatalf("something went while extracting the capture groups from the regex")
	}

	value, err := strconv.Atoi(captureGroups["amount"])
	if err != nil {
		log.Fatalf("invalid value provided in the input")
	}

	return instruction{
		value:     value,
		direction: direction(captureGroups["direction"]),
	}
}

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}

	return m
}

func floorDiv(a, b int) int {
	q := a / b
	r := a % b
	if r != 0 && ((a < 0) != (b < 0)) {
		q--
	}
	return q
}
