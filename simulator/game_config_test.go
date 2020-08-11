package simulator

import (
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
	"github.com/stretchr/testify/assert"
)

func sampleGameConfig() *GameConfig {
	song := models.NewDefaultSong("Default", 26, enum.AttrCool, 120000, 200)
	return NewGameConfig(
		&usermodel.Team{
			Ocards: [5]*usermodel.OwnedCard{
				&usermodel.OwnedCard{Vocal: 6954, Dance: 3779, Visual: 4564, Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrPassion}}},
				&usermodel.OwnedCard{Vocal: 3904, Dance: 4697, Visual: 7100, Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrPassion}}},
				&usermodel.OwnedCard{
					Vocal: 6970, Dance: 4602, Visual: 3821,
					Card: &models.Card{
						LeadSkill: &models.LeadSkillPassionVoice,
						Idol: &models.Idol{
							Attribute: enum.AttrPassion,
						},
						Rarity: &models.Rarity{Rarity: enum.RaritySSR},
					},
				},
				&usermodel.OwnedCard{Vocal: 7381, Dance: 3725, Visual: 4596, Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrPassion}}},
				&usermodel.OwnedCard{Vocal: 4694, Dance: 7106, Visual: 3897, Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrPassion}}},
			},
			LeaderIndex: 2,
		},
		[]*usermodel.OwnedCard{
			&usermodel.OwnedCard{Vocal: 6248, Dance: 4707, Visual: 4966, Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCool}}},
			&usermodel.OwnedCard{Vocal: 3744, Dance: 7356, Visual: 4608, Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCool}}},
			&usermodel.OwnedCard{Vocal: 4706, Dance: 7081, Visual: 3918, Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCool}}},
			&usermodel.OwnedCard{Vocal: 7381, Dance: 3725, Visual: 4596, Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCool}}},
			&usermodel.OwnedCard{Vocal: 3725, Dance: 4596, Visual: 7381, Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCool}}},
			&usermodel.OwnedCard{Vocal: 7100, Dance: 4697, Visual: 3904, Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCool}}},
			&usermodel.OwnedCard{Vocal: 4680, Dance: 7132, Visual: 3878, Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCool}}},
			&usermodel.OwnedCard{Vocal: 7179, Dance: 3855, Visual: 4642, Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCool}}},
			&usermodel.OwnedCard{Vocal: 7182, Dance: 3837, Visual: 4655, Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCool}}},
			&usermodel.OwnedCard{Vocal: 7204, Dance: 4646, Visual: 3817, Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCool}}},
		},
		&usermodel.OwnedCard{
			Vocal: 2171, Dance: 12141, Visual: 2468,
			Card: &models.Card{
				Idol:      &models.Idol{Attribute: enum.AttrPassion},
				LeadSkill: &models.LeadSkillBase,
				Rarity:    &models.Rarity{Rarity: enum.RaritySSR},
			},
		},
		&song,
	)
}

func TestAppeal(t *testing.T) {
	gc := sampleGameConfig()
	assert.Equal(t, 235059, gc.Appeal, "Error when calculating appeal!")
}

func TestRecalculateGameConfig(t *testing.T) {
	gc := sampleGameConfig()
	gc.SetGuest(&usermodel.OwnedCard{
		Vocal: 7063, Dance: 3839, Visual: 4637,
		Card: &models.Card{
			Idol:      &models.Idol{Attribute: enum.AttrPassion},
			LeadSkill: &models.LeadSkillPassionVoice,
			Rarity:    &models.Rarity{Rarity: enum.RaritySSR},
		},
	})
	assert.Equal(t, 271367, gc.Appeal, "Appeal is not recalculated when setting guest!")
	gc.SetSong(&models.Song{Attribute: enum.AttrAll})
	assert.NotEqual(t, 271367, gc.Appeal, "Appeal is not recalculated when setting song!")
}
