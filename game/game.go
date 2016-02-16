package game

import (
	"../grid"
	"fmt"
	"sync"
	"time"
)

type Game struct {
	master, slave *grid.Grid
}

var(
	game = &Game{}
)

func NewGame(path string, generations int) {
	var err error
	game.master, err = grid.NewGrid(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	game.slave = game.master.Clone()
	fmt.Printf(game.slave.String())

	// create a WaitGroup that will act as a barrier for the goroutines
	wg := &sync.WaitGroup{}

	// create a buffer that will hold the coordinates to be processed
	coordChannel := make(chan grid.Pos, game.master.Height * game.master.Width)

	// create as many goroutines as the Grid height
	for id := range make([]int, game.master.Height){
		go updateStateWorker(id, coordChannel, wg)
	}


	for range make([]int, generations) {

		// Add the number of coordinates in the WaitGroup
		wg.Add(game.master.Height * game.master.Width)

		// Add all the coordinates to the channel
		channelAddCoordinates(coordChannel, game.master)

		// Wait until all the coordinates have been updated
		wg.Wait()

		// Swap slave and master, so that we may work on the updated data
		game.master = game.slave

		// Print black lines
		fmt.Print("\n\n\n\n\n")

		// Print the new grid
		fmt.Println(game.master.String())

		// Sleep for 500 ms
		time.Sleep(350 * time.Millisecond)

	}

	close(coordChannel)

}

func channelAddCoordinates(coordChannel chan grid.Pos, src *grid.Grid) {
	pos := grid.Pos{}
	for lineIdx := range make([]int, src.Height) {
		pos.Row = lineIdx
		for columnIdx := range make([]int, src.Width) {
			pos.Column = columnIdx
			coordChannel <- pos
		}
	}
}

func updateStateWorker(id int, ci chan grid.Pos, wg *sync.WaitGroup) {

	// Keep getting new sets of coordinates
	for coord := range ci{

		// Update the element at the provided coordinates
		game.slave.SetTile(coord.Row, coord.Column, game.master.NextState(coord.Row, coord.Column))

		// Decrement the WaitGroup counter
		wg.Done()
	}

}
