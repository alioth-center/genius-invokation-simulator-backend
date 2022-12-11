package definition

// Affiliation 角色的归属势力
type Affiliation uint

const (
	AffiliationMondstadt Affiliation = iota // AscriptionMondstadt 蒙德人
	AffiliationLiyue                        // AscriptionLiyue 璃月人
	AffiliationInazuma                      // AscriptionInazuma 稻妻人
	AffiliationSumeru                       // AscriptionSumeru 须弥人
)
