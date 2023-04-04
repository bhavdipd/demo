package domain

type AlbumRepositoryStub struct {
	albums []Album
}

func (s AlbumRepositoryStub) FindAll() ([]Album, error) {
	return s.albums, nil
}

func NewAlbumRepositoryStub() AlbumRepositoryStub {
	albums := []Album{
		{"1001", "Test 1", "John", 12.99},
		{"1002", "Test 2", "Michale", 22.99},
		{"1003", "Test 3", "Dany", 32.99},
	}

	return AlbumRepositoryStub{albums}
}
