package main

// 掛け算九九
func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			print(i * j)
			print("\t")
		}
		print("\n")
	}
}
