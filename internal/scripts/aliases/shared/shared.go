package shared

import (
	"github.com/c2micro/c2mcli/internal/storage/beacon"
	"github.com/c2micro/c2mshr/defaults"
)

func BackendIsOs(id uint32, os defaults.BeaconOS) bool {
	b := beacon.Beacons.GetById(id)
	if b == nil {
		return false
	}
	return b.GetOs() == os
}

func BackendIsArch(id uint32, arch defaults.BeaconArch) bool {
	b := beacon.Beacons.GetById(id)
	if b == nil {
		return false
	}
	return b.GetArch() == arch
}
