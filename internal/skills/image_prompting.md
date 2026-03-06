# Image Prompting — AI Image Generation Master Guide

Complete reference for AI image generation prompt engineering on fal.ai (March 2026). Every endpoint, parameter, technique, and example prompt in this file is verified and actionable.

---

## Quick Reference

**Formula:** `[Subject] + [Action/Pose] + [Style/Medium] + [Lighting] + [Composition/Camera] + [Mood] + [Details]`

### Golden Rules

1. **Subject-first, always.** The first 20% of your prompt carries the highest weight. Put the most important visual information at the start.
2. **Be specific, not vague.** "85mm f/1.4" not "blurry background". "Rembrandt lighting" not "dramatic". "Crimson" not "reddish".
3. **Describe what you want, not what you don't.** Say "empty street" not "no people". Many modern models have no negative prompts.
4. **One style anchor per prompt.** Open with "editorial fashion photography" or "1970s Kodachrome snapshot" — it steers everything.
5. **40-80 words is the sweet spot.** Under 20 gives too much freedom. Over 120 dilutes focus and causes incoherence.
6. **Every word serves a visual purpose.** Remove filler words. Each token competes for model attention.

### Prompt Patterns

**Photorealistic portrait:**
`[Subject description], [pose/action], shot on [camera] with [lens], [film stock], [lighting], [mood]`

**Product photography:**
`Professional product photography of [product], [surface/background], [lighting setup], shot on Phase One IQ4 with 100mm macro at f/8, commercial quality`

**Cinematic / concept art:**
`[Character/scene], [environment], [atmospheric details], [lighting], [art direction], concept art, highly detailed, [color palette]`

**Text-in-image:**
`[Design description] with text "[YOUR TEXT]", [style], [colors], clean background`

**Image editing:**
`The [element] in the image is now [change]. Keep the [preserved elements] exactly the same.`

### Style Keyword Cheat Sheet

| Need | Keywords |
|------|----------|
| Warm portrait light | `Kodak Portra 400, golden hour, 85mm f/1.4` |
| Cinematic feel | `anamorphic lens flare, teal and orange grade, 2.39:1, film grain` |
| Moody / dramatic | `Rembrandt lighting, chiaroscuro, low-key, split lighting` |
| Clean product shot | `studio softbox, white sweep, Phase One IQ4, commercial grade` |
| Vintage analog | `Kodachrome 64, Kodak Tri-X 400, light leaks, film grain, 1970s` |
| Depth / subject pop | `85mm f/1.4, shallow depth of field, creamy bokeh, subject isolation` |
| Epic scale | `extreme wide shot, atmospheric perspective, bird's eye, vast, establishing` |
| Specific color | HEX codes work in some models: `"color #FF6B35"` — or use named colors: `"burnt sienna"` |

---

## 1. Universal Prompt Formula

```
[Subject] + [Action/Pose] + [Style/Medium] + [Lighting] + [Composition/Camera] + [Mood] + [Details]
```

### Golden Rules

1. **Subject-first.** All modern models weigh earlier tokens more heavily. The first 20% of your prompt defines the core subject (highest weight), the last 20% is fine details (lowest weight).
2. **Be specific.** "Crimson" not "reddish color". "85mm f/1.4" not "blurry background". "Rembrandt lighting" not "dramatic lighting".
3. **Every word serves a visual purpose.** Remove filler words ("a", "the", "very", "really"). Each token competes for model attention.
4. **Describe what you WANT, not what you don't want.** Most modern models have no negative prompts. Say "empty street" not "no people". Say "cloudless sky" not "no clouds".
5. **One style anchor per prompt.** Starting with "editorial fashion photography" or "1970s Kodachrome snapshot" steers everything that follows.
6. **40-80 words is the sweet spot.** Under 20 words gives the model too much freedom. Over 120 dilutes focus and causes incoherence.
7. **Use concrete references.** Camera bodies, lens focal lengths, film stocks, painter names, art movements are the most precise style controls available.
8. **Color specificity.** Some models support HEX codes for precise colors: `"primary color #FF6B35"`. Use named colors for broad compatibility: "burnt sienna", "cerulean blue", "sage green".
9. **Text rendering.** Place desired text in quotation marks: `"CAFE DEL SOL"`. Keep it to 1-4 words. Use search_models to find models with strong text rendering support.
10. **5-3-1 workflow.** Generate 5 quick variations on a fast/budget model, refine the best prompt 3 times, then do 1 final render on a premium model.

### Token Priority Map

```
Position in prompt → Weight influence:
  [0%-20%]   HIGHEST — Core subject, identity, main object
  [20%-40%]  HIGH    — Action, pose, key attributes
  [40%-60%]  MEDIUM  — Style, medium, artistic direction
  [60%-80%]  LOW     — Lighting, composition, camera
  [80%-100%] LOWEST  — Mood, fine details, quality boosters
```

Front-load the most important visual information. Quality boosters like "8K UHD" belong at the end.

---

## 3. Utility Models

Use `search_models` to find the best current utility models. Common categories:

### 3.1 Background Removal
Search for background removal models. Capabilities vary: some excel at complex edges (hair, fur, translucent objects), others are faster for simple shapes and bulk processing.

### 3.2 Upscaling
Search for image upscaling models. Options range from premium photo/face upscalers to fast 4x upscalers and specialized anime/illustration upscalers. Some models hallucinate detail (good for AI-generated images), others preserve original detail faithfully.

### 3.3 Segmentation
Search for image segmentation models (e.g., Segment Anything). These support point, box, or text prompts for isolating objects.

### 3.4 Outpainting
Search for image expansion/outpainting models to extend image boundaries in any direction.

### 3.5 Inpainting
Search for inpainting models for mask-based editing of specific image regions.

### 3.6 Virtual Try-On
Search for virtual try-on models to place garments onto person photos.

### 3.7 Face/Identity Preservation
Search for face identity preservation models that can generate new images while maintaining a person's likeness from a single reference photo.

---

## 4. Style Keywords Reference

### 4.1 Photorealistic

