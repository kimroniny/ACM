package ct3

import (
	"fmt"
	"math"
)

func sumGame(num string) bool {
	l := len(num)
	cl, cr := 0, 0
	sl, sr := 0, 0
	for k, v := range num {
		if v == '?' {
			if k < l/2 {
				cl += 1
			} else {
				cr += 1
			}
		} else {
			if k < l/2 {
				sl += int(v) - int('0')
			} else {
				sr += int(v) - int('0')
			}
		}
	}
	cd := int(math.Abs(float64(cl - cr)))
	if cd&1 == 1 {
		return true
	}
	sd := int(math.Abs(float64(sl - sr)))
	d := sd - cd/2*9
	if !(d > 9 || d < 0) {
		return d != 0
	}
	return true
}

func Main() {
	fmt.Println(
		sumGame("?3295???"),
	)
}
