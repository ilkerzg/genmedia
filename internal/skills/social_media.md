# Social Media — Platform-Optimized Content Creation Master Guide

Comprehensive reference for generating platform-optimized social media content with fal.ai models (2026).
Every specification, safe zone, prompt, and strategy in this document is verified and actionable.

---

## Quick Reference

**Platform Specs**

| Platform | Resolution | Aspect | Sweet Spot Duration |
|----------|-----------|--------|---------------------|
| TikTok | 1080x1920 | 9:16 | 15-60 seconds |
| Instagram Reels | 1080x1920 | 9:16 | 15-30 seconds |
| Instagram Feed | 1080x1350 | 4:5 | — (image) |
| YouTube Shorts | 1080x1920 | 9:16 | 30-60 seconds |
| YouTube Standard | 1920x1080 | 16:9 | 8-15 minutes |
| LinkedIn Video | 1920x1080 | 16:9 | Under 90 seconds |

> **Model Discovery:** Use `search_models` to find the best model for each task. Search by category: "text to video" for video clips, "image generation" for static posts, "text to image with text" for graphics with readable text, "text to speech" for voiceover, "music generation" for background music.

**Top 5 Social Media Content Tips**

1. Hook in 1-3 seconds — the algorithm tests on 300-500 viewers first; poor watch time kills reach.
2. Design for muted viewing — captions and text overlays are mandatory, not optional.
3. Saves beat likes — create reference-worthy, save-worthy content (lists, tutorials, templates).
4. Place all key visuals in the safe zone — UI overlays block the top ~150px and bottom ~270px on TikTok.
5. Post consistently over perfection — one average post daily outperforms one perfect post weekly.

**Key Prompt Patterns**

```
# POV / immersive
POV [action], [environment], handheld first-person perspective, [lighting], 9:16 vertical

# Satisfying process (high replay)
Extreme close-up of [hands/object] [doing tactile action], [material detail], shallow DOF, 9:16 vertical

# Surreal / pattern interrupt
[Ordinary subject in wrong context], cinematic handheld, documentary style, 9:16 vertical

# Loop-friendly (TikTok replays)
[Subject action], one complete cycle, ending where it began, smooth continuous motion, loop-friendly, 9:16 vertical
```

**Cheat Sheet — Key Prompt Words by Goal**

| Goal | Key Prompt Words |
|------|-----------------|
| Viral TikTok/Reel clip | cinematic, 9:16, handheld, hook opener |
| Human close-up / UGC | selfie angle, authentic, shallow DOF |
| Instagram carousel slide | clean layout, 1080x1350, bold text |
| Title card / text graphic | bold white sans-serif, centered, 1080x1920 |
| Product photo | commercial photography, key light, reflective surface |

---

## Platform Specifications

### TikTok

| Property | Value |
|----------|-------|
| Resolution | 1080x1920 (9:16) |
| Aspect ratio | 9:16 |
| Max duration | 10 minutes |
| Sweet spot duration | 15-60 seconds |
| File size limit | 287 MB |
| Autoplay | Yes, sound ON by default |

**Safe zones (critical for AI-generated content):**
- **Top 150px:** Username, follow button, and live indicator overlay here. Never place key visuals or text in this zone.
- **Bottom 270px:** Description text, hashtags, music ticker, and interaction buttons (like/comment/share/bookmark) overlay here. Keep subject faces and important text above this line.
- **Right 100px:** Like, comment, share, and bookmark icons stack vertically along the right edge.
- **True safe area:** The center rectangle from x:100 to x:980, y:150 to y:1650 is the only guaranteed clean zone.

**Algorithm signals (ranked by weight):**
1. Watch time / completion rate (strongest signal)
2. Replays (indicates compelling content)
3. Shares (especially to DMs)
4. Comments (volume and speed)
5. Likes (weakest signal but still counts)
6. Profile visits after viewing

**Generation notes:**
- TikTok compresses heavily. Generate at 1080x1920 minimum; upscaling from 720p looks noticeably soft.
- Sound-on by default means generated audio, music, or voiceover adds significant value.
- The algorithm tests content with small audiences first (300-500 views), then escalates. Hook quality in the first 1-3 seconds determines whether it scales.

---

### Instagram Reels

| Property | Value |
|----------|-------|
| Resolution | 1080x1920 (9:16) |
| Aspect ratio | 9:16 |
| Max duration | 180 seconds (3 min) |
| Sweet spot duration | 15-30 seconds |
| Thumbnail crop | Center 1080x1080 square shown in grid |

**Safe zones:**
- **Top 120px:** Camera, close-friends, and music controls.
- **Bottom 300px:** Caption, audio info, username, and action buttons. Taller overlay than TikTok.
- **Right 80px:** Like, comment, share, save, remix, and three-dot menu.
- **True safe area:** x:80 to x:1000, y:120 to y:1620.

