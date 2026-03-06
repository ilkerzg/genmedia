# World Building — Environment & World Design Master Guide

Comprehensive reference for creating environments, architectural visualizations, interior designs, landscapes, and fantasy/sci-fi worlds using AI generative models. Every section contains directly usable prompt keywords and complete prompt examples. Use `search_models` to find the best model for each task.

---

## Quick Reference

> **Model Discovery:** Use `search_models` to find the best model for each task. Search by category: "image generation" for environments and concept art, "text to video" for flythroughs, "image to video" for animating stills, "image editing" for weather/time-of-day changes, "outpainting" for expanding environments.

**Top 5 World Building Tips**

1. Use the formula: `[Setting Type] + [Architectural Style] + [Time of Day] + [Weather/Atmosphere] + [Lighting] + [Mood] + [Camera/Lens] + [Detail Layer]`
2. Always add a scale reference (a person, vehicle, or familiar object) so the viewer understands environment size.
3. Layer depth deliberately: foreground element, midground focal point, background with atmospheric perspective.
4. Lock color palette, time of day, and style keywords verbatim across all prompts in a series.
5. Generate a strong still first, then animate with I2V — this preserves your exact design when moving to video.

**Key Prompt Patterns**

Exterior: `[Wide shot] of [setting], [architectural style], [time of day] [weather], [lighting], [scale reference], [mood], cinematic, [lens], 8K`

Interior: `[Room type] in [interior style], [material/texture keywords], [light source], [key furniture], [mood], interior design photography, 8K`

Fantasy/Sci-Fi: `[Setting concept], [unique visual hook: bioluminescence / floating / overgrown], [color palette], [scale reference], [mood], concept art, cinematic lighting, 8K`

**Cheat Sheet — Key Prompt Elements by Scene Type**

- Photorealistic exterior day: 16:9, golden hour, tilt-shift or 24mm, add person for scale
- Photorealistic interior: 4:3, specify light source (fireplace / skylight / window)
- Fantasy landscape: 16:9, add "volumetric god rays, concept art, 8K"
- Moody/cinematic: blue hour or golden hour + "cinematic, anamorphic 2.39:1, film grain"
- Night urban: "wet reflective surfaces, neon color, volumetric light shafts, rain"
- Weather/season swap: use an image editing model with a single-sentence edit instruction

---

## 1. The Environment Prompt Formula

```
[Setting Type] + [Architectural Style] + [Time of Day] + [Weather/Atmosphere] + [Lighting] + [Mood] + [Camera/Lens] + [Detail Layer]
```

### Priority Order

| Position | Content | Weight | Example |
|----------|---------|--------|---------|
| 1-15% | Setting type and location | Highest | "A brutalist concrete library" |
| 15-30% | Architectural style and era | High | "in mid-century modern style, 1960s" |
| 30-50% | Atmosphere and lighting | Medium-High | "golden hour sunlight streaming through floor-to-ceiling windows" |
| 50-70% | Mood, color palette, weather | Medium | "warm amber tones, dust motes in light shafts" |
| 70-85% | Camera, lens, composition | Medium-Low | "wide angle 24mm, symmetrical composition" |
| 85-100% | Fine detail and quality | Lower | "8K, architectural photography, sharp focus" |

### Depth & Scale Techniques

Environments fail when they look flat. Use these layering strategies:

- **Foreground elements** — Place objects close to camera: leaves, railings, columns, furniture edges, water droplets on glass. This creates depth and frames the scene.
- **Midground subject** — The primary architectural or environmental subject sits here. This is where the eye lands first.
- **Background recession** — Distant elements should be hazier, less saturated, and cooler in tone. This is atmospheric perspective.
- **Scale indicators** — Include humans, vehicles, trees, birds, or furniture to communicate the size of structures. A cathedral without a person looks like a model.
- **Leading lines** — Roads, corridors, rivers, colonnades, and staircases all pull the eye into the scene toward the focal point.
- **Vertical layering** — For tall environments (canyons, skyscrapers, forests), layer bottom to top: ground detail, mid-height structures, canopy or skyline, sky.

### Environment-Specific Model Selection

Use `search_models` to find the best model for each task:

| Task | What to Search For |
|------|--------------------|
| Architectural visualization | "image generation" — look for photorealistic models with strong structural accuracy |
| Fantasy landscapes | "image generation" — look for models with strong imagination and coherent composition |
| Interior design | "image generation" — look for models with material fidelity and lighting accuracy |
| Environment concept art | "image generation" — look for models with painterly quality |
| Panoramic expansion | "outpainting" or "expand" — seamlessly extend existing environments |
| Environment detail upscale | "upscale" — recover architectural texture at high resolution |
| Environment video flythrough | "text to video" — smooth camera movement through spaces |
| Still-to-video environment | "image to video" — animate a static environment render |
| Environment editing | "image editing" — change time of day, weather, season |
| 3D environment asset | "image to 3d" — convert environment renders to 3D meshes |

---

## 2. Architectural Styles

### Historical Periods

#### Ancient (3000 BCE - 500 CE)
**Keywords:** monumental stone, massive columns, hieroglyphics, temple complex, stepped pyramid, colonnade, marble, open courtyard, sacred geometry, bas-relief carvings
- **Egyptian:** colossal pylons, obelisks, hypostyle hall, papyrus columns, desert sandstone, gold leaf, painted walls, palm-lined avenue
- **Greek:** Doric/Ionic/Corinthian columns, pediment, entablature, white marble, open agora, amphitheatre, olive trees, blue Aegean sky
- **Roman:** arched aqueducts, coffered dome (Pantheon), triumphal arch, mosaic floor, thermal baths, colosseum, travertine, brick and concrete

