package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// wrapper to avoid doing conversions in the middle of the logic
type id struct {
	i int
	s string
}

func (id *id) valid() bool {
	return true
}

type productID struct {
	firstID id
	lastID  id
}

func Exe2(lines []string, startingPoint int) int {
	return 0
}

func Exe1(lines []string, startingPoint int) int {
	accumulator := 0
	productIDs := fromLines(lines)

	for _, pID := range productIDs {
		if pID.firstID.valid() {
			accumulator += pID.firstID.i
		}

		if pID.lastID.valid() {
			accumulator += pID.lastID.i
		}
	}

	fmt.Printf("%v\n", productIDs)
	return 0
}

func fromLines(lines []string) []productID {
	productIDs := make([]productID, 0)
	for _, inputLine := range lines {
		strProductIDs := strings.Split(inputLine, ",") // get a list of all product IDs.
		for _, strPID := range strProductIDs {
			productIDs = append(productIDs, fromStrProductID(strPID))
		}
	}

	return productIDs
}

func fromStrProductID(strProductID string) productID {
	strIDs := strings.Split(strProductID, "-")

	if len(strIDs) != 2 {
		log.Fatalf("got something unexpected in the input line!")
	}

	i1, err := strconv.Atoi(strIDs[0])
	if err != nil {
		log.Fatalf("got something unexpected as an integert id from the input line!: %e", err)
	}

	i2, err := strconv.Atoi(strIDs[1])
	if err != nil {
		log.Fatalf("got something unexpected as an integert id from the input line!: %e", err)
	}

	return productID{
		firstID: id{
			i: i1,
			s: strIDs[0],
		},
		lastID: id{
			i: i2,
			s: strIDs[1],
		},
	}
}
