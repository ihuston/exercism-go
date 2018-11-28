// Package sieve finds primes.
package sieve

// Sieve finds prime numbers below a limit using the Sieve of Eratosthenes.
func Sieve(limit int) []int {
    notPrime := make(map[int]bool)
    primes := []int{}
    for i := 2; i <= limit; i++ {
        if notPrime[i] {
            continue
        }
        primes = append(primes, i)
        for factor := 2; i*factor <= limit; factor++ {
            notPrime[i*factor] = true
        }
    }
    return primes
}
