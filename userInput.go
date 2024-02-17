package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"syscall"

	"github.com/audrenbdb/goforeground"
	"github.com/kjk/w"
	"golang.design/x/hotkey"
)

type R struct {
	Titles []string  `json:"titles"`
	Handle []uintptr `json:"descs"`
	Id     []int     `json:"id"`
}

var r = &R{}

var choice = uintptr(0)

var result = getData(r)

func userInput() {

	wg := sync.WaitGroup{}
	wg.Add(9)
	go func() {
		defer wg.Done()

		err := listenHotkey(hotkey.Key1, hotkey.ModCtrl, hotkey.ModShift)
		if err != nil {
			log.Println(err)
		}
	}()
	go func() {
		defer wg.Done()

		err := listenHotkey(hotkey.Key2, hotkey.ModCtrl, hotkey.ModShift)
		if err != nil {
			log.Println(err)
		}
	}()
	go func() {
		defer wg.Done()

		err := listenHotkey(hotkey.Key3, hotkey.ModCtrl, hotkey.ModShift)
		if err != nil {
			log.Println(err)
		}
	}()
	go func() {
		defer wg.Done()

		err := listenHotkey(hotkey.Key4, hotkey.ModCtrl, hotkey.ModShift)
		if err != nil {
			log.Println(err)
		}
	}()
	go func() {
		defer wg.Done()

		err := listenHotkey(hotkey.Key5, hotkey.ModCtrl, hotkey.ModShift)
		if err != nil {
			log.Println(err)
		}
	}()
	go func() {
		defer wg.Done()

		err := listenHotkey(hotkey.Key6, hotkey.ModCtrl, hotkey.ModShift)
		if err != nil {
			log.Println(err)
		}
	}()
	go func() {
		defer wg.Done()

		err := listenHotkey(hotkey.Key7, hotkey.ModCtrl, hotkey.ModShift)
		if err != nil {
			log.Println(err)
		}
	}()
	go func() {
		defer wg.Done()

		err := listenHotkey(hotkey.Key8, hotkey.ModCtrl, hotkey.ModShift)
		if err != nil {
			log.Println(err)
		}
	}()
	go func() {
		defer wg.Done()

		err := listenHotkey(hotkey.Key9, hotkey.ModCtrl, hotkey.ModShift)
		if err != nil {
			log.Println(err)
		}
	}()
	wg.Wait()

	// prompt := &survey.Select{
	// 	Renderer:      survey.Renderer{},
	// 	Message:       "app list: ",
	// 	Options:       result.Titles,
	// 	Default:       nil,
	// 	PageSize:      0,
	// 	VimMode:       false,
	// 	FilterMessage: "",
	// }

	// survey.AskOne(prompt, &choice, survey.WithValidator(survey.Required))

	// fmt.Println("selected app: ", choice)

	// Assuming empty class name
	// w32.ShowWindow(w32.HWND(result.Handle[indexSelected]), w32.SW_SHOW) // does not fucking work idk
	// updateWindow(choice, result)

}

func updateWindow(choice int) {
	// indexSelected := findIndex(choice, result.Titles) // unneeded for now

	indexSelected := choice

	className, err := syscall.UTF16PtrFromString("")
	if err != nil {
		log.Fatal(err)
	}

	windowName, err := syscall.UTF16PtrFromString(result.Titles[indexSelected])
	if err != nil {
		log.Fatal(err)
	}

	hwnd := w.FindWindowWSys(className, windowName)
	// w.SetFocusSys(hwnd)
	w.UpdateWindowSys(hwnd)

	goforeground.Activate(result.Id[indexSelected])

	fmt.Println("selected: ", result.Titles[indexSelected])
}

func getData(r *R) R {

	f, err := os.Open("out.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	for decoder.More() {
		var data foundWindows
		err = decoder.Decode(&data)
		if err != nil {
			log.Fatalln(err)
		}

		r.Titles = append(r.Titles, data.Name)
		r.Handle = append(r.Handle, uintptr(data.Handle))
		r.Id = append(r.Id, data.Id)
	}

	return R{r.Titles, r.Handle, r.Id}
}
