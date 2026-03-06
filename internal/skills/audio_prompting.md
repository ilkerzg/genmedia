# Audio Prompting — AI Music, Sound Effects & Voice Master Guide

Complete reference for AI audio generation on fal.ai. Covers music composition, sound effects, text-to-speech, voice design, speech processing, audio utilities, and audio-for-video workflows. Use `search_models` to discover available models for each category.

---

## Quick Reference

### Golden Rules

1. **Describe the audio, do not command it** — "Warm acoustic folk, 105 BPM, G major" beats "Create a folk song."
2. **Always specify exact BPM and key** — vague terms like "fast" or "slow" produce inconsistent results.
3. **Front-load genre and mood** — models weight early tokens more heavily.
4. **Keep prompts to 15–40 words** — more than 40 unfocused words degrades quality.
5. **Use one primary genre** — mixing more than 2–3 genre descriptors causes models to average them into mush.
6. **Use `composition_plan` for structure** — for songs >30s, define sections with timings (intro, verse, chorus, etc.).

### Prompt Patterns

- **Music**: `[Genre] + [BPM] BPM + [Key] + [Mood] + [Instruments] + [Production quality] + [Use context]`
  - Example: "Warm acoustic folk, 105 BPM, G major, fingerpicked guitar, gentle harmonica, carefree, indie film"
- **Sound Effects**: `[Perspective] [Source] [Action/Character] [Acoustic environment]`
  - Example: "Close-up, heavy wooden door, slow creak, reverberant stone corridor"
- **Voice**: `[Tone/Style] [Pace] [Delivery note]` — use punctuation (em-dashes, ellipses, commas) to shape delivery
  - Example: "Warm, authoritative narrator... measured pace. Clear enunciation."

### Audio Categories — What to Search For

| Need | Search for |
|------|-----------|
| Instrumental background music | text-to-music, music generation |
| Music with vocals/lyrics | music generation with lyrics |
| Sound effects | sound effects generation |
| Video auto-SFX (no prompt) | video sound effects |
| Realistic voice-over | text-to-speech |
| Multi-speaker dialogue | dialogue TTS, multi-speaker |
| Transcription | speech-to-text |
| Audio upscaling | audio super resolution |

---

## Music Generation

### Music Generation Models

Use `search_models` to find available music generation models. Look for models that support text-to-music generation with features like composition plans, lyrics, and instrumental-only modes.

**Key features to look for:** composition plan support, lyrics/vocal generation, instrumental mode, long-form generation (up to 10 minutes), multiple output formats.

**Common parameters across music models:**
  - `prompt` (string, optional): Style, mood, and genre description. Front-load genre and mood.
  - `composition_plan` (string, optional): Structured plan defining sections with timing and instrumentation. Use for songs longer than 30 seconds.
  - `music_length_ms` (integer, optional): Duration 3000–600000 ms (3 seconds to 10 minutes).
  - `output_format` (string, optional): `mp3_44100_128` (default), `mp3_44100_192`, `pcm_16000`, `pcm_22050`, `pcm_24000`, `pcm_44100`, `opus_48000_32`, `opus_48000_64`, `opus_48000_128`
  - `respect_sections_durations` (boolean, default true): Whether to follow `composition_plan` timing exactly
  - `force_instrumental` (boolean, default false): Set true for no vocals
- **Tips**:
  - Use `prompt` alone for simple one-shot generation. Use `composition_plan` (if supported) for multi-section songs with precise timing.
  - Keep `prompt` focused on style/genre/mood; put structure in `composition_plan`.
  - For background music, use instrumental-only mode if the model supports it.
  - Many models support broadcast/commercial grade output quality.
- **Example — Simple prompt**:
  ```
  prompt: "Warm acoustic folk, 105 BPM, G major, fingerpicked guitar, gentle harmonica, light tambourine, wanderlust, golden hour, carefree, indie film"
  force_instrumental: true
  music_length_ms: 180000
  ```
- **Example — Composition plan**:
  ```
  prompt: "Epic orchestral trailer music, D minor, 100 BPM, massive, powerful"
  composition_plan: "Section 1 (0:00-0:10): Low ominous strings, single timpani pulse, tension building.
  Section 2 (0:10-0:25): Staccato string ostinato, layered brass, growing intensity.
  Section 3 (0:25-0:40): Full orchestra — soaring brass fanfare, thundering percussion, powerful choir, triumphant climax.
  Section 4 (0:40-0:50): Sudden silence, then one final massive impact with long reverb tail."
  force_instrumental: true
  music_length_ms: 50000
  ```

### Other Music Model Capabilities

Use `search_models` to discover the full range of music generation models. Different models excel at different tasks:

- **Lyrics/vocal support**: Some models accept lyric input for synchronized vocal tracks
- **Audio-to-audio**: Style transfer, remixing, and transformation of existing audio
- **Audio inpainting**: Repair or replace sections within existing audio
- **Fast generation**: Some models generate 30 seconds in under 2 seconds
- **Royalty-free output**: Some models specifically produce royalty-free music for commercial use
- **Multi-language music**: Some models excel at specific language music (e.g., Chinese, English pop)
- **High-fidelity output**: Look for models offering 48kHz WAV output

---

### Core Music Prompt Formula

```
Genre/Subgenre + BPM + Key + Mood + Instruments + Vocal Style + Production Quality + Structure + Use Context
```