#### Medieval (500 - 1500)
**Keywords:** fortress walls, crenellations, drawbridge, stained glass, flying buttresses, ribbed vault, rose window, cobblestone, timber frame
- **Romanesque:** thick walls, round arches, barrel vaults, fortress-like, small windows, stone masonry
- **Gothic:** pointed arches, flying buttresses, ribbed vaults, soaring nave, stained glass rose windows, gargoyles, spires, tracery
- **Half-timber:** exposed wooden beams, white plaster infill, steep gabled roofs, leaded windows, village square

#### Renaissance (1400 - 1600)
**Keywords:** symmetry, proportion, classical orders, dome, courtyard with arcades, rusticated stone, pediment, pilasters, loggia, palazzo
- Brunelleschi dome, Palladian villa, Florentine palazzo, formal gardens, geometric paving, marble inlay

#### Baroque (1600 - 1750)
**Keywords:** grandeur, dramatic curves, gilt ornament, trompe l'oeil ceiling, sweeping staircase, oval floor plan, cherubs, theatrical lighting
- Versailles Hall of Mirrors, Roman church facade, twisted columns, opulent ballroom, fresco ceiling

#### Victorian (1837 - 1901)
**Keywords:** ornate brickwork, wrought iron, bay windows, turrets, gingerbread trim, stained glass transom, slate roof, grand porch, terraced houses
- Queen Anne, Second Empire mansard roof, Gothic Revival, industrial warehouses, gaslight-era street

#### Art Nouveau (1890 - 1910)
**Keywords:** organic flowing curves, floral motifs, whiplash lines, stained glass, wrought iron, mosaic, ceramic tile, asymmetric facade
- Gaudi (Casa Batllo, Sagrada Familia), Horta, Guimard Metro entrances, Mucha decorative panels

#### Art Deco (1920 - 1940)
**Keywords:** geometric patterns, chevrons, sunburst motifs, stepped facade, gold and chrome, streamlined forms, ziggurat crown, terrazzo floors
- Chrysler Building, Miami Beach hotels, Egyptian Revival, Hollywood glamour, jazz-age opulence

#### Brutalist (1950 - 1975)
**Keywords:** raw exposed concrete (beton brut), massive geometric forms, repetitive modular units, angular overhangs, fortress-like, monolithic, no ornamentation
- Barbican Centre, Habitat 67, government buildings, university campuses, imposing stairwells

#### Modern / Contemporary
**Keywords:** glass curtain wall, steel frame, clean lines, open floor plan, cantilevered volumes, minimal ornamentation, flat roof, industrial materials
- Mies van der Rohe glass pavilion, Zaha Hadid curves, Tadao Ando concrete, Bjarke Ingels playful geometry, parametric facade

### Regional Architectural Traditions

#### Japanese
**Keywords:** zen garden, torii gate, tatami mat, shoji screen, sliding fusuma doors, engawa veranda, curved tiled roof, koi pond, bamboo fence, wabi-sabi imperfection, stone lantern, raked gravel, cedar wood, minimal ornament

#### Moroccan / North African
**Keywords:** riad courtyard, zellige mosaic tile, horseshoe arch, carved plaster (stucco), brass lanterns, fountain, bougainvillea, warm terracotta, blue city (Chefchaouen), souk market alley, keyhole doorway

#### Mediterranean
**Keywords:** whitewashed walls, terracotta roof tiles, bougainvillea climbing walls, blue shutters, courtyard with olive tree, stone staircase, wrought iron balcony, Santorini blue dome, cobblestone path

#### Scandinavian
**Keywords:** light wood (pine, birch, ash), floor-to-ceiling windows, white and grey palette, hygge warmth, wool textiles, nature integration, green roof, stave church, fjord-side cabin, minimalist clean lines

#### Middle Eastern
**Keywords:** muqarnas (honeycomb vault), arabesque geometric pattern, central courtyard with fountain, wind tower (badgir), mashrabiya screen, dome and minaret, desert fortress, palm oasis, calligraphy, blue and gold tile

#### Tropical / Southeast Asian
**Keywords:** elevated stilts, thatched palm roof, open-air pavilion, natural ventilation, carved teak, rice terrace, jungle canopy, Balinese temple gate, reflecting pool, orchid garden, rattan furniture

### Futuristic Architectural Visions

#### Solarpunk
**Keywords:** vertical garden facade, integrated solar panels, living green walls, organic curves merging with vegetation, community greenhouse, bamboo and recycled materials, rooftop farm, utopian, bright and hopeful, rewilded urban landscape

#### Cyberpunk
**Keywords:** neon signs reflecting on wet streets, holographic billboards, dense vertical mega-city, exposed pipes and cables, cramped market stalls, rain-slicked chrome, purple and cyan lighting, corporate megastructure towering above, street-level grime and chaos

#### Bio-organic / Biopunk
**Keywords:** grown architecture, living walls that breathe, bioluminescent corridors, mycelium-based structure, coral-like calcified building, symbiotic with nature, vine-threaded framework, translucent membrane walls, organic pod housing

#### Space Habitat
**Keywords:** rotating O'Neill cylinder, curved interior horizon, artificial gravity, modular hex panels, airlock doors, transparent dome viewing gallery, hydroponic garden ring, zero-gravity atrium, docking bay, navigation bridge

#### Underwater
**Keywords:** transparent geodesic dome, bioluminescent exterior lighting, coral reef integration, submarine docking tube, pressure-sealed glass tunnel, kelp farm, ocean floor foundation, bubble architecture, aquatic wildlife visible through walls

