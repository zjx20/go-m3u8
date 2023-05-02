package m3u8

import (
	"fmt"
	"strconv"
	"strings"
)

// SegmentItem represents EXTINF attributes with the URI that follows,
// optionally allowing an EXT-X-BYTERANGE tag to be set.
type SegmentItem struct {
	Duration        float64
	Segment         string
	Comment         *string
	ProgramDateTime *TimeItem
	ByteRange       *ByteRange
	Discontinuity   *DiscontinuityItem
	Attributes      *string
}

// NewSegmentItem parses a text line and returns a *SegmentItem
func NewSegmentItem(text string) (*SegmentItem, error) {
	var si SegmentItem
	line := strings.Replace(text, SegmentItemTag+":", "", -1)
	line = strings.Replace(line, "\n", "", -1)
	values := strings.Split(line, ",")
	parts := strings.Split(values[0], " ")
	d, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return nil, err
	}

	si.Duration = d
	if len(values) > 1 && values[1] != "" {
		si.Comment = &values[1]
	}
	if len(parts) > 1 {
		attrs := strings.Join(parts[1:], " ")
		si.Attributes = &attrs
	}

	return &si, nil
}

func (si *SegmentItem) String() string {
	date := ""
	if si.ProgramDateTime != nil {
		date = fmt.Sprintf("%v\n", si.ProgramDateTime)
	}
	byteRange := ""
	if si.ByteRange != nil {
		byteRange = fmt.Sprintf("%s:%v\n", ByteRangeItemTag, si.ByteRange.String())
	}

	comment := ""
	if si.Comment != nil {
		comment = *si.Comment
	}

	discontinuity := ""
	if si.Discontinuity != nil {
		discontinuity = si.Discontinuity.String() + "\n"
	}

	attrs := ""
	if si.Attributes != nil {
		attrs = " " + *si.Attributes
	}
	return fmt.Sprintf("%s:%v%s,%s\n%s%s%s%s", SegmentItemTag, si.Duration, attrs, comment,
		byteRange, date, discontinuity, si.Segment)
}