**Golden rules:**
- Write descriptions, not commands. "Ambient electronic, evolving pads, shimmer reverb, meditative, 60 BPM" — NOT "Create an ambient track."
- Sweet spot: 15–40 words for style descriptions. Longer compositions benefit from sectioned structure.
- Front-load the most important attributes (genre, mood) — models weight early tokens more heavily.
- Specify what you want, not what you do not want. "Instrumental only" beats "no vocals."
- Always use exact BPM ("85 BPM"), not vague terms ("slow tempo").
- Key signatures shape emotion: minor keys for melancholy/tension, major for bright/uplifting, modal for ambiguity.
- Layer specificity: "dusty lo-fi piano with tape wobble" is far better than "piano."
- One primary genre. Mixing more than 2–3 genre descriptors averages them into mush.

### Genre Keywords

**Electronic:**
- Ambient: ambient, dark ambient, space ambient, drone, ambient techno, new age
- Downtempo: downtempo, trip-hop, chillout, lounge, nu-jazz electronic
- Lo-Fi: lo-fi hip hop, lo-fi beats, bedroom pop, chillhop, jazzhop
- Synthwave: synthwave, retrowave, outrun, darksynth, dreamwave, vaporwave, future funk
- House: house, deep house, progressive house, tech house, acid house, tropical house, afro house, minimal house, disco house, soulful house
- Techno: techno, minimal techno, industrial techno, acid techno, dub techno, Detroit techno, melodic techno, hard techno
- Trance: trance, psytrance, progressive trance, uplifting trance, vocal trance, Goa trance
- Bass Music: dubstep, brostep, future bass, riddim, bass house, UK garage, 2-step, grime
- Drum and Bass: drum and bass, liquid DnB, neurofunk, jungle, breakcore
- EDM: EDM, big room, future house, electro house, hardstyle, happy hardcore
- Trap: trap, phonk, dark trap, melodic trap, cloud rap beats
- IDM: IDM, glitch, electronica, generative, granular

**Acoustic:**
- Classical: classical, baroque, romantic era, impressionist, neoclassical, minimalist classical, chamber music, solo piano, string quartet, symphonic
- Jazz: jazz, bebop, cool jazz, hard bop, modal jazz, free jazz, fusion, smooth jazz, acid jazz, gypsy jazz, swing, big band
- Latin Jazz: bossa nova, samba jazz, Latin jazz, Afro-Cuban jazz
- Folk: folk, indie folk, Celtic folk, Appalachian folk, Nordic folk, singer-songwriter, acoustic
- Blues: blues, delta blues, Chicago blues, electric blues, blues rock, swamp blues
- Country: country, outlaw country, country pop, bluegrass, Americana, honky-tonk, Western swing, alt-country, Nashville sound

**Contemporary:**
- Pop: pop, synth-pop, electropop, indie pop, dream pop, art pop, K-pop, J-pop, power pop, chamber pop, hyperpop
- Rock: rock, indie rock, post-rock, shoegaze, alternative rock, psychedelic rock, progressive rock, garage rock, surf rock, grunge, punk, post-punk, math rock, stoner rock, classic rock
- Hip Hop: hip hop, boom bap, conscious hip hop, old school hip hop, West Coast, East Coast, Southern hip hop, abstract hip hop
- R&B/Soul: R&B, neo-soul, contemporary R&B, classic soul, Motown, quiet storm, PBR&B, alternative R&B
- Funk: funk, P-funk, electro-funk, funk rock, disco funk, modern funk
- Metal: metal, heavy metal, thrash metal, black metal, death metal, doom metal, progressive metal, post-metal, djent, metalcore, symphonic metal

**Cinematic:**
- Film score, epic orchestral, trailer music, underscore, dramatic orchestral, adventure score, romantic score, suspense score, horror score, action score, fantasy score
- Retro/Game: 8-bit, chiptune, 16-bit, video game soundtrack, pixel art music, retro game, arcade

**World:**
- African: Afrobeat, Afropop, highlife, soukous, Ethio-jazz, amapiano, kwaito, juju, mbalax
- Latin: reggaeton, salsa, cumbia, bachata, merengue, tango, mariachi, norteno, Latin pop, dembow, Latin trap, corrido, bossa nova, samba, MPB
- Caribbean: reggae, dub, dancehall, ska, soca, calypso, zouk
- Middle Eastern: Arabic, Andalusian, oud music, Persian classical, Turkish, Sufi, belly dance
- South Asian: Indian classical, Hindustani, Carnatic, Bollywood, Bhangra, qawwali, Rajasthani folk, tabla solo, sitar raga
- East Asian: Japanese traditional (koto, shamisen, shakuhachi), Chinese classical (guzheng, erhu, pipa), Korean traditional (gayageum, pansori), K-pop, J-pop, C-pop, anime soundtrack
- Other: Polynesian, Aboriginal didgeridoo, Balkan brass, Klezmer, Flamenco, Fado, Celtic, Nordic/Viking folk

### Mood Keywords

**Positive/Energetic:** uplifting, euphoric, triumphant, joyful, playful, celebratory, warm, hopeful, bright, exuberant, carefree, vibrant, radiant, infectious, sunny, blissful, ecstatic, animated, spirited, jubilant

**Calm/Peaceful:** peaceful, serene, meditative, contemplative, tranquil, ethereal, dreamy, spacious, floating, gentle, soothing, lullaby, zen, still, delicate, hushed, ambient, relaxing, soft, introspective

**Dark/Intense:** brooding, ominous, sinister, menacing, aggressive, intense, gritty, dystopian, heavy, brutal, savage, punishing, oppressive, bleak, nihilistic, abrasive, chaotic, relentless, visceral, furious

**Dramatic/Emotional:** melancholic, bittersweet, nostalgic, haunting, epic, powerful, sweeping, majestic, cinematic, grandiose, soaring, stirring, poignant, wistful, tender, heartbreaking, yearning, emotional, passionate, longing

