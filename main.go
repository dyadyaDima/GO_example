package main

import (
    "./src/media"
    "./src/middleware"
    "./src/stream"
    "fmt"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "github.com/rs/cors"
    "log"
    "net/http"
    "os"
)

func main() {

    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    r := mux.NewRouter()


    // BEGIN ROUTES
    upload := r.PathPrefix("/upload").Subrouter()
    upload.HandleFunc("/media", media.Main).Methods("POST")
    upload.HandleFunc("/stream", stream.Main).Methods("POST")
    upload.Use(middleware.MyMiddleware)

    // END ROUTES

    corsOpts := cors.New(cors.Options{
        AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
        AllowedMethods: []string{
            http.MethodGet,//http methods for your app
            http.MethodPost,
            http.MethodPut,
            http.MethodPatch,
            http.MethodDelete,
            http.MethodOptions,
            http.MethodHead,
        },

        AllowedHeaders: []string{
            "Access-Control-Allow-Origin", os.Getenv("ACCESS_CONTROL_ALLOW_ORIGIN"),
        },
    })

    fmt.Println("Server is listening...")
    http.ListenAndServe(os.Getenv("LISTEN_IP")+":"+os.Getenv("LISTEN_PORT"), corsOpts.Handler(r))
}
