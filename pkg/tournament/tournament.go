package tournament

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"github.com/sharkoo0/Tic-Tac-Toe/pkg/game"
)

var players = [][]string{
	[]string{"1", "2"},
	[]string{"3", "4"},
}

type Tournament struct {
	Games sync.Map
}

func test(w *window.Window) string {
	root1, _ := w.GetRootElement()
	fmt.Println(root1)
	r, _ := root1.SelectById("#10")
	text, _ := r.Text()
	//w.DefineFunction("ai", AIPlays)
	return text
}

func createWindow(board game.Board) {

	rect := sciter.NewRect(250, 250, 250, 250)

	win, err := window.New(sciter.SW_TITLEBAR|
		sciter.SW_CONTROLS|
		sciter.SW_MAIN|
		sciter.SW_ENABLE_DEBUG, rect)
	if err != nil {
		log.Fatal(err)
	}

	fullpath, err := filepath.Abs("./style/main_html.html")
	if err != nil {
		log.Fatal(err)
	}
	win.LoadFile(fullpath)
	win.SetTitle("TIC-TAC-TOE")

	// win = window.Window(*w)
	//saveWin := win

	root, e := win.GetRootElement()
	if e != nil {
		fmt.Println(e.Error())
		log.Fatal()
	}

	buttonP, e := root.SelectById("player")
	if e != nil {
		fmt.Println(e.Error())
		log.Fatal()
	}
	buttonP.OnClick(func() {
		// win.SetTitle("VS PLAYER")
		// err = win.LoadFile("./style/UI_for_PvP.html")
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	log.Fatal()
		// }
		rect := sciter.NewRect(300, 300, 300, 300)
		childWin, err := window.New(sciter.SW_POPUP|
			sciter.SW_TITLEBAR|
			sciter.SW_CONTROLS|
			sciter.SW_MAIN|
			sciter.SW_ENABLE_DEBUG, rect)
		if err != nil {
			fmt.Println(err.Error())
			log.Fatal()
		}
		childWin.SetTitle("VS Player")
		childWin.LoadFile("./style/UI_for_PvP.html")
		win = childWin
		text := test(win)
		if text == "stop" {
			os.Exit(0)
		}
		childWin.Show()
		childWin.Run()
	})

	//win = saveWin

	win.Show()
	win.Run()

	//t.Start()
}

func (t *Tournament) Init() {
	for _, player := range players {
		b := game.Board{}
		b.Make()
		t.Games.Store(fmt.Sprintf("%s-%s", player[0], player[1]), b)
	}
}

func (t *Tournament) game(id string, wg *sync.WaitGroup) {
	//defer wg.Done()

	if b, ok := t.Games.Load(id); ok {
		createWindow(b.(game.Board))
	}

	//time.Sleep(2 * time.Second)

	fmt.Println("hujnq2")
}

func (t *Tournament) Start() {
	//var wg sync.WaitGroup

	index := 1

	t.Games.Range(func(key, value interface{}) bool {
		//wg.Add(index)
		index++
		t.game(key.(string), nil)
		return true
	})

	//wg.Wait()
}
