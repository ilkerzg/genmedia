<p align="center">
  <img src="https://v3b.fal.media/files/b/0a911fb0/5sW-6CopGrEuhioKnGDbd_image.png" width="720" />
</p>

<h1 align="center">genmedia</h1>

<p align="center">
  <img src="https://img.shields.io/badge/status-beta-orange" alt="beta" />
  <a href="https://fal.ai"><img src="https://img.shields.io/badge/powered%20by-fal.ai-blue" alt="powered by fal.ai" /></a>
  <img src="https://img.shields.io/github/license/ilkerzg/genmedia" alt="license" />
</p>

<p align="center">
  Open source AI media generation from the terminal.<br/>
  Powered by <a href="https://fal.ai">fal.ai</a>.
</p>

<p align="center">
  <a href="#install">Install</a> · <a href="#usage">Usage</a> · <a href="#commands">Commands</a> · <a href="#contributing">Contributing</a>
</p>

---

Generate images, videos, audio, and more — just describe what you want. genmedia picks the right model, handles the queue, and streams results back to your terminal.

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

Then just type:

```
> a cyberpunk city at sunset, neon reflections on wet streets
> generate a 10s video of ocean waves
> create lofi hip-hop background music
```

genmedia figures out which model to use, calls it, and renders the output — including inline image previews via [chafa](https://hpjansson.org/chafa/).

## Commands

| Command | Description |
|---------|-------------|
| `/model` | Switch LLM |
| `/search` | Search fal.ai models |
| `/info` | Model schema & details |
| `/price` | Check pricing |
| `/theme` | Switch color theme |
| `/default` | Set default models per category |
| `/usage` | Usage & cost |
| `/login` | Log in to fal.ai |
| `/clear` | Clear conversation |

**Themes:** Tokyo Night, Catppuccin, Nord, Gruvbox, Everforest

**Keyboard:** `Ctrl+L` clear, `Ctrl+C` quit, `PgUp/PgDn` scroll

## How it works

genmedia is a conversational TUI backed by an LLM with tool use. The LLM has access to fal.ai's full model catalog — it can search models, read schemas, submit generation jobs, and stream results. You talk to it naturally and it handles the rest.

Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) + [Lip Gloss](https://github.com/charmbracelet/lipgloss). Single static binary, zero dependencies, ~10ms startup.

## Contributing

This is an open source side project. Contributions, issues, and ideas are welcome.

```bash
git clone https://github.com/ilkerzg/genmedia
cd genmedia
go build ./cmd/genmedia
./genmedia
```

## License

MIT
