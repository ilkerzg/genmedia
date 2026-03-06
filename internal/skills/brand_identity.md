# Brand Identity — Visual Brand System Master Guide

Comprehensive knowledge base for creating logos, brand systems, color palettes, typography, and visual identity assets using AI generative models on fal.ai. Covers design theory, industry conventions, and ready-to-use prompt templates.

---

## Quick Reference

> **Model Discovery:** Use `search_models` to find the best model for each task. Search by category: "text to image with text" for logos with text/wordmarks, "image generation" for abstract symbols and mockups, "image editing" for iterating on assets, "vectorize" for SVG conversion, "background removal" for transparent PNGs, "upscale" for print-ready output.

### Top 5 Brand Identity Tips

1. **Always specify exact HEX codes** — Never use vague color names alone. Include explicit values like `#1B3A5C`.
2. **Put all text in quotes** — AI text rendering is most accurate when target string is quoted: `"BRAND NAME"`.
3. **Design for monochrome first** — A logo that works in single-color black works everywhere.
4. **Use a style anchor block** — Create a reusable descriptor (palette, typography, mood) and prepend it to every prompt.
5. **Limit your palette strictly** — Use `strict three-color palette` or `limited palette, no other colors` to prevent drift.

### Key Prompt Patterns

**Logos:** `[Style] logo for "BRAND NAME", [logo type], [typography desc], [color #HEX], flat vector style on [background] background, professional brand identity`

**Color Systems:** `strict palette of [color 1] #HEX, [color 2] #HEX, [color 3] #HEX — no other colors, [mood] brand identity`

### Cheat Sheet — Prompt Cores by Goal

| Goal | Prompt Core |
|------|-------------|
| Wordmark logo | `minimal wordmark logo for "NAME", geometric sans-serif, flat vector, white background` |
| Icon/symbol mark | `minimal abstract [shape] symbol, single color #HEX on white, scalable icon, no text` |
| Emblem / badge | `circular emblem badge for "NAME EST. YEAR", vintage style, single color #HEX` |
| Business card mockup | `premium business card mockup, "NAME" in [font style], [colors], marble surface, shallow DOF` |
| App icon | `rounded square app icon, [symbol] in white on gradient #HEX to #HEX, iOS style` |
| Iterate / edit logo | `Keep exact logo design and palette. Change [one thing] only.` |

---

## 1. Logo Design

### Logo Types

There are seven fundamental logo categories. Each serves a different strategic purpose and requires a different prompting approach.

| Type | Description | Famous Examples | Best For |
|------|-------------|-----------------|----------|
| **Wordmark / Logotype** | Brand name rendered in custom typography | Google, Coca-Cola, FedEx, Disney | Distinctive names, strong typography |
| **Lettermark / Monogram** | Initials or abbreviation | IBM, HBO, CNN, HP, NASA | Long company names, global recognition |
| **Pictorial / Symbol** | Recognizable icon or image | Apple, Twitter/X, Nike Swoosh, Target | Universal recognition, app icons |
| **Abstract** | Geometric or non-literal shape | Pepsi, Adidas trefoil, Airbnb Belo | Conveying concept without literal image |
| **Mascot** | Character or illustrated figure | Michelin Man, KFC Colonel, Duolingo Owl | Approachable brands, family audiences |
| **Emblem** | Text enclosed inside a symbol or badge | Starbucks, Harley-Davidson, NFL, BMW | Heritage, authority, premium feel |
| **Combination** | Symbol paired with text | Burger King, Lacoste, Doritos, Adidas | Flexibility (icon separates from text) |

### Logo Design Principles

1. **Simplicity** -- A logo must work at 16px (favicon) and on a billboard. If the shape is not recognizable as a silhouette, it is too complex. Reduce elements until nothing more can be removed.
2. **Memorability** -- The best logos have a distinctive silhouette that can be drawn from memory. Test: can you sketch it in 5 seconds? If not, simplify.
3. **Versatility** -- Must function on light backgrounds, dark backgrounds, in full color, single color, and reversed out of a photo. Design for the worst case first.
4. **Timelessness** -- Avoid trend-dependent elements (gradients of the decade, shadow styles, skeuomorphism). The Nike Swoosh has not changed since 1971.
5. **Appropriateness** -- A law firm logo should not look like a children's toy brand. The visual language must match the industry, audience, and brand personality.
6. **Scalability** -- Fine details disappear at small sizes. Thin strokes break apart. Test your logo at 32px, 128px, and 1024px.

