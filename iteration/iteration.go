package iteration

// import "testing"
import (
	"strings"
	"testing"
)

func Repeated(char string) string {
	var repeated strings.Builder
	var repeatedCount = 5

	for i := 0; i < repeatedCount; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeated("a")
	}
}
