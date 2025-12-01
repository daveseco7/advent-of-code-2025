package day1_test

import (
	"testing"

	"github.com/daveseco7/advent-of-code-2025/day1"
	"github.com/daveseco7/advent-of-code-2025/util"
)

func TestExe1(t *testing.T) {
	const startingPoint = 50

	t.Run("puzzle sample", func(_ *testing.T) {
		const expected = 3
		const filePath = "sample1.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day1.Exe1(lineArray, startingPoint)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})

	t.Run("puzzle input", func(_ *testing.T) {
		const expected = 1066
		const filePath = "input1.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day1.Exe1(lineArray, startingPoint)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
}

func TestExe2(t *testing.T) {
	const startingPoint = 50

	t.Run("puzzle sample", func(_ *testing.T) {
		const expected = 6
		const filePath = "sample1.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day1.Exe2(lineArray, startingPoint)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})

	t.Run("puzzle short sample", func(_ *testing.T) {
		tt := []struct {
			lines         []string
			startingPoint int
			expectation   int
		}{
			{
				lines:         []string{"R1000"},
				startingPoint: startingPoint,
				expectation:   10,
			},
			{
				lines:         []string{"L1000"},
				startingPoint: startingPoint,
				expectation:   10,
			},
			{
				lines:         []string{"L50", "R50"},
				startingPoint: startingPoint,
				expectation:   1,
			},
			{
				lines:         []string{"R50", "L50"},
				startingPoint: startingPoint,
				expectation:   1,
			},
			{
				lines:         []string{"L50", "L50"},
				startingPoint: startingPoint,
				expectation:   1,
			},
			{
				lines:         []string{"R50", "R50"},
				startingPoint: startingPoint,
				expectation:   1,
			},
			{
				lines:         []string{"L150", "L50"},
				startingPoint: startingPoint,
				expectation:   2,
			},
			{
				lines:         []string{"R150", "R50"},
				startingPoint: startingPoint,
				expectation:   2,
			},
			{
				lines:         []string{"L150", "R50"},
				startingPoint: startingPoint,
				expectation:   2,
			},
			{
				lines:         []string{"R150", "L50"},
				startingPoint: startingPoint,
				expectation:   2,
			},
			{
				lines:         []string{"R966", "L774", "R26", "R57"},
				startingPoint: startingPoint,
				expectation:   19,
			},
			{
				lines:         []string{"L50"},
				startingPoint: 11,
				expectation:   1,
			},
		}

		for _, tc := range tt {
			ans := day1.Exe2(tc.lines, tc.startingPoint)
			if ans != tc.expectation {
				t.Errorf("invalid response %d", ans)
			}
		}
	})

	t.Run("puzzle input", func(_ *testing.T) {
		const expected = 6223
		const filePath = "input1.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day1.Exe2(lineArray, startingPoint)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
}