**Algorithm signals (ranked by weight):**
1. Saves (strongest — indicates reference-worthy content)
2. Shares (especially to Stories or DMs)
3. Comments (length matters more than count)
4. Likes
5. Watch time / replays
6. Profile visits

**Generation notes:**
- Reels thumbnail defaults to a frame from the video. Aesthetic consistency with the profile grid matters — consider generating a custom thumbnail image separately.
- Instagram applies heavier color grading compression than TikTok. Oversaturate colors slightly (10-15%) in prompts to compensate.
- Trending audio is a major discovery lever. Pair generated video with trending audio using the music selection guidance below.

---

### Instagram Feed Post

| Format | Resolution | Aspect Ratio |
|--------|-----------|--------------|
| Square | 1080x1080 | 1:1 |
| Portrait (recommended) | 1080x1350 | 4:5 |
| Landscape | 1080x608 | 1.91:1 |
| Carousel | Up to 20 slides, any ratio above (must be consistent) |

**Algorithm signals:**
1. Saves (reference value)
2. Shares
3. Dwell time (how long someone looks at the post)
4. Comments
5. Likes

**Carousel best practices:**
- Slide 1 is the hook — it must stop the scroll. Use bold text, striking imagery, or a provocative question.
- Maintain consistent visual style across all slides (same color palette, typography, layout).
- End slide should be a clear CTA: "Save this," "Share with a friend," "Follow for more."
- Educational carousels (tips, tutorials, guides) are the highest-save format on Instagram.
- Use `search_models` to find a design-focused model for slides with text/graphics and a photorealistic image generation model for photographic slides.

---

### Instagram Stories

| Property | Value |
|----------|-------|
| Resolution | 1080x1920 (9:16) |
| Duration per story | Up to 60 seconds (auto-splits at 15s segments) |
| Sweet spot | 5-15 seconds per story |
| Lifespan | 24 hours (unless saved as Highlight) |

**Tap zones:**
- **Left 33%:** Tap to go back to previous story.
- **Center 34%:** Hold to pause the story.
- **Right 33%:** Tap to advance to next story.

**Sticker interaction zones:**
- Place poll stickers, question boxes, and sliders in the center vertical third for maximum interaction.
- Avoid placing interactive elements in the top 120px (close button, profile info) or bottom 100px (message bar, send-to).

---

### YouTube Shorts

| Property | Value |
|----------|-------|
| Resolution | 1080x1920 (9:16) |
| Max duration | 180 seconds (3 min) |
| Sweet spot duration | 30-60 seconds |
| Thumbnail | Auto-selected frame (custom thumbnail now available) |

**Safe zones:**
- **Bottom 200px:** Subscribe button, title, and channel info overlay here.
- **Right 80px:** Like, dislike, comment, share, and remix buttons.
- **Top 80px:** Search and camera icons.

**Algorithm signals (ranked):**
1. Click-through rate (thumbnail + title)
2. Watch time and completion rate
3. Engagement (likes, comments, shares)
4. Subscriber conversion (new subs from the Short)

**Generation notes:**
- YouTube Shorts are the strongest discovery tool for growing a YouTube channel. Optimize for subscriber conversion by including channel-relevant educational content.
- YouTube does NOT compress as aggressively as TikTok/Instagram. Higher quality source material shows a visible difference here.
- Shorts can drive traffic to long-form videos. Consider generating Shorts as teasers for full 16:9 content.

---

### YouTube Standard (Long-Form)

| Property | Value |
|----------|-------|
| Resolution | 1920x1080 (16:9), 4K preferred |
| Thumbnail | 1280x720 (required custom upload) |
| Aspect ratio | 16:9 |
| Sweet spot duration | 8-15 minutes |

**Thumbnail generation:**
- Use `search_models` to find a photorealistic image generation model for thumbnails, or a text-rendering-capable model if the thumbnail needs readable text.
- High contrast, saturated colors, expressive faces, and 3-5 large words perform best.
- Thumbnail + title is the single biggest factor in YouTube performance. Spend disproportionate effort here.

**Algorithm signals:**
1. Click-through rate (thumbnail + title)
2. Average view duration
3. Session time (do viewers keep watching YouTube after your video?)
4. Engagement velocity (likes/comments in first hour)

---

### LinkedIn

| Format | Resolution |
|--------|-----------|
| Video (landscape) | 1920x1080 (16:9) |
| Video (portrait) | 1080x1920 (9:16) |
| Feed image | 1080x1080 (1:1) or 1200x628 (1.91:1 for links) |
| Carousel (PDF) | 1080x1080 or 1080x1350 |

**Content guidelines:**
- Professional, B2B-oriented tone. Avoid overt entertainment or meme formats.
- Educational and thought-leadership content outperforms promotional content by 3-5x.
- Native video (uploaded directly) gets 5x more reach than YouTube links.
- Text posts with a single strong image often outperform video on LinkedIn.
- Use `search_models` to find a design-focused image generation model for professional infographics and data visualizations.

