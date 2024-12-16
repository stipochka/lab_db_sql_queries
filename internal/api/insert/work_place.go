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

func InsertIntoWorkPlace(db *database.Queries) error {
	file, err := os.Open("/home/stepa/lab_db_sql_queries/values_work_place.txt")
	if err != nil {
		return err
	}

	read := bufio.NewScanner(file)

	for read.Scan() {
		line := read.Text()

		fields := strings.Split(line, "\t")

		if len(fields) != 4 {
			fmt.Println("Skipping invalid line")
		}

		id, _ := strconv.Atoi(fields[0])
		institution := fields[1]
		address := fields[2]
		per := fields[3]

		params := database.CreateWorkplaceParams{
			ID:                    int32(id),
			Institution:           institution,
			Address:               address,
			Localbudgetpercentage: per,
		}

		_, err := db.CreateWorkplace(context.Background(), params)
		if err != nil {
			return err
		}
	}
	return nil
}
