# Motion Design — Motion Graphics & Animation Master Guide

Complete reference for AI-powered motion graphics, kinetic typography, logo animation,
and visual effects on fal.ai (March 2026). Every model, technique, and prompt in this
document is verified and actionable.

---

## Quick Reference

> **Model Discovery:** Use `search_models` to find the best model for each task. Search by category: "text to video" for kinetic typography and animation, "image to video" for logo reveals, "image generation" for static key-frame design, "text to image with text" for text-accurate frames.

**Top 5 Motion Design Tips**

1. Use slow motion — it dramatically reduces AI artifacts and looks more intentional.
2. Describe the background first, then the subject, then the motion. Order signals hierarchy.
3. Always specify easing: `eases into position`, `smoothly decelerates`, never leave motion speed implicit.
4. For logo reveals, use the two-step workflow: generate a static frame first, then animate with I2V.
5. Generate 3-5 variations per prompt — motion AI has high variance; pick the best, do not over-refine.

**Key Prompt Patterns**

*Kinetic Typography:* `Clean motion graphics on [BACKGROUND]. [FONT STYLE] text "[TEXT]" [ENTRANCE STYLE]. [EASING]. Minimal, professional kinetic typography.`

*Logo Reveal (I2V):* `[ASSEMBLY MECHANIC] into the logo shape against [BACKGROUND]. [LIGHTING]. [FINAL STATE]. Professional motion graphics.`

*Transition:* `[SCENE A] [TRANSITION MECHANIC] into [SCENE B]. Seamless, continuous motion. [SPEED/MOOD].`

*Abstract Loop:* `[PATTERN/FORM] in [COLOR PALETTE]. [MOTION BEHAVIOR]. Seamless looping animation. [STYLE ANCHOR].`

**Cheat Sheet — Prompt Cores by Effect**

- Bold announcement text: "Bold white sans-serif text scales up from center, speed ramp, explosive expansion, black background"
- Logo particle assembly: "Logo assembles from glowing particles, final flash as logo solidifies, dark background" (use I2V model)
- Glitch reveal: "RGB splitting, pixel fragmentation, gradually resolve to reveal [TEXT/LOGO]"
- Seamless abstract loop: "Slow organic color fields, gentle motion, seamless looping, meditative"

---

## Motion Graphics Fundamentals

Motion design is the art of giving graphic elements life through movement. In AI video
generation, the quality of your motion description directly determines the quality of
the output. This section maps classical animation principles to AI-prompt vocabulary.

### The 12 Principles of Motion (Adapted for AI Prompts)

Each principle below includes the exact keywords that AI video models respond to.

1. **Anticipation** — A small pull-back before the main action. Prompt keywords:
   `builds up`, `pulls back before`, `slight pause then`, `gathers energy`,
   `coils before launching`. Use this for logo reveals and text entrances.

2. **Follow-through & Overshoot** — Elements continue past their target then settle.
   Keywords: `overshoots and settles`, `bounces into place`, `elastic snap`,
   `wobbles to rest`, `momentum carry-through`. Essential for bouncing text.

3. **Ease in / Ease out** — Natural acceleration and deceleration. Keywords:
   `smoothly accelerates`, `gently decelerates`, `eases into position`,
   `soft landing`, `gradual start`. Avoid `linear motion` unless you want
   a mechanical feel.

4. **Squash & Stretch** — Organic deformation that adds weight and flexibility.
   Keywords: `elastic deformation`, `rubbery bounce`, `squash and stretch`,
   `organic squish`, `springy motion`. Works best with cartoon/playful styles.

5. **Secondary Motion** — Supporting elements that react to the primary action.
   Keywords: `particles trail behind`, `ripple effect spreads outward`,
   `debris scatters`, `light flares react`, `shadows follow`. Adds production value.

6. **Staging** — Clear visual hierarchy so the viewer knows where to look.
   Keywords: `centered composition`, `isolated on dark background`,
   `shallow depth of field draws focus`, `spotlight on subject`,
   `negative space surrounds`. Critical for title cards and logo reveals.

7. **Arcs** — Natural curved motion paths instead of straight lines.
   Keywords: `sweeps in an arc`, `curves gracefully`, `orbital path`,
   `follows a parabolic trajectory`, `arcing motion`. For swooping reveals.

8. **Timing** — The speed of action defines its weight and mood. Keywords:
   `slow and deliberate`, `rapid burst`, `measured pace`, `snappy timing`,
   `languid movement`. Slow motion reduces AI artifacts significantly.

