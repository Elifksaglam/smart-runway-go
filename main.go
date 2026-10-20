
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/redis/go-redis/v9"
)

func main() {
    // Read configuration from .env
    data, err := os.ReadFile(".env")
    if err != nil {
        log.Fatal("Cannot read .env: ", err)
    }

    for _, line := range strings.Split(string(data), "\n") {
        line = strings.TrimSpace(line)

        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }

        key, value, ok := strings.Cut(line, "=")
        if ok {
            os.Setenv(strings.TrimSpace(key), strings.TrimSpace(value))
        }
    }

    // PostgreSQL connection
    db, err := pgxpool.New(
        context.Background(),
        os.Getenv("DATABASE_URL"),
    )
    if err != nil {
        log.Fatal("PostgreSQL error: ", err)
    }
    defer db.Close()

    // Valkey connection
    cache := redis.NewClient(&redis.Options{
        Addr: os.Getenv("VALKEY_ADDR"),
    })
    defer cache.Close()

    // Main page
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")

        fmt.Fprint(w, `
            <html>
            <head><title>Smart Runway</title></head>
            <body style="font-family:Arial;text-align:center;margin-top:100px">
                <h1>Smart Runway</h1>
                <h2>Go Backend Running!</h2>
                <p>PostgreSQL + Valkey</p>
                <a href="/health">Check Database Connections</a>
            </body>
            </html>
        `)
    })

    // Database health checks
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
        defer cancel()

        status := map[string]string{
            "postgresql": "connected",
            "valkey":     "connected",
        }

        if err := db.Ping(ctx); err != nil {
            status["postgresql"] = "disconnected"
            log.Println("PostgreSQL:", err)
        }

        if err := cache.Ping(ctx).Err(); err != nil {
            status["valkey"] = "disconnected"
            log.Println("Valkey:", err)
        }

        w.Header().Set("Content-Type", "application/json")

        if status["postgresql"] != "connected" ||
            status["valkey"] != "connected" {
            w.WriteHeader(http.StatusServiceUnavailable)
        }

        json.NewEncoder(w).Encode(status)
    })

    log.Println("Smart Runway running at http://localhost:8080")
    log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
