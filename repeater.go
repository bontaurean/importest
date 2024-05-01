package importest

import "strings"

type Repeater struct{}

const TIMES = 5

func (Repeater) Repeat(s string) string {
	return strings.Repeat(s, TIMES)
}
