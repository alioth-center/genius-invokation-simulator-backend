/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "damage.go" LastUpdatedAt 2022/12/14 09:23:14
 */

package model

import "github.com/sunist-c/genius-invokation-simulator-backend/definition"

type Damage struct {
	Element definition.Element
	Amount  uint
}

type AttackDamageContext struct {
	Sender ISkill
	target *Player
	damage map[*Character]Damage
}

func (d AttackDamageContext) Effective() bool {
	damageAmount := uint(0)
	for _, damage := range d.damage {
		damageAmount += damage.Amount
	}
	return damageAmount != 0
}

func (d *AttackDamageContext) AddActiveDamage(amount uint) {
	damage := d.damage[d.target.ActiveCharacter]
	damage.Amount += amount
	d.damage[d.target.ActiveCharacter] = damage
}

func (d *AttackDamageContext) AddPenetratedDamage(amount uint) {
	for character, damage := range d.damage {
		if character != d.target.ActiveCharacter {
			damage.Amount += amount
		}
		d.damage[character] = damage
	}
}

func (d *AttackDamageContext) SubActiveDamage(amount uint) {
	damage := d.damage[d.target.ActiveCharacter]
	if damage.Amount > amount {
		damage.Amount -= amount
	} else {
		damage.Amount = 0
	}
	d.damage[d.target.ActiveCharacter] = damage
}

func (d *AttackDamageContext) SubPenetratedDamage(amount uint) {
	for character, damage := range d.damage {
		if character != d.target.ActiveCharacter {
			if damage.Amount > amount {
				damage.Amount -= amount
			} else {
				damage.Amount = 0
			}
			d.damage[character] = damage
		}
	}
}

func (d *AttackDamageContext) ChangeActiveDamageElement(element definition.Element) {
	damage := d.damage[d.target.ActiveCharacter]
	damage.Element = element
	d.damage[d.target.ActiveCharacter] = damage
}

func (d *AttackDamageContext) ChangeBackgroundDamageElement(element definition.Element) {
	for character, damage := range d.damage {
		if character != d.target.ActiveCharacter {
			damage.Element = element
		}
		d.damage[character] = damage
	}
}

func (d AttackDamageContext) Damage() map[*Character]Damage {
	return d.damage
}

func NewAttackDamageContext(target *Player, sender ISkill, element definition.Element) *AttackDamageContext {
	d := &AttackDamageContext{
		Sender: sender,
		target: target,
		damage: map[*Character]Damage{},
	}

	d.damage[target.ActiveCharacter] = Damage{
		Element: element,
		Amount:  0,
	}

	for _, character := range target.Characters {
		if character != target.ActiveCharacter {
			d.damage[character] = Damage{
				Element: definition.ElementNone,
				Amount:  0,
			}
		}
	}

	return d
}

type AttackDamageModifier func(ctx *ModifierContext[AttackDamageContext])

type DefenceDamageContext struct {
	sender *Player
	self   *Player
	damage map[*Character]Damage
}

func (d *DefenceDamageContext) Effective() bool {
	damageAmount := uint(0)
	for _, damage := range d.damage {
		damageAmount += damage.Amount
	}
	return damageAmount != 0
}

func (d *DefenceDamageContext) AddActiveDamage(amount uint) {
	damage := d.damage[d.self.ActiveCharacter]
	damage.Amount += amount
	d.damage[d.self.ActiveCharacter] = damage
}

func (d *DefenceDamageContext) AddPenetratedDamage(amount uint) {
	for character, damage := range d.damage {
		if character != d.self.ActiveCharacter {
			damage.Amount += amount
		}
		d.damage[character] = damage
	}
}

func (d *DefenceDamageContext) SubActiveDamage(amount uint) {
	damage := d.damage[d.self.ActiveCharacter]
	if damage.Amount > amount {
		damage.Amount -= amount
	} else {
		damage.Amount = 0
	}
	d.damage[d.self.ActiveCharacter] = damage
}

func (d *DefenceDamageContext) SubPenetratedDamage(amount uint) {
	for character, damage := range d.damage {
		if character != d.self.ActiveCharacter {
			if damage.Amount > amount {
				damage.Amount -= amount
			} else {
				damage.Amount = 0
			}
			d.damage[character] = damage
		}
	}
}

func (d *DefenceDamageContext) ChangeActiveDamageElement(element definition.Element) {
	damage := d.damage[d.self.ActiveCharacter]
	damage.Element = element
	d.damage[d.self.ActiveCharacter] = damage
}

func (d *DefenceDamageContext) ChangeBackgroundDamageElement(element definition.Element) {
	for character, damage := range d.damage {
		if character != d.self.ActiveCharacter {
			damage.Element = element
		}
		d.damage[character] = damage
	}
}

func (d DefenceDamageContext) Damage() map[*Character]Damage {
	return d.damage
}

func NewDefenceDamageContext(from *Player, attack *AttackDamageContext) *DefenceDamageContext {
	return &DefenceDamageContext{
		sender: from,
		self:   attack.target,
		damage: attack.damage,
	}
}

type DefenceDamageModifier func(ctx *ModifierContext[DefenceDamageContext])
