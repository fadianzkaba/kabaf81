package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	userID := 10
	val, err := fetchUserData(ctx, userID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result: ", val)
	fmt.Println("took: ", time.Since(start))
	fmt.Println(ctx.Value(start))

}

func fetchUserData(ctx context.Context, userID int) (int, error) {

	val, err := fetchThirdStuffWhichCanBeSlow()

	if err != nil {
		return 0, err
	}

	return val, err
}

func fetchThirdStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Microsecond * 500)
	return 666, nil
}
