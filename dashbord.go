package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ShowDashbod(a fyne.App) {
	mainMenu()
	win := myWindow
	comInfo, err := GetCompanyInfo("1")
	if err != nil {
		fmt.Println(err)
	}
	comName := comInfo["company_name"]
	comAddress := comInfo["address"]
	comWeb := comInfo["website"]
	comEmail := comInfo["email"]
	comMobile := comInfo["mobile"]

	// Header Section - More prominent and clean
	welcome := fmt.Sprintf("Welcome to %s", comName)
	onlineInfo := fmt.Sprintf("%s | %s | %s", comWeb, comEmail, comMobile)

	headLabel := widget.NewCard(
		"",
		"",
		container.NewVBox(
			widget.NewLabelWithStyle(welcome, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			widget.NewLabel(dataConveter(comAddress)),
			widget.NewSeparator(),
			widget.NewLabel(onlineInfo),
		),
	)

	totalClient := processAllClientData()
	totalClientCount := strconv.Itoa((len(totalClient) - 1))

	totalProduct := processAllProductData()
	totalProductCount := strconv.Itoa((len(totalProduct) - 1))

	clientCard := widget.NewCard(
		"Total Clients",
		"",
		container.NewVBox(
			widget.NewLabelWithStyle(totalClientCount, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			layout.NewSpacer(),
			widget.NewLabel("Active clients in system"),
		),
	)

	productCard := widget.NewCard(
		"Total Products",
		"",
		container.NewVBox(
			widget.NewLabelWithStyle(totalProductCount, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			layout.NewSpacer(),
			widget.NewLabel("Products available"),
		),
	)

	statsSection := container.NewGridWithColumns(2, clientCard, productCard)

	clientSection := widget.NewCard(
		"Client Management",
		"",
		container.NewVBox(
			widget.NewButtonWithIcon("Add New Client", theme.ContentAddIcon(), func() {
				ShowClient(myApp)
			}),
			widget.NewButtonWithIcon("View All Clients", theme.SearchIcon(), func() {
				ShowData(myApp)
			}),
		),
	)

	productSection := widget.NewCard(
		"Product Management",
		"",
		container.NewVBox(
			widget.NewButtonWithIcon("Add Product", theme.ContentAddIcon(), func() {
				ShowProductAdd(myApp)
			}),
			widget.NewButtonWithIcon("Add Product Group", theme.FolderNewIcon(), func() {
				ShowAddGroupItem(myApp)
			}),
			widget.NewButtonWithIcon("View All Products", theme.SearchIcon(), func() {
				ShowAllProduct(myApp)
			}),
		),
	)

	invoiceSection := widget.NewCard(
		"Invoice Management",
		"",
		container.NewVBox(
			widget.NewButtonWithIcon("Invoice Dashboard", theme.DocumentIcon(), func() {
				InvoiceDash(myApp)
			}),
		),
	)

	// Layout - Better organization
	actionButtons := container.NewVBox(
		clientSection,
		productSection,
		invoiceSection,
	)

	// Main content with better spacing
	mainContent := container.NewBorder(
		// Top
		container.NewVBox(
			headLabel,
			widget.NewSeparator(),
		),
		nil,
		// Left
		nil,
		// Right
		nil,
		// Center
		container.NewVBox(
			statsSection,
			widget.NewSeparator(),
			container.NewPadded(actionButtons),
		),
	)

	// Scrollable content for better responsiveness
	scrollContent := container.NewVScroll(mainContent)
	scrollContent.SetMinSize(fyne.NewSize(800, 600))

	win.SetContent(scrollContent)
	win.Resize(fyne.NewSize(900, 650))
	win.Show()
}
