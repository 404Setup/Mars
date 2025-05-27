package version

import "Mars/database/schemas"

// Pattern for version names
const Pattern = "[0-9.]+-?(?:pre|SNAPSHOT)?(?:[0-9.]+)?"

// CompareVersion Comparator for version families (implement as needed)
func CompareVersion(vf1, vf2 schemas.IVersion) int {
	if vf1.GetTime() != nil && vf2.GetTime() != nil {
		if vf1.GetTime().Before(*vf2.GetTime()) {
			return -1
		} else if vf1.GetTime().After(*vf2.GetTime()) {
			return 1
		}
	}
	return compareStrings(vf1.GetName(), vf2.GetName())
}

func compareStrings(s1, s2 string) int {
	if s1 < s2 {
		return -1
	} else if s1 > s2 {
		return 1
	}
	return 0
}
