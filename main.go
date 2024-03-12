package main

import (
	"fmt"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/xuri/excelize/v2"
)

func main() {

	type Details struct {
		CompanyName string
		PostingAge  string
		JobID       string
		Country     string
		Location    string
		Publication string
		SalaryMax   string
		SalaryMin   string
		SalaryType  string
		JobTitle    string
	}
	var company_id int

	file, error := excelize.OpenFile("Project2Data.xlsx")
	if error != nil {
		log.Fatal(error)
	}

	rowss, err := file.GetRows("Comp490 Jobs")
	if err != nil {
		fmt.Println(err)
		return
	}

	var rows = rowss[1:]

	var details []Details
	for _, v := range rows {
		details = append(details, Details{CompanyName: v[0], PostingAge: v[1], JobID: v[2],
			Country: v[3], Location: v[4], Publication: v[5], SalaryMax: v[6], SalaryMin: v[7],
			SalaryType: v[8], JobTitle: v[9]})
	}

	a := app.New()
	w := a.NewWindow("arshjotHomework3")
	w.Resize(fyne.NewSize(400, 400))

	// entry widget for the details of the new company if you want to enter
	entry_CompanyName := widget.NewEntry()
	entry_CompanyName.SetPlaceHolder("Enter company name here...")

	entry_PostingAge := widget.NewEntry()
	entry_PostingAge.SetPlaceHolder("Enter posting age here...")

	entry_JobID := widget.NewEntry()
	entry_JobID.SetPlaceHolder("Enter job ID here...")

	entry_Country := widget.NewEntry()
	entry_Country.SetPlaceHolder("Enter country here...")

	entry_Location := widget.NewEntry()
	entry_Location.SetPlaceHolder("Enter location here...")

	entry_Publication := widget.NewEntry()
	entry_Publication.SetPlaceHolder("Enter publication here...")

	entry_SalaryMax := widget.NewEntry()
	entry_SalaryMax.SetPlaceHolder("Enter salary max here...")

	entry_SalaryMin := widget.NewEntry()
	entry_SalaryMin.SetPlaceHolder("Enter salary min here...")

	entry_SalaryType := widget.NewEntry()
	entry_SalaryType.SetPlaceHolder("Enter salary type here...")

	entry_JobTitle := widget.NewEntry()
	entry_JobTitle.SetPlaceHolder("Enter job title here...")

	// this function is used to pt everything in the form of a list
	list := widget.NewList(
		func() int { return len(details) },

		func() fyne.CanvasObject { return widget.NewLabel("") },

		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(details[lii].CompanyName)
		},
	)

	// this is the function to select the spefic data row from the list
	list.OnSelected = func(id widget.ListItemID) {
		company_id = id
		entry_CompanyName.Text = details[id].CompanyName
		entry_CompanyName.Refresh()
		entry_PostingAge.Text = details[id].PostingAge
		entry_PostingAge.Refresh()
		entry_JobID.Text = details[id].JobID
		entry_JobID.Refresh()
		entry_Country.Text = details[id].Country
		entry_Country.Refresh()
		entry_Location.Text = details[id].Location
		entry_Location.Refresh()
		entry_Publication.Text = details[id].Publication
		entry_Publication.Refresh()
		entry_SalaryMax.Text = details[id].SalaryMax
		entry_SalaryMax.Refresh()
		entry_SalaryMin.Text = details[id].SalaryMin
		entry_SalaryMin.Refresh()
		entry_SalaryType.Text = details[id].SalaryType
		entry_SalaryType.Refresh()
		entry_JobTitle.Text = details[id].JobTitle
		entry_JobTitle.Refresh()
	}

	//this is used to insert the information if added
	insert_btn := widget.NewButton("Insert", func() {

		var details1 Details
		details1.CompanyName = entry_CompanyName.Text
		details1.PostingAge = entry_PostingAge.Text
		details1.JobID = entry_JobID.Text
		details1.Country = entry_Country.Text
		details1.Location = entry_Location.Text
		details1.Publication = entry_Publication.Text
		details1.SalaryMax = entry_SalaryMax.Text
		details1.SalaryMin = entry_SalaryMin.Text
		details1.SalaryType = entry_SalaryType.Text
		details1.JobTitle = entry_JobTitle.Text

		details = append(details, details1)

		sheetTotalRows := len(rowss) + 1
		CompanyNameCell := "A" + strconv.Itoa(sheetTotalRows)
		PostingAgeCell := "B" + strconv.Itoa(sheetTotalRows)
		JobIDCell := "C" + strconv.Itoa(sheetTotalRows)
		CountryCell := "D" + strconv.Itoa(sheetTotalRows)
		LocationCell := "E" + strconv.Itoa(sheetTotalRows)
		PublicationCell := "F" + strconv.Itoa(sheetTotalRows)
		SalaryMaxCell := "G" + strconv.Itoa(sheetTotalRows)
		SalaryMinCell := "H" + strconv.Itoa(sheetTotalRows)
		SalaryTypeCell := "I" + strconv.Itoa(sheetTotalRows)
		JobTitleCell := "J" + strconv.Itoa(sheetTotalRows)
		file.InsertRows("Comp490 Jobs", sheetTotalRows, 1)
		file.SetCellValue("Comp490 Jobs", CompanyNameCell, entry_CompanyName.Text)
		file.SetCellValue("Comp490 Jobs", PostingAgeCell, entry_PostingAge.Text)
		file.SetCellValue("Comp490 Jobs", JobIDCell, entry_JobID.Text)
		file.SetCellValue("Comp490 Jobs", CountryCell, entry_Country.Text)
		file.SetCellValue("Comp490 Jobs", LocationCell, entry_Location.Text)
		file.SetCellValue("Comp490 Jobs", PublicationCell, entry_Publication.Text)
		file.SetCellValue("Comp490 Jobs", SalaryMaxCell, entry_SalaryMax.Text)
		file.SetCellValue("Comp490 Jobs", SalaryMinCell, entry_SalaryMin.Text)
		file.SetCellValue("Comp490 Jobs", SalaryTypeCell, entry_SalaryType.Text)
		file.SetCellValue("Comp490 Jobs", JobTitleCell, entry_JobTitle.Text)
		file.Save()

		entry_CompanyName.Text = ""
		entry_PostingAge.Text = ""
		entry_JobID.Text = ""
		entry_Country.Text = ""
		entry_Location.Text = ""
		entry_Publication.Text = ""
		entry_SalaryMax.Text = ""
		entry_SalaryMin.Text = ""
		entry_SalaryType.Text = ""
		entry_JobTitle.Text = ""
		entry_CompanyName.Refresh()
		entry_PostingAge.Refresh()
		entry_JobID.Refresh()
		entry_Country.Refresh()
		entry_Location.Refresh()
		entry_Publication.Refresh()
		entry_SalaryMax.Refresh()
		entry_SalaryMin.Refresh()
		entry_SalaryType.Refresh()
		entry_JobTitle.Refresh()
		list.Refresh()
	})

	/// this is used to delete any of the information you want to delete-+
	delete_button := widget.NewButton("Delete", func() {
		for k := range details {
			if company_id == k {
				details = append(details[:company_id], details[company_id+1:]...)
				file.RemoveRow("Comp490 Jobs", company_id+2)
				file.Save()

				entry_CompanyName.Text = ""
				entry_PostingAge.Text = ""
				entry_JobID.Text = ""
				entry_Country.Text = ""
				entry_Location.Text = ""
				entry_Publication.Text = ""
				entry_SalaryMax.Text = ""
				entry_SalaryMin.Text = ""
				entry_SalaryType.Text = ""
				entry_JobTitle.Text = ""
				entry_CompanyName.Refresh()
				entry_PostingAge.Refresh()
				entry_JobID.Refresh()
				entry_Country.Refresh()
				entry_Location.Refresh()
				entry_Publication.Refresh()
				entry_SalaryMax.Refresh()
				entry_SalaryMin.Refresh()
				entry_SalaryType.Refresh()
				entry_JobTitle.Refresh()
				list.Refresh()
			}
		}
	})

	// to update the data
	update_button := widget.NewButton("Update", func() {

		var details1 Details
		details1.CompanyName = entry_CompanyName.Text
		details1.PostingAge = entry_PostingAge.Text
		details1.JobID = entry_JobID.Text
		details1.Country = entry_Country.Text
		details1.Location = entry_Location.Text
		details1.Publication = entry_Publication.Text
		details1.SalaryMax = entry_SalaryMax.Text
		details1.SalaryMin = entry_SalaryMin.Text
		details1.SalaryType = entry_SalaryType.Text
		details1.JobTitle = entry_JobTitle.Text

		details[company_id] = details1
		selectedRow := company_id + 2
		CompanyNameCell := "A" + strconv.Itoa(selectedRow)
		PostingAgeCell := "B" + strconv.Itoa(selectedRow)
		JobIDCell := "C" + strconv.Itoa(selectedRow)
		CountryCell := "D" + strconv.Itoa(selectedRow)
		LocationCell := "E" + strconv.Itoa(selectedRow)
		PublicationCell := "F" + strconv.Itoa(selectedRow)
		SalaryMaxCell := "G" + strconv.Itoa(selectedRow)
		SalaryMinCell := "H" + strconv.Itoa(selectedRow)
		SalaryTypeCell := "I" + strconv.Itoa(selectedRow)
		JobTitleCell := "J" + strconv.Itoa(selectedRow)
		file.SetCellValue("Comp490 Jobs", CompanyNameCell, entry_CompanyName.Text)
		file.SetCellValue("Comp490 Jobs", PostingAgeCell, entry_PostingAge.Text)
		file.SetCellValue("Comp490 Jobs", JobIDCell, entry_JobID.Text)
		file.SetCellValue("Comp490 Jobs", CountryCell, entry_Country.Text)
		file.SetCellValue("Comp490 Jobs", LocationCell, entry_Location.Text)
		file.SetCellValue("Comp490 Jobs", PublicationCell, entry_Publication.Text)
		file.SetCellValue("Comp490 Jobs", SalaryMaxCell, entry_SalaryMax.Text)
		file.SetCellValue("Comp490 Jobs", SalaryMinCell, entry_SalaryMin.Text)
		file.SetCellValue("Comp490 Jobs", SalaryTypeCell, entry_SalaryType.Text)
		file.SetCellValue("Comp490 Jobs", JobTitleCell, entry_JobTitle.Text)
		file.Save()

		entry_CompanyName.Text = ""
		entry_PostingAge.Text = ""
		entry_JobID.Text = ""
		entry_Country.Text = ""
		entry_Location.Text = ""
		entry_Publication.Text = ""
		entry_SalaryMax.Text = ""
		entry_SalaryMin.Text = ""
		entry_SalaryType.Text = ""
		entry_JobTitle.Text = ""
		entry_CompanyName.Refresh()
		entry_PostingAge.Refresh()
		entry_JobID.Refresh()
		entry_Country.Refresh()
		entry_Location.Refresh()
		entry_Publication.Refresh()
		entry_SalaryMax.Refresh()
		entry_SalaryMin.Refresh()
		entry_SalaryType.Refresh()
		entry_JobTitle.Refresh()
		list.Refresh()

	})

	w.SetContent(
		// creating HSplit container for display
		container.NewHSplit(

			list,

			container.NewVBox(
				entry_CompanyName, entry_PostingAge, entry_JobID, entry_Country, entry_Location,
				entry_Publication, entry_SalaryMax, entry_SalaryMin, entry_SalaryType,
				entry_JobTitle,
				delete_button, insert_btn, update_button,
			),
		),
	)

	w.ShowAndRun()
}
