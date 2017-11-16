package main

func main() {
	arr := elratosthenes(100)

	for i := range arr {
		if arr[i] == true {
			print(i)
			print("\t")
		}
	}
}

func elratosthenes(num int) []bool {
	prime_case := make([]bool, num)
	for i := range prime_case {
		prime_case[i] = true
	}
	prime_case[0], prime_case[1] = false, false

	for i := range prime_case {
		if prime_case[i] == true {
			screen(i, num, prime_case)
		}
	}
	return prime_case
}

func screen(num int, term int, arr []bool) {
	for i := num; i*num < term; i++ {
		arr[i*num] = false
	}
}
