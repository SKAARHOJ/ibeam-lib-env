package env

// This will always execute, due to the z in the name it will always do so last
func init() {
	if current == "" {
		current = Development
	}
}