9. **Exaggeration** — Pushing motion beyond realism for impact. Keywords:
   `dramatic`, `explosive`, `extreme close-up`, `impossibly fast`,
   `hyper-stylized motion`. Best for social media and attention-grabbing content.

10. **Solid Drawing / Dimensional Form** — 3D presence and volume. Keywords:
    `3D extruded`, `volumetric`, `dimensional depth`, `parallax layers`,
    `cast shadows reveal depth`. Gives flat graphics a cinematic feel.

11. **Appeal** — Visual attractiveness and clarity. Keywords:
    `clean design`, `polished`, `professional finish`, `satisfying motion`,
    `visually harmonious`. Combine with style anchors for brand alignment.

12. **Rhythm & Repetition** — Patterned motion that creates visual tempo.
    Keywords: `rhythmic pulse`, `repeating pattern`, `wave propagation`,
    `staggered timing`, `cascading sequence`. Perfect for kinetic typography.

### Motion Design Vocabulary — Master Keyword List

These keywords are understood by all major AI video models. Combine 3-5 per prompt.

**Speed & Timing:**
`slow motion`, `time lapse`, `speed ramp`, `bullet time`, `real-time`,
`hyper-speed`, `freeze frame`, `staccato rhythm`, `smooth continuous`,
`gradual acceleration`, `sudden burst`, `lingering pause`

**Direction & Path:**
`left to right`, `right to left`, `top to bottom`, `bottom-up`,
`diagonal sweep`, `radial outward`, `spiral inward`, `converging`,
`diverging`, `zigzag path`, `orbital rotation`, `pendulum swing`

**Scale & Transform:**
`scales up from zero`, `shrinks to nothing`, `grows organically`,
`expands outward`, `contracts inward`, `zoom punch`, `scale bounce`,
`size oscillation`, `progressive enlargement`, `micro to macro`

**Opacity & Visibility:**
`fades in`, `fades out`, `dissolve reveal`, `opacity pulse`,
`ghost trail`, `transparent layers`, `semi-transparent overlay`,
`materializes from nothing`, `gradually becomes visible`

**Rotation & Orientation:**
`360 rotation`, `half turn`, `spin clockwise`, `counter-clockwise spin`,
`tumbling`, `rolling`, `flipping`, `pivoting from corner`,
`axial rotation`, `barrel roll`, `gentle wobble`

**Deformation & Morph:**
`morphs into`, `liquid warp`, `distortion wave`, `elastic stretch`,
`ripple deformation`, `glitch displacement`, `pixel scatter`,
`smear effect`, `motion blur trails`, `shape-shifting`

---

## Kinetic Typography

Kinetic typography is text brought to life through motion. It is one of the
strongest capabilities of current AI video models because text + motion is
well-represented in training data from commercials, music videos, and social media.

### Text Animation Styles

#### 1. Word-by-Word Reveal
Words appear sequentially, building a sentence with dramatic timing.
Best for: quotes, lyrics, motivational content.

**Prompt template (T2V):**
```
Clean motion graphics on a solid black background. White bold sans-serif text
appears word by word in the center of the frame: "[YOUR TEXT]". Each word
fades in with a subtle scale-up animation, pausing briefly between words.
Minimal, professional kinetic typography. Smooth easing on every transition.
```
**Model:** Use `search_models` to find a text-to-video model suited for kinetic typography.

#### 2. Character-by-Character Typing
Letters appear one at a time like a typewriter or terminal cursor.
Best for: tech content, code reveals, hacker aesthetic.

**Prompt template (T2V):**
```
Dark terminal screen with a blinking green cursor. Monospace font characters
type out one by one: "[YOUR TEXT]". Each character appears with a subtle
digital flicker. Terminal green text on pure black background. Retro
computer typing animation with slight screen glow and scanline overlay.
```
**Model:** Use `search_models` to find a text-to-video model.

#### 3. Bouncing / Elastic Text
Text bounces into frame with playful overshoot and settle.
Best for: fun brands, kids content, social media hooks.

**Prompt template (T2V):**
```
Colorful motion graphics, bright gradient background. Bold rounded
sans-serif text "[YOUR TEXT]" bounces into frame from above, overshoots
downward with elastic squash and stretch, then bounces back to center
and settles. Playful, energetic kinetic typography with secondary
particle confetti burst on impact. Smooth 60fps animation.
```

#### 4. Scaling Impact Text
Text starts tiny and rapidly scales to fill the frame for emphasis.
Best for: announcements, countdowns, dramatic reveals.

**Prompt template (T2V):**
```
Dramatic motion graphics. On a dark cinematic background, large bold
white text "[YOUR TEXT]" rapidly scales up from a tiny point at the
center to fill the entire frame. Speed ramp — starts slow then
explosively fast. Screen shake on final scale. Lens flare and light
streaks accompany the expansion. High impact, cinematic typography.
```

