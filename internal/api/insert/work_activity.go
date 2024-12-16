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

func InsertIntoWorkActivity(db *database.Queries) error {
	file, err := os.Open("/home/stepa/lab_db_sql_queries/values_work_activity.txt")
	if err != nil {
		return err
	}

	read := bufio.NewScanner(file)

	for read.Scan() {
		line := read.Text()

		fields := strings.Split(line, "\t")

		if len(fields) != 7 {
			fmt.Println("Skipping invalid line")
		}

		contract, _ := strconv.Atoi(fields[0])
		date := fields[1]
		medpersonalID, _ := strconv.Atoi(fields[2])
		workplaceID, _ := strconv.Atoi(fields[3])
		operationID, _ := strconv.Atoi(fields[4])
		quantity, _ := strconv.Atoi(fields[5])
		payment := fields[6]

		params := database.CreateWorkActivityParams{
			Contract:      int32(contract),
			Date:          date,
			Medpersonalid: int32(medpersonalID),
			Workplaceid:   int32(workplaceID),
			Operationid:   int32(operationID),
			Quantity:      int32(quantity),
			Payment:       payment,
		}

		_, err := db.CreateWorkActivity(context.Background(), params)
		if err != nil {
			return err
		}
	}
	return nil
}
