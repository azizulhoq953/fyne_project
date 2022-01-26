package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowDashbord(a fyne.App) {
	win := myWindow

	headLable := canvas.NewText("Welcome To Minierp", color.RGBA{255, 20, 147, 255})
	headLable.TextSize = 40

	btn1 := widget.NewButton("Add New Client", func() {
		ShowClient(myApp)
	})
	btn2 := widget.NewButton("All Client", func() {
		ShowData(myApp)
	})
	totalClient := processAllClientData()
	totalClientCount := strconv.Itoa((len(totalClient) - 1))

	card1 := widget.NewCard(
		"Total Client",
		totalClientCount,
		canvas.NewRectangle(color.White),
	)

	card2 := widget.NewCard(
		"Total Product=54",
		"",
		canvas.NewImageFromFile("img/icon.png"),
	)

	card3 := widget.NewCard(
		"Total Client=54",
		"",
		canvas.NewLine(color.Opaque),
	)

	win.SetContent(
		container.NewVBox(
			container.NewGridWithColumns(1,
				container.NewGridWithColumns(1,
					container.NewCenter(headLable),
				),

				container.NewGridWithColumns(3,
					container.NewVBox(btn1, btn2),
					container.NewVBox(card1),
					container.NewVBox(card3),
				), container.NewGridWithColumns(3,
					container.NewVBox(),
					container.NewVBox(card1),
					container.NewVBox(card2),
				),
			),
		),
	)
	win.Show()
}
