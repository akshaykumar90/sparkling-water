// problem 8.5

package chapter8

import "fmt"

func _hanoi(n int, from, to, aux string) {
	if n > 0 {
		_hanoi(n-1, from, aux, to)
		fmt.Printf("Ring %d : %s => %s\n", n, from, to)
		_hanoi(n-1, aux, to, from)
	}
}

func TowerHanoi(n int) {
	_hanoi(n, "p1", "p2", "p3")
}
