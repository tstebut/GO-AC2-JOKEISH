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

// Joke joke
type Joke struct {
	ID    int    `json:"id" binding:"required"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke" binding:"required"`
}

/** we'll create a list of jokes */
var jokes = []Joke{
	Joke{1, 0, "Il y a 10 catégories de personnes: ceux qui comprennent le binaire...et les autres!"},
	Joke{2, 0, "Que dit une mère qui sert à dîner à son fils geek ? – ALT TAB"},
	Joke{3, 0, "Un homme met à jour le mot de passe de son ordinateur. Il tape « monpenis ». Sur l’écran, le message suivant s’affiche : « Erreur. Trop court. »"},
	Joke{4, 0, "Chuck Norris a tweeté une fois. Les serveurs ont crashés."},
	Joke{5, 0, "De nos jours, le zip ça devient rar(e)."},
	Joke{6, 0, "Linux a un noyau, Windows a des pépins."},
	Joke{7, 0, "Windows a détecté que vous n’aviez pas de clavier. Appuyez sur ‘F9′ pour continuer."},
	Joke{8, 0, "Quel animal a le plus de mémoire ? C’est l’agneau car il a une mémoire d’au moins 2 gigots."},
	Joke{9, 0, "Papa, dis-moi comment je suis né. \n – Très bien, il fallait bien que l’on en parle un jour ! Papa et maman se sont copier/coller dans un Chat sur MSN. Papa a fixé un rancard via e-mail à maman et se sont retrouvés dans les toilettes d’un cybercafé. Après, maman a fait quelques downloads du memory stick de papa. Lorsque papa fut prêt pour l’upload, nous avons découvert que nous n’avions pas mis de firewall. Il était trop tard pour faire delete, neuf mois plus tard le satané virus apparaissait."},
	Joke{10, 0, "Microsoft va encore s’enrichir et viendra supplanter toutes les autres … Effectivement ils vont déposer une demande de licence pour chacun de ses bugs"},
	Joke{11, 0, "Il parait qu’il existe une distribution Linux buguée, pour les nostalgiques de Windows."},
	Joke{12, 0, "La seule raison pour laquelle il est utile d’utiliser Windows, c’est pour tester un virus."},
}

var jwtMiddleWare *jwtmiddleware.JWTMiddleware

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Auth handler
	auth := new(auth.Auth)
	auth.Init()

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