**Mysterious/Atmospheric:** eerie, otherworldly, cosmic, surreal, hypnotic, tense, suspenseful, enigmatic, spectral, misty, foggy, twilight, liminal, uncanny, cryptic, arcane, ancient, primordial, alien, nebulous

### Instruments

**Strings:**
- Guitar: acoustic guitar, nylon-string guitar, steel-string guitar, 12-string guitar, electric guitar, clean electric guitar, overdriven guitar, distorted guitar, fingerpicked guitar, palm-muted guitar, slide guitar, pedal steel guitar
- Orchestral: violin, solo violin, viola, cello, solo cello, double bass, string quartet, string section, pizzicato strings, staccato strings, legato strings, tremolo strings
- Other: harp, Celtic harp, banjo, clawhammer banjo, ukulele, mandolin, sitar, oud, koto, guzheng, erhu, lute
- **Adjectives**: shimmering, soaring, weeping, warm, rich, lush, delicate, resonant, plucked, bowed, sustained

**Keys/Keyboards:**
- Piano: grand piano, upright piano, honky-tonk piano, prepared piano, solo piano, piano ballad
- Electric: Rhodes piano, Wurlitzer, Clavinet, Hammond B3 organ, Vox organ, church organ, pipe organ
- Synth: analog synth, Moog bass, Juno pad, Prophet lead, FM synth, wavetable synth, granular synth, modular synth, synth pad, synth lead, synth arpeggio, poly synth, supersaw
- Other: harpsichord, celesta, music box, kalimba, accordion, harmonium, toy piano, marimba, xylophone, vibraphone, glockenspiel
- **Adjectives**: lush, glassy, warm, bright, dark, evolving, pulsing, shimmering, detuned, filtered, resonant

**Percussion:**
- Acoustic: drum kit, brushed drums, jazz drums, rock drums, double kick, snare roll, hi-hat pattern, ride cymbal, crash cymbal, floor tom, timpani, orchestral percussion
- World: taiko drums, djembe, congas, bongos, cajon, tabla, frame drum, bodhran, steel drum, gamelan, talking drum, surdo, pandeiro
- Electronic: drum machine, TR-808, TR-909, LinnDrum, breakbeat, chopped break, boom bap drums, trap hi-hats, glitch percussion, sidechain kick
- Pitched: vibraphone, marimba, xylophone, glockenspiel, tubular bells, chimes, hand bells
- **Adjectives**: punchy, tight, loose, swinging, syncopated, driving, thundering, crisp, snappy, booming, rattling, rolling

**Wind/Brass:**
- Woodwinds: flute, piccolo, alto flute, clarinet, bass clarinet, oboe, English horn, bassoon, recorder, pan flute, bamboo flute, shakuhachi, duduk, bansuri
- Brass: trumpet, muted trumpet, flugelhorn, French horn, trombone, tuba, brass section, brass fanfare, brass stabs
- Reeds: saxophone (soprano/alto/tenor/baritone), harmonica, blues harp, accordion, bagpipes, melodica
- **Adjectives**: soaring, wailing, breathy, mellow, piercing, triumphant, muted, growling, smooth, jazzy

**Electronic/Production:**
- Bass: 808 bass, sub-bass, synth bass, acid bass, Reese bass, wobble bass, plucked bass, FM bass, distorted bass
- Textures: ambient pad, evolving texture, noise sweep, vinyl crackle, tape hiss, rain texture, field recording, found sound
- FX: riser, build-up, drop, impact, reverse cymbal, reverse reverb, white noise sweep, vocal chop, vocal stab, glitch effect, stutter edit, sidechain pump
- **Adjectives**: pulsing, throbbing, swelling, filtered, saturated, distorted, clean, airy, wide, deep, crushing

### BPM by Genre Table

| Genre | BPM Range | Sweet Spot | Genre | BPM Range | Sweet Spot |
|-------|-----------|------------|-------|-----------|------------|
| Ambient | 50–80 | 65 | Pop | 100–130 | 120 |
| Lo-Fi Hip Hop | 70–95 | 85 | Rock | 100–140 | 120 |
| Downtempo | 70–100 | 90 | Indie Rock | 110–140 | 125 |
| Jazz | 80–200 | 120 | House | 118–130 | 124 |
| R&B / Soul | 60–100 | 80 | Techno | 120–150 | 130 |
| Reggae / Dub | 60–90 | 75 | Trance | 125–150 | 138 |
| Hip Hop | 70–100 | 90 | Drum & Bass | 160–180 | 174 |
| Trap | 130–170 | 145 | Dubstep | 138–142 | 140 |
| Synthwave | 80–120 | 100 | Hardstyle | 150–160 | 155 |
| Bossa Nova | 100–130 | 115 | Afrobeat | 95–130 | 110 |
| Country | 100–140 | 115 | Reggaeton | 85–100 | 95 |
| Funk | 95–115 | 105 | Salsa | 150–250 | 180 |
| Blues | 60–120 | 85 | Film Score | 60–160 | variable |
| Metal | 100–200 | 140 | March/Military | 100–140 | 120 |
| Waltz | 80–120 | 100 (3/4) | Polka | 100–140 | 120 (2/4) |

### Production Quality Keywords

