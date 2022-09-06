package main

import "testing"

func TestCalculate(t *testing.T) {

	if Calculate(5, 1) != 12 {
		t.Error("expected 10*2=20")
	}

}

func TestInfo(t *testing.T) {

	result := "joyadhikary"
	if info("joy", "adhikary", 34) != result {
		t.Error("Error occured")
	}

}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input1   int
		input2   int
		expected int
	}{
		{1, 2, 6},
		{2, 3, 10},
		{1, 5, 12},
	}

	for _, test := range tests {
		result := Calculate(test.input1, test.input2)
		if result != test.expected {
			t.Error("Test Failed : {}input1 , {}input2 ,{}expected ,received:{}", test.input1, test.input2, test.expected, result)
		}
	}

}

func TestRepeated(t *testing.T) {

	var datas = []struct {
		input  int
		output int
	}{
		{5, 7},
		{1, 3},
		{109, 111},
	}

	for _, data := range datas {
		result := cal(data.input)
		if result != data.output {
			t.Error("Test Failed :{}input , {}output, received:{}", data.input, data.output, result)
		}
	}
}