```
photorealistic, shot on Canon EOS R5, 85mm f/1.4, Kodak Portra 400,
detailed skin texture, natural light, sharp focus, 8K UHD, shallow depth of field,
film grain, RAW photo, high dynamic range
```

**Sub-genre variations:**
- **Editorial**: `editorial photography, Vogue, clean composition, fashion-forward, high fashion`
- **Documentary**: `documentary photography, candid, authentic, available light, Leica M6, 35mm`
- **Street**: `street photography, decisive moment, 35mm, urban, natural, unposed, Magnum`
- **Fine Art Portrait**: `fine art portrait, Rembrandt lighting, painterly, textured backdrop, gallery print`
- **Commercial Product**: `commercial photography, studio lighting, product-ready, Phase One IQ4, white sweep`
- **Architectural**: `architectural photography, tilt-shift lens, deep DOF, golden hour, lines and geometry`
- **Wildlife/Nature**: `wildlife photography, 400mm telephoto, shallow DOF, natural habitat, National Geographic`

### 4.2 Cinematic

```
cinematic, film grain, anamorphic lens flare, teal and orange color grading,
Rembrandt lighting, 2.39:1 aspect ratio, directed by Roger Deakins,
shallow depth of field, atmospheric haze, 35mm film
```

**Director signatures (use for distinct visual styles):**
- `directed by Roger Deakins` — naturalistic, precise, atmospheric, minimal color grading
- `directed by Emmanuel Lubezki` — long takes, natural light, ethereal, golden hour
- `directed by Wes Anderson` — symmetrical, pastel palette, whimsical, centered framing
- `directed by Denis Villeneuve` — vast scale, muted tones, stark, geometric, imposing
- `directed by Wong Kar-wai` — saturated, neon, motion blur, romantic, step-printed
- `directed by Ridley Scott` — smoke, backlighting, epic, textured, industrial
- `directed by David Fincher` — desaturated, precise, dark, meticulous framing
- `directed by Terrence Malick` — magic hour, handheld, nature, spiritual, dreamlike
- `directed by Stanley Kubrick` — one-point perspective, cold, symmetrical, unsettling

### 4.3 Concept Art / Illustration

```
digital painting, concept art, trending on artstation, highly detailed,
matte painting, painterly brushstrokes, epic composition, fantasy illustration,
dynamic lighting, rich color palette
```

**Sub-styles:**
- **Environment Concept**: `environment concept art, vast landscape, cinematic scale, atmospheric perspective, epic vista`
- **Character Design Sheet**: `character design sheet, T-pose, turnaround, clean white background, multiple views, orthographic`
- **Creature Design**: `creature concept art, anatomical detail, evolutionary plausibility, multiple angles`
- **Prop/Weapon Design**: `prop concept art, orthographic views, material callouts, clean background`
- **Vehicle Design**: `vehicle concept art, 3/4 view, sleek design, surface reflections`

### 4.4 Anime / Manga

```
anime style, cel shading, Studio Ghibli color palette, clean lineart,
vibrant colors, expressive eyes, best quality, masterpiece, detailed background
```

**Sub-styles:**
- **Studio Ghibli**: `Studio Ghibli, Hayao Miyazaki, watercolor sky, lush greenery, whimsical, warm`
- **Makoto Shinkai**: `Makoto Shinkai, photorealistic backgrounds, dramatic sky, lens flare, emotional, detailed clouds`
- **90s Anime**: `90s anime aesthetic, VHS quality, cel animation, retro color palette, Cowboy Bebop, scanlines`
- **Manga Panel**: `manga panel, black and white, screentone, dynamic composition, speed lines, high contrast`
- **Chibi**: `chibi style, super deformed, oversized head, cute, simplified features, kawaii`
- **Dark Anime**: `dark anime, Berserk style, gritty, detailed armor, dramatic shadows, mature`

### 4.5 3D Render

```
octane render, unreal engine 5, physically based rendering, ray tracing,
HDRI lighting, product visualization, ultra-detailed, subsurface scattering,
global illumination, photogrammetry quality
```

**Sub-styles:**
- **Isometric**: `isometric 3D render, miniature world, tilt-shift, clean edges, low-poly, diorama`
- **Claymation**: `claymation, stop motion, soft diffused lighting, handmade texture, Aardman style`
- **Glass/Crystal**: `glass render, caustics, refraction, transparent material, studio lighting, dispersion`
- **Pixar/Disney**: `Pixar 3D style, subsurface scattering on skin, expressive, warm lighting, appealing design`

### 4.6 Watercolor / Oil Painting

**Watercolor:**
```
watercolor painting, wet-on-wet technique, bleeding edges, paper texture visible,
soft washes, luminous transparency, delicate brushwork, granulation
```

**Oil Painting:**
```
oil painting, thick impasto brushstrokes, canvas texture, rich pigment,
gallery quality, chiaroscuro, classical technique, alla prima, glazing layers
```

**Gouache:**
```
gouache painting, matte finish, opaque layers, flat color areas, illustration quality,
vintage children's book illustration, warm muted palette
```

### 4.7 Minimalist / Graphic Design

```
minimalist, clean lines, flat design, bold typography, limited color palette,
negative space, geometric shapes, Swiss design, Helvetica, grid-based layout
```

**Sub-styles:**
- **Swiss/International**: `Swiss International Typographic Style, grid, Helvetica, clean, rational`
- **Japanese Minimalism**: `Japanese minimalism, wabi-sabi, muted earth tones, sparse, zen`
- **Bauhaus**: `Bauhaus design, primary colors, geometric, functional, universal`

### 4.8 Vintage / Retro

```
vintage photograph, 1970s, faded colors, light leaks, Kodachrome,
analog film grain, retro aesthetic, nostalgic, warm color cast
```

**By decade:**
- **1920s**: `Art Deco, jazz age, sepia tone, geometric patterns, gold and black accents, glamour`
- **1950s**: `mid-century modern, Technicolor, pastel palette, pin-up aesthetic, chrome, optimistic`
- **1960s**: `psychedelic, Pop Art, bold patterns, mod fashion, saturated primary colors`
- **1970s**: `Kodachrome, warm grain, earth tones, sun-faded, analog, groovy, brown and orange`
- **1980s**: `synthwave, neon, VHS aesthetic, chrome, grid lines, hot pink and cyan, Memphis design`
- **1990s**: `grunge, muted colors, disposable camera, flash photography, lo-fi, Nirvana aesthetic`
- **Y2K**: `cyber Y2K, metallic silver, butterfly motifs, glossy, translucent plastic, futuristic optimism`

