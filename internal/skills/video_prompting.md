# Video Prompting — AI Video Generation Master Guide

Complete reference for AI video generation prompt engineering on fal.ai (March 2026).
Every endpoint, parameter, and technique in this document is verified and actionable.

---

## Quick Reference

### Golden Rules

1. **Front-load subject and action.** The first 20-30 words carry the most weight. Put your primary subject and its main motion at the very start.
2. **One motion, one camera move.** Pick one primary action and one camera movement per shot. Two competing motions cause distortion and warping.
3. **40-60 words sweet spot.** Too short loses control; too long dilutes focus.
4. **Over-describe static elements.** Anything not explicitly named may morph or shift color mid-clip. Lock down clothing, colors, materials, and background details.
5. **"Slow motion" reduces all artifacts.** The single most reliable quality booster — use it whenever motion complexity is high.
6. **Style anchor first.** Start with `"Cinematic 35mm"`, `"high-end commercial"`, `"photorealistic documentary"` to steer the entire aesthetic.
7. **Image-to-Video (I2V) is industry standard.** Generate a perfect still first, then animate it — far more control than pure T2V.

### Universal Prompt Formula

```
[Style anchor] + [Subject + action] + [Camera movement] + [Lighting] + [Environment] + [Mood] + [Technical specs]
```

### Most Commonly Needed Patterns

**Cinematic narrative:**
`[Genre/period], [shot type], [character with locked appearance], [single action], [3 anchored environment details], [one camera move], [film stock or director reference]`

**Product reveal:**
`[Product + material + surface], gentle orbit or slow dolly, single dramatic key light, [luxury/premium quality anchors], shallow depth of field, smooth motion`

**Nature / landscape:**
`Slow aerial drone shot over [landscape] at [time of day], [2-3 environmental motion elements], rising aerial reveal, [film/director reference], [audio cue]`

**Social media vertical (9:16):**
`[Centered subject], [simple relatable action], [portrait-filling background], warm natural light, gentle handheld, slow motion, 3-5 seconds`

### Cheat Sheet

- **Lock colors explicitly**: `"deep navy blue jacket"`, `"warm golden hour light"` — not just `"blue"` or `"sunset"`
- **Artifact prevention**: `temporally consistent, no flickering, stable composition, anatomically correct`
- **Quality realism stack**: `shot on ARRI Alexa 65, Kodak Vision3 500T, film grain, anamorphic lens`
- **Positive framing**: say "empty street" not "no people", "cloudless sky" not "no clouds"
- **Motion hierarchy**: subject motion → camera motion → environmental motion → light motion (describe in this order)
- **Workflow**: draft on a fast model → refine prompt → final render on a premium model. Use search_models to find the best current option.

---

## Universal Video Prompt Formula

```
[Style anchor] + [Subject description] + [Action/Motion] + [Camera movement] +
[Lighting] + [Environment/Setting] + [Mood/Atmosphere] + [Technical specs]
```

### Golden Rules

1. **Front-load the subject and action.** The first 20-30 words carry the most weight in every model. Put your primary subject and its main motion at the very beginning of the prompt.
2. **40-60 words is the sweet spot for most models.** Too short lacks control; too long dilutes focus. Some models handle longer prompts well, others prefer shorter ones — check model documentation via search_models.
3. **One primary motion per shot.** A single clear action always beats multiple simultaneous motions. Add at most one optional secondary motion (e.g., wind in hair while walking). Two competing primary motions cause distortion.
4. **One camera movement per shot.** Combining dolly + pan + crane causes warping. Pick one camera movement and commit to it.
5. **Over-describe static elements.** Anything not explicitly described may morph, shift color, or change shape between frames. Lock down clothing, background, colors, materials, and textures.
6. **Style anchor first.** Starting with "1970s romantic drama, 35mm warm halation" or "high-end luxury commercial, clean lighting" steers the entire generation aesthetic.
7. **"Slow motion" is a cheat code.** It reduces artifacts across all models by slowing the temporal complexity the model must resolve.
8. **Image-to-Video (I2V) is industry standard.** Generate a perfect still frame first, then animate it. This gives far more control over composition, character consistency, and style than pure text-to-video.
9. **Positive framing for negatives.** Say "empty street" not "no people". Say "cloudless sky" not "no clouds". Most models respond poorly to negation words.
10. **5-10-1 workflow.** Generate 5 variations on a fast/cheap model, refine the best prompt 10 times, then do 1 final render on a premium model.

