package commands

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/mikkeloscar/flise/backend"
	"github.com/mikkeloscar/flise/context"
)

// Exec implements the exec command.
type Exec string

// Exec executes the exec command.
func (e Exec) Exec(ctx context.Context) error {
	log.Debugf("Running command '%s'", e)
	backend.Get(ctx).Exec("/bin/sh", "-c", string(e))
	return nil
}

// String returns a string formatting of the exec command.
func (e Exec) String() string {
	return string(e)
}

// parseExec parses an exec command definition.
func parseExec(lex *lexer) (Executer, error) {
	var args []string

	for t := lex.nextItem(); t.typ == itemString; t = lex.nextItem() {
		args = append(args, fmt.Sprintf(`"%s"`, t.val))
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("no command defined")
	}

	return Exec(strings.Join(args, " ")), nil
}
