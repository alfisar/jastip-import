package paggination

func Paggination(page int, limit int) (pages int, limits int) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limits = 10
	} else {
		limits = limit
	}

	pages = (page - 1) * limits

	return
}
