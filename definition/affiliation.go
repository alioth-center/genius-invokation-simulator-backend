/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "affiliation.go" LastUpdatedAt 2022/12/12 10:19:12
 */

package definition

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
)
