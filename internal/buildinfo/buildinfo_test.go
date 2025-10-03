// Пакет buildinfo содержит тесты для buildinfo.
package buildinfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestString_Defaults тестирует строку с информацией о сборке по умолчанию.
func TestString_Defaults(t *testing.T) {
	Version, Commit, Date = "dev", "unknown", "unknown"
	got := String()
	want := "version=dev commit=unknown date=unknown"
	assert.Equal(t, want, got)
}

// TestString_Custom тестирует строку с информацией о сборке с настраиваемыми значениями.
func TestString_Custom(t *testing.T) {
	Version, Commit, Date = "1.2.3", "abc123", "2025-10-03T12:00:00Z"
	got := String()
	want := "version=1.2.3 commit=abc123 date=2025-10-03T12:00:00Z"
	assert.Equal(t, want, got)
}