---

## 5. Quality Boosters

### 5.1 Resolution Keywords
```
8K UHD, ultra high resolution, sharp focus, pixel-perfect, extremely detailed,
high-resolution scan, large format photography, medium format quality,
tack sharp, resolving fine detail
```

### 5.2 Detail Keywords
```
highly detailed, intricate details, hyper-detailed, razor sharp, fine detail,
meticulous, elaborate, ornate, precise, micro-detail visible
```

### 5.3 Texture Keywords
```
detailed skin texture, film grain, natural grain, subsurface scattering,
fabric weave visible, pore-level detail, surface imperfections, micro-texture,
material authenticity, tactile quality
```

### 5.4 Atmosphere Keywords
```
atmospheric, moody, ethereal, dreamy, haunting, serene, dramatic,
tension, melancholy, euphoric, contemplative, mysterious, ominous,
tranquil, energetic, nostalgic, intimate, vast, oppressive, liberating
```

### 5.5 Professional Quality Keywords
```
award-winning, gallery quality, museum print, exhibition quality,
professional photography, commercial grade, magazine cover quality,
Hasselblad Master, National Geographic photo of the year
```

---

## 6. Camera Bodies (Each Has a Distinct Look)

| Camera | Character | Best For |
|--------|-----------|----------|
| Canon EOS R5 | Sharp, accurate colors, clinical, reliable | Commercial, product, editorial |
| Canon EOS 5D Mark IV | Warm skin tones, versatile | Weddings, portraits, events |
| Sony A7IV | Clean, high dynamic range, neutral | Landscape, events, versatile |
| Sony A1 | Ultimate resolution, fast, precise | Sports, commercial, wildlife |
| Hasselblad 500C | Medium format, creamy tonality, depth | Fashion, fine art, portraits |
| Hasselblad X2D | Modern medium format, extreme detail | Commercial, landscape |
| Leica M6 | Organic character, classic, timeless | Street, documentary, editorial |
| Leica Q2 | Sharp, compact, character | Street, travel, candid |
| Phase One IQ4 | 150MP, ultimate detail, commercial grade | Architecture, commercial, fine art |
| Nikon Z9 | Punchy colors, sharp, reliable | Sports, wildlife, photojournalism |
| Fujifilm X-T5 | Film simulation colors, retro aesthetic | Lifestyle, travel, street |
| Fujifilm GFX 100S | Medium format, Fuji colors, detail | Portrait, landscape, commercial |
| Mamiya RZ67 | Medium format, smooth bokeh, classic | Portrait, fashion, editorial |
| Contax T2 | Warm tones, point-and-shoot charm, Zeiss | Casual, lifestyle, candid |
| Polaroid SX-70 | Instant film, soft, nostalgic, white border | Retro, artistic, personal |
| Rolleiflex 2.8F | Square format, waist-level, classic | Street, portrait, fine art |

---

## 7. Lenses

| Focal Length | Character | Best For |
|--------------|-----------|----------|
| 14mm f/2.8 | Ultra-wide, dramatic distortion, expansive | Architecture, interiors, dramatic landscapes |
| 20mm f/1.4 | Wide, immersive, environmental | Real estate, environmental, astrophotography |
| 24mm f/1.4 | Wide angle, environmental context, slight distortion | Architecture, environmental portraits, street |
| 35mm f/1.4 | Classic street perspective, versatile | Street, documentary, reportage, editorial |
| 50mm f/1.2 | Natural perspective, "normal" lens, dreamy wide open | General, street, documentary, low light |
| 50mm f/1.8 | Natural perspective, affordable classic | General, street, documentary |
| 85mm f/1.4 | Portrait king, beautiful bokeh, subject isolation | Portraits, fashion, beauty, headshots |
| 100mm f/2.8 macro | Extreme close-up, infinite detail | Product, nature, jewelry, food, insects |
| 105mm f/1.4 | Maximum subject isolation, dreamy | Portraits, fashion, fine art |
| 135mm f/2 | Compressed background, dreamy bokeh | Portraits, events, candid, editorial |
| 200mm f/2 | Massive compression, paper-thin DOF | Sports, wildlife, editorial |
| 300mm f/2.8 | Strong telephoto compression | Wildlife, sports |
| 400mm f/2.8 | Extreme telephoto, maximum isolation | Wildlife, sports, nature |
| 24mm tilt-shift | Selective focus, miniature effect, corrected verticals | Architecture, creative, product |
| 90mm tilt-shift | Selective focus plane, product isolation | Product, food, tabletop |

---

## 8. Film Stocks

