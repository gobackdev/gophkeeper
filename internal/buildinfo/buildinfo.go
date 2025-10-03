// Пакет buildinfo предоставляет информацию о сборке.
package buildinfo

import "fmt"

// Значения переписываются через -ldflags при сборке
var (
	Version = "dev"
	Commit  = "unknown"
	Date    = "unknown"
)

// String возвращает строку с информацией о сборке
func String() string {
	return fmt.Sprintf("version=%s commit=%s date=%s", Version, Commit, Date)
}