---

## 3. Interior Design

### Room Types & Prompt Vocabulary

**Residential:** living room, bedroom, master suite, kitchen, bathroom, home office, studio apartment, dining room, nursery, walk-in closet, reading nook, sunroom, conservatory, wine cellar, home theater
**Commercial:** hotel lobby, restaurant dining room, cocktail bar, co-working space, retail showroom, art gallery, museum hall, spa treatment room, rooftop lounge
**Specialty:** industrial loft, penthouse, yacht interior, private jet cabin, recording studio, photographer darkroom, artist atelier, greenhouse, treehouse interior, underground bunker

### Interior Styles with Keywords

#### Minimalist
Clean lines, monochromatic palette, negative space, hidden storage, single statement piece, natural light, uncluttered surfaces, matte finishes, recessed lighting, floating furniture

#### Scandinavian / Nordic
Light wood floors, sheepskin throws, pendant lights, white walls, pops of muted color, plants, functional beauty, cozy textiles, candles, ceramic vases, woven baskets

#### Industrial
Exposed brick walls, steel beams, concrete floor, Edison bulb pendant, factory windows, metal shelving, leather and raw wood, pipe fixtures, warehouse ceiling, distressed surfaces

#### Mid-Century Modern
Organic curves, tapered legs, teak and walnut, Eames chair, Nelson clock, sunken living room, shag rug, starburst mirror, avocado and mustard palette, low-slung sofa, built-in bookcase

#### Japanese Wabi-Sabi
Imperfect beauty, natural materials, handmade ceramics, low table, floor cushions, paper lantern, aged wood, linen curtains, single ikebana arrangement, muted earth tones, deliberate simplicity

#### Art Deco Interior
Geometric wallpaper, mirrored surfaces, velvet upholstery, gold and black palette, lacquered furniture, crystal chandelier, marble floor, fan motifs, chrome bar cart, statement fireplace

#### Bohemian / Maximalist
Layered textiles, patterned rugs, macrame, hanging plants, collected objects, global mix, warm tones, floor cushions, tapestries, vintage furniture, eclectic gallery wall

#### Luxury / Glam
Marble surfaces, crystal chandeliers, velvet in jewel tones, gold hardware, statement art, double-height ceiling, grand staircase, silk drapes, onyx countertop, faux fur, polished stone

### Material & Texture Keywords

**Wood:** white oak, dark walnut, honey maple, natural bamboo, reclaimed barn wood, charred shou sugi ban, whitewashed pine, raw cedar, ebony, teak, driftwood
**Stone:** Calacatta marble, black granite, grey slate, honey travertine, polished concrete, terrazzo with brass inlay, quartzite, sandstone, river stone, onyx
**Metal:** brushed brass, matte black iron, polished chrome, antique copper, raw steel, patinated bronze, gold leaf, hammered nickel, oxidized zinc
**Fabric:** plush velvet, raw linen, aged leather, heavy silk, chunky wool knit, boucle, cotton canvas, herringbone tweed, mohair, organic muslin
**Glass:** frosted glass partition, stained glass panel, etched crystal, smoked glass, hand-blown, beveled mirror, ribbed glass, colored glass vessel

### Lighting for Interiors
- **Natural light:** floor-to-ceiling windows, clerestory windows, skylights, light wells, sheer curtains diffusing sunlight
- **Ambient:** recessed cove lighting, indirect wall wash, pendant chandelier, paper lantern, backlit panels
- **Task:** desk lamp, under-cabinet strip, reading sconce, vanity mirror lights
- **Accent:** track lighting on art, LED strip under floating shelves, candles, fireplace glow, neon sign
- **Dramatic:** single overhead spot creating pool of light, silhouette against backlit window, candlelit room

---

## 4. Natural Landscapes

### Biomes with Visual Vocabulary

#### Forest
- **Tropical rainforest:** dense canopy, hanging vines, epiphytes, filtered green light, buttress roots, misty understory, exotic birds, giant ferns
- **Temperate deciduous:** oak and beech, dappled sunlight, leaf-covered forest floor, mushrooms on fallen logs, deer path, autumn colors
- **Boreal / Taiga:** spruce and pine, snow-laden branches, low grey sky, wolves, frozen stream, aurora overhead
- **Bamboo grove:** tall straight stalks, filtered green light, stone path, Japanese lantern, peaceful atmosphere
- **Enchanted old-growth:** moss-covered ancient trees, twisted roots, mist, shafts of golden light through canopy, fairy-tale atmosphere

#### Desert
- **Sand dune sea:** sweeping golden dunes, rippled sand texture, harsh midday shadows, single camel caravan, endless horizon
- **Rocky desert:** red sandstone formations, mesas, arches, sparse scrub, dramatic shadows, Monument Valley
- **Canyon:** layered geological strata, winding river below, dramatic depth, Antelope Canyon slot light, sandstone glow
- **Oasis:** palm cluster, clear pool, surrounded by sand, contrast of green and gold, mirage shimmer
- **Salt flat:** perfectly flat white surface, mirror reflection of sky, cracked hexagonal patterns, distant mountains

#### Mountain
- **Alpine meadow:** wildflowers, snow-capped peaks behind, clear stream, wooden cabin, blue sky with clouds
- **Volcanic:** dark basalt, lava flows (active or cooled), steam vents, crater lake, ash-grey sky, dramatic red glow
- **Cliff face:** vertical rock wall, climbers for scale, birds nesting, waterfall cascading, mist at base
- **Mountain pass:** winding path through peaks, prayer flags, thin atmosphere, panoramic valley view below

