package config

import (
	"fmt"
	"math/rand"
)

const (
	OpenRouterBase   = "https://fal.run/openrouter/router/openai/v1"
	DefaultModel     = "openai/gpt-5.4"
	DefaultTheme     = "tokyonight"
	InputPlaceholder = "Describe what to generate, or ask anything..."
)

var LogoGen = []string{
	" ██████   ████████  ██    ██",
	"██        ██        ███   ██",
	"██        ██        ████  ██",
	"██ ████   ██████    ██ ██ ██",
	"██   ██   ██        ██  ████",
	"██   ██   ██        ██   ███",
	" ██████   ████████  ██    ██",
}

var LogoMedia = []string{
	"██     ██  ████████  ████████   ██    ████",
	"███   ███  ██        ██    ██   ██   ██  ██",
	"████ ████  ██        ██    ██   ██  ██    ██",
	"██ ███ ██  ██████    ██    ██   ██  ████████",
	"██     ██  ██        ██    ██   ██  ██    ██",
	"██     ██  ██        ██    ██   ██  ██    ██",
	"██     ██  ████████  ████████   ██  ██    ██",
}

var LLMModels = []string{
	"openai/gpt-5.4",
	"google/gemini-3.1-pro-preview",
	"anthropic/claude-sonnet-4.6",
	"anthropic/claude-opus-4.6",
	"google/gemini-3-flash-preview",
}

type HelpCommand struct {
	Command     string
	Description string
	Keybind     string
}

var HelpCommands = []HelpCommand{
	{"/help", "show help", ""},
	{"/model", "switch model", ""},
	{"/theme", "switch theme", ""},
	{"/search", "search models", ""},
	{"/info", "model details", ""},
	{"/price", "check pricing", ""},
	{"/usage", "usage & cost", ""},
	{"/history", "request history", ""},
	{"/analytics", "performance stats", ""},
	{"/workflows", "list workflows", ""},
	{"/files", "list files", ""},
	{"/whoami", "auth identity", ""},
	{"/default", "set default models", ""},
	{"/login", "log in to fal.ai", ""},
	{"/logout", "log out of fal.ai", ""},
	{"/clear", "clear conversation", "ctrl+l"},
}

var Taglines = []string{
	"dreaming in pixels...",
	"connecting neurons...",
	"loading imagination...",
	"warming up GPUs...",
	"painting with math...",
	"brewing creativity...",
	"assembling photons...",
	"tuning latent space...",
	"summoning diffusion...",
	"composing reality...",
	"blending dimensions...",
	"rendering thoughts...",
	"distilling ideas...",
	"sculpting tensors...",
	"weaving embeddings...",
	"channeling inference...",
	"bootstrapping vision...",
	"calibrating aesthetics...",
	"sparking generation...",
	"igniting imagination...",
}

type ThemeConfig struct {
	Name           string
	Primary        string
	Secondary      string
	Accent         string
	Text           string
	TextMuted      string
	TextEmphasized string
	Bg             string
	BgSecondary    string
	BgDarker       string
	Border         string
	BorderFocused  string
	BorderDim      string
	Success        string
	Warning        string
	Error          string
	Info           string
}

// Theme is an alias for backwards compatibility.
type Theme = ThemeConfig

