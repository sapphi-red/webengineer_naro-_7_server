package router

import (
	"github.com/labstack/echo"
	"github.com/sapphi-red/webengineer_naro-_7_server/database"
	"github.com/sapphi-red/webengineer_naro-_7_server/database/posts"
	"net/http"
)

func getPostsHandler(c echo.Context) error {
	var posts []posts.Post
	err := database.Posts.GetPosts(posts)
	if err != nil {
		return return500(c, "getPostsError", err)
	}
	return c.JSON(http.StatusOK, posts)
}

func createPostsHandler(c echo.Context) error {
	post := new(posts.Post)
	err := c.Bind(post)
	if err != nil {
		return return500(c, "createPostsParseError", err)
	}

	sess, err := database.Sessions.Get(c)
	if err != nil {
		return return500(c, "checkSessionDBError", err)
	}
	id := sess.Values["id"]
	if id == nil {
		return c.NoContent(http.StatusForbidden)
	}

	if post.CreatedUser != id {
		return c.NoContent(http.StatusForbidden)
	}

	err = database.Posts.AddPost(post)
	if err != nil {
		return return500(c, "createPostsError", err)
	}

	return c.NoContent(http.StatusOK)
}
