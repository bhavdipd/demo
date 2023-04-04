package app

import (
	"github.com/bhavdipd/demo/domain"
	"github.com/bhavdipd/demo/service"
	"github.com/gin-gonic/gin"
)

func Start() {

	//Wiring
	// ah := AlbumHandlers{service.NewAlbumService(domain.NewAlbumRepositoryStub())}
	ah := AlbumHandlers{service.NewAlbumService(domain.NewAlbumRepositoryDb())}

	router := gin.Default()
	router.GET("/albums", ah.getAlbums)
	// router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