- **Analog Warmth**: warm analog, tape saturation, vinyl warmth, tube compression, retro mastering, analog summing, reel-to-reel, cassette tape, lo-fi warmth, worn-out vinyl
- **Modern Polish**: polished, radio-ready, mastered, compressed, punchy, crisp, bright, clean, tight, commercial-grade, broadcast-quality, streaming-optimized
- **Spatial/Reverb**: intimate room, concert hall reverb, plate reverb, spring reverb, wide stereo, 3D audio, binaural, cathedral echo, large warehouse, studio dry, close-miked, distant miking
- **Mix Character**: wide cinematic mix, punchy kick, crisp snare, deep sub-bass, airy highs, mid-forward, scooped mids, thick low-end, sparkling highs, balanced mix, in-your-face, lo-fi distortion
- **Era/Aesthetic**: 60s psychedelic, 70s disco production, 80s gated reverb, 90s boom bap grit, 2000s compressed pop, modern minimalist, future production, vintage recording, authentic live sound

### Musical Key Moods

| Key | Mood / Character |
|-----|-----------------|
| C major | Bright, clean, innocent, pure, happy |
| C minor | Dark, dramatic, passionate, stormy |
| D major | Triumphant, victorious, joyful, energetic |
| D minor | Serious, dark, melancholic, brooding, solemn |
| E major | Powerful, bright, cheerful, bold |
| E minor | Tender, sad, introspective, wistful |
| F major | Pastoral, warm, peaceful, gentle, graceful |
| F minor | Funereal, deeply sad, anguished, somber |
| G major | Simple, happy, rustic, marching, folk-like |
| G minor | Tragic, serious, discontent, moody |
| A major | Bright, confident, optimistic, sparkling |
| A minor | Neutral melancholy, tender, elegant, reflective |
| Bb major | Noble, grand, jovial, celebratory |
| Bb minor | Obscure, desolate, bleak, hopeless |
| Eb major | Majestic, heroic, bold, romantic |
| Ab major | Soft, dreamy, ethereal, warm velvet |

### Song Structure Tags

**Composition plan** (models that support structured sections):
```
composition_plan: "Section 1 (0:00-0:15): Soft piano intro, building anticipation.
Section 2 (0:15-0:45): Full band enters, driving rhythm, main melody.
Section 3 (0:45-1:15): Chorus — big, anthemic, layered vocals and synths.
Section 4 (1:15-1:30): Breakdown — stripped back to bass and drums.
Section 5 (1:30-2:00): Final chorus — everything at full power, triumphant ending."
```

**Natural language descriptions** (models that accept free-form structure prompts):
```
"Starts with a quiet atmospheric intro, sparse piano notes over ambient pad.
At 15 seconds the beat drops in with a driving four-on-the-floor kick and bassline.
Builds through the middle section with layered synths and arpeggios.
Climaxes at 1:30 with full orchestral swell. Fades out gently over the last 10 seconds."
```

**General structure terminology:**
```
Intro, Verse, Pre-Chorus, Chorus, Post-Chorus, Bridge, Breakdown,
Build-Up, Drop, Instrumental Break, Solo (Guitar Solo, Piano Solo, Sax Solo),
Outro, Fade Out, Cold Ending, Tag, Coda, Interlude, Vamp, Refrain
```

---

## Sound Effects

### Sound Effects Models

Use `search_models` to find available sound effects models. Search for "sound effects", "video to audio", or "SFX generation".

**Types of sound effects models available:**

- **Text-to-SFX**: Generate sound effects from text descriptions. Look for models supporting high-fidelity output (up to 48kHz), foley, ambient soundscapes, game audio, and UI sounds.
- **Video-to-audio**: Analyze video content and generate synchronized audio automatically, with no text prompt needed. Great for adding contextual sounds to video.
- **Video SFX sync**: Some models accept video input and return the video with a new soundtrack or matching sound effects.
- **General SFX**: Professional-grade sound effects for films, games, and multimedia. Some models are trained on millions of sound effects for realistic foley and environmental sounds.

### SFX Prompting Rules

1. **Specify material**: "wooden door" not just "door"; "glass bottle on marble floor" not just "bottle breaking"
2. **Describe size and weight**: "small tin can rolling" vs "massive steel barrel rolling"
3. **Set the environment/space**: "in a large cathedral" vs "in a small tiled bathroom" vs "outdoor open field"
4. **Include distance**: "close-up foley" vs "heard from 50 meters away" vs "distant, echoing"
5. **Describe temporal evolution**: "starts quietly, builds to a crash, then reverberates for 3 seconds"
6. **Quality markers**: add "professionally recorded," "foley quality," "high fidelity," "48kHz studio quality"
7. **Duration**: always specify "2 seconds," "10 seconds," "30 seconds, seamless loop"
8. **Layering hint**: for complex scenes, generate individual elements and layer them

### SFX Examples by Category

**Nature:**
```
"Steady heavy rain on a corrugated tin roof with occasional distant thunder rumbles, close-up recording, high fidelity, 30 seconds, seamless loop"

"Early morning temperate forest ambience, diverse birdsong chorus, gentle stream babbling over rocks, soft breeze through leaves, field recording, dawn, 30 seconds"

"Ocean waves crashing rhythmically on a rocky shore, sea spray, continuous cycle, wide stereo, coastal wind, field recording quality, 30 seconds seamless loop"

"Thunderstorm approaching across open plains, first distant rumbles, then closer cracks of lightning with immediate thunder, howling wind, heavy rain onset, dramatic, 30 seconds"
```

**Urban:**
```
"Busy outdoor European cafe, overlapping crowd conversation, clinking wine glasses, espresso machine steaming, distant church bells, afternoon ambience, stereo, 30 seconds"

"Subway train arriving at underground platform, distant rumble growing to roar, pneumatic brakes hissing, doors opening with chime, crowd murmur, echo, 10 seconds"

"City traffic intersection, honking taxis, bus air brakes, pedestrian crosswalk beeping, distant sirens, urban daytime, New York ambience, 20 seconds"

"Rainy city street at night, car tires on wet asphalt, neon sign buzzing, umbrella opening, footsteps on wet pavement, noir atmosphere, 15 seconds"
```

