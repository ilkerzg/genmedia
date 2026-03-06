# Character Design — Creation & Consistency Master Guide

Comprehensive reference for creating original characters and maintaining their visual identity across multiple images, scenes, outfits, and video sequences. Every section contains directly usable prompts and actionable techniques. Use `search_models` to find the best model for each task.

---

## Quick Reference

### Top 5 Character Design & Consistency Tips

1. **Write an anchor description once, paste it verbatim everywhere.** Never paraphrase — identical wording produces identical features. Lead with it in the first 30% of every prompt.
2. **Be specific enough to be unique.** "Chin-length chestnut brown bob with side-swept bangs, parted on the left" not just "brown hair".
3. **Include 8–12 physical attributes minimum.** Fewer gives the model too much freedom; more than 15 causes tokens to compete and drift.
4. **Always edit from the same canonical hero image.** Chain edits from a single source rather than cascading edits from edited versions.
5. **Separate permanent traits from variable ones.** The anchor covers face, build, and distinguishing marks. Scene, pose, lighting, and outfit come after.

### Key Prompt Patterns

```
# Anchor block (paste verbatim in every prompt)
[Name], [age]s, [skin tone], [gender],
[face shape] face, [eye shape] [eye color] eyes,
[hair color] [texture] [length] hair [style],
[build], [1-2 distinguishing features],

# Variable block (change per shot)
[action/pose], [location/scene],
[lighting], [camera framing], [style/mood]
```

**Edit pattern (for scene/outfit changes):**
`Keep [Name]'s face, hair, and [distinguishing feature] identical. Change only: [specific change]. Preserve all other details exactly.`

### Consistency Techniques (ranked by reliability)

1. **LoRA fine-tuning** (~95%) — Train on 15-20 images, use trigger word in all prompts. Best for long-term projects.
2. **Image editing / Kontext** (~90%) — Edit hero image to change scene/outfit. Best for quick variations.
3. **Face transfer / IP-Adapter** (~80-85%) — Paste a face from a photo into new scenes. No training needed.
4. **Reference-to-video** (~85%) — Pass a character reference image to video models. Best for multi-shot video.
5. **Seed locking** (~70%) — Same seed + same prompt = similar output. Fragile, breaks with prompt changes.
6. **Anchor description only** (~65%) — Verbatim text repetition. Works for style consistency, less for faces.

### Cheat Sheet — Workflows

| Goal | Approach |
|------|----------|
| New character, no reference | Write anchor → generate with text-to-image → pick best of 4 seeds |
| Same character, new outfit | Hero image → image editing model with "change outfit to X, keep face identical" |
| Same character, new scene | Hero image → image editing model with "place character in X, preserve appearance" |
| Consistent character in video | Reference image → reference-to-video model for each shot |
| Realistic face from a photo | Face transfer model with source face + scene prompt |
| Long-term consistency (5+ images) | Train a LoRA on 15-20 images → use trigger word in all prompts |

---

## 1. The Anchor Description Method

The single most important technique for character consistency without training. An **anchor description** is a fixed block of text describing your character's permanent physical traits. You paste it verbatim into every prompt, changing only the action, scene, and camera angle.

### Rules

1. **Write it once, reuse it everywhere.** Never paraphrase — identical wording triggers identical visual features.
2. **Lead with the anchor.** Place the character description in the first 30% of your prompt (highest token weight zone).
3. **Separate permanent from variable.** The anchor covers anatomy and identity. Scene, pose, lighting, and outfit go after.
4. **Be specific enough to be unique.** "Brown hair" matches millions of faces. "Chin-length chestnut brown bob with side-swept bangs, parted on the left" matches one.
5. **Include 8-12 physical attributes minimum.** Fewer and the model has too much freedom. More than 15 and tokens compete.

### Anchor Template

```
[Name/Label], [age]s, [ethnicity/skin tone], [gender presentation],
[face shape] face, [eye shape] [eye color] eyes, [nose type] nose,
[lip type] lips, [brow description],
[hair color] [hair texture] [hair length] hair [hair style],
[body build], [height impression],
[1-2 distinguishing features: scar, mole, freckles, dimples, etc.]
```

### Filled Example — "Maren"

```
Maren, late 20s, warm olive skin, woman,
oval face, large almond-shaped hazel-green eyes, straight refined nose,
full lips with a natural cupid's bow, dark expressive brows,
dark auburn wavy shoulder-length hair with a loose side part,
athletic slender build, average height,
faint constellation of freckles across nose and cheekbones, small beauty mark below left eye
```

### Filled Example — "Konstantin"

```
Konstantin, early 40s, Eastern European pale complexion, man,
square jaw with strong bone structure, deep-set steel-blue eyes with crow's feet,
prominent straight nose with a slight bump at the bridge,
thin firm lips, heavy dark brows,
salt-and-pepper short cropped hair buzzed at the sides,
broad-shouldered muscular build, tall imposing frame,
vertical scar through left eyebrow, five o'clock shadow stubble
```

### What to NEVER Change Between Images

- Bone structure (face shape, jaw, cheekbones)
- Eye color and shape
- Nose shape
- Skin tone
- Hair color (unless deliberately showing a style change)
- Distinguishing marks (scars, moles, freckles)
- Body proportions and build

### What CAN Change Between Images

- Expression and emotion
- Pose and body angle
- Outfit and accessories
- Hairstyling (up, down, braided — same hair color/texture)
- Lighting and color grading
- Background and environment
- Camera angle and focal length

---

## 2. Facial Features Vocabulary

Use these precise terms instead of vague descriptions. Models trained on captioned photography datasets respond strongly to anatomical vocabulary.

### Face Shape

