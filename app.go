package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) LaunchGame(build string) error {

	//if  {
	//
	//}

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
