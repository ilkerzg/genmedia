# Storytelling — Multi-Shot Narrative Production Master Guide

Complete reference for planning and producing multi-shot narrative videos and visual stories using fal.ai generative models. Use `search_models` to discover available models for each task.

---

## Quick Reference

> **Model Discovery:** Use `search_models` to find the best models for each task. Search by category: "text to video" for video generation, "image generation" for storyboard frames, "text to speech" for narration, "voice generation" for dialogue.

### Top 5 Storytelling Tips

1. **Lock a style anchor prefix.** Start every shot prompt with an identical style string (film stock, color grade, lens type, aspect ratio). Solves 80% of visual inconsistency.
2. **Follow the pacing cycle: Slow → Medium → Fast → IMPACT → Slow.** The "breathe" moment after each climax is mandatory.
3. **Wide → Medium → Close-Up → Insert is the universal grammar.** Open scenes wide, cut in to develop, close-up for emotion, inserts for detail.
4. **Character lock is non-negotiable.** Write an exact character description and copy it verbatim into every shot.
5. **Let music do narrative work.** Drop all music for 2-3s before a key reveal — silence creates more anticipation than any sound.

### Key Patterns

**Native multi-shot:** Some video models support `multi_prompt` arrays with per-shot `prompt` and `duration`. Use `search_models` to find models with multi-shot or native audio support.

**I2V pipeline (maximum control):** Generate storyboard key frames with an image generation model → animate each with an image-to-video model → assemble.

**First-last frame (controlled motion):** Generate shot start and end frames → feed both to a first-last-frame-to-video model.

**Reference continuity:** One clean character image → pass as reference to a reference-to-video model for every shot.

### Cheat Sheet

| Scenario | Approach |
|---|---|
| Quick 15-30s ad, audio needed | Multi-shot video model with `generate_audio: true` |
| Same character across many shots | Reference-to-video model per shot |
| Precise motion / start+end locked | First-last frame video model |
| 60s+ brand story, full control | Storyboard (image gen) → I2V pipeline |

**Shot duration rules of thumb:** Hook/insert = 2-3s. Medium action = 4-5s. Establishing/emotional = 6-10s. Fast-cut montage = 1.5-2s per cut.

---

## Narrative Structure for AI Video

Every compelling story, even a 15-second ad, follows structure. AI-generated video is no different. The difference is that each structural beat maps directly to one or more generated shots.

### Three-Act Structure

| Act | Purpose | Typical Shot Types | % of Total Duration |
|-----|---------|-------------------|---------------------|
| Act 1: Setup | Establish world, character, tone | Establishing wide, character intro medium, detail insert | 25% |
| Act 2: Confrontation | Introduce conflict, raise stakes, complicate | Action shots, reaction close-ups, montage, parallel cuts | 50% |
| Act 3: Resolution | Climax, resolve, final image | Climax action, emotional close-up, final wide or detail | 25% |

**Mapping acts to shots (30-second story, 6 shots):**
- Act 1 (7.5s): Shot 1 establishing wide (4s) + Shot 2 character intro (3.5s)
- Act 2 (15s): Shot 3 inciting incident (4s) + Shot 4 complication (4s) + Shot 5 climax build (7s)
- Act 3 (7.5s): Shot 6 resolution/final image (7.5s)

**Mapping acts to shots (60-second story, 10 shots):**
- Act 1 (15s): Shots 1-3 (establishing, character, world detail)
- Act 2 (30s): Shots 4-8 (incident, rising action, midpoint twist, complication, climax approach)
- Act 3 (15s): Shots 9-10 (climax, denouement/final image)

### Shot Sequence Planning

| Story Length | Recommended Shots | Avg Shot Duration | Pacing |
|-------------|-------------------|-------------------|--------|
| 15s (social hook) | 3-4 | 4-5s | Fast |
| 30s (ad/teaser) | 5-7 | 4-5s | Medium-fast |
| 60s (brand story) | 8-12 | 5-7s | Variable |
| 2-3 min (short film) | 20-30 | 5-8s | Variable |
| Music video (3-4 min) | 30-50 | 3-6s | Beat-synced |

**Duration guidelines per shot type:**
- Establishing wide: 4-6s (let the audience absorb the world)
- Character intro: 3-5s (enough to register face and posture)
- Action/motion: 3-5s (short, punchy, clear single action)
- Emotional close-up: 5-8s (let the emotion land)
- Montage element: 2-3s (rapid succession)
- Final shot: 5-8s (linger for impact)

### Scene Types and When to Use Them

**Establishing Shot** — Wide or extreme wide. Sets location, time of day, mood. Always the first shot of a new scene or location. Use slow camera movement (aerial descend, slow pan) to give the audience time to absorb.

**Character Introduction** — Medium or medium close-up. First appearance of the subject. Hold steady or use subtle push-in. Over-describe appearance to lock visual consistency for later shots.

**Reaction Shot** — Close-up on face or hands. Communicates emotion without dialogue. Critical after any plot event. Static camera or very slow push-in.

**Insert/Cutaway** — Extreme close-up on a detail, object, or texture. Breaks visual monotony, adds information, hides continuity jumps between shots. 2-3s is enough.

**Over-the-Shoulder** — Implies dialogue or relationship between two characters. Shallow depth of field, one character sharp, the other soft in foreground.

**Montage** — Rapid sequence of 2-3s shots showing passage of time, training, travel, or transformation. Use consistent color grade and camera style across all montage elements.

**Transitional** — Bridges two scenes. Can be a whip pan to black, a slow dissolve, a match cut on shape or movement. Often 2-4s.

---

## Storyboarding with AI

### Frame-by-Frame Planning

Before generating any video, plan every shot as a still frame. This gives you composition control, character locking, and style consistency that pure text-to-video cannot match.

**Recommended:** Use `search_models` to find image generation models for storyboard frames. Look for models with strong photorealism, high detail, and good prompt adherence.

### Visual Script Format

Use this format for every shot in your storyboard. This becomes both your planning document and your prompt source.

