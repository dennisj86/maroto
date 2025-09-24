package props

// List represents properties from a List.
type List struct {
	// MinimumRowsBypage is the minimum amount of rows that must fit on a page.
	// This limit is used to define whether the list will start on the current page or the next one.
	MinimumRowsBypage int

	// HeaderBottomMargin adds a vertical space (in the same unit as row heights)
	// immediately after the list header on every page where the header is rendered.
	// Use 0 for no additional spacing. Must be non-negative.
	HeaderBottomMargin float64
}

// MakeValid from List define default values for a list.
func (l *List) MakeValid() {
	minRows := 1
	if l.MinimumRowsBypage < minRows {
		l.MinimumRowsBypage = minRows
	}
	// Ensure non-negative margin; default 0 if not set or negative
	if l.HeaderBottomMargin < 0 {
		l.HeaderBottomMargin = 0
	}
}
