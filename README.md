# Conway's game of life

golang cli implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) 

### requirements
[golang](https://golang.org/doc/install)

### usage
`go run . [-file="centered_glider.txt"] [-w=16] [-h=16] [-gens=24]`
* file - input text file with initial state of playground
* w - width in characters of playground 
* h - height in lines of playground
* gens - ticks/generations to run

### limitations
* character to draw playground in CLI is `.` and `█`
* at least correct width is required to be specified
* playground file need to have equal length of lines (playground is a square)
* I didn't find quick solution to clean the CLI output between the ticks (this program was written on Windows)

### examples on how to run
* `go run .`
* `go run . -file="5x5.txt" -w=5 -h=5 -gens=4`
* `go run . -file="centered_glider.txt" -w=16 -h=16 -gens=24` (default params)

### what is not covered
As this is a coding task rather then production code, follow things are not provided:
* tests (unit, integration)
* some corner cases
* graceful shutdown
* linter rules (code is linted)

### benchmark
`go test -bench=. -benchmem` command will run single glider demo with 100 ticks