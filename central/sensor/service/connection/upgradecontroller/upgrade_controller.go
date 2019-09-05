package upgradecontroller

import (
	"context"

	"github.com/stackrox/rox/central/sensor/service/common"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/concurrency"
)

// UpgradeController controls auto-upgrading for one specific cluster.
type UpgradeController interface {
	ErrorSignal() concurrency.ReadOnlyErrorSignal
	// RegisterConnection registers a new connection from a sensor, and a handle to send messages to it.
	// The return value is a once-triggered error waitable that gets triggered if there is any critical issue
	// with the upgrade controller.
	RegisterConnection(sensorCtx context.Context, connection common.MessageInjector) concurrency.ErrorWaitable
	ProcessCheckInFromUpgrader(req *central.UpgradeCheckInFromUpgraderRequest) (*central.UpgradeCheckInFromUpgraderResponse, error)
	Trigger(ctx concurrency.Waitable) error
}

// ClusterStorage is the fragment of the cluster store interface that is needed by the upgrade controller.
type ClusterStorage interface {
	UpdateClusterUpgradeStatus(ctx context.Context, clusterID string, status *storage.ClusterUpgradeStatus) error
	GetCluster(ctx context.Context, id string) (*storage.Cluster, bool, error)
}

// New returns a new UpgradeController for the given cluster.
func New(clusterID string, storage ClusterStorage, autoTriggerEnabledFlag *concurrency.Flag) (UpgradeController, error) {
	u := &upgradeController{
		autoTriggerEnabledFlag: autoTriggerEnabledFlag,
		clusterID:              clusterID,
		errorSig:               concurrency.NewErrorSignal(),
		storage:                storage,
	}

	if err := u.initialize(); err != nil {
		return nil, err
	}
	return u, nil
}
