package tui

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Media type constants.
var (
	ImageExtensions = []string{".png", ".jpg", ".jpeg", ".webp", ".gif", ".bmp"}
	VideoExtensions = []string{".mp4", ".webm", ".mov", ".avi", ".mkv"}
	AudioExtensions = []string{".mp3", ".wav", ".ogg", ".flac", ".aac", ".m4a"}
)

// MediaState tracks loading state.
type MediaState int

const (
	MediaStateLoading MediaState = iota
	MediaStateLoaded
	MediaStateError
)

// MediaModel handles media preview rendering.
type MediaModel struct {
	url       string
	mediaType string // "image", "video", "audio"
	localPath string
	state     MediaState
	errMsg    string
	rendered  string // chafa output
	duration  string
	theme     *Theme
	width     int

	// Skeleton animation
	frame int
	timer time.Time
}

// NewMediaModel creates a new media preview.
func NewMediaModel(url, mediaType string, theme *Theme, width int) *MediaModel {
	return &MediaModel{
		url:       url,
		mediaType: mediaType,
		state:     MediaStateLoading,
		theme:     theme,
		width:     width,
		timer:     time.Now(),
	}
}

// TickMsg advances the skeleton animation.
type MediaTickMsg struct {
	URL string
}

// TickCmd returns a command for skeleton animation.
func (m *MediaModel) TickCmd() tea.Cmd {
	return tea.Tick(70*time.Millisecond, func(t time.Time) tea.Msg {
		return MediaTickMsg{URL: m.url}
	})
}

// LoadCmd returns a command that downloads and renders the media.
func (m *MediaModel) LoadCmd() tea.Cmd {
	url := m.url
	mediaType := m.mediaType
	return func() tea.Msg {
		localPath, _, _, err := downloadAndRender(url, mediaType)
		if err != nil {
			return MediaLoadedMsg{URL: url, Err: err}
		}
		return MediaLoadedMsg{
			URL:       url,
			MediaType: mediaType,
			LocalPath: localPath,
		}
	}
}

// SetLoaded updates the model with loaded content.
func (m *MediaModel) SetLoaded(localPath, rendered, duration string) {
	m.state = MediaStateLoaded
	m.localPath = localPath
	m.rendered = rendered
	m.duration = duration
}

// SetError marks the media as failed.
func (m *MediaModel) SetError(err string) {
	m.state = MediaStateError
	m.errMsg = err
}

// View renders the media preview.
func (m *MediaModel) View() string {
	switch m.state {
	case MediaStateLoading:
		return m.renderSkeleton()
	case MediaStateLoaded:
		return m.renderLoaded()
	case MediaStateError:
		return m.renderError()
	}
	return ""
}

