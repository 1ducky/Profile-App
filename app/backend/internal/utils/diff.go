package utils

func Diff(oldList, newList []string) []string {
	newSet := make(map[string]struct{}, len(newList))
	for _, v := range newList {
		newSet[v] = struct{}{}
	}

	var removed []string
	for _, v := range oldList {
		if _, ok := newSet[v]; !ok {
			removed = append(removed, v)
		}
	}
	return removed
}