| Film Stock | Character | Best For |
|------------|-----------|----------|
| Kodak Portra 160 | Ultra-smooth, fine grain, subtle warmth | Studio portraits, fashion, beauty |
| Kodak Portra 400 | Warm, flattering skin tones, versatile | Portraits, weddings, fashion, everyday |
| Kodak Portra 800 | Grainier Portra, works in low light | Indoor events, moody portraits, available light |
| Kodak Ektar 100 | Saturated, fine grain, vivid colors | Landscape, travel, product, bright daylight |
| Kodachrome 64 | Saturated, nostalgic warmth, punchy reds/blues | Vintage, editorial, street, travel |
| Kodak Gold 200 | Consumer warmth, accessible, nostalgic | Everyday, casual, snapshot aesthetic |
| Kodak Tri-X 400 | Contrasty B&W, gritty grain, iconic | Street, documentary, editorial, journalism |
| Kodak T-Max 400 | Fine-grain B&W, smooth tones | Studio B&W portraits, fine art |
| Ilford HP5 Plus 400 | Softer B&W, flexible, pushable | General B&W, portraits, street |
| Ilford Delta 3200 | Ultra-grainy B&W, atmospheric, moody | Night, concert, moody, low-light |
| Ilford FP4 Plus 125 | Fine grain B&W, smooth, detailed | Landscape B&W, architecture, still life |
| Cinestill 800T | Tungsten balance, red halation, neon glow | Night photography, neon, urban, cinematic |
| Cinestill 50D | Daylight cinema film, fine grain, rich color | Daylight portraits, golden hour, cinematic |
| Fuji Pro 400H | Soft pastels, muted, dreamy (discontinued) | Weddings, lifestyle, delicate, soft |
| Fuji Superia 400 | Green-shifted, nostalgic, consumer | Casual, snapshot aesthetic, everyday |
| Fuji Velvia 50 | Extreme saturation, vivid, slide film | Landscape, nature, sunrise/sunset |
| Fuji Provia 100F | Balanced slide film, accurate, neutral | Nature, product, accurate color |
| Polaroid 600 | Instant, soft, faded, white border, unique color | Creative, retro, personal, artistic |
| Lomography 800 | Cross-processed, wild colors, unpredictable | Experimental, artistic, playful |
| AgfaPhoto Vista 200 | European consumer, cool tones, blue shift | Travel, everyday, nostalgic |

---

## 9. Lighting Keywords

### 9.1 Natural Light
```
golden hour        — warm amber glow, 15-30 min before sunset, long soft shadows
blue hour          — cool twilight, 15-30 min after sunset, deep blue/purple sky
magic hour         — cinematic warmth, film industry term for golden hour
overcast diffused  — soft, even, no harsh shadows, cloudy sky, flattering
harsh midday sun   — hard shadows, high contrast, bleached highlights, graphic
dappled light      — light through leaves/blinds, pattern of light and shadow
window light       — soft directional, Vermeer-like, one-sided, interior, intimate
open shade         — soft, cool, even lighting in shadow of building/tree
reflected sunlight — warm bounce from walls/sand, fills shadows, golden
backlit sunrise    — halo effect, warm rim, lens flare, silhouette potential
```

### 9.2 Studio Lighting
```
softbox            — large, soft, even, wrapping light, commercial standard
key light          — main light source, defines form and dimension
fill light         — secondary, reduces shadow contrast, opens shadows
rim light          — edge/separation light, outlines subject against background
hair light         — highlights hair from behind, adds dimension
beauty dish        — punchy yet soft, specular quality, fashion/beauty standard
ring light         — even frontal, circular catchlight in eyes, beauty/YouTube
strip light        — narrow, elongated highlight, product photography, edge accent
butterfly lighting — glamour classic, shadow under nose, beauty/fashion (Paramount)
clamshell lighting — beauty dish above + reflector below, beauty commercial
para lighting      — parabolic modifier, focused center, soft edges, editorial
```

### 9.3 Dramatic Lighting
```
Rembrandt lighting — triangle of light on shadow-side cheek, classic portrait
split lighting     — half face lit, half in shadow, moody, dramatic, mysterious
chiaroscuro        — extreme light/dark contrast, Caravaggio style, painterly
low-key            — predominantly dark, selective highlights, noir, dramatic
high-key           — bright, minimal shadows, airy, clean, optimistic
silhouette         — subject backlit, dark shape against bright background
spotlight          — focused beam, theatrical, isolated subject, dark surround
noir lighting      — hard shadows, venetian blinds, high contrast B&W, mystery
cross lighting     — two sources from opposite sides, sculptural, dramatic
```

### 9.4 Atmospheric / Environmental
```
volumetric light       — visible light beams through fog/dust/smoke, atmospheric
god rays               — sunbeams through clouds or canopy, crepuscular rays
light shafts           — beams of light in dark space, cathedral/forest effect
neon glow              — colored light, urban night, cyberpunk, signs reflected in wet streets
candlelight            — warm, flickering, intimate, orange, low-key
firelight              — warm, dynamic dancing shadows, campfire, primal
moonlight              — cool, blue-white, subtle, night, romantic
bioluminescence        — self-lit organic glow, underwater, fantasy, ethereal
fog + backlight        — milky atmosphere, mysterious, diffused, horror/romance
dust particles in light — visible motes, warm beam, nostalgic, aged space
```

### 9.5 Light Direction
```
backlit         — light behind subject, halo/rim effect, dramatic silhouette potential
side-lit        — light from 90 degrees, reveals texture, dramatic shadows
top-lit         — light from above, deep eye shadows, theatrical, overhead sun
under-lit       — light from below, eerie, horror, campfire story, unnatural
front-lit       — flat, even, direct flash, paparazzi, passport photo
three-quarter   — classic, light at 45 degrees high, natural and flattering
cross-lit       — light from both sides, sculpted, editorial
Kicker          — accent light from behind and to the side, separation and edge
```

### 9.6 Light Quality
```
soft diffused       — gentle, wrapping, minimal shadows, overcast quality
hard directional    — sharp shadows, defined edges, midday sun quality
specular            — bright highlights, reflections, glossy surfaces
ambient             — overall environmental light, no strong direction
reflected / bounced — bounced light, warm/cool depending on surface color
mixed lighting      — multiple color temperatures, tungsten + daylight, urban
practical lighting  — in-frame light sources (lamps, screens, candles)
motivated lighting  — studio light mimicking a natural source visible in scene
```

---

## 10. Composition Keywords

### 10.1 Fundamental Compositions
```
rule of thirds        — subject at 1/3 intersection points, balanced dynamic
golden ratio          — fibonacci spiral, organic placement, naturally pleasing
centered composition  — symmetrical, formal, powerful, confrontational
diagonal composition  — dynamic energy, tension, movement, leads the eye
frame within frame    — doorway, window, arch, tunnel framing the subject
leading lines         — roads, rails, paths, rivers drawing eye to subject
negative space        — large empty areas around subject, minimal, breathing room
fill the frame        — subject filling entire frame, intimate, impactful
symmetry              — mirror balance, architectural, satisfying, formal
```