---

### Twitter / X

| Format | Resolution | Notes |
|--------|-----------|-------|
| Video | 1920x1080 or 1080x1920 | Max 2 min 20 sec |
| Single image | 1600x900 (16:9) | Displays full in timeline |
| Two images | 700x800 each | Side by side |
| Four images | 700x800 each | 2x2 grid |
| GIF | 1280x720 recommended | Autoplay, loops |

**Critical:** Twitter/X autoplays video MUTED. Captions or text overlays are mandatory for video content. Without them, scroll-past rate exceeds 90%.

**Algorithm signals:**
1. Replies and quote tweets
2. Retweets
3. Bookmark saves
4. Likes
5. Dwell time on media

---

## Content Hooks — First 1-3 Seconds

The hook is the single most important element in short-form content. It determines whether viewers watch or scroll. Every AI-generated video must open with a deliberate hook.

### Visual Hook Types

| Hook Type | Description | Best For |
|-----------|-------------|----------|
| Pattern interrupt | Unexpected, jarring, or surreal visual that breaks the scroll pattern | Entertainment, viral reach |
| Before/after reveal | Split screen or transition showing transformation | Product, tutorial, fitness |
| Satisfying process | Pouring, mixing, assembling, unwrapping — tactile and smooth | Product, ASMR, food |
| Bold text statement | Large text on screen with provocative claim or question | Educational, opinion |
| Face close-up | Extreme close-up of a face showing emotion | Storytelling, reaction |
| Action from frame 1 | Movement, impact, or physical action starting immediately | Sports, action, comedy |
| Mystery/curiosity gap | Partial reveal that demands the viewer watch to understand | All categories |
| Countdown/list | "5 things you didn't know about..." opening frame | Educational, tips |

### Hook Prompt Templates

**Pattern interrupt (surreal visual):**
```
A goldfish swimming through the air in a busy New York subway car, commuters
ignoring it completely, cinematic handheld camera, shallow depth of field,
warm fluorescent lighting, documentary style, 9:16 vertical format
```
**Model:** Use `search_models` — text-to-video.

**Before/after transformation:**
```
Split screen transition: left side shows a bare, empty concrete room, camera
slowly pans right crossing a sharp vertical dividing line to reveal the same
room fully designed with modern furniture, plants, warm lighting, and art on
walls, smooth lateral dolly, 9:16 vertical, interior design reveal
```
**Model:** Use `search_models` — text-to-video.

**Satisfying process:**
```
Extreme close-up of thick honey being poured in slow motion over a stack of
golden pancakes, viscous liquid catching warm morning sidelight, steam rising
from the pancakes, shallow DOF, food commercial quality, 9:16 vertical format
```
**Model:** Use `search_models` — text-to-video.

**Bold text opener (generate as image, then animate or use as first frame):**
```
Bold white sans-serif text "STOP SCROLLING" centered on pure black background,
slight 3D extrusion effect, dramatic red accent underline, clean graphic
design, 1080x1920 vertical
```
**Model:** Use `search_models` — text-to-image with strong text rendering.

**Face close-up with emotion:**
```
Extreme close-up of a young woman's face showing genuine surprise, mouth
slightly open, eyes wide, soft natural window light from camera-left, shallow
depth of field, slight handheld movement, documentary feel, 9:16 vertical
```
**Model:** Use `search_models` — text-to-video (look for strong human face rendering).

**Curiosity gap opener:**
```
A mysterious locked wooden chest sitting on a table, camera slowly pushing in,
dramatic overhead spotlight creating pool of light around the chest, dust
particles in the air, dark moody background, anticipation building, 9:16
```
**Model:** Use `search_models` — text-to-video.

---

## Viral Content Formats (2026)

### Format Reference Table

| Format | Duration | Platform | Model Category (use `search_models`) |
|--------|----------|----------|---------------------------------------|
| POV video | 15-30s | TikTok, Reels | text-to-video |
| Day in the life | 30-60s | TikTok, Reels, Shorts | text-to-video |
| Process/creation reveal | 15-45s | Reels, TikTok | text-to-video (good hand/object detail) |
| Transition/transformation | 10-20s | TikTok, Reels | text-to-video |
| Mini documentary | 60-180s | Shorts, Reels | text-to-video (long form) |
| Get ready with me | 30-60s | TikTok, Reels | text-to-video (good human rendering) |
| Tutorial/how-to | 30-90s | Shorts, Reels, TikTok | text-to-video |
| Comparison/versus | 15-30s | TikTok, Reels | image generation (stills), text-to-video (clips) |
| Behind the scenes | 15-60s | Reels, Stories, TikTok | text-to-video |
| Aesthetic/mood content | 10-30s | Reels, TikTok | text-to-video |
| Product showcase | 15-30s | Reels, TikTok, Shorts | text-to-video (good object detail) |
| Reaction content | 10-30s | TikTok, Shorts | text-to-video (good human rendering) |

