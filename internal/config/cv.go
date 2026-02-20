package config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// cvAPIResponse is the raw envelope returned by the rxresume public API.
// The actual content sits one level down inside "data".
type cvAPIResponse struct {
	Data CVData `json:"data"`
}

// CVData holds the resume content unpacked from the API envelope.
type CVData struct {
	Basics   CVBasics   `json:"basics"`
	Sections CVSections `json:"sections"`
}

type CVBasics struct {
	Name     string    `json:"name"`
	Headline string    `json:"headline"`
	Email    string    `json:"email"`
	Location string    `json:"location"`
	Picture  CVPicture `json:"picture"`
	URL      CVLink    `json:"url"`
}

type CVPicture struct {
	URL string `json:"url"`
}

type CVLink struct {
	Href  string `json:"href"`
	Label string `json:"label"`
}

type CVSections struct {
	Experience CVSection[CVExperienceItem] `json:"experience"`
	Skills     CVSection[CVSkillItem]      `json:"skills"`
	Languages  CVSection[CVLanguageItem]   `json:"languages"`
	Certs      CVSection[CVCertItem]       `json:"certifications"`
	Projects   CVSection[CVProjectItem]    `json:"projects"`
}

type CVSection[T any] struct {
	Items []T `json:"items"`
}

type CVExperienceItem struct {
	Company  string `json:"company"`
	Position string `json:"position"`
	Date     string `json:"date"`
	Location string `json:"location"`
	Summary  string `json:"summary"`
	URL      CVLink `json:"url"`
}

type CVSkillItem struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
}

type CVLanguageItem struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CVCertItem struct {
	Name   string `json:"name"`
	Issuer string `json:"issuer"`
	Date   string `json:"date"`
	URL    CVLink `json:"url"`
}

type CVProjectItem struct {
	Name     string   `json:"name"`
	Summary  string   `json:"summary"`
	Date     string   `json:"date"`
	URL      CVLink   `json:"url"`
	Keywords []string `json:"keywords"`
}

var cvClient = &http.Client{Timeout: 8 * time.Second}

// FetchCV retrieves live CV data from the rxresume public API.
func FetchCV(apiURL string) (CVData, error) {
	resp, err := cvClient.Get(apiURL)
	if err != nil {
		return CVData{}, fmt.Errorf("fetch cv: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return CVData{}, fmt.Errorf("fetch cv: unexpected status %d", resp.StatusCode)
	}

	var envelope cvAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		return CVData{}, fmt.Errorf("fetch cv: decode: %w", err)
	}
	return envelope.Data, nil
}
