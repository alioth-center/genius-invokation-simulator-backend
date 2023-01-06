package persistence

var (
	RuleSetPersistence   Persistence[RuleSet]
	CardPersistence      Persistence[Card]
	CharacterPersistence Persistence[CharacterInfo]
	PlayerPersistence    Persistence[PlayerInfo]
	SkillPersistence     Persistence[Skill]
)

type Persistence[T any] interface {
	QueryByID(id uint) (has bool, result T)
}
