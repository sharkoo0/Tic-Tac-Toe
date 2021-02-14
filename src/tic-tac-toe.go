//TIC-TAC-TOE project

package main

import (
	"fmt"
	"log"

	//"time"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"

	// "strconv"
	// "strings"

	"os"

	"github.com/sharkoo0/Tic-Tac-Toe/pkg/algorithms"
	"github.com/sharkoo0/Tic-Tac-Toe/pkg/game"
	"github.com/sharkoo0/Tic-Tac-Toe/pkg/tournament"
	//"github.com/sharkoo0/Tic-Tac-Toe/pkg/tournament"
)

//Test -> Helper function for player's turn and presenting it in the window
func Test(w *window.Window) string {
	root1, _ := w.GetRootElement()
	fmt.Println(root1)
	r, _ := root1.SelectById("#10")
	text, _ := r.Text()
	//w.DefineFunction("ai", AIPlays)
	return text
}

//Global variables used for GUI
var board game.Board
var win window.Window

func main() {

	t := tournament.Tournament{}
	t.Init()

	// time.Sleep(time.Second)

	board.Make()

	// rect := sciter.NewRect(250, 250, 250, 250)

	// win, err := window.New(sciter.SW_TITLEBAR|
	// 	sciter.SW_CONTROLS|
	// 	sciter.SW_MAIN|
	// 	sciter.SW_ENABLE_DEBUG, rect)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fullpath, err := filepath.Abs("./style/main_html.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// win.LoadFile(fullpath)
	// win.SetTitle("TIC-TAC-TOE")

	// // win = window.Window(*w)
	// //saveWin := win

	// root, e := win.GetRootElement()
	// if e != nil {
	// 	fmt.Println(e.Error())
	// 	log.Fatal()
	// }
	// buttonAI, err := root.SelectById("AI")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	log.Fatal()
	// }
	// buttonAI.OnClick(func() {
	// 	// rect := sciter.NewRect(300, 300, 300, 300)
	// 	// childWin, err := window.New(sciter.SW_POPUP|
	// 	// 	sciter.SW_TITLEBAR|
	// 	// 	sciter.SW_CONTROLS|
	// 	// 	sciter.SW_MAIN|
	// 	// 	sciter.SW_ENABLE_DEBUG, rect)
	// 	// if err != nil {
	// 	// 	fmt.Println(err.Error())
	// 	// 	log.Fatal()
	// 	// }

	// 	win.SetTitle("VS AI")
	// 	err = win.LoadFile("./style/UI_for_AI_play.html")
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		log.Fatal()
	// 	}

	// 	//win.DefineFunction("AIPlays", AIPlays)

	// 	//win = *childWin

	// 	// text := Test(win)
	// 	// if text == "stop" {
	// 	// 	os.Exit(0)
	// 	// }
	// 	//win.UpdateWindow()
	// 	// win.Show()
	// 	// win.Run()
	// 	// fmt.Println(*win.Sciter)
	// 	//win = window.Window(*childWin)
	// 	//log.Fatal()
	// 	//childWin.Show()
	// 	//childWin.Run()

	// 	// text := Test(&win)
	// 	// if text == "stop" {
	// 	// 	os.Exit(0)
	// 	// }
	// 	// fmt.Println(text)
	// })

	// buttonP, e := root.SelectById("player")
	// if e != nil {
	// 	fmt.Println(e.Error())
	// 	log.Fatal()
	// }
	// buttonP.OnClick(func() {
	// 	// win.SetTitle("VS PLAYER")
	// 	// err = win.LoadFile("./style/UI_for_PvP.html")
	// 	// if err != nil {
	// 	// 	fmt.Println(err.Error())
	// 	// 	log.Fatal()
	// 	// }
	// 	rect := sciter.NewRect(300, 300, 300, 300)
	// 	childWin, err := window.New(sciter.SW_POPUP|
	// 		sciter.SW_TITLEBAR|
	// 		sciter.SW_CONTROLS|
	// 		sciter.SW_MAIN|
	// 		sciter.SW_ENABLE_DEBUG, rect)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		log.Fatal()
	// 	}
	// 	childWin.SetTitle("VS Player")
	// 	childWin.LoadFile("./style/UI_for_PvP.html")
	// 	win = childWin
	// 	text := Test(win)
	// 	if text == "stop" {
	// 		os.Exit(0)
	// 	}
	// 	childWin.Show()
	// 	childWin.Run()
	// })

	// //win = saveWin

	// win.Show()
	// win.Run()

	t.Start()
}

//AIPlays -> The main AI function for the GUI
//(The signature is this and only this (by default) if you want to use the function is the .tis file)
func AIPlays(vals ...*sciter.Value) *sciter.Value {
	log.Fatal("fsdjklfsdkldfajlkf")
	if !board.IsFull() {
		os.Exit(0)
	}
	fmt.Println("before evaluate")

	winner := -1
	if board.Evaluate() != 0 {
		winner = 10
		fmt.Println("UNFORTUNATELY, YOU LOSE!!! THE COMPUTER WINS!!!")
		os.Exit(0)
		return sciter.NewValue(winner)
	}

	//Player's turn
	// text := Test(&win)
	// fmt.Println(text)
	// ij := strings.Split(text, " ")
	// returnValue1, _ := strconv.Atoi(ij[0])
	// returnValue2, _ := strconv.Atoi(ij[1])
	// board.SetPossition(game.PLAYER, returnValue1, returnValue2)
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
	board.PrintBoard()
	//end of AI turn

	if winner != -1 {
		fmt.Println("winner = ", winner)
		return sciter.NewValue(winner)
	}

	return sciter.NewValue(move.Row*3 + move.Col + 1)
}
