package handler

import (
	"github.com/labstack/echo"
	"encoding/json"
	"net/http"
)

func FollowingArtistsHandler(c echo.Context) error {

	res, err := Token.GetFollowingArtists()
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(out))
}

func FollowArtistsOrUsersHandler(c echo.Context) error {

	ids := []string{"5yKN5TpzvkQ17yDofamX90"}
	err := Token.FollowArtistsOrUsers("artist", ids)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "ok")
}

func UnfollowArtistsOrUsersHandler(c echo.Context) error {

	ids := []string{"5yKN5TpzvkQ17yDofamX90"}
	err := Token.UnfollowArtistsOrUsers("artist", ids)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "ok")
}

func CurrentFollowArtistsOrUsersHandler(c echo.Context) error {

	ids := []string{"5yKN5TpzvkQ17yDofamX90"}
	res, err := Token.CurrentFollowsArtistsOrUsers("artist", ids)
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}

func FollowPlaylistHandler(c echo.Context) error {

	userID := ""
	playlistID := ""

	err := Token.FollowPlaylist(userID, playlistID)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

func UnfollowPlaylistHandler(c echo.Context) error {

	userID := ""
	playlistID := ""

	err := Token.UnfollowPlaylist(userID, playlistID)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

func CheckFollowPlaylistHandler(c echo.Context) error {

	ownerID := ""
	playlistID := ""
	userIDs := []string{""}

	res, err := Token.CheckFollowPlaylist(ownerID, playlistID, userIDs)
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}
