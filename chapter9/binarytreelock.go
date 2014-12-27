// Problem 9.4

package chapter9

type BinaryTreeLock struct {
	Value                  int
	Left, Right, Parent    *BinaryTreeLock
	isLocked               bool
	lockedDescendantsCount int
}

func (btl *BinaryTreeLock) IsLock() bool {
	return btl.isLocked
}

func (btl *BinaryTreeLock) Lock() bool {
	if btl.isLocked {
		return false
	}

	if btl.lockedDescendantsCount > 0 {
		return false
	}

	var incrCountIfUnlocked func(*BinaryTreeLock) bool
	incrCountIfUnlocked = func(elem *BinaryTreeLock) bool {
		if elem == nil {
			return true
		} else if !elem.isLocked && incrCountIfUnlocked(elem.Parent) {
			elem.lockedDescendantsCount++
			return true
		} else {
			return false
		}
	}

	if incrCountIfUnlocked(btl.Parent) {
		btl.isLocked = true
	}

	return btl.isLocked
}

func (btl *BinaryTreeLock) Unlock() {
	if btl.isLocked {
		btl.isLocked = false
		for p := btl.Parent; p != nil; p = p.Parent {
			p.lockedDescendantsCount--
		}
	}
}
