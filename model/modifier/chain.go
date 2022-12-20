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

// clear 清理掉ModifierChain中无效的Modifier
func (c *Chain[data]) clear() {
	removes, index := make([]uint, c.size), 0
	for ptr := c.head; ptr != nil; ptr = ptr.next {
		if !ptr.modifier.Effective() {
			removes[index] = ptr.modifier.Info().id
			index++
		}
	}

	for i := 0; i < index; i++ {
		c.Remove(removes[i])
	}
}

// Append 向ModifierChain中加入一个Modifier，总共不能超过127个Modifier
func (c *Chain[data]) Append(modifier Modifier[data]) {
	// 如果已存在同id的handler，覆盖之
	for ptr := c.head; ptr != nil; ptr = ptr.next {
		if ptr.modifier.Info().id == modifier.Info().id {
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
	if c.size == 1 {
		// 只有一个node，直接判断
		if c.head.modifier.Info().id == id {
			c.head = nil
			c.rear = nil
			c.size = 0
		}
	} else if c.size > 1 {
		// 在队首，直接改
		if c.head.modifier.Info().id == id {
			c.head = c.head.next
			c.size -= 1
			return
		}

		// 否则单链表遍历移除元素
		for ptr := c.head; ptr.next != nil; ptr = ptr.next {
			if ptr.next.modifier.Info().id == id {
				ptr.next = ptr.next.next
				if ptr.next == nil {
					c.rear = ptr
				}
				c.size -= 1
				return
			}
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
		size:         0,
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

// Execute 使用给定的ContextData执行ModifierChain
func (c *Chain[data]) Execute(ctx *data) {
	context := newContext(ctx, *c)
	context.Next()
	c.clear()
}

// Effective 当前的ModifierChain中是否有有效
func (c Chain[data]) Effective() bool {
	return c.size != 0
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