### POV Videos

First-person perspective content. The camera IS the viewer.

**Prompt pattern:**
```
POV [action], [environment description], [what hands/body are doing],
[lighting], handheld first-person perspective, [mood], 9:16 vertical
```

**Example:**
```
POV walking through a neon-lit Tokyo alley at night, camera at eye level moving
forward, rain-soaked pavement reflecting pink and blue neon signs, slight
handheld sway, passing ramen shops and vending machines, cinematic atmosphere,
9:16 vertical format
```
**Model:** Use `search_models` — text-to-video (look for smooth forward movement and environmental detail).

### Process / Creation Reveal

Show something being made, assembled, or created. High save rate.

**Prompt pattern:**
```
[Overhead/close angle] of [hands/tools] [creating action], [materials],
[workspace setting], [lighting], satisfying process video, [speed note], 9:16
```

**Example:**
```
Overhead close-up of skilled hands assembling a miniature wooden house model,
placing tiny furniture inside with tweezers, warm desk lamp lighting, organized
craft workspace with tools, steady tripod shot, satisfying craft process,
9:16 vertical format
```
**Model:** Use `search_models` — text-to-video (look for detailed hand/object interactions).

### Transition / Transformation

Content that uses a visual transition to show change. High replay and share rate.

**Prompt pattern (generate as two separate clips, then combine):**

Clip 1:
```
[Subject in state A], [action moving toward camera / covering lens], 9:16
```

Clip 2:
```
[Subject in state B], [action moving away from camera / revealing], 9:16
```

**Example pair:**
```
Clip 1: A woman in casual pajamas walking toward the camera in a messy bedroom,
reaching her hand toward the lens to cover it, natural morning light, handheld,
9:16 vertical

Clip 2: The same woman in an elegant evening gown stepping back from the camera
in a glamorous ballroom, revealing the full scene, warm golden chandelier
lighting, steady shot, 9:16 vertical
```
**Model:** Use `search_models` — text-to-video (look for strong human consistency across clips).

---

## Platform-Specific Content Strategy

### TikTok Strategy

**Pacing:** Fast. Cut or change the visual every 2-3 seconds. Static shots longer than 4 seconds cause drop-off.

**Text overlays:** Essential. TikTok viewers expect text on screen reinforcing or adding to the visual. Use `search_models` to find a text-to-image model with strong text rendering for generating text graphics, then composite.

**Loop-friendly content:** The algorithm rewards replays. Design content where the end visually connects to the beginning:
```
A ceramic vase spinning on a potter's wheel, one complete rotation, hands
gently shaping the clay, overhead angle, warm studio lighting, smooth continuous
motion ending where it began, loop-friendly, 9:16 vertical, 5 seconds
```

**Comment-bait techniques:**
- Include a deliberate "mistake" or debatable choice that viewers will comment about.
- Ask a question in the text overlay: "Would you try this?"
- Show an incomplete process and ask viewers to guess the outcome.
- Place a subtle detail that rewards attentive viewers.

**Optimal posting cadence:** 1-3 videos per day for growth. Consistency matters more than individual post quality.

### Instagram Strategy

**Aesthetic consistency:** Instagram rewards cohesive visual identity. When generating multiple pieces:
- Lock down a color palette (specify exact tones in prompts).
- Maintain consistent lighting style across posts.
- Use the same style anchor phrase across related content.

**Carousel storytelling structure:**
1. **Slide 1 — Hook:** Bold statement, question, or striking image. This is the thumbnail.
2. **Slides 2-8 — Value:** Deliver on the hook. Tips, steps, information, or narrative progression.
3. **Final slide — CTA:** "Save this for later," "Share with someone who needs this," "Follow @handle for more."

**Reels strategy:**
- Pair AI-generated video with trending audio for discovery.
- Use text overlays heavily (Reels are often watched with sound off despite autoplay).
- Post Reels during high-engagement windows (typically 11am-1pm and 7-9pm local).
- Cross-post TikTok content but remove the TikTok watermark (Instagram deprioritizes watermarked content).

### YouTube Shorts Strategy

**Educational hooks dominate Shorts:**
```
"Did you know that..." / "Most people don't realize..." / "Here's what nobody tells you about..."
```

**Subscribe CTA:** Place a verbal or visual CTA in the last 3-5 seconds. Shorts are the fastest way to grow subscribers in 2026.

**Series content:** Create recurring formats (e.g., "Design tip #47") that viewers follow. YouTube recommends more of the same series once a viewer engages with one.

**Searchable topics:** Unlike TikTok/Reels, YouTube Shorts are heavily search-driven. Use descriptive titles and descriptions with keywords viewers actually search for.

### LinkedIn Strategy

**Content that performs on LinkedIn:**
- Industry insights and data visualization (use `search_models` to find a design-focused image generation model for charts and infographics).
- Before/after case studies with visual proof.
- Professional "day in the life" content.
- Thought-leadership with AI-generated supporting visuals.

