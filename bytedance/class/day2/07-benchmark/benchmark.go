package benchmark

import (
	"github.com/bytedance/gopkg/lang/fastrand"
)

var ServerIndex [10]int

func InitServerIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = i + 100
	}
}

func Select() int {
	return ServerIndex[fastrand.Intn(10)]
}