```
SHOT 1 — EWS (Extreme Wide Shot) — 5s
[Camera: Slow aerial descend over cityscape]
[Image prompt: Aerial view of a rain-soaked neon-lit Tokyo street at night,
  towering skyscrapers reflecting pink and blue neon on wet asphalt,
  tiny silhouettes of pedestrians with umbrellas, cyberpunk atmosphere,
  cinematic color grading, 2.39:1 anamorphic]
[Video model: an image-to-video model (use search_models)]
[Motion prompt: Camera slowly descends from above, revealing street below,
  rain falling, neon reflections rippling on wet ground]
[Audio: soft rain, distant city hum, muted synthwave bass]
[Transition: slow dissolve to Shot 2]
```

```
SHOT 2 — MCU (Medium Close-Up) — 4s
[Camera: Static, slight handheld micro-movement]
[Image prompt: A young woman in a black leather jacket standing under a
  neon pharmacy sign, rain dripping from the awning above, her face
  half-lit in pink neon half in shadow, determined expression,
  shallow depth of field, anamorphic bokeh]
[Video model: an image-to-video model (use search_models)]
[Motion prompt: Woman looks up slowly, rain continues falling,
  neon sign flickers once, subtle breath vapor]
[Audio: closer rain, neon buzz, single distant car horn]
[Transition: hard cut to Shot 3]
```

### Storyboard-to-Video Pipeline

This is the industry-standard workflow for maximum quality and control.

**Step 1: Write the visual script.** Plan every shot using the format above. Decide shot count, duration, camera, and audio for each.

**Step 2: Generate key frames.** Use an image generation model (find one via `search_models`) to generate the hero frame for each shot. Iterate until composition, character, and lighting are exactly right.

**Step 3: Animate each frame.** Feed each key frame into an Image-to-Video model with a motion-specific prompt. Choose model based on shot requirements (see model selection below).

**Step 4: Generate audio layers.** Create music, narration, and SFX as separate layers (see Audio Storytelling section).

**Step 5: Sequence and export.** Combine video shots in order, layer audio, apply transitions. The CLI handles assembly automatically when given the full shot sequence.

---

## Multi-Shot Video Production

### Native Multi-Shot Video (multi_prompt)

Use `search_models` to find video models with multi-shot/multi-prompt support.

Some video models natively support multi-shot sequences through a `multi_prompt` parameter. This is the simplest path to multi-shot video because the model handles inter-shot transitions internally.

**Key parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `multi_prompt` | array of objects | Each object has `prompt` (string) and `duration` (number, seconds) |
| `shot_type` | string | `"customize"` (you control cuts) or `"intelligent"` (model decides pacing) |
| `generate_audio` | boolean | Enable AI-generated audio track |
| `voice_ids` | array | Specific voice IDs for dialogue/narration |
| `aspect_ratio` | string | `"16:9"`, `"9:16"`, `"1:1"` |

**Example: 30-second Product Ad (native multi_prompt)**
```json
{
  "multi_prompt": [
    {
      "prompt": "Extreme wide shot of a pristine mountain lake at golden hour, mist rising from the water surface, towering pine forests on both sides, cinematic aerial descend, warm golden light, 4K commercial quality",
      "duration": 5
    },
    {
      "prompt": "Medium shot of a sleek titanium water bottle being placed on a mossy rock beside the lake, morning dew on the bottle surface, soft golden backlight creating rim light, shallow depth of field, product commercial lighting",
      "duration": 4
    },
    {
      "prompt": "Close-up of a woman's hand picking up the water bottle, condensation droplets on titanium surface, her reflection visible in the polished metal, slow motion, macro detail, luxury product photography",
      "duration": 4
    },
    {
      "prompt": "Medium close-up of the woman drinking from the bottle with the mountain lake behind her, golden hour backlight creating hair light and lens flare, satisfied expression, aspirational lifestyle commercial",
      "duration": 5
    },
    {
      "prompt": "Extreme wide aerial shot pulling back to reveal the full lake and mountain panorama, the woman tiny in the landscape, the bottle visible in her hand catching a glint of light, epic scale, cinematic drone pullback",
      "duration": 5
    }
  ],
  "shot_type": "customize",
  "generate_audio": true,
  "aspect_ratio": "16:9"
}
```

**Example: 20-second Social Media Story (native multi_prompt)**
```json
{
  "multi_prompt": [
    {
      "prompt": "Close-up of hands cracking an egg into a sizzling cast iron pan, steam rising, warm kitchen lighting, overhead camera angle, food photography, slow motion",
      "duration": 4
    },
    {
      "prompt": "Medium shot of a chef in a white apron flipping the pan with confident wrist motion, golden morning light streaming through kitchen window, bokeh background, professional culinary content",
      "duration": 4
    },
    {
      "prompt": "Close-up of a perfectly plated breakfast on a wooden table, eggs benedict with hollandaise dripping slowly, fresh herbs scattered, food commercial lighting, shallow depth of field, mouth-watering detail",
      "duration": 5
    }
  ],
  "shot_type": "customize",
  "aspect_ratio": "9:16"
}
```

### Multi-Shot Narrative Models

Use `search_models` to find models designed specifically for controllable multi-shot narrative with inter-shot consistency. These models maintain subject identity, environment coherence, and style continuity across shots automatically.

**Key capabilities to look for:**
- Variable shot counts and per-shot durations
- Subject consistency enforcement across all shots
- Motion and camera control per shot
- Background/environment customization per shot
- Style locking across the entire sequence

**When to use a dedicated multi-shot model vs a general multi_prompt model:**
- Dedicated multi-shot model: when character consistency across shots is critical, when you need fine control over subject appearance and motion per shot
- General multi_prompt: when you want fast results with auto-handled transitions, when audio generation is needed inline

### First/Last Frame Workflow

Use `search_models` to find models that support first-last-frame-to-video generation.

This workflow gives you maximum control over the start and end state of each shot. The model interpolates all motion between the two key frames.

**When to use:** Controlled transitions, character position changes, object transformations, guaranteed start/end compositions.

**Workflow:**
1. Generate the first frame of the shot (image model)
2. Generate the last frame of the shot (image model, same style/character but different pose/position)
3. Feed both frames to the first-last-frame endpoint
4. The model creates smooth motion between the two states

**Example: Door opening reveal**
- First frame: closed wooden door, warm light leaking from the crack beneath
- Last frame: same door now open, revealing a sunlit garden with flowers
- Result: smooth door-opening animation with consistent environment