### Motion Hierarchy

When describing motion, order matters. The model allocates attention budget in this priority:

1. **Subject motion** (highest priority) — what the main subject physically does
2. **Camera motion** — how the camera moves through or frames the scene
3. **Environmental motion** — wind, water, particles, ambient movement
4. **Light motion** — shifting shadows, flickering, light rays moving

Place them in this order in your prompt for best results.

---

## Specialized: Lip Sync / Avatar

Use `search_models` to find current lip sync, avatar, and talking head models. Common capabilities to search for:

| Capability | Description |
|-----------|-------------|
| Lip sync | Apply lip movements to existing video from audio |
| AI avatar | Generate talking avatar from image + audio |
| Full-body human animation | Realistic full-body motion from reference image |
| Enterprise video agent | Production-ready avatar for commercial use |

### Lipsync Workflow
1. **Generate or select a portrait image** — front-facing, neutral expression, good lighting, high resolution
2. **Prepare audio** — clean speech audio, minimal background noise, consistent volume, 16kHz+ sample rate, mono preferred
3. **Choose model** — use `search_models` to find a model matching your need:
   - Quick talking head: search for an AI avatar model
   - Lipsync on existing video: search for a lip sync model
   - Full-body realistic: search for a full-body human animation model
   - Commercial/ad avatar: search for an enterprise avatar model
4. **Portrait requirements:** Face clearly visible, centered, well-lit. Avoid extreme angles. Head-and-shoulders framing.
5. **Audio tips:** Normalize audio levels, keep under 60 seconds per clip, remove background noise

---

## Camera Movement Keywords

### Slow / Smooth — Safest, Highest Quality

| Keyword | Description | Best For |
|---------|-------------|----------|
| `static camera` | Locked-off, zero movement | Maximum stability, dialogue, product shots |
| `slow dolly in` | Camera glides forward toward subject | Building intimacy, focus pull |
| `slow dolly out` | Camera pulls back from subject | Reveal, creating emotional distance |
| `slow push-in` | Subtle forward creep | Building tension, anticipation |
| `slow pull-back` | Gradual backward reveal | Revealing context, surprise |
| `gentle orbit` | Camera slowly circles subject | Product showcase, hero shot |
| `smooth tracking shot` | Camera follows alongside action | Walking scenes, lateral motion |
| `slow crane up` | Camera rises smoothly | Establishing shot, grandeur reveal |
| `slow crane down` | Camera descends smoothly | Approaching subject, intimacy |
| `subtle handheld` | Slight natural shake | Documentary feel, realism, intimacy |
| `steadicam follow` | Smooth following movement | Walk-and-talk, exploration |
| `slow lateral pan` | Gentle horizontal sweep | Surveying a scene, panorama |
| `slow tilt up` | Camera angles upward slowly | Revealing height, power, wonder |
| `slow tilt down` | Camera angles downward slowly | Descending to subject, discovery |

### Dynamic — Higher Artifact Risk, Use with Shorter Durations

| Keyword | Description | Best For |
|---------|-------------|----------|
| `whip pan` | Fast horizontal sweep | Transitions, energy, surprise |
| `fast tracking` | Quick lateral follow | Action scenes, chase sequences |
| `zoom in` | Focal length increase (not physical move) | Dramatic emphasis, shock |
| `zoom out` | Focal length decrease | Context reveal, surprise |
| `aerial flyover` | Drone-style forward flight | Landscape establishing shots |
| `rising aerial` | Drone ascending reveal | Location reveal, epic scope |
| `first-person POV` | Subjective camera | Immersion, horror, action |
| `crash zoom` | Rapid zoom into detail | Comedy, shock, emphasis |
| `dutch angle` | Tilted horizon | Unease, disorientation, tension |
| `low angle push` | Ground-level forward movement | Power, intimidation |

### Specialized

| Keyword | Description | Best For |
|---------|-------------|----------|
| `360 orbit` | Full rotation around subject | Product hero, character reveal |
| `bullet time` | Frozen moment, camera orbits | Action highlight, dramatic pause |
| `rack focus` | Focus shift foreground to background | Directing attention, depth storytelling |
| `time-lapse camera` | Implied fast-forward | Passage of time, construction, weather |
| `vertigo effect` | Dolly zoom (zoom in + dolly out) | Psychological tension, disorientation |
| `macro lens extreme close-up` | Tight detail shot | Texture, product detail, nature |
| `bird's eye view` | Directly overhead | Maps, patterns, spatial context |
| `worm's eye view` | Looking straight up | Drama, architecture, power |
| `over-the-shoulder` | Framed past character's shoulder | Dialogue, POV context |