#### 5. 3D Extruded Text Rotation
Dimensional text rotates to reveal depth and material.
Best for: titles, brand names, premium feel.

**Prompt template (T2V):**
```
3D metallic chrome text "[YOUR TEXT]" slowly rotates on its vertical
axis against a dark gradient background. The text has beveled edges,
reflective chrome material catching studio lighting. Soft reflections
and caustics play across the surface as it turns. Smooth continuous
rotation, professional 3D motion graphics. Dramatic rim lighting.
```

#### 6. Glitch Text
Digital corruption effect with RGB split, scan lines, and displacement.
Best for: tech brands, cyberpunk aesthetic, digital art.

**Prompt template (T2V):**
```
Digital glitch motion graphics on black background. Bold text
"[YOUR TEXT]" appears with heavy glitch distortion — RGB channel
splitting, horizontal scan line displacement, pixel fragmentation,
and brief static bursts. The text stabilizes momentarily then
glitches again. Cyberpunk color palette: cyan, magenta, electric blue.
VHS tracking error aesthetic with chromatic aberration.
```

#### 7. Liquid / Morphing Text
Text formed from or dissolving into liquid.
Best for: beauty brands, water/beverage, organic products.

**Prompt template (T2V):**
```
Liquid chrome material forms into the text "[YOUR TEXT]" against
a pure black background. The reflective metallic liquid pools,
rises, and sculpts itself letter by letter into crisp typography.
Surface tension ripples and mercury-like reflections. Smooth,
mesmerizing fluid dynamics. High-end luxury motion graphics feel.
```

#### 8. Neon Text Flicker
Neon sign that flickers to life, buzzes, and stabilizes.
Best for: nightlife, entertainment, retro, restaurants.

**Prompt template (T2V):**
```
Dark brick wall at night. A neon sign reading "[YOUR TEXT]" flickers
to life — first a dim glow, then buzzing flashes, individual letters
turning on with electric arcs, finally the full text illuminated in
bright [COLOR] neon with a warm ambient glow on the surrounding wall.
Realistic neon tube lighting with gas discharge flicker and bloom.
```

### Font Style Keywords for AI Prompts

AI models respond to typographic descriptions. Use these to control font appearance:

| Style | Keywords | Visual Result |
|-------|----------|---------------|
| Bold Sans-Serif | `bold sans-serif, Helvetica-style, clean geometric` | Modern, corporate, strong |
| Serif | `elegant serif, Times-style, classical typography` | Editorial, literary, traditional |
| Handwritten | `handwritten script, brush lettering, calligraphic` | Personal, artistic, organic |
| Monospace | `monospace, terminal font, code typeface` | Technical, digital, developer |
| Display/Impact | `ultra-bold display font, impact typography, heavy weight` | Attention-grabbing, posters |
| Slab Serif | `slab serif, mechanical typeface, robust letters` | Industrial, vintage, sturdy |
| Condensed | `tall condensed font, narrow letterforms, vertical emphasis` | Space-efficient, editorial |
| Rounded | `rounded sans-serif, soft edges, friendly typeface` | Approachable, playful, apps |

### Text on Backgrounds — Readability Combos

| Background | Text Color | Prompt Addition |
|------------|-----------|-----------------|
| Solid black | White / neon | `on a pure black background` |
| Dark gradient | White / gold | `on a dark gradient from navy to black` |
| Textured concrete | White with shadow | `on a rough concrete texture with drop shadow` |
| Blurred video | White with outline | `overlaid on a soft bokeh background` |
| Color gradient | Dark text | `on a smooth pastel gradient from pink to orange` |

---

## Logo Animation

Logo animation takes a static logo and gives it a memorable entrance. The recommended
workflow is: create or upload the logo as an image, then use Image-to-Video (I2V) to
animate it. This gives you full control over the logo's exact appearance.

### Two-Step Logo Animation Workflow

**Step 1:** Generate or prepare a static logo frame.
- Use `search_models` to find a design-focused image generation model for creating logos
- Use `search_models` to find a background removal model to isolate an existing logo
- Place the logo centered on the target background color
- For logos containing specific text, find a model with strong text rendering capability

**Step 2:** Animate with I2V.
- Use `search_models` to find an image-to-video model. Look for options that offer high quality, reliable results, or stylized motion depending on your needs.

### Logo Reveal Styles & Prompt Templates

#### 1. Particle Assembly
Thousands of particles converge to form the logo.

