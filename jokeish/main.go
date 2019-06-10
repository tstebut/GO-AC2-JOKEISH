package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Response Response
type Response struct {
	Message string `json:"message"`
}

// Joke joke
type Joke struct {
	ID   int    `json:"id" binding:"required"`
	Joke string `json:"joke" binding:"required"`
}

/** we'll create a list of jokes */
var jokes = []Joke{
	Joke{1, "Il y a 10 catégories de personnes: ceux qui comprennent le binaire...et les autres!"},
	Joke{2, "Que dit une mère qui sert à dîner à son fils geek ? – ALT TAB"},
	Joke{3, "Un homme met à jour le mot de passe de son ordinateur. Il tape « monpenis ». Sur l’écran, le message suivant s’affiche : « Erreur. Trop court. »"},
	Joke{4, "Chuck Norris a tweeté une fois. Les serveurs ont crashés."},
	Joke{5, "De nos jours, le zip ça devient rar(e)."},
	Joke{6, "Linux a un noyau, Windows a des pépins."},
	Joke{7, "Windows a détecté que vous n’aviez pas de clavier. Appuyez sur ‘F9′ pour continuer."},
	Joke{8, "Quel animal a le plus de mémoire ? C’est l’agneau car il a une mémoire d’au moins 2 gigots."},
	Joke{9, "Papa, dis-moi comment je suis né. \n – Très bien, il fallait bien que l’on en parle un jour ! Papa et maman se sont copier/coller dans un Chat sur MSN. Papa a fixé un rancard via e-mail à maman et se sont retrouvés dans les toilettes d’un cybercafé. Après, maman a fait quelques downloads du memory stick de papa. Lorsque papa fut prêt pour l’upload, nous avons découvert que nous n’avions pas mis de firewall. Il était trop tard pour faire delete, neuf mois plus tard le satané virus apparaissait."},
	Joke{10, "Microsoft va encore s’enrichir et viendra supplanter toutes les autres … Effectivement ils vont déposer une demande de licence pour chacun de ses bugs"},
	Joke{11, "Il parait qu’il existe une distribution Linux buguée, pour les nostalgiques de Windows."},
	Joke{12, "La seule raison pour laquelle il est utile d’utiliser Windows, c’est pour tester un virus."},
}

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
		api.GET("/jokes", JokeHandler)
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
