package test

func Add(max int) {
	number := 0
	for {
		number++
		if number > max {
			break
		}
	}
}

//时间复杂度:o(2^n)
func Fb(i int) int {
	if i < 0 || i <= 2 {
		return 1
	}
	return Fb(i-1) + Fb(i-2)
}

//时间复杂度为:o(n)
func Fb2(i int64) int64 {
	//1, 1, 2, 3, 5, 8
	//v1, v2, v3, (v2+v3), (v3+(v2+v3)) ()
	var v1, v2, v3 int64
	v1, v2 = 1, 1
	v3 = 0
	var x int64
	for x = 0; x < i; x++ {
		v3 = v1 + v2
		v1 = v2
		v2 = v3
	}
	return v3
}