### Reference-to-Video

Use `search_models` to find models that support reference-to-video generation (maintaining character/scene consistency from reference images).

Use character or scene reference images to maintain visual consistency across independently generated shots. This is the primary tool for multi-shot stories where you generate each shot separately but need the same character throughout.

**Workflow:**
1. Generate or select a clear reference image of your character (front-facing, well-lit, distinctive features)
2. For each shot, provide the reference image plus a shot-specific prompt
3. The model preserves the character's appearance while executing the new action/composition

**Best practices:**
- Reference image should show the character's full face, hair, and clothing clearly
- Use a neutral background in the reference for cleanest identity transfer
- Include style anchors in every shot prompt to maintain overall aesthetic
- Works best when the character's clothing and hair are distinctive and describable

### Manual Multi-Shot (Any Model)

When using models without native multi-shot support, maintain consistency manually.

**Style anchor technique:** Start every shot prompt with the identical style prefix.
```
Style prefix: "Cinematic neo-noir, Kodak Vision3 500T film stock, teal and amber
color grading, anamorphic lens flare, 2.39:1 aspect ratio"

Shot 1: [style prefix] + ", extreme wide shot of a rain-soaked alley..."
Shot 2: [style prefix] + ", medium close-up of a detective in a trench coat..."
Shot 3: [style prefix] + ", insert shot of a gloved hand placing a photograph..."
```

**Character locking technique:** Describe the character identically in every shot.
```
Character lock: "a woman in her 30s with short black hair, sharp jawline,
wearing a navy wool peacoat with brass buttons, red scarf"

Use this exact description in every shot that features this character.
```

**I2V from consistent storyboard frames:** Generate all storyboard frames in a single batch session with the same image model, same style prompt, same seed prefix. Then animate each frame with I2V.

---

## Audio Storytelling

### Narration

Use `search_models` to find text-to-speech models. Look for high-quality single-narrator models for narration and multi-speaker models for dialogue scenes.

**Narration pacing guidelines:**
- Average speaking rate: 130-160 words per minute for narration
- For a 30s video, narration text should be 65-80 words
- For a 60s video, narration text should be 130-160 words
- Leave 2-3s of silence at the start and end for breathing room
- Pause narration during high-action shots to let visuals speak
- Sync narration breaks to shot transitions

**Multi-speaker dialogue:** Use a multi-speaker TTS model (find via `search_models`) for scenes with two or more characters speaking. Assign distinct voices per character. Generate each character's lines as separate audio segments timed to their respective shots.

### Sound Design Layers

Professional audio is built in layers. Each layer serves a distinct narrative function.

| Layer | Purpose | Model Type (use search_models) | Example |
|-------|---------|-------------------------------|---------|
| 1. Music | Emotional foundation, energy, mood | A music generation model | Soft piano underscore for emotional scene |
| 2. Ambient | Environmental grounding, sense of place | A text-to-audio model | Forest birds, city traffic, ocean waves |
| 3. Foley | Specific in-scene sounds tied to actions | A sound effects model | Footsteps on gravel, door creaking, glass clink |
| 4. Dialogue/Narration | Story information, character voice | A TTS model | Voice-over narration, character speech |
| 5. SFX | Impacts, transitions, emphasis | A sound effects model | Whoosh for transitions, boom for reveals |

**Layer mixing priority (loudest to quietest):**
1. Dialogue/narration (always on top)
2. Key SFX synced to on-screen action
3. Music (duck under dialogue, swell in gaps)
4. Ambient (constant low bed)
5. Foley (subtle, felt more than heard)

### Music as Narrative Tool

**Prompt patterns for narrative music:**

| Narrative Moment | Music Prompt Keywords |
|-----------------|----------------------|
| Opening/establishing | "ambient, atmospheric, slowly building, sparse piano, wide reverb" |
| Rising tension | "building tension, accelerating tempo, minor key, strings crescendo, pulsing bass" |
| Climax/impact | "epic orchestral hit, full crescendo, dramatic percussion, powerful brass" |
| Resolution/calm | "gentle resolution, major key resolve, soft strings, peaceful, fading" |
| Montage/progress | "upbeat, forward momentum, driving rhythm, positive energy, building layers" |
| Suspense | "low drone, dissonant undertones, sparse percussion, unsettling, minimal" |
| Joy/celebration | "bright major key, uplifting melody, claps, warm brass, celebratory" |
| Sadness/loss | "minor key piano, slow tempo, cello, intimate, melancholic, sparse arrangement" |

**Silence as a tool:** Dropping all music for 2-3s before a key moment creates massive impact when sound returns. Use this before reveals, climaxes, or emotional beats.

---

## Genre Templates

### Commercial / Ad (15-30s)

**Structure:** Hook (3s) --> Problem/Desire (5s) --> Product Intro (5s) --> Product in Action (7s) --> CTA/Logo (5s)

**5-shot, 25s Product Ad Template:**

```
SHOT 1 — CU — 3s — HOOK
[Camera: Static, slight push-in]
[Prompt: Extreme close-up of [problem visual], dramatic lighting,
  shallow depth of field, tension, premium commercial quality]
[Audio: Single sharp SFX to grab attention]

SHOT 2 — MS — 5s — DESIRE
[Camera: Slow dolly right]
[Prompt: Medium shot of [person experiencing the problem/desire],
  warm natural lighting, relatable setting, aspirational but authentic]
[Audio: Soft music begins, ambient environment sound]

SHOT 3 — MCU — 5s — PRODUCT REVEAL
[Camera: Slow push-in with slight rack focus]
[Prompt: [Product] being revealed/unboxed/held up, beautiful product
  lighting, clean background, premium materials visible, hero shot]
[Audio: Music lifts slightly, subtle "reveal" SFX]

SHOT 4 — MS — 7s — PRODUCT IN USE
[Camera: Tracking shot following action]
[Prompt: [Person] using [product] in [aspirational context],
  golden hour backlight, lifestyle commercial, showing key benefit]
[Audio: Music at full warmth, ambient environment]

SHOT 5 — STATIC — 5s — CTA
[Camera: Static, centered]
[Prompt: [Product] hero shot centered on clean background,
  brand colors, premium lighting, packshot commercial photography]
[Audio: Music resolves to final chord, silence before tagline]
```

