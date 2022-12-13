/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "reaction.go" LastUpdatedAt 2022/12/13 09:31:13
 */

package standard

import (
	d "github.com/sunist-c/genius-invokation-simulator-backend/definition"
)

var (
	// reactionDictionary 元素反应表
	reactionDictionary = map[reactivePair]d.Reaction{
		reactivePair{attached: d.ElementPyro, new: d.ElementCryo}:      d.ReactionMelt,              // 火冰融化
		reactivePair{attached: d.ElementCryo, new: d.ElementPyro}:      d.ReactionMelt,              // 冰火融化
		reactivePair{attached: d.ElementHydro, new: d.ElementPyro}:     d.ReactionVaporize,          // 火水蒸发
		reactivePair{attached: d.ElementPyro, new: d.ElementHydro}:     d.ReactionVaporize,          // 水火蒸发
		reactivePair{attached: d.ElementPyro, new: d.ElementElectro}:   d.ReactionOverloaded,        // 火雷超载
		reactivePair{attached: d.ElementElectro, new: d.ElementPyro}:   d.ReactionOverloaded,        // 雷火超载
		reactivePair{attached: d.ElementCryo, new: d.ElementElectro}:   d.ReactionSuperconduct,      // 冰雷超导
		reactivePair{attached: d.ElementElectro, new: d.ElementCryo}:   d.ReactionSuperconduct,      // 雷冰超导
		reactivePair{attached: d.ElementHydro, new: d.ElementCryo}:     d.ReactionFrozen,            // 冰水冻结
		reactivePair{attached: d.ElementCryo, new: d.ElementHydro}:     d.ReactionFrozen,            // 水冰冻结
		reactivePair{attached: d.ElementHydro, new: d.ElementElectro}:  d.ReactionElectroCharged,    // 水雷感电
		reactivePair{attached: d.ElementElectro, new: d.ElementHydro}:  d.ReactionElectroCharged,    // 雷水感电
		reactivePair{attached: d.ElementPyro, new: d.ElementDendro}:    d.ReactionBurning,           // 火草燃烧
		reactivePair{attached: d.ElementDendro, new: d.ElementPyro}:    d.ReactionBurning,           // 草火燃烧
		reactivePair{attached: d.ElementHydro, new: d.ElementDendro}:   d.ReactionBloom,             // 水草绽放
		reactivePair{attached: d.ElementDendro, new: d.ElementHydro}:   d.ReactionBloom,             // 草水绽放
		reactivePair{attached: d.ElementElectro, new: d.ElementDendro}: d.ReactionQuicken,           // 雷草激化
		reactivePair{attached: d.ElementDendro, new: d.ElementElectro}: d.ReactionQuicken,           // 草雷激化
		reactivePair{attached: d.ElementCryo, new: d.ElementAnemo}:     d.ReactionCryoSwirl,         // 冰风扩散
		reactivePair{attached: d.ElementHydro, new: d.ElementAnemo}:    d.ReactionHydroSwirl,        // 水风扩散
		reactivePair{attached: d.ElementPyro, new: d.ElementAnemo}:     d.ReactionPyroSwirl,         // 火风扩散
		reactivePair{attached: d.ElementElectro, new: d.ElementAnemo}:  d.ReactionElectroSwirl,      // 雷风扩散
		reactivePair{attached: d.ElementCryo, new: d.ElementGeo}:       d.ReactionCryoCrystalize,    // 冰岩结晶
		reactivePair{attached: d.ElementElectro, new: d.ElementGeo}:    d.ReactionElectroCrystalize, // 雷岩结晶
		reactivePair{attached: d.ElementHydro, new: d.ElementGeo}:      d.ReactionHydroCrystalize,   // 水岩结晶
		reactivePair{attached: d.ElementPyro, new: d.ElementGeo}:       d.ReactionPyroCrystalize,    // 火岩结晶
	}

	// elementAttachable 元素附着表
	elementAttachable = map[d.Element]bool{
		d.ElementAnemo:   false, // 风元素不可附着
		d.ElementCryo:    true,  // 冰元素可附着
		d.ElementDendro:  true,  // 草元素可附着
		d.ElementElectro: true,  // 雷元素可附着
		d.ElementGeo:     false, // 岩元素不可附着
		d.ElementHydro:   true,  // 水元素可附着
		d.ElementPyro:    true,  // 火元素可附着
	}
)

type reactivePair struct {
	attached d.Element
	new      d.Element
}

func elementAttach(elementNew d.Element, elementAttached []d.Element) (elementSurplus []d.Element) {
	if attachable := elementAttachable[elementNew]; attachable {
		for _, element := range elementAttached {
			if elementNew == element {
				return elementAttached
			}
		}

		return append(elementAttached, elementNew)
	} else {
		return elementAttached
	}
}

type ReactionCalculatorImplement struct{}

func (r ReactionCalculatorImplement) Type() d.RuleType {
	return d.RuleInGameModifier
}

func (r ReactionCalculatorImplement) Calculate(elementNew d.Element, elementAttached []d.Element) (reaction d.Reaction, elementSurplus []d.Element) {
	if len(elementAttached) == 0 {
		return d.ReactionNone, elementAttach(elementNew, elementAttached)
	} else if len(elementAttached) == 1 {
		pair := reactivePair{
			attached: elementAttached[0],
			new:      elementNew,
		}
		if reaction, reactive := reactionDictionary[pair]; reactive {
			return reaction, []d.Element{}
		} else {
			return d.ReactionNone, elementAttach(elementNew, elementAttached)
		}
	} else {
		for i, element := range elementAttached {
			pair := reactivePair{
				attached: element,
				new:      elementNew,
			}
			if reaction, reactive := reactionDictionary[pair]; reactive {
				elementSurplus = append(elementAttached[:i], elementAttached[i+1:]...)
				return reaction, elementSurplus
			}
		}

		return d.ReactionNone, elementAttach(elementNew, elementAttached)
	}
}
