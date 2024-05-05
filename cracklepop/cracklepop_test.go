package cracklepop

import (
	"strings"
	"testing"
)

func TestCracklePop(t *testing.T) {
	testCases := []struct {
		description string
		start       int
		end         int
		expected    string
		err         error
	}{
		{
			description: "negative starting number input",
			start:       -4,
			end:         5,
			expected:    "",
			err:         ErrNegativeInput(-4),
		},
		{
			description: "negative ending number input",
			start:       2,
			end:         -3,
			expected:    "",
			err:         ErrNegativeInput(-3),
		},
		{
			description: "starting number is greater than ending number",
			start:       10,
			end:         3,
			expected:    "",
			err:         ErrInvalidInput{start: 10, end: 3},
		},
		{
			description: "From 1 to 15",
			start:       1,
			end:         15,
			expected: strings.Join([]string{
				"1", "2", "Crackle", "4", "Pop", "Crackle", "7", "8",
				"Crackle", "Pop", "11", "Crackle", "13", "14", "CracklePop", "",
			}, "\n"),
			err: nil,
		},
		{
			description: "From 1 to 100",
			start:       1,
			end:         100,
			expected: strings.Join([]string{
				"1", "2", "Crackle", "4", "Pop", "Crackle", "7", "8",
				"Crackle", "Pop", "11", "Crackle", "13", "14", "CracklePop", "16",
				"17", "Crackle", "19", "Pop", "Crackle", "22", "23", "Crackle",
				"Pop", "26", "Crackle", "28", "29", "CracklePop", "31", "32",
				"Crackle", "34", "Pop", "Crackle", "37", "38", "Crackle", "Pop",
				"41", "Crackle", "43", "44", "CracklePop", "46", "47", "Crackle",
				"49", "Pop", "Crackle", "52", "53", "Crackle", "Pop", "56",
				"Crackle", "58", "59", "CracklePop", "61", "62", "Crackle", "64",
				"Pop", "Crackle", "67", "68", "Crackle", "Pop", "71", "Crackle",
				"73", "74", "CracklePop", "76", "77", "Crackle", "79", "Pop",
				"Crackle", "82", "83", "Crackle", "Pop", "86", "Crackle", "88",
				"89", "CracklePop", "91", "92", "Crackle", "94", "Pop", "Crackle",
				"97", "98", "Crackle", "Pop", "",
			}, "\n"),
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			outputStr, err := CracklePop(tc.start, tc.end)
			if err != tc.err {
				t.Fatalf("mismatch error value: got %v, expected %v", err, tc.err)
			}
			if outputStr != tc.expected {
				t.Fatalf("mismatch string output: got\n%v, expected\n%v", outputStr, tc.expected)
			}
		})
	}
}