#### Water
- **Ocean coast:** crashing waves, sea stack, lighthouse, rocky shore, seabirds, tide pools, dramatic sky
- **Calm lake:** mirror reflection, surrounded by mountains or forest, wooden dock, canoe, morning mist
- **Waterfall:** cascading tiers, mossy rocks, rainbow in spray, lush green surroundings, long exposure silk water
- **Fjord:** steep cliff walls, narrow water channel, village at base, overcast Nordic sky, dramatic scale
- **Coral reef (underwater):** colorful coral formations, tropical fish schools, sea turtles, light rays from surface, blue gradient depth

#### Plains & Open Spaces
- **Savanna:** golden grass, scattered acacia trees, dramatic cloud formations, wildlife silhouettes, warm sunset
- **Prairie:** rolling grassland, wildflowers, distant thunderstorm, farmstead, wind-bent grass, big open sky
- **Tundra:** frozen flat expanse, lichen-covered rocks, distant snow mountains, overcast sky, sparse vegetation
- **Rice paddy:** flooded terraces reflecting sky, brilliant green shoots, workers in conical hats, mountain backdrop

#### Polar & Glacial
- **Glacier:** cracked blue ice, crevasses, ice cave interior, meltwater stream, massive scale
- **Iceberg:** floating sculptural ice, turquoise underwater portion, arctic ocean, seals, cold blue atmosphere
- **Aurora borealis:** green and purple curtains of light, reflected in still water, snow-covered landscape, stars, silhouetted trees
- **Frozen lake:** transparent ice showing depth, air bubbles frozen in layers, cracked patterns, surrounded by snow-covered trees

### Weather & Atmosphere Keywords

| Condition | Keywords for Prompts |
|-----------|---------------------|
| Clear sky | crystal clear blue sky, not a cloud, sharp shadows, high visibility, vivid colors |
| Overcast | flat grey sky, even diffused light, no harsh shadows, muted tones, low ceiling |
| Rain | wet reflective surfaces, puddles, rain streaks, dark grey sky, umbrellas, dripping |
| Heavy storm | dramatic thunderheads, lightning bolt, wind-bent trees, horizontal rain, dark sky |
| Snow | soft falling snowflakes, white blanket, muffled quiet, warm windows glowing, bare branches |
| Fog / Mist | low visibility, silhouettes emerging, mystery, muted colors, moisture on surfaces |
| Dust storm | orange-brown haze, reduced visibility, gritty texture, desert setting, dramatic |
| Volumetric light | god rays through clouds, light shafts through trees, visible beams in dust or mist |
| Heat haze | shimmering air above hot surface, distorted background, summer desert road |

### Time of Day — Visual Characteristics

| Time | Light Quality | Color Palette | Shadow Character | Best For |
|------|--------------|---------------|-----------------|----------|
| Pre-dawn | Cool blue-violet, faint | Deep blue, indigo, grey | Nearly absent | Mystery, stillness |
| Dawn | Soft warm pink-orange, low angle | Pink, peach, lavender, gold | Long, soft-edged | Hope, new beginnings |
| Golden hour (morning) | Warm directional, long shadows | Gold, amber, warm white | Long, warm-toned | Warmth, romance |
| Midday | Harsh overhead, high contrast | Neutral, vivid saturation | Short, directly below | Drama, clarity |
| Afternoon | Warm but less extreme | Warm neutral | Medium length | Everyday realism |
| Golden hour (evening) | Deep warm gold, raking light | Deep gold, orange, warm shadow | Very long, dramatic | Cinematic beauty |
| Blue hour | Cool ambient, no direct sun | Deep blue, purple, cyan | Minimal, soft | Melancholy, calm |
| Dusk / Twilight | Fading, mixed warm-cool | Orange horizon fading to blue zenith | Disappearing | Transition, nostalgia |
| Night (moonlit) | Cool silver-blue, dim | Blue-grey, silver, black | Faint, soft | Mystery, solitude |
| Night (artificial) | Warm pools of light, dark surrounds | Orange sodium, white LED, neon color | Hard-edged from point sources | Urban, noir, energy |

### Seasonal Visual Indicators

- **Spring:** cherry blossoms, new green buds, rain showers, wildflowers, soft light, lambs, renewal
- **Summer:** lush full green, harsh bright sunlight, heat haze, thunderstorms, long days, vibrant
- **Autumn:** red-orange-gold foliage, fallen leaves, low golden light, harvest, pumpkins, wood smoke, mist
- **Winter:** bare branches, snow cover, ice, short days, warm indoor glow, breath visible, stark contrast

---

## 5. Fantasy Worlds

### Classic Fantasy Environments

#### Enchanted Forest
**Keywords:** ancient trees with glowing runes, bioluminescent mushrooms, fairy lights, mist between trunks, twisted roots forming archways, crystal-clear stream, soft ethereal glow, mythical creatures, overgrown stone ruins, canopy of silver leaves
```
Enchanted ancient forest with massive gnarled oak trees, bioluminescent mushrooms casting soft blue-green glow along a mossy path, fireflies and fairy lights floating in the misty air, a crystal stream winding through twisted roots, shafts of ethereal golden light piercing the dense canopy above, fantasy concept art, 8K, detailed, magical atmosphere
```

#### Floating Islands
**Keywords:** landmasses suspended in sky, waterfalls cascading into void, chain bridges connecting islands, cloud level vegetation, ancient ruins on peaks, underside rock and roots visible, birds circling below, rainbow mist, impossible geography
```
Vast floating islands suspended in a golden sunset sky, waterfalls pouring from the edges into clouds below, ancient stone bridges and chain links connecting the islands, lush green vegetation on top with exposed rock and dangling roots underneath, birds circling between islands, fantasy landscape painting, epic scale, volumetric clouds, 8K
```

