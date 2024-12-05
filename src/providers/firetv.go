package providers

func decodeFireTV(_ string, id string) (string, error) {
	return "https://d21m0ezw6fosyw.cloudfront.net/" + id + ".jpg", nil
}
