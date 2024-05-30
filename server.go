package main

import (
	"database/sql"
	"fintech-app/graph"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	// Подключение к бд
	connStr := "user=postgres password=qwerty123 dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Проверка соединения с базой данных
	err = db.Ping()
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	log.Println("Успешное подключение к базе данных")

	// Чтение SQL-запроса из файла
	sqlFilePath := filepath.Join("storage", "entities.sql")
	sqlFile, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		log.Fatalf("Ошибка чтения файла SQL: %v", err)
	}

	// Выполнение SQL-запроса
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		log.Fatalf("Ошибка выполнения SQL-запроса: %v", err)
	}
	log.Println("Таблицы успешно созданы")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := graph.NewResolver(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
