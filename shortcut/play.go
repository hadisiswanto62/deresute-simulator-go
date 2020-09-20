package shortcut

import (
	"fmt"
	"time"

	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/simulator"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodelmanager"
)

// shortcut to play the game given config and customCardParams
func Play(config BaseGameConfig, customCardParams *usermodelmanager.CustomOwnedCardParameters, useDefaultCards bool) {
	gc, err := ToGameConfig(config, customCardParams, useDefaultCards)
	if err != nil {
		panic(err)
	}

	defer helper.MeasureTime(time.Now(), "Play")
	// game := simulator.NewGame(gc)
	// result := game.Play()
	result := simulator.Simulate(gc, 1000)
	fmt.Printf("%f %d %d\n", result.Average, result.Min, result.Max)
	// fmt.Printf("%d\n", result.Score)
}

// Shortcut to create GameConfig
func GameConfig(config BaseGameConfig,
	customCardParams *usermodelmanager.CustomOwnedCardParameters, useDefaultCards bool) *simulator.GameConfig {
	gc, err := ToGameConfig(config, customCardParams, useDefaultCards)
	if err != nil {
		panic(err)
	}
	return gc
	// result := simulator.Simulate(gc, 1000)
	// fmt.Printf("%f %d\n", result.Average, result.Appeal)
}
