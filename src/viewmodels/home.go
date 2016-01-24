package viewmodels

type Home struct {
	Title  string
	Active string
}

func GetHome() Home {
	result := Home{
		Title:  "Vedado Control Center",
		Active: "home",
	}

	return result
}