---

## Motion & Action Keywords

### Speed Control
`slow motion` | `ultra slow motion` | `real-time` | `time-lapse` | `hyper-speed` | `frozen moment` | `bullet time` | `half speed` | `double speed` | `ramping from slow to fast` | `speed ramp` | `variable speed`

### Natural / Environmental Motion
`wind blowing hair` | `fabric flowing in breeze` | `water rippling` | `leaves rustling` | `fire flickering` | `smoke drifting upward` | `clouds moving slowly` | `grass swaying` | `snow falling gently` | `rain streaming down glass` | `dust particles floating in light beam` | `fog rolling in` | `waves crashing` | `sand blowing across dunes` | `aurora borealis pulsing` | `tide slowly receding` | `mist rising from lake surface`

### Human Motion
`walking confidently` | `running` | `turning head slowly` | `gesturing with hands` | `reaching forward` | `dancing gracefully` | `breathing deeply` | `blinking` | `smiling slowly` | `standing up from chair` | `leaning against wall` | `crossing arms` | `brushing hair behind ear` | `looking over shoulder` | `nodding` | `shrugging` | `tilting head` | `furrowing brow` | `raising eyebrow`

### Liquid / Fluid
`pouring liquid` | `splashing water` | `dripping slowly` | `flowing downstream` | `swirling in glass` | `cascading waterfall` | `bubbling` | `ink dispersing in water` | `mercury pooling` | `oil on water rainbow` | `condensation forming` | `melting ice` | `honey dripping`

### Light / Shadow Motion
`flickering candlelight` | `moving shadows` | `light rays shifting` | `neon pulsing` | `sun moving across floor` | `headlights sweeping` | `lightning flash` | `sunrise gradual brightening` | `spotlight following` | `lens flare drifting` | `shadow crawling across wall`

### Particle / Material
`sparks flying` | `embers rising` | `confetti falling` | `petals scattering` | `glass shattering` | `paper floating` | `feathers drifting` | `bubbles rising` | `glitter dispersing` | `sand pouring through fingers` | `ash falling` | `snowflakes accumulating`

### Weather
`rain falling` | `storm approaching` | `lightning striking` | `thunder clouds rolling` | `wind gusting` | `blizzard whiteout` | `fog bank advancing` | `rainbow forming` | `hail bouncing` | `drizzle on water surface`

### Mechanical / Technical
`gears turning` | `clock hands moving` | `pistons pumping` | `wheels spinning` | `conveyor belt rolling` | `doors opening` | `elevator rising` | `robotic arm extending` | `engine revving` | `switches flipping` | `dials rotating`

---

## Prompt Engineering by Use Case

### Cinematic Narrative

**Template:**
```
[Time period/genre] + [Shot type from distance] + [Character with specific clothing/features] +
[Specific action with direction] + [Environment with 3+ anchored details] + [Single camera move] +
[Lighting setup] + [Film reference or director name] + [Mood word]
```

**Example — Film Noir:**
```
1940s detective noir, medium shot, a weary private investigator in a rumpled grey fedora and
rain-soaked trench coat steps into a dimly lit office, rain dripping from his hat brim, he
slowly removes his gloves and tosses them onto a cluttered wooden desk, single desk lamp
casting hard shadows across venetian blind patterns on the wall, slow push-in, Kodak Tri-X
film grain, David Fincher atmosphere, world-weary tension
```

### Product Reveal

**Template:**
```
[Product type] + [Material/finish description] + [Surface/background] + [gentle orbit or
slow dolly] + [Single dramatic light source] + [Quality anchors: premium, luxury, commercial]
```

**Example — Perfume:**
```
Elegant crystal perfume bottle with gold cap on polished white marble surface, gentle
360-degree orbit, single soft key light from upper left creating delicate caustic reflections
through the crystal, soft gradient background fading to warm cream, golden liquid catching
light as the camera moves, luxury fragrance commercial, premium product photography, shallow
depth of field, smooth motion
```

### Character Action

**Template:**
```
[Character description with locked appearance] + [Single clear physical action with direction] +
[One secondary ambient motion] + [Environment with 3 static anchors] + [Camera: tracking or
push-in] + [Lighting] + [Emotional tone]
```