| Shape | Description | Prompt Keywords |
|-------|-------------|-----------------|
| Oval | Balanced proportions, gently narrowing at forehead and jaw | `oval face, balanced proportions` |
| Round | Equal width and length, soft jawline, full cheeks | `round face, soft jawline, full cheeks` |
| Square | Strong angular jaw, wide forehead, equal width top and bottom | `square jaw, angular face, strong bone structure` |
| Heart | Wide forehead, high cheekbones, narrow pointed chin | `heart-shaped face, pointed chin, wide cheekbones` |
| Diamond | Narrow forehead, wide cheekbones, narrow jaw | `diamond face shape, prominent cheekbones, narrow jaw` |
| Oblong | Longer than wide, straight cheeks, balanced jaw | `long face, high forehead, straight cheek line` |
| Rectangular | Like square but elongated, strong jaw, angular | `rectangular face, strong angular jaw, long proportions` |

### Eyes

**Shape:** almond, round, hooded, monolid, deep-set, upturned, downturned, wide-set, close-set, protruding, heavy-lidded

**Color vocabulary:**
- Blues: ice blue, cerulean, steel blue, navy, cornflower, pale blue-grey
- Greens: emerald, sage, olive, moss green, sea green, jade, hazel-green
- Browns: warm amber, chocolate brown, dark espresso, honey brown, golden brown, mahogany
- Greys: silver grey, storm grey, pale grey, blue-grey, green-grey
- Rare: violet, heterochromia (two different colors), amber gold

**Expression modifiers:**
- `bright alert eyes` — engaged, present
- `heavy-lidded drowsy eyes` — tired, seductive, bored
- `wide startled eyes` — surprise, fear
- `narrowed suspicious eyes` — distrust, anger, focus
- `soft warm gaze` — kindness, love, contentment
- `piercing intense stare` — determination, menace, confidence
- `glassy tear-filled eyes` — sadness, emotion, overwhelm
- `crinkled laughing eyes` — genuine joy, warmth

### Nose

| Type | Prompt Keywords |
|------|-----------------|
| Aquiline | `aquiline nose, prominent bridge, curved downward` |
| Button | `small button nose, rounded upturned tip` |
| Straight | `straight refined nose, clean profile line` |
| Roman | `Roman nose, high bridge, slightly convex` |
| Wide/Broad | `wide nose, broad flat bridge, round nostrils` |
| Pointed | `narrow pointed nose, sharp tip, thin bridge` |
| Snub | `snub nose, short upturned, slightly concave profile` |
| Bulbous | `rounded bulbous nose tip, wide nostrils` |

### Lips

| Type | Prompt Keywords |
|------|-----------------|
| Full | `full plump lips, generous volume top and bottom` |
| Thin | `thin precise lips, defined lip line` |
| Cupid's bow | `pronounced cupid's bow, defined peaks on upper lip` |
| Wide | `wide mouth, broad smile line` |
| Heart-shaped | `heart-shaped lips, fuller center tapering to corners` |
| Downturned | `downturned lip corners, serious resting expression` |
| Asymmetric | `slightly crooked smile, asymmetric lip line` |

### Skin

**Tone vocabulary:** porcelain, ivory, fair, peach, light beige, warm beige, tan, golden, caramel, olive, light brown, warm brown, medium brown, bronze, copper, deep brown, dark brown, espresso, ebony, mahogany

**Texture modifiers:** smooth, clear complexion, freckled, sun-kissed, weathered, textured, pockmarked, ruddy, flushed cheeks, matte, dewy, glowing

**Age markers:** fine lines around eyes, laugh lines, forehead creases, crow's feet, nasolabial folds, age spots, sun damage

### Hair

**Color:** platinum blonde, golden blonde, strawberry blonde, honey blonde, ash blonde, dirty blonde, light brown, chestnut brown, chocolate brown, dark brown, auburn, copper red, ginger, fiery red, burgundy, jet black, blue-black, raven, silver, salt-and-pepper, white, grey-streaked

**Texture:** pin-straight, stick-straight, slight wave, loose waves, beachy waves, tight curls, coils, kinky curls, afro-textured, springy ringlets, frizzy, sleek, silky, coarse, fine, thick, wispy

**Length:** buzzcut, closely cropped, pixie cut, ear-length, chin-length, jaw-length, shoulder-length, collarbone-length, mid-back, waist-length, hip-length

**Style:** slicked back, side-parted, center-parted, swept to one side, messy tousled, windswept, pulled back in ponytail, high bun, low bun, braided, French braid, twin braids, loose flowing, half-up half-down, pinned with clips, under a hat, with headband, dreadlocks, locs, cornrows, box braids, twist-out, bantu knots

### Facial Hair (Male Characters)

| Type | Prompt Keywords |
|------|-----------------|
| Clean-shaven | `clean-shaven, smooth jaw` |
| Five o'clock shadow | `five o'clock shadow, light stubble` |
| Heavy stubble | `heavy stubble, 3-day beard growth` |
| Short boxed beard | `short trimmed beard, neat beard line, boxed beard` |
| Full beard | `full thick beard, natural beard` |
| Long beard | `long flowing beard, wizard beard` |
| Goatee | `goatee, chin beard, no sideburns` |
| Van Dyke | `Van Dyke beard, pointed goatee with disconnected mustache` |
| Handlebar mustache | `handlebar mustache, curled waxed tips` |
| Mutton chops | `mutton chop sideburns, clean chin` |

---

## 3. Body Types & Proportions

### Build Vocabulary

| Build | Female Keywords | Male Keywords |
|-------|----------------|---------------|
| Very thin | `waif-like, slender, narrow frame, willowy` | `lean, wiry, thin frame, narrow shoulders` |
| Athletic | `athletic build, toned, defined muscles, fit` | `athletic build, muscular definition, V-taper torso` |
| Average | `average build, soft contours, natural proportions` | `average build, moderate frame, natural proportions` |
| Curvy | `curvy figure, hourglass proportions, full hips` | — |
| Muscular | `muscular, powerful build, visible muscle definition` | `heavily muscular, broad-shouldered, thick-necked, bodybuilder` |
| Stocky | `stocky, compact, sturdy build, broad` | `stocky, barrel-chested, thick, compact powerful build` |
| Heavy | `plus-size, full-figured, soft rounded build` | `heavyset, large frame, thick build` |
| Petite | `petite, small-framed, delicate bone structure` | `short, compact frame, small-boned` |
| Tall | `tall, long-limbed, statuesque, leggy` | `tall, towering, long-limbed, imposing height` |

