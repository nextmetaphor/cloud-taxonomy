package model

type (
	// Category TODO
	Category struct {
		ID          string `json:"ID" yaml:"ID"`
		Title       string `json:"Title" yaml:"title"`
		Description string `json:"Description" yaml:"description"`
	}
)
