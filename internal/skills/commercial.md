# Commercial — Advertising & Product Content Production Guide

Comprehensive guide for AI-generated commercial content: product photography, advertising video, stock footage, social media formats, and motion graphics. Optimized for fal.ai models (2026).

## Quick Reference

### Top 5 Commercial Production Tips

1. **Image-to-Video always beats text-to-video for products.** Generate a perfect still first, then animate with I2V. This locks product appearance and prevents morphing.
2. **Leave negative space for ad copy** — don't ask AI to render text. Specify `negative space on the [left/right/top] for headline placement`.
3. **Subject always comes first in the prompt.** Open with a hyper-specific product description (material, color, form factor).
4. **Lighting vocabulary is the single biggest quality lever.** Replace "good lighting" with `single large softbox from upper-left` or `rim lighting separating subject from background`.
5. **Lock style consistency across campaigns** by copying the identical technical block into every shot prompt. Vary only subject action and camera angle.

### Key Prompt Patterns

**Hero shot:** `[Product + material + color], on [surface], [lighting setup], [lens + aperture], [negative space for copy], commercial product photography, 8K`

**Luxury / dark:** `[Product], single spotlight from 45 degrees, deep shadows, rim light, polished black acrylic surface, low-key, 8K`

**E-commerce packshot:** `[Product], light tent, even 360-degree diffused lighting, pure white seamless background, packshot, 8K`

**Food / beverage:** `[Food], [action mid-motion], warm directional light from upper-left, shallow DOF 85mm f/2.0, appetite appeal, food editorial, 8K`

**Video reveal:** `Darkness. [Single light source] gradually reveals [product]. Camera [one movement]. [Surface]. [Style modifier]. 4K.`

### Cheat Sheet — Surface / Lighting / Lens by Scenario

| Scenario | Surface | Lighting | Lens | Style Tag |
|----------|---------|----------|------|-----------|
| Premium any product | Black seamless / obsidian | Low-key single spotlight + rim | 85–100mm | `luxury commercial photography` |
| E-commerce / packshot | White seamless | Light tent / high-key | 50mm | `packshot, e-commerce product photography` |
| Food / beverage | Wood / slate / marble | Warm sidelight + steam | 85mm f/2.0 | `food editorial, appetite appeal` |
| Beauty / cosmetics | Marble / terrazzo | Clamshell / large softbox | 85mm f/1.8 | `beauty campaign, skincare editorial` |
| Tech / gadgets | Dark gradient / white | Rim light + precision studio | 85mm f/2.8 | `Apple product page aesthetic` |
| Fashion editorial | Location / cyclorama | Hard spotlight / dramatic | 85mm medium format | `Vogue editorial, high fashion` |

---

## Product Photography Techniques

### 1. Hero Shot (Single Product, Dramatic)

The hero shot is the cornerstone of product advertising. One product, one focal point, maximum impact.

**Key principles:**
- Single dominant light source with optional fill (2:1 to 4:1 ratio)
- Product occupies 40-60% of frame
- Rule of thirds or dead center (centered = authority, off-center = editorial)
- Shallow depth of field (f/1.4 to f/2.8) to isolate from background
- Surface: reflective (glass, acrylic) for luxury; matte (stone, wood) for organic

**Critical prompt descriptors:** `hero shot`, `single overhead softbox`, `studio product photography`, `commercial grade`, `shallow DOF`, `negative space`, `advertisement quality`

**Image prompts:**
```
Hero shot of premium wireless headphones on polished obsidian slab, single large
softbox overhead creating broad specular highlight on ear cups, subtle warm fill
from camera-left, shallow DOF f/2.0 isolating product, dark charcoal gradient
background fading to black, commercial product photography, 8K resolution
```

```
Hero shot of artisanal perfume bottle on raw marble block, dramatic Rembrandt
lighting from upper right, deep shadows with golden rim light on glass edges,
liquid amber visible through crystal-cut glass, negative space left side for
ad copy, luxury brand photography, medium format, 8K
```

```
Minimalist hero shot of smartwatch displaying fitness screen, floating above
matte white surface with soft shadow beneath, even diffused lighting from
360-degree light tent, pure white background, Apple product page aesthetic,
ultra-clean, packshot, 8K
```

**Video prompts:**
```
Slow cinematic rotation of premium headphones on obsidian pedestal. Single
overhead softbox creating moving specular highlights across brushed aluminum.
360-degree turntable rotation over 8 seconds. Black gradient background.
Smooth, deliberate, commercial product reveal. 4K.
```

```
Camera slowly pushes in on perfume bottle from medium shot to extreme close-up.
Spotlight sweeps from left to right revealing prismatic refractions through
crystal glass. Marble surface. Dark background. Luxury fragrance commercial. 4K.
```

```
Top-down camera slowly descending toward smartwatch on white surface. Watch
screen illuminates as camera approaches. Soft ambient lighting. Clean minimal
aesthetic. Tech product reveal. 4K.
```

### 2. Tabletop / Flat Lay

Overhead perspective showing curated product arrangements. Essential for beauty, cosmetics, stationery, food ingredients, and lifestyle brands.

**Key principles:**
- Camera perfectly overhead (90 degrees) or slight angle (80-85 degrees)
- Even, diffused lighting from both sides to eliminate harsh shadows
- Grid or organic arrangement (grids = professional, organic = lifestyle)
- Props must support, never compete: dried flowers, fabric swatches, tools of the trade
- Negative space for text overlay areas
- Color palette coordination: 2-3 colors maximum

**Critical prompt descriptors:** `flat lay`, `overhead camera`, `birds-eye view`, `even diffused lighting`, `curated arrangement`, `editorial styling`, `knolling` (for grid arrangements)

**Image prompts:**
```
Professional flat lay of luxury skincare collection on blush pink terrazzo
surface, bottles and jars arranged in golden ratio spiral, dried eucalyptus
sprigs and cotton gauze as props, perfectly even lighting with no shadows,
overhead camera, editorial beauty photography, Glossier aesthetic, 8K
```

```
Knolling flat lay of mechanical watch components precisely arranged on navy
blue suede, each part equidistant, tools of the watchmaker surrounding
the disassembled timepiece, warm studio lighting, overhead view, technical
craftsmanship, commercial photography, 8K
```

```
Breakfast flat lay on white marble surface, artisanal sourdough toast with
avocado, poached eggs with runny yolk, fresh berries in ceramic bowl, linen
napkin, coffee in handmade mug, morning sunlight from window, food editorial,
overhead shot, 4K
```

**Video prompts:**
```
Top-down camera slowly panning across curated flat lay of skincare products
on terrazzo surface. Hands enter frame and gently pick up serum bottle, tilt
to show golden liquid inside, place it back. Soft even lighting. Beauty
editorial style. 4K.
```

```
Stop-motion style flat lay assembly. Items appear one by one on white surface
in perfect grid arrangement. Watch, wallet, sunglasses, keys, phone. Each
placed with precision. Overhead camera. Clean minimal. 4K.
```

### 3. Splash / Liquid Photography

High-energy product shots featuring liquid dynamics. Essential for beverages, skincare, cleaning products, and any brand wanting to convey freshness.

**Key principles:**
- Freeze motion at 1/4000s to 1/8000s shutter speed equivalent
- Crown splash (drop impact), impact splash (object plunging), swirl, pour
- Rim lighting or backlight to illuminate liquid edges and droplets
- Dark backgrounds make liquid pop; light backgrounds feel fresh and clean
- Condensation on product surface adds realism
- Liquid color should complement or contrast the product

**Critical prompt descriptors:** `crown splash`, `suspended droplets`, `high-speed photography`, `water tension`, `hyper-realistic liquid refraction`, `condensation beads`, `1/8000s freeze motion`, `liquid dynamics`

**Image prompts:**
```
Sleek aluminum energy drink can plunging into crystal-clear water, violent
symmetrical crown splash rising around the can, thousands of individual
droplets suspended mid-air catching rim light, condensation beads on can
surface, dark slate background, two-point rim lighting from behind, 100mm
macro, 1/8000s, commercial beverage photography, 8K
```

```
Fresh strawberry falling into milk, creamy white splash erupting upward in
organic tendrils, individual milk droplets frozen in air, strawberry partially
submerged with visible liquid displacement, clean white background, high-key
lighting, food photography, 1/5000s, 8K
```

```
Premium face serum bottle with golden liquid pouring in continuous stream onto
marble surface, liquid pooling and spreading with visible viscosity, backlit
to show transparency and golden hue, droplets scattered on marble, black
background, luxury skincare commercial, macro, 8K
```

```
Ice cubes dropping into glass of whiskey, amber liquid erupting with dynamic
splash, ice cubes captured mid-bounce, condensation on glass exterior, single
warm spotlight from above creating caramel light through liquid, dark oak bar
surface, beverage commercial, 8K
```

**Video prompts:**
```
Extreme slow motion of orange juice pouring into glass from above. Liquid
impacts surface and erupts in dynamic splash, droplets rising and falling.
Camera at glass level. Warm backlight illuminating orange liquid. White
background. Fresh, vibrant. 240fps slow motion aesthetic. 4K.
```

```
Slow-motion close-up of water bottle being opened, cap twisting off with mist
escaping. Water pours over the bottle exterior creating rivulets. Individual
water beads rolling down the surface. Dramatic rim lighting. Dark background.
Refreshment commercial. 4K.
```

### 4. Food Photography

Food photography demands appetite appeal above all. Every element must make the viewer hungry.

