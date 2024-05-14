package resources

import "code.cloudfoundry.org/cli/types"

type Sidecar struct {
	GUID          string               `json:"guid,omitempty"`
	Name          string               `json:"name"`
	Command       types.FilteredString `json:"command"`
	ProcessTypes  []string             `json:"process_types"`
	MemoryInMB    *int                 `json:"memory_in_mb,omitempty"`
	Origin        *string              `json:"origin,omitempty"`
	Relationships Relationships        `json:"relationships,omitempty"`
}
