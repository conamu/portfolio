package config

import (
	"encoding/json"
	"io/fs"
)

// Theme holds all visual configuration loaded from data/theme.json.
type Theme struct {
	Light        ThemeColors `json:"light"`
	Dark         ThemeColors `json:"dark"`
	FontFamily   string      `json:"font_family"`
	BorderRadius string      `json:"border_radius"`
	BorderWidth  string      `json:"border_width"`
	ShadowOffset string      `json:"shadow_offset"`
	LogoText     string      `json:"logo_text"`
	SiteTitle    string      `json:"site_title"`
}

type ThemeColors struct {
	Background  string `json:"background"`
	Surface     string `json:"surface"`
	Primary     string `json:"primary"`
	PrimaryDark string `json:"primary_dark"`
	Accent      string `json:"accent"`
	AccentDark  string `json:"accent_dark"`
	Text        string `json:"text"`
	TextMuted   string `json:"text_muted"`
	Border      string `json:"border"`
	Shadow      string `json:"shadow"`
	NavBg       string `json:"nav_bg"`
	TagBg       string `json:"tag_bg"`
	TagText     string `json:"tag_text"`
}

// Project is one entry in data/projects.json.
type Project struct {
	ID                string   `json:"id"`
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	Image             string   `json:"image"`
	UnsplashAuthor    string   `json:"unsplash_author"`     // display name, e.g. "John Doe"
	UnsplashAuthorURL string   `json:"unsplash_author_url"` // link to author's Unsplash profile
	Tags              []string `json:"tags"`
	URL               string   `json:"url"`
	GitHub            string   `json:"github"`
}

// AboutLink is a social/external link on the about page.
type AboutLink struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

// About holds the content for data/about.json.
type About struct {
	Name     string      `json:"name"`
	Tagline  string      `json:"tagline"`
	Bio      string      `json:"bio"`
	Photo    string      `json:"photo"`
	CVURL    string      `json:"cv_url"`     // public rxresume URL (open in new tab button)
	CVApiURL string      `json:"cv_api_url"` // rxresume stealthy API endpoint to fetch CV data
	CVLabel  string      `json:"cv_label"`
	ShowCV   bool        `json:"show_cv"`
	Links    []AboutLink `json:"links"`
}

func LoadTheme(fsys fs.FS, path string) (Theme, error) {
	var t Theme
	return t, loadJSON(fsys, path, &t)
}

func LoadProjects(fsys fs.FS, path string) ([]Project, error) {
	var p []Project
	return p, loadJSON(fsys, path, &p)
}

func LoadAbout(fsys fs.FS, path string) (About, error) {
	var a About
	return a, loadJSON(fsys, path, &a)
}

func loadJSON(fsys fs.FS, path string, v any) error {
	f, err := fsys.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(v)
}
