package std

import (
	"time"

	"github.com/art-media-platform/amp-sdk-go/amp"
	"github.com/art-media-platform/amp-sdk-go/stdlib/tag"
)

var (
	LoginSpec           = amp.AttrSpec.With("Login").ID
	LoginChallengeSpec  = amp.AttrSpec.With("LoginChallenge").ID
	LoginResponseSpec   = amp.AttrSpec.With("LoginResponse").ID
	LoginCheckpointSpec = amp.AttrSpec.With("LoginCheckpoint").ID

	CellChildren   = amp.AttrSpec.With("children.TagID") // ID suffix denotes SeriesIndex is used to store a CellID
	CellProperties = amp.AttrSpec.With("cell-properties")
	LaunchURL      = amp.AttrSpec.With("LaunchURL").ID

	CellProperty      = tag.Spec{}.With("cell-property")
	CellLabel         = CellProperty.With("text.Tag.label").ID
	CellCaption       = CellProperty.With("text.Tag.caption").ID
	CellSynopsis      = CellProperty.With("text.Tag.synopsis").ID
	CellCollection    = CellProperty.With("text.Tag.collection").ID
	CellGlyphs        = CellProperty.With("Tag.glyphs").ID
	CellContentLink   = CellProperty.With("content-link").ID
	CellAuthor        = CellProperty.With("author").ID
	CellExternalLinks = CellProperty.With("external-link").ID
	CellFileInfo      = CellProperty.With("FileInfo").ID
	CellPortrait      = CellProperty.With("portrait").ID
)

const (
	// URL prefix for a glyph and is typically followed by a media (mime) type.
	GenericGlyphURL = "amp:glyph/"

	GenericImageType = "image/*"
	GenericAudioType = "audio/*"
	GenericVideoType = "video/*"
)

// Common universal glyphs
var (
	GenericFolderGlyph = &amp.Tag{
		URL: GenericGlyphURL + "application/x-directory",
	}
)

type PinnableAttr struct {
	Spec tag.Spec
}

func (v *Position) MarshalToStore(in []byte) (out []byte, err error) {
	return amp.MarshalPbToStore(v, in)
}

func (v *Position) TagSpec() tag.Spec {
	return amp.AttrSpec.With("Position")
}

func (v *Position) New() tag.Value {
	return &Position{}
}

func (v *FSInfo) MarshalToStore(in []byte) (out []byte, err error) {
	return amp.MarshalPbToStore(v, in)
}

func (v *FSInfo) TagSpec() tag.Spec {
	return amp.AttrSpec.With("FSInfo")
}

func (v *FSInfo) New() tag.Value {
	return &FSInfo{}
}

func (v *FSInfo) SetModifiedAt(t time.Time) {
	tag := tag.FromTime(t, false)
	v.ModifiedAt = int64(tag[0])
}

func (v *FSInfo) SetCreatedAt(t time.Time) {
	tag := tag.FromTime(t, false)
	v.CreatedAt = int64(tag[0])
}
