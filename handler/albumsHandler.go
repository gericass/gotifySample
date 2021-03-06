package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"
)

func AlbumsHandler(c echo.Context) error {

	albumIDs := []string{"4sgYpkIASM1jVlNC8Wp9oF", "1c9Sx7XdXuMptGyfCB6hHs"}

	res, err := Token.GetAlbums(albumIDs)
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(out))
}

func AlbumsTracksHandler(c echo.Context) error {

	albumID := "4sgYpkIASM1jVlNC8Wp9oF"

	res, err := Token.GetAlbumsTracks(albumID)
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(out))
}