**Mechanical:**
```
"Sports car engine cold start, ignition crank, deep V8 rumble idling, revving through gears to high RPM, exhaust crackle on deceleration, close-up, 10 seconds"

"Intricate clockwork mechanism, tiny gears meshing, springs unwinding, escapement ticking precisely, steampunk aesthetic, foley quality, 5 seconds"

"Heavy industrial vault door opening slowly, massive steel hinges groaning, hydraulic pistons releasing, metallic echoing, dust settling, 8 seconds"

"Factory assembly line, robotic arm servos whirring, spot welding sparks, conveyor belt rolling, pneumatic press stamp, industrial rhythm, 15 seconds"
```

**Fantasy/Sci-Fi:**
```
"Magical spell casting, gathering energy with rising crystalline tones, sparkling chime particles swirling, explosive release with deep reverberant whoosh, fantasy RPG, 4 seconds"

"Ancient dragon roar, deep guttural growl building through chest to thunderous bellow with fire-breath crackle overtone, cave echo, epic scale, 5 seconds"

"Sci-fi laser blaster single shot, bright energy charge-up, sharp plasma discharge, ionized air sizzle, echo in metal corridor, futuristic, 1.5 seconds"

"Spaceship warp drive engaging, low-frequency power building, electromagnetic hum intensifying, reality-tearing whoosh with Doppler stretch, deep space, 6 seconds"
```

**UI/Interface:**
```
"Pleasant notification chime, two ascending bell tones (C5 to E5), bright, digital, warm, clean, 0.5 seconds"

"Tactile button click, crisp mechanical micro-switch feel, short pop with subtle bottom-out, satisfying, 0.1 seconds"

"Level-up achievement fanfare, ascending brass and strings, mini-orchestral flourish, triumphant, rewarding, game UI, 3 seconds"

"Error alert — two low descending tones, gentle but attention-getting, not harsh, minimal, 0.6 seconds"
```

**Horror:**
```
"Creaking wooden floorboards under slow, heavy footsteps in abandoned house, dust settling, absolute silence between steps, close-up foley, dread, 8 seconds"

"Distant whispered voices layered and overlapping, unintelligible, coming from all directions, wet reverb, paranormal, deeply unsettling, 10 seconds"

"Rusty metal gate swinging open in wind, prolonged screech, chain rattling, desolate graveyard ambience, fog, 5 seconds"

"Something large breathing in darkness, wet inhale and exhale, guttural undertone, cave acoustics, creature unknown, horror, 6 seconds"
```

**Comedy:**
```
"Classic cartoon slip and fall, ascending slide whistle, splat impact, spring boing, exaggerated, Looney Tunes style, 2 seconds"

"Comedic dramatic sting — dun dun DUUUN — low brass with timpani, over-the-top melodrama, sitcom reveal, 2 seconds"

"Record scratch abruptly stopping music, vinyl needle scrape, awkward silence, freeze-frame moment, 1 second"

"Wah-wah-wah sad trombone, three descending notes, comedic failure, game show buzzer energy, 2 seconds"
```

**Sports:**
```
"Baseball bat hitting ball, sharp crack of wood, ball flight whoosh, distant crowd erupting in cheers, stadium ambience, 4 seconds"

"Boxing ring bell, three sharp dings, echoing in arena, crowd anticipation murmur, 3 seconds"

"Swimming pool splash, competitive dive entry, clean water impact, underwater bubbles, surface settling, Olympic quality, 3 seconds"

"Soccer goal celebration, net rustling, crowd erupting, stadium horn blast, chanting, euphoric, 8 seconds"
```

**Foley/Everyday:**
```
"Ceramic coffee mug placed on wooden desk, solid thunk with brief ring, close-up, studio quality, 0.5 seconds"

"Paper pages turning in a hardcover book, crisp rustle, intimate, library quiet, ASMR quality, 3 seconds"

"Glass of ice water poured, ice clinking against glass, liquid filling, final settling fizz, kitchen, close-up, 5 seconds"

"Zipper on leather jacket, full zip up, metal teeth, smooth pull, close-up, 2 seconds"
```

---

## Voice & Speech (Text-to-Speech)

### TTS Models

Use `search_models` to find available text-to-speech models. Search for "text to speech", "TTS", "voice synthesis", or "dialogue".

**Types of TTS capabilities available:**

- **High-quality single narrator**: Best for narration, commercials, character voices, audiobooks, professional voice-over. Look for models supporting inline emotion tags (`[laughs]`, `[whispers]`, `[sighs]`, `[gasps]`, `[clears throat]`, `[crying]`, `[sniffles]`) and voice parameter controls (`stability`, `similarity_boost`, `voice_id`).

- **Fast/turbo variants**: Optimized for speed with good quality. Best for chatbots, live applications, rapid iteration, prototyping.

- **Multilingual**: 20-30+ language support with natural pronunciation and code-switching within a single generation. Best for non-English content, localization, multilingual projects.

- **Text-to-dialogue**: Multi-speaker conversation generation from script. Natural turn-taking, interruptions, rhythm. Assign different voices to different speakers. Best for podcasts, audio dramas, interview simulations, scripted dialogue scenes.

- **Multi-speaker TTS**: Native multi-speaker support with natural turn-taking and voice presets. Best for podcasts, natural conversation, long-form generation. Some models offer multiple size variants (larger for production quality, smaller for prototyping).

