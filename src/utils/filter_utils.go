package utils

func IsAllowedExtension(allowedExtensions []string, extension string) bool {
	if len(allowedExtensions) == 1 && allowedExtensions[0] == "*" {
		return true
	}

	for _, allowedExtension := range allowedExtensions {
		if allowedExtension == extension {
			return true
		}
	}

	return false
}
