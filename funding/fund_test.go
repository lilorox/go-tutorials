package funding

import "testing"

func BenchmarkFund(b *testing.B) {
    // Adds as much funds as we have iterations
    fund := NewFund(b.N)
}