### Short Film (1-3 min)

**Structure:** World establish (10s) --> Character intro (10s) --> Inciting incident (10s) --> Rising action (30s) --> Midpoint shift (10s) --> Complications (20s) --> Climax (15s) --> Resolution (15s)

**Key principles:**
- Compress ruthlessly. Every shot must advance plot or character.
- One character, one conflict, one resolution. No subplots.
- Visual storytelling over dialogue. Show, never tell.
- End on a powerful final image that lingers.
- Use the I2V storyboard pipeline for maximum visual control.

**10-shot, 90s Short Film Template:**

```
SHOT 1 — EWS — 8s — WORLD
Slow aerial establishing the world. No characters yet. Set tone with
lighting and environment. Use music to establish emotional key.

SHOT 2 — MS — 6s — CHARACTER INTRO
First appearance of protagonist. Reveal through action, not exposition.
What are they doing? This tells us who they are.

SHOT 3 — CU — 5s — CHARACTER DETAIL
Insert shot: hands working, eyes looking, a personal object. Adds depth
and intimacy. Static camera, shallow depth of field.

SHOT 4 — MS — 8s — INCITING INCIDENT
Something changes. A discovery, arrival, disruption. Camera movement
shifts from calm to dynamic to mirror the narrative shift.

SHOT 5 — MCU — 7s — REACTION
Character's emotional response to the incident. This is the most
important acting beat. Hold on the face. Let the audience feel it.

SHOT 6 — WS — 10s — RISING ACTION
Character takes action in response. Wider shot to show them moving
through the world with new purpose. Tracking or following camera.

SHOT 7 — MS — 8s — COMPLICATION
The plan hits a wall. New obstacle, reversal, or cost revealed.
Camera work becomes more urgent. Music tension increases.

SHOT 8 — CU — 6s — DECISION MOMENT
Close on the character as they make the critical choice. Minimal
movement. This is the emotional climax. Music drops to silence.

SHOT 9 — WS/MS — 12s — CLIMAX ACTION
The choice plays out. Longest shot. Can be dynamic or deliberately
slow depending on genre. Full music returns. Decisive motion.

SHOT 10 — EWS — 10s — FINAL IMAGE
Pull back to wide. Mirror the opening establishing shot but with
something changed. The world is different now. Music resolves.
Hold this frame. Let it breathe.
```

### Music Video

**Structure:** Beat-synced cuts aligned to BPM. Alternate between performance shots and narrative shots. Visual motifs repeat and evolve.

**Key principles:**
- Calculate shot timing from BPM: at 120 BPM, beats land every 0.5s. Common cut points: every 2 beats (1s), every 4 beats (2s), every 8 beats (4s).
- Verse: narrative/storytelling shots, slower pacing, 4-8 beat cuts
- Chorus: performance/energy shots, faster pacing, 2-4 beat cuts
- Bridge: visual shift, different location/color/style, longest held shots
- Establish 2-3 recurring visual motifs that evolve (e.g., flower blooming progressively, light changing from cool to warm)

### Documentary Style

**Structure:** Opening hook (arresting image or statement) --> Context (establishing shots, narration) --> Evidence (B-roll sequences) --> Testimonial/Interview (talking head) --> Resolution (synthesis, final thought)

**4-shot Documentary Opening Template:**

```
SHOT 1 — EWS — 6s — ARRESTING IMAGE
[Prompt: Breathtaking or provocative wide shot of the documentary
  subject, natural lighting, raw authenticity, 16mm film grain,
  handheld micro-movement, documentary cinematography]
[Audio: Natural ambient sound only, no music]

SHOT 2 — MS — 5s — CONTEXT
[Prompt: Medium shot establishing the environment/community/situation,
  observational camera, natural available light, candid framing,
  documentary realism]
[Audio: Narrator begins: "In [place], [context statement]..."]

SHOT 3 — CU — 5s — HUMAN ELEMENT
[Prompt: Close-up of a person's weathered hands working, or eyes
  looking into distance, shallow depth of field, intimate, textured,
  natural skin detail, documentary portrait]
[Audio: Narration continues, ambient under]

SHOT 4 — WS — 6s — SCALE/IMPACT
[Prompt: Wide shot revealing the full scope of the subject, drone or
  elevated angle, golden hour or overcast natural light, showing
  scale and impact, documentary aerial]
[Audio: Music enters softly, narration pauses to let image speak]
```

### Social Media Story (15-60s, Vertical)

**Structure:** Hook (first 2s) --> Content payload (8-40s) --> CTA/payoff (3-5s)

**The hook is everything.** If the first 2s do not arrest scrolling, the rest does not matter. Start with the most visually striking or emotionally provocative shot.

**3-shot, 15s Vertical Hook Template:**
```json
{
  "multi_prompt": [
    {
      "prompt": "Extreme close-up dramatic reveal of [striking subject], camera pulling back to reveal context, bold saturated colors, high contrast, vertical 9:16 framing, social media content, punchy and immediate",
      "duration": 4
    },
    {
      "prompt": "Dynamic medium shot of [subject in action], energetic camera movement, bright vibrant lighting, fast-paced, engaging vertical composition, contemporary style",
      "duration": 6
    },
    {
      "prompt": "[Final payoff shot: result, transformation, or punchline], satisfying conclusion, clean framing, bold colors, memorable final image, vertical format",
      "duration": 5
    }
  ],
  "shot_type": "customize",
  "aspect_ratio": "9:16"
}
```

### Product Launch

**Structure:** Tease (dark/mysterious) --> Reveal (dramatic unveil) --> Features (detail shots) --> Aspiration (lifestyle context) --> Brand close