**Key principles:**
- "Hero food" should be at peak doneness, texture, and color
- Single directional light (typically from 10-11 o'clock position) with bounce fill
- Steam, drizzle, condensation, garnish movement = life and freshness
- 45-degree angle (most flattering), overhead for flat dishes, eye-level for tall items
- Complementary props: wooden boards, linen, vintage utensils, fresh herbs
- Color temperature: warm (2800-3500K) for comfort food, neutral for fresh/healthy
- Sauce/drizzle always mid-pour or just landing for dynamic energy

**Critical prompt descriptors:** `food photography`, `appetite appeal`, `hero food`, `steam wisps`, `drizzle mid-pour`, `fresh garnish`, `moody lighting`, `shallow DOF`, `food editorial`

**Image prompts:**
```
Steaming bowl of tonkotsu ramen at 45-degree angle, rich milky pork broth with
rendered fat globules, chashu pork slices with caramelized edges, soft-boiled
egg with jammy orange yolk, fresh green onions and nori, wisps of steam
catching warm light from upper-left, dark wooden table, moody restaurant
atmosphere, 85mm f/2.0 shallow DOF, food photography, 8K
```

```
Drizzle of dark chocolate sauce mid-pour over stack of golden Belgian waffles,
sauce creating glossy ribbons and pooling at base, fresh raspberries and
powdered sugar dusting, warm morning light from side window, rustic wooden
board, whipped cream with visible texture, comfort food editorial, 4K
```

```
Overhead shot of Neapolitan pizza fresh from wood-fired oven, leopard-spotted
charred crust, San Marzano tomato sauce, buffalo mozzarella with pools of
melted cheese and basil leaves, steam rising, rustic wooden pizza peel, flour
dusting on dark surface, Italian food photography, 8K
```

```
Eye-level shot of towering gourmet burger on brioche bun, layers visible in
cross-section: smashed patty with melted American cheese dripping, crisp
lettuce, thick tomato slice, caramelized onions, special sauce dripping down
side, toasted sesame bun, dark background with dramatic side lighting,
food commercial, shallow DOF, 8K
```

**Video prompts:**
```
Close-up of honey being drizzled from wooden dipper onto stack of golden
pancakes. Honey stretches in glistening thread, pools on top pancake and
begins to cascade over edges. Steam rising from fresh pancakes. Warm morning
light from left. Breakfast food commercial. Slow motion. 4K.
```

```
45-degree angle of sizzling steak on cast iron skillet. Fat rendering and
bubbling, wisps of smoke rising. Butter melting alongside with fresh thyme
and garlic cloves. Camera slowly pushes in. Warm dramatic lighting.
Restaurant quality. Food commercial. 4K.
```

```
Overhead shot of hands assembling poke bowl. Rice placed first, then salmon
cubes, avocado slices fanned out, edamame, pickled ginger, sesame seeds
sprinkled from above. Each ingredient placed with care. Bright natural
lighting. Fresh, healthy food commercial. 4K.
```

### 5. Tech Products (Apple-Style Minimal)

The defining aesthetic of modern technology advertising. Absolute purity — the product IS the design.

**Key principles:**
- Seamless gradient or pure solid background (white, black, or brand color)
- Product floating or on invisible surface — no visible support
- Precise, controlled multi-source lighting that reveals material properties
- Zero lifestyle elements — no hands, no context, no props
- If the screen is shown, it must display content that demonstrates the product
- Material differentiation: aluminum reads different from glass reads different from ceramic
- Edge lighting and rim glow to separate product from background

**Critical prompt descriptors:** `Apple product photography`, `seamless background`, `floating product`, `precision lighting`, `ultra-clean`, `minimal`, `packshot`, `material rendering`, `studio lighting`

**Image prompts:**
```
Latest-generation smartphone floating against pure matte black gradient,
soft cool-white accent light from left edge creating thin rim glow on aluminum
frame, screen displaying vibrant abstract wallpaper, secondary warm fill from
right revealing camera module detail, no surface, no shadow, Apple commercial
aesthetic, ultra-clean, product hero, 8K
```

```
Premium laptop open at 110 degrees on seamless white background, screen
displaying color-accurate photo, keyboard backlit with subtle glow, precise
studio lighting from above and slightly behind revealing thin profile and
aluminum unibody, product floating with soft shadow, tech product photography,
Apple website aesthetic, 8K
```

```
Pair of wireless earbuds in open charging case against deep navy gradient,
case interior illuminated with subtle LED glow, earbuds reflecting controlled
studio lighting, glossy and matte material contrast, product floating, dark
premium tech aesthetic, 8K
```

**Video prompts:**
```
Darkness. A single point of light expands to reveal the edge of a premium
smartphone. Camera slowly orbits 180 degrees as lighting gradually builds,
revealing brushed aluminum frame, ceramic back, camera lenses. Pure black
void. Precision lighting. Apple keynote reveal aesthetic. 4K.
```

```
Laptop slowly opening from closed position. Screen illuminates the keyboard
and surrounding darkness. Camera at desk level, slight low angle. Screen
content comes alive with color. Smooth, deliberate, 6-second reveal. Tech
product commercial. Minimal. 4K.
```

### 6. Jewelry & Luxury Goods

Jewelry and luxury items demand mastery of light interaction with precious materials. Sparkle, fire, luster, and material authenticity are paramount.

**Key principles:**
- Single controlled light source with precise positioning to create "fire" in gemstones
- Black negative fill to deepen shadows and increase contrast
- Macro or close-macro for detail (100mm+ equivalent)
- Surface material matters enormously: velvet, silk, leather, stone
- Specular highlights must be sharp and positioned on edges/facets
- Multiple internal reflections visible in transparent gems
- Gold reads warm (3200K), silver/platinum reads cool (5600K), diamonds read neutral

**Specular highlights vocabulary:** `anisotropic highlights` (stretched along grain — brushed metal), `specular roll-off` (how gradually a highlight fades), `caustics` (light patterns cast through/by transparent objects), `prismatic glints` (rainbow flashes from gem dispersion), `Fresnel effect` (increased reflection at glancing angles), `fire` (spectral dispersion in diamonds), `luster` (soft glow of pearls), `brilliance` (white light reflection from gem surface)

**Material rendering tips:**
- Gold: `warm specular highlights, rich yellow reflections, polished mirror finish` or `brushed gold with anisotropic highlights along grain`
- Silver/Platinum: `cool white specular highlights, mirror-like reflections of environment, blue-shifted shadows`
- Diamonds: `internal fire, spectral dispersion visible in facets, white brilliance on table facet, rainbow prismatic glints from pavilion`
- Pearls: `soft diffuse luster, orient (surface iridescence), nacre visible in close-up, cream undertones`
- Leather: `rich patina, pebbled grain texture visible, subtle sheen on high points, warm brown tones`

**Image prompts:**
```
Brilliant-cut diamond solitaire ring on folded black silk, single focused
spotlight from 45 degrees above-right creating sharp specular highlight on
platinum band, internal fire and prismatic dispersion visible through diamond
facets, caustic light patterns cast onto silk surface, deep bokeh background
with warm golden circles, macro photography 100mm f/2.8, luxury jewelry
commercial, 8K
```

```
Gold chain necklace draped over raw rose quartz crystal, warm directional
lighting revealing polished mirror finish with anisotropic highlights along
chain links, soft pink fill from quartz, Fresnel effect visible at glancing
angles on each link, shallow DOF, luxury editorial, 8K
```

```
Luxury Swiss chronograph on weathered leather surface, dial catching focused
light revealing applied indices and fine guilloché pattern, sapphire crystal
with anti-reflective coating showing slight blue tint at edges, steel case
with alternating brushed and polished surfaces, horological photography,
extreme detail, 8K
```

```
Pearl strand on black velvet display bust, soft diffused overhead lighting
creating gentle luster on each pearl, orient (surface iridescence) visible
in macro detail, graduated sizing from 7mm to 11mm, nacre quality visible,
dark dramatic background, fine jewelry photography, 8K
```

**Video prompts:**
```
Extreme slow rotation of diamond engagement ring on black velvet. Spotlight
creates prismatic fire flashing from diamond facets as it turns. Caustic
patterns dance on velvet surface. Camera slowly pushes in from medium to
macro. Single light source. Luxury jewelry commercial. 4K.
```

```
Close-up of hands fastening luxury watch clasp on wrist. Camera at wrist
level. Watch face catches light and reflects. Brushed steel contrasts with
polished bezel. Slow, deliberate movement. Warm natural lighting. Premium
timepiece commercial. 4K.
```

### 7. Automotive Photography

Automotive photography combines architectural scale with precise surface rendering. Cars are large reflective objects that mirror their entire environment.

**Key principles:**
- Studio: light painting technique (long exposure + moving light) for smooth, even reflections
- Location: golden hour, blue hour, or wet surfaces at night
- Camera height: belt-line level (1m) for authority, ground-level for aggression, elevated for context
- Three-quarter front is the classic "hero angle" showing front and side
- Wet surfaces multiply visual interest through reflections
- Environment is critical — the car reflects EVERYTHING around it
- Body lines and design language should be emphasized through light placement

**Critical prompt descriptors:** `automotive photography`, `light painting`, `rig shot`, `three-quarter view`, `wet surface reflections`, `rim lighting`, `dynamic lines`, `car commercial`

**Image prompts:**
```
Cinematic three-quarter front view of matte black luxury SUV in vast darkened
studio, professional light painting technique revealing body lines and contours,
continuous specular highlight along belt-line, subtle blue accent on lower body
panels, polished concrete floor with soft reflection, automotive commercial
photography, 8K
```

```
Low-angle dramatic shot of sports car on rain-soaked city street at night,
neon signs reflecting in wet asphalt creating color ribbons, car headlights
cutting through light mist, rim lighting from streetlamps defining rear
fender curves, long-exposure light trails from passing traffic, cinematic
automotive, anamorphic 2.39:1, 8K
```

```
Aerial three-quarter view of electric vehicle on salt flat at golden hour,
vast empty landscape stretching to horizon, warm directional sunlight
creating long shadow, vehicle paint reflecting golden sky, minimalist
composition, environmental automotive photography, 8K
```

**Video prompts:**
```
Tracking shot alongside sports car driving on mountain road at golden hour.
Camera mounted at wheel level on rig, car sharp with motion-blurred background.
Sunlight flashing through trees creating strobe effect on car body. Dust
particles in air. Cinematic automotive commercial. 4K.
```

```
Slow cinematic orbit around parked luxury sedan in dark studio. Light painting
effect — continuous moving highlight sweeps across body panels revealing curves
and design lines. Reflections in polished paint. Dark dramatic atmosphere.
Premium car commercial. 4K.
```

```
Front view of electric car driving straight toward camera on desert highway.
Heat shimmer on asphalt. Car grows from distant dot to filling the frame.
Camera low, centered. Wide open landscape. Blue sky. Dramatic approach.
Automotive freedom commercial. 4K.
```

### 8. Fashion Photography

Fashion divides cleanly into two modes: editorial (art-driven, magazine) and commercial (product-driven, e-commerce). Each has distinct rules.

**Key principles — Editorial:**
- Story and mood over product clarity
- Dramatic, unconventional lighting (hard shadows, color gels, mixed sources)
- Strong poses, exaggerated movement, fabric in motion
- Location or set design as narrative element
- Post-processing: color grading, contrast, desaturation for mood

**Key principles — Commercial:**
- Garment clarity is paramount — every detail must be visible and accurate
- Even, flattering lighting (large softboxes or diffused window light)
- Model poses that show fit, drape, and construction
- Neutral or complementary backgrounds that do not compete
- Color accuracy is critical for e-commerce (customer expectation management)

**Critical prompt descriptors:**
- Editorial: `high fashion editorial`, `Vogue`, `avant-garde`, `dramatic lighting`, `fashion story`, `medium format`
- Commercial: `e-commerce`, `catalog photography`, `clean studio`, `neutral background`, `garment detail`, `fit model`

**Image prompts:**
```
High-fashion editorial, model in structured oversized black coat striding
through fog-filled industrial warehouse, dramatic single hard spotlight from
above creating sharp shadow, strong directional wind blowing coat open to
reveal red lining, desaturated cool palette with warm highlights on face,
medium format camera, Vogue Italia editorial, 8K
```

```
Editorial beauty shot, close-up of model with geometric metallic face paint,
high-contrast studio lighting with deep shadows, wet-look slicked-back hair,
strong cheekbones catching light, editorial makeup, bold unconventional,
beauty magazine cover, 8K
```

```
Clean commercial fashion photo, female model wearing cream cashmere sweater
and tailored navy trousers, standing in bright airy studio with natural
window light from camera-left, light gray seamless background, relaxed
confident pose with one hand in pocket, garment details clearly visible,
wrinkle-free, e-commerce product photography, 4K
```

```
E-commerce flat lay of men's premium sneaker, three-quarter angle showing
sole profile and upper detail, pure white background, even shadow-free
lighting, laces styled naturally, material textures visible (suede, mesh,
rubber), product packshot, shoe commercial, 8K
```

**Video prompts:**
```
Fashion editorial video, model walking toward camera in slow motion wearing
flowing silk gown, fabric billowing dramatically behind, industrial concrete
space, strong backlight creating silhouette and fabric translucency, wind
machine effect, editorial fashion film, cinematic. 4K.
```

```
E-commerce product video, 360-degree rotation of sneaker on white turntable.
Even diffused lighting. Camera at shoe level. Smooth continuous rotation
showing all angles. Clean white background. Product detail video. 4K.
```

### 9. Cosmetics & Beauty

A sub-specialty that bridges product photography and fashion, demanding both product accuracy and aspirational beauty.

**Key principles:**
- Texture is everything: cream must look creamy, matte must look matte, glossy must gleam
- Swatches, smears, and texture shots show product consistency
- Skin must look real — pores visible but flawless, natural luminosity
- Lighting: broad, soft, and slightly directional for skin; harder and more focused for product
- Color accuracy critical (lipstick shade, foundation match)

**Image prompts:**
```
Luxury lipstick bullet extended from gold case, extreme macro showing creamy
texture and precise molded tip, single warm spotlight creating hot specular
on metal case, lipstick color rich burgundy with visible moisture, dark
background with subtle gradient, beauty product photography, 8K
```

```
Textural smear of four foundation shades on frosted glass surface, ranging
light to deep, creamy consistency visible with natural spreads and finger
marks, soft even overhead lighting, minimalist beauty editorial, clean
composition, 8K
```

```
Close-up beauty shot of model's face with dewy skin and natural makeup, single
soft catchlight in each eye, visible pore texture with luminous healthy glow,
subtle highlight on cheekbone and nose bridge, soft focus on hair, beauty
campaign, skincare brand, 85mm f/1.8, 8K
```

### 10. Beverage Photography

Beverage photography demands mastery of condensation, translucency, carbonation, and pour dynamics. The glass or bottle is a lens — light passing through the liquid is the primary visual story.

**Key principles:**
- Backlight or strong side-light to illuminate liquid color and clarity
- Condensation must look natural — varying droplet sizes, gravity-pulled drip trails
- Carbonation trails should be visible rising through the liquid
- Surface: wet slate, bar wood, or reflective black for premium spirits
- Ice must look crystalline, not opaque white
- Pour shots require implied motion: cascading liquid, dynamic foam, rising bubbles

**Critical prompt descriptors:** `condensation droplets`, `dew beads`, `frost texture`, `effervescence`, `carbonation trails`, `meniscus`, `laminar pour`, `cascade pour`, `head retention`, `lacing`, `backlit translucency`, `color gradient through liquid`

**Image prompts:**
```
Ice-cold craft IPA being poured into branded tulip glass, golden amber liquid
cascading with dense white head forming and carbonation trails rising, condensation
droplets running down glass exterior, dark wooden bar surface with single warm
backlight illuminating beer color, craft brewery commercial, 85mm f/2.8, 4K
```

```
Tall glass of freshly squeezed orange juice on white marble surface, morning
sunlight streaming through the liquid from behind showing suspended pulp, bright
vibrant orange, condensation running down glass, halved orange and scattered ice
cubes beside, bright kitchen background, breakfast beverage commercial, 90mm f/3.5, 8K
```

```
Espresso shot being pulled from bottomless portafilter into clear glass,
rich tiger-stripe crema forming, honey-thick stream of dark brown coffee,
steam rising, close-up macro shot, cafe environment blurred behind, warm
ambient lighting, specialty coffee commercial, 100mm macro f/2.8, 4K
```

```
Premium gin and tonic in copa glass with abundant ice, lime wheel, and rosemary
sprig, tonic bubbles rising in visible effervescent streams, blue juniper berries
resting on ice, backlit to reveal pale botanical tint of the gin, wet dark slate
surface, moody bar atmosphere, cocktail photography, 85mm f/2.0, 8K
```

**Video prompts:**
```
Slow-motion overhead pour of golden whiskey into crystal tumbler with large
spherical ice ball. Liquid cascading over ice, amber color glowing in warm
backlight. Camera at slight angle above glass. Whiskey settles, ice ball rotates
slightly. Dark oak bar surface. Spirits commercial. 120fps slow-motion. 4K.
```

### 11. Real Estate / Architecture Photography

Real estate and architectural photography serves a dual purpose: documenting space accurately and making it aspirational. AI generation excels at creating idealized interior and exterior visualizations.

**Key principles:**
- Wide-angle lens (16-24mm) with corrected verticals (no converging lines)
- Interior: HDR window pull — expose for both interior warmth and exterior view
- Twilight exterior: the "magic hour" for real estate — warm interior glow against deep blue sky
- Staging: furnished rooms sell better than empty ones
- Multiple heights: waist-level for interiors (eye level feels too high), drone for establishing
- Color temperature: warm interiors (2800-3500K feel), neutral-cool for modern/minimal

**Critical prompt descriptors:** `wide-angle interior`, `corrected verticals`, `window pull HDR`, `twilight exterior`, `staged interior`, `architectural detail`, `drone aerial`, `natural light flood`

**Interior:**
```
Professional real estate interior of modern open-plan penthouse living room,
floor-to-ceiling windows showing city skyline at blue hour, warm interior
lighting from recessed ceiling and floor lamps, contemporary furniture in
neutral palette, white oak hardwood floors reflecting warm light, shot at
16mm with corrected verticals, HDR window pull showing both interior and
exterior detail, architectural photography, 4K
```

```
Luxury bathroom interior, freestanding soaking tub centered below large window
with garden view, Calacatta marble walls and floor, brass fixtures catching warm
overhead light, folded white towels and single orchid, steam on mirror,
spa-like atmosphere, interior design photography, 20mm wide-angle, 4K
```

**Exterior:**
```
Twilight exterior of contemporary hillside home, warm interior light glowing
through large glass walls against deep blue sky with last traces of sunset
gradient, infinity pool in foreground reflecting the house and sky, landscape
lighting illuminating native plants and stone pathway, luxury real estate
photography, 24mm, 4K
```

**Aerial:**
```
Aerial drone photograph of Mediterranean villa estate, terracotta roof tiles,
central courtyard with fountain, surrounded by cypress trees and olive groves,
late afternoon golden light casting long shadows, turquoise swimming pool,
winding gravel driveway, real estate aerial photography, 4K
```

**Video prompts:**
```
Smooth cinematic walkthrough of luxury modern apartment. Camera at waist height
gliding forward on gimbal through entryway into open-plan living space.
Floor-to-ceiling windows reveal cityscape at golden hour, warm light flooding
the space. Camera arcs gently right through the kitchen showing marble waterfall
island, then continues toward primary bedroom. Steady, unhurried movement.
Wide-angle. Interior design showcase. 4K.
```

```
Drone descending from 100m altitude toward luxury coastal property at sunset.
Camera tilts down revealing estate layout — main house, guest house, pool,
gardens, private beach access. Golden hour light. Mediterranean architecture.
Smooth descent ending at eye level at the front entrance. Real estate
establishing shot. 4K.
```

---

## Commercial Video Styles

### 1. Apple-Style Product Reveal
**Characteristics:** Darkness-to-light, single product focus, slow deliberate movement, precision lighting, silence or minimal ambient sound, extreme restraint, every frame is a photograph.

```
Complete darkness. A hairline of light appears on a curved edge. Camera slowly
pulls back as overhead spotlight gradually increases, revealing premium
titanium laptop. Slow 180-degree orbit. Each surface catches light sequentially
— brushed top, polished hinge, keyboard illumination. Pure black void.
Apple keynote aesthetic. Ultra-smooth. 4K.
```

```
Single wireless earbud resting in darkness. Point light source slowly expands,
revealing translucent case around it. Camera rises as case opens. Second
earbud lifts out magnetically. Floating hover. Precision lighting on glossy
and matte surfaces. Tech product reveal. 4K.
```

### 2. Nike / Sports Brand Dynamic
**Characteristics:** Extreme slow motion, athlete in peak action, sweat/water/dust particles, dramatic rim lighting, ground-level or unusual angles, high energy through frozen moments.

```
Ultra slow motion of basketball player's shoes during crossover dribble,
rubber sole gripping polished court with visible flex, sweat droplets
launching from ankle, dramatic low-angle rim lighting, stadium lights
creating starburst bokeh in background, sports commercial, 4K
```

```
Slow motion of sprinter exploding from starting blocks, track surface
particles erupting from cleat impact, muscle definition visible with sweat
glistening, extreme low angle from track level, dark stadium with single
dramatic overhead light, athletic brand commercial, 4K
```

### 3. Perfume / Fragrance Luxury
**Characteristics:** Sensual, slow, ethereal. Fog, particles, reflections, flowing fabric or liquid. Camera movements are smooth and graceful. Color palette is controlled and rich.

```
Crystal perfume bottle centered on infinite black reflective floor. Spotlight
descends from above in slow crane movement. Gold particles drift lazily through
beam of light. Fog rolls across floor surface. Bottle refracts light into
prismatic patterns. Camera slowly orbits. Ultra-luxury fragrance commercial. 4K.
```

```
Extreme close-up of perfume being sprayed in slow motion, fine mist expanding
in golden backlight, individual droplets visible as they disperse, camera
pulling back to reveal the bottle silhouette, warm amber tones, sensual
luxury aesthetic, 4K.
```

### 4. Real Estate / Architecture Walkthrough
**Characteristics:** Smooth gimbal movement at waist height, wide-angle lens, natural light emphasis, slow and steady pace, lingering on views and architectural details.

```
Smooth cinematic walkthrough entering modern luxury apartment through front
door. Camera at waist height gliding forward through entry hall into open-plan
living area. Floor-to-ceiling windows reveal city skyline. Golden hour light
streaming across hardwood floors. Wide-angle gimbal movement. Real estate
video. 4K.
```

```
Aerial establishing shot of Mediterranean villa, camera descending from high
angle toward courtyard pool. Transition to ground-level walkthrough through
arched doorways, natural stone walls, into bright white kitchen. Warm natural
sunlight. Luxury property tour. 4K.
```

### 5. Social Media Ad (Direct Response)
**Characteristics:** Fast-paced, attention in first 2 seconds, product prominent immediately, vertical format, high energy, clear value proposition shown visually, loop-friendly ending.

```
Vertical 9:16. Close-up of hand grabbing product off shelf. Quick cut to
product in use. Satisfying result shown. Bright, poppy colors. Fast-paced
editing feel with multiple angles in 5 seconds. Mobile-first, TikTok/Reels
native. Product always visible. 4K.
```

```
Vertical 9:16. Dramatic before/after reveal. Left side shows problem, hand
swipes right to reveal solution using product. Clean bright lighting. Product
centered in frame. Quick, punchy, satisfying. Social media ad. 4K.
```

### 6. Unboxing / Product Demo
**Characteristics:** Top-down or slightly angled view, hands interacting with product, slow deliberate movements, ASMR-satisfying textures and sounds implied visually, premium packaging focus.

```
Top-down shot of hands slowly opening matte black magnetic-closure box.
Interior reveals product nestled in custom foam insert. Hands lift product
out reverently. Close-up details. Tissue paper, embossed logo on box.
Slow, deliberate. Soft overhead lighting. Clean desk. Premium unboxing
experience. 4K.
```

```
Tabletop demo video, product placed on clean surface. Hands demonstrate
key feature — twist, click, unfold. Camera alternates between wide context
and close-up detail. Even lighting. Simple clear background. Product
tutorial style. 4K.
```

### 7. Testimonial / Lifestyle
**Characteristics:** Natural, authentic feel. Real environments (homes, offices, outdoors). Warm natural lighting. Handheld or gentle gimbal movement. Person interacting naturally with product.

```
Woman in bright modern kitchen using blender, natural window light, candid
relaxed smile, product on countertop clearly visible, shallow DOF background,
warm lifestyle aesthetic, gentle handheld camera movement, authentic brand
storytelling, 4K
```

```
Man sitting in home office, picks up premium coffee tumbler, takes a sip with
satisfied expression, returns to laptop. Natural daylight from window. Warm,
authentic, aspirational lifestyle. Shallow DOF. Gentle camera drift. 4K.
```

### 8. Music Video / Brand Film
**Characteristics:** Highly stylized, strong color grading, creative camera movements, narrative or abstract, art direction over product focus, emotional resonance.

```
Stylized brand film, person walking through field of tall golden wheat at
sunset, camera tracking alongside at shoulder height, anamorphic lens flare
from sun, strong color grade with lifted shadows and warm highlights, 35mm
film grain, cinematic brand storytelling, 2.39:1 widescreen, 4K
```

```
Abstract brand film, overhead shot of dancer in flowing white fabric
performing on black floor, fabric creating sculptural shapes with each
movement, dramatic single spotlight following the dancer, artistic,
evocative, high-fashion brand film aesthetic, 4K
```

### 9. Beverage Commercial
**Characteristics:** Liquid dynamics, condensation, ice, pour shots, refreshment cues. Bright and vibrant for soft drinks, dark and moody for spirits.

```
Close-up of ice-cold craft beer being poured into frosted glass, golden
liquid cascading with dynamic foam formation, bubbles rising through amber
body, condensation running down glass exterior, warm pub lighting with
dark wood background, beverage commercial, slow motion, 4K
```

```
Cocktail being prepared in slow motion. Ice cubes dropping into crystal glass.
Spirit poured over ice creating amber cascade. Citrus peel twisted and dropped,
releasing oils. Overhead camera looking down into glass. Dark bar. Dramatic
single light. Premium spirits commercial. 4K.
```

### 10. Food Commercial
**Characteristics:** Appetite appeal above all, warm color temperature, implied texture and taste through visual cues, steam/drizzle/cheese pull as hero moments, fast-paced multi-shot edits or single indulgent slow-motion.

```
Slow-motion overhead shot of pizza being pulled apart, long stretchy mozzarella
cheese pull between two halves. Camera follows the cheese stretch. Warm oven-glow
lighting. Rustic wooden board. Steam rising from fresh dough. Basil leaves
settling. Authentic artisan pizzeria commercial. 120fps slow motion. 4K.
```

```
Macro close-up of chocolate sauce being drizzled in slow spiral over stack of
fluffy pancakes. Sauce pooling and cascading off edges. Fresh berries on top.
Single warm directional light from upper-left. Wisps of steam. Breakfast
restaurant commercial. Indulgent, appetite appeal. 4K.
```

```
Fast-paced sushi preparation montage. Chef's knife slicing through salmon
in one clean motion. Rice being shaped by expert hands. Wasabi dabbed with
precision. Final beauty shot: omakase platter arrangement on black lacquer
surface with dramatic top-down lighting. Japanese restaurant commercial. 4K.
```

### 11. Fashion Runway / Lookbook
**Characteristics:** Full-body movement, garment visibility paramount, clean backdrops, consistent even lighting, confident model walk, outfit changes between cuts, smooth tracking or static camera.

```
Model walks toward camera on clean white cyclorama backdrop. Fluid, confident
stride. Camera at hip height, slowly dollying backward to maintain distance.
Outfit: tailored double-breasted blazer over silk camisole, wide-leg trousers,
statement gold earrings. Garment movement visible — trouser fabric swaying,
blazer opening with each step. Flat even lighting from overhead panels. Model
reaches center mark, pauses, quarter-turn showing back detail. Fashion lookbook
video, NET-A-PORTER style. 4K.
```

```
Close-up tracking shot of model walking on runway, camera focused on lower body
showing fabric movement of pleated midi skirt, shoes striking runway floor,
dramatic side lighting creating shadows in pleats, audience bokeh visible on
sides, fashion week runway footage, 4K.
```

### 12. Tech Demo / SaaS Product
**Characteristics:** Screen content visible but secondary to human context, modern workspace environments, natural daylight, collaborative energy, clean and optimistic tone, product-in-use rather than product-alone.

```
Close-up of hands typing on laptop keyboard, screen showing clean SaaS dashboard
with data visualizations. Camera slowly orbits from over-the-shoulder to side
angle revealing the user's focused expression. Modern collaborative workspace
with team members in background, natural daylight, plants, warm productive
atmosphere. Clean, professional, optimistic. Corporate tech commercial, Slack
or Notion brand video aesthetic. 4K.
```

```
Split-screen style: left side shows designer working on tablet with stylus,
right side shows the design appearing in real-time on large wall-mounted
display. Modern creative studio. Team reacting positively in background.
Natural light. Creative SaaS product demo. Collaborative energy. 4K.
```

---

## Stock Footage Styles

### 1. Corporate / Business
**Key characteristics:** Diverse teams, bright modern offices, collaboration, positive energy, natural light, shallow DOF.

```
Diverse team of five professionals collaborating around glass whiteboard in
bright modern open-plan office, natural light from large windows, shallow
DOF on gesturing speaker with others listening engaged, warm positive
atmosphere, corporate stock footage, 4K
```

```
Close-up of hands typing on laptop keyboard, screen reflecting in eyeglasses
of focused professional, shallow DOF, modern desk with plant and coffee,
warm natural sidelight, business productivity, stock footage, 4K
```

```
Wide establishing shot of modern glass office building exterior at sunrise,
warm golden light reflecting off facade, professionals entering through
revolving doors, urban corporate environment, time-lapse clouds, 4K
```

### 2. Lifestyle / Wellness
**Key characteristics:** Natural settings, authentic moments, warm color grading, gentle movement, aspirational but relatable.

```
Woman practicing yoga on wooden deck overlooking misty mountain lake at
sunrise, tree pose silhouette against golden sky, gentle morning breeze
moving hair, peaceful serene atmosphere, wellness lifestyle, gimbal
tracking shot, 4K
```

```
Friends gathered around outdoor dinner table at golden hour, string lights
above, laughing and passing dishes, wine glasses clinking, shallow DOF on
warm faces, Mediterranean lifestyle aesthetic, organic natural movement, 4K
```

```
Parent and child walking hand-in-hand through autumn forest path, golden
leaves falling, dappled sunlight through canopy, warm color grading, slow
motion, emotional connection, lifestyle stock footage, 4K
```

### 3. Abstract / Technology
**Key characteristics:** Particles, data visualization, neural networks, geometric forms, dark backgrounds with vibrant accents, futuristic.

```
Abstract visualization of data flowing through neural network, glowing blue
and purple particles connecting in organic web structure, camera slowly
pulling through the network, dark void background, depth of field on
nearest nodes, futuristic tech abstract, seamless loop, 4K
```

```
Geometric crystal structure assembling from scattered polygons, each face
catching iridescent light, floating in dark space, slow hypnotic rotation,
abstract beauty, generative art aesthetic, seamless loop, 4K
```

```
Macro shot of circuit board with data flowing as light pulses through traces,
blue and gold, shallow DOF with bokeh from distant LEDs, technology abstract,
stock footage, 4K
```

### 4. Nature / Environment
**Key characteristics:** Epic landscapes, aerial perspectives, golden hour, dramatic weather, pristine environments, documentary quality.

```
Aerial drone shot sweeping over Norwegian fjord at sunrise, mist filling
valley between towering emerald cliffs, mirror-still water reflecting
mountains, golden directional light, epic landscape, nature documentary,
4K
```

```
Close-up time-lapse of flower blooming, petals unfurling from tight bud to
full bloom over 10 seconds, soft studio lighting revealing petal texture and
color transition, black background, botanical time-lapse, 4K
```

```
Powerful ocean wave forming and breaking in ultra slow motion, sunlight
penetrating turquoise water revealing internal structure, spray and foam
in dramatic backlight, cinematic ocean footage, 4K
```

### 5. Urban / City
**Key characteristics:** Architecture, street life, light trails, day-to-night transitions, energy and movement, diverse perspectives.

```
Hyperlapse of busy city intersection from overhead, pedestrians crossing in
organized chaos, vehicle light trails during transition from day to night,
illuminated skyscrapers in background, urban energy, 4K
```

```
Rain-soaked city street at night, neon signs reflecting in puddles, taxi
passing through frame creating spray, pedestrians with umbrellas, cinematic
urban atmosphere, Blade Runner aesthetic, 4K
```

```
Golden hour rooftop establishing shot of city skyline, camera slowly panning
180 degrees revealing full panorama, warm light on building facades, long
shadows between towers, urban beauty, stock footage, 4K
```

### 6. Technology / Innovation
**Key characteristics:** Clean labs, screens with data, robotic movement, futuristic interfaces, precision and progress.

```
Close-up of 3D printer nozzle depositing material layer by layer, camera
slowly pulling out to reveal complex geometric object taking shape, blue
LED ambient lighting, shallow DOF, technology innovation, 4K
```

```
Robotic arm in automotive factory performing precise welding, sparks creating
dynamic light in dark industrial space, camera tracking alongside robot's
smooth movement, manufacturing technology, 4K
```

---

## Motion Graphics Keywords & Prompts

### Kinetic Typography
Movement-based text animation. Text as the primary visual element.

**Keywords:** `animated text`, `kinetic typography`, `word-by-word reveal`, `text explosion`, `typewriter effect`, `bouncing text`, `text tracking in/out`, `letter-by-letter assembly`

```
Bold white sans-serif text slamming into frame word by word against pure
black background, each word creating subtle impact ripple, cinematic kinetic
typography, dramatic pacing, motion graphics, 4K
```

### Logo Animation / Reveal
Brand mark appearing through creative animation.

**Keywords:** `logo reveal`, `particle assembly`, `logo morph`, `emergence from particles`, `3D logo rotation`, `liquid metal logo`, `light streak logo formation`, `brand motion identity`

```
Thousands of golden particles swirling in dark void, gradually converging and
assembling into company logo shape, final particles settling with subtle
shimmer, logo solidifies with metallic sheen, dark background, premium
logo reveal animation, 4K
```

```
Liquid chrome flowing across black surface, forming into logo shape through
surface tension, metallic reflections, satisfying fluid dynamics settling
into final mark, minimal, premium brand reveal, 4K
```

### Transitions
Creative shot-to-shot transitions for commercial editing.

**Keywords:** `morph transition`, `zoom transition`, `whip pan`, `seamless transition`, `liquid wipe`, `glitch transition`, `light leak transition`, `particle dissolve`

```
Seamless zoom transition through coffee cup on desk, camera pushes into dark
liquid surface which becomes aerial shot of city at night, lights matching
the dark-to-light pattern, creative seamless transition, 4K
```

### Particle Systems
Atmospheric and decorative particle effects.

**Keywords:** `floating embers`, `sparkle dust`, `particle rain`, `confetti burst`, `firefly particles`, `snow particles`, `ash falling`, `pollen floating`, `light orbs`

```
Thousands of warm golden ember particles floating upward against pure black
background, soft focus, varying sizes, gentle random drift, warm glow on
each particle, cinematic atmosphere, seamless loop, 4K
```

### Liquid Motion
Fluid simulation and organic movement.

**Keywords:** `fluid simulation`, `chrome liquid`, `mercury blob`, `ink drop`, `paint splash`, `oil on water`, `liquid morphing`, `viscous flow`, `lava lamp`

```
Chrome liquid blob floating in zero gravity, morphing through organic shapes,
reflecting studio environment in mirror-like surface, slow satisfying
deformation, dark background, abstract liquid motion, seamless loop, 4K
```

### Geometric Patterns
Mathematical and generative visual patterns.

**Keywords:** `sacred geometry`, `fractal zoom`, `tessellation`, `Voronoi pattern`, `isometric grid`, `geometric transformation`, `platonic solids`, `wireframe morph`

```
Infinite fractal zoom into geometric Mandelbrot pattern, shifting colors from
deep blue to golden, endless detail at every scale, hypnotic mathematical
beauty, abstract generative art, seamless loop, 4K
```

---

## Platform Formats — Complete Reference

| Platform | Format | Aspect Ratio | Resolution | Max Length | Safe Zone |
|----------|--------|-------------|-----------|-----------|-----------|
| Instagram Reels | Vertical | 9:16 | 1080x1920 | 180s (3 min) | 150px top/bottom |
| Instagram Stories | Vertical | 9:16 | 1080x1920 | 60s per story | 250px top, 200px bottom |
| Instagram Feed (Square) | Square | 1:1 | 1080x1080 | 60s | Full frame |
| Instagram Feed (Portrait) | Portrait | 4:5 | 1080x1350 | 60s | Full frame |
| Instagram Carousel | Varies | 1:1 or 4:5 | 1080x1080 / 1080x1350 | 60s per slide | Full frame |
| TikTok | Vertical | 9:16 | 1080x1920 | 10 min | 150px top, 270px bottom |
| YouTube Long | Landscape | 16:9 | 1920x1080 / 3840x2160 | 12 hr | Full frame |
| YouTube Shorts | Vertical | 9:16 | 1080x1920 | 180s (3 min) | 150px top/bottom |
| Twitter/X Video | Landscape | 16:9 | 1920x1080 | 140s | Full frame |
| Twitter/X Vertical | Vertical | 9:16 | 1080x1920 | 140s | Full frame |
| LinkedIn Video | Landscape | 16:9 | 1920x1080 | 15 min | Full frame |
| LinkedIn Vertical | Vertical | 9:16 | 1080x1920 | 15 min | Full frame |
| Facebook Feed | Landscape | 16:9 | 1920x1080 | 240 min | Full frame |
| Facebook Reels | Vertical | 9:16 | 1080x1920 | 90s | 150px top/bottom |
| Pinterest Video | Vertical | 2:3 or 9:16 | 1000x1500 / 1080x1920 | 15 min | Full frame |
| Cinema (Widescreen) | Ultra-wide | 2.39:1 | 2560x1070 | -- | Full frame |
| Cinema (IMAX) | Tall wide | 1.43:1 | 5616x3924 | -- | Full frame |

### Platform-Specific Prompt Modifiers

**TikTok / Instagram Reels:**
```
Vertical 9:16 framing, subject centered in middle 60% of frame, safe zones
150px from top and bottom for UI overlay, mobile-first composition, high
contrast for small screen, bold visual hook in first frame, fast-paced
```

**YouTube (landscape):**
```
Widescreen 16:9, rule of thirds composition, cinematic horizontal framing,
detailed wide shots that reward large screen viewing, standard definition
safe area maintained
```

**Cinema (2.39:1 anamorphic):**
```
Anamorphic 2.39:1 widescreen, dramatic horizontal letterbox frame, subjects
positioned at power points, wide environmental context, cinematic depth,
oval bokeh from anamorphic glass, horizontal lens flares
```

**Instagram Stories (ephemeral):**
```
Vertical 9:16, text-safe zone in center 60% of frame (250px from top,
200px from bottom for username/reply bar), bold graphics, high contrast,
immediate visual impact, designed for quick consumption
```

**LinkedIn (professional):**
```
16:9 landscape preferred, professional tone, well-lit, corporate color
palette, clean composition, text overlays readable at small size, business
context
```

---

## Lighting for Products

### Key Lighting Setups

| Setup | Description | Best For | Prompt Terms |
|-------|------------|----------|-------------|
| **Single large softbox** | One large diffused source from 45deg above-left; soft shadows, gentle highlight gradient | General products, cosmetics, food | `single large softbox from upper-left, soft shadows, gradient highlight` |
| **Two-light (key + fill)** | Main softbox plus weaker fill opposite; reduces shadow density while maintaining dimension | E-commerce, anything needing even detail | `key light from left, gentle fill from right, controlled shadows` |
| **Three-point** | Key + fill + hair/rim light behind subject; maximum dimension and separation | Fashion, portrait-style product, lifestyle | `three-point lighting, rim light from behind separating subject` |
| **Rim / edge light only** | Light behind and beside product, front remains dark; creates dramatic silhouette outline | Premium tech, dark moody aesthetics, luxury | `rim lighting only, dark front face, glowing edges, dramatic silhouette` |
| **Top-down / overhead** | Large diffuser panel directly above; even illumination, minimal visible shadows | Flat lay, tabletop, overhead food | `overhead diffused lighting, even illumination from above, no side shadows` |
| **Low-key** | Single focused light on product, everything else falls to deep black; maximum drama | Luxury, spirits, watches, jewelry, premium | `low-key lighting, single spotlight, deep blacks, dramatic shadows` |
| **High-key** | Multiple large sources filling scene with light; white or bright background | Clean beauty, health, corporate, e-commerce | `high-key lighting, bright even illumination, white background, minimal shadows` |
| **Backlight / transillumination** | Primary light behind the subject; reveals internal color, translucency, and structure | Beverages, glass bottles, gemstones, liquids | `backlit to reveal liquid color, transillumination, glowing from behind` |
| **Strip light** | Narrow rectangular softbox; creates long thin specular highlights on curved surfaces | Cylindrical products: bottles, cans, tubes | `strip light creating long specular highlight along curve, controlled reflection` |
| **Clamshell** | Two lights above and below face/product at equal angles | Beauty portraits, cosmetics on face | `clamshell lighting, catchlight above and below, even facial illumination` |
| **Light tent / light box** | Product surrounded by diffused light from all sides; zero harsh shadows | E-commerce packshots, small products | `light tent, even 360-degree diffused lighting, no harsh shadows, clean` |

### Material-Specific Lighting

| Material | Lighting Approach | Key Prompt Terms |
|----------|------------------|-----------------|
| **Glass / transparent** | Backlight or side-light on white; dark-field technique for edge definition | `backlit translucency, edge definition, internal refraction, dark-field glass photography` |
| **Metal / chrome** | Large soft sources for broad, smooth reflections; gradient lighting | `gradient reflection on metal, studio environment reflected in chrome, controlled specular highlights` |
| **Brushed metal** | Directional soft light aligned with or perpendicular to grain for anisotropic effect | `anisotropic highlight following grain direction, brushed aluminum sheen, satin finish` |
| **Fabric / textile** | Side-light at raking angle to reveal weave and texture | `raking light revealing fabric texture, visible weave pattern, textile macro, thread detail` |
| **Matte surfaces** | Soft wrap-around light; matte materials absorb rather than reflect | `soft diffused wrap lighting, matte finish, no specular highlights, even illumination` |
| **Food** | Backlight or strong side-light for steam and texture; warm 3200K color temp | `warm backlight through steam, glistening texture, appetite appeal lighting, golden warmth` |
| **Liquid in glass** | Backlight for color and clarity; side-light for surface texture and bubbles | `backlit liquid showing color depth, side-lit effervescence, surface tension highlights` |
| **Skin (beauty)** | Beauty dish or large softbox from above; fill from below to minimize chin shadow | `beauty dish lighting, soft specular skin highlights, butterfly shadow under nose` |
| **Leather** | Moderate side-light to reveal grain and patina without washing out | `directional light revealing leather grain, warm highlights on patina, rich depth` |
| **Wood** | Side-light at 30-45deg to reveal grain pattern and surface texture | `raking light on wood grain, warm tones, visible texture, natural material` |

---

## Surface & Background Guide

| Surface Type | Prompt Description | Best For |
|-------------|-------------------|----------|
| **White seamless** | `clean seamless white background, infinite white, no horizon line` | E-commerce, medical, tech, food, packshots |
| **Black seamless** | `pure black seamless background, dark void, no visible surface or horizon` | Luxury, jewelry, watches, premium tech, spirits |
| **Dark gradient** | `dark charcoal gradient background fading to black at edges` | Hero shots, dramatic product photography |
| **Colored gradient** | `smooth [color] gradient background, seamless color transition` | Tech products, modern branding, creative |
| **Calacatta marble** | `polished Calacatta marble surface with soft grey veining` | Luxury, beauty, jewelry, spirits, cosmetics |
| **Terrazzo** | `pink terrazzo surface with multicolored aggregate chips` | Beauty, modern lifestyle, cosmetics |
| **Polished concrete** | `raw polished concrete surface with subtle industrial texture` | Industrial, tech, minimalist, architectural |
| **Reclaimed wood** | `rustic reclaimed wood table surface with visible grain, knots, and patina` | Food, artisan, craft products, organic |
| **Dark slate** | `wet dark slate surface with subtle sheen from moisture` | Beverages, food, premium products, spirits |
| **Natural linen** | `draped natural linen surface in [color], visible loose weave texture` | Flat lay, lifestyle, organic, artisan products |
| **Reflective black acrylic** | `polished black acrylic surface creating perfect mirror reflection below product` | Automotive, luxury, tech, high-end |
| **Obsidian / volcanic** | `polished obsidian block, glossy black natural stone` | Jewelry, watches, premium small goods |
| **Sand** | `fine white sand surface with subtle ripple patterns` | Summer products, sunscreen, resort, outdoor |
| **Tropical foliage** | `bed of fresh green monstera and palm leaves as surface` | Natural products, organic, botanical brands |
| **Velvet** | `draped [color] velvet fabric with rich light-absorbing folds` | Jewelry, luxury, cosmetics, premium goods |
| **Copper / bronze** | `aged patina copper sheet surface with verdigris accents` | Artisan, heritage, craft spirits, food |

**Choosing backgrounds by product type:**
- **Tech products**: Black gradient, white seamless, colored gradient, polished concrete
- **Food & beverage**: Wood, slate, marble, linen, raw stone
- **Beauty / cosmetics**: Marble, terrazzo, pink gradient, white seamless, velvet
- **Luxury goods**: Black seamless, reflective black, obsidian, silk, velvet
- **Lifestyle**: Linen, wood, concrete, natural environment, sand
- **Spirits / beverages**: Wet slate, dark wood bar surface, reflective black, ice surface
- **Jewelry**: Black silk, obsidian, velvet, reflective black acrylic

---

## Color Psychology in Commercial Content

| Color | Hex Reference | Association | Commercial Use Cases |
|-------|--------------|------------|---------------------|
| **Red** | #E63946 | Urgency, appetite, passion, energy, power | Food brands, sale events, fast-fashion, sports, entertainment |
| **Orange** | #FF6B35 | Enthusiasm, creativity, affordability, warmth | Youth brands, food delivery, CTA buttons, budget-friendly brands |
| **Yellow** | #FFD166 | Optimism, warmth, attention-grabbing, happiness | Kids' products, clearance sales, sunshine brands, caution |
| **Green** | #2D6A4F | Nature, health, sustainability, money, growth | Organic food, wellness, finance, eco brands, plant-based |
| **Blue** | #1D3557 | Trust, professionalism, calm, security, technology | Corporate, finance, tech, healthcare, insurance, SaaS |
| **Purple** | #7B2D8E | Luxury, creativity, spirituality, mystery | Premium brands, beauty, creative tools, confectionery |
| **Gold** | #D4A574 | Luxury, prestige, quality, exclusivity, heritage | High-end watches, spirits, fashion houses, awards |
| **Black** | #1A1A1A | Premium, power, sophistication, edge, authority | Luxury fashion, tech, automotive, professional services |
| **White** | #FFFFFF | Clean, simple, pure, modern, medical, minimal | Skincare, tech, minimalist brands, healthcare, bridal |
| **Pink** | #FFB4C2 | Playful, feminine, romantic, youthful, sweet | Beauty, confectionery, romantic brands, Gen Z targeting |
| **Teal** | #2EC4B6 | Modern, fresh, digital, innovative, balanced | Tech startups, health-tech, modern wellness, fintech |
| **Burgundy** | #800020 | Sophistication, wine, warmth, heritage, depth | Wine brands, heritage fashion, luxury interiors, fine dining |

**Using color in prompts:** Specify hex codes for brand-critical colors. AI models (especially the Flux family) respond well to hex color codes: `"product packaging in brand teal #2EC4B6 against neutral background"`. This produces more accurate color matching than descriptive words alone.

**Color combinations for commercial tone:**
- **Premium/Luxury**: Black + Gold, Navy + Gold, Black + White
- **Fresh/Natural**: Green + White, Earth tones + Cream
- **Energetic/Youth**: Bright Orange + Electric Blue, Hot Pink + Yellow
- **Corporate/Trust**: Navy + Light Blue, Dark Blue + White
- **Warm/Comfort**: Amber + Cream, Terracotta + Sage

---

## Ad Copy Integration

When generating images intended for advertising, you must plan for text overlay in post-production. AI should never generate the text itself (except with Recraft V4 or Ideogram V3 which handle text) — only the visual canvas with space reserved.

### Designing for Text Overlay

**Rule of thirds for ad layouts:**
- **Left third empty, product right**: Headline and body copy go left — classic magazine layout
- **Top third empty, product centered-low**: Headline above, CTA below — billboard layout
- **Bottom third empty, product fills top**: Product hero with CTA bar beneath — web banner layout
- **Center empty, product framing edges**: Text in center surrounded by product/lifestyle — social media

**Prompt additions for text-safe images:**
```
generous negative space on the [left/right/top] side of the composition for
text overlay placement, clean uncluttered area suitable for headline text,
advertising layout with room for copy, no visual elements in [position]
```

### Safe Zone Prompting Examples

**Magazine / print ad layout:**
```
Product positioned in right two-thirds of frame, left third is clean [color]
background with no detail — reserved for headline and body copy placement,
advertising layout composition, print ad ready, 4K
```

**Billboard / OOH layout:**
```
Product hero shot centered in lower portion of frame, upper 40% is clean
gradient sky or background with no competing elements — space reserved for
brand logo and headline, billboard-ready composition, high impact, 4K
```

**Social media ad layout:**
```
Product on clean background with 60% negative space around it, centered
subject with breathing room on all sides for platform-specific text overlays,
social media advertising layout, mobile-optimized, 4K
```

### Ad-Ready Composition Examples

```
Premium skincare bottle on white marble surface positioned in right third of
frame, left two-thirds is soft warm gradient from peach to white with no
elements — clean space for headline text and brand logo, advertising layout,
beauty commercial, 85mm f/2.8, 4K
```

```
Athletic shoe in dynamic angle filling bottom-right corner of frame,
upper-left two-thirds is dramatic dark gradient with subtle motion blur
streaks — space designed for bold typography overlay, Nike-style advertising
layout, high energy, 4K
```

```
Luxury watch on reflective black surface, centered in bottom third of frame,
upper two-thirds is pure black with single accent light streak — space for
brand name and tagline in gold serif typography (to be added in post), premium
watch advertising layout, 8K
```

---

## The 5-Part Commercial Prompt Formula

Every effective commercial prompt contains these five components. Order matters — models weight earlier tokens more heavily.

### Part 1: SUBJECT (The Hero)
The product or person. This MUST come first. Be hyper-specific about materials, colors, brand cues, and form factor.

**Weak:** `a bottle of perfume`
**Strong:** `translucent crystal-cut perfume bottle with gold cap, amber liquid visible through faceted glass, brand label catching light`

### Part 2: ACTION (The Verb)
What is happening? Static shots need implied energy. Video needs explicit motion description.

**For images — implied action:** `mid-pour`, `suspended mid-air`, `fabric caught in wind`, `steam rising`
**For video — explicit motion:** `slow 360-degree rotation`, `liquid pouring from bottle`, `camera pushing in from medium to close-up`

### Part 3: ENVIRONMENT (The Stage)
Surface, background, atmospheric elements. Environment sets the emotional context.

**Studio:** `on polished marble surface, dark gradient background, controlled studio environment`
**Location:** `in sun-drenched modern bathroom with marble countertops, morning light through frosted window`
**Abstract:** `floating in dark void, particle dust catching rim light`

### Part 4: TECHNICAL (The Craft)
Lens, aperture, lighting setup, camera movement, resolution. This is what separates amateur from professional.

**Lighting vocabulary:**
- `single overhead softbox` — broad, even, flattering
- `Rembrandt lighting` — dramatic 45-degree with triangle shadow on cheek
- `rim lighting` / `edge lighting` — backlight defining product silhouette
- `clamshell lighting` — beauty standard (above + below fill)
- `practical lighting` — light sources visible in scene (lamps, candles, screens)
- `natural window light` — soft, directional, authentic
- `hard spotlight` — focused beam, sharp shadows, drama
- `two-point lighting` — main + fill, standard product setup

**Lens vocabulary:**
- `24mm wide-angle` — environmental context, real estate, establishing
- `35mm` — street, documentary, natural perspective
- `50mm` — standard, true-to-eye, versatile
- `85mm f/1.4` — portrait, beauty, product isolation, creamy bokeh
- `100mm macro` — detail, jewelry, texture, extreme close-up
- `200mm telephoto` — compression, isolation, sports, automotive

**Camera movement vocabulary (video):**
- `static locked-off tripod` — stability, authority, product focus
- `slow dolly in` — building attention, approaching the subject
- `slow dolly out` — revealing context, expanding the view
- `orbit / arc` — 3D exploration of product, showing all angles
- `crane up` — establishing grandeur, revealing scale
- `tracking shot` — following alongside subject movement
- `gimbal float` — smooth natural movement, walkthrough feel
- `whip pan` — fast transition, energy, excitement

### Part 5: STYLE MODIFIER (The Aesthetic)
The overall creative direction. Reference specific aesthetics, publications, or brand styles.

**Examples:** `Apple product page aesthetic`, `Vogue Italia editorial`, `Nike campaign energy`, `Wes Anderson color palette`, `luxury fragrance commercial`, `editorial beauty photography`, `food editorial for Bon Appetit`, `architectural digest real estate`

### Complete Formula Examples

**Skincare Commercial (Image):**
```
[SUBJECT] Close-up of translucent green serum bottle with gold dropper cap,
liquid catching light through glass.
[ACTION] Golden serum mid-drop from dropper, thread of liquid stretching.
[ENVIRONMENT] Sun-drenched modern bathroom, white marble countertop, soft
morning light from frosted window.
[TECHNICAL] 85mm f/1.8, shallow DOF, warm natural sidelight with soft fill,
bokeh from bathroom fixtures.
[STYLE] Premium skincare brand campaign, editorial beauty, magazine
advertisement quality, 8K.
```

**Running Shoe Commercial (Video):**
```
[SUBJECT] Performance running shoe with knit upper and reactive foam sole.
[ACTION] Extreme slow motion of shoe striking wet pavement, water erupting
from impact, sole compressing and rebounding.
[ENVIRONMENT] Dark urban street at dawn, wet surfaces, distant city lights.
[TECHNICAL] Ground-level camera, 100mm macro, dramatic rim lighting from
street lamps, 240fps slow-motion feel.
[STYLE] Nike campaign energy, athletic performance commercial, dynamic
sports advertising, 4K.
```

**Luxury Watch (Image):**
```
[SUBJECT] Swiss chronograph with midnight blue dial, applied rose gold
indices, exhibition caseback.
[ACTION] Static, positioned at 10:10 classic time display, crown catching
light.
[ENVIRONMENT] On weathered driftwood, ocean-misted background, seaside
luxury concept.
[TECHNICAL] 100mm macro f/5.6, focus-stacked for full dial sharpness, single
warm directional light, subtle blue fill.
[STYLE] Horological photography, Hodinkee editorial, luxury timepiece
advertising, 8K.
```

---

## Negative Prompts by Product Category

Use these to prevent common AI generation artifacts in commercial content.

**Food:**
```
unappetizing, cold-looking, wilted, artificial food, plastic food, frozen,
overcooked, undercooked, wrong color temperature, dull colors, unnatural
food coloring, messy presentation, dirty plate, fingerprints on glass
```

**Fashion:**
```
mannequin-like, stiff pose, distorted limbs, warped fabric, extra fingers,
wrong proportions, plastic skin, dead eyes, uncanny valley, garment clipping
through body, impossible drape, wrinkled in wrong places, seam distortion
```

**Automotive:**
```
toy car, distorted proportions, wrong reflections, plastic body panels,
mismatched panel gaps, impossible wheel angles, cartoon-like, wrong number
of wheels, distorted headlights, unrealistic scale, incorrect badge
placement, floating car
```

**Jewelry:**
```
costume jewelry, dull metal, plastic gemstones, no sparkle, wrong
proportions, blurry facets, missing prongs, impossible stone setting,
cheap-looking, no light interaction, flat rendering, missing reflections,
uniform color (gems should have depth)
```

**Tech Products:**
```
thick bezels, wrong proportions, dated design, scratched surface, distorted
screen content, wrong number of cameras, asymmetric design, melted edges,
plastic where metal should be, wrong port placement, text garbled on
screen, impossible thinness
```

**Cosmetics / Beauty:**
```
cakey makeup, unblended, wrong skin undertone, plastic skin texture,
missing pores, airbrushed to uncanny valley, asymmetric features,
product label illegible, wrong viscosity, separated formulas, dried out
```

**Beverages:**
```
flat beer (no bubbles), warm-looking cold drinks, wrong liquid viscosity,
floating ice that defies physics, unnatural liquid color, glass half-empty
when should be full, no condensation on cold glass, perfectly still liquid
(should have surface tension meniscus)
```

---

## Common Issues & Fixes

| Problem | Cause | Fix |
|---------|-------|-----|
| Product morphs during video | Insufficient product description | Over-describe static elements; use I2V with reference frame generated first |
| Metal looks plastic | Missing environment reflection cues | Add `reflecting studio environment`, `anisotropic highlights`, `high dynamic range` |
| Highlights too harsh / blown | Light source too focused | Use `large diffused softbox` or `light tent` instead of `spotlight` |
| Highlights flicker in video | Multiple competing light sources | Commit to single primary light source; use I2V for consistency |
| Background too busy / distracting | Too much environmental detail | Add `clean background`, `negative space`, `seamless backdrop`, `shallow DOF` |
| Text is garbled / unreadable | AI text generation limitation | Use Recraft V4 or Ideogram V3 for text; otherwise overlay text in post |
| Food looks artificial / unappetizing | Generic food descriptors | Use specific culinary terms: `caramelized`, `charred`, `al dente`, `jammy yolk` |
| Fabric looks stiff / unnatural | Missing movement and drape cues | Add `natural drape`, `fabric catching wind`, `soft folds`, `gravity-pulled` |
| Skin looks waxy / plastic | Over-smoothing tendency | Add `visible pore texture`, `natural skin luminosity`, `subsurface scattering` |
| Product floating unrealistically | Missing surface/shadow cues | Add `contact shadow on surface`, `subtle reflection on [surface material]` |
| Jewelry has no sparkle | Missing light interaction terms | Add `specular highlights`, `caustics`, `prismatic fire`, `facet reflections` |
| Car reflections look wrong | Missing environment context | Describe what the car should reflect: `reflecting studio softboxes`, `reflecting sky gradient` |
| Liquid looks like gel | Wrong viscosity descriptors | Specify `water-like fluidity`, `low viscosity`, `splashing readily` or `viscous, slow-moving, honey-like` |
| Background and product same tone | No separation lighting | Add `rim light separating subject from background`, `edge glow`, `backlight` |
| Video has temporal flickering | Model instability with complex scenes | Simplify scene, reduce number of moving elements, use `consistent lighting` |
| Colors look desaturated | Missing vibrancy cues | Add `rich saturated colors`, `vibrant`, or specific color names like `deep crimson` |

---

## Commercial Production Workflows

### Workflow 1: Image-to-Video Product Commercial Pipeline

The most reliable method for consistent commercial video. Generate a perfect still first, then animate.

1. **Generate hero image** using Flux 2 Pro or Imagen 4
   - Perfect composition, lighting, product placement
   - Generate 3-5 variations, select the strongest
2. **Upscale if needed** using Topaz Upscale
3. **Generate video from image** using Veo 3.1 or Kling v3 Pro I2V
   - Use the hero image as the first frame / reference
   - Describe ONLY the motion — the image already defines the look
   - Keep camera movement simple: one movement type
4. **Optional: Generate additional angles** as separate shots
5. **Sequence shots** for multi-shot commercial feel

### Workflow 2: Background Replacement Pipeline

Replace product background while maintaining lighting consistency.

1. **Generate or photograph product** on any background
2. **Remove background** using Bria RMBG 2.0 (general) or BiRefNet v2 (fine detail / hair)
3. **Generate new background** separately using Flux 2 Pro or Imagen 4
   - Match lighting direction and color temperature to the product
4. **Composite** using Flux 2 Edit or GPT-Image 1.5 Edit
   - Describe the composite: "Place this product on [new environment], matching lighting"
5. **Refine** if needed with another edit pass for shadow and reflection matching

### Workflow 3: Multi-Shot Video Commercial

Create a cohesive multi-shot commercial by generating consistent shots.

1. **Define the shot list** (typically 4-8 shots for a 15-30s commercial):
   - Shot 1: Establishing / environment
   - Shot 2: Product reveal / hero
   - Shot 3: Detail / feature close-up
   - Shot 4: Product in use / lifestyle
   - Shot 5: Final beauty shot / pack shot
2. **Generate hero image** first as the "look" reference
3. **Generate each shot as I2V** maintaining consistent:
   - Color temperature and lighting direction
   - Product appearance (use same product description across all prompts)
   - Style modifier (keep identical across shots)
4. **Maintain consistency** by using identical style/technical blocks across all prompts

### Workflow 4: Product Photography Batch (E-Commerce)

Generate a complete product image set for e-commerce listing.

1. **Hero shot** — three-quarter angle, dramatic lighting (Flux 2 Pro)
2. **Front view** — flat, even lighting, pure white background
3. **Back view** — same setup as front
4. **Detail shots** (2-3) — macro on key features, textures, materials
5. **Lifestyle / in-context** — product in use or in environment
6. **Scale shot** — product next to recognizable object for size reference
7. **Remove backgrounds** on product-only shots using Bria RMBG 2.0
8. **Upscale all** using Topaz for print-ready resolution

### Workflow 5: Social Media Content Batch

Generate a complete social media campaign from one product.

1. **Hero product image** (1:1 for feed, 4:5 for Instagram portrait)
2. **Lifestyle image** (product in context, same aspect ratios)
3. **Video: product reveal** (9:16 vertical, 15-30s for Reels/TikTok)
4. **Video: product in use** (9:16 vertical, 15-30s)
5. **Carousel images** (4-6 images at 1:1 or 4:5, showing different angles/features)
6. **Story-format content** (9:16, designed for ephemeral consumption)
7. Apply platform-specific prompt modifiers (safe zones, framing) to each

### Workflow 6: Virtual Try-On for Fashion E-Commerce

Generate product-on-model images without physical photoshoots.

1. **Generate or obtain model image** (or use customer-uploaded photo)
2. **Generate garment flat lay** or use existing product image
3. **Apply try-on** using FASHN v1.6 with garment and model images
4. **Generate additional poses/angles** with Flux Kontext for consistency
5. **Remove and replace backgrounds** for different lifestyle contexts

---

## Luxury & Premium Content Techniques

### The Luxury Lighting Trinity

1. **Key light** (single, controlled, directional) — creates the drama
2. **Negative fill** (black card/flag removing light) — deepens shadows for contrast
3. **Accent/rim light** (edge separation) — defines form against background

Premium content uses MORE shadow than light. Restraint is luxury.

### Material-Specific Rendering

**Brushed metal:** `anisotropic highlights stretching along brush direction, soft satin sheen, fine parallel scratches visible in macro, controlled studio reflections`

**Polished chrome/mirror:** `sharp specular highlights, mirror-like reflections of studio environment, high dynamic range, pristine surface, no fingerprints`

**Matte surfaces:** `soft diffuse reflection, no specular hotspots, velvet-like light absorption, subtle surface texture visible in raking light`

**Glass/crystal:** `internal refractions visible, caustic light patterns cast on surface below, Fresnel effect at glancing edges, visible thickness, trapped light at edges`

**Leather:** `rich patina developing at wear points, visible grain texture, subtle sheen on high points catching directional light, warm organic tones`

**Silk/satin fabric:** `specular highlights following drape contours, deep shadows in folds, lustrous sheen, visible thread count at macro, color shift between highlight and shadow`

**Velvet:** `light-absorbing shadows, glowing highlights where nap catches light, rich depth of color, texture visible as directional sheen`

**Marble:** `translucent edges where thin, visible veining pattern, polished surface reflections, cool undertones, substantial weight implied`

### Color Grading Descriptors for Premium

- **Luxury warm:** `rich shadows, warm highlight roll-off, slight amber shift in midtones, deep blacks`
- **Premium cool:** `steel blue shadows, neutral highlights, desaturated midtones, high contrast`
- **Editorial muted:** `lifted shadows, desaturated pastels, low contrast, dreamy, film-like`
- **High fashion:** `extreme contrast, deep blacks, vivid selective color, dramatic`
- **Heritage/classic:** `warm sepia undertones, soft contrast, film grain, timeless`

---

## 20+ Ready-to-Use Commercial Prompt Examples

### 1. Skincare Serum — Hero (Image)
```
Premium glass serum bottle with rose gold dropper cap on white Carrara marble
slab, single drop of golden serum suspended mid-fall from dropper, bottle
three-quarter angle showing translucent amber liquid, soft directional light
from upper left with warm fill, white studio background with subtle gradient,
luxury skincare product photography, 85mm f/2.8, 8K
```

### 2. Sneaker Launch — Dynamic (Image)
```
Latest-generation running sneaker captured mid-air against pure black
background, dynamic angle showing sole tread and knit upper, colorful dust
particles exploding from shoe as if from impact, dramatic rim lighting
from both sides in brand orange and white, floating hero product shot,
athletic footwear commercial, 8K
```

### 3. Coffee Brand — Lifestyle (Image)
```
Hands cupping ceramic mug of fresh cappuccino with latte art, steam wisps
rising in warm morning sunlight from side window, cozy cafe setting with
soft bokeh background, wooden table, warm natural color palette, lifestyle
coffee brand photography, 50mm f/1.8, shallow DOF, 4K
```

### 4. Luxury Watch — Editorial (Image)
```
Swiss dive watch on wet volcanic rock by the ocean, water droplets on sapphire
crystal and brushed steel case, golden hour sunlight from low angle creating
long shadows, ocean spray mist in background, watch showing 10:10 display,
rugged luxury aesthetic, adventure timepiece photography, 100mm, 8K
```

### 5. Craft Beer — Pour Shot (Video)
```
Dramatic slow-motion pour of amber craft beer into branded pint glass. Golden
liquid cascading with dynamic foam formation. Camera at glass level, slightly
below. Backlit to illuminate beer color. Condensation on glass. Dark pub
background with warm bokeh. Beverage commercial. 4K.
```

### 6. Electric Vehicle — Reveal (Video)
```
Darkness. Floor-level camera. Single line of light appears, tracing the
silhouette of an electric sedan from front to rear. Full studio lighting
builds gradually over 6 seconds revealing matte gray finish. Camera slowly
rises from ground level to belt-line. Black studio. Automotive reveal. 4K.
```

### 7. Perfume — Luxury (Video)
```
Crystal perfume bottle on infinite black reflective surface. Single beam of
light descends from above in slow motion. Gold particles drift through light
beam. Bottle refracts light creating prismatic rainbow patterns on floor.
Camera slowly orbiting. Fog at ground level. Ultra-luxury fragrance ad. 4K.
```

### 8. Pizza Restaurant — Food (Image)
```
Fresh Neapolitan pizza emerging from wood-fired oven on long wooden peel,
leopard-spotted charred crust bubbling, mozzarella stretching as slice
separates, basil leaves wilting from heat, orange firelight from oven
illuminating scene, dark rustic pizzeria, steam and smoke wisps, authentic
Italian food photography, 8K
```

### 9. Fitness App — Social Media (Video)
```
Vertical 9:16. Fit woman in athletic wear checking smartwatch, starts running
in urban park at sunrise. Camera tracking alongside at shoulder height.
Morning golden light. Trees and path in background with bokeh. Energetic
but smooth. Subject centered in safe zone. Fitness lifestyle. Social media
ad format. 4K.
```

### 10. Headphones — Tech (Image)
```
Over-ear wireless headphones floating against matte dark gray gradient,
left ear cup angled toward camera showing driver detail, soft cool-white
lighting from left revealing brushed aluminum and memory foam, thin rim
glow separating from background, no surface, no shadow, tech product
hero, Apple product page aesthetic, 8K
```

### 11. Candle Brand — Lifestyle (Image)
```
Premium soy candle in minimalist concrete vessel, lit with warm flickering
glow illuminating immediate surroundings, placed on linen-draped bedside
table, soft evening atmosphere, cozy bedroom background in deep blur,
warm intimate color palette, hygge lifestyle, candle brand photography,
50mm f/1.4, 4K
```

### 12. Sunglasses — Fashion Editorial (Image)
```
Close-up of model wearing oversized sunglasses, lenses reflecting tropical
beach scene, strong midday sun creating sharp shadow from frames across
cheekbones, golden tanned skin, slight head tilt, fashion editorial,
poolside luxury, vacation brand campaign, 85mm, shallow DOF, 8K
```

### 13. Smartphone — Feature Demo (Video)
```
Close-up of hand holding smartphone, thumb swiping through camera interface.
Phone raised to eye level, captures photo — screen flash visible. Hand lowers
phone to review shot on screen. Clean modern interior background with soft
natural light. Smooth, natural gestures. Tech product demo. 4K.
```

### 14. Wine — Product (Image)
```
Bottle of premium red wine beside filled crystal glass on aged oak barrel
top, deep garnet liquid visible through glass catching warm sidelight, wine
bottle label in sharp focus, dark cellar background with single shaft of
light, grape vine leaves as prop, traditional wine photography with modern
lighting, 8K
```

### 15. Real Estate — Establishing (Video)
```
Drone ascending from garden level of modern glass-and-timber mansion,
revealing infinity pool, then surrounding landscape of hills and vineyards.
Golden hour. Camera rises smoothly from eye level to 50 meters. Warm
directional sunlight. Luxury property establishing shot. Real estate
commercial. 4K.
```

### 16. Backpack — Adventure (Image)
```
Technical hiking backpack resting against a rock on mountain summit, dramatic
alpine panorama in background at golden hour, pack slightly open showing
organized compartments, trekking poles leaning alongside, warm directional
sunlight, adventure brand campaign, environmental product photography,
24mm wide-angle, 8K
```

### 17. Baby Product — Tender Lifestyle (Video)
```
Soft close-up of parent gently applying baby lotion to infant's hand.
Baby's tiny fingers curling around parent's thumb. Warm soft window light.
White and cream nursery setting in background blur. Tender, intimate,
emotional. Gentle slow motion. Baby care brand commercial. 4K.
```

### 18. Chocolate — Indulgence (Image)
```
Dark chocolate bar broken in half revealing interior snap texture, cocoa
powder dusting scattered on dark slate surface, chocolate shavings and
whole cocoa beans as props, single warm spotlight creating rich deep shadows,
moody indulgent atmosphere, artisanal chocolate brand photography,
macro detail on snap texture, 8K
```

### 19. Hotel / Travel — Aspirational (Video)
```
Slow tracking shot through luxury resort, starting at infinity pool edge
with water meeting ocean horizon. Camera drifts right past lounge chairs
with white cushions, through tropical garden, arriving at open-air restaurant
with ocean view. Golden hour. Gentle breeze moving palm fronds. Travel
luxury commercial. 4K.
```

### 20. Laptop — Workspace (Image)
```
Premium laptop open on clean minimal desk, screen displaying creative
software interface, warm desk lamp providing accent light from left, small
potted plant and ceramic coffee cup as lifestyle props, natural window
light from right, bright modern workspace, tech lifestyle product
photography, 35mm environmental, 4K
```

### 21. Lipstick — Beauty Close-Up (Video)
```
Extreme close-up of lips being painted with deep red lipstick. Bullet
applying color in slow deliberate stroke from center outward. Creamy texture
visible. Lips slightly parted. Macro camera. Ring light catchlight visible
in lip gloss. Beauty commercial. Slow motion. 4K.
```

### 22. Sports Car — Motion (Video)
```
Low rig shot tracking alongside sports car on coastal highway. Car sharp,
background motion-blurred. Ocean and cliffs visible through blur. Late
afternoon light creating dynamic shadows across body panels. Engine sound
implied through visual vibration. Camera at wheel height. Automotive
commercial. Anamorphic 2.39:1. 4K.
```

### 23. Plant-Based Milk — Splash (Image)
```
Glass of oat milk with dramatic splash frozen mid-air, oat grains and whole
oats suspended around the splash, creamy white liquid with visible texture,
single oat stalk as prop, bright clean white background, fresh healthy
energy, high-speed photography 1/5000s, plant-based beverage commercial, 8K
```

### 24. Interior Paint — Transformation (Video)
```
Time-lapse style video of room transforming. Walls change from dull gray
to rich emerald green as paint roller sweeps across. Furniture fades in.
Natural light shifts from cold to warm. Room goes from empty to styled.
Before-and-after transformation feel. Home improvement commercial. 4K.
```

### 25. Wireless Earbuds — Lifestyle (Video)
```
Young professional on morning commute, removes earbuds from case in pocket.
Inserts one earbud, music implied by subtle head nod and relaxed smile.
City street background with morning light and soft bokeh pedestrians.
Smooth gimbal tracking. Natural, authentic. Audio brand lifestyle. 4K.
```

---

## Quick Reference: Prompt Modifier Cheat Sheet

### Resolution & Quality
`8K`, `4K`, `ultra-high resolution`, `commercial grade`, `print quality`, `magazine quality`, `advertisement quality`

### Lighting Shortcuts
- Drama: `single hard spotlight`, `Rembrandt lighting`, `chiaroscuro`
- Soft/Beauty: `large softbox`, `diffused window light`, `beauty dish`, `clamshell lighting`
- Product: `light tent`, `two-point studio lighting`, `controlled studio`
- Atmosphere: `golden hour`, `neon`, `candle light`, `overcast diffused`

### Composition Shortcuts
- `rule of thirds`, `centered composition`, `negative space for copy`, `overhead flat lay`
- `three-quarter angle`, `eye level`, `low angle hero`, `high angle establishing`

### Mood/Atmosphere
- `moody and dramatic`, `bright and airy`, `warm and cozy`, `cool and clinical`
- `energetic and dynamic`, `calm and serene`, `luxurious and opulent`, `raw and authentic`

### Camera/Lens
- Wide context: `24mm wide-angle`
- Natural: `35mm`, `50mm`
- Portrait/Product: `85mm f/1.4`, `100mm macro`
- Compression: `200mm telephoto`
- Film look: `anamorphic`, `35mm film`, `medium format`