### Posture Keywords

- `perfect upright posture, shoulders back` — confidence, military, royalty
- `relaxed casual stance, weight on one hip` — approachable, youthful
- `hunched forward, rounded shoulders` — tired, defeated, scholarly
- `wide power stance, arms crossed` — authority, defiance
- `leaning against [surface], one knee bent` — casual cool, nonchalance
- `coiled ready stance, weight on balls of feet` — combat readiness, tension
- `graceful elongated posture, dancer's bearing` — elegance, training

---

## 4. Expression & Emotion Library

### Complete Expression Prompts

Each entry is a drop-in replacement for the expression portion of your prompt.

**Joy / Happiness:**
```
genuine warm smile reaching the eyes, crinkled eye corners, raised cheeks, relaxed open expression, bright sparkling eyes
```

**Sadness / Grief:**
```
downcast eyes, slightly furrowed brow, downturned lip corners, glistening unshed tears, hollow distant gaze, jaw tight with restrained emotion
```

**Anger / Fury:**
```
intense glare, deeply furrowed brow, flared nostrils, clenched jaw, tight pressed lips, narrowed eyes burning with rage, tense neck tendons
```

**Fear / Terror:**
```
wide eyes with visible whites, raised eyebrows, parted trembling lips, pale complexion, pupils dilated, frozen expression, tense pulled-back posture
```

**Surprise / Shock:**
```
raised eyebrows, wide open eyes, parted mouth forming an O, slightly pulled-back head, hand rising toward face
```

**Disgust / Revulsion:**
```
wrinkled nose, curled upper lip, narrowed eyes, turned-away chin, one eyebrow raised, mouth pulled to one side
```

**Contempt / Disdain:**
```
one corner of the mouth raised in a smirk, half-lidded eyes, slightly raised chin, cool dismissive gaze, one eyebrow arched
```

**Determination / Resolve:**
```
set jaw, steady focused gaze, slightly narrowed eyes, firm closed lips, squared shoulders, subtle forward lean, unwavering stare
```

**Serenity / Peace:**
```
soft closed-lip smile, relaxed half-closed eyes, smooth unwrinkled brow, slightly tilted head, calm steady breathing expression, gentle gaze
```

**Mischief / Playfulness:**
```
asymmetric grin, one raised eyebrow, twinkling conspiratorial eyes, slight head tilt, knowing smirk, dimples visible
```

**Exhaustion / Weariness:**
```
heavy-lidded half-closed eyes, dark circles, slack jaw, pale complexion, unfocused distant stare, hand rubbing temple, slouched posture
```

**Concentration / Focus:**
```
slightly furrowed brow, eyes locked on target, pursed lips, narrowed gaze, head tilted slightly forward, jaw set, absorbed expression
```

---

## 5. Costume & Outfit Design

### Period-Accurate Vocabulary

**Medieval (5th-15th century):**
`chainmail hauberk, leather bracers, woolen tunic, linen undertunic, fur-trimmed cloak, leather belt with brass buckle, pointed leather boots, heraldic tabard, iron helm, padded gambeson`

**Renaissance (15th-17th century):**
`doublet with slashed sleeves, ruff collar, brocade vest, puffed trunk hose, velvet cap with feather, embroidered bodice, farthingale skirt, pearl-strung hair, high-collared gown`

**Victorian (1837-1901):**
`high-collared white shirt, waistcoat, pocket watch chain, top hat, tailcoat, cravat, bustle dress, corset-shaped bodice, mutton-chop sleeves, lace parasol, riding boots, morning coat`

**1920s Art Deco:**
`beaded flapper dress, long pearl necklace, feathered headband, T-strap heels, finger waves, dropped waist, geometric beading, silk fringe, fur stole, pin-striped three-piece suit, fedora, wingtip oxfords`

**1950s:**
`full circle skirt, petticoat, cardigan twin set, saddle shoes, cat-eye glasses, pearl earrings, pompadour hairstyle, leather jacket, white T-shirt, cuffed denim, penny loafers`

**1970s:**
`bell-bottom jeans, platform shoes, halter top, wide collar polyester shirt, suede fringe vest, tinted aviator sunglasses, macramé accessories, corduroy flares, peasant blouse`

**1980s:**
`oversized blazer with shoulder pads, acid-washed denim, leg warmers, neon windbreaker, Members Only jacket, high-waisted jeans, chunky gold jewelry, power suit, scrunchie`

**Modern / Contemporary:**
`tailored slim-fit blazer, raw denim jeans, minimalist white sneakers, cashmere turtleneck, oversized coat, crossbody leather bag, aviator sunglasses, Chelsea boots, linen shirt`

**Cyberpunk / Futuristic:**
`holographic jacket, LED-embedded clothing, matte black tactical vest, chrome visor, neon-piped bodysuit, augmented reality contact lenses, carbon-fiber arm guards, transparent raincoat with circuit patterns`

### Material Keywords

Models respond well to specific fabric descriptions that carry visual and textural information:

`raw silk, crushed velvet, distressed leather, waxed canvas, brushed denim, cashmere knit, sheer organza, matte neoprene, quilted satin, crinkled linen, oiled suede, patent leather, chunky cable knit, metallic lamé, sequined mesh, washed chambray, stretch lycra, perforated pleather`

### Fit Descriptions

