package gumble

// Version represents a Mumble client or server version.
type Version struct {
	// The semantic version information as a single unsigned double.
	//
	// Bits 0-15 are the major version, bits 16-31 are the minor version, and
	// bits 32-47 are the patch version.
	// https://github.com/mumble-voip/mumble/blob/master/src/Version.h
	Version uint64
	// The name of the client.
	Release string
	// The operating system name.
	OS string
	// The operating system version.
	OSVersion string
}

// SemanticVersion returns the version's semantic version components.
func (v *Version) SemanticVersion() (major, minor, patch uint16) {
	major = uint16(v.Version>>48) & 0xFFFF
	minor = uint16(v.Version>>32) & 0xFFFF
	patch = uint16(v.Version>>16) & 0xFFFF
	return
}