### 10.2 Camera Angles
```
eye level           — natural, relatable, neutral, conversational
low angle           — looking up at subject, powerful, heroic, imposing, tall
high angle          — looking down at subject, vulnerable, diminished, small
bird's eye view     — directly overhead, pattern, map-like, graphic, God's view
worm's eye view     — extreme low angle, dramatic, towering, monumental
Dutch angle         — tilted horizon, unease, dynamic energy, tension, disorientation
over-the-shoulder   — voyeuristic, narrative, cinematic, intimate
POV (point of view) — first person perspective, immersive, subjective
```

### 10.3 Shot Types
```
extreme close-up (ECU) — eyes, lips, texture details only, macro-like
close-up (CU)          — face fills frame, emotional, intimate
medium close-up (MCU)  — head and shoulders, conversational
medium shot (MS)       — waist up, standard dialogue framing
medium wide (MW)       — knees up, action visible with context
full shot (FS)         — entire body visible, some environment
wide shot (WS)         — subject in environment, establishing context
extreme wide (EWS)     — vast landscape, subject small, epic scale
establishing shot      — wide, sets location, mood, context, opening shot
detail shot            — isolated object/texture, narrative emphasis
```

### 10.4 Depth and Layering
```
shallow depth of field — subject sharp, background blurred (creamy bokeh)
deep depth of field    — everything sharp front-to-back, f/11+, landscape
foreground element     — object in front of subject, adds depth and dimension
depth layering         — distinct foreground, midground, background planes
atmospheric perspective — distant objects lighter, bluer, hazier, adds depth
bokeh                  — out-of-focus highlights, circular/hexagonal, dreamy
tilt-shift effect      — selective focus plane, miniature world illusion
```

---

## 11. Art Movements Quick Reference

### 11.1 Classical

| Movement | Key Artist | Visual Signature |
|----------|-----------|-----------------|
| Renaissance | Da Vinci, Raphael | Sfumato, linear perspective, divine proportion, idealized |
| Baroque | Caravaggio, Rembrandt | Dramatic chiaroscuro, rich detail, dynamic composition |
| Rococo | Fragonard, Boucher | Pastel, ornate, playful, delicate, aristocratic |
| Romanticism | Turner, Friedrich | Sublime nature, emotion, dramatic skies, vast scale |
| Pre-Raphaelite | Waterhouse, Rossetti | Vivid color, medieval themes, extreme detail, poetic |
| Art Nouveau | Mucha, Klimt | Organic curves, floral motifs, gold leaf, decorative borders |
| Art Deco | Erte, Cassandre | Geometric, gold/black, streamlined, luxury, machine-age |
| Ukiyo-e | Hokusai, Hiroshige | Flat color, bold outlines, woodblock texture, waves, landscapes |

### 11.2 Modern

| Movement | Key Artist | Visual Signature |
|----------|-----------|-----------------|
| Impressionism | Monet, Renoir | Visible brushstrokes, light and color, plein air, soft |
| Post-Impressionism | Van Gogh, Cezanne | Bold color, structure, emotional intensity, thick paint |
| Expressionism | Munch, Kirchner | Distortion, raw emotion, intense color, anxiety |
| Surrealism | Dali, Magritte | Dream logic, impossible scenes, juxtaposition, melting |
| Cubism | Picasso, Braque | Multiple perspectives, fragmented, geometric, angular |
| Pop Art | Warhol, Lichtenstein | Bold, commercial, halftone dots, bright, repetition |
| Minimalism | Judd, Flavin | Reduction, geometry, essential forms, monochrome |
| Abstract Expressionism | Pollock, Rothko | Gesture, color field, emotional, large scale |

### 11.3 Genre Aesthetics

| Genre | Visual Signature |
|-------|-----------------|
| Cyberpunk | Neon, rain-slicked streets, augmentation, megacities, teal/magenta, holograms |
| Steampunk | Brass, gears, Victorian fashion, goggles, airships, mechanical, copper |
| Solarpunk | Green technology, organic architecture, optimistic, gardens on buildings |
| Vaporwave | Pastel gradients, Greek statues, 90s nostalgia, glitch, pink/purple/cyan |
| Synthwave | Neon grid, chrome, sunset, retro-futuristic, 80s, DeLorean, laser |
| Dark Academia | Libraries, tweed, candlelight, classical, autumn, gothic architecture |
| Cottagecore | Pastoral, wildflowers, linen, handmade, cozy, cottage, baking |
| Gothic | Dark, ornate, cathedrals, shadows, romantic decay, ravens, moonlight |
| Afrofuturism | African motifs, sci-fi, vibrant patterns, futuristic, cultural pride |
| Dieselpunk | 1940s, industrial, military vehicles, Art Deco machines, propaganda poster |
| Atompunk | 1950s-60s futurism, atomic age, Jetsons, chrome, tailfins, optimistic |
| Solarpunk | Sustainable tech, biomimicry, green cities, hopeful, organic + digital |

---

## 12. Character Consistency Techniques

### 12.1 Anchor Description Method
Create a fixed character block and reuse it verbatim across all prompts:
```
ANCHOR: "A 30-year-old woman with shoulder-length auburn hair, green eyes,
light freckles across her nose, wearing a cream wool turtleneck and brown leather jacket"
```

**Rules:**
- Repeat the EXACT phrases every time — "shoulder-length auburn hair" not "reddish medium hair"
- Lock camera setup (same focal length, same distance, same angle)
- Specify left/right for asymmetric features: "scar above left eyebrow"
- Only change the action/setting between images
- Include unique identifying details (specific jewelry, tattoo, accessory)

### 12.2 LoRA Fine-Tuning
- Train on 15-30 high-quality, diverse images of the subject
- Use `search_models` to find models that support LoRA training and inference
- Use a unique trigger word: `ohwx_person`, `sks_style`, `xyz_character`
- Weight scale: 0.6-0.9 for character likeness, 0.4-0.7 for artistic styles
- Include diverse angles, expressions, and lighting in training images
- 15 images minimum, 30 images ideal for faces

### 12.3 Reference Images
- Use `search_models` to find models that support reference images for consistency
- Some models accept edit-mode reference images to describe changes
- Some models support 2-4 reference images for maximum consistency across scenes
- IP-Adapter models provide single reference, no training, instant style/character transfer
- Face identity preservation models maintain likeness from one photo