**Video on LinkedIn:**
- Keep under 90 seconds for feed video.
- Add captions (autoplay is muted).
- Landscape 16:9 performs better than vertical on LinkedIn desktop (majority of LinkedIn usage).
- Professional, polished look. Avoid TikTok-style effects and transitions.

---

## Content Types and Ready-to-Use Prompts

### Product Showcase (Vertical 9:16)

**Hero product rotation:**
```
Elegant slow-motion rotation of a matte black wireless earbud case on a
reflective dark surface, single key light from above creating clean specular
highlights, subtle particle dust in the air catching the light, premium product
commercial, 9:16 vertical, dark moody background
```
**Model:** Use `search_models` — text-to-video (look for good object consistency during rotation).

**Product in context:**
```
A hand picking up a sleek thermos from a wooden desk in a modern home office,
warm afternoon sunlight streaming through large windows, camera follows the hand
with a subtle rack focus, lifestyle product placement, 9:16 vertical format
```
**Model:** Use `search_models` — text-to-video.

**Product flat-lay (image for carousel):**
```
Overhead flat-lay product photography of skincare collection: three glass
bottles and two cream jars arranged symmetrically on pale marble surface, fresh
eucalyptus sprigs and water droplets as props, soft diffused lighting, clean
minimalist aesthetic, 1080x1350 portrait format
```
**Model:** Use `search_models` — image generation.

### Lifestyle / Aspirational

**Travel aspirational:**
```
A woman in a flowing white linen dress walking along the edge of an infinity
pool overlooking the Amalfi Coast at golden hour, gentle breeze moving her hair
and dress, turquoise Mediterranean Sea in the background, warm cinematic color
grade, slow motion, aspirational travel content, 9:16 vertical
```
**Model:** Use `search_models` — text-to-video.

**Morning routine aesthetic:**
```
Close-up of hands pouring oat milk into a ceramic cup of matcha latte, creating
a swirling green and white pattern, steam rising, morning sunlight on a clean
white marble countertop, ASMR-style satisfying pour, warm tones, 9:16 vertical
```
**Model:** Use `search_models` — text-to-video.

**Fitness / wellness:**
```
A silhouette of a person doing yoga tree pose on a rooftop at sunrise, city
skyline in the background, golden hour light creating a warm glow around the
figure, slight slow motion, peaceful and aspirational, 9:16 vertical format
```
**Model:** Use `search_models` — text-to-video.

### Educational / Tutorial

**Step-by-step visual (image carousel):**
```
Clean instructional diagram showing Step 3 of a coffee brewing guide: a hand
pouring hot water in circular motion over a pour-over dripper, arrows showing
the circular pattern, minimal flat illustration style, warm brown and cream
color palette, bold "Step 3: The Pour" text at top, 1080x1350
```
**Model:** Use `search_models` — design-focused image generation with text rendering (best for graphics with text and diagrams).

**Tutorial video:**
```
Overhead close-up of a hand demonstrating calligraphy, writing the word
"breathe" in flowing cursive with a brush pen on cream paper, ink flowing
smoothly, steady tripod shot, warm desk lamp lighting, tutorial video style,
9:16 vertical format
```
**Model:** Use `search_models` — text-to-video.

### Entertainment / Humor

**Surreal comedy:**
```
A golden retriever sitting at a desk in a corporate office, wearing a tiny
necktie, looking directly at the camera with a deadpan expression, coworkers
walking past in the background unfazed, fluorescent office lighting, mockumentary
style handheld camera, 9:16 vertical
```
**Model:** Use `search_models` — text-to-video (handles surreal scenarios with photographic realism).

**Dramatic mundane:**
```
Extreme slow-motion of a person dramatically biting into a crispy piece of
toast, crumbs exploding outward catching golden morning sidelight, shallow depth
of field, epic cinematic score feeling, shot like a high-budget food commercial,
9:16 vertical format
```
**Model:** Use `search_models` — text-to-video.

### Behind the Scenes

**Creative process:**
```
A ceramic artist's hands covered in wet clay shaping a bowl on a spinning
wheel, camera slowly pulling back to reveal the full messy workshop with
shelves of finished pottery, natural window light, authentic behind-the-scenes
documentary feel, 9:16 vertical
```
**Model:** Use `search_models` — text-to-video.

### User-Generated Content (UGC) Style

UGC-style content intentionally looks raw and authentic, not polished. It outperforms commercial-looking content on TikTok and Reels because viewers trust it more.

**Key prompt descriptors for UGC feel:** `handheld camera`, `slightly shaky`, `natural lighting`, `casual framing`, `front-facing camera`, `selfie angle`, `authentic`, `unpolished`, `home setting`

**UGC product review style:**
```
Front-facing selfie-angle video of a young woman holding up a small skincare
bottle toward the camera in her bathroom, natural overhead bathroom lighting,
slightly messy counter in background, authentic unpolished feel, casual handheld
movement, genuine excitement on her face, 9:16 vertical
```
**Model:** Use `search_models` — text-to-video.

