package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"

	"fyne.io/fyne/v2/widget"
)

var (
	myApp    fyne.App    = app.New()
	myWindow fyne.Window = myApp.NewWindow("MiniErp")
	db       *sql.DB
	err      error
)

func init() {

	// Connect to database
	dbcon()
	// db, err = sql.Open("sqlite3", "E:/GOLANG/src/master_academy/golang/minierp/mini-erp.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// db.SetMaxOpenConns(1)

	// log.Println("db connection successful")

}

func main() {

	myWindow.Resize(fyne.NewSize(800, 600))
	mainMenu()

	emailEntry := widget.NewEntry()
	emailEntry.PlaceHolder = "Enter Your Email"

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.PlaceHolder = "Eter Your Password"

	email := widget.NewFormItem("Email", emailEntry)
	password := widget.NewFormItem("Password", passwordEntry)

	loginForm := widget.NewForm(email, password)
	loginForm.SubmitText = "Login"

	loginForm.OnSubmit = func() {
		appEmail := ""
		appPass := ""
		if emailEntry.Text == appEmail && passwordEntry.Text == appPass {
			ShowDashbord(myApp)
		} else {
			dialog.NewInformation("Invalid Email Or Password", " ", myWindow).Show()
		}
	}

	myWindow.SetContent(
		loginForm,
	)
	myWindow.ShowAndRun()
}
