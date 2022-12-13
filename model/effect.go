/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "effect.go" LastUpdatedAt 2022/12/13 16:57:13
 */

package model

import "github.com/sunist-c/genius-invokation-simulator-backend/definition"

// EffectSet Effect映射表，TriggerType->Effect单链表，Effect链表带有头节点
type EffectSet struct {
	chain map[definition.TriggerType]*EffectChainNode
}

func (e *EffectSet) AddEffect(trigger definition.TriggerType, effect IEffect) {
	if _, ok := e.chain[trigger]; ok {
		// 遍历链表寻找是否有同名Effect，如果有，则重置Effect
		var ptr *EffectChainNode
		for ptr = e.chain[trigger]; ptr.NextNode != nil; ptr = ptr.NextNode {
			if ptr.NextNode.Name == effect.Name() {
				ptr.NextNode.Effect = effect
				return
			}
		}

		// 如果没有，则添加到链表尾
		ptr.NextNode.NextNode = &EffectChainNode{
			Name:     effect.Name(),
			Effect:   effect,
			NextNode: nil,
		}
		return
	} else {
		// 如果没有触发链表，新增头节点
		e.chain[trigger] = &EffectChainNode{
			Name:     "",
			Effect:   nil,
			NextNode: nil,
		}
		e.AddEffect(trigger, effect)
	}
}

func (e *EffectSet) RemoveEffect(trigger definition.TriggerType, effect IEffect) {
	if _, ok := e.chain[trigger]; ok {
		for ptr := e.chain[trigger]; ptr.NextNode != nil; ptr = ptr.NextNode {
			if ptr.NextNode.Name == effect.Name() {
				ptr.NextNode = ptr.NextNode.NextNode
			}
		}
	}
}

func (e EffectSet) Execute(trigger definition.TriggerType, ctx *Context) {
	if _, ok := e.chain[trigger]; ok {
		for ptr := e.chain[trigger]; ptr.NextNode != nil; ptr = ptr.NextNode {
			ptr.NextNode.Effect.Effect(trigger, ctx)
		}
	}
}

type EffectChainNode struct {
	Name     string
	Effect   IEffect
	NextNode *EffectChainNode
}
