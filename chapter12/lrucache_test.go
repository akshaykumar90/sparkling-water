package chapter12

import "testing"

type Book struct {
	Isbn  string
	Price int
}

func (b Book) Key() string {
	return b.Isbn
}

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Set(Book{"a", 100})
	cache.Set(Book{"a", 150})

	if cache.Get("b") != nil {
		t.Fatalf("got b, want nil")
	}

	if cache.Get("a") == nil {
		t.Fatalf("want a, got nil")
	}

	if p := cache.Get("a").(Book).Price; p != 150 {
		t.Fatalf("p = %d, want 150", p)
	}

	cache.Set(Book{"b", 10})
	cache.Set(Book{"c", 15})
	cache.Set(Book{"d", 15})
	cache.Set(Book{"e", 15})

	if cache.Get("a") != nil {
		t.Fatalf("got a, want nil")
	}
}