### Logo Style Keywords for AI Prompts

Group these keywords to steer AI models toward your desired aesthetic.

**Minimal / Modern:**
- minimal, geometric, flat, clean, Swiss style, Bauhaus, modernist, reductive, grid-based, whitespace, negative space, simple forms

**Premium / Luxury:**
- luxury, premium, elegant, sophisticated, gold foil, silver emboss, serif typography, refined, haute, upscale, bespoke, monogram

**Tech / Digital:**
- tech, digital, circuit, pixel, binary, holographic, neon, gradient mesh, futuristic, cyber, data-driven, node graph, wireframe

**Vintage / Heritage:**
- vintage, retro, hand-drawn, distressed, letterpress, woodcut, Art Deco, Art Nouveau, Victorian, mid-century, aged, patina, heritage badge

**Organic / Natural:**
- organic, flowing, natural, botanical, leaf, water, wave, biomorphic, earthy, hand-crafted, artisanal, calligraphic

**3D / Dimensional:**
- 3D, glossy, metallic, embossed, chrome, glass, isometric, extruded, volumetric, rendered, reflective surface

**Playful / Friendly:**
- playful, rounded, bubbly, colorful, whimsical, cartoon, illustrated, chunky, bouncy, vibrant, fun, approachable

### Logo Prompt Templates

#### Wordmark Logo

**Model:** Use `search_models` to find a text-to-image model with strong typography/text rendering.

```
Minimal wordmark logo for "BRAND NAME", custom sans-serif lettering, clean geometric letterforms, balanced kerning, flat vector style on pure white background, professional brand identity design, no icon, text only, high contrast black typography
```

```
Elegant serif wordmark logo for "BRAND NAME", refined thin-to-thick stroke contrast, luxury fashion brand aesthetic, gold #C5A258 lettering on charcoal #2D2D2D background, minimalist, timeless, premium feel
```

#### Lettermark / Monogram

**Model:** Use `search_models` to find a text-to-image model with strong text rendering for lettermarks.

```
Monogram logo design, letters "AB" intertwined in a geometric pattern, minimal line weight, single color navy blue #1B3A5C, clean vector style on white background, professional corporate identity, balanced proportions
```

```
Modern lettermark logo "XY", bold geometric sans-serif, negative space design, interlocking forms, flat minimal style, tech startup brand identity, electric blue #0066FF on white background
```

#### Pictorial / Symbol Logo

**Model:** Use `search_models` to find a high-quality image generation model for symbol/icon design.

```
Minimal abstract bird symbol logo, single continuous line, geometric simplification, flat vector style, single color black on white background, professional logo design, clean and modern, scalable icon
```

```
Pictorial logo of a mountain peak, geometric triangular form, negative space river flowing through center, minimal flat design, teal #008B8B and white, outdoor adventure brand identity
```

#### Abstract Logo

**Model:** Use `search_models` to find a design-focused image generation model.

```
Abstract geometric logo mark, three overlapping circles forming a unified shape, gradient from deep blue #1A237E to cyan #00BCD4, flat modern design, tech company brand identity, clean vector on white background
```

#### Mascot Logo

**Model:** Use `search_models` to find an image generation model strong at illustrated characters.

```
Friendly owl mascot logo for an education brand, simple geometric illustration style, large expressive eyes, wearing a graduation cap, flat vector design, limited palette of navy #1B2A4A and gold #F4B942 on white background
```

#### Emblem Logo

**Model:** Use `search_models` to find a text-to-image model with strong typography/text rendering.

```
Circular emblem logo for "BRAND NAME EST. 2024", badge style with text wrapping around a central icon of a compass rose, vintage heritage design, single color dark green #1B4332, detailed but clean line work, on white background
```

#### Combination Logo

**Model:** Use `search_models` to find a text-to-image model with strong text rendering (text + icon needed).

```
Combination logo with a minimal leaf icon to the left of the text "GREENLEAF" in clean sans-serif typography, flat vector style, forest green #2D6A4F text and icon, horizontal layout, professional brand identity on white background
```

### Logo Variations Workflow

Every brand needs multiple logo formats. Generate them systematically:

| Variation | Purpose | Prompt Modifier |
|-----------|---------|-----------------|
| **Primary (full color)** | Main usage | Base prompt |
| **Reversed** | Dark backgrounds | Add: "on dark navy #0A0F1E background, white and light colored elements" |
| **Monochrome** | Single color print | Add: "single color black, no gradients, pure monochrome" |
| **Icon only** | App icon, favicon, social | Add: "icon mark only, no text, square format, centered" |
| **Horizontal lockup** | Website headers, letterhead | Add: "horizontal layout, icon left text right" |
| **Vertical lockup** | Social profiles, signage | Add: "vertical stacked layout, icon above text" |

