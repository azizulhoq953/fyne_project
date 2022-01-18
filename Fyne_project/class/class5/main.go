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
	name.PlaceHolder = "Enter Name"
	mobile := widget.NewEntry()
	mobile.PlaceHolder = "Enter Your Mobile"
	email := widget.NewEntry()
	email.PlaceHolder = "Enter Your Email"
	Address := widget.NewMultiLineEntry()
	Address.PlaceHolder = "Enter Your Address"
	row1 := widget.NewFormItem("Name", name)
	row2 := widget.NewFormItem("mobile", mobile)
	row3 := widget.NewFormItem("email", email)
	row4 := widget.NewFormItem("Address", Address)

	wform := widget.NewForm(row1, row2, row3, row4)
	wform.SubmitText = "Save"
	wform.OnSubmit = func() {
		name := name.Text
		message := mobile.Text
		// email := email.Text
		// Address := Address.Text
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