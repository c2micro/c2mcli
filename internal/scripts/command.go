package scripts

import (
	"fmt"
	"strings"

	"github.com/c2micro/c2mcli/internal/scripts/aliases"
	"github.com/c2micro/c2mcli/internal/scripts/aliases/alias"
)

// обработка и выполнение алиаса
func ProcessCommand(bid uint32, cmd string) error {
	cmd = strings.TrimSpace(cmd)

	c := strings.Split(cmd, " ")
	if aliases.IsAliasExists(c[0]) {
		return alias.BackendAlias(bid, cmd)
	}
	return fmt.Errorf("unknown alias '%s'", c[0])
}