- `tailored, perfectly fitted, sharp lines` — professional, precise
- `oversized, baggy, hanging off shoulders` — streetwear, comfort, casual
- `skin-tight, form-fitting, second-skin` — athletic, futuristic, bold
- `flowing, loose draping, billowing` — ethereal, romantic, fantasy
- `structured, architectural, sculptural silhouette` — high fashion, avant-garde
- `distressed, torn, patched, worn` — post-apocalyptic, punk, lived-in

### Color Coordination

- **Monochromatic:** `dressed entirely in shades of navy blue, tonal outfit`
- **Complementary:** `orange jacket against deep blue backdrop`
- **Analogous:** `warm outfit in rust, amber, and golden yellow`
- **Accent pop:** `all-black outfit with a single crimson red scarf`

---

## 6. Character Sheet Generation

Character sheets are multi-view or multi-expression images on a single canvas. They are essential for pre-production and establishing a character before generating scenes.

### Multi-View Turnaround Sheet

Use `search_models` to find the best image generation model for character sheets.

```
Character turnaround reference sheet, white background, five views of the same character,
front view, three-quarter view, side profile, three-quarter back, back view,
all on a clean white studio background with consistent lighting,

[INSERT ANCHOR DESCRIPTION],

wearing [outfit description],
concept art style, character design sheet, blueprint layout,
evenly spaced views with consistent proportions, professional character reference
```

### Expression Sheet

Use `search_models` to find the best image generation model for expression sheets.

```
Character expression sheet, 3x3 grid on white background,
nine different facial expressions of the same character,
top row: neutral, happy smile, laughing,
middle row: angry, sad, surprised,
bottom row: scared, disgusted, determined,

[INSERT ANCHOR DESCRIPTION],

close-up headshots, consistent lighting, concept art style,
expression reference sheet, each expression clearly distinct and readable
```

### Outfit Variation Sheet

```
Outfit variation sheet, 1x4 horizontal layout on white background,
four full-body views of the same character in different outfits,
left to right: casual everyday wear, formal business attire, athletic sportswear, evening glamour,

[INSERT ANCHOR DESCRIPTION],

consistent character proportions across all four, concept art style,
fashion design reference sheet, clean studio lighting
```

---

## 7. Consistency Techniques — Ranked by Reliability

### Technique 1: LoRA Fine-Tuning (BEST — 95%+ Consistency)

Train a dedicated model weight on your character's face and body. This is the gold standard.

**Image Requirements:**
- 15-30 images of your character (generated or real photos with consent)
- Diverse angles: front, 3/4 left, 3/4 right, profile, slight up-angle, slight down-angle
- Diverse lighting: natural light, studio light, warm, cool, high-key, low-key
- Diverse backgrounds: solid colors, outdoor, indoor, simple, complex
- Consistent identity across all training images
- Minimum 512x512 resolution, ideally 1024x1024
- No heavy filters, watermarks, or text overlays

**Training:** Use `search_models` with "lora training" to find available LoRA training endpoints. Options include fast training, portrait-specialized training, and architecture-specific trainers for different inference pipelines.

**Inference (using trained LoRA):** Use `search_models` with "lora" to find inference endpoints that support LoRA weights. Available options include image generation with LoRA, image editing with LoRA, and video generation with LoRA (both text-to-video and image-to-video).

**LoRA Strength Tuning:**
- `0.4-0.5` — Subtle influence, character "inspired by" the LoRA. Useful for blending multiple LoRAs.
- `0.6-0.7` — Moderate likeness. Face recognizable but allows more variation in style.
- `0.8-0.9` — Strong identity lock. Recommended default range for character consistency.
- `1.0` — Maximum fidelity. Can cause rigidity or artifacts. Use when identity must be exact.
- `>1.0` — Overfitting artifacts. Avoid unless testing.

**Workflow:**
1. Generate 15-30 diverse images of your character using anchor description + seed locking.
2. Curate best images: remove duplicates, artifacts, inconsistent faces.
3. Submit to training endpoint with trigger word (e.g., `ohwx_maren`).
4. Inference: include trigger word in prompt + set LoRA strength to 0.8.

```
ohwx_maren standing in a sunlit Tuscan vineyard, golden hour light,
wearing a white linen sundress, wind in her hair,
warm olive skin, hazel-green eyes, auburn wavy hair, freckles across nose,
cinematic photograph, shot on 85mm f/1.4, shallow depth of field
```

### Technique 2: Kontext Editing (90% Consistency)

Start from one perfect reference image and iteratively edit pose, scene, outfit, and expression while preserving identity.

Use `search_models` with "kontext" or "image editing" to find the best editing model. Choose the highest-fidelity option for character consistency work.

**Three-Part Prompt Structure:**
```
1. IDENTIFY: "The woman in the image..."
2. CHANGE: "is now standing in a rainy Tokyo street at night..."
3. PRESERVE: "maintaining the same face, hair, and body proportions"
```

**Example chain — Same character, four scenes:**

Starting image: Maren in a studio headshot.

Edit 1 — Change scene:
```
The woman in the image is now standing on a windswept cliff overlooking the ocean at sunset,
wearing a navy peacoat and white scarf, her auburn hair blowing in the wind,
same face and features, same freckles and hazel-green eyes, cinematic wide shot
```

Edit 2 — Change outfit:
```
The woman in the image is now sitting at a modern café table, wearing a black turtleneck
and gold hoop earrings, holding an espresso cup, same face and exact features,
warm interior lighting, medium shot
```

Edit 3 — Change expression:
```
The woman in the image now has a serious determined expression, jaw set,
eyes focused and intense, same face and all features preserved exactly,
dramatic low-key lighting, close-up portrait
```

**Pro tips:**
- Always reference "the woman/man in the image" to anchor identity to the input.
- Use the highest-fidelity editing model available for best identity preservation.
- Chain at most 3-4 edits from one source before drift accumulates. Return to the original reference periodically.
- Keep edits focused — change one major element per step for best results.

