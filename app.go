package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	// "runtime"
	// "github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/runtime"
    // "github.com/wailsapp/wails/v2/pkg/options"
  	// "github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// sartup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// const (
// 	Windows string = "windows"
// 	Mac        = "mac"
// 	Linux        = "linux"
// )

// let currentOS = runtime.GOOS

// https://archive.org/download/flashplayer32_0r0_363_win_sa

// type ImportEvent struct {
// 	Status  string   `json:"status"`
// 	Objekte []Objekt `json:"objekte"`
// }


func (a *App) InitializeApp() error {
	// First - get current info
	// Get OS info

	os := runtime.Environment(a.ctx)
	runtime.EventsEmit(a.ctx, "infoLog", fmt.Sprintf("Platform: %s %v", os.Platform, os.Arch))
    
	// Linux, Windows, Mac - architecture, version, etc

	latest, err := getLatestBuild()

	if err != nil {
		runtime.EventsEmit(a.ctx, "infoLog", fmt.Sprintf("Could not retrieve latest build %s", err))
		return nil
	}
	localVersion, _ := createBuildFolderAndVersionFile()
	runtime.EventsEmit(a.ctx, "infoLog", fmt.Sprintf("Latest version: %d \n", latest.ID))
	runtime.EventsEmit(a.ctx, "infoLog", fmt.Sprintf("Current version: %d ", localVersion))
	
	// Get server info - is online, latest version, runtime info + links

	// 


	// Pass to frontend
	return nil
}

func (a *App) LaunchGame(build string,runtimeName string) error {

	latest, err := patcher()
	if err != nil {
		return errors.New("error on getting latest files")
	}

	fmt.Print(latest.ID)

	fPath := filepath.Join(".", buildFolder, fmt.Sprintf("bymr-%s.swf", build))

	if !fileExists(fPath) {
		fmt.Print("cannot find file: ", fPath)
		return errors.New("cannot find swf build")
	}

	pPath := filepath.Join(".", "flashplayer", "flashplayer_32.exe")
	if !fileExists(pPath) {
		fmt.Print("cannot find file: ", fPath)
		return errors.New("cannot find flashplayer")
	}
	cmd := exec.Command(pPath, fPath)
	//     cmd.SysProcAttr = &syscall.SysProcAttr{
	//         HideWindow:    true,
	//         CreationFlags: 0x08000000,
	//     }
	if err := cmd.Start(); err != nil {
		log.Println("[BYMR LAUNCHER] Failed to start BYMR build %s: %v", build, err)
		return err
	}
	return nil
}