var Themes = map[string]Theme{
	"tokyonight": {
		Name:           "Tokyo Night",
		Primary:        "#82aaff",
		Secondary:      "#c099ff",
		Accent:         "#ff966c",
		Text:           "#c8d3f5",
		TextMuted:      "#636da6",
		TextEmphasized: "#ffc777",
		Bg:             "#222436",
		BgSecondary:    "#1e2030",
		BgDarker:       "#191b29",
		Border:         "#3b4261",
		BorderFocused:  "#82aaff",
		BorderDim:      "#2f334d",
		Success:        "#c3e88d",
		Warning:        "#ff966c",
		Error:          "#ff757f",
		Info:           "#82aaff",
	},
	"catppuccin": {
		Name:           "Catppuccin",
		Primary:        "#89b4fa",
		Secondary:      "#cba6f7",
		Accent:         "#fab387",
		Text:           "#cdd6f4",
		TextMuted:      "#6c7086",
		TextEmphasized: "#b4befe",
		Bg:             "#1e1e2e",
		BgSecondary:    "#181825",
		BgDarker:       "#11111b",
		Border:         "#313244",
		BorderFocused:  "#89b4fa",
		BorderDim:      "#45475a",
		Success:        "#a6e3a1",
		Warning:        "#fab387",
		Error:          "#f38ba8",
		Info:           "#89b4fa",
	},
	"nord": {
		Name:           "Nord",
		Primary:        "#88c0d0",
		Secondary:      "#b48ead",
		Accent:         "#ebcb8b",
		Text:           "#eceff4",
		TextMuted:      "#4c566a",
		TextEmphasized: "#d8dee9",
		Bg:             "#2e3440",
		BgSecondary:    "#3b4252",
		BgDarker:       "#242933",
		Border:         "#434c5e",
		BorderFocused:  "#88c0d0",
		BorderDim:      "#3b4252",
		Success:        "#a3be8c",
		Warning:        "#ebcb8b",
		Error:          "#bf616a",
		Info:           "#88c0d0",
	},
	"gruvbox": {
		Name:           "Gruvbox",
		Primary:        "#83a598",
		Secondary:      "#d3869b",
		Accent:         "#fe8019",
		Text:           "#ebdbb2",
		TextMuted:      "#665c54",
		TextEmphasized: "#fabd2f",
		Bg:             "#282828",
		BgSecondary:    "#1d2021",
		BgDarker:       "#141414",
		Border:         "#3c3836",
		BorderFocused:  "#83a598",
		BorderDim:      "#504945",
		Success:        "#b8bb26",
		Warning:        "#fe8019",
		Error:          "#fb4934",
		Info:           "#83a598",
	},
	"everforest": {
		Name:           "Everforest",
		Primary:        "#a7c080",
		Secondary:      "#d699b6",
		Accent:         "#e69875",
		Text:           "#d3c6aa",
		TextMuted:      "#859289",
		TextEmphasized: "#dbbc7f",
		Bg:             "#2d353b",
		BgSecondary:    "#272e33",
		BgDarker:       "#1e2326",
		Border:         "#475258",
		BorderFocused:  "#a7c080",
		BorderDim:      "#374145",
		Success:        "#a7c080",
		Warning:        "#e69875",
		Error:          "#e67e80",
		Info:           "#7fbbb3",
	},
}

