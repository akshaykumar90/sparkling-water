// problem 5.12

package chapter5

type Rectangle struct {
	rx, ry, rw, rh int
}

func testIntersection(x11, x12, x21, x22 int) (int, int, bool) {
	if x21 < x11 {
		// swap positions
		x11, x21 = x21, x11
		x12, x22 = x22, x12
	}

	x, w := 0, 0

	if x21 < x12 {
		x = x21
		if x22 < x12 {
			w = x22 - x21
		} else {
			w = x12 - x21
		}
		return x, w, true
	} else {
		return x, w, false
	}
}

func IntersectRectangle(r1, r2 Rectangle) *Rectangle {
	x, w, xIntersection := testIntersection(r1.rx, r1.rx+r1.rw, r2.rx, r2.rx+r2.rw)
	y, h, yIntersection := testIntersection(r1.ry, r1.ry+r1.rh, r2.ry, r2.ry+r2.rh)

	if xIntersection && yIntersection {
		intersection := new(Rectangle)
		intersection.rx, intersection.ry = x, y
		intersection.rw, intersection.rh = w, h
		return intersection
	} else {
		return nil
	}
}
