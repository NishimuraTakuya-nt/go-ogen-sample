package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/NishimuraTakuya-nt/go-ogen-sample/middlewares"
	petstore "github.com/NishimuraTakuya-nt/go-ogen-sample/petstore"
	"go.uber.org/zap"
)

func main() {
	// Create service instance.
	h := &handler{
		pets: map[int64]petstore.Pet{},
	}

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	// Create generated server.
	srv, err := petstore.NewServer(
		h,
		petstore.WithMiddleware(middlewares.Logging(logger)),
		petstore.WithErrorHandler(func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
			// エラーが発生した場合、ここで処理される
			log.Printf("Error occurred: %v", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error handler": err.Error(),
			})
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8180", srv); err != nil {
		log.Fatal(err)
	}
}
