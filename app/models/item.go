package models

import (
	"context"
	"log"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
)

type Mybook struct {
	Id string
}

func GetItem(ctx context.Context) string {
	var result = "hello world"
	log.Println(DbConnection)

	cmd := `SELECT * FROM mybook`

	rows, err := DbConnection.QueryContext(ctx, cmd)
	if err != nil {
		log.Println("Error!")
		log.Fatalln(err)
	}
	defer rows.Close()

	// log.Println(rows)

	//var result []string
	//for rows.Next() {
	//	var i Items
	//	err := rows.Scan(&i.Id, &i.Name, &i.Price)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	v, _ := json.Marshal(i)
	//	result = append(result, string(v))
	//}

	return result
}
