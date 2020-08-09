package jsonmodels

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

const (
	rawAllCardsDir = "resources/all_cards.json"
	// AllCardsDir is directory where clean Cards data is stored
	allCardsDir   = "data/all_cards.json"
	allCardsDir2  = "data/all_cards2.json"
	leadSkillPath = "data/lead_skill.json"
	cardPath      = "data/cards.json"
	rarityPath    = "data/rarity.json"
	idolPath      = "data/idol.json"
	skillPath     = "data/skill.json"
)

// JSONDataParser manages card data that is stored in JSON form
type JSONDataParser struct {
}

var instance *JSONDataParser

// ParserInstance get instance of JSONDataParser
func ParserInstance() *JSONDataParser {
	if instance == nil {
		instance = &JSONDataParser{}
	}
	return instance
}

// InitData initializes data to be parsed
func (p JSONDataParser) InitData() error {
	text, err := ioutil.ReadFile(rawAllCardsDir)
	if err != nil {
		return fmt.Errorf("read file error: %v", err)
	}

	var cards []TmpCard
	if err := json.Unmarshal(text, &cards); err != nil {
		return fmt.Errorf("cannot parse json: %v", err)
	}

	tmpRarities := make(map[TmpRarity]bool)
	tmpLeadSkills := make(map[TmpLeadSkill]bool)
	tmpIdols := make(map[TmpIdol]bool)
	tmpSkills := make(map[TmpSkill]bool)
	for _, card := range cards {
		tmpRarities[card.TmpRarity] = true
		tmpLeadSkills[card.TmpLeadSkill] = true
		tmpIdols[card.TmpIdol] = true
		tmpSkills[card.TmpSkill] = true
	}

	var lst []TmpRarity
	for i := range tmpRarities {
		lst = append(lst, i)
	}
	if err := save(lst, rarityPath); err != nil {
		return fmt.Errorf("could not save rarities: %v", err)
	}

	var lst2 []TmpLeadSkill
	for i := range tmpLeadSkills {
		lst2 = append(lst2, i)
	}
	if err := save(lst2, leadSkillPath); err != nil {
		return fmt.Errorf("could not save lead skills: %v", err)
	}

	var lst3 []TmpSkill
	for i := range tmpSkills {
		lst3 = append(lst3, i)
	}
	if err := save(lst3, skillPath); err != nil {
		return fmt.Errorf("could not save skills: %v", err)
	}

	var lst4 []TmpIdol
	for i := range tmpIdols {
		lst4 = append(lst4, i)
	}
	if err := save(lst4, idolPath); err != nil {
		return fmt.Errorf("could not save idols: %v", err)
	}

	save(cards, cardPath)
	return nil
}

// Parse parse card data into `[]*models.Card`
func (p JSONDataParser) Parse() ([]*models.Card, error) {
	var ret []*models.Card

	// Parsing skills
	var tmpSkills []TmpSkill
	var skills []*models.Skill
	text, err := ioutil.ReadFile(skillPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file (%s): %v", skillPath, err)
	}
	if err = json.Unmarshal(text, &tmpSkills); err != nil {
		return nil, fmt.Errorf("cannot unmarshal (%s): %v", rarityPath, err)
	}
	for _, tmpSkill := range tmpSkills {
		skill, err := makeSkill(tmpSkill)
		if err != nil {
			return nil, fmt.Errorf("cannot make skill: %v", err)
		}
		skills = append(skills, skill)
	}

	// Parsing rarities
	var tmpRarities []TmpRarity
	var rarities []*models.Rarity
	text, err = ioutil.ReadFile(rarityPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file (%s): %v", rarityPath, err)
	}
	if err = json.Unmarshal(text, &tmpRarities); err != nil {
		return nil, fmt.Errorf("cannot unmarshal (%s): %v", rarityPath, err)
	}
	for _, tmpRarity := range tmpRarities {
		rarities = append(rarities, makeRarity(tmpRarity))
	}

	// Parsing idols
	var tmpIdols []TmpIdol
	var idols []*models.Idol
	text, err = ioutil.ReadFile(idolPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file (%s): %v", idolPath, err)
	}
	if err = json.Unmarshal(text, &tmpIdols); err != nil {
		return nil, fmt.Errorf("cannot unmarshal (%s): %v", rarityPath, err)
	}
	for _, tmpIdol := range tmpIdols {
		idols = append(idols, makeIdol(tmpIdol))
	}

	// Parsing cards
	var cards []TmpCard
	text, err = ioutil.ReadFile(cardPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file (%s): %v", cardPath, err)
	}
	if err = json.Unmarshal(text, &cards); err != nil {
		return nil, fmt.Errorf("cannot unmarshal %s : %v", cardPath, err)
	}

	for _, tmpCard := range cards {
		leadSkill, err := makeLeadSkill(tmpCard.TmpLeadSkill)
		if err != nil {
			return nil, fmt.Errorf("cannot get lead skill for card id (%d): %v", tmpCard.ID, err)
		}
		var chosenRarity *models.Rarity
		for _, rarity := range rarities {
			if rarity.ID == tmpCard.TmpRarity.Rarity {
				chosenRarity = rarity
			}
		}
		if chosenRarity == nil {
			return nil, fmt.Errorf("cannot get rarity for card id (%d)", tmpCard.ID)
		}

		var chosenIdol *models.Idol
		for _, idol := range idols {
			if idol.ID == tmpCard.TmpIdol.ID {
				chosenIdol = idol
				break
			}
		}
		if chosenIdol == nil {
			return nil, fmt.Errorf("cannot get idol for card id (%d)", tmpCard.ID)
		}

		var chosenSkill *models.Skill
		for _, skill := range skills {
			if skill.ID == tmpCard.TmpSkill.ID {
				chosenSkill = skill
				break
			}
		}
		if chosenSkill == nil {
			return nil, fmt.Errorf("cannot get skill for card id (%d)", tmpCard.ID)
		}

		card := models.Card{
			ID:          tmpCard.ID,
			SeriesID:    tmpCard.SeriesID,
			Idol:        chosenIdol,
			Rarity:      chosenRarity,
			LeadSkill:   leadSkill,
			Skill:       chosenSkill,
			BonusDance:  tmpCard.BonusDance,
			BonusHp:     tmpCard.BonusHp,
			BonusVisual: tmpCard.BonusVisual,
			BonusVocal:  tmpCard.BonusVocal,
			DanceMax:    tmpCard.DanceMax,
			DanceMin:    tmpCard.DanceMin,
			HpMax:       tmpCard.HpMax,
			HpMin:       tmpCard.HpMin,
			VisualMax:   tmpCard.VisualMax,
			VisualMin:   tmpCard.VisualMin,
			VocalMax:    tmpCard.VocalMax,
			VocalMin:    tmpCard.VocalMin,
		}
		ret = append(ret, &card)
	}

	return ret, nil
}

func save(obj interface{}, filename string) error {
	data, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		return fmt.Errorf("could not marshal: %v", err)
	}
	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("write failed: %v", err)
	}
	return nil
}