### Aesthetic / Mood Board

**Dark academia:**
```
A leather-bound book lying open on a mahogany desk in a dim library, candlelight
flickering and casting dancing shadows on the pages, dust particles floating in
the warm light beam from a high window, vintage dark academia aesthetic,
moody warm tones, 9:16 vertical
```
**Model:** Use `search_models` — text-to-video.

**Coastal grandmother:**
```
Close-up of linen napkins fluttering gently on a clothesline overlooking a
calm Mediterranean sea, soft pastel blue sky, warm golden afternoon light,
gentle breeze, coastal aesthetic, neutral and cream tones, peaceful, 9:16
```
**Model:** Use `search_models` — text-to-video.

---

## Text Overlay Best Practices

Text overlays are mandatory for social media video. Most platforms autoplay muted on scroll, and even sound-on platforms (TikTok) show better retention with reinforcing text.

### Text Design Rules

| Property | Recommendation |
|----------|---------------|
| Font | Bold sans-serif (Montserrat, Inter, Futura). Never thin weights. |
| Color | White with black stroke/outline (2-3px) for universal readability |
| Position | Center or upper third. Never bottom 25% (UI overlays). |
| Size | Large enough to read on a phone held at arm's length |
| Duration | 2-3 words per second. Hold each line 1.5-3 seconds. |
| Animation | Fade-in or pop-in. Avoid complex animations that distract. |
| Style | Word-by-word highlight (karaoke style) for voiceover sync |

### Generating Text Graphics

For text that must be pixel-perfect and readable, generate as separate image overlays:

```
Bold white text "THIS CHANGES EVERYTHING" on transparent background, thick
black outline, modern sans-serif font, centered, 1080x300 banner format
```
**Model:** Use `search_models` — text-to-image with strong text rendering (highest accuracy for readable text).

For stylized text that is part of the visual design:
```
Neon sign text "OPEN LATE" glowing pink against a dark brick wall, realistic
neon tube letter style, slight bloom and haze, moody nighttime atmosphere
```
**Model:** Use `search_models` — text-to-image with strong text rendering.

---

## Audio Strategy

### Music Selection by Content Type

| Content Type | Genre/BPM | Mood |
|-------------|-----------|------|
| Product showcase | Lo-fi hip hop, 70-90 BPM | Chill, aspirational |
| Tutorial/educational | Upbeat acoustic, 100-120 BPM | Friendly, approachable |
| Fitness/action | EDM/trap, 130-150 BPM | High energy, motivating |
| Lifestyle/aesthetic | Ambient/indie, 80-100 BPM | Dreamy, atmospheric |
| Comedy/entertainment | Quirky/playful, 110-130 BPM | Fun, lighthearted |
| Luxury/premium | Orchestral/minimal piano, 60-80 BPM | Elegant, sophisticated |
| Behind the scenes | Acoustic/folk, 90-110 BPM | Authentic, warm |

**Generating background music:**
```
Model: Use search_models — music generation
Prompt: "Upbeat lo-fi hip hop instrumental with warm vinyl crackle, soft Rhodes
piano chords, mellow drum pattern, relaxed cafe atmosphere, 85 BPM"
```

```
Model: Use search_models — audio generation
Prompt: "Cinematic ambient electronic music, slow build, ethereal synth pads,
subtle bass pulse, dreamy and aspirational mood, 90 BPM, 30 seconds"
```

### BPM Matching for Cuts

Sync visual cuts to the beat for professional feel. Common approach:
- At 120 BPM, one beat = 0.5 seconds. Cut every 1-2 beats (0.5-1.0 seconds) for fast content.
- At 90 BPM, one beat = 0.67 seconds. Cut every 2-3 beats (1.3-2.0 seconds) for medium pace.
- At 60 BPM, one beat = 1.0 second. Cut every 2-4 beats (2-4 seconds) for slow, cinematic content.

### Sound Effects

Add subtle SFX to elevate content quality:
```
Model: Use search_models — sound effects generation
Prompt: "Satisfying swoosh transition sound effect, clean and modern, 1 second"
```

```
Model: Use search_models — sound effects generation
Prompt: "Soft camera shutter click, DSLR, crisp and clean"
```

### Voiceover

**TikTok/Reels narration (conversational):**
```
Model: Use search_models — text-to-speech
Voice: Select a casual, young, energetic voice
Text: Keep sentences short. Pause between thoughts. Sound like you're talking
to a friend, not reading a script.
```

**YouTube/LinkedIn narration (professional):**
```
Model: Use search_models — text-to-speech
Voice: Select a warm, authoritative, clear voice
Text: Slightly longer sentences are fine. Maintain confident pacing. Enunciate
clearly.
```

---

## Engagement Optimization

### Call-to-Action (CTA) Placement

