package insert

import (
	"bufio"
	"context"
	"fmt"
	"lab_db_sql_queries/internal/database"
	"os"
	"strconv"
	"strings"
)

func InsertIntoOperationTypes(db *database.Queries) error {
	file, err := os.Open("/home/stepa/lab_db_sql_queries/values_operation_types.txt")
	if err != nil {
		return err
	}

	read := bufio.NewScanner(file)

	for read.Scan() {
		line := read.Text()

		fields := strings.Split(line, "\t")
		if len(fields) != 5 {
			fmt.Println("Skipping invalid line")
		}

		id, _ := strconv.Atoi(fields[0])
		name := fields[1]
		basePoint := fields[2]
		stock, _ := strconv.Atoi(fields[3])
		cost := fields[3]

		params := database.CreateOperationTypesParams{
			ID:        int32(id),
			Name:      name,
			Basepoint: basePoint,
			Stock:     int32(stock),
			Cost:      cost,
		}

		_, err := db.CreateOperationTypes(context.Background(), params)
		if err != nil {
			return err
		}
	}
	return nil
}