- **Two-speaker dialogue**: Realistic two-speaker dialogue with emotion and natural nonverbals. Use speaker prefixes like `[S1]` and `[S2]`. Emotion/nonverbal tags: `(laughs)`, `(sighs)`, `(clears throat)`, `(gasps)`, `(cries)`, `(groans)`, `(yawns)`.

- **Emotionally expressive TTS**: Models that understand emotional context from surrounding text. May support emotion tags like `<laugh>`, `<sigh>`, `<gasp>`, `<cry>`, `<groan>`, `<yawn>`, `<cough>`. Some models infer emotion from the text content itself without explicit tags.

- **Voice cloning**: Clone a voice from 10-30 seconds of clean audio. Zero-shot cloning available on some models.

- **Voice design**: Create a custom voice from a text description (no audio sample needed). Example: "warm female alto, professional, slightly breathy, aged 35".

- **Speech-to-speech**: Voice transformation — convert one voice to another while preserving speech content, pacing, and emotion.

- **Chinese/Mandarin-optimized**: Some models have excellent Chinese and English support with dedicated HD and turbo variants.

**Common TTS parameters:**
  - `voice_id`: Pre-built or cloned voice identifier
  - `stability`: Lower values = more expressive delivery
  - `similarity_boost`: How closely to match the target voice
  - `speed`, `vol`, `pitch`: Granular voice controls (model-dependent)
  - Temperature control (higher = more creative), exaggeration factor (0.0-1.0)

### Voice Characteristics

**Pitch:**
- Deep: deep bass, baritone, low resonant, rumbling, booming, subterranean
- Mid: warm tenor, mellow, medium, balanced, clear mid-range
- High: bright soprano, airy, light, crystal clear, ringing, silvery

**Texture:**
- Warm: warm, velvety, honeyed, buttery, smooth, rich, golden
- Crisp: crisp, articulate, precise, clean, sharp, clear-cut, polished
- Rough: gravelly, husky, raspy, scratchy, gritty, sandpaper, whiskey-soaked, weathered
- Smooth: silky, flowing, liquid, effortless, satin, creamy

**Age/Character:**
- Young: youthful, fresh, bright, innocent, energetic, adolescent, child-like
- Mature: experienced, seasoned, wise, authoritative, distinguished, aged
- Character types: authoritative, commanding, friendly, approachable, soothing, trustworthy, playful, mischievous, stern, gentle, intense, detached, quirky

### Speaking Styles

| Style | Keywords | Best For |
|-------|----------|----------|
| Conversational | natural, relaxed, informal, chatty, easygoing, off-the-cuff | Podcasts, vlogs, casual content |
| Narration | measured, clear, storytelling, even pace, engaging, literary | Audiobooks, documentaries, explainers |
| Announcement | bold, projected, dynamic, attention-grabbing, authoritative | Trailers, ads, event promos |
| Newscast | professional, neutral, clear diction, steady pace, informative | News, reports, corporate |
| Meditation | slow, gentle, breathy, hypnotic, calming, spacious pauses | Wellness apps, guided meditation |
| ASMR | whispered, close-miked, intimate, soft, delicate, mouth sounds | ASMR content, intimate narration |
| Character Voice | exaggerated, animated, distinctive, theatrical, unique | Animation, games, storytelling |
| Commercial | upbeat, confident, persuasive, warm, trustworthy, polished | Ads, promos, brand content |
| Academic | measured, authoritative, explanatory, patient, didactic | E-learning, courses, tutorials |
| Sportscaster | energetic, fast-paced, excited, building tension, exclamatory | Sports content, hype reels |
| Bedtime Story | warm, gentle, rhythmic, soothing, slower tempo, tender | Children's content, sleep stories |

### Emotion Tags — Common Patterns

Different TTS models support different emotion tag formats. Check the model's documentation for supported tags.

**Square bracket tags** (common in many TTS models):
- `[laughs]`, `[whispers]`, `[sighs]`, `[gasps]`, `[clears throat]`, `[crying]`, `[sniffles]`
- Descriptive text also works: "said excitedly", "whispered nervously"
- Voice settings: stability (0.0-1.0), similarity_boost, style, use_speaker_boost
- Lower stability = more expressive/variable; higher stability = more consistent/flat

**Parenthetical tags** (common in dialogue TTS models):
- `(laughs)`, `(sighs)`, `(clears throat)`, `(gasps)`, `(cries)`, `(groans)`, `(yawns)`
- Multi-speaker: `[S1]` and `[S2]` prefixes

**Angle bracket tags** (common in expressive TTS models):
- `<laugh>`, `<sigh>`, `<gasp>`, `<cry>`, `<groan>`, `<yawn>`, `<cough>`
- Some models read emotional cues from surrounding text automatically

**Implicit emotion** (some models):
- No explicit tags needed — infers emotion from text content

**Voice parameter controls** (some models):
- `voice_setting` with `speed`, `vol`, `pitch` controls
- Voice IDs for pre-built voices, or use voice cloning/design

### Punctuation as Delivery Control

| Punctuation | Effect |
|-------------|--------|
| Period `.` | Natural pause, falling intonation, finality |
| Comma `,` | Brief pause, continuation expected |
| Ellipsis `...` | Thoughtful pause, trailing off, hesitation |
| Question mark `?` | Rising intonation |
| Exclamation `!` | Increased energy, emphasis |
| ALL CAPS | Stress/emphasis on word |
| Em-dash `—` | Abrupt pause, interruption, dramatic break |
| Parentheses `()` | Softer aside, lowered volume |
| Semicolon `;` | Medium pause, connected thoughts |
| Double line break | Longer pause between paragraphs |

