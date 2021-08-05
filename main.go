package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	fileName := "game_score_log.csv"

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	// ヘッダーの書き込み
	header := []string{"player_id", "score"}
	writer.Write(header)

	rowCount := 100
	for i := 0; i < rowCount; i++ {
		rand.Seed(time.Now().UnixNano())

		playerId := strconv.Itoa(rand.Intn(15))
		score := strconv.Itoa(rand.Intn(100) * 100)

		// 行の書き込み
		row := []string{playerId, score}
		writer.Write(row)
	}

	writer.Flush()

	log.Printf("generated %v", fileName)
}