#### Underwater Kingdom
**Keywords:** coral palace, bioluminescent architecture, bubble domes, kelp forest corridors, pearl gates, mermaid architecture, light filtering from surface, fish schools as living decoration, shell and mother-of-pearl materials, pressure-proof glass
```
An underwater kingdom city built from living coral and iridescent shell, bioluminescent towers glowing in deep blue ocean, schools of colorful fish swimming between ornate archways, light rays filtering down from the distant surface, giant kelp forest framing the scene, pearl-encrusted gates, fantasy underwater concept art, cinematic lighting
```

#### Dragon's Lair
**Keywords:** vast cavern, treasure hoard (gold coins, gems, artifacts), volcanic vents with orange glow, massive scale (dragon skeleton or sleeping dragon), scorched stone, stalactites, eerie light, smoke and embers, carved ancient warnings
```
Enormous underground cavern serving as a dragon's lair, mountains of gold coins and jeweled artifacts, a massive sleeping dragon coiled around the treasure hoard, volcanic vents casting orange glow on the cave walls, stalactites hanging from the vast ceiling, smoke and embers drifting, epic fantasy, cinematic dramatic lighting, 8K detail
```

#### Elven City
**Keywords:** organic tree-integrated architecture, silver and white spires, living wood bridges, waterfalls within the city, delicate arches, leaf-shaped windows, bioluminescent lanterns, harmony with nature, ethereal atmosphere, ancient and elegant

#### Dwarven Halls
**Keywords:** massive underground carved stone halls, pillars with runic inscriptions, forge fires, rivers of molten metal, crystal veins in rock walls, geometric precision, vaulted ceilings lost in darkness, war machines, treasure vaults, echo of hammers

### Dark Fantasy

**Cursed ruins:** crumbling castle overgrown with dead thorns, green-black sky, spectral mist, shattered stained glass, crows circling, boneyard courtyard, oppressive atmosphere

**Haunted manor:** Victorian mansion, perpetual fog, dead garden, candlelit windows, cobwebs, tilted paintings, creaking staircase, full moon, twisted iron gate

**Dead forest:** skeletal leafless trees, ashen grey palette, no wildlife, toxic pools, fungal growths, perpetual twilight, path of bones, ancient curse corruption

**Underworld / Necropolis:** vast underground city of the dead, obsidian architecture, green ghostfire lanterns, floating spirits, endless catacombs, river of souls, oppressive darkness

### High Fantasy

**Crystal palace:** structure built entirely of prismatic crystal, rainbow light refractions, floating staircase, celestial-themed throne room, aurora visible through transparent walls

**Celestial realm:** clouds as solid ground, golden architecture, radiant beings, infinite sky, divine light source, cosmic backdrop of nebulae and stars, eternal dawn

**Elemental planes:** pure fire landscape (rivers of magma, flame geysers), water plane (infinite ocean, floating kelp islands), earth plane (crystal caverns, living stone), air plane (cloud towers, wind currents visible)

---

## 6. Sci-Fi Environments

### Space Environments

#### Space Station Interior
**Keywords:** curved corridors with handrails, LED strip lighting, modular panels, viewport windows showing stars, zero-gravity floating objects, control consoles with holographic displays, airlock doors, life support vents, clinical white and grey with colored accents
```
Interior corridor of a large orbital space station, curved walls with modular white panels, soft blue LED strip lighting, handrails along both sides, a large viewport window showing Earth below and stars, holographic directional signs, clean minimalist sci-fi design, cinematic lighting, 8K, concept art
```

#### Alien Planet Surface
**Keywords:** non-Earth sky color (green, purple, dual suns, ring system), unusual vegetation (crystalline, bioluminescent, oversized), strange geological formations, alien atmosphere (thick haze, floating particles), unfamiliar scale
```
Surface of an alien planet with a deep purple sky and two suns, crystalline tree-like formations growing from turquoise soil, bioluminescent floating spores in the thick atmosphere, massive ring system visible on the horizon, a lone astronaut in a white suit for scale, sci-fi concept art, awe-inspiring, Ridley Scott visual style, 8K
```

#### Spaceship Bridge
**Keywords:** command chair, wraparound viewport, holographic star map, crew stations, ambient lighting, hull curvature, tactical display, warp streaks through window, clean futuristic design

### Dystopian Environments

#### Ruined Megacity
**Keywords:** collapsed skyscrapers, overgrown with vegetation, flooded streets, rusted vehicles, faded billboards, nature reclaiming concrete, dramatic sky, survivors or emptiness, post-apocalyptic atmosphere
```
A ruined megacity decades after collapse, crumbling skyscrapers covered in creeping vines and moss, flooded streets reflecting an overcast sky, rusted cars half-submerged, a single deer standing on a highway overpass, nature reclaiming civilization, post-apocalyptic, photorealistic, golden hour light breaking through clouds, 8K cinematic
```

#### Wasteland
**Keywords:** barren cracked earth, dust storms, makeshift settlements from scrap, water scarcity, irradiated zones, harsh sunlight, survival aesthetic, Mad Max visual language

#### Underground Bunker
**Keywords:** concrete walls, flickering fluorescent lights, water pipes, canned food storage, bunk beds, radio equipment, sealed blast door, claustrophobic, stale air atmosphere

### Cyberpunk City

