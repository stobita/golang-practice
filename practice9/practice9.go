package main

func main() {
	println(polygonalNum(3, 1))
	println(polygonalNum(3, 2))
	println(polygonalNum(3, 3))
	println(polygonalNum(4, 1))
	println(polygonalNum(4, 2))
	println(polygonalNum(4, 3))
	println(polygonalNum(5, 1))
	println(polygonalNum(5, 2))
	println(polygonalNum(5, 3))
}

func polygonalNum(p, n int) int {
	answer := 1
	up := p - 1
	for i := 1; i < n; i++ {
		answer += up
		up += (p - 2)
	}
	return answer
}
