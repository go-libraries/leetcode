package _07reverse_integer

const IntMax = 2147483647 //math.MaxInt32
const max  = (1<<31)-1
const IntMin = -2147483648
const min  = -(1<<31)

func reverse(x int) int {
	//x1 := int32(x)

	if x < 10 && x > -1 {
		return x
	}

	isLtZero := false
	if x < 0 {
		x = x * -1
		isLtZero = true
	}
	var i int
	var v int
	var result int
	for {
		t := i * 10
		if x < t {
			break
		}
		if t == 0 {
			v = x % 10
		} else {
			v = x / t % 10
		}
		if i == 0 {
			i = 1
			result = v
		}else{
			i = i *10
			result = result * 10 + v
		}
	}
	if isLtZero {
		result *= -1
	}
	if result > IntMax {
		return 0
	}
	if result < IntMin {
		return 0
	}

	return result
}