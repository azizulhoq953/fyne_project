package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var db *sql.DB
var err error

func init() {

	// Connect to database
	db, err = sql.Open("sqlite3", "E:/GOLANG/src/master_academy/golang/minierp/mini-erp.db")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	defer db.Close()
	log.Println("db connection successful")

	// data := make(url.Values)
	// data.Set("table", "client")
	// data.Set("dbtype", "sqlite3")

	// data.Set("name", "AZIZUR")
	// data.Set("mobile", "01706257588")
	// data.Set("email", "AZIZULHOQ4305@GMAIL.COM")
	// data.Set("address", "House #3pa")
	// pid, err := msql.InsertIntoAnyTable(data, db)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println("successfully Inserted", pid)
}

func main() {

	addClient("azizul", "01706257588", "azizulhoq4305.com", "Barishal")

	os.Exit(1)

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
		name := strings.ToUpper(name.Text)

		mobile := mobile.Text
		email := strings.ToLower(email.Text) //user to lower case convert
		Address := Address.Text

		myData := fmt.Sprintf(`%v %v %s %s`, name, mobile, email, Address)
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
