// Problem 11.12

package chapter11

func FindMinMax(a []int) (int, int) {
	if len(a) <= 1 {
		return a[0], a[0]
	} else {
		var gm, gM int
		if a[0] <= a[1] {
			gm, gM = a[0], a[1]
		} else {
			gm, gM = a[1], a[0]
		}

		for i := 2; i+1 < len(a); i += 2 {
			var m, M int
			if a[i] <= a[i+1] {
				m, M = a[i], a[i+1]
			} else {
				m, M = a[i+1], a[i]
			}
			if m < gm {
				gm = m
			}
			if M > gM {
				gM = M
			}
		}

		n := len(a)
		if n&1 == 1 {
			if a[n-1] < gm {
				gm = a[n-1]
			}
			if a[n-1] > gM {
				gM = a[n-1]
			}
		}

		return gm, gM
	}
}
