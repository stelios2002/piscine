package piscine

// stole solution from another's chatgpt solution :/

var prime_cache []int

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	prime_cache = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	prime_found := false
	for prime_found == false {
		if IsPrime2(nb) == true {
			return nb
		}
		nb += 1
	}
	return 0
}

func IsPrime2(nb int) bool {
	for cache_i := 0; cache_i < len(prime_cache); cache_i++ {
		if nb == prime_cache[cache_i] {
			return true
		}
		if nb%prime_cache[cache_i] == 0 {
			return false
		}
	}
	for i := prime_cache[len(prime_cache)-1]; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	prime_cache = append(prime_cache, nb)
	return true
}