**Example — Warrior:**
```
A battle-scarred female warrior in dented bronze armor and a torn crimson cloak draws a
longsword from a worn leather scabbard at her hip, wind whipping her braided dark hair,
standing atop a rocky cliff overlooking a misty valley at dawn, slow tracking shot from her
side, golden hour rim light creating a warm halo, epic fantasy atmosphere, Lord of the Rings
cinematic quality
```

### Nature / Landscape

**Template:**
```
[Aerial or establishing shot type] + [Primary landscape feature] + [Time of day and weather] +
[2-3 environmental motion elements] + [Slow camera move] + [Atmospheric quality keywords] +
[Scale reference]
```

**Example — Mountain Lake:**
```
Sweeping aerial drone shot over a pristine alpine lake surrounded by snow-capped mountains at
golden hour, mist rising from the turquoise water surface, a solitary wooden rowboat drifting
near the shore, pine forests reflected perfectly in the still water, rising aerial revealing
the full valley, IMAX quality, Denis Villeneuve atmosphere, the sound of gentle wind and
distant birdsong
```

### Urban / City

**Template:**
```
[Urban setting with time] + [Street-level or aerial perspective] + [Human activity or traffic] +
[Architectural details] + [Lighting: neon, street lamps, car lights] + [Camera move] +
[Weather/atmosphere]
```

**Example — Tokyo Night:**
```
Neon-drenched Shibuya crossing at night, tracking shot at street level following a lone figure
with a red umbrella walking through crowds, rain-slicked pavement reflecting thousands of
colorful signs, steam rising from manhole covers, taxis and pedestrians blurred in background,
shallow depth of field, Wong Kar-wai atmosphere, Cinestill 800T tungsten warmth, melancholic
urban beauty
```

### Abstract / Artistic

**Template:**
```
[Surreal concept] + [Material/texture transformation] + [Non-literal motion] + [Color palette] +
[Lighting mood] + [Art reference] + [slow motion strongly recommended]
```

**Example — Metamorphosis:**
```
A porcelain mask slowly cracks and dissolves into a flock of white paper butterflies that
spiral upward into a sky made of flowing watercolor paint, each butterfly trailing trails of
gold leaf that scatter into luminous dust, warm amber and deep ultramarine color palette,
soft volumetric light, slow motion, surreal magical realism, Rene Magritte meets studio Ghibli
```

### Social Media (Vertical 9:16)

**Template:**
```
[Subject centered for vertical frame] + [Action that works in portrait orientation] +
[Background that fills vertical space] + [Warm/inviting lighting] + [Relatable emotion] +
[Keep 3-5 seconds]
```

**Example — Lifestyle:**
```
A young woman in an oversized cream sweater walks toward the camera through a sun-dappled
autumn forest path, golden leaves falling around her, she looks up and smiles warmly, soft
natural backlighting creating a warm glow, shallow depth of field, gentle handheld camera
movement, warm cozy aesthetic, Instagram autumn mood, smooth slow motion
```

---

## Artifact Prevention

### Common Issues & Fixes

| Problem | Cause | Fix |
|---------|-------|-----|
| **Subject morphing** | Insufficient visual description | Over-describe all static features (clothing, colors, materials). Use I2V with reference image. |
| **Background warping** | Too much camera motion | Use `static camera` or only `slow` movements. Describe background elements explicitly. |
| **Face distortion** | Complex angles, profile views | Keep faces front-facing or 3/4 angle. Add "clear facial features, portrait angle". |
| **Flickering/strobing** | Multiple light sources, complex lighting | Use "single key light", simplify lighting. Add to negative: "flickering, strobing". |
| **Temporal inconsistency** | Long duration | Keep clips 3-5 seconds. Chain shorter clips in post. |
| **Unnatural motion** | Vague action description | Specify exact motion: "slowly raises right hand palm-up" not "moves hand". |
| **Extra limbs/fingers** | Hand close-ups, complex poses | Avoid hand detail shots. Frame to crop hands. Add "anatomically correct". |
| **Text corruption** | Text elements in scene | Never generate text in video — overlay in post. Describe signage as "blurred signs". |
| **Color shifting** | No color anchoring | Explicitly state colors: "deep navy blue suit", "warm golden hour light". |
| **Jittery motion** | High complexity + long duration | Add "smooth motion, fluid, consistent". Reduce simultaneous actions. Use slow motion. |
| **Object duplication** | Ambiguous spatial layout | Clearly describe spatial relationships: "single red car in center of frame". |
| **Audio desync** | Complex action + native audio | Simplify the visual action when using audio generation. Describe sounds that naturally match the motion. |
| **Motion blur excess** | Fast motion + low frame rate | Use "sharp detail, crisp motion" or slow down the action. |
| **Style inconsistency** | No style anchor | Add persistent style keywords at prompt start: "photorealistic, shot on 35mm film". |

