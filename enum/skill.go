package enum

// SkillType 技能类型
type SkillType byte

const (
	SkillPassive        SkillType = iota // SkillPassive 被动技能
	SkillNormalAttack                    // SkillNormalAttack 普通攻击
	SkillElementalSkill                  // SkillElementalSkill 元素战技
	SkillElementalBurst                  // SkillElementalBurst 元素爆发
	SkillCooperative                     // SkillCooperative 协同攻击技能
)
