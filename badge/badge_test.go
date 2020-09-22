package badge

import (
	"testing"

	"github.com/gofiber/fiber/v2/utils"
)

// go test -v -run=^$ -bench=Benchmark_Template -benchmem -count=4
func Benchmark_Template(b *testing.B) {
	var badge []byte
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		badge = Generate("viewcounter", "10000", "0000")
	}
	utils.AssertEqual(b, 1157, len(badge))
}
