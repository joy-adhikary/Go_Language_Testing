package Basic_Concept

func add(x, y int) int {
	result := x + y
	return result
}

func Sub(x, y int) int {
	if x > y {
		return x - y
	} else {
		return 0
	}
}

func Multi(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	} else {
		return x * y
	}
}

func Div(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	} else {
		return x / y
	}
}

// 		methods

type joy struct {
	name string
}

func (j *joy) PrintName() string {
	return j.name
}
