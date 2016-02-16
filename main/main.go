package main
import(
	"flag"
	"../game"
)

// Command-line flags.
var (
	fname   	= flag.String("file", "pi_orbital.cells", "Initial state filename")
	generations = flag.Int("generations", 100, "Number of generations to emulate")
)

func main() {
	flag.Parse()

	if len(*fname) > 0 {
		game.NewGame("presets/" + *fname, *generations)
	} else {
		println("File parameter not aquired")
	}




}

