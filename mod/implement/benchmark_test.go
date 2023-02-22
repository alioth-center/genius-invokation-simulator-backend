package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"testing"
)

func BenchmarkTestNewCharacterWithOpts(b *testing.B) {
	SetDebugFlag(true)

	for i := 0; i < b.N; i++ {
		NewCharacterWithOpts(
			WithCharacterID(1),
			WithCharacterName("Ganyu"),
			WithCharacterAffiliation(enum.AffiliationLiyue),
			WithCharacterHP(10),
			WithCharacterMP(2),
			WithCharacterVision(enum.ElementCryo),
			WithCharacterWeapon(enum.WeaponBow),
		)
	}
}
