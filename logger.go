package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db, errDB        = sql.Open("sqlite3", "stats.db")
	clients          = make(chan string, 30)
	blocked          = make(chan string, 30)
	cacheClients     = make(map[string]int)
	cacheClientsLock = &sync.Mutex{}
	cacheBlocked     = make(map[string]int)
	cacheBlockedLock = &sync.Mutex{}
)

func loggerStart() {
	if errDB != nil {
		log.Fatal(errDB)
	}
	table, _ := db.Prepare("CREATE TABLE IF NOT EXISTS clients (ip TEXT PRIMARY KEY, queries INTEGER NOT NULL)")
	table.Exec()
	table, _ = db.Prepare("CREATE TABLE IF NOT EXISTS blocked (ip TEXT PRIMARY KEY, queries INTEGER NOT NULL)")
	table.Exec()

	go loggerClients()
	go loggerBlocked()
	go loggerStats()
}

func loggerEnd() {
	db.Close()
}

func loggerClients() {
	for {
		client := <-clients
		cacheClientsLock.Lock()
		cacheClients[client] = cacheClients[client] + 1
		cacheClientsLock.Unlock()
	}
}

func loggerBlocked() {
	for {
		badguy := <-blocked
		cacheBlockedLock.Lock()
		cacheBlocked[badguy] = cacheBlocked[badguy] + 1
		cacheBlockedLock.Unlock()
	}
}

func loggerStats() {
	for {
		loggerSync()
		podium := 1
		fmt.Println("---------------- TOP BLOCKED ----------------")
		row, _ := db.Query("SELECT * FROM blocked ORDER BY queries DESC LIMIT 5")
		defer row.Close()
		for row.Next() { // Iterate and fetch the records from result cursor
			var ip string
			var queries int
			row.Scan(&ip, &queries)
			fmt.Println("[", podium, "]", ip, "with", queries, "failed queries!")
			podium++
		}
		fmt.Println()
		time.Sleep(time.Duration(conf.LoggerPeriod) * time.Second)
	}
}

func loggerSync() {
	defer cacheClientsLock.Unlock()
	defer cacheBlockedLock.Unlock()
	cacheClientsLock.Lock()
	cacheBlockedLock.Lock()
	cursor, _ := db.Begin()
	for client, queries := range cacheClients {
		delete(cacheClients, client)
		res, _ := cursor.Exec("UPDATE clients SET queries=queries+? WHERE ip=?", queries, client)
		affected, _ := res.RowsAffected()
		if affected == 0 {
			cursor.Exec("INSERT INTO clients VALUES (?, 1)", client, queries)
		}
	}
	for badguy, queries := range cacheBlocked {
		delete(cacheBlocked, badguy)
		res, _ := cursor.Exec("UPDATE blocked SET queries=queries+? WHERE ip=?", queries, badguy)
		affected, _ := res.RowsAffected()
		if affected == 0 {
			cursor.Exec("INSERT INTO blocked VALUES (?, ?)", badguy, queries)
		}
	}
	cursor.Commit()
}
