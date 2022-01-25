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
	return current == Development
}

func IsProd() bool {
	return current == Production
}

func IsSkaarOSDev() bool {
	return current == SkaarOSDevelopment
}

func IsSkaarOSProd() bool {
	return current == SkaarOSProduction
}

func IsSkaarOS() bool {
	return current == SkaarOSProduction || current == SkaarOSDevelopment
}
