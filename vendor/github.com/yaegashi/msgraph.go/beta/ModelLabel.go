// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// LabelActionBase undocumented
type LabelActionBase struct {
	// Object is the base model of LabelActionBase
	Object
	// Name undocumented
	Name *string `json:"name,omitempty"`
}

// LabelDetails undocumented
type LabelDetails struct {
	// Object is the base model of LabelDetails
	Object
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Color undocumented
	Color *string `json:"color,omitempty"`
	// Sensitivity undocumented
	Sensitivity *int `json:"sensitivity,omitempty"`
	// Tooltip undocumented
	Tooltip *string `json:"tooltip,omitempty"`
	// IsActive undocumented
	IsActive *bool `json:"isActive,omitempty"`
}

// LabelPolicy undocumented
type LabelPolicy struct {
	// Object is the base model of LabelPolicy
	Object
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
}