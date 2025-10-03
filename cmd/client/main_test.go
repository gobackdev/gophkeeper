// Пакет main содержит тесты для клиента.
package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestClientVersion_NoConfig тестирует версию клиента без конфигурации.
func TestClientVersion_NoConfig(t *testing.T) {
	bin := filepath.Join("..", "..", "bin", "client")
	if _, err := os.Stat(bin); err != nil {
		t.Skip("client binary not built; run make client")
	}
	cmd := exec.Command(bin, "--version")
	out, err := cmd.CombinedOutput()
	require.NoError(t, err, "client --version failed: %s", string(out))
	assert.NotEmpty(t, out, "expected some output")
}

// TestClientVersion_WithConfig тестирует версию клиента с конфигурацией.
func TestClientVersion_WithConfig(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skip on windows path quoting differences")
	}
	bin := filepath.Join("..", "..", "bin", "client")
	if _, err := os.Stat(bin); err != nil {
		t.Skip("client binary not built; run make client")
	}
	cfg := t.TempDir()
	cfgPath := filepath.Join(cfg, "agent.json")
	err := os.WriteFile(cfgPath, []byte(`{"version":{"version":"X","date":"Y"}}`), 0644)
	require.NoError(t, err, "write config")

	cmd := exec.Command(bin, "--version", "--config", cfgPath)
	out, err := cmd.CombinedOutput()
	require.NoError(t, err, "client --version --config failed: %s", string(out))

	got := string(out)
	assert.Equal(t, "version=X date=Y\n", got)
}
