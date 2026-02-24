package msi

import (
	"github.com/google/uuid"
)

type PackageType int

const (
	PackageTypeInstaller PackageType = iota
	PackageTypePatch
	PackageTypeTransform
)

func PackageTypeFromCLSID(clsid uuid.UUID) PackageType {

	switch clsid {
	case uuid.MustParse(InstallerPackageClsid):
		return PackageTypeInstaller
	case uuid.MustParse(PatchPackageClsid):
		return PackageTypePatch
	case uuid.MustParse(TransformPackageClsid):
		return PackageTypeTransform
	default:
		return -1
	}
}

func (p PackageType) CLSID() uuid.UUID {
	switch p {
	case PackageTypeInstaller:
		return uuid.MustParse(InstallerPackageClsid)
	case PackageTypePatch:
		return uuid.MustParse(PatchPackageClsid)
	case PackageTypeTransform:
		return uuid.MustParse(TransformPackageClsid)
	default:
		return uuid.Nil
	}
}

func (p PackageType) String() string {
	switch p {
	case PackageTypeInstaller:
		return "Installation Database"
	case PackageTypePatch:
		return "Patch"
	case PackageTypeTransform:
		return "Transform"
	default:
		return "Unknown"
	}
}