### Quality Maximizers

**For Realism:**
```
photorealistic, hyperrealistic, shot on RED Monstro 8K, shot on ARRI Alexa 65,
Kodak Vision3 500T film stock, natural lighting, film grain, shallow depth of field,
lens distortion, chromatic aberration, bokeh, anamorphic lens flare, RAW footage,
IMAX quality, 8K resolution, photographic detail
```

**For Stability:**
```
smooth motion, fluid movement, consistent lighting, coherent, steady camera,
stable composition, uniform color palette, temporally consistent, no flickering,
seamless, continuous, controlled motion
```

**For Cinematic Look:**
```
cinematic, cinematic color grading, Hollywood production quality, feature film look,
dramatic lighting, three-point lighting, chiaroscuro, Rembrandt lighting, golden hour,
lens flare, volumetric light, god rays, haze, atmospheric, deep shadows, rich highlights
```

**Style-Specific Anchors:**
```
35mm film — warm, organic, slightly grainy
16mm film — indie, gritty, handheld feel
IMAX 70mm — epic scale, extreme detail, immersive
anamorphic — horizontal lens flares, oval bokeh, widescreen compression
Kodachrome — warm reds, rich greens, vintage feel
Fujifilm Velvia — saturated colors, vivid landscapes
teal and orange grading — blockbuster action movie look
desaturated — muted, somber, serious tone
high contrast black and white — noir, dramatic, timeless
pastel palette — soft, dreamy, gentle
neon-lit — cyberpunk, nightlife, urban
```

---

## Negative Prompt Keywords

For models supporting negative prompts (check model documentation):

### Standard Negative Prompt
```
blurry, low resolution, distorted, morphing, flickering, jittery, unnatural motion,
extra limbs, bad anatomy, watermark, text overlay, oversaturated, artifacts, glitch,
frame skip, duplicate frames, compression artifacts, noise, excessive grain
```

### For Human Subjects
```
deformed face, extra fingers, fused fingers, bad hands, distorted proportions,
cross-eyed, asymmetric face, uncanny valley, plastic skin, dead eyes,
unnatural skin texture, floating hair
```

### For Environments
```
warping background, inconsistent shadows, floating objects, impossible geometry,
texture pop-in, z-fighting, clipping, repeating textures, tiling artifacts
```

### For Motion
```
jittery camera, stuttering motion, sudden jumps, teleporting objects,
inconsistent speed, rubber banding, sliding feet, moon walking
```

---

## Resolution & Duration Reference

### By Platform

| Platform | Resolution | Aspect Ratio | Duration |
|----------|-----------|--------------|----------|
| TikTok / Reels / Shorts | 1080x1920 | 9:16 | 3-15s |
| Instagram Feed | 1080x1080 | 1:1 | 3-60s |
| YouTube Standard | 1920x1080 | 16:9 | any |
| YouTube Shorts | 1080x1920 | 9:16 | up to 60s |
| Twitter/X | 1920x1080 | 16:9 | up to 140s |
| LinkedIn | 1920x1080 | 16:9 | 3-10min |
| Cinema Widescreen | 2560x1080 | 21:9 | any |
| Cinema Standard | 1920x816 | 2.39:1 | any |

### Model Resolution Support

Use `search_models` to find current models and their resolution, duration, and audio capabilities. Key parameters to check:
- **Resolution**: ranges from 480p to 4K depending on the model
- **Max duration**: typically 5-15 seconds per clip
- **Native audio**: some models can generate synchronized audio; others require post-production audio

---

## Workflow Patterns

### I2V Standard — Best Quality Pipeline
1. Generate a perfect still frame with an image model (Flux, Midjourney, etc.)
2. Ensure the image matches the target video aspect ratio exactly
3. Upload the image to the I2V endpoint of your chosen model
4. In the I2V prompt, describe only the motion — the image provides all visual context
5. Keep camera movements minimal — the model must extrapolate unseen areas

