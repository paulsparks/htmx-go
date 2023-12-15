package helper_functions

type TodoItem struct {
	Id         int
	ItemString string
}

type Link struct {
	RouteName string
	URL       string
}

type NavbarProps struct {
	Links []Link
}
