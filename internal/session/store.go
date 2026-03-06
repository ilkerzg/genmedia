package session

import (
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

// SessionStore manages SQLite persistence for chat sessions.
type SessionStore struct {
	db *sql.DB
}

func defaultDBPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".cache", "fal-dev", "chat_sessions.db")
}

// NewSessionStore creates a new store. Pass empty string for default path.
func NewSessionStore(dbPath string) (*SessionStore, error) {
	if dbPath == "" {
		dbPath = defaultDBPath()
	}
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	s := &SessionStore{db: db}
	if err := s.init(); err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}

func (s *SessionStore) init() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS sessions (
			id TEXT PRIMARY KEY,
			title TEXT NOT NULL DEFAULT '',
			model TEXT NOT NULL DEFAULT '',
			created_at TEXT NOT NULL,
			updated_at TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS messages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			session_id TEXT NOT NULL REFERENCES sessions(id),
			role TEXT NOT NULL,
			content TEXT,
			tool_calls TEXT,
			tool_call_id TEXT,
			created_at TEXT NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_messages_session ON messages(session_id);
	`)
	return err
}

// CreateSession creates a new session and returns its ID.
func (s *SessionStore) CreateSession(model, title string) string {
	id := uuid.New().String()[:12]
	now := time.Now().UTC().Format(time.RFC3339)
	s.db.Exec(
		"INSERT INTO sessions (id, title, model, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		id, title, model, now, now,
	)
	return id
}

// SaveMessage persists a message dict to the database.
func (s *SessionStore) SaveMessage(sessionID string, msg map[string]interface{}) {
	now := time.Now().UTC().Format(time.RFC3339)
	role, _ := msg["role"].(string)
	content, _ := msg["content"].(string)

	var toolCallsJSON *string
	if tc, ok := msg["tool_calls"]; ok && tc != nil {
		data, err := json.Marshal(tc)
		if err == nil {
			str := string(data)
			toolCallsJSON = &str
		}
	}

	var toolCallID *string
	if tcid, ok := msg["tool_call_id"].(string); ok && tcid != "" {
		toolCallID = &tcid
	}

	s.db.Exec(
		"INSERT INTO messages (session_id, role, content, tool_calls, tool_call_id, created_at) VALUES (?, ?, ?, ?, ?, ?)",
		sessionID, role, content, toolCallsJSON, toolCallID, now,
	)
	s.db.Exec("UPDATE sessions SET updated_at = ? WHERE id = ?", now, sessionID)
}

// LoadMessages loads all messages for a session as dicts.
func (s *SessionStore) LoadMessages(sessionID string) []map[string]interface{} {
	rows, err := s.db.Query(
		"SELECT role, content, tool_calls, tool_call_id FROM messages WHERE session_id = ? ORDER BY id",
		sessionID,
	)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var messages []map[string]interface{}
	for rows.Next() {
		var role, content string
		var toolCallsStr, toolCallID sql.NullString
		if err := rows.Scan(&role, &content, &toolCallsStr, &toolCallID); err != nil {
			continue
		}
		msg := map[string]interface{}{
			"role":    role,
			"content": content,
		}
		if toolCallsStr.Valid && toolCallsStr.String != "" {
			var tc interface{}
			if err := json.Unmarshal([]byte(toolCallsStr.String), &tc); err == nil {
				msg["tool_calls"] = tc
			}
		}
		if toolCallID.Valid && toolCallID.String != "" {
			msg["tool_call_id"] = toolCallID.String
		}
		messages = append(messages, msg)
	}
	return messages
}

// UpdateTitle updates a session's title.
func (s *SessionStore) UpdateTitle(sessionID, title string) {
	now := time.Now().UTC().Format(time.RFC3339)
	s.db.Exec("UPDATE sessions SET title = ?, updated_at = ? WHERE id = ?", title, now, sessionID)
}

// UpdateModel updates a session's model.
func (s *SessionStore) UpdateModel(sessionID, model string) {
	now := time.Now().UTC().Format(time.RFC3339)
	s.db.Exec("UPDATE sessions SET model = ?, updated_at = ? WHERE id = ?", model, now, sessionID)
}

// GetSession returns session info as a map, or nil if not found.
func (s *SessionStore) GetSession(sessionID string) map[string]interface{} {
	var id, title, model, createdAt, updatedAt string
	err := s.db.QueryRow(
		"SELECT id, title, model, created_at, updated_at FROM sessions WHERE id = ?",
		sessionID,
	).Scan(&id, &title, &model, &createdAt, &updatedAt)
	if err != nil {
		return nil
	}
	return map[string]interface{}{
		"id":         id,
		"title":      title,
		"model":      model,
		"created_at": createdAt,
		"updated_at": updatedAt,
	}
}

// ListSessions returns recent sessions.
func (s *SessionStore) ListSessions(limit int) []map[string]interface{} {
	rows, err := s.db.Query(
		"SELECT id, title, model, created_at, updated_at FROM sessions ORDER BY updated_at DESC LIMIT ?",
		limit,
	)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var sessions []map[string]interface{}
	for rows.Next() {
		var id, title, model, createdAt, updatedAt string
		if err := rows.Scan(&id, &title, &model, &createdAt, &updatedAt); err != nil {
			continue
		}
		sessions = append(sessions, map[string]interface{}{
			"id":         id,
			"title":      title,
			"model":      model,
			"created_at": createdAt,
			"updated_at": updatedAt,
		})
	}
	return sessions
}

// GetLastSessionID returns the most recent session ID, or empty string.
func (s *SessionStore) GetLastSessionID() string {
	var id string
	err := s.db.QueryRow("SELECT id FROM sessions ORDER BY updated_at DESC LIMIT 1").Scan(&id)
	if err != nil {
		return ""
	}
	return id
}

// Close closes the database.
func (s *SessionStore) Close() {
	if s.db != nil {
		s.db.Close()
	}
}
