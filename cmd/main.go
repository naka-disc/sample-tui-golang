package main

import (
	"log"
	"os"

	"github.com/rivo/tview"
)

// メニュー用の構造体
type Menu struct {
	name  string
	value string
}

// エントリーポイント
func main() {
	// 大元になるアプリケーションインスタンス生成
	app := tview.NewApplication()

	// 左のペイン用のリストを作成
	menuList := tview.NewList().ShowSecondaryText(false)
	menuList.SetBorder(true).SetTitle("Menu")
	// 左のペインでESCキーが押されたら終了
	menuList.SetDoneFunc(func() {
		app.Stop()
		os.Exit(0)
	})

	// 右のペイン用のテーブルを作成
	mainTable := tview.NewTable().SetBorders(true)
	mainTable.SetBorder(true).SetTitle("Contents")

	var menuListvalueList []Menu
	menuListvalueList = append(menuListvalueList, Menu{name: "Favorite", value: "S"})
	for _, v := range menuListvalueList {
		menuList.AddItem(
			v.name, "", 0, func() {
				mainTable.Clear()

				mainTable.SetCellSimple(0, 0, "Test")
				mainTable.SetCellSimple(0, 1, "これはテストです。")

				mainTable.SetCellSimple(1, 0, "Test")
				mainTable.SetCellSimple(1, 1, "本当にテストです。")

				mainTable.SetCellSimple(2, 0, "Test")
				mainTable.SetCellSimple(3, 0, "Test")
				mainTable.SetCellSimple(4, 0, "Test")

			})
	}

	// ウィジェットをボックスに配置
	flex := tview.NewFlex()
	flex.AddItem(menuList, 0, 1, true)
	flex.AddItem(mainTable, 0, 5, false)

	// アプリケーション駆動
	if err := app.SetRoot(flex, true).Run(); err != nil {
		log.Printf("Error running application: %s\n", err)
	}
}