### Technique 3: IP-Adapter / PuLID (80-85% Consistency)

Transfer facial identity from a single reference image without any training. Fast and zero-shot.

Use `search_models` with "face transfer" or "IP-Adapter" to find available identity transfer models. Options include:
- **PuLID-based models** — Inject a specific face into new scenes with strong facial identity.
- **IP-Adapter + ControlNet models** — Combine reference style, pose control, and identity in one call.

**PuLID workflow:**
1. Provide a clear, well-lit frontal face photo as the reference.
2. Write a full scene prompt describing the desired output.
3. The model replaces the generated face with identity features from the reference.

```
Prompt: "A woman standing in a neon-lit cyberpunk alley, wearing a leather jacket,
rain falling, dramatic backlighting, cinematic photograph"

Reference: [upload clear frontal photo of character]
```

**When to use PuLID vs LoRA:**
- PuLID: Quick one-off identity transfer, no training budget, testing character in a new scene.
- LoRA: Production pipeline, 10+ images needed, highest consistency requirements.

### Technique 4: Reference-to-Video (85% Consistency)

Use a character reference image to generate video where the character appears with consistent identity.

Use `search_models` with "reference to video" to find models that accept a character reference image and generate video maintaining that character's identity.

**Workflow:**
1. Generate one perfect hero image of your character.
2. Submit as reference image with a video prompt describing the action.
3. The model generates video maintaining the character's identity.

```
Reference image: [hero portrait of Maren]
Prompt: "A woman walks through an autumn forest, leaves falling around her,
she pauses to pick up a golden leaf and smiles, warm afternoon light filtering through trees,
cinematic tracking shot, shallow depth of field"
```

### Technique 5: Seed Locking (70% Consistency)

Same seed + identical description tokens = similar facial geometry. Cheapest method but lowest reliability.

**How it works:**
- The random seed determines the noise pattern that initializes generation.
- If every other variable is identical, the same seed produces near-identical output.
- Change only the action/scene/outfit words. Keep character description and style tokens identical.

**Example — Same seed, two scenes:**

Prompt A (seed: 42):
```
Maren, late 20s, warm olive skin, oval face, large almond-shaped hazel-green eyes,
straight refined nose, full lips with cupid's bow, dark auburn wavy shoulder-length hair,
athletic slender build, freckles across nose and cheekbones,
sitting in a library reading a leather-bound book, warm lamplight,
cinematic photograph, 85mm lens, shallow depth of field
```

Prompt B (seed: 42):
```
Maren, late 20s, warm olive skin, oval face, large almond-shaped hazel-green eyes,
straight refined nose, full lips with cupid's bow, dark auburn wavy shoulder-length hair,
athletic slender build, freckles across nose and cheekbones,
walking through a rain-soaked city street at night, neon reflections,
cinematic photograph, 85mm lens, shallow depth of field
```

**Limitations:** Even small prompt changes shift the noise trajectory. Expect ~70% facial similarity. Body proportions may vary. Best as a supplement to other methods, not standalone.

### Technique 6: Anchor Description Repetition (65% Consistency)

The baseline method. No reference image, no training, no seed locking. Pure prompt engineering.

Repeat the exact anchor description in every prompt. Lock the camera setup, lighting style, and artistic direction alongside the character description.

**Boost consistency by also fixing:**
- Art style: `"cinematic photograph, shot on ARRI Alexa, 2.39:1 aspect ratio"`
- Lighting: `"Rembrandt lighting, warm key light from camera-left"`
- Color grade: `"warm amber tones with deep teal shadows"`
- Lens: `"85mm f/1.4, shallow depth of field"`

---

## 8. Character Types & Archetypes

### Realistic Humans (Photographic)

Use `search_models` with "image generation" to find the best photorealistic image generation model.

Prefix your prompt with photographic anchors:
```
editorial photograph, natural light, shot on Canon EOS R5 with RF 85mm f/1.2L,
ISO 200, slight film grain, skin texture visible, natural imperfections
```

### Stylized — Anime / Manga

```
anime character illustration, Studio Ghibli style, cel-shaded, clean linework,
large expressive eyes, small nose and mouth, soft pastel coloring,
[character anchor description adapted: emphasize eye color/size, hair style, outfit details]
```

### Stylized — Pixar / 3D Animation

```
3D rendered character, Pixar animation style, subsurface scattering on skin,
slightly exaggerated proportions, large expressive eyes, rounded features,
smooth plastic-like skin texture, vibrant saturated colors, studio lighting
```

### Stylized — Comic Book

```
comic book character illustration, bold ink outlines, flat cel shading,
dynamic pose, halftone dot shading, vibrant primary colors,
Marvel/DC art style, clean graphic novel rendering
```

### Fantasy Races

**Elf:**
```
high elf, tall and slender, elongated pointed ears, angular ethereal features,
high cheekbones, large luminous eyes, flawless pale skin with faint glow,
long straight silver hair, graceful willowy build, ageless appearance
```

**Dwarf:**
```
mountain dwarf, short and stocky, barrel-chested powerful build,
broad flat nose, deep-set eyes under heavy brow ridge, thick braided beard
with metal clasps, ruddy wind-burnt skin, powerful forearms, about 4 feet tall
```

**Orc:**
```
orc warrior, green-grey skin, broad flat face, prominent lower jaw with tusks,
small deep-set yellow eyes, heavy brow ridge, battle scars across face,
thick muscular neck, massive heavily-built frame, tribal tattoos
```

### Sci-Fi Characters

**Cyberpunk:**
```
cyberpunk character, chrome cybernetic implants on temple and jaw,
LED-lit iris augmentations glowing ice blue, shaved sides with neon-streaked mohawk,
scarred pale skin, mixed tech-organic aesthetic, dystopian near-future
```