### Multi-Shot Narrative
- Some models support a `multi_prompt` array with 3-5 scene descriptions, each maintaining character description consistency
- Search for a "multi-shot" or "narrative" video model that handles shot composition and transitions automatically
- **Manual chaining**: Generate individual clips with consistent style anchors, assemble in post

### Audio-Synced Video
- Some models support `generate_audio: true` — describe ambient sounds, dialogue, and SFX naturally in the prompt
- Some models support `voice_ids` for character-specific dialogue voices
- Some models accept an `audio_url` parameter to sync video to an external audio file
- **Post-production**: Generate video without audio, then use `workflowutils/join-audio-video` to add a separate audio track

### First/Last Frame Control
- Search for models that support "first-last-frame-to-video" — provide start and end frame images, model generates the transition
- Use case: precise camera moves, character pose changes, scene transitions with known endpoints

### Reference-to-Video (Character Consistency)
- Search for models that support "reference-to-video" — provide reference images for style/character without them being the first frame
- Some models support multi-reference for identity consistency across shots

### LoRA Custom Style
- Search for video models that support LoRA — apply custom LoRA weights for branded or artistic styles
- Some models are trainable for fine-tuning on custom content

### Fast Iteration Pipeline (5-10-1)
1. **5 drafts**: Generate 5 variations on a fast/budget model — test concepts quickly
2. **10 refinements**: Take the best prompt, iterate 10 times on a mid-tier model — dial in the details
3. **1 final render**: Render the perfected prompt on a premium model at maximum quality

---

## Director & Style References

Using director names as style anchors is one of the most powerful techniques across all models:

| Director/Reference | Visual Style Activated |
|--------------------|----------------------|
| Denis Villeneuve | Vast landscapes, orange/teal, atmospheric haze, IMAX scale, slow deliberate camera |
| Roger Deakins | Masterful natural lighting, single-source motivation, painterly composition |
| Wes Anderson | Perfect symmetry, pastel palette, centered framing, dollhouse aesthetic |
| Christopher Nolan | IMAX practical, deep focus, temporal complexity, blue/steel tones |
| Ridley Scott | Atmospheric smoke/fog, backlit silhouettes, dystopian beauty, layered depth |
| Wong Kar-wai | Neon-soaked, slow motion, step-printed motion blur, romantic melancholy |
| Terrence Malick | Golden hour, magic hour, natural light only, whispered voiceover feel |
| David Fincher | Desaturated green/yellow, precise camera movement, clinical framing |
| Andrei Tarkovsky | Poetic slow cinema, water elements, long takes, spiritual atmosphere |
| Stanley Kubrick | One-point perspective, symmetrical framing, cold precise lighting |
| Studio Ghibli | Hand-painted watercolor, gentle motion, nature reverence, magical realism |
| Blade Runner | Neon noir, rain-soaked streets, holographic ads, dystopian cyberpunk |
| Emmanuel Lubezki | Natural light only, long takes, golden hour, wide-angle immersion |
| Akira Kurosawa | Weather as character, deep staging, bold graphic composition |

---

## Lighting Keywords Reference

| Keyword | Effect |
|---------|--------|
| `golden hour` | Warm, low-angle sunlight, long shadows, magical |
| `blue hour` | Cool twilight tones, pre-dawn or post-sunset |
| `magic hour` | Combined golden/blue, cinematic transition light |
| `Rembrandt lighting` | Triangle of light on shadowed cheek, dramatic portrait |
| `butterfly lighting` | Symmetrical shadow under nose, glamour/beauty |
| `rim lighting` | Bright edge outline separating subject from background |
| `backlight` | Light behind subject creating silhouette or glow |
| `chiaroscuro` | Extreme light/dark contrast, Caravaggio style |
| `volumetric light` | Visible light rays through atmosphere/dust/fog |
| `god rays` | Dramatic light beams through clouds or windows |
| `neon lighting` | Colored artificial light, urban/cyberpunk |
| `practical lighting` | In-scene light sources (lamps, candles, screens) |
| `three-point lighting` | Key + fill + back, professional studio standard |
| `high key` | Bright, minimal shadows, clean/happy/commercial |
| `low key` | Dark, dramatic shadows, moody/noir/thriller |
| `motivated lighting` | Light that has a visible or logical source |
| `bounce light` | Soft reflected light, natural fill |
| `dappled light` | Light filtered through leaves/blinds, organic patterns |
| `split lighting` | Half-face shadow, duality, dramatic portrait |
| `silhouette` | Extreme backlight, subject as dark shape |

