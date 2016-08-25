package viewmodels

type ErrorPage struct {
	Title  string
  Code string
}

func Get404() ErrorPage {
	result := ErrorPage{
		Title:  "Page Not Found",
    Code: "404",
	}

	return result
}
