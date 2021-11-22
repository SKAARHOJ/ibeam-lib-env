//go:build skaaros_dev

package env

func init() {
	if current != "" {
		panic("do not combine env build tags")
	}
	current = SkaarOSDevelopment
}
