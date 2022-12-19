package modifier

type Chain[data any] struct {
	size         byte
	currentIndex byte
	currentPtr   *node[data]
	head         *node[data]
	rear         *node[data]
}

// index 找出链表第index个元素，已针对遍历进行优化
func (c *Chain[data]) index(index byte) *node[data] {
	// 可以从缓存指针开始查找，缩短路径
	if index >= c.currentIndex && c.currentPtr != nil {
		for ; c.currentPtr != nil; c.currentPtr, c.currentIndex = c.currentPtr.next, c.currentIndex+1 {
			if c.currentIndex == index {
				return c.currentPtr
			}
		}
	}

	// 缓存指针无效，从头查找
	if c.size != 0 {
		for c.currentPtr, c.currentIndex = c.head, 0; c.currentPtr != nil; c.currentPtr, c.currentIndex = c.currentPtr.next, c.currentIndex+1 {
			if c.currentIndex == index {
				return c.currentPtr
			}
		}
	}

	return nil
}

// Append 向ModifierChain中加入一个Modifier，总共不能超过127个Modifier
func (c *Chain[data]) Append(modifier Modifier[data]) {
	// 如果已存在同id的handler，覆盖之
	for ptr := c.head; ptr != nil; ptr = ptr.next {
		if ptr.modifier.Info().ID == modifier.Info().ID {
			ptr.modifier = modifier
			return
		}
	}

	// 如果没有该id的handler，则在队尾新增handler
	node := &node[data]{
		modifier: modifier,
		next:     nil,
	}
	if c.rear != nil {
		c.rear.next = node
		c.rear = node
		c.size += 1
	} else {
		c.head = node
		c.rear = node
		c.currentIndex = 0
		c.currentPtr = node
		c.size += 1
	}
}

// Remove 根据id删除ModifierChain中的某个Modifier
func (c *Chain[data]) Remove(id uint) {
	for ptr := c.head; ptr.next != nil; ptr = ptr.next {
		if ptr.next.modifier.Info().ID == id {
			if c.currentPtr == ptr.next {
				c.currentPtr = c.head
				c.currentIndex = 0
			}
			ptr.next = ptr.next.next
			c.size -= 1
		}
	}
}

// Size 返回ModifierChain的Modifier数量
func (c *Chain[data]) Size() byte {
	return c.size
}

// ResetModifiers 执行ModifierChain中所有Modifier的RoundReset
func (c Chain[data]) ResetModifiers() {
	for ptr := c.head; ptr != nil; ptr = ptr.next {
		ptr.modifier.RoundReset()
	}
}

// Clone 将当前ModifierChain复制一份，供Preview使用
func (c Chain[data]) Clone() Chain[data] {
	chain := Chain[data]{
		size:         c.size,
		currentIndex: 0,
		currentPtr:   nil,
		head:         nil,
		rear:         nil,
	}

	for ptr := c.head; ptr != nil; ptr = ptr.next {
		chain.Append(ptr.modifier.Clone())
	}

	return chain
}

func NewChain[data any]() Chain[data] {
	return Chain[data]{
		size:         0,
		currentIndex: 0,
		head:         nil,
		rear:         nil,
		currentPtr:   nil,
	}
}

type node[data any] struct {
	modifier Modifier[data]
	next     *node[data]
}