---

## Additional Example Prompts

### Sci-Fi VFX
**Recommended:** a high-quality text-to-video model | **Resolution:** 1080p | **Duration:** 6s
```
A massive alien mothership slowly emerges from thick storm clouds above a deserted highway,
blue-white light beams cutting through rain and fog, cars abandoned on the road below, the
ship's surface covered in bioluminescent patterns pulsing rhythmically, low rumbling vibration
shakes puddles on the asphalt, slow crane up from ground level to reveal the full scale,
Ridley Scott meets Denis Villeneuve, anamorphic lens flare, deep bass rumble and eerie
atmospheric sound design
```

### Food / Culinary
**Recommended:** a text-to-video model with good detail rendering | **Duration:** 6s
```
A chef's hands crack a fresh egg into a sizzling cast iron skillet, the white immediately
bubbles and sets while the golden yolk remains perfectly runny, steam rising in soft morning
window light, rustic wooden kitchen counter, overhead camera angle slowly pulling back,
warm appetizing color palette, food photography perfection
```

### Multi-Shot Jazz Club
**Recommended:** a model supporting multi_prompt | **Duration:** 12s
```json
{
  "multi_prompt": [
    "Wide establishing shot of a dimly lit jazz club at night, warm amber lighting, smoke drifting through spotlight beams, a pianist in a black vest sits at a grand piano, audience silhouettes in foreground, 1960s atmosphere",
    "Close-up of the pianist's hands beginning to play, fingers moving gracefully over ivory keys, warm spotlight from above, shallow depth of field, the sound of a melancholic jazz melody beginning",
    "Medium shot from the side, the pianist closes his eyes and leans into the music, swaying gently, golden rim light on his profile, smoke curling past, deep emotion on his face, cinematic color grading"
  ],
  "duration": "12",
  "aspect_ratio": "16:9",
  "generate_audio": true
}
```

### Fashion Editorial
**Recommended:** a text-to-video model supporting ultra-widescreen | **Duration:** 8s | **Aspect:** 21:9
```
Ultra-widescreen fashion editorial, a model in a structured ivory blazer and wide-brimmed hat
walks slowly down an empty modernist corridor with floor-to-ceiling windows, harsh geometric
shadows from afternoon sun creating dramatic light patterns on the marble floor, camera fixed,
she turns her head slightly toward the camera as she passes, editorial fashion photography,
clean minimal aesthetic, Helmut Newton meets Wes Anderson
```

### Horror Atmosphere
**Recommended:** a text-to-video model with negative prompt support | **Duration:** 5s | **Resolution:** 1080p
```
A long dark hospital corridor, single overhead fluorescent light flickering at the far end,
the camera slowly dollies forward, a child's wheelchair sits abandoned in the center of the
hallway, dust particles float through the stuttering light, peeling paint on institutional
green walls, something shifts in the shadows at the end of the corridor, low key lighting,
unsettling atmosphere, The Shining meets Silent Hill
```
**Negative:** `bright colors, cheerful, cartoon, blurry, distorted, watermark`

### Timelapse City
**Recommended:** a premium text-to-video model with 4K support | **Resolution:** 4K | **Duration:** 8s
```
Hyper-lapse from a rooftop overlooking a major city skyline transitioning from golden sunset
to deep blue night, clouds streaking across the sky, building lights flickering on one by
one, car headlights creating red and white light trails on the highways below, the city
transforming into a glowing grid of lights, shot on ARRI Alexa with time-lapse photography,
Christopher Nolan atmosphere, ambient city hum transitioning to night soundscape
```

---

## Quick Reference Card

**Fastest path to quality video:**
1. Write prompt using the Universal Formula (40-60 words)
2. Start with a style anchor
3. One subject, one action, one camera move
4. Over-describe everything static
5. Use "slow motion" if artifacts appear
6. Use I2V when possible (generate still first, then animate)
7. Draft on a fast model, refine on a mid-tier model, final render on a premium model. Use search_models to find the best current options.

**Audio in video:**
- Some models support native audio generation (ambient + dialogue + SFX)
- Some models support `voice_ids` for character-specific dialogue
- Some models accept an `audio_url` to sync video to external audio
- Use search_models to find models with audio capabilities

**Key specs vary by model** — use search_models to find current models with the best duration, resolution, and audio support for your needs.