func (m *MediaModel) renderSkeleton() string {
	m.frame++
	skeleton := []string{
		"╭──────────────────────────────────╮",
		"│  ◇ · · · ◇ · · · ◇ · · · ◇ · · │",
		"│  · · ◇ · · · ◇ · · · ◇ · · · ◇ │",
		"│  ◇ · · · ◇ · · · ◇ · · · ◇ · · │",
		"│  · · ◇ · · · ◇ · · · ◇ · · · ◇ │",
		"│  ◇ · · · ◇ · · · ◇ · · · ◇ · · │",
		"│  · · ◇ · · · ◇ · · · ◇ · · · ◇ │",
		"│  ◇ · · · ◇ · · · ◇ · · · ◇ · · │",
		"│  · · ◇ · · · ◇ · · · ◇ · · · ◇ │",
		"╰──────────────────────────────────╯",
	}

	shinePos := m.frame % 36
	var sb strings.Builder
	for _, line := range skeleton {
		runes := []rune(line)
		for j, r := range runes {
			dist := abs(j - shinePos)
			if dist < 3 {
				sb.WriteString(lipgloss.NewStyle().Foreground(m.theme.Primary).Render(string(r)))
			} else if dist < 6 {
				sb.WriteString(lipgloss.NewStyle().Foreground(m.theme.TextMuted).Render(string(r)))
			} else {
				sb.WriteString(lipgloss.NewStyle().Foreground(m.theme.BorderDim).Render(string(r)))
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (m *MediaModel) renderLoaded() string {
	if m.rendered == "" {
		return lipgloss.NewStyle().Foreground(m.theme.TextMuted).Render("[media loaded: " + m.url + "]")
	}

	var sb strings.Builder
	sb.WriteString(m.rendered)
	sb.WriteByte('\n')

	// Footer
	icon := "🖼"
	if m.mediaType == "video" {
		icon = "🎬"
	} else if m.mediaType == "audio" {
		icon = "🎵"
	}

	footer := icon + " " + m.mediaType
	if m.duration != "" {
		footer += " " + m.duration
	}
	footer += "  " + lipgloss.NewStyle().Foreground(m.theme.TextMuted).Render(m.url)

	sb.WriteString(lipgloss.NewStyle().Foreground(m.theme.TextMuted).Render(footer))
	return sb.String()
}

func (m *MediaModel) renderError() string {
	return lipgloss.NewStyle().Foreground(m.theme.Error).Render("✗ " + m.errMsg)
}

// downloadAndRender downloads media and renders it via chafa.
func downloadAndRender(url, mediaType string) (localPath, duration, rendered string, err error) {
	// Check chafa
	chafaPath, err := exec.LookPath("chafa")
	if err != nil {
		return "", "", "", fmt.Errorf("chafa not found (install: brew install chafa)")
	}

	// Download to temp file
	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", "", "", fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	ext := filepath.Ext(url)
	if ext == "" || len(ext) > 6 {
		ext = ".tmp"
	}

	tmpFile, err := os.CreateTemp("", "genmedia-*"+ext)
	if err != nil {
		return "", "", "", fmt.Errorf("temp file failed: %w", err)
	}

	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return "", "", "", fmt.Errorf("download failed: %w", err)
	}
	tmpFile.Close()
	localPath = tmpFile.Name()

	switch mediaType {
	case "image":
		out, err := exec.Command(chafaPath, "--size", "60x20", "--animate=off", "--format=symbols", "--color-space=din99d", localPath).Output()
		if err != nil {
			return localPath, "", "", fmt.Errorf("chafa render failed: %w", err)
		}
		rendered = strings.TrimRight(string(out), "\n")

	case "video":
		// Get duration
		dur, _ := exec.Command("ffprobe", "-v", "quiet", "-show_entries", "format=duration",
			"-of", "default=noprint_wrappers=1:nokey=1", localPath).Output()
		duration = strings.TrimSpace(string(dur)) + "s"

		// Extract thumbnail
		thumbPath := localPath + ".thumb.png"
		exec.Command("ffmpeg", "-y", "-i", localPath, "-vframes", "1", "-q:v", "2", thumbPath).Run()

		if _, err := os.Stat(thumbPath); err == nil {
			out, err := exec.Command(chafaPath, "--size", "60x20", "--animate=off", "--format=symbols", "--color-space=din99d", thumbPath).Output()
			if err == nil {
				rendered = strings.TrimRight(string(out), "\n")
			}
			os.Remove(thumbPath)
		}

	case "audio":
		// Get duration
		dur, _ := exec.Command("ffprobe", "-v", "quiet", "-show_entries", "format=duration",
			"-of", "default=noprint_wrappers=1:nokey=1", localPath).Output()
		duration = strings.TrimSpace(string(dur)) + "s"

		// Generate waveform
		wavePath := localPath + ".wave.png"
		exec.Command("ffmpeg", "-y", "-i", localPath, "-filter_complex",
			"showwavespic=s=640x240:colors=#4a9eff", "-frames:v", "1", wavePath).Run()

		if _, err := os.Stat(wavePath); err == nil {
			out, err := exec.Command(chafaPath, "--size", "60x10", "--animate=off", "--format=symbols", "--color-space=din99d", wavePath).Output()
			if err == nil {
				rendered = strings.TrimRight(string(out), "\n")
			}
			os.Remove(wavePath)
		}
	}

	return localPath, duration, rendered, nil
}

// ExtractMediaURLs recursively walks a result dict to find media URLs.
func ExtractMediaURLs(data interface{}) []MediaURL {
	var urls []MediaURL
	extractURLs(data, &urls)
	// Deduplicate
	seen := make(map[string]bool)
	var unique []MediaURL
	for _, u := range urls {
		if !seen[u.URL] {
			seen[u.URL] = true
			unique = append(unique, u)
		}
	}
	return unique
}

// MediaURL holds a URL and its detected type.
type MediaURL struct {
	URL       string
	MediaType string
}

func extractURLs(data interface{}, urls *[]MediaURL) {
	switch v := data.(type) {
	case map[string]interface{}:
		// Check priority keys
		for _, key := range []string{"url", "image", "images", "output", "audio", "video"} {
			if val, ok := v[key]; ok {
				extractURLs(val, urls)
			}
		}
		// Check all other values
		for k, val := range v {
			switch k {
			case "url", "image", "images", "output", "audio", "video":
				continue // already handled
			}
			extractURLs(val, urls)
		}
	case []interface{}:
		for _, item := range v {
			extractURLs(item, urls)
		}
	case string:
		if isMediaURL(v) {
			mt := detectMediaType(v)
			if mt != "" {
				*urls = append(*urls, MediaURL{URL: v, MediaType: mt})
			}
		}
	}
}

func isMediaURL(s string) bool {
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}

func detectMediaType(rawURL string) string {
	lower := strings.ToLower(rawURL)
	// Strip query string for extension check
	path := lower
	if idx := strings.Index(path, "?"); idx >= 0 {
		path = path[:idx]
	}
	// Check by path suffix
	for _, ext := range ImageExtensions {
		if strings.HasSuffix(path, ext) {
			return "image"
		}
	}
	for _, ext := range VideoExtensions {
		if strings.HasSuffix(path, ext) {
			return "video"
		}
	}
	for _, ext := range AudioExtensions {
		if strings.HasSuffix(path, ext) {
			return "audio"
		}
	}
	// Check by domain
	if strings.Contains(lower, "fal.media") || strings.Contains(lower, "fal-cdn") {
		return "image" // default for fal URLs
	}
	return ""
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
