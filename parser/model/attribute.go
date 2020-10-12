package model

type (
	// Attribute TODO
	Attribute struct {
		ID          string `json:"ID" yaml:"ID"`
		Title       string `json:"Title" yaml:"title"`
		Description string `json:"Description" yaml:"description"`
	}
)