### Logo Generation Workflow

1. **Concept generation** -- Use `search_models` to find a text-to-image model with strong text rendering for initial concepts. Generate 4-8 variations. Models with high typography accuracy work best.
2. **Refinement** -- Use `search_models` to find an image editing model for iterating on specific elements while maintaining the core design. Use prompts like "change the color to deep red" or "make the typography bolder".
3. **Background removal** -- Use `search_models` to find a background removal model to get a transparent PNG.
4. **Vectorization** -- Use `search_models` to find a vectorization model. Convert to SVG for infinite scalability.
5. **Upscale for print** -- Use `search_models` to find an upscaling model if a high-res raster is needed for sharp edges.

---

## 2. Color Systems

### Color Psychology

Color is the fastest communicator in brand identity. Viewers form an opinion about a brand within 90 seconds, and 62-90% of that judgment is based on color alone.

| Color | Associations | Industries | HEX Reference |
|-------|-------------|------------|---------------|
| **Red** | Energy, passion, urgency, danger, appetite | Food, entertainment, retail, sports | `#E63946` warm, `#C1121F` deep |
| **Orange** | Friendly, confident, creative, youthful, affordable | Tech, food, creative, children | `#FF6B35` vibrant, `#E76F51` muted |
| **Yellow** | Optimistic, warm, attention, caution, happiness | Food, transportation, retail | `#FFB703` golden, `#FFC300` bright |
| **Green** | Nature, growth, health, wealth, stability | Finance, health, eco, agriculture | `#2D6A4F` forest, `#40916C` sage |
| **Blue** | Trust, professional, calm, technology, security | Tech, finance, healthcare, corporate | `#0066FF` electric, `#1B3A5C` navy |
| **Purple** | Luxury, creative, spiritual, wisdom, royalty | Beauty, luxury, education, creative | `#7B2D8E` royal, `#9B5DE5` soft |
| **Pink** | Playful, feminine, compassion, modern, bold | Fashion, beauty, food, lifestyle | `#FF006E` hot, `#FFB4C2` soft |
| **Black** | Premium, powerful, elegant, sophisticated | Luxury, fashion, tech, automotive | `#0D0D0D` soft, `#000000` pure |
| **White** | Clean, minimal, pure, medical, space | Healthcare, tech, minimal brands | `#FFFFFF` pure, `#F8F9FA` warm |
| **Gold** | Luxury, success, prestige, quality, wealth | Finance, luxury, awards, premium | `#D4A853` classic, `#C5A258` muted |

### Color Palette Types

**Monochromatic** -- One hue, multiple tones (tints, shades, tones). Creates a clean, cohesive, sophisticated look. Easiest to maintain consistency.
- Example: Navy `#1B3A5C`, Medium Blue `#3A6EA5`, Light Blue `#6BA3D6`, Pale Blue `#C0D6EB`

**Analogous** -- Adjacent hues on the color wheel. Harmonious and pleasing, low contrast. Good for calm, unified brands.
- Example: Teal `#008B8B` + Blue `#0077B6` + Indigo `#3A0CA3`

**Complementary** -- Opposite hues on the wheel. Maximum contrast and visual tension. Bold and energetic.
- Example: Blue `#0066FF` + Orange `#FF6600`

**Split-Complementary** -- One base + two colors adjacent to its complement. Strong contrast but more nuanced than pure complementary.
- Example: Blue `#0077B6` + Yellow-Orange `#FFB703` + Red-Orange `#E63946`

**Triadic** -- Three evenly spaced hues. Vibrant and balanced. Works well for playful or creative brands.
- Example: Red `#E63946` + Blue `#457B9D` + Yellow `#F1C40F`

**Neutral + Accent** -- Mostly gray, white, and black with one strong accent color. Modern, corporate, clean. The accent does all the heavy lifting.
- Example: Charcoal `#2D2D2D` + Light Gray `#E5E5E5` + White `#FFFFFF` + Accent Coral `#FF6B6B`

### Industry Color Conventions

These are conventions, not rules. Breaking them intentionally is a valid brand strategy, but you should understand the norms first.

