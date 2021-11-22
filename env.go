package env

type SKENV string

const (
	Development        SKENV = "development"
	Production         SKENV = "production"
	SkaarOSDevelopment SKENV = "skaaros_dev"
	SkaarOSProduction  SKENV = "skaaros_prod"
)

var current SKENV = ""

func GetCurrent() SKENV {
	return current
}

// Helpers

func IsDev() bool {
	return current
}

func IsProd() bool {
	return current
}

func IsSkaarOSDev() bool {
	return current
}

func IsSkaarOSProd() bool {
	return current
}