#### Street Level
```
Dense cyberpunk city street at night in the rain, neon signs in Japanese and English reflecting on wet asphalt, holographic billboards advertising corpo products, steam rising from grates, street food vendor with a glowing cart, crowded with diverse pedestrians under umbrellas, towering megastructures vanishing into low clouds above, purple and cyan color palette, Blade Runner atmosphere, cinematic, 8K
```

#### Megacorp Tower
```
Massive corporate headquarters tower rising above the clouds, sleek black glass and chrome facade, holographic company logo projected into the sky, flying vehicles docking at upper levels, cold authoritarian architecture, viewed from street level looking up, rain falling, neon-lit base contrasting with clean upper floors, cyberpunk, dramatic vertical composition, 8K
```

### Solarpunk / Utopian
```
Solarpunk city of the future, buildings covered in vertical gardens and integrated solar panels, elevated greenway walkways connecting organic-curved towers, community greenhouse domes on rooftops, clean light rail gliding silently, people cycling and walking among lush vegetation, warm golden afternoon sunlight, hopeful and utopian atmosphere, architectural visualization, 8K
```

---

## 7. Cinematic Environment Shots

### Establishing Shot Formula
An establishing shot sets the scene. Use wide framing, clear time-of-day indicators, and scale references.
```
Template: Extreme wide shot of [environment], [time of day], [weather], [atmosphere], [tiny human or vehicle for scale], [mood keyword], cinematic, [aspect ratio], [lens]
```
```
Extreme wide shot of a gothic cathedral perched on a cliff edge overlooking a stormy sea at dusk, lightning illuminating dark thunderclouds, crashing waves far below, a lone figure with a lantern approaching the entrance, ominous and awe-inspiring, cinematic, 2.39:1 anamorphic, 24mm wide angle
```

### Walkthrough / Fly-Through (Video)
For video generation, use `search_models` with "text to video" to find a high-quality video generation model:
```
Template: Smooth steadicam/drone [moving through/flying over] [environment], [passing by details], [revealing new areas], [lighting and atmosphere], cinematic, [speed]
```
```
Smooth steadicam walk-through of a grand abandoned Art Deco hotel lobby, camera gliding past a dust-covered marble reception desk, through a fallen chandelier of crystal and brass, past cracked mirrored columns reflecting golden light from a shattered skylight above, dust particles floating in the light shafts, slow and reverent pace, cinematic
```
```
Aerial drone shot slowly ascending over a dense bamboo forest, camera rising through the canopy to reveal a hidden ancient temple complex on a mountain peak, morning mist clinging to the valleys below, golden sunrise light, smooth continuous movement, cinematic, 4K
```

### Reveal Shot (Video)
Build anticipation then expose the full environment:
```
Camera pushing through dense fog that slowly clears to reveal a massive floating steampunk city, airships docked at multiple levels, smokestacks and brass spires catching the golden sunset light, the city stretching to the horizon, awe-inspiring reveal, cinematic, slow dramatic movement
```
```
A heavy stone door grinding open, camera moving forward through the widening gap into a vast underground dwarven hall, pillars carved with runes receding into the distance, forge fires glowing orange in the deep background, dramatic reveal, fantasy cinematic
```

---

## 8. Environment Consistency Across Multiple Generations

When building a world across multiple images or video sequences, maintain visual coherence:

1. **Lock your color palette** — Define 3-5 dominant colors and reference them in every prompt. Use HEX codes for precision: `"primary stone color #8B7D6B, accent teal #2A9D8F"`.
2. **Fix time of day and weather** — Every shot should state the same lighting conditions unless a deliberate time shift is intended.
3. **Repeat architectural style keywords** — Use the exact same style descriptors across prompts. If you said "brutalist concrete with geometric hex panels," repeat that phrase verbatim.
4. **Use image-to-video for animation** — Generate a strong still environment with an image generation model, then animate it with an image-to-video model (find both via `search_models`). This preserves the exact environment design.
5. **Expand with outpainting** — Use an outpainting/expand model (find via `search_models` with "outpainting" or "expand") to extend an environment image in any direction, revealing more of the same world.
6. **Edit for variation** — Use an image editing model (find via `search_models` with "image editing") to change time of day, add weather effects, or swap seasons while keeping the same base environment.
7. **Upscale for detail** — After establishing the composition, use an upscaling model (find via `search_models` with "upscale") to bring out architectural textures and fine environmental detail.

---

## 9. Complete Environment Prompts — Ready to Use

### 01: Modern Architecture Exterior
```
A contemporary art museum with a dramatic cantilevered glass and white concrete facade, set against a clear blue sky, reflecting pool in the foreground mirroring the structure, a few visitors walking along the minimalist landscaped entrance path, late afternoon golden hour sunlight casting long shadows, architectural photography, shot on Hasselblad, 24mm tilt-shift lens, sharp detail throughout, 8K
```
**Recommended aspect ratio:** 16:9. Use a high-quality image generation model (find via `search_models`).

### 02: Cozy Scandinavian Interior
```
A cozy Scandinavian living room in winter, large floor-to-ceiling window showing snow falling outside at dusk, warm interior lit by a crackling fireplace and soft pendant lamp, light oak floors, a cream boucle sofa with sheepskin throw and knit cushions, a low walnut coffee table with a ceramic mug and open book, a fiddle leaf fig plant in the corner, hygge atmosphere, interior design photography, warm natural tones, 8K
```
**Recommended aspect ratio:** 4:3 or 3:2. Use a high-quality image generation model (find via `search_models`).

