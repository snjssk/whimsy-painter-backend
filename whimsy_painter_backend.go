package main

import (
	"context"
	"fmt"

	"github.com/snjssk/whimsy-painter-backend/app/models"
)

func main()  {
	ctx := context.Background()
	result := models.GetItem(ctx)
	fmt.Println("main", result)
}
