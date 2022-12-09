package pkg

const (
	media_url = "https://api.kron-x.ru/media/"
)

func PhotoCheckerForCms(photo string) string {
	if photo == "" {
		return ""
	} else {
		return media_url + photo
	}
}