### 12.4 Seed Locking
- Same seed + identical character description = very similar face/pose
- Only change the action, expression, or setting between generations
- Works most reliably with diffusion-based models that support seed parameters
- Not guaranteed to be consistent across different model versions or endpoints
- Combine with anchor description for best results

---

## 13. Negative Prompts

### Which Models Support Negative Prompts?

Some models support negative prompts while many modern models do not. Check the model documentation or use search_models to find models with negative prompt support.

For models without negative prompts, rephrase positively:
- Instead of "no blur" -> use "sharp focus, crisp details"
- Instead of "no people" -> use "empty scene, deserted"
- Instead of "no watermark" -> use "clean image, pristine"
- Instead of "no artifacts" -> use "flawless, perfect quality"

### 13.1 Universal Negative (for models with negative prompt support)
```
ugly, poorly drawn hands, poorly drawn face, deformed, blurry, watermark,
signature, low quality, extra limbs, bad anatomy, disfigured, mutated,
jpeg artifacts, text, error, cropped, worst quality, low quality, normal quality
```

### 13.2 Photorealistic Negative
```
drawing, painting, cartoon, 3d render, CGI, blurry, noise, jpeg artifacts,
overexposed, underexposed, oversaturated, plastic skin, mannequin, doll-like,
illustration, digital art, anime
```

### 13.3 Portrait Negative
```
deformed face, extra fingers, mutated hands, cross-eyed, bad anatomy,
asymmetric eyes, unnatural skin tone, waxy skin, dead eyes, uncanny valley,
plastic surgery look, over-smoothed, airbrushed
```

### 13.4 Anime/Illustration Negative
```
photorealistic, photograph, 3d render, low quality, bad anatomy,
deformed, extra limbs, blurry, watermark, signature, worst quality,
realistic face, photo, 3D
```

### 13.5 Landscape/Architecture Negative
```
people, figures, text, watermark, frame, border, signature, oversaturated,
HDR artifacts, chromatic aberration, lens distortion, tilted horizon,
blurry foreground, overprocessed
```

---

## 14. Example Prompts by Category

### 14.1 Product Photography

**Luxury Watch:**
```
Professional product photography of a luxury chronograph watch with brushed
titanium case and midnight blue dial, floating on a dark reflective surface,
dramatic side lighting from camera-left revealing surface textures and brushed
metal finish, shot on Phase One IQ4 with 100mm macro at f/8, ultra-sharp,
commercial quality, dark moody background with subtle blue gradient
```

**Wireless Earbuds:**
```
Professional product photography of wireless earbuds, matte white finish with
rose gold accents #B76E79, floating on clean white surface, soft studio lighting
from above and left with gentle shadow, shot on Phase One IQ4 with 100mm macro
at f/8, ultra-sharp, commercial quality, minimalist, Apple-style product shot
```

**Skincare Bottle:**
```
High-end cosmetics product shot of a frosted glass serum bottle with gold dropper
cap, morning light streaming through a sheer curtain behind, fresh dewy leaves
and water droplets as props, shallow depth of field, shot on Hasselblad with
120mm, beauty lighting, soft, luxurious, editorial, Sephora campaign quality
```

**Sneaker:**
```
Dynamic sneaker product shot, white and neon green running shoe suspended mid-air
against deep black background, dramatic rim lighting from both sides revealing
mesh texture and sole details, slight motion blur on floating dust particles,
shot on Canon EOS R5 with 85mm macro, commercial photography, Nike campaign style
```

### 14.2 Cinematic Portrait

**Weathered Fisherman:**
```
Cinematic portrait of a weathered fisherman mending nets at dawn, deep wrinkles
and sun-weathered skin, golden hour light hitting from camera-right, shot on
Leica M6 with 35mm f/1.4, Kodak Tri-X 400, shallow depth of field, documentary
photography, film grain, harbor background out of focus, authentic, contemplative
```

**Noir Detective:**
```
Film noir portrait of a detective in a dimly lit office, fedora casting shadow
across eyes, cigarette smoke curling upward catching a sliver of light, Venetian
blind shadows striping across the wall and face, black and white, Kodak Tri-X
400, 50mm f/1.4, hard chiaroscuro lighting, 1940s atmosphere, heavy grain
```

**Fashion Editorial:**
```
High fashion editorial portrait of a model in an oversized structured blazer
in electric blue #0047AB, standing in brutalist concrete architecture, overcast
flat light creating soft shadows, shot on Hasselblad 500C with 80mm f/2.8,
Kodak Portra 160, muted background, strong geometric composition, Vogue Italia
```

**Elderly Hands:**
```
Extreme close-up of elderly hands holding a worn leather-bound book, visible
veins and age spots, warm afternoon window light from camera-left, shallow
depth of field, shot on Canon EOS R5 with 100mm macro at f/2.8, Kodak Portra
400, documentary, intimate, time and wisdom, fine detail in skin texture
```

### 14.3 Fantasy Character

**Elven Warrior Queen:**
```
Ancient elven warrior queen in ornate silver armor with living vines growing
through the metalwork, glowing emerald eyes, misty enchanted forest at dawn,
volumetric light streaming through ancient trees, concept art, highly detailed,
painterly brushstrokes, epic composition, rich color palette
```

**Cyberpunk Street Samurai:**
```
Cyberpunk street samurai crouching on a rain-slicked rooftop at night, holographic
katana crackling with energy, cybernetic arm with exposed chrome mechanisms, neon
signs reflecting in puddles below, teal and magenta lighting, anamorphic lens
flare, concept art, hyper-detailed, atmospheric rain, volumetric fog
```

**Dark Fantasy Necromancer:**
```
Dark fantasy necromancer in tattered black robes standing in a decrepit cathedral,
green spectral energy swirling from outstretched skeletal hands, skull-adorned
staff, stained glass windows casting colored light through dust motes, gothic
architecture, dramatic low-angle composition, dark moody atmosphere, concept art
```