**Space Explorer:**
```
space explorer, form-fitting EVA suit with mission patches,
helmet visor reflecting starfield, practical utility belt with tools,
weather-worn face with goggle tan lines, determined expression,
zero-gravity floating hair
```

### Mascots & Brand Characters

```
brand mascot character design, simple memorable silhouette, bold flat colors,
friendly approachable expression, exaggerated proportions (large head, small body),
clean vector-style rendering, works at any scale, white background,
brand color palette: [specify HEX codes]
```

---

## 9. Character Animation Workflow

### Still-to-Video Pipeline

**Step 1: Generate the perfect character still**

Use a high-quality image generation model (find one via `search_models`) with your full anchor description. Generate multiple candidates and select the best one.

**Step 2: Remove background (if needed)**

Use `search_models` with "background removal" to find a background removal model. Look for one with clean production-ready output and good hair segmentation.

**Step 3: Animate with Image-to-Video**

Use `search_models` with "image to video" to find the best image-to-video model. Look for high quality motion, cinematic results, and strong prompt adherence.

**Animation prompt tips:**
- Describe motion minimally: `"she slowly turns her head to the right and smiles"`.
- Avoid complex multi-step actions in one clip. One motion per generation.
- Keep the character description from your still in the video prompt for coherence.
- Specify camera: `"static camera"` or `"slow push-in"` — unspecified camera = random movement.

**Step 4: Talking head / Lipsync (if needed)**

Use `search_models` with "talking head" or "lipsync" to find models that generate talking avatars from image + audio. Options may include portrait-focused and full-body animation models.

### Multi-Shot Character Consistency in Video

**Method A: Reference-to-Video models**
Use `search_models` with "reference to video" to find a reference-to-video model. Provide the same hero reference image across all shots. Vary only the action prompt.

**Method B: LoRA-based video generation**
Train a LoRA on your character, then use it directly in video generation. Use `search_models` with "video lora" to find video generation endpoints that support LoRA weights. Both text-to-video and image-to-video variants may be available — image-to-video with LoRA provides maximum consistency.

### Virtual Try-On

Place your character in specific clothing using a virtual try-on model (find one via `search_models` with "try on"). Provide a character image + garment image to get the character wearing that garment.

### Upscaling Final Output

Before delivery, upscale for maximum quality. Use `search_models` with "upscale" to find an image upscaling model. Look for detail-enhancing options that work well with faces.

---

## 10. Age & Aging

### Age Descriptors

| Age Range | Keywords |
|-----------|----------|
| Child (5-12) | `young child, round soft features, large eyes relative to face, baby fat cheeks, small nose, smooth skin, missing or growing-in teeth` |
| Teen (13-17) | `teenager, slightly angular features emerging, clear skin, youthful glow, lanky proportions, not fully mature bone structure` |
| Young adult (18-25) | `young adult, smooth skin, full lips, bright clear eyes, taut jawline, peak collagen, minimal lines` |
| Late 20s-30s | `late twenties / early thirties, first subtle expression lines, mature features, defined bone structure, confident bearing` |
| 40s | `early forties, visible laugh lines, crow's feet beginning, possible grey streaks, mature distinguished look` |
| 50s | `fifties, prominent nasolabial folds, forehead lines, grey hair or salt-and-pepper, skin losing elasticity, distinguished character` |
| 60s | `sixties, deep wrinkles, weathered skin texture, thin lips, age spots, white or grey hair, jowls forming` |
| 70s+ | `elderly, deeply lined face, paper-thin skin, sunken cheeks, white hair, gentle clouded eyes, liver spots, wisdom in expression` |

### Aging a Character Prompt

To show the same character at different ages while maintaining identity, keep your anchor description and add age modifiers:

```
[Anchor description for Maren], now in her early sixties,
her once-auburn hair now silver-streaked and cut shorter to chin length,
laugh lines deeply etched around her still-hazel-green eyes,
same constellation of freckles now faded, warm weathered skin,
reading glasses perched on her nose, wearing a soft wool cardigan
```

---

## 11. Lighting for Character Portraits

Lighting fundamentally changes character mood and perceived personality. Use these setups as prompt fragments.

### Portrait Lighting Setups

**Rembrandt Lighting:**
```
Rembrandt lighting, single key light from 45 degrees above and to the side,
triangle of light on the shadow-side cheek, dramatic chiaroscuro,
warm directional light, deep shadows on one side of the face
```

**Butterfly / Paramount Lighting:**
```
butterfly lighting, key light directly above and in front of the face,
small shadow under the nose, glamorous even illumination,
flattering for high cheekbones, classic Hollywood portrait lighting
```

**Split Lighting:**
```
split lighting, half the face lit and half in deep shadow,
dramatic 90-degree side light, mysterious and intense,
strong contrast, moody noir atmosphere
```

**Broad Lighting:**
```
broad lighting, the wider side of the face toward camera is lit,
open friendly illumination, soft fill light, approachable look
```

**Rim / Edge Lighting:**
```
rim lighting, bright edge light outlining the subject's silhouette,
backlit with hair light creating a glowing halo effect,
dark moody background, dramatic separation from background
```

**Natural Window Light:**
```
soft natural window light from camera left, gentle directional illumination,
soft shadows, warm daylight color temperature, diffused and flattering,
no harsh highlights, natural skin rendering
```

**Neon / Colored Gel Lighting:**
```
colored gel lighting, magenta light from the left and cyan from the right,
neon-lit face, colorful dramatic portrait, mixed color temperature,
urban night aesthetic, bold color contrast on skin
```

---

## 12. Complete Example Prompts

### Realistic Woman — Editorial Portrait

Use a high-quality image generation model (find via `search_models`).

