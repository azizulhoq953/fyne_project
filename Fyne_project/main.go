// package main

// import (
// 	"image/color"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	//"fyne.io/fyne/v2/layout"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Container")
// 	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

// 	text1 := canvas.NewText("Hello", green)
// 	text2 := canvas.NewText("There", green)
// 	text2.Move(fyne.NewPos(20, 20))
// 	content := container.NewWithoutLayout(text1, text2)
// 	// content := container.New(layout.NewGridLayout(2), text1, text2)

// 	myWindow.SetContent(content)
// 	myWindow.ShowAndRun()
// }



/*package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	myCanvas.SetContent(text)
	go changeContent(myCanvas)

	myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.ShowAndRun()
}

func changeContent(c fyne.Canvas) {
	time.Sleep(time.Second * 2)

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	c.SetContent(canvas.NewRectangle(blue))

	time.Sleep(time.Second * 2)
	c.SetContent(canvas.NewLine(color.Gray{Y: 180}))

	time.Sleep(time.Second * 2)
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)

	time.Sleep(time.Second * 2)
	c.SetContent(canvas.NewImageFromResource(theme.FyneLogo()))
}
*/

 package main

 import (
 	"fyne.io/fyne/v2/app"
 	"fyne.io/fyne/v2/container"
 	"fyne.io/fyne/v2/widget"
 )

 func main() {
 	a := app.New()
	w := a.NewWindow("Hello")

 	hello := widget.NewLabel("Hello Azizul I'm Fyne!")
 	w.SetContent(container.NewVBox(
 		hello,
 		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
 		}),
 	))

 	w.ShowAndRun()
 }