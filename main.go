package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/getlantern/systray"
)

type AIStatus int

const (
	StatusIdle AIStatus = iota
	StatusThinking
	StatusDone
	StatusError
)

type App struct {
	status  AIStatus
	watcher *fsnotify.Watcher
}

var app *App

func main() {
	app = &App{
		status: StatusIdle,
	}
	
	systray.Run(onReady, onExit)
}

func onReady() {
	// Set initial menu bar icon
	updateMenuBarIcon()
	
	systray.SetTitle("AI Done")
	systray.SetTooltip("AI IDE Activity Monitor")
	
	// Setup menu
	mStatus := systray.AddMenuItem("Status: Idle", "Current AI status")
	mStatus.Disable()
	
	systray.AddSeparator()
	
	mQuit := systray.AddMenuItem("Quit", "Quit the app")
	
	// Start watching for AI IDE activity
	go watchAIActivity()
	
	// Handle menu clicks
	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func onExit() {
	if app.watcher != nil {
		app.watcher.Close()
	}
}

func updateMenuBarIcon() {
	var icon string
	switch app.status {
	case StatusIdle:
		icon = "💤"
	case StatusThinking:
		icon = "🤔"
	case StatusDone:
		icon = "✅"
	case StatusError:
		icon = "❌"
	}
	
	systray.SetTitle(icon)
}

func watchAIActivity() {
	var err error
	app.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	
	// Watch Kiro hooks directory
	homeDir, _ := os.UserHomeDir()
	kiroHooksDir := filepath.Join(homeDir, ".kiro", "hooks")
	
	if _, err := os.Stat(kiroHooksDir); err == nil {
		app.watcher.Add(kiroHooksDir)
		log.Printf("Watching: %s\n", kiroHooksDir)
	} else {
		log.Printf("Kiro hooks directory not found: %s\n", kiroHooksDir)
	}
	
	// Watch for hook executions
	for {
		select {
		case event, ok := <-app.watcher.Events:
			if !ok {
				return
			}
			
			if event.Op&fsnotify.Write == fsnotify.Write {
				handleHookEvent(event.Name)
			}
			
		case err, ok := <-app.watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}

func handleHookEvent(filename string) {
	// Read the hook file to determine event type
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	
	var hook map[string]interface{}
	if err := json.Unmarshal(data, &hook); err != nil {
		return
	}
	
	// Check if this is an agentStop event
	if when, ok := hook["when"].(map[string]interface{}); ok {
		if eventType, ok := when["type"].(string); ok {
			if eventType == "agentStop" {
				onAIDone()
			}
		}
	}
}

func onAIDone() {
	// Update status
	app.status = StatusDone
	updateMenuBarIcon()
	
	// Play sound
	playSound()
	
	// Show notification
	showNotification("AI Done", "Code generation complete!")
	
	// Reset to idle after 3 seconds
	time.AfterFunc(3*time.Second, func() {
		app.status = StatusIdle
		updateMenuBarIcon()
	})
}

func playSound() {
	cmd := exec.Command("afplay", "/System/Library/Sounds/Glass.aiff")
	cmd.Run()
}

func showNotification(title, message string) {
	script := fmt.Sprintf(`display notification "%s" with title "%s" sound name "Glass"`, message, title)
	cmd := exec.Command("osascript", "-e", script)
	cmd.Run()
}