**Steampunk Inventor:**
```
Steampunk inventor in a cluttered workshop, brass goggles pushed up on forehead,
leather apron covered in oil stains, surrounded by gears and clockwork mechanisms,
warm gaslight illumination, copper and bronze color palette, detailed mechanical
contraptions, Victorian aesthetic, digital painting, intricate details
```

### 14.4 Architecture

**Modern Hillside Residence:**
```
Modern hillside residence at golden hour, cantilevered concrete and floor-to-ceiling
glass, infinity pool reflecting amber sky, Mediterranean landscaping with olive
trees, 24mm tilt-shift lens, architectural photography, warm interior lights
glowing through windows, aspirational, sharp throughout, Architectural Digest
```

**Gothic Cathedral Interior:**
```
Interior of a Gothic cathedral with soaring ribbed vaults, sunlight streaming
through an enormous rose window casting colorful light patterns on stone floor,
atmospheric dust particles floating in light beams, 14mm ultra-wide lens, deep
depth of field, symmetrical composition down the nave, awe-inspiring vertical
scale, architectural photography, sacred atmosphere
```

**Brutalist Parking Structure:**
```
Abandoned brutalist parking structure, raw concrete spiraling upward, single
shaft of golden afternoon light cutting through an opening, long shadows on
stained concrete, 24mm lens, deep focus, symmetrical composition, urban decay,
moody, industrial beauty, shot on Leica Q2, geometric patterns
```

### 14.5 Food Photography

**Artisan Pasta:**
```
Overhead food photography of handmade pappardelle pasta with wild mushroom ragu,
shaved Parmigiano-Reggiano, fresh thyme garnish, on a rustic ceramic plate,
reclaimed wood table, warm natural window light from upper left, shallow depth
of field at edges, shot on Sony A7IV with 50mm f/1.4, editorial food photography,
appetizing, Bon Appetit style
```

**Coffee Pour:**
```
Close-up of espresso being poured into a ceramic cup, rich crema forming on
surface, steam rising and catching backlight, barista hands visible at top of
frame, warm cafe lighting, shallow depth of field, shot on Canon EOS R5 with
85mm f/1.4, moody earth tones, coffee shop atmosphere, texture of liquid surface
```

**Sushi Platter:**
```
Overhead view of an elaborate sushi platter on a black slate board, fresh nigiri
and maki rolls with glistening fish, wasabi and pickled ginger as accents, single
chopstick pair resting at edge, soft diffused natural light, clean negative space,
shot on Fujifilm GFX with 63mm, Japanese minimalism, premium, Monocle magazine
```

### 14.6 Fashion Editorial

**Streetwear:**
```
Fashion editorial, model in oversized vintage denim jacket and wide-leg trousers,
standing in empty parking garage with concrete pillars, harsh fluorescent overhead
light mixed with warm sunset through entrance, strong shadow play, shot on Contax
T2, Kodak Portra 800, full body, editorial pose, i-D Magazine aesthetic, raw
```

**Haute Couture:**
```
Haute couture fashion photography, model in an avant-garde sculptural white gown
with dramatic ruffled layers, standing in a minimalist white gallery space, soft
diffused overhead lighting, shot on Hasselblad with 120mm, clean background,
architectural silhouette, high fashion, Vogue aesthetic, editorial pose
```

**Athleisure Campaign:**
```
Athletic fashion campaign, model in motion running through desert landscape at
sunset, flowing activewear in terracotta and sand tones, hair and fabric catching
wind, golden hour backlighting creating halo effect, shot on Sony A1 with 70-200mm
at 200mm f/2.8, motion blur on feet, sharp on face, Nike campaign quality
```

### 14.7 Abstract Art

**Color Field:**
```
Abstract color field painting in the style of Mark Rothko, two large rectangular
blocks of deep crimson #8B0000 and burnt orange #CC5500 on a warm ochre background,
soft edges where colors meet and bleed, canvas texture visible, oil paint, luminous
color depth, contemplative, meditative, gallery quality, large scale
```

**Geometric:**
```
Abstract geometric composition, interlocking triangles and circles in navy blue
#1B365D and gold #C5A558 on cream background, clean precise edges, Bauhaus
inspired, screenprint texture, limited color palette, modernist, balanced
asymmetry, graphic design quality, grid-based structure
```

**Fluid Art:**
```
Abstract fluid art, swirling marble-like patterns in deep teal #006D6F and
metallic gold #D4AF37 on black background, paint pouring technique, organic
cellular patterns, resin-coated glossy finish, macro photography detail,
mesmerizing, contemporary gallery art
```

### 14.8 Typography / Logo Design

**Cafe Signage (use a model with strong text rendering, e.g., search for text-in-image models):**
```
Vintage hand-painted cafe signage reading "CAFE DEL SOL" in warm terracotta
and cream colors, ornate serif font with decorative flourishes, aged wood
texture background, Mediterranean style, artisan quality, warm lighting,
nostalgic, hand-lettered character
```

**Tech Logo (use a logo generation model with style "logo_raster" or similar):**
```
Modern minimalist logo for a technology company, the letters "NOVA" in clean
geometric sans-serif font, gradient from deep blue #1a1a4e to electric purple
#7b2ff7, isolated on pure white background, vector quality, crisp edges,
balanced letterforms, professional, Silicon Valley aesthetic
```

**Poster (use a model with strong text rendering):**
```
Vintage travel poster for "TOKYO" in bold Art Deco typography at the top,
Mount Fuji and cherry blossoms in background, bullet train in foreground,
limited color palette of red #C41E3A, cream, and navy, retro screen-print
texture, 1960s Japanese tourism poster style, clean graphic design
```

### 14.9 Landscape / Nature

**Mountain Dawn:**
```
Majestic mountain landscape at dawn, jagged snow-capped peaks above a cloud
inversion layer, first golden light hitting summits turning them amber, shot
on Sony A7IV with 24-70mm at 24mm f/11, deep depth of field, Fuji Velvia 50
color saturation, foreground alpine wildflowers, epic scale, pristine wilderness,
National Geographic quality
```

