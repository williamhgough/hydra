package shields

import "strings"

// Shield structure
type Shield struct {
	front bool
	back  bool
	right bool
	left  bool
}

// ShieldBuilder structure
type ShieldBuilder struct {
	code string
}

// NewShieldBuilder - return new shield builder
func NewShieldBuilder() *ShieldBuilder {
	return new(ShieldBuilder)
}

// RaiseFront - Raise the front shields, captain!
func (sb *ShieldBuilder) RaiseFront() *ShieldBuilder {
	sb.code += "F"
	return sb
}

// RaiseBack - Raise the front shields, captain!
func (sb *ShieldBuilder) RaiseBack() *ShieldBuilder {
	sb.code += "B"
	return sb
}

// RaiseRight - Raise the front shields, captain!
func (sb *ShieldBuilder) RaiseRight() *ShieldBuilder {
	sb.code += "R"
	return sb
}

// RaiseLeft - Raise the front shields, captain!
func (sb *ShieldBuilder) RaiseLeft() *ShieldBuilder {
	sb.code += "L"
	return sb
}

// Build - Add the shields!
func (sb *ShieldBuilder) Build() *Shield {
	code := sb.code
	return &Shield{
		front: strings.Contains(code, "F"),
		back:  strings.Contains(code, "B"),
		right: strings.Contains(code, "R"),
		left:  strings.Contains(code, "L"),
	}
}
