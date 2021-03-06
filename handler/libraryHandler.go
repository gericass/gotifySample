package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"
)

func SaveTracksHandler(c echo.Context) error {

	ids := []string{"0udZHhCi7p1YzMlvI4fXoK"}

	err := Token.SaveTracks(ids)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}


func GetUsersSavedTracksHandler(c echo.Context) error {


	res, err := Token.GetUsersSavedTracks()
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}

func RemoveUsersSavedTracksHandler(c echo.Context) error {

	ids := []string{"0udZHhCi7p1YzMlvI4fXoK"}

	err := Token.RemoveUsersSavedTracks(ids)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

func CheckUsersSavedTracksHandler(c echo.Context) error {

	ids := []string{"0udZHhCi7p1YzMlvI4fXoK"}

	res, err := Token.CheckUsersSavedTracks(ids)
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}

func SaveAlbumsHandler(c echo.Context) error {

	ids := []string{"72N66a5gEVxk2xtSeymwIM"}

	err := Token.SaveAlbums(ids)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

func GetUsersSavedAlbumsHandler(c echo.Context) error {

	res, err := Token.GetUsersSavedAlbums()
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}

func RemoveAlbumsForCurrentUserHandler(c echo.Context) error {

	ids := []string{"72N66a5gEVxk2xtSeymwIM"}

	err := Token.RemoveAlbumsForCurrentUser(ids)
	if err != nil {
		return err
	}


	return c.String(http.StatusOK, "OK")
}

func CheckUsersSavedAlbumsHandler(c echo.Context) error {

	ids := []string{"72N66a5gEVxk2xtSeymwIM"}

	res, err := Token.CheckUsersSavedAlbums(ids)
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}