| CTA Type | Best Placement | Platform |
|----------|---------------|----------|
| "Follow for more" | Last 3 seconds | TikTok, Reels |
| "Save this" | Last slide or last 5 seconds | Instagram (high save rate) |
| "Subscribe" | Last 5 seconds | YouTube Shorts |
| "Comment below" | Middle of content (after value delivery) | All platforms |
| "Share with someone" | End | Instagram, TikTok |
| "Link in bio" | End, with pointing gesture/arrow | Instagram, TikTok |

### Comment-Bait Techniques

1. **Polarizing opinion:** "I think X is better than Y. Fight me in the comments." Generates debate.
2. **Incomplete list:** "Here are 4 of the 5 best tips... can you guess #5?" Forces engagement.
3. **Spot the difference:** Include a subtle intentional detail and ask viewers to find it.
4. **This or that:** "Which do you prefer: A or B?" Simple binary choice drives comments.
5. **Wrong answer only:** "Tell me what this looks like. Wrong answers only." Comedy engagement.

### Save-Worthy Content Triggers

Content that gets saved (highest-value engagement on Instagram):
- Step-by-step tutorials and how-to guides
- Reference lists ("10 camera angles every creator should know")
- Templates and frameworks
- Travel guides and restaurant recommendations
- Recipe carousels
- Outfit inspiration grids
- Tool recommendations

### Share Triggers

Content that gets shared:
- **Relatable:** "This is so me" — generates tag-a-friend sharing
- **Surprising:** Unexpected facts or visuals — generates "you need to see this" sharing
- **Useful:** Practical tips — generates "this will help you" sharing
- **Emotional:** Content that moves people — generates empathetic sharing
- **Identity:** Content that says something about the sharer's values or taste

---

## Batch Content Workflow

Generating 10-20 content pieces efficiently using fal.ai models.

### Step 1: Define the Content Pillar

Choose one theme or topic (e.g., "morning coffee aesthetics" or "productivity tips").

### Step 2: Generate Base Assets

**Image batch (for carousels, thumbnails, static posts):**
Use `search_models` to find a photorealistic image generation model for photographic content or a design-focused model for graphic/text content. Generate 10-15 variations by adjusting:
- Lighting (morning, golden hour, blue hour, studio)
- Angle (overhead, 45-degree, eye-level, low angle)
- Color palette (warm, cool, neutral, vibrant)
- Composition (centered, rule of thirds, negative space left, negative space right)

### Step 3: Generate Video Content

Pick the 3-5 strongest images and animate them using image-to-video:
Use `search_models` to find an image-to-video model for cinematic quality or a text-to-video model for human subjects.

### Step 4: Generate Audio

Create 2-3 music tracks and select the best fit:
Use `search_models` to find a music generation model for full tracks or an audio generation model for ambient/background.

Generate voiceover if needed:
Use `search_models` to find a text-to-speech model for narration.

### Step 5: Variations and Edits

Use editing models to create variations of top-performing assets:
Use `search_models` to find an image editing model for scene-aware edits (change background, swap props, adjust colors) or an inpainting model for targeted edits.

### Step 6: Platform Adaptation

From a single hero asset, create platform-specific versions:
- **9:16 (TikTok, Reels, Shorts):** Primary vertical content
- **1:1 (Instagram Feed, Twitter):** Center-crop or regenerate as square
- **4:5 (Instagram Feed):** Slight crop from 9:16 or regenerate
- **16:9 (YouTube, LinkedIn, Twitter):** Regenerate horizontal or letterbox the vertical
- **1280x720 (YouTube Thumbnail):** Generate separately with an image generation model + a text-rendering model for text overlay (use `search_models`)

---

## Complete Platform-Ready Examples

### Example 1: TikTok Coffee Shop Aesthetic

**Video (hook clip):**
```
Extreme close-up slow-motion of espresso pouring from a bottomless portafilter
into a clear glass, rich crema forming, warm cafe lighting, shallow depth of
field, steam rising, satisfying coffee process, 9:16 vertical, 5 seconds
```
**Model:** Use `search_models` — text-to-video.

**Music:**
```
Warm lo-fi hip hop beat with vinyl crackle, soft jazz piano, gentle brush
drums, coffee shop ambiance, 80 BPM, 30 seconds
```
**Model:** Use `search_models` — music generation.

---

### Example 2: Instagram Carousel — Design Tips

**Slide 1 (hook):**
```
Bold modern graphic design: large white text "5 DESIGN RULES" on deep navy
background, subtle gold geometric accents, clean sans-serif typography,
minimalist, 1080x1350
```
**Model:** Use `search_models` — text-to-image with strong text rendering.

**Slides 2-5 (content — generate each separately):**
```
Clean infographic layout, white background, left side shows [design principle]
illustrated with simple shapes, right side shows bold text "Rule [N]: [Title]"
with 2-line explanation below, consistent navy and gold color scheme, 1080x1350
```
**Model:** Use `search_models` — design-focused image generation.

