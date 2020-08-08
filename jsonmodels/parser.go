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
	if err := save(lst, "data/rarity.json"); err != nil {
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
	if err := save(lst3, "data/skill.json"); err != nil {
		return fmt.Errorf("could not save skills: %v", err)
	}

	var lst4 []TmpIdol
	for i := range tmpIdols {
		lst4 = append(lst4, i)
	}
	if err := save(lst4, "data/idol.json"); err != nil {
		return fmt.Errorf("could not save idols: %v", err)
	}

	save(cards, cardPath)
	return nil
}

// Parse parse card data into `[]models.Card`
func (p JSONDataParser) Parse() ([]models.Card, error) {
	var ret []models.Card

	// Parsing lead skills
	text, err := ioutil.ReadFile(leadSkillPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file (%s): %v", leadSkillPath, err)
	}
	var tmpLeadSkills []TmpLeadSkill
	if err = json.Unmarshal(text, &tmpLeadSkills); err != nil {
		return nil, fmt.Errorf("cannot unmarshal lead skill json: %v", err)
	}
	leadSkills, err := p.parseLeadSkills(tmpLeadSkills)
	if err != nil {
		return nil, err
	}
	fmt.Println(len(leadSkills))

	// Parsing skills
	// Parsing idols
	// Parsing rarities

	// Parsing cards
	var cards []TmpCard
	text, err = ioutil.ReadFile(cardPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %v", err)
	}
	if err = json.Unmarshal(text, &cards); err != nil {
		return nil, fmt.Errorf("cannot unmarshal %s : %v", cardPath, err)
	}

	for _, tmpCard := range cards {
		isEvolved := tmpCard.TmpRarity.IsEvolved()
		card := models.Card{
			ID:          tmpCard.ID,
			SeriesID:    tmpCard.SeriesID,
			Idol:        nil,
			Rarity:      nil,
			LeadSkill:   nil,
			IsEvolved:   isEvolved,
			MaxLevel:    tmpCard.TmpRarity.BaseMaxLevel,
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
		ret = append(ret, card)
	}

	return ret, nil
}

func (p JSONDataParser) parseLeadSkills(tmpLeadSkills []TmpLeadSkill) ([]*models.LeadSkill, error) {
	var leadSkills []*models.LeadSkill
	for _, tmpLeadSkill := range tmpLeadSkills {
		leadSkill, err := models.GetLeadSkill(tmpLeadSkill.Name)
		if err != nil {
			return nil, fmt.Errorf("cannot get lead skill for (%s): %v", tmpLeadSkill.Name, err)
		}
		leadSkills = append(leadSkills, leadSkill)
	}
	return leadSkills, nil
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
