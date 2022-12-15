/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "player.go" LastUpdatedAt 2022/12/12 15:37:12
 */

package model

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/definition"

	"errors"
)

type Player struct {
	UID               uint
	Name              string
	Deck              PlayerDeck
	HoldingCards      CardDeck
	Elements          definition.ElementSet
	ActiveCharacter   *Character
	Characters        []*Character
	StackCards        *CardStack
	CooperativeSkills []ISkill
}

type PlayerChainNode struct {
	Self       *Player
	PrevPlayer *PlayerChainNode
	NextPlayer *PlayerChainNode
}

type PlayerChain struct {
	resetPtr *PlayerChainNode
	lastPtr  *PlayerChainNode
	Head     *PlayerChainNode
}

func (pc *PlayerChain) append(node *PlayerChainNode) {
	if node == nil {
		return
	}

	if pc.Head != nil && pc.lastPtr != nil {
		node.PrevPlayer = pc.lastPtr
		node.NextPlayer = pc.Head
		pc.lastPtr = node
	} else {
		node.PrevPlayer = nil
		node.NextPlayer = nil
		pc.Head = node
		pc.lastPtr = node
		pc.resetPtr = pc.Head
	}
}

func (pc *PlayerChain) Reset() {
	pc.Head = pc.resetPtr
	if pc.Head != nil && pc.Head.PrevPlayer != nil {
		pc.lastPtr = pc.Head.PrevPlayer
	}
}

func (pc *PlayerChain) RemoveNode(node *PlayerChainNode) {
	if node == nil {
		return
	}

	if node.PrevPlayer == nil && node.NextPlayer == nil {
		pc.Head = nil
		pc.resetPtr = nil
		pc.lastPtr = nil
		return
	}

	if node.PrevPlayer != nil {
		node.PrevPlayer.NextPlayer = node.NextPlayer
	}

	if node.NextPlayer != nil {
		node.NextPlayer.PrevPlayer = node.PrevPlayer
	}

	if pc.Head == node {
		pc.Head = node.NextPlayer
		pc.resetPtr = node.NextPlayer
	}

	if pc.lastPtr == node {
		pc.lastPtr = node.PrevPlayer
	}
}

func (pc *PlayerChain) MoveNodeToNextChain(node *PlayerChainNode, nextChain PlayerChain) {
	pc.RemoveNode(node)
	nextChain.append(node)
}

func GeneratePlayerChain(players []*Player) (result PlayerChain, err error) {
	if len(players) < 2 {
		return PlayerChain{}, errors.New("incorrect playerList")
	}

	ShufflePlayerChainFunction.Shuffle(0, len(players), players)

	// 初始化头节点
	head := &PlayerChainNode{
		Self:       players[0],
		PrevPlayer: nil,
		NextPlayer: nil,
	}
	result = PlayerChain{
		Head:     head,
		lastPtr:  nil,
		resetPtr: head,
	}

	// 构造循环链表
	prev := head
	for i := 1; i < len(players); i++ {
		next := &PlayerChainNode{
			Self:       players[i],
			PrevPlayer: prev,
			NextPlayer: nil,
		}
		prev.NextPlayer = next
		prev = next
	}
	prev.NextPlayer = head
	head.PrevPlayer = prev
	result.lastPtr = prev

	return result, nil
}