**5-shot Product Launch Template:**
```
SHOT 1 — CU — 4s — TEASE
Dark frame with a single light source illuminating just an edge or
silhouette of the product. Mysterious. Anticipation. Slow push-in.

SHOT 2 — MS — 5s — REVEAL
Dramatic lighting change reveals the full product. Slow rotation or
camera orbit. Hero product shot. Premium materials, clean background.

SHOT 3 — ECU — 4s — FEATURE DETAIL
Macro close-up of the key differentiating feature. Texture, material,
mechanism. Shallow DOF. Product photography lighting.

SHOT 4 — WS — 5s — ASPIRATION
Product in its ideal context. Lifestyle shot. Beautiful environment.
Person using it effortlessly. Golden hour. Aspirational.

SHOT 5 — STATIC — 4s — BRAND CLOSE
Product centered, brand colors, clean. Logo appears. Premium.
Music resolves to final note.
```

### Tutorial / How-To

**Structure:** Problem statement (show the pain) --> Step 1 --> Step 2 --> Step 3 --> Result (transformation)

**Key principles:**
- Show, do not explain. Visual clarity over verbal complexity.
- Each step is one shot. One action per shot, clearly framed.
- Use insert close-ups for critical details.
- Before/after bookend shots for maximum impact.

### Brand Story

**Structure:** Origin (where it started) --> Challenge (what stood in the way) --> Breakthrough (the turning point) --> Impact (what changed) --> Values (what drives us)

**8-shot, 60s Brand Story Template:**
```
SHOT 1 — EWS — 6s — ORIGIN
Wide shot of humble beginnings. A garage, workshop, small room.
Warm nostalgic lighting. Grain. Intimate scale.

SHOT 2 — CU — 5s — FOUNDER
Close-up of hands working. Craft, dedication, attention to detail.
The human behind the brand.

SHOT 3 — MS — 5s — EARLY DAYS
Medium shot of early work. Small team or solo effort. Natural light.
Documentary feel. Authentic.

SHOT 4 — WS — 6s — CHALLENGE
Visual metaphor for obstacle. Storm, wall, vastness, complexity.
Tone shifts darker. Music builds tension.

SHOT 5 — MCU — 5s — DETERMINATION
Character reaction. Resolve. The moment before the breakthrough.
Lighting transitioning from shadow to light.

SHOT 6 — MS — 7s — BREAKTHROUGH
The turning point. Action, creation, discovery. Dynamic camera.
Energy shifts positive. Music lifts.

SHOT 7 — WS — 8s — IMPACT
Wide shot showing the result at scale. What the brand built.
Community, product in the world, transformation. Epic scope.

SHOT 8 — STATIC — 5s — VALUES
Final image that embodies the brand's core value. Simple, powerful,
iconic. Logo. Music resolves.
```

---

## Transition Techniques

### Visual Transitions

| Transition | When to Use | Prompt Technique |
|-----------|-------------|-----------------|
| Hard cut | Same scene, different angle. Energy. Urgency. | Default between shots. No special handling needed. |
| Dissolve | Time passage, memory, dream, gentle scene change. | End Shot N and begin Shot N+1 with overlapping visual elements (same color, same shape). |
| Fade to black | End of act, major time skip, finality. | End Shot N with subject moving into shadow or darkness. Begin Shot N+1 from black. |
| Whip pan | Energy, surprise, comedy, fast location change. | End Shot N with fast camera pan blur. Begin Shot N+1 with matching directional blur settling. |
| Match cut | Thematic connection between two different subjects. | End Shot N on a specific shape/movement. Begin Shot N+1 on the same shape/movement in a new context. |
| Zoom transition | Modern, social media, dynamic energy. | End Shot N pushing into a detail. Begin Shot N+1 pulling back from a similar detail in a new scene. |
| Light flash | Supernatural, memory, time travel, impact. | End Shot N with bright light filling frame. Begin Shot N+1 from white fading to new scene. |

### Narrative Transitions

**Time skip:** Use fade to black + establishing shot of new time. Change lighting (day to night, season change) to signal temporal shift visually.

**Location change:** Cut to establishing wide of new location. Or use match cut on a shared visual element between locations.

**Flashback/forward:** Shift color grade dramatically. Flashbacks: warmer, more grain, softer focus. Flash-forwards: cooler, sharper, slightly desaturated.

**Parallel action:** Alternate between two simultaneous storylines with matched pacing. Use similar shot types (wide-wide, close-close) to create rhythmic parallel.

---

## Pacing and Rhythm

### Fast Pacing (Action, Excitement, Comedy)

- Shot duration: 2-4s
- Camera: Dynamic movement, handheld energy, whip pans, quick zooms
- Cuts: Hard cuts, every 2-3s
- Music: High BPM (120-160+), percussive, driving rhythm
- Motion prompt keywords: "fast motion, dynamic, energetic, quick, rapid, explosive"

### Slow Pacing (Drama, Emotion, Horror, Luxury)

- Shot duration: 5-10s
- Camera: Static or very slow dolly/push-in, steady, deliberate
- Cuts: Minimal, let shots breathe
- Music: Low BPM (60-90), atmospheric, sparse, resonant
- Motion prompt keywords: "slow motion, deliberate, graceful, measured, contemplative"

### Variable Pacing (The Professional Pattern)

The most compelling narratives vary their pacing. The fundamental pattern:

```
Slow (establish) → Medium (develop) → Fast (intensify) → IMPACT → Slow (breathe)
```

Repeat this cycle for each major narrative beat. The "breathe" moment after each impact is critical. It gives the audience time to process and resets their sensitivity for the next build.

**Example pacing map for a 60s brand story:**
```
0-10s:  SLOW  — Establishing, world-building, ambient music
10-20s: MEDIUM — Character action, rising music, purpose emerges
20-35s: FAST   — Montage, quick cuts, energy building, music driving
35-40s: IMPACT — Climax moment, biggest visual, music peak
40-50s: SLOW   — Resolution, emotional landing, music softens
50-60s: SLOW   — Final image, brand, held shot, music resolves
```

---

## Visual Continuity Checklist

When generating multiple shots for a narrative, check every item before finalizing.