**I2V Prompt (applied to static logo image):**
```
The logo is initially scattered as thousands of glowing particles
floating in dark space. The particles stream inward, converging and
assembling into the final logo shape. Each particle leaves a faint
light trail. The assembly accelerates as it completes, with a final
flash of light as the logo solidifies. Cinematic particle effects,
dark background, professional motion graphics.
```

#### 2. Light Trace / Laser Draw
A beam of light draws the logo outline then fills it in.

**I2V Prompt:**
```
A single bright beam of light traces the outline of the logo on a
dark background, drawing each contour with precision. The light
leaves a glowing trail that persists. Once the outline is complete,
the interior fills with a smooth gradient of light. Subtle lens
flare follows the light point. Professional logo reveal animation.
```

#### 3. Liquid Metal Morph
Chrome liquid pools and sculpts itself into the logo shape.

**I2V Prompt:**
```
Reflective liquid chrome pools on a dark surface, then rises and
morphs into the logo shape. The metallic surface catches dramatic
studio lighting with sharp reflections and caustics. Surface tension
ripples settle as the form solidifies. Luxury, high-end motion
graphics with a polished chrome finish on black background.
```

#### 4. 3D Rotation Reveal
Logo rotates from edge-on (invisible) to face-on.

**I2V Prompt:**
```
The logo is a 3D object with depth and material. It begins turned
away from camera, edge-on, appearing as a thin line. It slowly
rotates on its vertical axis to reveal the full face of the logo.
Dramatic studio lighting with rim light on edges. Cast shadow on
the surface below. Smooth, confident rotation with subtle ease-out.
```

#### 5. Glitch / Digital Reveal
Digital corruption resolves into the clean logo.

**I2V Prompt:**
```
Heavy digital glitch artifacts — pixel fragmentation, RGB splitting,
data corruption, and scan line displacement — gradually resolve and
stabilize to reveal the clean logo. The glitch effects decrease in
intensity over time, with final subtle flickers after the logo is
formed. Dark background, cyberpunk color accents of cyan and magenta.
```

#### 6. Nature / Organic Growth
Branches, vines, or roots grow to form the logo shape.

**I2V Prompt:**
```
On a textured paper background, delicate botanical vines and
branches grow organically from the center outward, gradually
forming the shape of the logo. Small leaves and buds sprout along
the branches. The growth is natural and unhurried, with subtle
movement continuing after the logo is formed. Elegant, organic
motion with earthy green tones.
```

#### 7. Smoke / Ink Reveal
Smoke or ink in water coalesces into the logo.

**I2V Prompt:**
```
Dense smoke wisps swirl in slow motion against a jet black
background. The smoke gradually gathers, compresses, and forms
the logo shape. Volumetric lighting from above catches the smoke
edges. The final form has a solid core with wispy smoke edges that
continue to drift subtly. Atmospheric, moody, cinematic quality.
```

#### 8. Shatter / Reverse Explode
Fragments fly in from all directions and assemble into the logo.

**I2V Prompt:**
```
Hundreds of sharp geometric fragments fly inward from all
directions at high speed, colliding and assembling into the logo
shape at the center. A shockwave ripple pulses outward on impact.
The fragments lock into place with metallic clinks. Dramatic speed
ramp — fast approach, sudden stop. Dark background with motion
blur on the fragments. Cinematic impact.
```

#### 9. Neon Sign Flicker On
Logo as a neon sign that powers on with characteristic flicker.

**I2V Prompt:**
```
The logo rendered as a neon tube sign mounted on a dark wall.
It powers on section by section with electrical flicker and buzz.
Some segments flash multiple times before holding steady. Warm
color glow bleeds onto the wall surface. Realistic neon gas
discharge behavior with bloom and haze. Nighttime atmosphere.
```

#### 10. Stamp / Press Impact
Logo stamps down with force, creating an impression.

**I2V Prompt:**
```
A metallic stamp descends from above and presses the logo into
a soft surface with satisfying impact. The surface deforms on
contact — a ring of displaced material pushes outward. Dust
particles scatter from the impact point. The stamp lifts to
reveal the crisp embossed logo. Dramatic top-down lighting.
Physical, tactile motion with weight and authority.
```

---

## Transitions & Visual Effects

### Scene Transitions

Transitions connect two scenes. For AI video, describe the outgoing scene, the
transition mechanic, and the incoming scene in a single prompt.

#### Transition Prompt Templates

**Morph Transition:**
```
A close-up of [SCENE A SUBJECT] smoothly morphs and transforms
into [SCENE B SUBJECT]. The forms blend fluidly — edges melt,
colors shift, and shapes reorganize seamlessly. Continuous smooth
camera, no cuts. Surreal transformation effect.
```

