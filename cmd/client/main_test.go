// Пакет main содержит тесты для клиента.
package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

// TestClientVersion_NoConfig тестирует версию клиента без конфигурации.
func TestClientVersion_NoConfig(t *testing.T) {
	bin := filepath.Join("..", "..", "bin", "client")
	if _, err := os.Stat(bin); err != nil {
		t.Skip("client binary not built; run make client")
	}
	cmd := exec.Command(bin, "--version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("client --version failed: %v, out=%s", err, string(out))
	}
	if len(out) == 0 {
		t.Fatalf("expected some output, got empty")
	}
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
	if err := os.WriteFile(cfgPath, []byte(`{"version":{"version":"X","date":"Y"}}`), 0644); err != nil {
		t.Fatalf("write config: %v", err)
	}
	cmd := exec.Command(bin, "--version", "--config", cfgPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("client --version --config failed: %v, out=%s", err, string(out))
	}
	got := string(out)
	if got != "version=X date=Y\n" {
		t.Fatalf("unexpected output: %q", got)
	}
}
