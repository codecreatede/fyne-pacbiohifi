/*

 A pacbiohifi profiler written in Go, which streams directly the reads using the API also
 and summarizes the information either using the PacBio or the OxfordNanopore.
 Author Gaurav Sablok
 Universitat Potsdam
 Date 2024-8-8

*/

// Welcome window made

package main

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {

	app := app.New()
	windowapp := app.NewWindow("Welcome to PacBioHifi")
	windowapp.SetContent(widget.NewLabel("Welcome to PacBioHifi"))
	windowapp.SetContent(canvas.NewText(" This is a PacBioHifi read analyzer written in Go Programming language", color.Black))
	// load the image here
	img := canvas.NewImageFromURI(storage.NewFileURI("FILEPATH"))
	// enabled the grid for the analysis
	Readmer := widget.NewLabel("Plot reads")
	PlotPacBioHifi := widget.NewLabel("PlotPacBioHifi")
	grid := container.New(layout.NewFormLayout(), Readmer, PlotPacBioHifi)
	windowapp.SetContent(grid)
	// adding a new button based window to add the reads path and confirming it for the analysis.
	windowapp.SetContent(widget.NewButton("Upload the PacbioHifi reads"),
		// adding a callback function and other callables
		func() string {
			fmt.Println(path)
		})

	// added a callable click default disable button for the path input.
	addClick := widget.NewButton("Analyze", func() {
		fmt.Println("Analyze PacbioHifi reads")
	})
	addClick.Disable()
	click.OnChanged = func(path string) string {
		click.Disable()
		if path == nil {
			log.Fatal(err)
			fmt.Println("Path to the PacbioHifi reads cant be empty")
		}
		if path != nil {
			click.Enable()
		}
	}

	windowapp.ShowAndRun()
}

// read uploader window and button func to be merged

func openFile (uri fyne.URI) (fyne.URIReadCloser, error) {
    readfile := fyne.CurrentApp().Driver().FileReaderForURI(uri)
    readbuffer := bufio.NewScanner(readfile)
	header := []string{}
	sequences := []string{}
	header := []string{}
	sequences := []string{}
	for readbuffer.Scan() {
		line := readbuffer.Text()
		if string(line[0]) == "A" || string(line[0]) == "T" || string(line[0]) == "G" || string(line[0]) == "C" {
			sequences = append(sequences, line)
		}
		if string(line[0]) == "@" {
			header = append(header, line)
		}
	}
}

// read summarizer window and button and save 

func saveFile (uri fyne.URI) (fyne.URIWriteCloser, error) {
	return fyne.CurrentApp().Driver().FileWriteForURI(uri)
}  

// read plotter window and button.

package main 

import (
	charts "github.com/vicanso/go-charts/v2"
)

func main () {

	func clustereads () {
		// sequence plotter 
		seqCount := []int{}
		seqHeaders := []string{}
		seqgcCount := []int{}
		seqplot := []int{}
		for i,j := range mapmer {
			seqHeaders = append(seqHeaders,i)
			seqCount = append(seqCount, (strings.Count(mapmer[i], "A")+ strings.Count(mapmer[i], "T")+strings.Count(mapmer[i], "G")+strings.Count(mapmer[i], "C")))
		    seqgcCount = append(seqgcCount, strings.Count(mapmer[i], "G")+strings.Count(mapmer[i], "C"))
		}
		for i := range  seqgcCount {
			seqplot = append(seqplot, int(seqgcCount[i])/int(seqCount[i]))
		}
     values, err := charts.LineRender(
     	seqCount, 
     	charts.TitleTextOptionFunc("Line")
     	charts.XAxisDataOptionFunc(seqHeaders)
     ), 
        charts.LegendLabelOptionFunc([]string{
        	"Reads"
        	"BaseCount"
        }, charts.PositionCenter),
    ), 
    if err != nil {
    	panic(err)
    	log.Fatal(err)
    }
}
 