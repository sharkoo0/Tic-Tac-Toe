//TIC-TAC-TOE project

package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"

	"strconv"
	"strings"

	"os"

	"../pkg/game"
	"../pkg/algorithms"
)

//Test -> Helper function for player's turn and presenting it in the window
func Test(w window.Window) string {
	root1, _ := w.GetRootElement()
	r, _ := root1.SelectById("#10")
	text, _ := r.Text()
	w.DefineFunction("ai", AIPlays)
	return text
}

//Global variables used for GUI
var board game.Board
var win window.Window

func main() {

	board.Make()

	rect := sciter.NewRect(250,250,250,250)

	w, err := window.New(sciter.SW_TITLEBAR |
							sciter.SW_CONTROLS |
							sciter.SW_MAIN | 
							sciter.SW_ENABLE_DEBUG, rect)
	if err != nil {
		log.Fatal(err)
	}
	win = window.Window(*w)
	fullpath, err := filepath.Abs("../style/main_html.html")
	if err != nil {
		log.Fatal(err)
	}
	w.LoadFile(fullpath)
	w.SetTitle("TIC-TAC-TOE")

	

	saveWin := win

	root, _ := win.GetRootElement()
	buttonAI, _ := root.SelectById("AI")
	buttonAI.OnClick(func() {
		rect := sciter.NewRect(300,300,300,300)
		childWin, _ := window.New(sciter.SW_POPUP | 
									sciter.SW_TITLEBAR |
									sciter.SW_CONTROLS |
									sciter.SW_MAIN | 
									sciter.SW_ENABLE_DEBUG, rect)
		childWin.SetTitle("VS AI")
		childWin.LoadFile("../style/UI_for_AI_play.html")
		win = *childWin
		text := Test(win)
		if text == "stop" {
			os.Exit(0)
		}
		childWin.Show()
		childWin.Run()
	})

	buttonP, _ := root.SelectById("player")
	buttonP.OnClick(func() {
		rect := sciter.NewRect(300,300,300,300)
		childWin, _ := window.New(sciter.SW_POPUP | 
									sciter.SW_TITLEBAR |
									sciter.SW_CONTROLS |
									sciter.SW_MAIN |
									sciter.SW_ENABLE_DEBUG, rect)
		childWin.SetTitle("VS Player")
		childWin.LoadFile("../style/UI_for_PvP.html")
		win = *childWin
		text := Test(win)
		if text == "stop" {
			os.Exit(0);
		}
		childWin.Show()
		childWin.Run()
	})

	win = saveWin

	w.Show()
	w.Run()
}

//AIPlays -> The main AI function for the GUI 
//(The signature is this and only this (by default) if you want to use the function is the .tis file)
func AIPlays(vals ...*sciter.Value) *sciter.Value {
	if !board.IsFull() {
		os.Exit(0)
	}

	winner := -1
	if board.Evaluate() != 0 {
		winner = 10
		fmt.Println("UNFORTUNATELY, YOU LOSE!!! THE COMPUTER WINS!!!")
		os.Exit(0)
		return sciter.NewValue(winner)
	}

	//Player's turn
	text := Test(win)
	ij := strings.Split(text, " ")
	returnValue1, _ := strconv.Atoi(ij[0])
	returnValue2, _ := strconv.Atoi(ij[1])
	board.SetPossition(game.PLAYER, returnValue1, returnValue2)
	//end of player turn

	if !board.IsFull() {
		os.Exit(0)
	}
	
	//AI's turn
	move := algorithms.BestMove(board, false)
	if board.Evaluate() != 0 {
		winner = 0
		fmt.Println("CONGRANTS, YOU WIN!!!")
		os.Exit(0)
		return sciter.NewValue(winner)
	}
	
	board.SetPossition(game.OPPONENT, move.Row, move.Col)
	//end of AI turn

	if winner != -1 {
		fmt.Println("winner = ", winner)
		return sciter.NewValue(winner)
	}

	return sciter.NewValue(move.Row * 3 + move.Col + 1)
}
