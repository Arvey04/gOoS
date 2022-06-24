package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
    "fyne.io/fyne/v2"	
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"	
	"fyne.io/fyne/v2/canvas"

)

func showGalleryApp() {
	a := app.New()
	w := a.NewWindow("Go Image Viewer")
	w.Resize(fyne.NewSize(1200, 720))
	root_src := "C:\\Users\\rohit verma\\Desktop\\Pepcoding\\images"
	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}
	var picsArr []string
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpeg" {
				picsArr = append(picsArr, root_src+"\\"+file.Name());
			}

		}
	}
	
	tabs := container.NewAppTabs(
		
		// container.NewTabItem("Image1", )),
	    container.NewTabItem("Image2", canvas.NewImageFromFile(picsArr[0])),
	)
	for i:= 0;i<len(picsArr);i++{
        tabs.Append(container.NewTabItem("Image",canvas.NewImageFromFile(picsArr[i])))

		}
		tabs.SetTabLocation((container.TabLocationLeading))
	w.SetContent(tabs);
	w.ShowAndRun()
}