const SystemPromptBase = `You are genmedia, the best AI creative engine on the planet. You are an interactive CLI tool that generates images, videos, audio, and other media on the fal.ai platform. You execute actions — you do not explain, teach, or write code.

IMPORTANT: Always respond in English regardless of what language the user writes in.

# Tone and style
- Be concise. Minimize output tokens while maintaining quality. No preamble or postamble.
- After generation completes, show the output URL and key details (dimensions, duration, time). Nothing else.
- One-liner acknowledgments are fine. Don't narrate your tool calls.
- Never output code snippets, SDK examples, or API usage. You are not a coding assistant. If the user shares code, extract the endpoint and parameters and call generate yourself.

# Tools
- search_models: Find models by query, category, or sort order
- model_info: Get model schema — input parameters, types, defaults, constraints, valid values
- generate: Submit a job to a model endpoint and wait for results
- ask_user: Show an interactive choice picker to the user (arrow keys + Enter). MANDATORY for any question.
- get_skill: Load domain knowledge progressively — start broad, go deeper only when needed:
  1. get_skill(name) → TOC + quick reference cheat sheet. Often enough for simple tasks.
  2. get_skill(name, section="key") → section details. Large sections show a sub-TOC instead.
  3. get_skill(name, section="subsection_key") → specific subsection (e.g. "veo_3_1", "dolly").
  The quick reference covers 80% of use cases. Only drill deeper for specific model params.
  Available skills:
  * cinematography — Hollywood techniques, camera moves, lighting, composition, lenses
  * video_prompting — AI video prompt engineering, model-specific tips, artifact prevention
  * image_prompting — AI image prompt engineering, style keywords, quality boosters
  * commercial — Advertising production, product shots, social media formats
  * audio_prompting — Music/audio generation, genres, moods, sound effects
  * character_design — Character creation & consistency, expressions, costumes, LoRA, IP-Adapter
  * storytelling — Multi-shot narrative, storyboarding, scene structure, pacing, transitions
  * social_media — Platform-optimized content (TikTok, Reels, Shorts), hooks, trends, viral formats
  * motion_design — Motion graphics, kinetic typography, logo reveals, transitions, particles
  * world_building — Environment design, architecture, interiors, fantasy/sci-fi worlds, atmosphere
  * brand_identity — Logo design, color systems, typography, brand guidelines, visual consistency
  * workflow_utils — fal.ai processing tools (resize, overlay, subtitle, merge, mask, etc.)
- get_pricing / estimate_cost: Check costs before expensive operations
- check_usage / get_analytics / request_history: Usage and performance data
- list_workflows / list_files / whoami: Platform utilities

# Doing tasks
- When the user asks to generate something, keep going until the generation is complete. Do not stop to explain or confirm — just execute.
- ALWAYS call model_info before generate. Read the input schema to build correct parameters. This is not optional. Skipping this causes 422 errors.
- If you don't know which model to use, call search_models first, pick the best match, then model_info, then generate.
- Use smart defaults: 16:9 for video, square for images, highest available quality. Don't ask the user about technical parameters unless they specifically requested custom settings.
- For domain expertise: call get_skill to get the TOC first, then load 1-2 specific sections. The quick reference in the TOC often has enough info for simple tasks.
- For complex pipelines (background removal, video editing, subtitles, etc.), load the workflow_utils skill to find the right utility endpoints.

# ask_user — interactive questions
- NEVER write questions or options as plain text. ALWAYS use the ask_user tool. The user sees a visual picker with arrow-key navigation.
- Use ask_user when the request is genuinely ambiguous: which model, which style, what subject, what creative direction.
- Do NOT use ask_user for technical defaults, error recovery, or confirming actions.

# Error recovery
When a tool call fails, diagnose and fix it yourself:
1. Read the error message carefully — it tells you exactly what's wrong.
2. If generate returns 422: call model_info to check correct parameter names, types, and valid enum values. Fix the input and retry. Do NOT ask the user about validation errors.
3. If a model is not found: search_models to find alternatives, pick the closest match, proceed.
4. If authentication fails: tell the user to run ` + "`fal auth login`" + `.
5. Never give up after one failure. Try at least 2 different approaches before reporting to the user.

# Working with user code
When the user pastes code containing fal endpoints (fal.subscribe, fal.run, fal.queue.submit):
1. Extract the endpoint_id and input parameters from the code.
2. Call model_info to verify the parameters against the schema.
3. Call generate with the corrected parameters. If any parameter doesn't match the schema, fix it silently based on model_info.`

func RandomTagline() string {
	return Taglines[rand.Intn(len(Taglines))]
}

func BuildSystemPrompt(prefs *Preferences) string {
	prompt := SystemPromptBase
	if prefs != nil {
		defaults := prefs.FormatForSystemPrompt()
		if defaults != "" {
			prompt += fmt.Sprintf("\n\n# User's default model preferences\n%s", defaults)
		}
	}
	return prompt
}

func GetTheme(name string) ThemeConfig {
	if t, ok := Themes[name]; ok {
		return t
	}
	return Themes[DefaultTheme]
}

