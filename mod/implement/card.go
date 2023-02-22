package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
)

type CardImpl struct {
	EntityImpl
	cardType enum.CardType
	cost     map[enum.ElementType]uint
}

func (impl *CardImpl) CardType() enum.CardType {
	return impl.cardType
}

func (impl *CardImpl) Cost() map[enum.ElementType]uint {
	if impl.cost == nil {
		impl.cost = map[enum.ElementType]uint{}
	}

	return impl.cost
}

type CardOptions func(option *CardImpl)

func WithCardID(id uint16) CardOptions {
	return func(option *CardImpl) {
		option.InjectTypeID(uint64(id))
	}
}

func WithCardType(cardType enum.CardType) CardOptions {
	return func(option *CardImpl) {
		option.cardType = cardType
	}
}

func WithCardCost(cost map[enum.ElementType]uint) CardOptions {
	return func(option *CardImpl) {
		option.cost = cost
	}
}

func NewCardWithOpts(options ...CardOptions) definition.Card {
	impl := &CardImpl{}
	for _, option := range options {
		option(impl)
	}

	return impl
}

type EventCardImpl struct {
	CardImpl
	event definition.Event
}

func (impl *EventCardImpl) Event() definition.Event {
	return impl.event
}

type EventCardOptions func(option *EventCardImpl)

func WithEventCardID(id uint16) EventCardOptions {
	return func(option *EventCardImpl) {
		opt := WithCardID(id)
		opt(&option.CardImpl)
	}
}

func WithEventCardCost(cost map[enum.ElementType]uint) EventCardOptions {
	return func(option *EventCardImpl) {
		opt := WithCardCost(cost)
		opt(&option.CardImpl)
	}
}

func WithEventCardSubType(subType enum.CardSubType) EventCardOptions {
	return func(option *EventCardImpl) {
		opt := WithCardType(subType)
		opt(&option.CardImpl)
	}
}

func WithEventCardEvent(event definition.Event) EventCardOptions {
	return func(option *EventCardImpl) {
		option.event = event
	}
}

func NewEventCardWithOpts(options ...EventCardOptions) definition.EventCard {
	impl := &EventCardImpl{}
	cardTypeOption := WithCardType(enum.CardEvent)
	cardTypeOption(&impl.CardImpl)

	for _, option := range options {
		option(impl)
	}

	return impl
}

type FoodCardImpl struct {
	EventCardImpl
}

type FoodCardOptions func(option *FoodCardImpl)

func WithFoodCardID(id uint16) FoodCardOptions {
	return func(option *FoodCardImpl) {
		opt := WithEventCardID(id)
		opt(&option.EventCardImpl)
	}
}

func WithFoodCardCost(cost map[enum.ElementType]uint) FoodCardOptions {
	return func(option *FoodCardImpl) {
		opt := WithCardCost(cost)
		opt(&option.EventCardImpl.CardImpl)
	}
}

func WithFoodCardEvent(event definition.Event) FoodCardOptions {
	return func(option *FoodCardImpl) {
		opt := WithEventCardEvent(event)
		opt(&option.EventCardImpl)
	}
}

func NewFoodCardWithOpts(options ...FoodCardOptions) definition.EventCard {
	impl := &FoodCardImpl{}
	cardTypeOption := WithEventCardSubType(enum.CardFood)
	cardTypeOption(&impl.EventCardImpl)

	for _, option := range options {
		option(impl)
	}

	return impl
}

type ElementalResonanceCardImpl struct {
	EventCardImpl
}

type ElementalResonanceCardOptions func(option *ElementalResonanceCardImpl)

func WithElementalResonanceCardID(id uint16) ElementalResonanceCardOptions {
	return func(option *ElementalResonanceCardImpl) {
		opt := WithEventCardID(id)
		opt(&option.EventCardImpl)
	}
}

func WithElementalResonanceCardCost(cost map[enum.ElementType]uint) ElementalResonanceCardOptions {
	return func(option *ElementalResonanceCardImpl) {
		opt := WithEventCardCost(cost)
		opt(&option.EventCardImpl)
	}
}

func WithElementalResonanceCardEvent(event definition.Event) ElementalResonanceCardOptions {
	return func(option *ElementalResonanceCardImpl) {
		opt := WithEventCardEvent(event)
		opt(&option.EventCardImpl)
	}
}

func NewElementalResonanceCardWithOpts(options ...ElementalResonanceCardOptions) definition.EventCard {
	impl := &ElementalResonanceCardImpl{}
	cardTypeOption := WithEventCardSubType(enum.CardElementalResonance)
	cardTypeOption(&impl.EventCardImpl)

	for _, option := range options {
		option(impl)
	}

	return impl
}

type EquipmentCardImpl struct {
	CardImpl
	equipmentType enum.EquipmentType
	modify        definition.Event
}

func (impl *EquipmentCardImpl) EquipmentType() enum.EquipmentType {
	return impl.equipmentType
}

func (impl *EquipmentCardImpl) Modify() definition.Event {
	return impl.modify
}

type EquipmentCardOptions func(option *EquipmentCardImpl)

func WithEquipmentCardID(id uint16) EquipmentCardOptions {
	return func(option *EquipmentCardImpl) {
		opt := WithCardID(id)
		opt(&option.CardImpl)
	}
}