| Industry | Primary Colors | Reasoning |
|----------|---------------|-----------|
| **Tech / SaaS** | Blue, purple, teal, gradients | Trust, innovation, digital-native |
| **Food / Restaurant** | Red, orange, yellow, green | Appetite stimulation, warmth, freshness |
| **Finance / Banking** | Navy, dark blue, green, gold | Trust, stability, wealth, authority |
| **Health / Medical** | Blue, green, white, soft teal | Calm, cleanliness, healing, trust |
| **Fashion / Luxury** | Black, gold, white, muted tones | Sophistication, premium, timelessness |
| **Eco / Sustainability** | Green, earth tones, sky blue | Nature, responsibility, organic |
| **Entertainment / Media** | Red, black, neon, vibrant | Energy, excitement, boldness |
| **Education** | Blue, navy, green, warm tones | Knowledge, growth, tradition |
| **Real Estate** | Blue, green, gold, charcoal | Trust, growth, premium, stability |
| **Children / Toys** | Primary colors, rainbow, pastels | Fun, energy, approachability |

### Using Color in AI Prompts

Always be explicit about color. Vague color descriptions produce inconsistent results.

**Specify HEX codes directly:**
```
primary color #FF6B35, secondary color #1B3A5C
```

**Describe the palette relationship:**
```
limited palette of navy blue #1B3A5C and warm gold #D4A853 with white #FFFFFF accents
```

**Reference well-known brand colors for anchor:**
```
Tiffany blue teal accent, clean white background
```

**Control palette size:**
```
strict three-color palette: charcoal #2D2D2D, coral #FF6B6B, and white #FAFAFA, no other colors
```

### Color Palette Visualization Prompt

**Model:** Use `search_models` to find an image generation model suited for graphic design assets.

```
Professional color palette presentation, five vertical color swatches side by side, colors from left to right: deep navy #1B3A5C, teal #008B8B, warm gold #D4A853, light gray #E8E8E8, off-white #FAFAFA, clean minimal design, each swatch labeled with its hex code, brand identity design asset
```

---

## 3. Typography

### Font Category Guide

Typography carries as much brand meaning as color. The typeface you choose signals whether a brand is serious or playful, traditional or modern, accessible or exclusive.

| Category | Personality | Famous Examples | Best For |
|----------|------------|-----------------|----------|
| **Serif** | Traditional, trustworthy, editorial, established | Times New Roman, Garamond, Playfair Display, Baskerville | Law, finance, editorial, luxury, heritage |
| **Sans-Serif** | Modern, clean, accessible, tech-forward | Helvetica, Inter, Futura, Montserrat, DM Sans | Tech, startups, healthcare, corporate |
| **Slab Serif** | Bold, industrial, confident, grounded | Rockwell, Courier, Museo Slab, Roboto Slab | Construction, manufacturing, bold statements |
| **Display** | Decorative, attention-grabbing, distinctive | Impact, Lobster, Bangers, Bebas Neue | Headlines, posters, event branding |
| **Script** | Elegant, personal, flowing, handcrafted | Pacifico, Great Vibes, Dancing Script, Allura | Wedding, beauty, boutique, hospitality |
| **Monospace** | Technical, precise, code-like, systematic | Fira Code, JetBrains Mono, IBM Plex Mono, Courier | Developer tools, data, technical brands |
| **Handwritten** | Casual, authentic, approachable, personal | Caveat, Indie Flower, Permanent Marker | Cafes, personal brands, crafts, journals |

### Typography Hierarchy

A brand system needs at least two levels, ideally three to four:

| Level | Role | Typical Treatment | Example Pairing |
|-------|------|-------------------|-----------------|
| **Display / Hero** | Main headlines, splash screens | Large, bold or ultra weight, distinctive | Playfair Display Bold |
| **Heading** | Section headers, card titles | Medium-large, semi-bold to bold | Inter Semi-Bold |
| **Body** | Paragraphs, descriptions | Regular weight, optimized for readability | Inter Regular |
| **Caption / Detail** | Labels, metadata, fine print | Small, light or regular weight, often uppercase | Inter Light or Mono |

### Typography Pairing Rules

1. **Contrast, not conflict** -- Pair fonts that differ clearly (serif heading + sans-serif body) rather than two fonts that are similar but not identical.
2. **Shared x-height** -- Fonts with similar x-heights sit well together on the same baseline grid.
3. **Two is safe, three is maximum** -- One display, one body. Add a third only for a specific role (code blocks, captions, accents).
4. **Weight over variety** -- One font family in multiple weights (Light, Regular, Bold, Black) often creates better consistency than multiple families.