**Zoom Transition:**
```
Camera pushes into an extreme close-up of [DETAIL IN SCENE A]
until it fills the frame and becomes abstract. Without cutting,
the abstract forms resolve into [SCENE B] as the camera pulls
back to a wide shot. Seamless zoom transition, continuous motion.
```

**Whip Pan Transition:**
```
Camera rapidly whip-pans to the right with heavy motion blur,
the scene smearing into horizontal streaks. As the motion slows,
a completely new scene resolves — [SCENE B DESCRIPTION]. Fast,
energetic camera transition with natural motion blur.
```

**Light Flash Transition:**
```
A brilliant flash of white light overexposes the entire frame,
washing out [SCENE A]. As the light fades, [SCENE B] is revealed
in its place. Clean, simple transition with a bright flare peak.
Smooth exposure ramp up and down.
```

**Liquid / Fluid Transition:**
```
[SCENE A] dissolves into flowing liquid — colors and forms melt
downward like thick paint. The liquid flows reform upward into
the shapes and colors of [SCENE B]. Viscous, organic transition
with fluid dynamics. Art directed color mixing.
```

**Glitch Transition:**
```
[SCENE A] breaks apart with digital glitch artifacts — pixel
sorting, data moshing, frame tearing, RGB displacement. Through
the corruption, [SCENE B] emerges and stabilizes. Brief, punchy
digital transition lasting under one second.
```

### Particle & Fluid Effects

#### Particle System Prompts

**Ember / Spark Particles:**
```
Glowing orange and red ember particles drift upward in slow motion
against a dark background. Each particle pulses with warm light and
leaves a faint trail. Shallow depth of field with some particles
softly out of focus. Cinematic atmosphere, campfire warmth.
```

**Confetti Burst:**
```
An explosive burst of colorful confetti erupts from the center of
the frame. Hundreds of small rectangular pieces in red, blue, gold,
green, and pink tumble and flutter downward with realistic air
resistance. Celebration, victory, festive atmosphere. Bright lighting.
```

**Digital Data Particles:**
```
Streams of glowing cyan data particles flow through a dark
three-dimensional space, forming grid lines, network connections,
and data pathways. Binary numbers and code fragments float among
the particles. Technology, AI, digital infrastructure visualization.
Blue-cyan-white color palette on black.
```

**Dust Motes in Light:**
```
Soft golden light beams stream through a window into a dark room.
Millions of tiny dust particles catch the light, drifting slowly
in gentle air currents. Visible volumetric light shafts. Peaceful,
meditative, nostalgic atmosphere. Shallow depth of field.
```

#### Fluid Motion Prompts

**Ink Drop in Water:**
```
A drop of vibrant [COLOR] ink falls into clear water in slow motion.
The ink blooms and expands in organic tendrils, creating complex
swirling patterns. Shot from the side in a glass tank with bright
backlighting. Mesmerizing fluid dynamics, macro photography feel.
```

**Chrome Liquid Morph:**
```
A blob of reflective liquid chrome floats in zero gravity against
a dark studio background. It slowly morphs through organic shapes —
stretching, splitting, and recombining. Studio lighting creates
sharp highlights and reflections on the metallic surface. Satisfying,
abstract, high-end material study.
```

**Paint Splash / Splatter:**
```
Vibrant paint splashes in slow motion against a white background.
Multiple colors — [COLOR 1], [COLOR 2], [COLOR 3] — collide and
mix mid-air, creating dynamic abstract patterns. Each splash has
thick viscosity and fine spray droplets. High-speed photography
aesthetic, commercial production quality.
```

---

## Lower Thirds & Title Cards

### Lower Third Styles

Lower thirds are text overlays in the bottom third of the frame. For AI generation,
create them as standalone motion graphics on a solid background, then composite.

**Minimal Line + Text:**
```
Clean motion graphics on black background. A thin white horizontal
line slides in from the left, stopping at center-left. Above the
line, "[NAME]" in clean sans-serif fades in. Below the line,
"[TITLE]" in smaller lighter weight text fades in with a 0.2s
delay. Minimal, professional, broadcast quality lower third.
```

**Geometric Shape + Text:**
```
On black background, a sharp rectangular bar in [BRAND COLOR]
slides in from the left edge. The text "[NAME]" appears inside
the bar in white bold sans-serif. A secondary thinner bar below
carries "[TITLE]" in smaller text. Clean geometric motion with
crisp edges. Corporate broadcast graphics.
```

