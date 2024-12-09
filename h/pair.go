package h

import "cmp"

type Pair[T cmp.Ordered] struct {
	F T
	S T
}
