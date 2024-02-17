package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/AllenDang/w32"
	"golang.design/x/hotkey/mainthread"
	"golang.org/x/sys/windows"
)

type foundWindows struct {
	Name   string     `json:"name"`
	Handle w32.HANDLE `json:"handle"`
	Id     int        `json:"id"`
}

func main() {
	file, err := os.OpenFile("out.json", os.O_WRONLY|os.O_TRUNC, 0644) // clear out the file
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	enumWindowsCallback := syscall.NewCallback(enumWindowsProc)
	err = windows.EnumWindows(enumWindowsCallback, nil)
	if err != nil {
		panic(err)
	}

	mainthread.Init(userInput)
	userInput()
}

func enumWindowsProc(hwnd uintptr, lparam w32.LPARAM) uintptr {
	// extendedStyle := w32.GetWindowLong(hwnd, w32.GWL_EXSTYLE)

	// if w32.IsWindowVisible(w32.HWND(hwnd)) {
	windowName := w32.GetWindowText(w32.HWND(hwnd))                                       // get the name of the window
	windowName = strings.TrimSpace(windowName)                                            // Trim white space
	windowFile, err := os.OpenFile("out.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644) // needs to be opened again couse of the limitation of the winapi
	if err != nil {
		log.Fatal(err)
	}
	defer windowFile.Close()

	if IsAltTabWindow(hwnd) {
		handle, id := w32.GetWindowThreadProcessId(w32.HWND(hwnd)) // get id's needed for identifiyng processes
		if err != nil {
			panic(err)
		}

		output, err := json.Marshal(foundWindows{
			Name:   windowName,
			Handle: handle,
			Id:     id,
		})
		if err != nil {
			log.Fatalln(err)
		}

		_, err = windowFile.Write(output)
		if err != nil {
			panic(err)
		}
	}

	return 1
}

func IsAltTabWindow(hwnd uintptr) bool {
	shellWindow := windows.GetShellWindow()

	title := w32.GetWindowText(w32.HWND(hwnd))
	// className, err := windows.GetClassName(hwnd, nil, 0) // left over from copied code xd 😎
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	if windows.HWND(hwnd) == shellWindow || len(title) == 0 || !w32.IsWindowVisible(w32.HWND(hwnd)) {
		return false
	}

	var childHwnd int32
	if childHwnd = w32.GetWindowLong(w32.HWND(hwnd), w32.WS_CHILDWINDOW); childHwnd != 0 && childHwnd != int32(hwnd) {
		return false
	}

	style := w32.GetWindowLong(w32.HWND(hwnd), w32.GWL_STYLE)
	if (style & w32.WS_DISABLED) != 0 {
		return false
	}

	// var cloaked int32
	// hrTemp, _ := w32.DwmGetWindowAttribute(w32.HWND(hwnd), w32.DWMWA_CLOAKED)
	// if hrTemp == windows.S_OK && cloaked == w32.DWMWA_CLOAKED {
	// 	return false
	// }

	return true
}
