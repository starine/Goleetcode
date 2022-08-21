package main

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return true //
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0 //
}

/*
284. 窥探迭代器 :https://leetcode-cn.com/problems/peeking-iterator/
*/
//不会测试
type PeekingIterator struct {
	iter     *Iterator
	_hasNext bool
	_next    int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter, iter.hasNext(), iter.next()}
}

func (this *PeekingIterator) hasNext() bool {
	return this._hasNext
}

func (this *PeekingIterator) next() int {
	ret := this._next
	this._hasNext = this.iter.hasNext()
	if this._hasNext {
		this._next = this.iter.next()
	}
	return ret
}

func (this *PeekingIterator) peek() int {
	return this._next
}

//func mian () {
//	it := []int{1,2,3}
//	var peekingIterator PeekingIterator
//	peekingIterator = Constructor(it); // [1,2,3]
//	peekingIterator.next();    // 返回 1 ，指针移动到下一个元素 [1,2,3]
//	peekingIterator.peek();    // 返回 2 ，指针未发生移动 [1,2,3]
//	peekingIterator.next();    // 返回 2 ，指针移动到下一个元素 [1,2,3]
//	peekingIterator.next();    // 返回 3 ，指针移动到下一个元素 [1,2,3]
//	peekingIterator.hasNext(); // 返回 False
//}
