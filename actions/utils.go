package actions

func actionError(err error) map[string]string {
	return map[string]string{"error": err.Error()}
}
