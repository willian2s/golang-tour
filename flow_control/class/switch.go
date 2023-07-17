package class

func Switch(os string) string {
	switch os {
	case "linux":
		return "Linux"
	case "darwin":
		return "OSX"
	case "windows":
		return "Windows"
	default:
		return "Unknown"
	}
}