### Classic Typography Pairings

| Heading | Body | Personality |
|---------|------|-------------|
| Playfair Display | Source Sans Pro | Editorial luxury |
| Montserrat | Open Sans | Modern corporate |
| Bebas Neue | Lato | Bold tech |
| DM Serif Display | DM Sans | Balanced contemporary |
| Oswald | Roboto | Strong industrial |
| Lora | Inter | Warm professional |
| Space Grotesk | Space Mono | Developer / technical |

### Typography in AI Prompts

**Critical rule: always put text in quotes.** AI models render text most accurately when the target string is quoted.

**Specify font style explicitly:**
```
bold condensed sans-serif typography, uppercase letterforms
```

```
elegant thin serif lettering with high stroke contrast
```

**Finding models with good text rendering:**

Use `search_models` to find models by category "text to image with text" for best typography accuracy. Models specialized in text rendering handle multi-word strings, stylized fonts, and complex layouts. For short text (1-3 words), most high-quality image generation models work well. For longer text or precise typography, look for models specifically noted for text rendering capability.

---

## 4. Brand Applications and Mockups

Once a logo, color palette, and typography direction are established, generate realistic mockups to visualize the brand in context. This is where brands come to life.

### Stationery

**Business Card:**
```
Photorealistic business card mockup on a marble surface, clean minimal design, "BRAND NAME" in bold sans-serif, contact details in light weight below, primary color deep navy #1B3A5C with gold #D4A853 accent line, white card stock, subtle shadow, professional photography style
```

**Letterhead:**
```
Professional letterhead mockup, A4 white paper on a dark wood desk, minimal logo "BRAND NAME" top left in navy #1B3A5C, thin gold #D4A853 line below header, body text placeholder in light gray, clean corporate design, top-down photography angle
```

**Envelope:**
```
Branded envelope mockup, C5 size, white stock with navy #1B3A5C logo "BRAND NAME" on the flap, minimal clean design, placed on a neutral gray surface, professional product photography
```

### Digital Applications

**Website Hero Section:**
```
Website hero section mockup displayed on a modern MacBook Pro, clean landing page design with bold headline typography, brand colors navy #1B3A5C and coral #FF6B6B, large hero image, CTA button, minimal navigation bar, professional UI/UX design, 3/4 angle view of laptop
```

**Social Media Profile:**
```
Instagram profile mockup grid, 9 posts in cohesive brand style, consistent color palette of teal #008B8B and warm neutrals, mix of photography and graphic design posts, profile picture showing the brand logo icon, professional social media branding
```

**App Icon:**

**Model:** Use `search_models` to find a design-focused image generation model (clean vector style).

```
Mobile app icon design, rounded square format, minimal abstract symbol in white on gradient background from blue #0066FF to purple #7B2D8E, clean flat design, iOS app icon style, single centered icon element, no text
```

**Favicon:**
```
Favicon design, 32x32 pixel grid, single letter "B" in bold geometric sans-serif, white on deep blue #1B3A5C square, pixel-perfect, minimal, high contrast, simple recognizable form
```

### Packaging

**Product Box:**
```
Premium product packaging box mockup, matte black box with gold #D4A853 foil stamped logo "BRAND NAME", minimal design, clean typography, luxury unboxing aesthetic, professional product photography on white surface, soft studio lighting
```

**Bottle Label:**
```
Minimalist bottle label design mockup on a glass bottle, white label with black typography "BRAND NAME" in serif font, simple line illustration of botanical elements, clean Scandinavian design aesthetic, professional product photography, neutral background
```

**Shopping Bag:**
```
Branded paper shopping bag mockup, kraft brown bag with navy #1B3A5C logo "BRAND NAME" printed large on the side, cotton handles, standing upright on white surface, clean commercial photography, minimal design
```

### Environmental

**Storefront Signage:**
```
Modern storefront signage mockup, illuminated channel letters spelling "BRAND NAME" mounted on dark gray building facade, warm LED glow at dusk, architectural photography style, urban setting, professional commercial signage
```

**Vehicle Wrap:**
```
Professional van wrap mockup, white delivery van with navy #1B3A5C and teal #008B8B brand graphics, logo "BRAND NAME" large on the side, contact info on rear, clean modern vehicle branding, photographed on city street
```

**Office Wall:**
```
Branded office reception wall, large 3D logo "BRAND NAME" in brushed metal mounted on white wall, clean modern interior, recessed lighting illuminating the logo, minimalist corporate environment, architectural interior photography
```