**Underwater:**
```
Underwater photograph of a coral reef at noon, sunbeams penetrating turquoise
water from directly above, colorful tropical fish school around brain coral and
purple sea fans in foreground, crystal clear water, shot on Nikon Z9 in underwater
housing, wide angle 14mm, vibrant, teeming with life, BBC Earth quality, blue
```

**Desert Storm:**
```
Dramatic desert landscape with approaching sandstorm, massive wall of dust
advancing across red sand dunes, single twisted dead tree in foreground, last
golden light before storm hits, deep contrast between warm sand and dark storm
cloud, shot on Canon EOS R5 with 35mm, documentary landscape, ominous, epic scale
```

### 14.10 Interior Design

**Scandinavian Living Room:**
```
Scandinavian interior design photography of a minimalist living room, large windows
with sheer linen curtains, warm oak flooring, pale gray linen sofa with cream
throw pillows, single monstera plant in ceramic pot, natural daylight filling
the space, clean and airy, shot on Canon EOS R5 with 24mm tilt-shift, Architectural
Digest quality, hygge atmosphere, calming
```

**Japanese Tea Room:**
```
Traditional Japanese tea room (chashitsu), tatami mats, tokonoma alcove with
single ikebana arrangement, shoji screens filtering soft natural light, wooden
beams, wabi-sabi aesthetic, minimal and intentional, shot on Fujifilm GFX with
23mm, peaceful, meditative, earth tones, tea ceremony in progress
```

---

## 16. Advanced Techniques

### 16.1 Multi-Step Quality Workflow
```
Step 1: Generate 4 variations with a fast/budget model (explore concepts)
Step 2: Pick best, refine prompt, regenerate with a premium model (quality)
Step 3: Remove background with a background removal model (if needed)
Step 4: Upscale 4x with an upscaling model (final resolution boost)
```

### 16.2 Style Transfer Pipeline
```
Step 1: Generate a reference image in your desired style
Step 2: Use as style reference via an IP-Adapter model (search for IP-Adapter support)
Step 3: Generate new subjects in that consistent style with different prompts
```

### 16.3 Product Turnaround
```
Step 1: Generate hero shot with a premium image generation model
Step 2: Use a multi-angle generation model for additional viewpoints
Step 3: Remove all backgrounds with a background removal model
Step 4: Upscale with an upscaler suited for design work (sharp edges)
```

### 16.4 Consistent Character Series
```
Step 1: Generate initial character with detailed anchor description + locked seed
Step 2: Verify consistency by regenerating with same seed
Step 3: Use a multi-reference model with 2-3 reference images for new scenes
Step 4: If doing 10+ images, train a LoRA for best consistency
```

### 16.5 Image-to-Image Refinement
```
Step 1: Quick text-to-image draft (use a fast/budget model)
Step 2: Image-to-image refinement (use a model with I2I support)
Step 3: Targeted editing (use an image editing model)
Step 4: Final upscale (photo upscaler for photos, AI art upscaler for AI-generated images)
```

### 16.6 Design Asset Pipeline
```
Step 1: Generate illustration with a model supporting illustration styles (style: "digital_illustration")
Step 2: Vectorize with a vectorization model for SVG output
Step 3: Or decompose with a layered output model for PSD-like output
```

---

## 17. Common Mistakes and Fixes

| Mistake | Fix |
|---------|-----|
| Using `(text:1.3)` weighting in modern models | Remove weights, use natural emphasis: "with particular focus on..." |
| Adding negative prompts to models that don't support them | Remove them, rephrase positively |
| Prompt over 150 words | Cut to 40-80 words, remove filler |
| "Beautiful, stunning, amazing" as descriptors | Replace with specific visual descriptors (camera, light, color) |
| "Realistic" as only style cue | Specify camera body + lens + film stock + lighting |
| Too many style references competing | Pick ONE style anchor, be consistent throughout |
| Low resolution + no upscale | Generate at native res, upscale with an upscaling model |
| Wrong model for text rendering | Switch to a model with strong text rendering (use search_models) |
| Inconsistent characters across images | Use anchor description + multi-reference model + optional LoRA |
| Generic lighting description | Specify direction, quality, color temperature, and type |
| Forgetting `"quotes"` for text | Always quote desired text: `"OPEN 24 HOURS"` |
| Using premium model for brainstorming | Use a fast/budget model for iteration, premium model for finals |
| Ignoring composition keywords | Add shot type + camera angle + composition rule |
| All prompts starting the same | Vary your style anchor to explore different looks |

---

## 18. Prompt Templates

### Quick-Start Template (Modern Models)
```
[Style anchor], [subject with specific details], [action or pose],
[lighting type and direction], [camera/lens if photographic],
[setting/environment], [mood/atmosphere], [output quality keywords]
```

### Template for Models with Weights and Negatives
```
Prompt: (main subject:1.2), detailed description of appearance, style keywords,
(lighting type:1.1), composition and framing, quality boosters

Negative: ugly, deformed, blurry, watermark, low quality, (specific unwanted elements:1.3)
```

### Editing Template (Image Edit Models)
```
Change [specific element] from [current state] to [desired state],
while preserving [list everything that must stay unchanged: expression,
pose, background, lighting, clothing, camera angle]
```

### Product Photography Template
```
Professional product photography of [product with exact materials and colors],
on [surface/background description], [lighting setup with direction],
shot on [camera body] with [lens] at [aperture], [composition style],
commercial quality, [mood/brand feeling]
```

### Portrait Template
```
[Portrait style] portrait of [subject with age, features, expression],
[clothing and accessories], [setting/environment], [lighting with direction],
shot on [camera] with [focal length] at [aperture], [film stock],
[depth of field], [mood], [quality keywords]
```

### Concept Art Template
```
[Genre] concept art of [subject with detailed description], [action/pose],
[environment with atmosphere], [lighting — volumetric, dramatic, etc.],
[color palette], [composition], highly detailed, painterly, epic
```

### Output Format Selection
- **JPEG**: Smaller file, lossy, good for photos. Best for web and social media.
- **PNG**: Larger file, lossless, supports transparency. Best for design assets and compositing.
- **WebP**: Modern efficient format, good compression with quality. Best for web delivery.