### 03: Epic Fantasy Landscape
```
An epic fantasy landscape at sunrise, a massive ancient stone bridge spanning a deep misty gorge, a walled medieval city visible on distant cliffs, a winding river far below reflecting the pink and gold sky, a lone cloaked traveler on horseback crossing the bridge, eagles circling above, lush green mountains receding into atmospheric haze, fantasy concept art, cinematic composition, volumetric god rays, 8K, painted quality with photorealistic detail
```
**Recommended aspect ratio:** 21:9 or 16:9. Use an image generation model (find via `search_models`).

### 04: Cyberpunk City Night
```
A dense cyberpunk city intersection at night in heavy rain, neon signs in multiple languages reflecting on wet pavement in streaks of magenta, cyan, and gold, holographic advertisements flickering above, street vendors with glowing stalls, dense crowds with translucent umbrellas, towering dark megastructures disappearing into low clouds, steam and smoke rising, Blade Runner meets Ghost in the Shell, cinematic night photography, 8K
```
**Recommended aspect ratio:** 16:9. Use a high-quality image generation model.

### 05: Ancient Roman Forum
```
The Roman Forum at its peak, grand marble temples with Corinthian columns lining a paved stone avenue, senators in white togas walking among market stalls, a triumphal arch in the midground, the Colosseum visible in the distant background, clear Mediterranean sky, warm afternoon sunlight, historical reconstruction, photorealistic, cinematic wide shot, 8K
```
**Recommended aspect ratio:** 16:9. Use a high-quality image generation model.

### 06: Volcanic Alien Landscape
```
Surface of a volcanic alien world with a deep crimson sky and three moons visible, rivers of bright orange lava flowing between obsidian rock formations, bioluminescent alien flora growing near thermal vents, towering basalt columns in the background, toxic green atmosphere particles floating, an exploration rover in the distance for scale, sci-fi concept art, dramatic lighting from lava glow, 8K
```
**Recommended aspect ratio:** 16:9. Use an image generation model.

### 07: Underwater Coral City
```
A fantastical underwater city built into a massive coral reef, architecture of living coral and mother-of-pearl in iridescent colors, bioluminescent pathways connecting shell-shaped buildings, schools of tropical fish swimming between ornate arches, a giant sea turtle gliding past, light rays filtering from the ocean surface above creating caustic patterns, deep blue to turquoise color gradient, fantasy concept art, 8K
```
**Recommended aspect ratio:** 16:9. Use a high-quality image generation model.

### 08: Aerial Drone View — Autumn Forest
```
Aerial top-down drone shot of a winding river cutting through a dense autumn forest, trees in brilliant red, orange, gold, and remaining green, the river reflecting the overcast sky like liquid silver, a small wooden footbridge crossing at a narrow point, morning mist lingering in the valleys, nature photography, DJI Mavic, 8K, vivid autumn palette
```
**Recommended aspect ratio:** 1:1 or 4:3. Use a high-quality image generation model.

### 09: Neon-Lit Night Market
```
A bustling night market in a narrow alley, hundreds of colorful paper lanterns strung overhead creating warm orange glow, street food stalls with steam and smoke rising, neon signs in Chinese characters, wet cobblestones reflecting all the lights, crowded with people browsing, a cat sitting on a stall counter, shallow depth of field focusing on midground, night street photography, 35mm lens, 8K
```
**Recommended aspect ratio:** 9:16 (vertical) or 3:4. Use a high-quality image generation model.

### 10: Blizzard Mountain Pass
```
A treacherous mountain pass during a raging blizzard, visibility reduced to meters, a stone watchtower barely visible through the whiteout, snow-covered path with a single set of footprints being rapidly filled, wind-blown snow streaming horizontally, dark grey sky merging with white ground, a faint warm orange glow from the watchtower window, survival atmosphere, dramatic weather photography, 8K
```
**Recommended aspect ratio:** 16:9. Use a high-quality image generation model.

### 11: Solarpunk Rooftop Garden
```
A solarpunk rooftop community garden atop a curved green-architecture building, raised beds of vegetables and flowers, integrated solar panel canopy providing dappled shade, a small greenhouse dome, neighbors tending plants together, city skyline of other green buildings in the background, blue sky with white clouds, warm afternoon sunlight, hopeful utopian atmosphere, architectural visualization, 8K
```
**Recommended aspect ratio:** 16:9. Use a high-quality image generation model.

### 12: Gothic Cathedral Interior
```
Interior of a vast Gothic cathedral, soaring ribbed vaulted ceiling disappearing into shadow, massive stained glass rose window casting colored light across stone columns, rows of wooden pews, candlelight flickering in iron candelabras, incense smoke catching the colored light beams, a single figure kneeling in prayer in the distance, reverent and awe-inspiring atmosphere, architectural photography, 14mm ultra-wide lens, 8K
```
**Recommended aspect ratio:** 9:16 (vertical). Use a high-quality image generation model.

### 13: Japanese Zen Garden at Dawn
```
A traditional Japanese zen garden at dawn, raked white gravel in perfect concentric circles around moss-covered rocks, a single cherry blossom tree in full bloom dropping petals, a stone lantern, bamboo water feature (shishi-odoshi), low morning mist, distant pagoda silhouette, soft pink and gold dawn light, perfect tranquility, minimalist composition, 8K
```
**Recommended aspect ratio:** 3:2. Use a high-quality image generation model.

### 14: Post-Apocalyptic Overgrown City
```
A post-apocalyptic city street decades after humanity vanished, skyscrapers covered in thick ivy and trees growing from windows, cracked asphalt with grass and wildflowers pushing through, a rusted bus covered in moss, deer grazing where cars once drove, dramatic cumulus clouds in a clean blue sky, nature triumphant over civilization, golden hour sunlight, photorealistic, 8K
```
**Recommended aspect ratio:** 16:9. Use a high-quality image generation model.

