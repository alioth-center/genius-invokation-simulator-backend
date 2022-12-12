/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "deck.go" LastUpdatedAt 2022/12/12 14:44:12
 */

package model

// CardStack 牌堆
type CardStack struct {
	cards []ICard
	index uint
}

// Length 牌堆的剩余可用牌数
func (cs CardStack) Length() uint {
	return uint(len(cs.cards)) - cs.index
}

// Takeout 从牌堆中取出若干张牌，若数量不够，则取出剩余的所有卡牌
func (cs *CardStack) Takeout(number uint) (cards []ICard, take uint) {
	cards = make([]ICard, 0)

	if cs.index+number < uint(len(cs.cards)) {
		cards = append(cards, cs.cards[cs.index:]...)
		length := uint(len(cards))
		cs.index = cs.index + length
		return cards, length
	} else {
		cards = append(cards, cs.cards[cs.index:cs.index+number]...)
		cs.index = cs.index + number
		return cards, number
	}
}

// Shuffle 将牌堆打乱，调用ShuffleCardStackInterface，需指定实现
func (cs *CardStack) Shuffle(start, end int) {
	ShuffleCardStackFunction.Shuffle(start, end, cs.cards)
}

// CardDeck 卡组，键为卡牌，值为数量
type CardDeck map[ICard]uint

// GenerateCardStack 使用玩家的卡组生成一个乱序的牌堆
func (cd CardDeck) GenerateCardStack() CardStack {
	result := CardStack{
		cards: make([]ICard, len(cd), 0),
		index: 0,
	}

	for card, count := range cd {
		for i := uint(0); i < count; i++ {
			result.cards = append(result.cards, card)
		}
	}

	result.Shuffle(0, len(result.cards))

	return result
}

// CharacterDeck 角色列表，键为角色，值为数量
type CharacterDeck map[ICharacter]uint

// PlayerDeck 玩家的出战卡组
type PlayerDeck struct {
	Name       string
	Characters CharacterDeck
	Cards      CardDeck
}