**Slide 6 (CTA):**
```
Minimal graphic: white text "SAVE THIS FOR LATER" centered on navy background,
small bookmark icon above text, subtle gold border, 1080x1350
```
**Model:** Use `search_models` — text-to-image with strong text rendering.

---

### Example 3: YouTube Shorts — Tech Review

**Video:**
```
Close-up of hands unboxing a sleek new smartphone from a minimalist white box,
carefully peeling back tissue paper to reveal the device, overhead angle, clean
desk with warm sidelight, satisfying unboxing process, 9:16 vertical
```
**Model:** Use `search_models` — text-to-video.

**Voiceover:**
```
"This might be the best phone of 2026. Let me show you why in 60 seconds."
```
**Model:** Use `search_models` — text-to-speech.

---

### Example 4: LinkedIn Professional Post

**Image:**
```
Clean data visualization showing upward growth trend, modern flat design style,
blue and white corporate color palette, title text "AI Adoption in Enterprise:
2024 vs 2026" at top, bar chart with clear labels, professional infographic
quality, 1200x628
```
**Model:** Use `search_models` — design-focused image generation.

---

### Example 5: Twitter/X Thread Visual

**Image (attention-grabbing header):**
```
Dramatic split comparison image: left half shows a cluttered chaotic desk with
papers and mess labeled "Before", right half shows the same desk perfectly
organized and minimal labeled "After", clean overhead shot, bright even
lighting, 1600x900
```
**Model:** Use `search_models` — image generation for the base image, text-rendering model for the text labels.

---

### Example 6: Instagram Reels — Fashion

**Video:**
```
A woman walking confidently toward the camera on a sunlit European cobblestone
street, wearing an oversized beige blazer and white wide-leg trousers, golden
hour backlighting creating a warm glow, slow motion, slight handheld sway,
fashion editorial feel, 9:16 vertical, 10 seconds
```
**Model:** Use `search_models` — text-to-video.

**Music:**
```
Chic minimal electronic beat, French house influence, light and airy synths,
subtle groove, fashion runway vibe, 110 BPM, 15 seconds
```
**Model:** Use `search_models` — music generation.

---

### Example 7: TikTok Surreal / Viral

**Video:**
```
A tiny chef (3 inches tall) cooking a full meal on a regular-sized kitchen
counter, carefully stirring a massive pot with a comically large spoon, camera
at counter-level eye-line with the tiny chef, tilt-shift miniature effect,
warm kitchen lighting, whimsical and surreal, 9:16 vertical
```
**Model:** Use `search_models` — text-to-video (handles surreal scale manipulation well).

---

### Example 8: YouTube Thumbnail

**Image:**
```
Shocked face of a young man with mouth wide open and hands on cheeks, vibrant
yellow background, large bold red text "I TESTED IT" at the top, high contrast,
saturated colors, clean cutout look, YouTube thumbnail style, 1280x720
```
**Model:** Use `search_models` — text-to-image with text rendering for text + face, or generate face with an image generation model then add text with a text-rendering model.

---

### Example 9: Instagram Stories — Poll Engagement

**Background image:**
```
Split image: left side shows iced coffee with whipped cream, right side shows
hot black coffee in a ceramic mug, both on a marble surface, warm café lighting,
top space for "WHICH ONE?" text and bottom space for poll sticker, 1080x1920
```
**Model:** Use `search_models` — image generation.

---

### Example 10: Multi-Platform Product Launch

**Hero product image (base asset):**
```
Premium wireless headphones floating at a slight angle against a gradient
background transitioning from deep purple to black, dramatic volumetric light
rays from behind, subtle lens flare, product hovering with soft shadow below,
ultra-clean commercial photography, 1080x1080
```
**Model:** Use `search_models` — image generation.

**Adapt to platforms (use `search_models` to find appropriate models):**
- **Instagram Feed:** Use 1080x1080 as-is or extend to 1080x1350 with an image editing model to add negative space above/below for text.
- **TikTok/Reels:** Animate with an image-to-video model adding slow rotation and light movement at 9:16.
- **YouTube Thumbnail:** Crop/extend to 1280x720 with an image editing model, add bold text with a text-rendering model.
- **Twitter/X:** Extend to 1600x900 with an image editing model.
- **LinkedIn:** Extend to 1200x628 with an image editing model, add professional copy space.

---

### Example 11: SFX for Transitions

**Swoosh transition:**
```
Model: Use search_models — sound effects generation
Prompt: "Quick cinematic whoosh transition, modern, clean, 0.5 seconds"
```

**Impact hit for text reveal:**
```
Model: Use search_models — sound effects generation
Prompt: "Deep bass impact hit with subtle reverb tail, cinematic, 1 second"
```

**Satisfying click for product placement:**
```
Model: Use search_models — sound effects generation
Prompt: "Crisp magnetic snap click, premium product sound, satisfying, short"
```

