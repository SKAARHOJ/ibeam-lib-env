//go:build prod

package env

func init() {
	if current != "" {
		panic("do not combine env build tags")
	}
	current = Production
}
