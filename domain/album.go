package domain

type Album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

type AlbumRepository interface {
	FindAll() ([]Album, error)
}
