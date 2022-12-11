package definition

// Reaction 反应类型
type Reaction uint

const (
	ReactionMelt           Reaction = iota // ReactionMelt 融化
	ReactionVaporize                       // ReactionVaporize 蒸发
	ReactionOverloaded                     // ReactionOverloaded 超载
	ReactionSuperconduct                   // ReactionSuperconduct 超导
	ReactionFrozen                         // ReactionFrozen 冻结
	ReactionElectroCharged                 // ReactionElectroCharged 感电
	ReactionBurning                        // ReactionBurning 燃烧
	ReactionBloom                          // ReactionBloom 绽放
	ReactionQuicken                        // ReactionQuicken 激化
	ReactionCryoSwirl                      // ReactionCryoSwirl 冰扩散
	ReactionHydroSwirl                     // ReactionHydroSwirl 水扩散
	ReactionPyroSwirl                      // ReactionPyroSwirl 火扩散
	ReactionElectroSwirl                   // ReactionElectroSwirl 电扩散
	// todo: Complete Geo Reaction
)