- [ ] **Lighting direction** — Is the key light coming from the same side in all shots of the same scene? Inconsistent light direction breaks immersion instantly.
- [ ] **Color grade** — Are all shots graded to the same palette? Use the same style anchor prefix in every prompt.
- [ ] **Character appearance** — Same hair, same clothes, same skin tone, same accessories across all shots. Use character lock descriptions.
- [ ] **Screen direction** — If a character is moving left-to-right in one shot, they should continue left-to-right in the next (180-degree rule). Reversing direction implies they turned around.
- [ ] **Time of day** — Golden hour does not become midday between cuts. Lock the lighting description: "low-angle warm golden hour sunlight" in every shot of the same scene.
- [ ] **Weather/atmosphere** — Rain, fog, snow, clear sky must be consistent within a scene. If it is raining in the wide shot, it must be raining in the close-up.
- [ ] **Background elements** — If a mountain is visible behind the character in the wide shot, it should still be there in tighter shots from the same angle.
- [ ] **Aspect ratio** — All shots in a sequence must use the same aspect ratio.
- [ ] **Film stock/grain** — If using a film look, apply it to every shot. Mixing clean digital and grainy film breaks cohesion.
- [ ] **Camera era/lens** — If using anamorphic lens characteristics (oval bokeh, flare), use them in every shot.

---

## Complete Story Examples with Full Prompts

### Example 1: 30-Second Coffee Brand Ad (Native Multi-Prompt)

```json
{
  "multi_prompt": [
    {
      "prompt": "Extreme close-up of a single coffee bean falling in slow motion against a deep black background, the bean rotating slowly to reveal its texture and oil sheen, dramatic single spotlight from above, luxury product commercial, macro photography, shallow depth of field",
      "duration": 4
    },
    {
      "prompt": "Medium shot of coffee beans cascading into a copper hand grinder, warm amber side lighting, rustic wooden countertop, morning sunlight streaming through a window, artisanal coffee preparation, slow motion pour, warm color grading",
      "duration": 4
    },
    {
      "prompt": "Close-up of a hand slowly turning the grinder handle, fresh ground coffee falling into the drawer below, visible aromatic steam rising, warm golden morning light, tactile craft, shallow depth of field on the grounds, artisanal detail",
      "duration": 5
    },
    {
      "prompt": "Medium shot of hot water from a gooseneck kettle being poured in a slow circular motion over coffee grounds in a ceramic pour-over dripper, steam rising beautifully, morning kitchen, warm tones, slow motion, satisfying pour pattern",
      "duration": 5
    },
    {
      "prompt": "Close-up of dark rich coffee streaming into a white ceramic cup, the last drops creating ripple rings on the surface, warm amber backlight creating a glow through the steam, beautiful dark liquid, premium coffee brand commercial, 4K detail",
      "duration": 4
    },
    {
      "prompt": "Medium shot of the finished cup of coffee on the wooden counter beside the grinder and bag of beans, soft morning light, steam gently rising, cozy warm atmosphere, hero product shot, premium artisanal brand, clean composition with negative space for logo",
      "duration": 5
    }
  ],
  "shot_type": "customize",
  "generate_audio": true,
  "aspect_ratio": "16:9"
}
```

### Example 2: 60-Second Brand Story — Architecture Firm (I2V Pipeline)

This uses the storyboard-to-video pipeline for maximum visual control.

**Style anchor (use in every image and video prompt):**
```
Cinematic architectural photography, clean modern aesthetic, natural
light, concrete and glass, cool neutral tones with warm accent light,
2.39:1 anamorphic, shallow depth of field
```

```
SHOT 1 — EWS — 7s
[Image: Aerial view of a vast empty plot of land at dawn, single
  architect figure standing in the center with rolled blueprints,
  construction site stakes marking the perimeter, misty morning,
  cinematic architectural photography, cool neutral tones]
[I2V model: an image-to-video model (use search_models)]
[Motion: Slow aerial descent toward the figure, morning mist drifting,
  wind moving the architect's jacket slightly]
[Audio: Ambient dawn, distant birds, soft atmospheric pad beginning]

SHOT 2 — CU — 5s
[Image: Close-up of architect's hands unrolling a large blueprint on
  a temporary table, pencil marks and measurements visible, morning
  light raking across the paper surface, shallow depth of field]
[I2V model: an image-to-video model (use search_models)]
[Motion: Hands smoothing the blueprint flat, finger tracing a line,
  paper rustling slightly in the breeze]
[Audio: Paper rustle, pencil tap, music building slowly]

SHOT 3 — MS — 6s
[Image: Medium shot of the architect gesturing toward the empty site
  while speaking to a small construction team, hard hats, early morning
  light creating long shadows, documentary feel]
[I2V model: an image-to-video model (use search_models)]
[Motion: Architect pointing, team nodding, natural body language,
  camera slowly dollying right]
[Audio: Music rising, narration: "Every structure begins as an idea"]

SHOT 4 — MONTAGE — 8s (2 sub-shots at 4s each)
[Image A: Steel beams being lifted by crane against a blue sky,
  construction in progress, dramatic upward angle, industrial scale]
[Image B: Close-up of a worker welding, sparks flying, protective
  mask, dramatic chiaroscuro lighting from the welding arc]
[I2V model: an image-to-video model (use search_models) for both]
[Motion A: Crane lifting, beam swinging slowly into position]
[Motion B: Welding sparks cascading, subtle camera push-in]
[Audio: Construction ambient, metal clangs, music driving forward]

SHOT 5 — WS — 7s
[Image: Wide shot of the now-complete building at golden hour, stunning
  modern architecture, glass facade reflecting sunset, landscaped
  surroundings, people walking in the entrance plaza, aspirational]
[I2V model: an image-to-video model (use search_models)]
[Motion: Slow dolly forward toward the building entrance, people
  walking naturally, golden light shifting on glass facade]
[Audio: Music at emotional peak, narration: "Built to endure"]

SHOT 6 — MS — 6s
[Image: Medium shot looking through the glass facade into the building
  interior, warm interior lights contrasting with blue hour exterior,
  silhouettes of people inside living and working, beautiful reflection]
[I2V model: an image-to-video model (use search_models)]
[Motion: Camera slowly pulling back, interior lights warm and alive,
  exterior transitioning to evening blue]
[Audio: Music resolving warmly, ambient evening, narration: firm name]
```

### Example 3: Music Video Sequence (6 shots, beat-synced)

**BPM: 100 (beat = 0.6s). Cut on every 8 beats (4.8s, rounded to 5s).**