### 15: Luxury Penthouse at Blue Hour
```
A luxury modern penthouse living room at blue hour, floor-to-ceiling windows with panoramic city skyline view of twinkling lights, interior lit by warm recessed lighting and a linear fireplace, Italian leather sectional sofa, marble coffee table, statement contemporary art on the wall, polished dark hardwood floors reflecting the city lights, sophisticated and intimate atmosphere, interior design photography, 8K
```
**Recommended aspect ratio:** 16:9. Use a high-quality image generation model.

### 16: Frozen Waterfall Ice Cave
```
Inside a massive ice cave behind a frozen waterfall, translucent blue and white ice walls with trapped air bubbles, icicles hanging from the ceiling, the frozen waterfall visible as a curtain of ice with faint daylight filtering through, a narrow ice path leading deeper into the cave, cool blue ambient light with occasional warm golden rays, pristine and otherworldly, nature photography, 8K
```
**Recommended aspect ratio:** 9:16. Use a high-quality image generation model.

### 17: Moroccan Riad Courtyard
```
A traditional Moroccan riad courtyard, ornate zellige mosaic tile floor in geometric blue and white pattern, a central marble fountain with trickling water, horseshoe arches with carved plaster walls, potted orange trees, brass lanterns casting warm patterned shadows, bougainvillea climbing the upper balcony railing, warm afternoon light from above, peaceful and exotic atmosphere, architectural photography, 8K
```
**Recommended aspect ratio:** 1:1. Use a high-quality image generation model.

### 18: Space Station Observation Deck
```
The observation deck of a massive orbital space station, a curved panoramic window showing Earth's blue marble and the Milky Way, the interior has sleek white and grey panels with blue accent lighting, floating plants in zero-gravity terrariums, a few crew members in jumpsuits gazing out the window, their reflections visible in the glass, serene and contemplative atmosphere, sci-fi concept art, cinematic, 8K
```
**Recommended aspect ratio:** 21:9. Use a high-quality image generation model.

### 19: Enchanted Library
```
A vast enchanted library inside a hollow tree, curved wooden shelves growing organically from the trunk walls reaching up many stories, floating glowing orbs providing warm light, books with faintly glowing spines, a spiral staircase of living wood, a reading nook with moss cushions, magical runes carved into bark, dust motes and tiny sparkles in the air, fantasy interior concept art, warm amber and green tones, 8K
```
**Recommended aspect ratio:** 9:16. Use an image generation model.

### 20: Desert Oasis at Sunset
```
A hidden desert oasis at sunset, a crystal-clear natural pool surrounded by date palms and lush vegetation, golden sand dunes in the background catching the last orange and pink light, the water perfectly reflecting the colorful sky and palm silhouettes, a small traditional tent camp with warm lantern light on the shore, camels resting nearby, serene and magical atmosphere, landscape photography, 8K
```
**Recommended aspect ratio:** 16:9. Use a high-quality image generation model.

### 21: Art Deco Hotel Lobby
```
A grand Art Deco hotel lobby from the 1930s restored to perfection, geometric patterned marble floor in black and gold, a massive sunburst chandelier, mirrored columns with chrome trim, velvet emerald green seating, a curved reception desk with inlaid wood, elevator doors with etched geometric motifs, warm ambient lighting, jazz-age glamour and sophistication, interior photography, 24mm, 8K
```
**Recommended aspect ratio:** 3:2. Use a high-quality image generation model.

### 22: Bioluminescent Cave System
```
A vast underground cave system illuminated entirely by bioluminescence, glowing blue and purple fungi on the walls, a subterranean lake reflecting the light like a starfield, crystal formations growing from floor and ceiling with internal glow, a natural stone bridge over the luminous water, no artificial light source, otherworldly and alien atmosphere, fantasy nature photography, 8K
```
**Recommended aspect ratio:** 16:9. Use an image generation model.

---

## 10. Environment Video Prompt Patterns

For video models (use `search_models` with "text to video" to find available options):

### Slow Pan Across Landscape
```
Slow cinematic pan from left to right across a vast mountain valley at golden hour, revealing snow-capped peaks, a winding river, and a small village with smoke rising from chimneys, atmospheric haze in the distance, warm golden light, birds flying across frame, peaceful and majestic
```

### Ascending Drone Reveal
```
Drone camera starting close to the ground in a field of lavender, slowly ascending vertically to reveal an enormous Provencal countryside landscape stretching to the horizon, a medieval hilltop village visible in the distance, patchwork fields of purple and gold, dramatic cloud formations, golden hour
```

### Walking Into an Interior
```
First-person steadicam walking through the front door of a cozy cabin, entering a warm firelit living room with wooden beams, past a kitchen with copper pots, through a hallway with family photos, to a window seat overlooking a snowy mountain landscape, warm interior contrasting with cold blue exterior, smooth continuous movement
```

### Day-to-Night Timelapse
```
Timelapse of a city skyline from afternoon to night, the sun setting behind skyscrapers, golden light transitioning to blue hour then to full darkness with building lights turning on one by one, car light trails appearing on highways below, stars appearing above, smooth 10-second timelapse, cinematic
```

### Weather Change
```
A peaceful countryside meadow under blue sky, clouds gradually building and darkening, wind picking up bending the tall grass, rain beginning to fall, a dramatic thunderstorm rolling in with visible lightning, then slowly clearing to reveal a rainbow, cinematic weather progression
```