```
Editorial fashion photograph of Maren, late 20s, warm olive skin, woman,
oval face, large almond-shaped hazel-green eyes, straight refined nose,
full lips with natural cupid's bow, dark expressive brows,
dark auburn wavy shoulder-length hair with loose side part,
athletic slender build, faint freckles across nose and cheekbones,
small beauty mark below left eye,

wearing an oversized camel wool coat over a cream cashmere turtleneck,
gold minimal hoop earrings, no other jewelry,

standing on a foggy cobblestone bridge at dawn, soft diffused golden light,
shallow depth of field, shot on Hasselblad X2D with XCD 90mm f/2.5,
natural film grain, muted earth-tone palette with warm amber highlights,
Vogue editorial style, 3:4 portrait orientation
```

### Realistic Man — Cinematic

Use a high-quality image generation model (find via `search_models`).

```
Cinematic portrait of Konstantin, early 40s, Eastern European pale complexion, man,
square jaw with strong bone structure, deep-set steel-blue eyes with crow's feet,
prominent straight nose with slight bump at the bridge,
thin firm lips, heavy dark brows,
salt-and-pepper short cropped hair buzzed at the sides,
broad-shouldered muscular build, tall imposing frame,
vertical scar through left eyebrow, five o'clock shadow stubble,

wearing a worn dark leather jacket over a black henley shirt,
standing in a dimly lit industrial warehouse, rain visible through broken windows,
moody blue-grey color grading, single shaft of cold daylight from above,
Rembrandt lighting on face, shot on ARRI Alexa with Cooke S4 50mm,
cinematic 2.39:1 widescreen, Denis Villeneuve atmosphere
```

### Anime Character

Use an image generation model (find via `search_models`).

```
Anime character illustration of a young woman warrior, cel-shaded style,
large expressive violet eyes with star-shaped highlights,
long flowing silver hair in twin tails reaching her waist with black ribbon ties,
heart-shaped face, small pointed nose, determined expression with slight smirk,
fair skin with a bandage across her left cheek,

wearing a cropped dark blue military jacket with gold epaulettes and silver buttons,
white blouse underneath, pleated grey skirt with hidden blade holster,
knee-high black leather boots, fingerless gloves,

dynamic three-quarter pose, one hand on hip, wind blowing hair dramatically,
cherry blossom petals floating in background, sunset sky in orange and purple,
clean bold linework, vibrant saturated anime coloring, Studio Bones art style
```

### Pixar-Style Character

Use a high-quality image generation model (find via `search_models`).

```
3D Pixar animation style character, a cheerful elderly inventor,
round friendly face, large bright blue eyes behind thick round spectacles,
bulbous red nose, wide warm smile with gap between front teeth,
wild Einstein-like white hair sticking out in all directions,
bushy white eyebrows, rosy cheeks, wrinkled but lively expression,
short and round body, small hands, wearing a stained lab coat over a
green argyle sweater vest, bow tie, holding a glowing contraption,

standing in a cluttered workshop full of impossible machines and bubbling beakers,
warm golden lighting from hanging Edison bulbs, subsurface scattering on skin,
slightly exaggerated proportions, vibrant Pixar color palette, smooth 3D render
```

### Fantasy Warrior

Use a high-quality image generation model (find via `search_models`).

```
Epic fantasy character portrait of a battle-scarred orc chieftain,
green-grey skin with ritual scarification patterns across broad flat face,
prominent lower jaw with two polished bone tusks, small deep-set amber eyes,
heavy brow ridge casting shadows over eyes, broken nose healed crooked,
thick muscular neck with iron torc, massive heavily-built frame,
long black hair in warrior braids tied with bronze rings,

wearing layered fur and leather armor, massive pauldron made from a dragon skull
on the left shoulder, runic tattoos glowing faintly on forearms,
holding a jagged two-handed greataxe resting on the ground,

standing on a windswept mountain peak, storm clouds gathering behind,
dramatic low-angle shot, rim lighting from lightning flashes,
dark fantasy art style, painted realism, Frank Frazetta inspired composition
```

### Sci-Fi Pilot

Use a high-quality image generation model (find via `search_models`).

```
Science fiction character portrait of a starfighter pilot,
Southeast Asian woman in her mid-30s, warm golden-brown skin,
sharp angular face, intense dark brown monolid eyes, straight nose,
thin determined lips pressed together, high cheekbones,
jet black hair in a tight functional braid tucked into collar,
lean muscular build from zero-G training, small scar on right temple,

wearing a form-fitting white flight suit with blue mission patches,
neural interface port visible behind right ear glowing soft blue,
holding helmet under one arm, standing beside the cockpit of her ship,

dramatic hangar lighting, cool blue ambient with warm orange from engine exhaust,
lens flare from overhead lights, cinematic medium shot,
hard sci-fi aesthetic, Syd Mead inspired production design, 16:9 widescreen
```

### Historical Figure — 1920s Jazz Singer

Use a high-quality image generation model (find via `search_models`).

```
Vintage-style portrait photograph of a 1920s jazz singer,
African American woman in her late 20s, deep brown skin with golden undertones,
oval face, large warm dark brown eyes with kohl liner,
full lips painted in dark berry, high sculpted cheekbones, beauty mark on right cheek,
finger-waved short black hair pressed close to the head with a jeweled headband,
slender graceful build, long elegant neck,

wearing a beaded champagne-gold flapper dress with long fringe,
long strand of pearls knotted at the waist, elbow-length silk gloves,
diamond drop earrings catching the light,

standing at a vintage microphone in a smoky speakeasy,
warm amber spotlight from above, hazy atmosphere, dark background,
sepia-toned with selective warm gold highlights,
photographed in the style of James Van Der Zee, period-accurate 1920s aesthetic
```

### Character in Multiple Outfits (Consistency via Kontext)

