package main

// 素数計算
func main() {
	print(2)
	print("\t")
	divisible := false
	for i := 3; i <= 1000; i += 2 {
		divisible = false
		for j := 3; j < i; j += 2 {
			if i%j == 0 {
				divisible = true
				break
			}
		}
		if !divisible {
			print(i)
			print("\t")
		}
	}
}
