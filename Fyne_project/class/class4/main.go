package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("My Form")
	myWindow.Resize(fyne.NewSize(400, 600))

	name := widget.NewEntry()
	message := widget.NewMultiLineEntry()
	row1 := widget.NewFormItem("Name", name)
	row2 := widget.NewFormItem("Message", message)

	wform := widget.NewForm(row1, row2)
	wform.SubmitText = "Save"
	wform.OnSubmit = func() {
		name := name.Text
		message := message.Text
		myData := fmt.Sprintf(`%v %v`, name, message)
		dialog.NewInformation("confirmation", myData, myWindow).Show()

	}
	wform.OnCancel = func() {}

	// labelName := widget.NewLabel("Name")
	// inputName := widget.NewEntry()

	// row1 := container.NewHBox(labelName, inputName)
	myWindow.SetContent(wform)
	//container.NewHBox(row1)

	myWindow.ShowAndRun()
}
