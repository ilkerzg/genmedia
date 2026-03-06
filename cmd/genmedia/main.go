package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fal-ai/genmedia/internal/auth"
	"github.com/fal-ai/genmedia/internal/config"
	"github.com/fal-ai/genmedia/internal/session"
	"github.com/fal-ai/genmedia/internal/tui"
)

func main() {
	modelFlag := flag.String("m", config.DefaultModel, "LLM model to use")
	sessionFlag := flag.String("s", "", "Session ID to resume ('last' for most recent)")
	flag.Parse()

	// Auth
	authKey := auth.GetCredentials()
	if authKey == "" {
		fmt.Fprintln(os.Stderr, "Warning: Not authenticated. Use /login inside genmedia or set FAL_KEY.")
	}

	// Session store
	store, err := session.NewSessionStore("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not open session database: %v\n", err)
	}
	defer func() {
		if store != nil {
			store.Close()
		}
	}()

	// Handle -s flag
	sessionID := ""
	if *sessionFlag != "" && store != nil {
		if *sessionFlag == "last" {
			sessionID = store.GetLastSessionID()
			if sessionID == "" {
				fmt.Fprintln(os.Stderr, "No previous sessions found.")
			}
		} else {
			sessionID = *sessionFlag
			if sess := store.GetSession(sessionID); sess == nil {
				fmt.Fprintf(os.Stderr, "Session %s not found.\n", sessionID)
				sessionID = ""
			}
		}
	}

	// Preferences
	prefs, _ := config.LoadPreferences()

	// Create app
	app := tui.NewAppModel(tui.AppConfig{
		Model:       *modelFlag,
		AuthKey:     authKey,
		ThemeKey:    config.DefaultTheme,
		SessionID:   sessionID,
		Store:       store,
		Preferences: prefs,
	})

	// Run
	p := tea.NewProgram(app, tea.WithAltScreen())
	app.SetProgram(p)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
