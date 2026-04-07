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
	"github.com/roblillack/spot"
)

type AIStatus int

const (
	StatusIdle AIStatus = iota
	StatusThinking
	StatusDone
	StatusError
)

type App struct {
	app    spot.App
	status AIStatus
	watcher *fsnotify.Watcher
}

func main() {
	app := &App{
		status: StatusIdle,
	}

	// Initialize the menu bar app
	spot.Run(app)
}

func (a *App) Init(app spot.App) {
	a.app = app
	
	// Set initial menu bar icon
	a.updateMenuBarIcon()
	
	// Start watching for AI IDE activity
	go a.watchAIActivity()
	
	// Setup menu
	a.setupMenu()
}

func (a *App) setupMenu() {
	menu := a.app.NewMenu()
	
	// Status item
	statusItem := menu.AddItem("Status: Idle", "", nil)
	statusItem.SetEnabled(false)
	
	menu.AddSeparator()
	
	// Preferences
	menu.AddItem("Preferences...", "p", func() {
		a.showPreferences()
	})
	
	menu.AddSeparator()
	
	// Quit
	menu.AddItem("Quit", "q", func() {
		a.app.Quit()
	})
	
	a.app.SetMenu(menu)
}

func (a *App) updateMenuBarIcon() {
	var icon string
	switch a.status {
	case StatusIdle:
		icon = "💤"
	case StatusThinking:
		icon = "🤔"
	case StatusDone:
		icon = "✅"
	case StatusError:
		icon = "❌"
	}
	
	a.app.SetTitle(icon)
}

func (a *App) watchAIActivity() {
	var err error
	a.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer a.watcher.Close()
	
	// Watch Kiro hooks directory
	homeDir, _ := os.UserHomeDir()
	kiroHooksDir := filepath.Join(homeDir, ".kiro", "hooks")
	
	if _, err := os.Stat(kiroHooksDir); err == nil {
		a.watcher.Add(kiroHooksDir)
	}
	
	// Watch for hook executions
	for {
		select {
		case event, ok := <-a.watcher.Events:
			if !ok {
				return
			}
			
			if event.Op&fsnotify.Write == fsnotify.Write {
				a.handleHookEvent(event.Name)
			}
			
		case err, ok := <-a.watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}

func (a *App) handleHookEvent(filename string) {
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
				a.onAIDone()
			}
		}
	}
}

func (a *App) onAIDone() {
	// Update status
	a.status = StatusDone
	a.updateMenuBarIcon()
	
	// Play sound
	a.playSound()
	
	// Show notification
	a.showNotification("AI Done", "Code generation complete!")
	
	// Reset to idle after 3 seconds
	time.AfterFunc(3*time.Second, func() {
		a.status = StatusIdle
		a.updateMenuBarIcon()
	})
}

func (a *App) playSound() {
	cmd := exec.Command("afplay", "/System/Library/Sounds/Glass.aiff")
	cmd.Run()
}

func (a *App) showNotification(title, message string) {
	script := fmt.Sprintf(`display notification "%s" with title "%s" sound name "Glass"`, message, title)
	cmd := exec.Command("osascript", "-e", script)
	cmd.Run()
}

func (a *App) showPreferences() {
	// TODO: Implement preferences window
	a.showNotification("Preferences", "Preferences coming soon!")
}
