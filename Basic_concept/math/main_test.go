package math

import "testing"

type AddData struct {
	x, y   int
	result int
}

func TestAdd(t *testing.T) {

	result := Add(1, 3)

	if result != 4 {
		t.Errorf("Add(1, 3) FAILED. Expected %d, got %d\n", 4, result)
	} else {
		t.Logf("Add(1, 3) PASSED. Expected %d, got %d\n", 4, result)
	}
}

func TestADD(t *testing.T) {
	testData := []AddData{
		{1, 2, 3},
		{3, 5, 8},
		{7, -4, 3},
	}

	for _, data := range testData {
		result := Add(data.x, data.y)

		if result != data.result {
			t.Errorf("Add(%d, %d) FAILED. Expected %d got %d\n", data.x, data.y, data.result, result)
		}
	}
}

func TestDIV(t *testing.T) {
	result := Div(10, 5)

	if result != 2 {
		t.Errorf("Div(10, 5) FAILED. Expected %d, got %f\n", 2, result)
	} else {
		t.Logf("Div(10, 5) PASSED. Expected %d, got %f\n", 2, result)
	}
}
