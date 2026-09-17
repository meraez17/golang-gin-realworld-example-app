package common

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("JWT_SECRET", "unit-test-secret-with-at-least-32-characters")
	os.Exit(m.Run())
}

func TestJWTSecretRejectsMissingAndShortValues(t *testing.T) {
	original := os.Getenv("JWT_SECRET")
	t.Cleanup(func() { _ = os.Setenv("JWT_SECRET", original) })
	_ = os.Unsetenv("JWT_SECRET")
	if _, err := JWTSecret(); err == nil { t.Fatal("missing secret must fail") }
	_ = os.Setenv("JWT_SECRET", "short")
	if _, err := JWTSecret(); err == nil { t.Fatal("short secret must fail") }
}
