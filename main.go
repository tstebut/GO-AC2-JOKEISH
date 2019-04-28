package main

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"./auth"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Response Response
type Response struct {
	Message string `json:"message"`
}

// Jwks Jwks
type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

// JSONWebKeys JSONWebKeys
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// Joke joke
type Joke struct {
	ID    int    `json:"id" binding:"required"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke" binding:"required"`
}

/** we'll create a list of jokes */
var jokes = []Joke{
	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	Joke{2, 0, "What do you call a fake noodle? An Impasta."},
	Joke{3, 0, "How many apples grow on a tree? All of them."},
	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
	Joke{8, 0, "Dad, did you get a haircut? No I got them all cut."},
	Joke{9, 0, "What do you call a Mexican who has lost his car? Carlos."},
	Joke{10, 0, "Dad, can you put my shoes on? No, I don't think they'll fit me."},
	Joke{11, 0, "Why did the scarecrow win an award? Because he was outstanding in his field."},
	Joke{12, 0, "Why don't skeletons ever go trick or treating? Because they have no body to go with."},
}

var jwtMiddleWare *jwtmiddleware.JWTMiddleware

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve the frontend
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.GET("/jokes", auth.AuthMiddleware(), JokeHandler)
		api.POST("/jokes/like/:jokeID", auth.AuthMiddleware(), LikeJoke)
	}
	// Start the app
	router.Run(":" + os.Getenv("LISTEN_PORT"))
}

// JokeHandler returns a list of jokes available (in memory)
func JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, jokes)
}

// LikeJoke to perform likes
func LikeJoke(c *gin.Context) {
	// Check joke ID is valid
	if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		// find joke and increment likes
		for i := 0; i < len(jokes); i++ {
			if jokes[i].ID == jokeid {
				jokes[i].Likes = jokes[i].Likes + 1
			}
		}
		c.JSON(http.StatusOK, &jokes)
	} else {
		// the jokes ID is invalid
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// getJokesByID returns a single joke
func getJokesByID(id int) (*Joke, error) {
	for _, joke := range jokes {
		if joke.ID == id {
			return &joke, nil
		}
	}
	return nil, errors.New("Joke not found")
}
