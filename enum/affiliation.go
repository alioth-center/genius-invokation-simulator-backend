package enum

// Affiliation 角色的归属势力
type Affiliation byte

const (
	AffiliationMondstadt Affiliation = iota        // AffiliationMondstadt 蒙德人
	AffiliationLiyue                               // AffiliationLiyue 璃月人
	AffiliationInazuma                             // AffiliationInazuma 稻妻人
	AffiliationSumeru                              // AffiliationSumeru 须弥人
	AffiliationFatui     Affiliation = 1<<4 + iota // AffiliationFatui 愚人众
	AffiliationHilichurl                           // AffiliationHilichurl 丘丘人
	AffiliationMonster                             // AffiliationMonster 魔物
	AffiliationUndefined Affiliation = 1<<8 - 1    // AffiliationUndefined 未定义
)
