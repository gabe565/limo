package main

// Reset data dir
//go:generate rm -rf data
//go:generate mkdir data
// Run Goose migrations
//go:generate goose -dir=internal/migrations sqlite3 data/limo.db up
// Create SQLBoiler models
//go:generate sqlboiler sqlite3
