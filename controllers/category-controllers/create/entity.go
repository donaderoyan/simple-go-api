package createCategory

type InputCreateCategory struct {
	ParentID  string `json:"parent_id" validate:""`
	SectionID string `json:"section_id" validate:""`
	Name      string `json:"name" validate:"required"`
	Slug      string `json:"slug" validate:""`
}
