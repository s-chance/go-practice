package slice

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestGetLastBySlice(t *testing.T) {
	testGetLast(t, GetLastBySlice)
}

func TestGetLastByCopy(t *testing.T) {
	testGetLast(t, GetLastByCopy)
}

func testGetLast(t *testing.T, f func([]int) []int) {
	result := make([][]int, 0)
	for k := 0; k < 100; k++ {
		origin := generateWithCap(128 * 1024)
		result = append(result, f(origin))
	}
	printMem(t)
	_ = result
}

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	origin := make([]int, 0, n)
	for i := 0; i < n; i++ {
		origin = append(origin, rand.Int())
	}
	return origin
}

func printMem(t *testing.T) {
	t.Helper()
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	t.Logf("%.2f MB", float64(rtm.Alloc)/1024./1024.)
}
