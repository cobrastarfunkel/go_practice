package maze

import (
	"bufio"
	"fmt"
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

func (m *Maze) makeRooms(lines []string) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {

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