### Merchandise

**T-Shirt:**
```
White t-shirt mockup on a flat surface, minimal chest print of brand logo "BRAND NAME" in black, clean screen-print style, folded neatly, professional apparel photography, neutral gray background
```

**Mug:**
```
Ceramic coffee mug mockup, white mug with brand logo "BRAND NAME" in navy #1B3A5C, minimal clean design, placed on a wooden table with coffee inside, warm morning light, lifestyle product photography
```

**Tote Bag:**
```
Canvas tote bag mockup, natural cotton fabric, large centered logo "BRAND NAME" screen-printed in black, minimal design, flat lay on white surface, professional product photography
```

---

## 5. Brand Style Guide Components

### What a Complete Brand Guide Contains

1. **Brand story** -- Mission, vision, values, and brand personality in narrative form.
2. **Logo usage rules** -- Clear space requirements (minimum padding = half the logo height around all sides), minimum reproduction size (typically 24px digital, 12mm print), prohibited modifications.
3. **Color palette** -- Primary (1-2 colors), secondary (2-3 colors), accent (1 color), neutral (2-3 grays/whites/blacks). Each with HEX, RGB, CMYK, and Pantone values.
4. **Typography** -- Primary font (headings), secondary font (body), accent font (optional). Sizes, weights, line heights, and letter spacing for each usage context.
5. **Photography style** -- Mood, subject guidelines, lighting preference, color treatment, composition rules.
6. **Iconography** -- Icon style (line, filled, duotone), stroke weight, corner radius, grid system.
7. **Voice and tone** -- How the brand speaks (formal/casual, technical/simple, serious/playful).
8. **Do's and don'ts** -- Visual examples of correct and incorrect usage for logo, color, and typography.

### Generating Brand Guide Assets with AI

**Logo clear space diagram:**
```
Technical diagram showing logo clear space rules, brand logo "BN" centered with dotted lines indicating minimum padding on all sides labeled "X", clean vector diagram style, black and white, professional brand guidelines format, grid overlay
```

**Color palette documentation:**

**Model:** Use `search_models` to find a text-to-image model with accurate text/number rendering.

```
Brand color palette documentation page, five large color swatches in a row: "#1B3A5C" navy, "#008B8B" teal, "#D4A853" gold, "#E8E8E8" gray, "#FAFAFA" off-white, each swatch labeled with hex code below, clean minimal grid layout, professional brand guidelines design
```

**Typography specimen sheet:**

**Model:** Use `search_models` to find a text-to-image model with strong text rendering.

```
Typography specimen sheet for brand guidelines, heading font "Montserrat Bold" showing full alphabet Aa Bb Cc in large size, body font "Inter Regular" showing paragraph text sample below, font sizes and weights listed, clean white page, professional typographic layout
```

---

## 6. Visual Consistency Techniques

Maintaining visual consistency across multiple AI generations is the hardest challenge in AI-assisted brand work. Use these strategies:

### Anchor Phrases

Create a "style anchor" -- a block of descriptive text that you prepend to every prompt for the same brand. This ensures visual coherence.

```
[STYLE ANCHOR] Clean modern brand aesthetic, color palette: navy #1B3A5C, teal #008B8B, gold #D4A853, off-white #FAFAFA. Typography: bold geometric sans-serif headings, clean sans-serif body text. Style: minimal, professional, generous whitespace, flat design with subtle shadows.
```

Prepend this to every prompt for the brand, then add the specific asset request.

### Kontext Editing for Brand Iteration

Use `search_models` to find an image editing model that can take an existing brand asset and modify it while preserving the established look:

```
[Provide existing brand image as input]
Keep the exact same logo design, color palette, and typography style. Change the background to a dark navy #0A0F1E. Make the logo elements white. Add a subtle gradient glow behind the logo.
```

### Consistent Color Control

- Always specify exact HEX codes in every prompt, never rely on color names alone.
- Use "strict color palette" or "limited palette" language to prevent the model from introducing extra colors.
- Validate output colors match your specification. Regenerate if they drift.

### Reference Image Chaining

For multi-asset brand projects, use each generated asset as a reference for the next:

1. Generate logo with a text-to-image model that has strong text rendering (use `search_models`)
2. Use that logo as input to an image editing model for business card mockup (use `search_models`)
3. Use the business card as style reference for letterhead generation
4. Each step inherits visual DNA from the previous output

---