**Glassmorphism Panel:**
```
On a dark background, a frosted glass panel with subtle blur and
soft white border slides up from below. Text "[NAME]" and "[TITLE]"
fade in on the translucent panel. Soft light refracts through the
glass edges. Modern UI design aesthetic, Apple-style glassmorphism.
```

### Title Card Styles

**Cinematic Wide Title:**
```
Black letterboxed frame (2.39:1 aspect ratio). Text "[TITLE]" in
elegant thin serif font fades in at center, letter-spaced widely.
A subtle lens flare drifts slowly across. Below, "[SUBTITLE]" in
smaller sans-serif appears. Premium, cinematic, film title sequence
quality. Measured timing, confident pacing.
```

**Split Screen Title:**
```
Frame divided vertically — left half is [COLOR/TEXTURE], right half
is [CONTRASTING COLOR/TEXTURE]. Large bold text "[TITLE]" spans
both halves, with the text color inverted on each side for contrast.
A vertical line at the division pulses subtly. Bold, graphic,
editorial design. Slow zoom creates depth.
```

---

## Abstract & Generative Motion Art

### Abstract Motion Styles

**Geometric Pattern Animation:**
```
Intricate geometric pattern of interlocking triangles and hexagons
tessellates and rotates in perfect synchronization. The pattern
shifts colors smoothly — deep blue to purple to magenta. Sacred
geometry aesthetic with mathematical precision. Hypnotic, seamless
loop. Black background, sharp vector-clean edges.
```

**Organic Flow / Cellular:**
```
Abstract organic forms reminiscent of microscopic biology — cell
division, membrane ripples, organelle movement. Soft translucent
spheres and tendrils in warm amber and deep teal float and interact.
Bioluminescent glow from within each form. Scientific beauty,
macro scale, slow contemplative motion.
```

**Fractal Zoom:**
```
Infinite fractal zoom — camera continuously pushes forward into an
ever-more-detailed fractal structure. Mandelbrot-style patterns
with vivid color mapping: electric blue, gold, deep crimson.
Self-similar patterns repeat at every scale. Mesmerizing, infinite,
psychedelic mathematics. Smooth continuous forward motion.
```

**Grid Distortion:**
```
A perfect white grid on black background begins to distort — waves
ripple through the grid lines, creating 3D depth illusion. The
distortion follows a smooth sine wave pattern moving from left to
right. Wireframe aesthetic, retro-futuristic, vaporwave influence.
Clean vector lines with subtle glow.
```

---

## Color & Style Direction

### Color Transition Techniques

**Gradient Shift:**
```
Smooth abstract gradient background slowly transitions from
[COLOR A] to [COLOR B] to [COLOR C]. The color changes flow like
aurora borealis — organic, soft, continuous blending with no hard
edges. Gentle motion of color fields. Atmospheric, ambient mood.
```

**Monochrome to Color Reveal:**
```
Scene begins in desaturated black and white. Color gradually bleeds
in from the center outward like watercolor spreading on paper,
revealing the full vibrant palette. The transition from monochrome
to full color takes several seconds. Dramatic reveal moment.
```

### Style Anchor Keywords for Motion Graphics

Prepend these to any motion design prompt to set the overall aesthetic:

| Style | Anchor Keywords |
|-------|----------------|
| Clean / Minimal | `minimalist motion design, clean lines, ample negative space, monochrome palette` |
| Retro / Vintage | `1970s motion graphics, analog feel, warm grain, retro color palette, rounded fonts` |
| Neon / Cyberpunk | `cyberpunk neon glow, dark background, electric blue and magenta, digital rain` |
| Corporate / Professional | `professional broadcast graphics, clean corporate design, brand-safe, polished` |
| Playful / Cartoon | `playful cartoon animation, bouncy motion, bright primary colors, rounded shapes` |
| Luxury / Elegant | `luxury brand animation, gold and black, slow deliberate motion, serif typography` |
| Tech / Futuristic | `futuristic HUD interface, holographic elements, data visualization, blue-white glow` |
| Organic / Natural | `organic motion, flowing natural forms, earth tones, botanical, hand-crafted feel` |
| Brutalist / Bold | `brutalist graphic design in motion, heavy type, high contrast, raw, confrontational` |
| Vaporwave / Retro-Future | `vaporwave aesthetic, 80s sunset gradient, chrome, roman busts, pixel sort, VHS` |

---

## Production Workflows

### Static Design to Animation Pipeline

This is the recommended professional workflow for motion graphics on fal.ai:

**Step 1: Design the Key Frame**
Generate a static frame that represents the final state of your motion graphic. Use `search_models` to find:
- A design-focused image generation model for vector aesthetics
- A text-to-image model with strong text rendering when the frame contains specific typography
- A photorealistic image generation model for cinematic frames

