package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Création du routeur par défaut (avec logger et recovery middleware)
	r := gin.Default()

	// Définition de la route avec un paramètre nommé ":pattern"
	r.GET("/:pattern", func(c *gin.Context) {
		// Extraction du paramètre directement par son nom
		pattern := c.Param("pattern")

		// Réponse en texte brut (ou HTML)
		c.String(http.StatusOK, "Hello, %s!", pattern)
	})

	// Lancement du serveur sur le port 8080
	r.Run(":8080")
}