## 7. Industry-Specific Brand Examples

### Tech Startup

**Palette:** Electric Blue `#0066FF`, Dark `#0D1117`, White `#FFFFFF`, Light Gray `#F0F0F0`
**Typography:** Geometric sans-serif (Inter, DM Sans, Space Grotesk)
**Logo style:** Abstract or lettermark, clean geometry, gradients allowed
**Mood:** Innovative, trustworthy, forward-thinking

```
Minimal abstract logo for a SaaS startup "NEXUS", geometric overlapping shapes forming an N, gradient from electric blue #0066FF to purple #7B2D8E, clean flat design on dark #0D1117 background, modern tech brand identity
```

### Restaurant / Food Brand

**Palette:** Warm Red `#C1121F`, Cream `#FFF8E7`, Forest Green `#2D6A4F`, Charcoal `#333333`
**Typography:** Combination of serif display and clean sans-serif
**Logo style:** Combination or emblem, often with illustration element
**Mood:** Warm, appetizing, inviting, authentic

```
Restaurant logo for "TERRA Kitchen", combination mark with a minimal wheat stalk icon beside hand-lettered serif text, warm red #C1121F and forest green #2D6A4F on cream #FFF8E7 background, artisanal farm-to-table feel, organic natural aesthetic
```

### Fashion / Luxury Brand

**Palette:** Black `#000000`, White `#FFFFFF`, Gold `#C5A258`, Nude `#D4B8A0`
**Typography:** Elegant high-contrast serif or ultra-thin sans-serif, generous letter spacing
**Logo style:** Wordmark or monogram, typically uppercase
**Mood:** Exclusive, refined, confident, timeless

```
Luxury fashion wordmark logo "MAISON CLAIRE", ultra-thin elegant serif typography, wide letter spacing, all uppercase, gold #C5A258 on black #000000 background, high fashion brand identity, refined and timeless, Vogue editorial aesthetic
```

### Wellness / Health Brand

**Palette:** Sage Green `#A7C4A0`, Soft Blue `#B8D4E3`, Warm White `#FFF9F0`, Blush `#E8C4B8`
**Typography:** Rounded sans-serif or soft serif, approachable and calm
**Logo style:** Organic shapes, leaf/nature motifs, flowing lines
**Mood:** Calming, natural, nurturing, balanced

```
Wellness brand logo for "BLOOM Wellness", minimal organic leaf shape forming a B, soft sage green #A7C4A0, rounded sans-serif text, gentle and calming design, clean flat style on warm white #FFF9F0 background, holistic health brand identity
```

### Professional Services (Law, Consulting, Finance)

**Palette:** Navy `#1B2A4A`, Dark Gold `#B8860B`, Charcoal `#2D2D2D`, Light Gray `#E5E5E5`, White `#FFFFFF`
**Typography:** Classic serif for firm name, clean sans-serif for supporting text
**Logo style:** Wordmark, lettermark, or emblem. Conservative and authoritative.
**Mood:** Trustworthy, established, professional, authoritative

```
Professional law firm logo for "HARRISON & ASSOCIATES", classic serif typography, navy #1B2A4A text, thin gold #B8860B horizontal line underneath, clean minimal design, white background, corporate brand identity, trustworthy and established
```

---

## 8. Complete Brand Prompt Examples

### Example 1: Logo with Specific Typography

**Model:** Use `search_models` — text-to-image with strong text rendering.

```
Minimal modern logo for "ARCADIA Studios", clean geometric sans-serif typography, lowercase, letter "a" has a distinctive notch detail, flat design, charcoal #2D2D2D on white background, professional creative agency brand identity, no icon, pure wordmark
```

### Example 2: Abstract Symbol for Tech

**Model:** Use `search_models` — design-focused image generation.

```
Abstract technology logo symbol, three connected hexagonal nodes forming a network pattern, gradient from teal #00BFA5 to deep blue #1A237E, minimal flat vector design, no text, clean tech brand mark, centered on pure white background
```

### Example 3: Color Palette Visualization

**Model:** Use `search_models` — image generation.

```
Professional brand color palette presentation card, five large circular swatches arranged horizontally: deep forest green #1B4332, sage #52796F, cream #F2E8CF, terracotta #BC6C25, dark brown #3E2723, clean modern design on white background, design asset
```

### Example 4: Branded Social Media Post

**Model:** Use `search_models` — high-quality image generation.