```
SHOT 1 — EWS — 5s (Verse 1 start)
[Prompt: Wide shot of an empty desert highway stretching to the horizon
  at blue hour, a single vintage red Mustang parked on the shoulder,
  headlights glowing, dust in the air, Americana cinematic, Roger
  Deakins natural lighting, 35mm film grain, 2.39:1]
[Audio: Song verse 1 begins, acoustic guitar intro]

SHOT 2 — MCU — 5s (Verse 1 continue)
[Prompt: Medium close-up of a young woman sitting on the hood of the
  red Mustang, cowboy boots, denim jacket, looking at the horizon,
  blue hour light on her face, wind moving her hair, contemplative
  expression, shallow depth of field, warm film tones, 35mm grain]
[Audio: Verse 1 continues, vocal enters]

SHOT 3 — CU — 5s (Pre-chorus)
[Prompt: Close-up of the woman's hand turning the car ignition key,
  dashboard lights flickering on, warm amber glow on her fingers,
  vintage car interior detail, shallow depth of field, cinematic grain]
[Audio: Pre-chorus builds, drums enter, bass rises]

SHOT 4 — WS — 5s (Chorus)
[Prompt: Wide tracking shot of the red Mustang speeding down the desert
  highway at golden hour, dust trail behind it, sun low on the horizon
  creating dramatic lens flare, dynamic motion, freedom and speed,
  cinematic aerial tracking shot, 35mm Kodak warmth]
[Audio: Chorus full energy, all instruments, driving rhythm]

SHOT 5 — MS — 5s (Chorus continue)
[Prompt: Medium shot from inside the car, woman driving with windows
  down, hair blowing, sun visor shadow across her face, one hand on
  the wheel, desert landscape blurring past in the window, warm golden
  light flooding the interior, 35mm film, joy and freedom]
[Audio: Chorus continues, vocal peak, instruments full]

SHOT 6 — EWS — 5s (Chorus end / Instrumental)
[Prompt: Extreme wide aerial shot following the red Mustang as a tiny
  dot on the vast desert highway, the road curving toward distant
  mountains, sunset painting the sky in orange and purple gradients,
  epic cinematic scale, drone pullback, 2.39:1 anamorphic]
[Audio: Chorus resolves, instruments pull back, atmospheric outro]
```

### Example 4: Documentary Opening — Ocean Conservation (4 shots)

```
SHOT 1 — EWS — 7s
[Image prompt: Breathtaking aerial shot of a turquoise tropical ocean
  with a dark coral reef visible beneath the surface, morning light
  creating jewel-like reflections, the scale of the reef visible from
  altitude, documentary nature cinematography, 16mm film texture]
[I2V: an image-to-video model (use search_models)]
[Motion: Slow aerial glide over the reef, gentle ocean surface movement,
  light shifting on water]
[Audio: Ocean ambient only, no music. Pure natural sound.]

SHOT 2 — UW/MS — 6s
[Image prompt: Underwater medium shot of vibrant coral reef teeming
  with tropical fish, sunlight filtering through water creating god
  rays, blue water, colorful coral and anemones, National Geographic
  underwater photography, natural available light]
[I2V: an image-to-video model (use search_models)]
[Motion: Camera slowly gliding through the reef, fish swimming naturally,
  light rays shifting through water]
[Audio: Muffled underwater ambient, gentle bubbles, narrator begins:
  "Beneath the surface, a world fights for survival."]

SHOT 3 — CU — 5s
[Image prompt: Close-up of a sea turtle's ancient face looking directly
  at camera, gentle wise eyes, barnacles on shell, ocean blue-green
  background blurred, shallow depth of field, intimate wildlife
  portrait, 16mm texture, natural underwater light]
[I2V: an image-to-video model (use search_models)]
[Motion: Turtle turns head slowly, blinks once, subtle water current
  moving around the shell, intimate eye contact with lens]
[Audio: Narration continues, subtle music enters — low strings, sparse]

SHOT 4 — EWS — 6s
[Image prompt: Wide shot split between underwater coral reef below and
  a grey industrial skyline on the horizon above, the juxtaposition
  of nature and industry, half-underwater half-above composition,
  overcast moody light on the surface, vivid life below, documentary]
[I2V: an image-to-video model (use search_models)]
[Motion: Slow pull back revealing the split, waves lapping at the
  surface dividing line, smoke from distant industrial stacks]
[Audio: Music shifts minor key, tension, narration: "But for how long?"]
```

### Example 5: Social Media Hook — Cooking (3 shots, vertical)

```json
{
  "multi_prompt": [
    {
      "prompt": "Extreme close-up of a knife slicing through a ripe mango in slow motion, juice bursting from the cut, vibrant orange flesh against dark cutting board, dramatic overhead lighting, food photography, 9:16 vertical, satisfying ASMR visual, 4K macro detail",
      "duration": 4
    },
    {
      "prompt": "Top-down overhead shot of colorful tropical ingredients being arranged on a dark plate, mango slices, passion fruit halves, lime wedges, coconut shavings, mint leaves, hands working quickly, bright studio lighting, food styling, vertical format",
      "duration": 6
    },
    {
      "prompt": "Medium shot of the completed tropical fruit bowl being placed on a sunlit marble counter, condensation on a glass of juice beside it, morning light creating beautiful highlights on the fruit, final hero food shot, vibrant colors, appetizing, vertical 9:16",
      "duration": 5
    }
  ],
  "shot_type": "customize",
  "aspect_ratio": "9:16"
}
```

### Example 6: Cinematic Short Scene — Reunion (I2V Pipeline, 6 shots)

**Style anchor:**
```
Independent drama, 35mm Kodak Portra 400, warm naturalistic lighting,
shallow depth of field, gentle handheld, intimate, emotional, 2.39:1
```