**Step 2: Prepare the Asset**
Use `search_models` to find:
- A background removal model to isolate elements from backgrounds
- A vectorization model to convert raster designs to clean vectors for scaling

**Step 3: Animate with I2V**
Use the static frame as the starting image for I2V generation. Use `search_models` to find an image-to-video model suited for your needs (high motion quality for complex reveals, reliable consistency, or strong stylization).

**Step 4: Iterate**
Generate 3-5 variations, pick the best motion, refine the prompt, render final.

### Motion Graphics for Social Media

| Platform | Aspect Ratio | Duration | Notes |
|----------|-------------|----------|-------|
| Instagram Reels | 9:16 | 3-7s | Hook in first 1s, bold text, high contrast |
| TikTok | 9:16 | 3-10s | Fast pacing, trending effects, text overlays |
| YouTube Shorts | 9:16 | 5-15s | Slightly longer, can breathe more |
| Twitter/X | 16:9 or 1:1 | 3-6s | Auto-plays muted, text must carry the message |
| LinkedIn | 16:9 or 1:1 | 5-10s | Professional, clean, corporate style |
| Instagram Stories | 9:16 | 3-5s | Quick impact, swipe-up friendly |

### Looping Motion Graphics

For seamless loops (social media backgrounds, website heroes, digital signage):

**Prompt addition for loopability:**
```
Seamless looping animation. The motion returns to its starting
position by the end of the clip. Continuous cyclical movement
with no visible start or end point.
```

Best subjects for loops: rotating objects, flowing particles, gradient shifts,
oscillating patterns, floating elements, pulsing glows, drifting clouds.

---

## Complete Example Prompts — Ready to Use

### 1. Kinetic Typography — Motivational Quote
**Model:** Use `search_models` — text-to-video.
```
Cinematic motion graphics on a deep navy background. Bold white
sans-serif text animates word by word: "CREATE SOMETHING BEAUTIFUL."
Each word scales up with elastic bounce, holds for a beat, then the
next word appears. After all words are visible, a subtle gold
underline draws itself beneath the full phrase. Smooth easing,
professional kinetic typography, inspired by Apple keynote design.
```

### 2. Logo Reveal — Tech Startup
**Model:** Use `search_models` — image-to-video (provide logo image as input).
```
The logo assembles from thousands of tiny glowing blue data particles
streaming in from all directions against a black void. The particles
accelerate as they approach the center, forming precise edges.
A pulse of light radiates outward as the final particle locks into
place. The completed logo holds steady with a subtle ambient glow.
Clean, futuristic tech company reveal. Professional motion graphics.
```

### 3. Scene Transition — Product Demo
**Model:** Use `search_models` — text-to-video.
```
A smartphone floats in clean white studio space, slowly rotating to
show its design. The camera pushes into the phone screen, zooming
through the glass into the interface. Inside, app screens scroll
smoothly, demonstrating features with highlighted touch points.
Seamless zoom transition from physical product to digital interface.
Apple-style product video, pristine lighting.
```

### 4. Abstract Background — Event Visuals
**Model:** Use `search_models` — text-to-video.
```
Slow-moving abstract color fields in deep purple, electric blue,
and soft magenta blend and flow like aurora borealis. Gentle organic
motion with no sharp edges. Subtle particle dust floats through the
color fields catching soft light. Ambient, atmospheric, perfect as
a stage background or event visual. Seamless loop, meditative pace.
```

### 5. Title Sequence — YouTube Channel
**Model:** Use `search_models` — text-to-video.
```
Dynamic title sequence: geometric shapes — circles, triangles,
rectangles — in bold primary colors fly in from all edges of
the frame and assemble into a mosaic pattern. The shapes snap
into place with satisfying timing. The mosaic then flattens and
the negative space between shapes forms the channel name
"[CHANNEL NAME]". Energetic, fast-paced, Saul Bass inspired
motion graphics with a modern twist.
```

### 6. Lower Third — Interview Setup
**Model:** Use `search_models` — text-to-video.
```
Professional broadcast lower third on transparent-style dark
background. A thin gold line extends horizontally from center.
Above it, "SARAH CHEN" in crisp white sans-serif slides in from
left. Below, "Chief Technology Officer" in lighter weight text
fades in after a brief delay. The gold line pulses subtly once.
Broadcast news quality, clean corporate design.
```

