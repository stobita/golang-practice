package main

func main() {
	primeFactor(24)
}

func primeFactor(n int) {
	// まずは2で割れるところまで割る
	for ; n%2 == 0; n /= 2 {
		println(2)
	}
	// 3以降の奇数で割り続ける
	for i := 3; i <= n; {
		if n%i == 0 {
			println(i)
			n /= i
		} else {
			i += 2
		}
	}
}
