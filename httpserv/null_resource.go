package httpserv

import (
	"fmt"
)

type NullResource struct {
}

func (n *NullResource) ContentType() string {
	return "text/plain"
}

func (n *NullResource) HandleCommand(comm string) string {
	return fmt.Sprintf("Succeed to Execute:%s", comm)
}
