package control

import (
	"log"
	"testing"

	"github.com/zxysilent/utils"
)

func TestRandStr(t *testing.T) {
	log.Println(RandStr(10))
}
func BenchmarkRandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStr(10)
	}
}
func BenchmarkRandStr11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.RandStr(10)
	}
}
