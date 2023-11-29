package main

// Welcome to
// __________         __    __  .__                               __
// \______   \_____ _/  |__/  |_|  |   ____   ______ ____ _____  |  | __ ____
//  |    |  _/\__  \\   __\   __\  | _/ __ \ /  ___//    \\__  \ |  |/ // __ \
//  |    |   \ / __ \|  |  |  | |  |_\  ___/ \___ \|   |  \/ __ \|    <\  ___/
//  |________/(______/__|  |__| |____/\_____>______>___|__(______/__|__\\_____>
//
// This file can be a nice home for your Battlesnake logic and helper functions.
//
// To get you started we've included code to prevent your Battlesnake from moving backwards.
// For more info see docs.battlesnake.com

import (
	"github.com/BattlesnakeOfficial/starter-snake-go/modifiers"
	"log"
)

// info is called when you create your Battlesnake on play.battlesnake.com
// and controls your Battlesnake's appearance
// TIP: If you open your Battlesnake URL in a browser you should see this data
func info() BattlesnakeInfoResponse {
	log.Println("INFO")

	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "",        // TODO: Your Battlesnake username
		Color:      "#888888", // TODO: Choose color
		Head:       "default", // TODO: Choose head
		Tail:       "default", // TODO: Choose tail
	}
}

// start is called when your Battlesnake begins a game
func start(state GameState) {
	log.Println("GAME START")
}

// end is called when your Battlesnake finishes a game
func end(state GameState) {
	log.Printf("GAME OVER\n\n")
}

// move is called on every turn and returns your next move
// Valid moves are "up", "down", "left", or "right"
// See https://docs.battlesnake.com/api/example-move for available data
func move(state GameState) BattlesnakeMoveResponse {

	movePredictions := [4]float64{
		100, 100, 100, 100,
	}

	// We've included code to prevent your Battlesnake from moving backwards
	myHead := state.You.Body[0] // Coordinates of your head
	myNeck := state.You.Body[1] // Coordinates of your "neck"
	// Left, Right, Up, Down
	if myNeck.X < myHead.X { // Neck is left of head, don't move left
		movePredictions[0] = 0

	} else if myNeck.X > myHead.X { // Neck is right of head, don't move right
		movePredictions[1] = 0

	} else if myNeck.Y < myHead.Y { // Neck is below head, don't move down
		movePredictions[3] = 0

	} else if myNeck.Y > myHead.Y { // Neck is above head, don't move up
		movePredictions[2] = 0
	}

	// TODO: Step 1 - Prevent your Battlesnake from moving out of bounds
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height

	var mov [4]float64
	mov[0] = float64(myHead.X)
	mov[1] = float64(boardWidth - 1 - myHead.X)
	mov[3] = float64(myHead.Y)
	mov[2] = float64(boardHeight - 1 - myHead.Y)

	for i, m := range mov {
		if i < 2 {
			m /= float64(boardWidth)
		} else {
			m /= float64(boardHeight)
		}
		m *= modifiers.PositionModifier
	}

	for mod := range movePredictions {
		movePredictions[mod] *= mov[mod]
	}

	//TODO: Step 2 - Prevent your Battlesnake from colliding with itself
	mybody := state.You.Body
	count := [4]float64{0, 0, 0, 0}

	for i, coords := range mybody {
		if i == 0 {
			continue
		}
		if coords.X < myHead.X {
			count[0]++
		} else if coords.X > myHead.X {
			count[1]++
		}
		if coords.Y < myHead.Y {
			count[3]++
		} else if coords.Y > myHead.Y {
			count[2]++
		}
	}

	for mod := range movePredictions {
		count[mod]++
		movePredictions[mod] /= count[mod] * modifiers.BodyModifier
	}
	//
	maximum := 0.0
	for _, prediction := range movePredictions {
		if prediction > maximum {
			maximum = prediction
		}
	}

	for p := range movePredictions {
		if maximum == movePredictions[p] {
			if p == 0 {
				log.Printf("MOVE %d: %s\n", state.Turn, "left")
				return BattlesnakeMoveResponse{Move: "left"}
			} else if p == 1 {
				log.Printf("MOVE %d: %s\n", state.Turn, "right")
				return BattlesnakeMoveResponse{Move: "right"}
			} else if p == 2 {
				log.Printf("MOVE %d: %s\n", state.Turn, "up")
				return BattlesnakeMoveResponse{Move: "up"}
			} else {
				log.Printf("MOVE %d: %s\n", state.Turn, "down")
				return BattlesnakeMoveResponse{Move: "down"}
			}
		}
	}
	
	log.Printf("MOVE %d: %s\n", state.Turn, "up")
	return BattlesnakeMoveResponse{Move: "up"}
}

func main() {
	RunServer()
}
