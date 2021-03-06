package sieve

const testVersion = 1

// Sieve return an slice of size max containing
// a list of primes up to max -1
// Algo from user Hamain in PE. Translated from C.
func Sieve(max int) []int {
	primes := make([]bool, max)
	for i := 0; i < max; i++ {
		primes[i] = true
	}
	primes[0] = false
	primes[1] = false
	primes[2] = true

	var i int
	for {
		for i < max && !primes[i] {
			i++
		}
		for j := i << 1; j < max; j += i {
			primes[j] = false
		}
		i++
		if i > max {
			break
		}
	}

	primesList := make([]int, 0)
	for i, p := range primes {
		if p {
			primesList = append(primesList, i)
		}
	}

	return primesList
}