---

## Audio Utilities

### Speech-to-Text

Use `search_models` to find speech-to-text models. Look for features like word-level timestamps, speaker identification, and multi-language support.

### Audio Processing

Use `search_models` to find audio processing models. Available capabilities include:

- **Audio isolation**: Separate vocals from music, clean noisy audio, extract voice from background
- **Voice changer**: Convert one voice to another while preserving speech content, pacing, and emotion
- **Dubbing**: Translate and re-voice video/audio content, preserving speaker characteristics across languages

### Audio Enhancement

Use `search_models` to find audio enhancement/super-resolution models. Capabilities include upscaling speech audio to high-quality 48 kHz output.

---

## Audio for Video

### Video Models with Native Audio

Some video generation models support built-in audio generation. Use `search_models` to find video models with audio capabilities. Common patterns:

- **`generate_audio` parameter**: Enables native audio matched to video content. Some models also support `voice_ids` for dialogue scenes.
- **Native audio built in**: Some models generate audio automatically as part of video generation.
- **`audio_url` parameter**: Provide custom audio and the video syncs to it. Generate audio first, then pass it to the video model.

### Background Music Templates by Genre

**Product Launch / Commercial:**
`Uplifting indie pop, 120 BPM, bright acoustic guitar, warm female vocals humming, cheerful claps, commercial, polished, radio-ready`

**Corporate / B2B:**
`Light positive corporate, 110 BPM, soft piano melody, gentle clean guitar, subtle warm strings, inspiring, instrumental, professional`

**Tutorial / How-To:**
`Minimal lo-fi beats, 85 BPM, soft Rhodes keys, muted kick, unobtrusive background music, no vocals, seamless loop`

**Fashion / Luxury:**
`Minimal deep house, 122 BPM, sultry sub-bassline, dark atmospheric pads, crisp hi-hats, sophisticated, runway, high-end`

**Horror / Thriller:**
`Dark ambient drone, dissonant string cluster, creeping tension, 60 BPM, reverse reverb swells, unsettling whispers, found footage`

**Social Media / Reels:**
`Catchy energetic, 128 BPM, trending pop beat, bright supersaw synths, punchy sidechained kick, hook-driven, attention-grabbing, 15 seconds`

**Travel / Vlog:**
`Warm acoustic folk, 105 BPM, fingerpicked guitar, light tambourine, gentle whistling melody, wanderlust, golden hour, carefree`

**Fitness / Sports:**
`High-energy EDM, 140 BPM, heavy bass drop, aggressive synth leads, driving four-on-the-floor, motivational, powerful, intense`

**Documentary / Nature:**
`Ambient cinematic, 80 BPM, evolving pads, gentle piano motif, wide atmospheric mix, contemplative, awe-inspiring, National Geographic feel`

**Tech / Startup:**
`Modern electronic, 118 BPM, clean synth arpeggios, subtle glitch percussion, futuristic, innovative, minimal, forward-thinking`

**Kids / Animation:**
`Playful bouncy, 130 BPM, pizzicato strings, xylophone melody, toy piano, cheerful woodwinds, cartoon, silly, fun`

**Podcast Intro:**
`Chill groovy, 95 BPM, funky bass riff, clean guitar, light drums, lo-fi warmth, personality-driven, 10 seconds`

**Wedding / Romance:**
`Romantic piano ballad, 72 BPM, solo grand piano, gentle string quartet swell, tender, emotional, cinematic, intimate`

**Gaming / Esports:**
`Aggressive electronic, 145 BPM, distorted synth leads, heavy 808 bass, rapid hi-hats, cyberpunk, neon, adrenaline, competitive`

### Sound Design Layers

For professional video, build audio in layers. Generate each element separately and composite:

1. **Room Tone / Ambience** — the acoustic space:
   `Clean modern office room tone, air conditioning hum, subtle, 30 seconds seamless loop`

2. **Primary Foley** — key object interactions:
   `Hands unboxing premium product, crisp cardboard, tissue paper crinkling, ASMR close-up quality`

3. **Secondary Foley** — supporting sounds:
   `Laptop keyboard typing, soft mechanical keys, rhythmic, close-up`

4. **UI / Transition Sounds** — digital interactions:
   `Soft digital interface glow activation, warm ascending chime, futuristic, 1 second`

5. **Risers and Impacts** — transitions:
   `Cinematic tension riser, filtered noise sweep building over 3 seconds, leading to deep sub-bass impact hit with bright shimmer decay`

6. **Background Music** — the score, mixed lower than dialogue, ducking on voice.

7. **Dialogue / Voiceover** — primary voice, mixed front and center.

### Sync Considerations

- Generate music at known BPM so you can cut on the beat in your editor
- For reveal moments, use a separate "impact" SFX timed to land on the visual reveal
- If using a video model with `audio_url`, provide the audio at generation time so the video syncs to it
- Voice-over pacing: generate TTS first, then cut video to match speech timing — not the reverse
- For dialogue scenes, use a dialogue TTS model for natural rhythm, then edit video to match

---

## Ready-to-Use Prompt Examples

### Music

**Lo-Fi Study Beats:**
```
prompt: "Lo-fi hip hop, 85 BPM, C major, vinyl crackle and tape hiss, nostalgic Rhodes piano chords, warm analog bass, soft brushed snare, mellow, cozy, late night study session, instrumental only"
model: [use search_models to find a music generation model]
force_instrumental: true
music_length_ms: 180000
```

**Synthwave Driving Track:**
```
prompt: "Synthwave, 108 BPM, A minor, analog Juno pads, driving arpeggiated bassline, gated reverb snare, neon-lit 80s retrowave, pulsing sidechain, cinematic, outrun aesthetic, instrumental"
model: [use search_models to find a music generation model]
force_instrumental: true
music_length_ms: 240000
```

