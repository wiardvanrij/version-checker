package api

import (
	"regexp"
	"time"
)

const (
	EnableAnnotationKey = "enable.version-checker.io"

	UseSHAAnnotationKey     = "use-sha.version-checker.io"
	MatchRegexAnnotationKey = "match-regex.version-checker.io"

	// MetaData is defined as a tag containing anything after the patch digit.
	// e.g. v1.0.1-gke.3 v1.0.1-alpha.0, v1.2.3.4
	UseMetaDataAnnotationKey = "use-metadata.version-checker.io"

	PinMajorAnnotationKey = "pin-major.version-checker.io"
	PinMinorAnnotationKey = "pin-minor.version-checker.io"
	PinPatchAnnotationKey = "pin-patch.version-checker.io"

	// TODO: set OS + arch options
)

// Options is used to describe what restrictions should be used for determining
// the latest image.
type Options struct {
	// UseSHA cannot be used with any other options
	UseSHA bool `json:"use-sha,omitempty"`

	MatchRegex *string `json:"match-regex,omitempty"`

	// UseMetaData defines whether tags with '-alpha', '-debian.0' etc. is
	// permissible.
	UseMetaData bool `json:"use-metadata,omitempty"`

	PinMajor *int64 `json:"pin-major,omitempty"`
	PinMinor *int64 `json:"pin-minor,omitempty"`
	PinPatch *int64 `json:"pin-patch,omitempty"`

	RegexMatcher *regexp.Regexp `json:"-"`
}

// ImageTag describes a container image tag.
type ImageTag struct {
	Tag          string    `json:"tag"`
	SHA          string    `json:"sha"`
	Timestamp    time.Time `json:"timestamp"`
	Architecture string    `json:"architecture,omitempty"`
	OS           string    `json:"os,omitempty"`
}
