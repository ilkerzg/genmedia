<p align="center">
  <img src="https://v3b.fal.media/files/b/0a912003/kLRUpnF9PU8HzIHGx3ylP_image.png" width="720" />
</p>

<h1 align="center">genmedia</h1>

<p align="center">
  <img src="https://img.shields.io/badge/status-beta-orange" alt="beta" />
  <a href="https://fal.ai"><img src="https://img.shields.io/badge/powered%20by-fal.ai-blue" alt="powered by fal.ai" /></a>
  <img src="https://img.shields.io/github/license/ilkerzg/genmedia" alt="license" />
</p>

<p align="center">
  Open source AI media generation from the terminal.
</p>

<p align="center">
  <a href="#install">Install</a> · <a href="#usage">Usage</a> · <a href="#commands">Commands</a> · <a href="#contributing">Contributing</a>
</p>

---

Generate images, videos, audio, and more. Just describe what you want. genmedia picks the right model, handles the queue, and streams results back to your terminal.

## Install

```bash
# Homebrew
brew install ilkerzg/tap/genmedia

# Go
go install github.com/ilkerzg/genmedia/cmd/genmedia@latest

# npx (no install)
npx genmedia
```

## Auth

```bash
# Option 1: Browser login
genmedia
/login

# Option 2: API key (from fal.ai/dashboard/keys)
export FAL_KEY="your-key-here"
genmedia
```

## Usage

```
genmedia                  # start a new session
genmedia -s last          # resume last session
genmedia -m anthropic/claude-sonnet-4.6  # use a specific LLM
```

Then just type what you want:

```
> a cyberpunk city at sunset, neon reflections on wet streets
> generate a 10s video of ocean waves crashing on rocks
> create lofi hip-hop background music, chill vibes
> remove the background from this image and make it transparent
> upscale my photo to 4K
> add cinematic subtitles to my video
> search for the best anime style image model
> how much does flux pro cost per image?
```

genmedia figures out which model to use, calls it, and renders the output inline via [chafa](https://hpjansson.org/chafa/).

## Commands

| Command | Description |
|---------|-------------|
| `/model` | Switch LLM |
| `/search` | Search fal.ai models |
| `/info` | Model schema and details |
| `/price` | Check pricing |
| `/theme` | Switch color theme |
| `/default` | Set default models per category |
| `/usage` | Usage and cost |
| `/login` | Log in to fal.ai |
| `/clear` | Clear conversation |

**Themes:** Tokyo Night, Catppuccin, Nord, Gruvbox, Everforest

**Keyboard:** `Ctrl+L` clear, `Ctrl+C` quit, `PgUp/PgDn` scroll

## How it works

genmedia is a conversational TUI backed by an LLM with tool use. The LLM has access to fal.ai's full model catalog. It can search models, read schemas, submit generation jobs, and stream results. You talk to it naturally and it handles the rest.

Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lip Gloss](https://github.com/charmbracelet/lipgloss). Single static binary, zero dependencies, ~10ms startup.

## Contributing

Open source side project. Contributions, issues, and ideas are welcome.

```bash
git clone https://github.com/ilkerzg/genmedia
cd genmedia
go build ./cmd/genmedia
./genmedia
```

## License

MIT
