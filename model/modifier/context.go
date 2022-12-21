package modifier

const abortIndex byte = 1 << 7

type Context[data any] struct {
	index     byte
	chain     Chain[data]
	data      *data
	extendMap map[string]any
}

// Next 让Context继续执行事件链
func (c *Context[data]) Next() {
	c.index += 1
	for c.index <= c.chain.size {
		c.chain.index(c.index - 1).modifier.Handler()(c)
		c.index += 1
	}
}

// Abort 中断Context事件链的调用，但是在终止之前的调用仍会生效
func (c *Context[data]) Abort() {
	c.index = abortIndex
}

// IsAborted 判断Context是否被某一handler中断了
func (c Context[data]) IsAborted() bool {
	return c.index >= abortIndex
}

// Data 获取Context中的数据，只读
func (c Context[data]) Data() *data {
	return c.data
}

// Get 在Context的extendMap中获取数据，不建议使用这种方法传递信息
func (c Context[data]) Get(key string) (value any, ok bool) {
	value, ok = c.extendMap[key]
	return value, ok
}

// Set 在Context的extendMap中写入数据，不建议使用这种方法传递信息
func (c *Context[data]) Set(key string, value any) {
	c.extendMap[key] = value
}

// newContext 使用给定的data和handlers生成Context
func newContext[data any](ctx *data, handlers Chain[data]) *Context[data] {
	return &Context[data]{
		index:     0,
		chain:     handlers,
		data:      ctx,
		extendMap: map[string]any{},
	}
}