**Step 1 — Generate hero image** with a high-quality image generation model:
```
Studio portrait photograph of Maren, late 20s, warm olive skin, oval face,
large almond-shaped hazel-green eyes, straight refined nose, full lips with cupid's bow,
dark auburn wavy shoulder-length hair, freckles across nose and cheekbones,
wearing a simple white T-shirt, neutral expression, clean white background,
medium shot, even studio lighting, 85mm lens
```

**Step 2 — Edit outfit** with a high-fidelity image editing model:
```
The woman in the image is now wearing an elegant black evening gown with a deep V-neckline,
diamond pendant necklace, hair swept up in a loose chignon, subtle smokey eye makeup,
standing in a grand marble ballroom with chandelier lighting,
same face, same freckles, same hazel-green eyes, same features exactly preserved
```

**Step 3 — Edit to another outfit** with the same editing model (from original):
```
The woman in the image is now wearing athletic gear, black sports bra and leggings,
hair in a high ponytail, light sweat on skin, boxing gloves hanging around her neck,
standing in a gritty boxing gym, overhead fluorescent lighting,
same face, same freckles, same hazel-green eyes, same features exactly preserved
```

### Character Sheet — Turnaround Reference

Use a high-quality image generation model (find via `search_models`).

```
Professional character turnaround reference sheet on clean white background,
five evenly-spaced views: front, three-quarter left, profile left, three-quarter back, back,

Maren, late 20s, warm olive skin, oval face, large hazel-green eyes,
dark auburn wavy shoulder-length hair with side part, athletic slender build,
freckles across nose and cheekbones,

wearing a fitted dark green utility jacket, grey henley shirt, dark jeans, brown boots,

consistent proportions and lighting across all five views,
character design concept art, clean technical reference sheet layout,
even studio lighting, no shadows on background, professional character design document
```

### Expression Sheet — 9 Emotions

Use a high-quality image generation model (find via `search_models`).

```
Character expression reference sheet, 3x3 grid layout on white background,
nine headshot portraits of the same woman showing different emotions,

Maren, late 20s, warm olive skin, oval face, large hazel-green eyes,
dark auburn wavy shoulder-length hair, freckles across nose,

top row left to right: neutral calm, warm genuine smile, laughing with eyes crinkled,
middle row: cold anger with clenched jaw, deep sadness with downcast eyes, wide-eyed surprise,
bottom row: disgusted wrinkled nose, terrified wide eyes, fierce determination,

consistent face across all nine, same lighting, same angle (front-facing),
character design expression sheet, clean presentation
```

### Character Across Scenes — Full Consistency Pipeline

**Generate hero image** with a high-quality image generation model, then use **an image editing model** for each scene:

Scene 1 — Morning coffee:
```
The woman in the image is sitting at a sunny kitchen table, morning light streaming through windows,
holding a ceramic mug of coffee, wearing a cozy oversized sweater, hair messy from sleep,
soft warm smile, domestic warm atmosphere, same face and all features preserved
```

Scene 2 — Office meeting:
```
The woman in the image is standing at a conference room whiteboard, mid-presentation,
wearing a tailored navy blazer and white blouse, hair neatly styled,
confident professional expression, modern office environment with glass walls,
same face and all features preserved
```

Scene 3 — Evening run:
```
The woman in the image is jogging along a waterfront path at sunset,
wearing running clothes, wireless earbuds, hair in a ponytail bouncing mid-stride,
focused athletic expression, golden hour light, city skyline in background,
same face and all features preserved
```

Scene 4 — Formal event:
```
The woman in the image is at a rooftop cocktail party at twilight,
wearing a deep emerald silk slip dress, delicate gold jewelry, hair in loose waves,
holding a champagne flute, soft city lights bokeh in background,
candlelit warm atmosphere, same face and all features preserved
```

---

## 14. Common Mistakes & Fixes

| Mistake | Fix |
|---------|-----|
| Vague description: "pretty woman with brown hair" | Use full anchor: specify face shape, eye shape/color, nose, lips, hair texture/length/style, skin tone, distinguishing marks |
| Changing anchor wording between prompts | Copy-paste identical text every time. Even "hazel eyes" vs "hazel-green eyes" causes drift |
| Too many changes in one Kontext edit | Change one major element per edit step. Scene OR outfit OR expression, not all three |
| LoRA strength too high (1.0+) | Start at 0.8, reduce to 0.7 if seeing artifacts or plastic-looking skin |
| Training LoRA with too few images | Use 15-30 images minimum with diverse angles and lighting |
| Training LoRA with watermarked images | Use clean, unprocessed images. No text overlays, no heavy filters |
| Not specifying camera/style consistently | Lock your "camera rig" phrase and repeat it: same lens, same film stock, same color grade |
| Expecting seed locking to work alone | Seed locking is a supplement. Combine with anchor description for best results |
| Using negative prompts for character traits | Describe what you want, not what you don't. "Clear smooth skin" not "no acne, no blemishes" |
| Forgetting to specify expression | Default is ambiguous neutral. Always state the emotion explicitly |

---

## 15. Character Design Checklist

Before generating, verify your character spec covers:

- [ ] **Name / identifier** — For your own tracking and LoRA trigger words
- [ ] **Age range** — Specific decade, not "young" or "old"
- [ ] **Skin tone** — Use specific vocabulary from Section 2
- [ ] **Face shape** — One of the seven types
- [ ] **Eyes** — Shape + color + any distinctive quality
- [ ] **Nose** — Type from vocabulary
- [ ] **Lips** — Type and any color note
- [ ] **Brows** — Shape, thickness, color
- [ ] **Hair** — Color + texture + length + style (4 attributes minimum)
- [ ] **Build** — Body type + height impression
- [ ] **Distinguishing marks** — At least 1-2 unique features (scars, freckles, moles, dimples, birthmarks)
- [ ] **Default expression** — The character's resting face
- [ ] **Default outfit** — The signature look
- [ ] **Color palette** — 2-3 colors associated with the character (optionally as HEX codes)
