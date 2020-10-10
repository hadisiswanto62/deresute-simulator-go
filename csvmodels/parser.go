package csvmodels

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gocarina/gocsv"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type CSVDataParser struct {
}

const (
	ownedCardPath           = "userdata/cards.csv"
	songDifficultyIdPath    = "data/song_difficulty_id.csv"
	songDifficultyLevelPath = "data/song_difficulty_level.csv"
	songNamePath            = "data/song_name.csv"
	baseMusicScoresPath     = "data/song/"
)

func (p CSVDataParser) ParseOwnedCardRawData(path string) ([]*usermodel.OwnedCardRawData, error) {
	if path == "" {
		path = ownedCardPath
	}
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("cannot read %s: %v", ownedCardPath, err)
	}
	defer file.Close()

	tmps := []*TmpOwnedCardRawData{}
	if err := gocsv.UnmarshalFile(file, &tmps); err != nil {
		return nil, fmt.Errorf("cannot unmarshal %s: %v", ownedCardPath, err)
	}
	// var ret []*usermodel.OwnedCardRawData
	ret := make([]*usermodel.OwnedCardRawData, 0, len(tmps))
	for _, tocd := range tmps {
		ocd := &usermodel.OwnedCardRawData{
			CardID:     tocd.CardID,
			SkillLevel: tocd.SkillLevel,
			StarRank:   tocd.StarRank,
			PotVisual:  tocd.PotVisual,
			PotDance:   tocd.PotDance,
			PotVocal:   tocd.PotVocal,
			PotHp:      tocd.PotHp,
			PotSkill:   tocd.PotSkill,
		}
		ret = append(ret, ocd)
	}
	return ret, nil
}

func (p CSVDataParser) parseNotes(data string) ([]*TmpNoteRawData, error) {
	ret := []*TmpNoteRawData{}
	if err := gocsv.UnmarshalString(data, &ret); err != nil {
		return nil, fmt.Errorf("cannot parse: %v", err)
	}
	return ret, nil
}

var diffMap = map[int]enum.SongDifficulty{
	1: enum.SongDifficultyDebut,
	2: enum.SongDifficultyRegular,
	3: enum.SongDifficultyPro,
	4: enum.SongDifficultyMaster,
	5: enum.SongDifficultyMasterPlus,
}

func (p CSVDataParser) InitSongRawData() ([]*models.Song, error) {
	songDifficultyIds := []*TmpSongDifficultyId{}
	songDifficultyLevels := []*TmpSongDifficultyLevel{}
	musicNames := []*TmpMusicNameData{}
	// #region initializing
	// SongDifficultyID
	path := songDifficultyIdPath
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("cannot read %s: %v", path, err)
	}
	defer file.Close()
	if err = gocsv.UnmarshalFile(file, &songDifficultyIds); err != nil {
		return nil, fmt.Errorf("cannot unmarshal %s: %v", path, err)
	}

	// SongDifficultyLevel
	path = songDifficultyLevelPath
	file, err = os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("cannot read %s: %v", path, err)
	}
	defer file.Close()
	if err = gocsv.UnmarshalFile(file, &songDifficultyLevels); err != nil {
		return nil, fmt.Errorf("cannot unmarshal %s: %v", path, err)
	}

	// SongName
	path = songNamePath
	file, err = os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("cannot read %s: %v", path, err)
	}
	defer file.Close()
	if err = gocsv.UnmarshalFile(file, &musicNames); err != nil {
		return nil, fmt.Errorf("cannot unmarshal %s: %v", path, err)
	}
	// #endRegion

	type OneSong struct {
		ID        int
		Name      string
		Attribute enum.Attribute
	}
	oneSongs := []*OneSong{}
	for _, songDifficultyId := range songDifficultyIds {
		for _, musicName := range musicNames {
			if songDifficultyId.MusicDataID == musicName.MusicDataID {
				oneSongs = append(oneSongs, &OneSong{
					ID:        songDifficultyId.SongID,
					Name:      musicName.Name,
					Attribute: enum.GetAttribute(songDifficultyId.Type),
				})
			}
		}
	}
	songDiffs := make(map[*OneSong]map[enum.SongDifficulty]int, 0)
	for _, oneSong := range oneSongs {
		songDiffs[oneSong] = make(map[enum.SongDifficulty]int, 0)
		for _, songDifficultyLevel := range songDifficultyLevels {
			if oneSong.ID == songDifficultyLevel.SongID {
				diff, err := enum.GetSongDifficulty(songDifficultyLevel.DifficultyInt)
				if err != nil {
					continue
				}
				songDiffs[oneSong][*diff] = songDifficultyLevel.Level
			}
		}
	}

	songList := []*models.Song{}
	r, _ := regexp.Compile("^musicscores/m([0-9]*)/([0-9]*)_([0-9]*).csv$")
	for song, difficultyLevelMap := range songDiffs {
		songPath := baseMusicScoresPath + getSongPath(song.ID)
		baseDiffData := make(map[int][]*TmpNoteRawData, 0)

		// #region sql shit
		db, err := sql.Open("sqlite3", songPath)
		if err != nil {
			return nil, fmt.Errorf("cannot parse %s: %v", songPath, err)
		}
		defer db.Close()

		rows, err := db.Query(`SELECT * FROM blobs;`)
		if err != nil {
			fmt.Printf("cannot run query on %s\n", songPath)
			continue
		}
		defer rows.Close()
		for rows.Next() {
			var name string
			var data []byte
			err = rows.Scan(&name, &data)
			if err != nil {
				return nil, fmt.Errorf("invalid data: %v", err)
			}
			diffId := -1
			if r.MatchString(name) {
				diffIdStr := r.FindAllStringSubmatch(name, -1)[0][3]
				var err error
				diffId, err = strconv.Atoi(diffIdStr)
				if err != nil {
					return nil, fmt.Errorf("cannot parse filename: %v", err)
				}
			}
			if diffId == -1 {
				continue
			}
			baseDiffData[diffId], err = p.parseNotes(string(data))
			if err != nil {
				return nil, fmt.Errorf("cannot parse %s: %v", name, err)
			}
		}
		// #endregion

		for baseDiff, rawNotesData := range baseDiffData {
			notesData := []models.Note{}
			for _, rawNote := range rawNotesData {
				note := rawNote.toNote()
				if note == nil {
					continue
				}
				notesData = append(notesData, *rawNote.toNote())
			}
			difficulty, err := enum.GetSongDifficulty(baseDiff)
			if err != nil {
				// fmt.Printf("Difficulty %d not handled. %v\n", baseDiff, err)
				continue
			}
			level, ok := difficultyLevelMap[*difficulty]
			if !ok {
				fmt.Printf("Difficulty %d not found in map.\n", level)
			}
			chart := models.Song{
				Name:       song.Name,
				Level:      level,
				Difficulty: *difficulty,
				Attribute:  song.Attribute,
				DurationMs: notesData[len(notesData)-1].TimestampMs + 1000,
				Notes:      notesData,
			}
			songList = append(songList, &chart)
		}
	}

	return songList, nil
}

func getSongPath(id int) string {
	idStr := strconv.Itoa(id)
	for len(idStr) < 3 {
		idStr = "0" + idStr
	}
	return fmt.Sprintf("musicscores_m%s.db", idStr)
}
