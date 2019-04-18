package singleton

import (
	"time"

	"github.com/stackrox/rox/pkg/buildinfo"
	"github.com/stackrox/rox/pkg/license/validator"
	"github.com/stackrox/rox/pkg/utils"
)

func init() {
	utils.Must(
		validatorInstance.RegisterSigningKey(
			validator.EC256,
			[]byte{
				0x30, 0x59, 0x30, 0x13, 0x06, 0x07, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x02,
				0x01, 0x06, 0x08, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x03, 0x01, 0x07, 0x03,
				0x42, 0x00, 0x04, 0x67, 0x91, 0x74, 0xc6, 0x87, 0x11, 0xf8, 0x89, 0xb8,
				0x30, 0x4b, 0xec, 0x31, 0x74, 0x96, 0x5a, 0x3e, 0x2b, 0x5d, 0xa6, 0x7d,
				0xb3, 0x04, 0x75, 0xde, 0xd0, 0x7e, 0xac, 0xe7, 0xa6, 0xe0, 0xbb, 0x12,
				0x2b, 0x15, 0xbc, 0x88, 0x26, 0x59, 0x9d, 0xdb, 0xe2, 0x06, 0x90, 0x71,
				0xa4, 0xc6, 0x48, 0xb5, 0xa7, 0x7a, 0xb6, 0xcd, 0xd7, 0xb3, 0xcb, 0xbe,
				0x36, 0x40, 0x8a, 0xb9, 0x6b, 0x1e, 0x21,
			},
			validator.SigningKeyRestrictions{
				EarliestNotValidBefore:        buildinfo.BuildTimestamp(),
				LatestNotValidAfter:           buildinfo.BuildTimestamp().Add(48 * time.Hour),
				MaxDuration:                   6 * time.Hour,
				AllowOffline:                  true,
				MaxNodeLimit:                  10,
				AllowNoBuildFlavorRestriction: true,
				DeploymentEnvironments:        []string{"gcp/stackrox-ci", "aws/051999192406"},
			}),
	)
}
