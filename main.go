package main

import (
	"os"
	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	mainWindow := widgets.NewQMainWindow(nil, 0)
	mainWindow.SetWindowTitle("Hallo, Qt in Go!")

	mainWindow.Show()

	widgets.QApplication_Exec()
}