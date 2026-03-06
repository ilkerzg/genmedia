package tools

import "github.com/fal-ai/genmedia/internal/skills"

// GetSkillTool loads domain knowledge progressively.
type GetSkillTool struct{}

func (t *GetSkillTool) Name() string { return "get_skill" }
func (t *GetSkillTool) Description() string {
	return "Load domain knowledge progressively — start broad, drill down only as needed. " +
		"1) get_skill(name) → TOC + quick reference cheat sheet. Often enough for simple tasks. " +
		"2) get_skill(name, section='key') → section content. Small sections load fully. " +
		"   Large sections return a sub-TOC listing their subsections. " +
		"3) get_skill(name, section='subsection_key') → specific subsection (e.g. 'veo_3_1'). " +
		"Start with the TOC. The quick reference covers 80% of use cases. " +
		"Only drill deeper when you need specific model parameters or detailed techniques."
}
func (t *GetSkillTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{
			"skill_name": map[string]interface{}{
				"type": "string",
				"description": "Skill to load. Available: cinematography, video_prompting, " +
					"image_prompting, commercial, audio_prompting, workflow_utils, " +
					"character_design, storytelling, social_media, motion_design, " +
					"world_building, brand_identity",
				"enum": []string{
					"cinematography", "video_prompting", "image_prompting", "commercial",
					"audio_prompting", "workflow_utils", "character_design", "storytelling",
					"social_media", "motion_design", "world_building", "brand_identity",
				},
			},
			"section": map[string]interface{}{
				"type": "string",
				"description": "Section or subsection key to load. Omit to get TOC. " +
					"Use a key from the TOC or sub-TOC. Supports fuzzy matching " +
					"(e.g. 'veo' matches 'veo_3_1_google_flagship').",
			},
		},
		"required": []string{"skill_name"},
	}
}

func (t *GetSkillTool) Execute(args map[string]interface{}, onProgress func(map[string]interface{})) (map[string]interface{}, error) {
	name, _ := args["skill_name"].(string)
	section, _ := args["section"].(string)

	if section != "" {
		content := skills.LoadSkillSection(name, section)
		if content == "" {
			available := skills.ListAvailableSkills()
			return map[string]interface{}{
				"ok":               false,
				"error":            "Skill '" + name + "' not found or not available",
				"available_skills": available,
			}, nil
		}
		return map[string]interface{}{
			"ok":      true,
			"skill":   name,
			"section": section,
			"content": content,
		}, nil
	}

	content := skills.LoadSkillTOC(name)
	if content == "" {
		available := skills.ListAvailableSkills()
		return map[string]interface{}{
			"ok":               false,
			"error":            "Skill '" + name + "' not found or not available",
			"available_skills": available,
		}, nil
	}
	return map[string]interface{}{
		"ok":      true,
		"skill":   name,
		"section": "toc",
		"content": content,
	}, nil
}
