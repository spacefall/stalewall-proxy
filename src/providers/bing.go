package providers

func decodeBing(_ string, id string) (string, error) {
	return "https://bing.com/th?id=" + id + ".jpg&p=0&pid=hp&qlt=100", nil
}
