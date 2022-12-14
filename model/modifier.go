/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "modifier.go" LastUpdatedAt 2022/12/14 13:32:14
 */

package model

const (
	abortIndex byte = 1<<8 - 1
)

type ModifierHandler[C any] func(ctx *ModifierContext[C])

// ModifierContext 供Modifier提供上下文的Context
type ModifierContext[C any] struct {
	index byte
	chain []ModifierHandler[C]
	Data  C
}

func (c *ModifierContext[C]) Continue() {
	c.index++
	for c.index < byte(len(c.chain)) {
		c.chain[c.index](c)
		c.index++
	}
}

func (c *ModifierContext[C]) Abort() {
	c.index = abortIndex
}

func (c *ModifierContext[C]) IsAborted() bool {
	return c.index >= abortIndex
}

func NewContext[C any](modifiers ModifierChain[C], ctx C) *ModifierContext[C] {
	return &ModifierContext[C]{
		index: 0,
		chain: modifiers.generateModifierHandlerChain(),
		Data:  ctx,
	}
}

type ModifierChain[C any] struct {
	names []string
	chain []ModifierHandler[C]
}

func (m *ModifierChain[C]) AddModifierHandler(name string, handler ModifierHandler[C]) {
	m.names = append(m.names, name)
	m.chain = append(m.chain, handler)
}

func (m *ModifierChain[C]) RemoveModifierHandler(name string) {
	for i, n := range m.names {
		if name == n {
			m.names = append(m.names[:i], m.names[i+1:]...)
			m.chain = append(m.chain[:i], m.chain[i+1:]...)
		}
	}
}

func NewModifierChain[C any]() *ModifierChain[C] {
	return &ModifierChain[C]{
		names: []string{},
		chain: []ModifierHandler[C]{},
	}
}

func (m *ModifierChain[C]) generateModifierHandlerChain() []ModifierHandler[C] {
	return m.chain
}