```
SHOT 1 — WS — 8s
[Prompt: Wide shot of a small rural train station platform at late
  afternoon, golden dust motes in slanted sunlight, wooden bench,
  one person sitting and waiting, empty tracks stretching into the
  distance, 35mm Kodak Portra warm tones, lonely atmosphere]
[Motion: Very slow push-in toward the waiting figure, dust drifting,
  tree leaves rustling in gentle wind]
[Audio: Distant birds, wind through dry grass, train tracks creaking.
  No music yet.]

SHOT 2 — MCU — 6s
[Prompt: Medium close-up of a middle-aged woman on the bench, worn
  hands folded in her lap, sun-weathered face, nervous anticipation
  in her eyes, warm golden sidelight, shallow DOF, 35mm grain,
  Portra tones, emotional intimacy]
[Motion: Static camera, micro-handheld movement, woman glances down
  the tracks, swallows, looks down at her hands, wind moves her hair]
[Audio: Quiet ambient, a distant train horn. Music: single sustained
  piano note begins.]

SHOT 3 — WS — 5s
[Prompt: Wide shot down the train tracks, a train approaching in the
  heat shimmer distance, the platform waiting figure small in the
  right third of frame, golden hour light, dust and atmosphere,
  anticipation, cinematic rail shot, 35mm warmth]
[Motion: Train grows larger, approaching. Camera static, grounded.
  Heat shimmer ripples. Woman stands from bench.]
[Audio: Train approaching, growing louder. Piano joined by strings.]

SHOT 4 — MS — 7s
[Prompt: Medium shot of the train doors opening, passengers stepping
  down, the waiting woman scanning faces with desperate hope, golden
  backlight from the setting sun behind the train, shallow DOF with
  passengers moving through focus, emotional searching, 35mm grain]
[Motion: People walking past the camera, the woman's eyes tracking,
  searching, rack focus between faces and her expression]
[Audio: Train idle, footsteps, voices. Music building with cello.]

SHOT 5 — CU — 6s
[Prompt: Close-up of the woman's face as recognition and emotion hit,
  eyes filling with tears, mouth trembling into a smile, golden
  sidelight catching a tear, incredible emotional detail, shallow
  depth of field, 35mm Portra, the most important frame in the story]
[Motion: Nearly static. A single tear falls. Her smile breaks through.
  Barely perceptible nod. Golden light catches the tear.]
[Audio: Music swells — full strings, piano, emotional crescendo]

SHOT 6 — WS — 10s
[Prompt: Wide shot of two people embracing on the empty train platform,
  the train pulling away in the background, late golden hour light
  creating long shadows, dust settling, the world around them fading
  to warm blur, 35mm Kodak Portra, final image of reunion, intimate
  figures in vast landscape, beautiful and simple, 2.39:1]
[Motion: Figures holding each other, slight rocking, train slowly
  departing behind them, dust and light particles drifting, camera
  very slowly pulling back to show the empty platform around them]
[Audio: Music at emotional peak, then slowly fading. Train departing.
  Final moments: just wind and the sound of the empty station.]
```

---

## Single-Prompt Multi-Shot Sequences

Some video models handle multi-shot sequences through detailed prompt descriptions that include shot transitions within a single generation. Describe the full sequence in one prompt, including camera changes and scene shifts. Use `search_models` to find text-to-video models with this capability.

**Example prompt for single-prompt multi-shot:**
```
A cinematic sequence: First, an aerial wide shot descends over a foggy
mountain forest at dawn. Cut to a close-up of a deer lifting its head
from a stream, water droplets catching morning light. Cut to a medium
shot tracking through the forest as the deer runs gracefully between
trees, dappled sunlight. Final shot pulls back to an extreme wide as
the deer emerges at the edge of a cliff overlooking a vast misty valley.
Natural documentary lighting, 4K cinematic, atmospheric.
```

---

## Production Workflow Summary

### From Concept to Final Output

**Phase 1: Story Design (no generation yet)**
1. Define the story in one sentence: who, what, where, why.
2. Choose a genre template from the section above.
3. Decide total duration and shot count.
4. Write the three-act beat sheet (3-5 bullet points per act).
5. Choose aspect ratio (16:9 for cinematic, 9:16 for social, 1:1 for square).

**Phase 2: Visual Script**
6. Write out every shot using the visual script format.
7. For each shot, specify: shot type, duration, camera, prompt, audio layer notes.
8. Run the visual continuity checklist across all shots.
9. Decide: native multi-shot (use `search_models` to find models with multi-shot support) or I2V pipeline.

**Phase 3: Image Generation (I2V pipeline only)**
10. Generate key frames for every shot using an image generation model (use `search_models`).
11. Iterate each frame until composition, character, and lighting are correct.
12. Ensure character and style consistency across all frames before proceeding.

**Phase 4: Video Generation**
13. For native multi-shot: Submit the full multi_prompt or sequence prompt.
14. For I2V pipeline: Animate each key frame with the appropriate I2V model.
15. For first/last frame shots: Generate both frames, submit to interpolation endpoint.
16. For reference-to-video shots: Provide reference images for character consistency.

**Phase 5: Audio Production**
17. Generate music track matching the emotional arc (use a music generation model via `search_models`).
18. Generate narration/dialogue if needed (use a TTS model via `search_models`).
19. Generate SFX for specific moments (use a sound effects model via `search_models`).
20. Generate ambient beds for environmental grounding (use a text-to-audio model via `search_models`).

**Phase 6: Assembly**
21. Sequence all video shots in order.
22. Layer audio: ambient bed, music, SFX, narration (from bottom to top).
23. Time narration to shot transitions and key visual moments.
24. Review the full sequence for pacing, continuity, and emotional arc.
25. Export final output.

### Model Selection Decision Tree

```
Need multi-shot in one call?
├── Yes, with audio → a multi-prompt video model with generate_audio support
├── Yes, character consistency critical → a dedicated multi-shot narrative model
├── Yes, complex sequence → a model supporting full sequence in one prompt
└── No, maximum control per shot → I2V Pipeline
    ├── Need start+end control → a first-last-frame-to-video model
    ├── Need character consistency → a reference-to-video model
    ├── Need highest quality → a top-tier I2V model
    ├── Need fast iteration → a fast I2V model
    └── Need stylized/animated look → a stylized I2V model

Use search_models to find the best model for each capability.
```

### Audio Model Selection

```
Need music?
├── Full song/score → a music generation model (search_models)
└── Short ambient/texture → a text-to-audio model (search_models)

Need voice?
├── Single narrator, high quality → a high-quality TTS model (search_models)
└── Multiple speakers/characters → a multi-speaker TTS model (search_models)

Need sound effects?
└── All SFX → a sound effects model (search_models)

Need synced audio with video?
└── A video model with generate_audio support (search_models)
```
