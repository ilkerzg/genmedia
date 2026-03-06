package skills

import (
	"embed"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

//go:embed *.md
var skillFS embed.FS

// SkillEntry describes a skill.
type SkillEntry struct {
	Filename    string
	Description string
	Keywords    []string
}

// SkillCatalog maps skill names to metadata.
var SkillCatalog = map[string]SkillEntry{
	"cinematography":   {"cinematography.md", "Hollywood directing techniques, camera movements, lighting, composition, lens choices, color grading", []string{"cinema", "film", "movie", "director", "camera", "lighting", "shot", "lens", "composition", "color grade"}},
	"video_prompting":  {"video_prompting.md", "AI video generation prompt engineering — model-specific tips, motion keywords, artifact prevention", []string{"video", "animate", "motion", "clip", "wan", "kling", "veo", "sora", "luma", "minimax", "hunyuan", "seedance", "ltx", "hailuo"}},
	"image_prompting":  {"image_prompting.md", "AI image generation prompt engineering — style keywords, quality boosters, model-specific tips", []string{"image", "photo", "picture", "illustration", "render", "flux", "sdxl", "midjourney", "dall-e", "portrait", "landscape"}},
	"commercial":       {"commercial.md", "Advertising and commercial production — product shots, brand videos, social media formats", []string{"product", "ad", "commercial", "brand", "marketing", "social media", "instagram", "tiktok", "advertisement"}},
	"audio_prompting":  {"audio_prompting.md", "AI audio/music generation — genre, mood, instruments, sound effects, voice generation", []string{"audio", "music", "sound", "voice", "speech", "song", "soundtrack", "sfx", "tts"}},
	"workflow_utils":   {"workflow_utils.md", "fal.ai workflow utilities — image/video/audio processing, compositing, format conversion", []string{"workflow", "merge", "split", "overlay", "subtitle", "resize", "crop", "mask", "composite", "convert", "gif", "audio extract"}},
	"character_design": {"character_design.md", "Character creation & consistency — anchor descriptions, multi-view sheets, expressions, costumes, LoRA training, IP-Adapter", []string{"character", "person", "face", "consistent", "identity", "costume", "outfit", "expression", "pose", "character sheet", "turnaround"}},
	"storytelling":     {"storytelling.md", "Multi-shot narrative video production — storyboarding, scene structure, pacing, transitions, shot sequences", []string{"story", "narrative", "storyboard", "scene", "sequence", "multi-shot", "transition", "pacing", "plot", "script"}},
	"social_media":     {"social_media.md", "Platform-optimized content creation — TikTok, Reels, Shorts, hooks, trends, viral formats, engagement", []string{"tiktok", "reels", "shorts", "viral", "hook", "trend", "engagement", "social", "influencer", "content creator"}},
	"motion_design":    {"motion_design.md", "Motion graphics & animation — kinetic typography, logo reveals, transitions, particle effects, UI animation", []string{"motion graphics", "typography", "kinetic", "logo reveal", "transition", "particle", "animation", "title", "lower third", "intro"}},
	"world_building":   {"world_building.md", "Environment & world design — architecture, interior design, landscapes, fantasy/sci-fi worlds, atmosphere", []string{"environment", "world", "architecture", "interior", "landscape", "city", "fantasy world", "sci-fi", "building", "room", "space"}},
	"brand_identity":   {"brand_identity.md", "Visual brand identity — logo design, color systems, typography, brand guidelines, visual consistency", []string{"logo", "brand", "identity", "typography", "color palette", "brand guide", "visual identity", "corporate", "design system"}},
}

var (
	skillCache  = make(map[string]string)
	skillCacheMu sync.RWMutex
)

// LoadSkill loads a skill's full markdown content.
func LoadSkill(name string) string {
	skillCacheMu.RLock()
	if content, ok := skillCache[name]; ok {
		skillCacheMu.RUnlock()
		return content
	}
	skillCacheMu.RUnlock()

	entry, ok := SkillCatalog[name]
	if !ok {
		return ""
	}
	data, err := skillFS.ReadFile(entry.Filename)
	if err != nil {
		return ""
	}
	content := string(data)

	skillCacheMu.Lock()
	skillCache[name] = content
	skillCacheMu.Unlock()
	return content
}

type section struct {
	Slug      string
	Header    string
	Body      string
	LineCount int
}

var nonAlphaRe = regexp.MustCompile(`[^a-z0-9\s]`)
var spaceRe = regexp.MustCompile(`\s+`)

func slugify(headerText string) string {
	text := strings.TrimLeft(headerText, "# ")
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	text = nonAlphaRe.ReplaceAllString(text, " ")
	text = spaceRe.ReplaceAllString(strings.TrimSpace(text), "_")
	return strings.Trim(text, "_")
}

func parseSections(content string, level string) []section {
	prefix := level + " "
	lines := strings.Split(content, "\n")
	var sections []section
	var currentHeader, currentSlug string
	var currentLines []string
	started := false

	for _, line := range lines {
		if strings.HasPrefix(line, prefix) && !strings.HasPrefix(line, prefix+"#") {
			if started {
				sections = append(sections, section{
					Slug:      currentSlug,
					Header:    currentHeader,
					Body:      strings.Join(currentLines, "\n"),
					LineCount: len(currentLines),
				})
			}
			currentHeader = line
			currentSlug = slugify(line)
			currentLines = []string{line}
			started = true
		} else if started {
			currentLines = append(currentLines, line)
		}
	}
	if started {
		sections = append(sections, section{
			Slug:      currentSlug,
			Header:    currentHeader,
			Body:      strings.Join(currentLines, "\n"),
			LineCount: len(currentLines),
		})
	}
	return sections
}

const largeSectionThreshold = 5000

// LoadSkillTOC loads a skill's table of contents with quick reference.
func LoadSkillTOC(name string) string {
	content := LoadSkill(name)
	if content == "" {
		return ""
	}

	lines := strings.Split(content, "\n")
	var title string
	var introLines []string

	for _, line := range lines {
		if strings.HasPrefix(line, "# ") && !strings.HasPrefix(line, "## ") && title == "" {
			title = line
		} else if strings.HasPrefix(line, "## ") {
			break
		} else if title != "" {
			introLines = append(introLines, line)
		}
	}
	intro := strings.TrimSpace(strings.Join(introLines, "\n"))

	sections := parseSections(content, "##")

	var quickRef string
	for _, s := range sections {
		if s.Slug == "quick_reference" {
			quickRef = s.Body
			break
		}
	}

	var parts []string
	parts = append(parts, title)
	if intro != "" {
		parts = append(parts, intro)
	}
	parts = append(parts, "")

	if quickRef != "" {
		parts = append(parts, quickRef)
		parts = append(parts, "")
	}

	parts = append(parts, "## Available Sections")
	for _, s := range sections {
		if s.Slug == "quick_reference" {
			continue
		}
		headerText := strings.TrimSpace(strings.TrimLeft(s.Header, "#"))
		parts = append(parts, fmt.Sprintf("- **%s** — %s (%d lines)", s.Slug, headerText, s.LineCount))
	}
	parts = append(parts, "")
	parts = append(parts, fmt.Sprintf(`Call get_skill("%s", section="<key>") to load a specific section.`, name))

	return strings.Join(parts, "\n")
}

func findSection(sections []section, key string) *section {
	// Exact match
	for i := range sections {
		if sections[i].Slug == key {
			return &sections[i]
		}
	}
	// Token prefix match
	for i := range sections {
		tokens := strings.Split(sections[i].Slug, "_")
		for _, t := range tokens {
			if strings.HasPrefix(t, key) {
				return &sections[i]
			}
		}
	}
	return nil
}

func buildSectionTOC(name, h2Slug, h2Header string, h3Sections []section) string {
	var parts []string
	parts = append(parts, strings.TrimSpace(h2Header))
	parts = append(parts, "")
	parts = append(parts, fmt.Sprintf("This section has %d subsections. Load one at a time:", len(h3Sections)))
	parts = append(parts, "")
	for _, s := range h3Sections {
		headerText := strings.TrimSpace(strings.TrimLeft(s.Header, "#"))
		parts = append(parts, fmt.Sprintf("- **%s** — %s (%d lines)", s.Slug, headerText, s.LineCount))
	}
	parts = append(parts, "")
	parts = append(parts, fmt.Sprintf(`Call get_skill("%s", section="%s") to load a subsection.`, name, h3Sections[0].Slug))
	return strings.Join(parts, "\n")
}

// LoadSkillSection loads a specific section from a skill by slug key.
func LoadSkillSection(name, sectionKey string) string {
	content := LoadSkill(name)
	if content == "" {
		return ""
	}

	h2Sections := parseSections(content, "##")
	key := strings.ToLower(strings.TrimSpace(sectionKey))

	// First: try h3 subsection across all h2
	for _, h2 := range h2Sections {
		h3Sections := parseSections(h2.Body, "###")
		if len(h3Sections) > 0 {
			if match := findSection(h3Sections, key); match != nil {
				return match.Body
			}
		}
	}

	// Second: try h2
	if match := findSection(h2Sections, key); match != nil {
		if len(match.Body) < largeSectionThreshold {
			return match.Body
		}
		h3Sections := parseSections(match.Body, "###")
		if len(h3Sections) > 0 {
			return buildSectionTOC(name, match.Slug, match.Header, h3Sections)
		}
		return match.Body
	}

	// Not found
	var available []string
	for _, h2 := range h2Sections {
		available = append(available, h2.Slug)
		for _, h3 := range parseSections(h2.Body, "###") {
			available = append(available, "  "+h3.Slug)
		}
	}
	return fmt.Sprintf("Section '%s' not found. Available sections:\n%s", sectionKey, strings.Join(available, "\n"))
}

// ListAvailableSkills returns names of available skills.
func ListAvailableSkills() []string {
	var names []string
	for name, entry := range SkillCatalog {
		if _, err := skillFS.ReadFile(entry.Filename); err == nil {
			names = append(names, name)
		}
	}
	return names
}
