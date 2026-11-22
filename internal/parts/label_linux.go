package parts

import (
	"syscall"

	"github.com/franzudev/echolocal/internal/layout"
)

func setLabel(path string) error {
	return syscall.Setxattr(path, "security.selinux", append([]byte(layout.OurLabel), 0), 0)
}
