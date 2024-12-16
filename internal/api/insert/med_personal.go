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

func InsertIntoMedPersonal(db *database.Queries) error {
	file, err := os.Open("/home/stepa/lab_db_sql_queries/values_medpersonal.txt")
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
		name := fields[1]
		address := fields[2]
		tax := fields[3]

		params := database.CreateMedPersonalParams{
			ID:            int32(id),
			Lastname:      name,
			Address:       address,
			Taxpercentage: tax,
		}

		_, err := db.CreateMedPersonal(context.Background(), params)
		if err != nil {
			return err
		}
	}
	return nil
}