**Jazz Club Night:**
```
prompt: "Smoky jazz club, 120 BPM, Bb major, walking upright bass, brushed drums with ride cymbal, mellow tenor saxophone melody, warm Rhodes comping, intimate, late-night, analog recording warmth"
model: [use search_models to find a music generation model]
force_instrumental: true
music_length_ms: 180000
```

**Afrobeat Groove:**
```
prompt: "Afrobeat groove, 110 BPM, G minor, complex polyrhythmic percussion, funky guitar chops, deep bass groove, brass stabs, Fela Kuti inspired, warm analog production, energetic, danceable"
model: [use search_models to find a music generation model]
music_length_ms: 240000
```

**Ambient Meditation:**
```
prompt: "Deep ambient meditation, 60 BPM, ethereal evolving pads, Tibetan singing bowl resonance, soft granular textures, spacious reverb, cosmic, transcendent, healing frequencies, no percussion"
model: [use search_models to find a music/audio generation model]
```

### Sound Effects

**Thunderstorm:**
```
prompt: "Intense thunderstorm, heavy rain on window glass, close lightning cracks with immediate thunder, wind gusting, continuous storm, dramatic, cinematic quality, 30 seconds seamless loop"
model: [use search_models to find a sound effects model]
```

**Sci-Fi Door:**
```
prompt: "Futuristic spaceship airlock door opening, hydraulic hiss, pneumatic release, heavy metal sliding, pressurization equalization, sci-fi, high-tech, clean recording, 3 seconds"
model: [use search_models to find a sound effects model]
```

**Medieval Battle:**
```
prompt: "Medieval battlefield ambience, distant clash of swords and shields, war horns blowing, horses galloping, soldiers shouting battle cries, arrows whistling overhead, epic scale, 15 seconds"
model: [use search_models to find a sound effects model]
```

### Voice

**Documentary Narrator (high-quality TTS):**
```
text: "In the depths of the Pacific Ocean, at pressures that would crush steel, life persists. These creatures have evolved over millions of years... adapting to total darkness, near-freezing temperatures, and the crushing weight of the water above."
model: [use search_models to find a high-quality TTS model]
stability: 0.6
similarity_boost: 0.8
```

**Dialogue Scene (two-speaker TTS):**
```
text: "[S1] Did you hear that? (whispers) Something is moving in the hallway.
[S2] (sighs) It is probably just the cat again. Go back to sleep.
[S1] No, this is different. It sounds like... (gasps) footsteps. Heavy footsteps.
[S2] (clears throat) Okay. Stay here. I will go check."
model: [use search_models to find a dialogue/multi-speaker TTS model]
```

**Emotional Narration (expressive TTS):**
```
text: "She opened the letter with trembling hands. <sigh> After all these years... the words blurred through her tears. <cry> He had written it the night before he left — and she was only reading it now."
model: [use search_models to find an emotionally expressive TTS model]
```

**Commercial Voice-Over:**
```
text: "Introducing the all-new Horizon series. Crafted for those who demand more. More precision. More elegance. More of everything that matters. Horizon — redefine your standard."
model: [use search_models to find a high-quality TTS model]
stability: 0.5
similarity_boost: 0.75
```

**Meditation Guide:**
```
text: "Take a deep breath in... and slowly let it go. Feel your body soften... your shoulders dropping... your jaw releasing. With each breath... you sink a little deeper... into stillness... into peace."
model: [use search_models to find a high-quality TTS model]
stability: 0.7
similarity_boost: 0.6
```

**Multi-Speaker Dialogue (dialogue TTS):**
```
INTERVIEWER: "So what first drew you to deep-sea exploration?"
SCIENTIST: "Honestly? A Jacques Cousteau documentary when I was seven. I remember sitting on the living room floor, completely mesmerized."
INTERVIEWER: "And now you have spent, what, fifteen years at the bottom of the ocean?"
SCIENTIST: "[laughs] Well, not continuously. But yes, roughly fifteen years of expeditions."
model: [use search_models to find a text-to-dialogue TTS model]
```

**Chinese Voice-Over:**
```
text: "欢迎来到我们的频道。今天我们将探讨人工智能如何改变音乐创作的方式。"
model: [use search_models to find a TTS model with Chinese/Mandarin support]
```

---

## Pro Tips and Common Pitfalls

### Do

- Generate audio elements separately and layer them in post — do not try to get everything in one generation
- Use exact BPM, key signatures, and duration values
- Front-load the most important descriptors (genre, mood) in prompts
- For video, generate audio first and edit video to match (especially TTS)
- Test with shorter durations first, then scale up
- Use `force_instrumental: true` (or equivalent) when you need no vocals
- Specify "seamless loop" for background audio that needs to repeat
- Use `composition_plan` (if supported) for any track longer than 30 seconds
- Match BPM to your video edit timing for beat-synced cuts

### Do Not

- Do not use vague tempo words ("fast," "slow") — always specify BPM
- Do not request contradictory attributes ("aggressive lullaby," "dark uplifting")
- Do not overload prompts — 40 words is better than 100 unfocused words
- Do not expect perfect sync from text descriptions alone — use `audio_url` on models that support it for sync-critical work
- Do not mix more than 2–3 genre descriptors — models will average them into mush
- Do not forget to specify duration — models have different defaults and maximums
- Do not use command language ("Create a...", "Generate a...") — describe the audio itself
- Do not skip the layering approach for video — single-pass audio always sounds flat compared to layered design
