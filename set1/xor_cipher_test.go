package main

import (
    "testing"
)

type testXorAgainstCharPairs struct {
    inputString string
    inputChar byte
    expected string
}

type testScorePlaintextPairs struct {
    input string
    expected float32
}

var testXorAgainstCharCases = []testXorAgainstCharPairs {
    {
        inputString: "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
        inputChar: 99, //c
        expected: "7854545052555c1b76781c481b5752505e1b5a1b4b544e555f1b545d1b595a585455",
    },
    {
        inputString: "",
        inputChar: 99,
        expected: "",
    },
    {
        inputString: "",
        inputChar: 0,
        expected: "",
    },
}

var textScorePlaintextCases = []testScorePlaintextPairs {
    {
        input: "1234556",
        expected: float32(0),
    },
    {
        input: "",
        expected: float32(0),
    },
    {
        input: "eee",
        expected: float32(38.1),
    },
    {
        input: "this is a test",
        expected: float32(99),
    },
}

func TestXorAgainstChar(t *testing.T) {

    for _, testCase := range testXorAgainstCharCases {
        actual, _, _ := XorAgainstChar(testCase.inputString, testCase.inputChar)

        if actual != testCase.expected {
            t.Error("Expected:", testCase.expected, "but got:", actual)
        }
    }
}

func TestScorePlaintext(t *testing.T) {

    for _, testCase := range textScorePlaintextCases {
        actual := ScorePlaintext(testCase.input)

        if actual != testCase.expected {
            t.Error("Expected:", testCase.expected, "but got:", actual)
        }

    }
}