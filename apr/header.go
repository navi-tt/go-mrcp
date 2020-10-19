package apr

/** An opaque array type */
type ArrayHeader struct {
	Stack *Stack // todo(换成[]T)
}

func NewArrayHeader(n int) *ArrayHeader {
	return &ArrayHeader{
		Stack: NewStack(n),
	}
}

func ArrayCopy(src *ArrayHeader) *ArrayHeader {
	a := &ArrayHeader{
		Stack: src.Stack.Clone(),
	}
	//
	//for e := src.Stack.Pop(); e != nil; e = src.Stack.Pop() {
	//	a.Stack.Push(e)
	//}
	return a
}

func (a *ArrayHeader) ArrayHeaderIndex(id int) interface{} {
	return a.Stack.Index(id)
}

func (a *ArrayHeader) Append(ah *ArrayHeader) {
	a.Stack.Append(ah.Stack)
}
