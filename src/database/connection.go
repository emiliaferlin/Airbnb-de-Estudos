package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

const (
	DatabaseName = "match_estudos"
	URI          = "mongodb://admin:admin123@localhost:27017/?authSource=admin"
)

// Connect abre a conexão com o MongoDB e armazena o cliente globalmente.
func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB não respondeu ao ping: %v", err)
	}

	Client = client
	DB = client.Database(DatabaseName)
	fmt.Println("Conectado ao MongoDB com sucesso!")
}

// GetCollection retorna uma coleção pelo nome.
func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}

// Disconnect encerra a conexão com o banco.
func Disconnect() {
	if err := Client.Disconnect(context.Background()); err != nil {
		log.Printf("Erro ao desconectar do MongoDB: %v", err)
	}
}
