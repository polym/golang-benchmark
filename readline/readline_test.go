package readline

import (
	"os"
	"testing"
)

var fd *os.File

func BenchmarkReadLineByBufioScanner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadLineByBufioScanner(fd)
	}
}

func BenchmarkReadLineByBufioReadLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadLineByBufioReadLine(fd)
	}
}
func BenchmarkReadLineByBufioReadSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadLineByBufioReadSlice(fd)
	}
}

func BenchmarkReadLineByDIY(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadLineByDIY(fd)
	}
}

func init() {
	fd, _ = os.Open("/tmp/t.go")
}
