/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "skill.go" LastUpdatedAt 2022/12/12 10:20:12
 */

package definition

// SkillType 技能类型
type SkillType byte

const (
	SkillPassive        SkillType = iota // SkillPassive 被动技能
	SkillNormalAttack                    // SkillNormalAttack 普通攻击
	SkillElementalSkill                  // SkillElementalSkill 元素战技
	SkillElementalBurst                  // SkillElementalBurst 元素爆发
)
