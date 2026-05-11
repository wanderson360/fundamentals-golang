package main

import (
	"api-product/controller"
	usecase "api-product/useCase"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Carrega variáveis do .env (na raiz do projeto)
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Aviso: não foi possível carregar .env, usando variáveis de ambiente do sistema")
	}

	// Monta string de conexão
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	// Abre conexão com o banco
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	// Testa conexão
	if err := db.Ping(); err != nil {
		log.Fatalf("Banco não respondeu: %v", err)
	}

	// Inicializa servidor
	server := gin.Default()
	server.SetTrustedProxies(nil) // remove warning de proxy

	productUseCase := usecase.NewProductUseCase(db)
	productController := controller.NewProductController(productUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	server.GET("/products", productController.GetProducts)

	// Porta do servidor vinda do .env
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000" // valor padrão
	}
	log.Printf("Servidor rodando na porta %s 🚀", port)
	server.Run(":" + port)
}
