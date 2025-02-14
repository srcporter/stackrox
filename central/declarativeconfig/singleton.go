package declarativeconfig

import "github.com/stackrox/rox/pkg/sync"

var (
	once     sync.Once
	instance Manager
)

// ManagerSingleton provides the instance of Manager to use.
func ManagerSingleton() Manager {
	once.Do(func() {
		instance = New()
	})
	return instance
}
