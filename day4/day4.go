package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	numberCount int // All non-used numbers added together
	rows        [][]string
	winner      bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Print("Starting day 4")
	answer := Part2("input/data")
	fmt.Printf("Answer = %d\n", answer)

}

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstLine := true
	var numbersCalled []string
	rows := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		// First line is numbers called
		if firstLine {
			numbersCalled = strings.Split(line, ",")
			firstLine = false
			continue
		}

		if line == "" {
			continue
		}

		row := strings.Fields(line)
		rows = append(rows, row)
	}

	fmt.Print(numbersCalled)

	for i, row := range rows {
		fmt.Printf("Row %d = %q\n", i, row)
	}

	lastRow := len(rows)
	// Board1 = 0:5
	// Board2 = 5:10
	boards := make([]Board, 0)
	start := 0
	end := 5
	for {
		board := Board{rows: rows[start:end]}
		boards = append(boards, board)

		if end >= lastRow {
			break
		}

		start += 5
		end += 5
	}

	// Call numbers until someone wins or theres none left

	for _, num := range numbersCalled {
		fmt.Printf("CHECKING NUMBER %q\n", num)
		for _, board := range boards {
			updateBoard(board, num)

			if board.checkForBingo() {
				fmt.Printf("BINGO!")

				intNum, err := strconv.Atoi(num)
				check(err)
				return intNum * board.calculateScore()
			}
		}
	}

	fmt.Print("No bingo found")
	return -1
}

func Part2(filePath string) int {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstLine := true
	var numbersCalled []string
	rows := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		// First line is numbers called
		if firstLine {
			numbersCalled = strings.Split(line, ",")
			firstLine = false
			continue
		}

		if line == "" {
			continue
		}

		row := strings.Fields(line)
		rows = append(rows, row)
	}

	fmt.Print(numbersCalled)

	for i, row := range rows {
		fmt.Printf("Row %d = %q\n", i, row)
	}

	lastRow := len(rows)
	// Board1 = 0:5
	// Board2 = 5:10
	boards := make([]Board, 0)
	start := 0
	end := 5
	for {
		board := Board{rows: rows[start:end]}
		boards = append(boards, board)

		if end >= lastRow {
			break
		}

		start += 5
		end += 5
	}

	// Call numbers until someone wins or theres none left
	winningBoards := make([]Board, 0)
	for _, num := range numbersCalled {
		fmt.Printf("CHECKING NUMBER %q\n", num)

		for _, board := range boards {
			updateBoard(board, num)

			if board.checkForBingo() {
				fmt.Printf("BINGO!")

				fmt.Printf("BOARD.WINNER = %t\n", board.winner)

				if !board.inWinningBoards(winningBoards) {
					winningBoards = append(winningBoards, board)
				}
				// if board.winner == false {
				// 	fmt.Printf("Adding to winning boards!")
				// 	board.winner = true
				// 	fmt.Printf("Board.winner = %t", board.winner)
				// 	winningBoards = append(winningBoards, board)
				// 	fmt.Printf("Number of winning boards= %d", len(winningBoards))
				// } else {
				// 	fmt.Printf("That board already won. Skipping")
				// }

				if len(winningBoards) == len(boards) {

					fmt.Printf("\nThat was the last board!")
					printBoard(board)

					fmt.Printf("\nBOARD SUPPOSED TO LOST")
					printBoard(boards[1])
					intNum, err := strconv.Atoi(num)
					check(err)
					return intNum * board.calculateScore()
				}

				fmt.Print("\n")
				printBoard(board)
				fmt.Print("\n")

			}
		}
	}

	fmt.Print("No bingo found")
	return -1
}

func printBoard(board Board) {
	fmt.Print("\n")
	for _, row := range board.rows {
		fmt.Print(row)
		fmt.Print("\n")
	}
}

func (board Board) calculateScore() int {
	score := 0
	for _, row := range board.rows {
		for _, val := range row {
			if val != "X" {
				valInt, err := strconv.Atoi(val)
				check(err)
				score += valInt
			}
		}
	}
	return score
}

func (boardToCheck Board) inWinningBoards(winningBoards []Board) bool {

	for _, board := range winningBoards {
		if boardToCheck.equals(board) {
			return true
		}
	}
	return false
}

func (board Board) equals(otherBoard Board) bool {
	for i, row := range board.rows {
		for j, _ := range row {
			if board.rows[i][j] != otherBoard.rows[i][j] {
				return false
			}
		}
	}

	return true

}

// Looks for the number on the board and marks it as an X if it is there
func updateBoard(board Board, num string) {

	for i, row := range board.rows {
		for j, val := range row {
			if val == num {
				board.rows[i][j] = "X"
			}
		}
	}
}

func (board Board) checkForBingo() bool {
	// Horizontal and vertical Check
	for i, row := range board.rows {
		if completeRow(row) {
			return true
		}
		if completeRow(board.boardColumn(i)) {
			return true
		}
	}
	return false
}

func completeRow(row []string) bool {

	// fmt.Printf("\n Checking row %q", row)
	for _, val := range row {
		if val != "X" {
			return false
		}
	}
	return true
}

func (board Board) boardColumn(columnIndex int) []string {
	column := make([]string, 0)
	for _, row := range board.rows {
		column = append(column, row[columnIndex])
	}

	return column
}
