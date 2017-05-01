package hamming

import "errors"

const testVersion = 5

func Distance(a, b string) (int, error) {
	if len(a)-len(b) != 0 {
		return -1, errors.New("two DNA strands not equel")
	}
	dis := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dis = dis + 1
		}
	}
	return dis, nil
}
