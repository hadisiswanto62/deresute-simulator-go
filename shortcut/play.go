package shortcut

import (
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
	game := simulator.NewGame(gc)
	game.Play()
	// result := simulator.Simulate(gc, 1000)
	// fmt.Printf("%f %d\n", result.Average, result.Appeal)
}
