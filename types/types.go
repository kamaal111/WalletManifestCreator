package types

// Manifest represents the structure of the manifest.json file
type Manifest struct {
	Icon         string `json:"icon.png,omitempty"`
	Icon2x       string `json:"icon@2x.png,omitempty"`
	Logo         string `json:"logo.png,omitempty"`
	Logo2x       string `json:"logo@2x.png,omitempty"`
	Background   string `json:"background.png,omitempty"`
	Background2x string `json:"background@2x.png,omitempty"`
	Footer       string `json:"footer.png,omitempty"`
	Footer2x     string `json:"footer@2x.png,omitempty"`
	Strip        string `json:"strip.png,omitempty"`
	Strip2x      string `json:"strip@2x.png,omitempty"`
	Thumbnail    string `json:"thumbnail.png,omitempty"`
	Thumbnail2x  string `json:"thumbnail@2x.png,omitempty"`
	Pass         string `json:"pass.json,omitempty"`
}
