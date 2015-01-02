package chapter7

import "testing"

func verifyPostingsList(fst, snd *PostingsListElement) bool {
	if fst == nil && snd == nil {
		return true
	} else if fst != nil && snd != nil {
		if fst.Value != snd.Value {
			return false
		}
		if snd.Jump == nil || fst.Jump.Value != snd.Jump.Value {
			return false
		}
		return verifyPostingsList(fst.Next, snd.Next)
	} else {
		return false
	}
}

func TestCopyingPostingsList(t *testing.T) {
	d := &PostingsListElement{Value: 4}
	c := &PostingsListElement{Value: 3, Next: d}
	b := &PostingsListElement{Value: 2, Next: c}
	a := &PostingsListElement{Value: 1, Next: b}

	a.Jump = c
	b.Jump = d
	c.Jump = b
	d.Jump = d

	expected := a
	actual := CopyingPostingsList(a)

	if !verifyPostingsList(expected, actual) {
		t.Fatalf("CopyingPostingsList: verifyPostingsList failed!")
	}
}
