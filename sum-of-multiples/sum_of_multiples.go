// Package summultiples provides calculations of sum of unique multiples.
package summultiples

// SumMultiples calculates the sum of the unique multiples below a value.
func SumMultiples(limit int, divisors ...int) int {
    sum := 0
    for i := 1; i < limit; i++ {
        for _, divisor := range(divisors) {
            if i % divisor == 0 {
                sum += i
                break
            }
        }
    }
    return sum
}
