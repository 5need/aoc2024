package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Coordinates struct {
	x int
	y int
}

type Node struct {
	pos  Coordinates
	kind rune
}

func main() {
	r := strings.Split(input, "\n")
	fmt.Println(input)

	nodes := []Node{}

	for y, row := range r {
		for x, char := range row {
			if unicode.IsLetter(char) || unicode.IsNumber(char) {
				nodes = append(nodes, Node{pos: Coordinates{x, y}, kind: char})
			}
		}
	}

	antinodes := make(map[Coordinates]bool)

	height := len(r)
	width := len(r[0])

	for _, node := range nodes {
		for _, otherNode := range nodes {
			if node == otherNode {
				continue
			}
			if node.kind != otherNode.kind {
				continue
			}
			// node and otherNode are not the same and also share the same kind

			dx := otherNode.pos.x - node.pos.x
			dy := otherNode.pos.y - node.pos.y
			gcd := gcd(dx, dy)

			xStep := dx / gcd
			yStep := dy / gcd

			i := 0
			antinodePos := Coordinates{node.pos.x + i*xStep, node.pos.y + i*yStep}

			// forward
			for antinodePos.x >= 0 && antinodePos.x < width && antinodePos.y >= 0 && antinodePos.y < height {
				antinodes[antinodePos] = true

				i++
				antinodePos = Coordinates{node.pos.x + i*xStep, node.pos.y + i*yStep}
			}

			i = 0
			antinodePos = Coordinates{node.pos.x + -i*xStep, node.pos.y + -i*yStep}

			// backwards
			for antinodePos.x >= 0 && antinodePos.x < width && antinodePos.y >= 0 && antinodePos.y < height {
				antinodes[antinodePos] = true

				i++
				antinodePos = Coordinates{node.pos.x + -i*xStep, node.pos.y + -i*yStep}
			}
		}
	}

	fmt.Println("")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if antinodes[Coordinates{x, y}] == true {
				fmt.Print("x")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
	fmt.Println(len(antinodes))
}

func gcd(a, b int) int {
	// euclidean algo
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

var input = `.....y..........................p................r
........I.........................................
......................4.s.........................
..........4.......................................
....y.............................................
......................................p.........r.
..........0..s......N..................1.....p....
..y........4.......................p..............
...............0..................................
..............0....t....N....h....................
.............N....................................
......j...................s............H...l..O...
..........q.................H................O....
..f...e.qj.....y...0..............................
...........t..........................k..Q..r.....
.........6................Q..s...x......W.........
....2..b...e....t..4.........c.....xW.j...........
...e....................w................1.....O..
..e..j..5...........................c.............
.........B..2...............MK................H...
...2......b...g..X...q..........h...............O.
...q...2..........m....k...i...............QV.x...
...................i.........W.k.............HQ...
........b...X...............D..........c...N......
................................l..........h.....I
.m...........g......l.......c.............3......V
....X.......m........g...V.K...7......F.d.........
.........b.X...U..........................C.......
.....................l..............o.1....C......
............u.............K..............3...d....
......................i.T....f................V...
..............................1.k.................
.B.....E......9..m....K..5.M......................
...P...............M...95....o..i........I........
...............................S......3......wI...
.....EP...........9........5..T.R.................
.P..........v..9......f.............R.Co..w3......
..........h...SG..v.E...7..f.T....................
..........6..........L.................Y.......d..
..........B...............U........D..............
....B................U.....8..M....n...J..........
.........................L................Fw......
....L6E.P.................7.UG....J.....Y.D.......
........t........v...SJ........n..d...............
......................8v.....uG...................
..................L.....n.........................
...............T..............n......D............
..............o.........8................J.Y.R....
..................S...............u....F.......R..
........6..............u.....7.8..........Y..F....`
var input2 = `............
............
............
............
....0...0...
............
............
....0.......
....0.......
............
............
............`
