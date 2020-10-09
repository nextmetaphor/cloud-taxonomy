package model

type (
	// Service TODO
	Service struct {
		ID          string `json:"ID" yaml:"ID"`
		CategoryID  string `json:"Category-ID" yaml:"category-ID"`
		Title       string `json:"Title" yaml:"title"`
		Description string `json:"Description" yaml:"description"`
	}
)