func WithEquipmentCardSubType(subType enum.CardSubType) EquipmentCardOptions {
	return func(option *EquipmentCardImpl) {
		opt := WithCardType(subType)
		opt(&option.CardImpl)
	}
}

func WithEquipmentCardCost(cost map[enum.ElementType]uint) EquipmentCardOptions {
	return func(option *EquipmentCardImpl) {
		opt := WithCardCost(cost)
		opt(&option.CardImpl)
	}
}

func WithEquipmentCardEquipmentType(equipmentType enum.EquipmentType) EquipmentCardOptions {
	return func(option *EquipmentCardImpl) {
		option.equipmentType = equipmentType
	}
}

func WithEquipmentCardModify(modify definition.Event) EquipmentCardOptions {
	return func(option *EquipmentCardImpl) {
		option.modify = modify
	}
}

func NewEquipmentCardWithOpts(options ...EquipmentCardOptions) definition.EquipmentCard {
	impl := &EquipmentCardImpl{}
	cardTypeOption := WithCardType(enum.CardEquipment)
	cardTypeOption(&impl.CardImpl)

	for _, option := range options {
		option(impl)
	}

	return impl
}

type ArtifactCardImpl struct {
	EquipmentCardImpl
}

type ArtifactCardOptions func(option *ArtifactCardImpl)

func WithArtifactCardID(id uint16) ArtifactCardOptions {
	return func(option *ArtifactCardImpl) {
		opt := WithEquipmentCardID(id)
		opt(&option.EquipmentCardImpl)
	}
}

func WithArtifactCardCost(cost map[enum.ElementType]uint) ArtifactCardOptions {
	return func(option *ArtifactCardImpl) {
		opt := WithEquipmentCardCost(cost)
		opt(&option.EquipmentCardImpl)
	}
}

func WithArtifactCardModify(modify definition.Event) ArtifactCardOptions {
	return func(option *ArtifactCardImpl) {
		opt := WithEquipmentCardModify(modify)
		opt(&option.EquipmentCardImpl)
	}
}

func NewArtifactCardWithOpts(options ...ArtifactCardOptions) definition.EquipmentCard {
	impl := &ArtifactCardImpl{}
	cardTypeOption := WithCardType(enum.CardEquipment)
	cardTypeOption(&impl.CardImpl)
	equipmentTypeOption := WithEquipmentCardEquipmentType(enum.EquipmentArtifact)
	equipmentTypeOption(&impl.EquipmentCardImpl)

	for _, option := range options {
		option(impl)
	}

	return impl
}

type TalentCardImpl struct {
	EquipmentCardImpl
}

type TalentCardOptions func(option *TalentCardImpl)

func WithTalentCardID(id uint16) TalentCardOptions {
	return func(option *TalentCardImpl) {
		opt := WithEquipmentCardID(id)
		opt(&option.EquipmentCardImpl)
	}
}

func WithTalentCardCost(cost map[enum.ElementType]uint) TalentCardOptions {
	return func(option *TalentCardImpl) {
		opt := WithEquipmentCardCost(cost)
		opt(&option.EquipmentCardImpl)
	}
}

func WithTalentModify(modify definition.Event) TalentCardOptions {
	return func(option *TalentCardImpl) {
		opt := WithEquipmentCardModify(modify)
		opt(&option.EquipmentCardImpl)
	}
}

func NewTalentCardWithOpts(options ...TalentCardOptions) definition.EquipmentCard {
	impl := &TalentCardImpl{}
	cardTypeOption := WithCardType(enum.CardEquipment)
	cardTypeOption(&impl.CardImpl)
	equipmentTypeOption := WithEquipmentCardEquipmentType(enum.EquipmentTalent)
	equipmentTypeOption(&impl.EquipmentCardImpl)

	for _, option := range options {
		option(impl)
	}

	return impl
}

type WeaponCardImpl struct {
	EquipmentCardImpl
	weaponType enum.WeaponType
}

func (impl *WeaponCardImpl) WeaponType() enum.WeaponType {
	return impl.weaponType
}

type WeaponCardOptions func(option *WeaponCardImpl)

func WithWeaponCardID(id uint16) WeaponCardOptions {
	return func(option *WeaponCardImpl) {
		opt := WithEquipmentCardID(id)
		opt(&option.EquipmentCardImpl)
	}
}

func WithWeaponCardCardCost(cost map[enum.ElementType]uint) WeaponCardOptions {
	return func(option *WeaponCardImpl) {
		opt := WithEquipmentCardCost(cost)
		opt(&option.EquipmentCardImpl)
	}
}

func WithWeaponCardModify(modify definition.Event) WeaponCardOptions {
	return func(option *WeaponCardImpl) {
		opt := WithEquipmentCardModify(modify)
		opt(&option.EquipmentCardImpl)
	}
}

func WithWeaponCardWeaponType(weaponType enum.WeaponType) WeaponCardOptions {
	return func(option *WeaponCardImpl) {
		option.weaponType = weaponType
	}
}

func NewWeaponCardWithOpts(options ...WeaponCardOptions) definition.WeaponCard {
	impl := &WeaponCardImpl{}
	cardTypeOption := WithCardType(enum.CardEquipment)
	cardTypeOption(&impl.CardImpl)
	equipmentTypeOption := WithEquipmentCardEquipmentType(enum.EquipmentWeapon)
	equipmentTypeOption(&impl.EquipmentCardImpl)

	for _, option := range options {
		option(impl)
	}

	return impl
}
