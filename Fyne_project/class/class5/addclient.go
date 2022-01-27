package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowClient(a fyne.App) {
	win := myWindow
	btnHead := widget.NewButton("Dashbord", func() {
		ShowDashbord(myApp)
	})

	nameEntry := widget.NewEntry()
	nameEntry.PlaceHolder = "Enter Client Name"
	mobileEntry := widget.NewEntry()
	mobileEntry.PlaceHolder = "Enter Client Number"

	emailEntry := widget.NewEntry()
	emailEntry.PlaceHolder = "Enter client Email Address"

	addressEntry := widget.NewEntry()
	addressEntry.PlaceHolder = "client Address"

	name1 := widget.NewFormItem("name", nameEntry)
	mobile1 := widget.NewFormItem("Mobile", mobileEntry)
	email1 := widget.NewFormItem("email", emailEntry)
	address1 := widget.NewFormItem("address", addressEntry)
	clientForm := widget.NewForm(name1, mobile1, email1, address1)

	clientForm.SubmitText = "Save"

	clientForm.OnSubmit = func() {
		name := nameEntry.Text
		mobile := mobileEntry.Text
		email := emailEntry.Text
		address := addressEntry.Text

		id, err := AddClient(name, mobile, email, address)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(id)
	}
	win.SetContent(container.NewVBox(btnHead, clientForm))

	win.Show()
}
