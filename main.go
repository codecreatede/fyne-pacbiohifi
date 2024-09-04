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
	_ "github.com/vicanso/go-charts/v2"
)

func main() {

	app := app.New()
	windowapp := app.NewWindow("Welcome to PacBioHifi")
	windowapp.SetContent(widget.NewLabel("Welcome to PacBioHifi Genome Profiler"))
	windowapp.SetContent(canvas.NewText(" This is a PacBioHifi read analyzer written in Go Programming language", color.Black))
	windowapp.Resize(fyne.NewSize(400,400))
	// load the image here
	// img := canvas.NewImageFromURI(storage.NewFileURI("FILEPATH"))
	windowapp.SetContent(widget.NewButton("Pacbiohifi reads path", func (){
		input := widget.NewEntry()
		input.SetPlaceHolder("Please provide the path for the PacBiohifi reads")
		content := container.NewVBox(input, widget.NewButton("Save the path", func (){
			log.Println("The provided path is:", input.Text)
		}))
	windowapp.SetContent(content)
	}))
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
		charts.TitleTextOptionFunc("Line"),
		charts.XAxisDataOptionFunc(seqHeaders)
	),
        charts.LegendLabelOptionFunc([]string{
        	"Reads"
        	"BaseCount"
        }, charts.PositionCenter),
	),
	if err != nil {
    	log.Fatal(err)
	}
  // saving the file after reading it.
	func saveFile (uri fyne.URI) (fyne.URIWriteCloser, error) {
	return fyne.CurrentApp().Driver().FileWriteForURI(uri)
	}
}
