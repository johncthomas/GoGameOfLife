package main

import (
	"bufio"
	"fmt"
	"os"
)

// more info on interfaces
// http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go


type grid struct {
	// cells is the grid where the states will be kept
	cells 	[][]int
	xdim	int
	ydim	int
}

type indexer interface {
	xy()		int
	occupancy()	[4]int
	neighbours() int
	print()
} 

func (g *grid) xy(x, y int) *int {
	cell := &g.cells[x][y]
	return cell
}

func (g *grid) print() {
	for _, row := range g.cells {
		fmt.Println(row)
	}
}


// return the number of occupied neighbouring cells to x,  y
func (g *grid) neighbours(x, y int) int {
	offsetter := [3]int{-1, 0, 1}
	neighbours := 0
	for _, xoff := range offsetter {
		for _, yoff := range offsetter {
			if !(xoff == 0 && yoff == 0){
				nx := x+xoff
				ny := y+yoff
				if nx < g.xdim && ny < g.ydim && nx > 0 && ny > 0{
					neighbours += g.cells[x+xoff][y+yoff]
				}
			}
		}
	}
	return neighbours
}

func emptyGrid(xdim, ydim int) *grid {
	// calling this xs as ys don't really exist yet
	xs := new(grid)
	xs.xdim = xdim
	xs.ydim = ydim
	// add the y slices
	for xi := 0; xi < xdim; xi++ {
		y := make([]int, ydim)
		xs.cells = append(xs.cells, y)
	}
	
	return xs
}

func step(g grid) *grid {
	// new grid with updated cells, the returned (pointer) value
	nug := emptyGrid(g.xdim, g.ydim)
	// go through the cells and apply the rules of Conway's game of life
	for xi := 0; xi < g.xdim; xi++ {
		for yi := 0; yi < g.ydim; yi++ {
			adj := g.neighbours(xi, yi)
			// is the cell alive or dead?
			switch {
			case adj < 2 || adj > 3:
				nug.cells[xi][yi] = 0
			case adj == 3:
				nug.cells[xi][yi] = 1
			}
		}
	}
	return nug
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// create a grid out of arrangement of '0' & '1' in fn
// probaly rotates it 90deg. Non-rectangular arrangements
// are not caught, and will cause problems.
func readInGrid(fn string ) *grid {

	file, e := os.Open(fn)
	check(e)
	scanner := bufio.NewScanner(file)
	// set the split function
	scanner.Split(bufio.ScanLines)
	// calling scanner.Scan reads the file up to the split, 
	// ending the loop when done
	xs := [][]int{}
	rowi := 0
	var coli int
	for ; scanner.Scan(); rowi++ {
		ys := []int{}
		coli = 0
		for _, ch := range scanner.Text() {
			coli++
			switch ch {
			case '0':
				ys = append(ys, 0)
			case '1':
				ys = append(ys, 1)
			default:
				panic("Invalid char in grid file")
			}
		}
		xs = append(xs, ys)
	} // finished creating the sliceoslices
	g := emptyGrid(coli, rowi)
	g.cells = xs
	return g
}


func main() {
	// So `g` is a pointer so I_think that's why i can change the values
	//fn := os.Args[1]

	fn := os.Args[1]
	g := readInGrid(fn)
	g.print()

	for turn:=0;turn < 5; turn++ {
		g = step(*g)
		fmt.Println("*****************")
		g.print()
	}

}
