package chapter14

import "testing"

// Figure 14.2 - A BST with duplicate keys.
var d = &TreeNode{-14, nil, nil}
var e = &TreeNode{2, nil, nil}
var c = &TreeNode{-10, d, e}
var f = &TreeNode{108, nil, nil}
var b = &TreeNode{108, c, f}
var j = &TreeNode{401, nil, nil}
var i = &TreeNode{285, nil, j}
var h = &TreeNode{243, nil, nil}
var g = &TreeNode{285, h, i}
var a = &TreeNode{108, b, g}

var searchBSTForFirstOccurrenceTests = []struct {
	k        int
	expected *TreeNode
}{
	{108, b},
	{285, g},
	{143, nil},
	{2, e},
}

func TestSearchBSTForFirstOccurrenceRecursive(t *testing.T) {
	for _, tt := range searchBSTForFirstOccurrenceTests {
		actual := SearchBSTForFirstOccurrenceRecursive(a, tt.k)
		if actual != tt.expected {
			t.Errorf("SearchBSTForFirstOccurrenceRecursive(%v, %d): expected %v, actual %v",
				a, tt.k, tt.expected, actual)
		}
	}
}

func TestSearchBSTForFirstOccurrenceIterative(t *testing.T) {
	for _, tt := range searchBSTForFirstOccurrenceTests {
		actual := SearchBSTForFirstOccurrenceIterative(a, tt.k)
		if actual != tt.expected {
			t.Errorf("SearchBSTForFirstOccurrenceIterative(%v, %d): expected %v, actual %v",
				a, tt.k, tt.expected, actual)
		}
	}
}
