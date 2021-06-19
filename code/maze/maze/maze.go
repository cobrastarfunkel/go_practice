package maze

import (
	"bufio"
	"fmt"
	"maze/passage"
	"maze/room"
	"os"
	"strconv"
	"strings"
)

type Maze struct {
	Rooms      [][]room.Room
	rows, cols int
	MazeFile   string
}

func (m *Maze) setMazeSize(row []string) {
	rows, err := strconv.Atoi(row[0])

	if err != nil {
		panic(err)
	}

	cols, err := strconv.Atoi((row[1]))

	if err != nil {
		panic(err)
	}

	m.rows, m.cols = rows, cols
}

func (m *Maze) initRooms() {
	m.Rooms = make([][]room.Room, m.rows)
	for i := range m.Rooms {
		m.Rooms[i] = make([]room.Room, m.cols)
	}
}

func createPassage(id string) *passage.Passage {
	tempPass := passage.Passage{}
	switch id {
	case "+":
		tempPass.IsOpen = true
	case "-":
		tempPass.IsOpen = false
	default:
		tempPass.IsOpen = false
		tempPass.Key = id
	}
	return &tempPass
}

func setupPassages(curRoom *room.Room, lines []string) {
	curRoom.NorthPassage = createPassage(lines[0])
	curRoom.EastPassage = createPassage(lines[1])
	curRoom.SouthPassage = createPassage(lines[2])
	curRoom.WestPassage = createPassage(lines[3])
}

func getItems(item string, curRoom *room.Room) {
	for _, s := range strings.Fields(item) {
		curRoom.AddItem(s)
	}
}

func (m *Maze) makeRooms(lines []string) {
	curLine := 0
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			spl_line := strings.Split(lines[curLine], ",")
			m.Rooms[i][j] = room.Room{Name: spl_line[0]}

			getItems(spl_line[len(spl_line)-1], &m.Rooms[i][j])
			setupPassages(&m.Rooms[i][j], spl_line[1:len(spl_line)-1])

			curLine++
		}
	}
}

func (m *Maze) buildMaze(lines []string) {
	m.initRooms()
	m.makeRooms(lines)
}

func (m *Maze) LoadMaze() {
	file, err := os.Open(m.MazeFile)

	if err != nil {
		fmt.Printf("File Open error %s", err)
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	line := 0

	for scanner.Scan() {

		if line == 0 {
			row := strings.Fields(scanner.Text())
			m.setMazeSize(row)
		} else {
			lines = append(lines, scanner.Text())
		}

		line++
	}
	m.buildMaze(lines)
}
