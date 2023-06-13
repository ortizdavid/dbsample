package helpers

func Contains(items []string, value string) bool {
	for _, item := range items {
		if value == item {
			return true
		}
	}
	return false
}

func GetItem(items []any, value string) any {
	for _, item := range items {
		if value == item {
			return item
		}
	}
	return nil
}