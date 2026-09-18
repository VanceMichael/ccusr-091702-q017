package main

import (
    "database/sql"
    "fmt"
    "os"
    "path/filepath"
    _ "modernc.org/sqlite"
)

func main() {
    path := os.Getenv("DATABASE_PATH")
    if path == "" { path = filepath.FromSlash("data/visitor_dispute.sqlite3") }
    if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil { panic(err) }
    migration, err := os.ReadFile("migrations/001_init.sql"); if err != nil { panic(err) }
    database, err := sql.Open("sqlite", path); if err != nil { panic(err) }
    defer database.Close()
    if _, err = database.Exec(string(migration)); err != nil { panic(err) }
    fmt.Println(path)
}