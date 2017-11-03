package main

func main() {
	println(sumOfInt(1, 4))
	println(sumOfSquare(1, 3))
	println(sumOfCube(1, 3))

}
func sumOfInt(n, m int) int {
	sum := 0
	for i := n; i <= m; i++ {
		sum = sum + i
	}
	return sum
}
func sumOfSquare(n, m int) int {
	sum := 0
	for i := n; i <= m; i++ {
		sum = sum + i*i
	}
	return sum
}
func sumOfCube(n, m int) int {
	sum := 0
	for i := n; i <= m; i++ {
		sum = sum + i*i*i
	}
	return sum
}
