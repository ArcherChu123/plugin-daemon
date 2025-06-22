package exception

import "github.com/ArcherChu123/plugin-daemon/pkg/entities"

type PluginDaemonError interface {
	error

	ToResponse() *entities.Response
}
