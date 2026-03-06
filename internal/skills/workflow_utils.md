# Workflow Utilities — fal.ai Media Processing Tools

Post-processing, compositing, and utility endpoints for building media pipelines.
Use these AFTER generation to manipulate, combine, or transform media outputs.

## Quick Reference

> **Model Discovery:** Use `search_models` to find the best generation models for any task (image generation, video generation, audio, etc.). The workflow utilities below are post-processing tools — they manipulate outputs from generation models.

**Most Common Operations**

- Segment → isolate subject: segment → invert mask → composite
- Image to video pipeline: generate image → resize → I2V → join audio → auto-subtitle
- Social resize batch: resize (1:1 Instagram) + resize (9:16 Stories) → compress
- Product shot: segment → composite → add text → image grid

**Key Tips**

- Prefix matters: some endpoints use `workflowutils/` and others use `workflow-utilities/` — do not mix them up.
- Get image dimensions first when downstream steps require exact pixel dimensions.
- Segmentation produces masks; combine with invert, blur, grow, or shrink mask operations to refine edges before compositing.
- Auto-subtitle handles both transcription and karaoke-style highlighting in one call.
- Image compression accepts a quality value and auto-resizes — best final step before any web delivery.
- Video speed change range is 0.25x–4x; stay within this range to avoid errors.

---

## Image Utilities

### Resize & Transform
- `workflowutils/resize-image` — Resize image to specific dimensions
- `workflowutils/resize-to-max-pixels` — Resize maintaining aspect ratio to max pixel count
- `workflowutils/crop-image` — Crop using percentage-based coordinates
- `workflowutils/image-size` — Get width/height of an image
- `workflow-utilities/compress-image` — Compress with quality control and auto-resize

### Compositing & Layout
- `workflowutils/composite-image` — Composite (layer) images together
- `workflow-utilities/overlay-image` — Overlay image on background with position, scale, opacity, stroke
- `workflow-utilities/concat-image` — Concatenate images horizontally or vertically with spacing/alignment
- `workflow-utilities/image-grid` — Create grid layouts from multiple images (columns, spacing, fit modes)
- `workflow-utilities/add-text-to-image` — Add text with font, color, background, stroke options

### Format Conversion
- `workflowutils/rgba-to-rgb` — Convert RGBA to RGB with background color
- `workflow-utilities/split` — Split multiple images into individual outputs with format conversion
- `workflow-utilities/merge` — Merge multiple images into single array output

### Masking & Segmentation
- `workflowutils/sam-hq` — SAM HQ: high-quality zero-shot segmentation (generate masks)
- `workflowutils/invert_mask` — Invert a mask
- `workflowutils/blur_mask` — Blur a mask (soften edges)
- `workflowutils/grow_mask` — Expand mask region
- `workflowutils/shrink_mask` — Shrink mask region
- `workflowutils/transparent-image-to_mask` — Convert transparent image to mask

### Edge Detection
- `workflowutils/teed` — TEED: fast SOTA edge detector
- `workflowutils/canny` — Canny edge detector
- `workflowutils/image-preprocessors-canny` — Canny edge detector (alternative endpoint)

### Face & Detection
- `workflowutils/face` — Face detection and processing

### Content Safety
- `workflowutils/brand-checker` — Detect brands in image (returns error if found)
- `workflowutils/llm-nsfw-checker` — Analyze text for NSFW content

## Video Utilities

### Video Manipulation
- `workflow-utilities/overlay-video` — Overlay video on another with position, scale, opacity
- `workflow-utilities/setpts-video` — Change playback speed (0.25x to 4x)
- `workflow-utilities/add-subtitles-to-video` — Add customizable subtitles (TikTok-style formatting)
- `workflow-utilities/auto-subtitle` — Auto-generate subtitles with karaoke-style word highlighting
- `workflowutils/join-audio-video` — Merge audio track with video

### Video Conversion
- `workflow-utilities/video-to-gif` — Convert video segment to animated GIF
- `workflow-utilities/gif-to-video` — Convert GIF to video with format/transparency controls

## Audio Utilities

### Audio Manipulation
- `workflow-utilities/extract-audio` — Extract audio from video (format + bitrate options)
- `workflow-utilities/split-audio` — Split audio at timestamps into segments
- `workflow-utilities/merge-audio` — Merge audio files with optional gaps
- `workflow-utilities/amix-audio` — Mix audio streams with weights and normalization

## Text Utilities

- `workflow-utilities/merge-text` — Merge text strings with custom separator
- `workflow-utilities/split-text` — Split text using custom separator
- `workflowutils/text-concat` — Concatenate text strings, combine LoRA triggers, merge prompts

## Generation (Inference)

### Image Generation
- `workflowutils/any-stable-diffusion-xl` — Run SDXL with dedicated APIs
- `workflowutils/shopify-lora` — Run any SD model with customizable LoRA weights
- `workflowutils/flux-1-img-2-img` — FLUX.1 image-to-image (12B params, outstanding aesthetics)

### Video Generation
- `workflowutils/luma-dream-machine-image-to-video` — Luma Dream Machine v1.5 image-to-video

## Common Workflows

### Image → Video Pipeline
1. Generate image: use a text-to-image model (use `search_models`)
2. Upscale if needed: use an upscaling model (use `search_models`)
3. Generate video: use an image-to-video model (use `search_models`)
4. Add audio: `workflowutils/join-audio-video`
5. Add subtitles: `workflow-utilities/auto-subtitle`

### Background Replacement
1. Segment subject: `workflowutils/sam-hq`
2. Invert mask: `workflowutils/invert_mask`
3. Generate new background: use a text-to-image model (use `search_models`)
4. Composite: `workflowutils/composite-image`

### Multi-shot Video
1. Generate multiple video clips from same source image
2. Join with transitions using overlay-video
3. Add background music: `workflowutils/join-audio-video`
4. Add subtitles: `workflow-utilities/auto-subtitle`

### Product Photography
1. Generate product image: use a text-to-image model (use `search_models`)
2. Segment product: `workflowutils/sam-hq`
3. Remove background: use mask to isolate
4. Add new background: `workflowutils/composite-image`
5. Add text overlay: `workflow-utilities/add-text-to-image`
6. Create grid for catalog: `workflow-utilities/image-grid`

### Social Media Batch
1. Generate hero image at 16:9
2. Resize for Instagram: `workflowutils/resize-image` (1:1)
3. Resize for Stories: `workflowutils/resize-image` (9:16)
4. Compress for web: `workflow-utilities/compress-image`
