package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
)

type Card interface {
	ID() uint
	Type() enum.CardType
}

type FoodCard interface {
	Card
	ExecuteModify(ctx *context.ModifierContext)
}

// CardDeck 牌堆，记录顺序的是一个数组，正常情况下已取出的牌永远在队列的一端
type CardDeck struct {
	cards  kv.Map[uint, Card]
	used   kv.Map[uint, bool]
	queue  []uint
	offset int
	remain uint
}

// arrange 整理CardDeck中offset之后的部分
func (c *CardDeck) arrange() {
	for i := c.offset; i < len(c.queue); i++ {
		if c.used.Get(c.queue[i]) {
			c.queue[c.offset], c.queue[i] = c.queue[i], c.queue[c.offset]
			c.offset++
			c.remain--
		}
	}
}

// takeOne 将CardDeck的队列中第index张牌标记为已取出
func (c *CardDeck) takeOne(index int) {
	c.used.Set(c.queue[index], true)
	c.queue[index], c.queue[c.offset] = c.queue[c.offset], c.queue[index]
	c.offset++
	c.remain--
}

// Shuffle 将CardDeck中的未取出部分进行洗牌
func (c *CardDeck) Shuffle() {
	random.Shuffle(len(c.queue)-c.offset, func(i, j int) {
		c.queue[i+c.offset], c.queue[j+c.offset] = c.queue[j+c.offset], c.queue[i+c.offset]
	})
}

// GetOne 从CardDeck中取出一张牌
func (c *CardDeck) GetOne() (result Card, success bool) {
	if c.remain != 0 {
		for i := c.offset; i < len(c.queue); i++ {
			if !c.used.Get(c.queue[i]) {
				result = c.cards.Get(c.queue[i])
				c.takeOne(i)
				return result, true
			}
		}
	}

	return nil, false
}

// FindOne 从CardDeck中取出一张指定类型的牌
func (c *CardDeck) FindOne(cardType enum.CardType) (result Card, success bool) {
	for i := c.offset; i < len(c.queue); i++ {
		if !c.used.Get(c.queue[i]) && c.cards.Get(c.queue[i]).Type() == cardType {
			result = c.cards.Get(c.queue[i])
			c.takeOne(i)
			return result, true
		}
	}

	return nil, false
}

// Reset 将牌堆中除了holding之外的牌全部标记为未取出，此方法没有洗牌逻辑
func (c *CardDeck) Reset(holding []uint) {
	c.used.Range(func(key uint, value bool) bool {
		c.used.Set(key, false)
		return true
	})

	c.remain = uint(len(c.queue))
	c.offset = 0
	for _, id := range holding {
		c.used.Set(id, true)
	}

	c.arrange()
}

// Remain 获取CardDeck还可以获取多少张牌
func (c CardDeck) Remain() uint {
	return c.remain
}

func NewCardDeck(cards []Card) *CardDeck {
	cardDeck := &CardDeck{
		cards:  kv.NewSimpleMap[Card](),
		used:   kv.NewSimpleMap[bool](),
		queue:  []uint{},
		offset: 0,
		remain: 0,
	}

	for _, card := range cards {
		cardDeck.cards.Set(card.ID(), card)
		cardDeck.used.Set(card.ID(), false)
		cardDeck.queue = append(cardDeck.queue, card.ID())
		cardDeck.remain++
	}

	return cardDeck
}
