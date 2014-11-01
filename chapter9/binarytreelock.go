// Problem 9.4

package chapter9

type BinaryTreeLock struct {
	Value                  int
	Left, Right, Parent    *BinaryTreeLock
	IsLocked               bool
	LockedDescendantsCount int
}

func (btl BinaryTreeLock) isLock() bool {
	return btl.IsLocked
}

func (btl *BinaryTreeLock) lock() bool {
	if btl.IsLocked {
		return false
	}

	if btl.LockedDescendantsCount > 0 {
		return false
	}

	var incrCountIfUnlocked func(*BinaryTreeLock) bool
	incrCountIfUnlocked = func(elem *BinaryTreeLock) bool {
		if elem == nil {
			return true
		} else if !elem.IsLocked && incrCountIfUnlocked(elem.Parent) {
			elem.LockedDescendantsCount++
			return true
		} else {
			return false
		}
	}

	if incrCountIfUnlocked(btl.Parent) {
		btl.IsLocked = true
	}

	return btl.IsLocked
}

func (btl *BinaryTreeLock) unlock() {
	if btl.IsLocked {
		btl.IsLocked = false
		for p := btl.Parent; p != nil; p = p.Parent {
			p.LockedDescendantsCount--
		}
	}
}