### 7. Particle Effect — Product Hero
**Model:** Use `search_models` — image-to-video (provide product image as input).
```
The product is surrounded by a burst of golden sparkle particles
that spiral outward in slow motion. Each particle catches dramatic
studio lighting, creating brilliant points of light. Shallow depth
of field with bokeh on distant particles. The product remains sharp
and stable at center while the particle system adds luxury and magic.
Premium product photography brought to life.
```

### 8. Liquid Typography — Beauty Brand
**Model:** Use `search_models` — text-to-video.
```
Luxury beauty brand motion graphics. On a cream background, liquid
gold flows from the top of frame, pooling and forming elegant serif
letters spelling "[BRAND NAME]". The liquid has realistic viscosity,
surface tension, and reflections. After forming, the golden text
holds with subtle surface ripples. Warm, premium, sensual aesthetic.
High-end cosmetics brand feel.
```

### 9. Glitch Reveal — Music Artist
**Model:** Use `search_models` — text-to-video.
```
Black screen erupts with aggressive digital glitch effects — RGB
splitting, pixel sorting, data moshing in neon pink and cyan. Through
the chaos, bold condensed text "[ARTIST NAME]" fights to appear,
glitching and stabilizing repeatedly. The corruption intensifies then
suddenly cuts to clean — the artist name holds solid in white on
black. Silence after chaos. Music release announcement aesthetic.
```

### 10. Minimal Countdown — Event Promo
**Model:** Use `search_models` — text-to-video.
```
Clean white background. A large bold number "3" appears at center
with a subtle ripple effect. It dissolves into particles that reform
as "2". Another dissolution reforms as "1". The final "1" explodes
outward as the event name "[EVENT NAME]" scales up from zero to fill
the frame with dramatic impact. Minimal design, satisfying timing,
black text on white, Swiss design aesthetic.
```

### 11. Holographic UI — Sci-Fi Promo
**Model:** Use `search_models` — text-to-video.
```
Futuristic holographic interface floating in dark space. Translucent
blue panels, circular diagrams, and data streams animate into view.
Text labels and numbers update in real-time. A central holographic
globe rotates slowly. Scan lines and subtle flicker add authenticity.
Iron Man HUD inspired, blue-white color scheme, volumetric light
scatter. Professional sci-fi motion graphics.
```

### 12. Nature Logo — Eco Brand
**Model:** Use `search_models` — image-to-video (provide logo image as input).
```
Time-lapse style: green moss and tiny ferns grow across a stone
surface, gradually filling in the shape of the logo. Small flowers
bloom at key points. Morning dew drops form on the leaves. The
organic growth is patient and beautiful, taking the full duration
of the clip. Natural daylight, macro photography feel, shallow
depth of field on edges. Sustainable, eco-friendly brand aesthetic.
```

### 13. Gradient Typography — Social Hook
**Model:** Use `search_models` — text-to-video.
```
Bold oversized text "DID YOU KNOW?" fills the frame against a
smoothly animated gradient background shifting from coral to violet.
The text has a subtle 3D shadow that shifts with the gradient.
After a beat, the text slides up and out, replaced by smaller text
with the hook content. Social media motion graphics, vertical 9:16
format, attention-grabbing, clean modern design.
```

### 14. Data Visualization — Infographic
**Model:** Use `search_models` — text-to-video.
```
Animated infographic on a clean dark blue background. A bar chart
grows upward, each bar rising sequentially in a staggered cascade.
Numbers count up at the top of each bar. Percentage labels fade in.
A trend line draws itself smoothly across the chart peaks. Clean
sans-serif labels, white and cyan on dark blue. Professional data
visualization motion graphics, broadcast quality.
```

### 15. Cinematic Reveal — Film Title
**Model:** Use `search_models` — text-to-video.
```
Extreme slow motion: fine metallic gold dust particles drift through
a beam of light in an otherwise dark void. The particles are so
fine they move like smoke. Gradually, the density of particles
increases in the center, coalescing into the faintest suggestion
of letters. The title "[FILM TITLE]" materializes from the gold
dust, ethereal and luminous. Thin elegant serif typography. The
dust continues to drift after the title forms. Atmospheric,
prestigious, award-season film title sequence quality.
```

### 16. Retro VHS — Nostalgia Content
**Model:** Use `search_models` — text-to-video.
```
VHS recording aesthetic: heavy analog noise, tracking errors,
and color bleeding. Through the distortion, bold 1980s-style
chrome text "[TITLE]" with a neon pink outline appears. Scan
lines roll across the frame. A timestamp flickers in the corner.
The text has a reflective chrome gradient and subtle glow. Retro
synthwave color palette — deep purple, hot pink, chrome silver.
Vaporwave nostalgia, analog warmth, lo-fi charm.
```