```
Instagram post design for a coffee brand, square format, bold headline "MORNING RITUAL" in white condensed sans-serif, warm moody photograph of steaming coffee cup in background, dark brown #3E2723 overlay with 60% opacity, brand logo small bottom right, professional social media design
```

### Example 5: Business Card Mockup

**Model:** Use `search_models` — high-quality image generation.

```
Premium business card mockup, two cards on a dark slate surface, front card shows minimal design with "JAMES CHEN" in thin sans-serif and "Creative Director" below, back card shows large logo mark in coral #FF6B6B on charcoal #2D2D2D, thick 600gsm stock with edge painting in coral, professional product photography, shallow depth of field
```

### Example 6: Packaging Design

**Model:** Use `search_models` — high-quality image generation.

```
Luxury skincare packaging design, frosted glass bottle with pump dispenser, minimal white label with "AURA" in thin uppercase serif, small botanical line illustration, muted sage green #A7C4A0 accent details, clean Scandinavian aesthetic, product photography on marble surface, soft natural lighting
```

### Example 7: Brand Pattern / Texture

**Model:** Use `search_models` — image generation.

```
Seamless geometric brand pattern, repeating minimal hexagonal grid, thin line weight, navy #1B3A5C lines on white background, subtle and elegant, suitable for business stationery background, clean vector style, tileable pattern design
```

### Example 8: Emblem with Vintage Feel

**Model:** Use `search_models` — text-to-image with strong text rendering.

```
Vintage circular emblem logo for "NORTHBOUND Brewing Co. EST. 2020", hand-drawn style compass rose in center, text wraps around the circle, hops and wheat decorative elements, single color rust orange #C65D07, distressed letterpress texture, heritage craft brewery brand identity, on cream background
```

### Example 9: App Icon with Brand Colors

**Model:** Use `search_models` — design-focused image generation.

```
Mobile app icon, rounded square, abstract folded paper airplane symbol in white, background gradient from electric blue #0066FF to violet #7C3AED, minimal flat design, iOS style, clean single focal element, no text, brand icon for a messaging app
```

### Example 10: Brand Photography Style Guide

**Model:** Use `search_models` — high-quality image generation.

```
Brand photography style guide page showing four example photographs in a grid, cohesive warm golden hour lighting across all images, subjects: product flat lay, lifestyle portrait, workspace detail, outdoor texture, consistent warm color grading with lifted shadows, brand color accent coral #FF6B6B visible in each shot, professional editorial photography style
```

### Example 11: Iconography Set

**Model:** Use `search_models` — design-focused image generation.

```
Set of 6 minimal line icons in a grid layout: email envelope, phone, location pin, calendar, user profile, settings gear, consistent 2px stroke weight, rounded line caps, navy #1B3A5C on white background, clean brand icon set, unified design language
```

### Example 12: Brand Typography Poster

**Model:** Use `search_models` — text-to-image with strong text rendering.

```
Typography poster design, large display text "THE QUICK BROWN FOX" in bold serif Didot style at top, body text paragraph sample in clean sans-serif below, font size hierarchy demonstration, black text on white background, professional type specimen design, brand guidelines format
```

### Example 13: Storefront Visualization

**Model:** Use `search_models` — high-quality image generation.

```
Modern cafe storefront mockup at golden hour, large glass windows, illuminated interior sign reading "DOSE" in minimal sans-serif, warm interior lighting visible through glass, outdoor seating with branded umbrellas in forest green #2D6A4F, architectural photography, clean modern design
```

### Example 14: Branded Presentation Slide

**Model:** Use `search_models` — image generation.

```
Corporate presentation slide design, dark navy #0A0F1E background, large white bold headline "Our Vision" top left, three columns below with icon and text placeholder in each, accent color teal #00BFA5 used for icon circles and underlines, minimal clean layout, professional pitch deck design, 16:9 aspect ratio
```

### Example 15: Complete Brand Board

**Model:** Use `search_models` — high-quality image generation.

```
Professional brand board layout on white background, organized grid showing: logo in top left, color palette swatches (navy, gold, white, gray) top right, typography specimens center left, brand photography mood images center right, pattern sample bottom left, business card mockup bottom right, cohesive brand identity overview, clean design presentation
```

### Example 16: Logo Animation Starting Frame

**Model:** Use `search_models` — image generation.

```
Logo reveal animation keyframe, minimal "N" lettermark in white centered on deep black #0D0D0D background, subtle particle dispersion effect around the letter as if materializing, faint blue #0066FF glow, cinematic motion graphics still, dark atmospheric brand reveal moment
```

