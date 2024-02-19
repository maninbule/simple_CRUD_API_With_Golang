package Type

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (m Movie) IsValid() bool {
	if len(m.Isbn) == 0 || len(m.Title) == 0 || len(m.ID) == 0 {
		return false
	}
	if m.Director == nil || len(m.Director.FirstName) == 0 || len(m.Director.LastName) == 0 {
		return false
	}
	return true
}
