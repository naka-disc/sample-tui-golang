package main

import (
	"log"
	"os"

	"github.com/rivo/tview"
)

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

	// ウィジェットをボックスに配置
	flex := tview.NewFlex()
	flex.AddItem(menuList, 0, 1, true)
	flex.AddItem(mainTable, 0, 5, false)

	// アプリケーション駆動
	if err := app.SetRoot(flex, true).Run(); err != nil {
		log.Printf("Error running application: %s\n", err)
	}
}
