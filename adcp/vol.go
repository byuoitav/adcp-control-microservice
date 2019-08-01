package adcp

// the volume level that the projectors put out is only really
// useful from 0-50(ish). above 50 or so, the volume seems to stay
// somewhat constant. these functions map the given 0-100 volume
// to the min and the max that we set.

const (
	minAdcp = 0
	maxAdcp = 50
)

func normalToAdcpVolume(level int) int {
	switch {
	case level >= 0 && level <= 100:
		return level / 2
	case level < 0:
		return minAdcp
	case level > 100:
		return maxAdcp
	default:
		return level
	}
}

func adcpToNormalVolume(level int) int {
	switch {
	case level >= minAdcp && level <= maxAdcp:
		return level * 2
	case level < minAdcp:
		return minAdcp
	case level > maxAdcp:
		return maxAdcp
	default:
		return level
	}
}
