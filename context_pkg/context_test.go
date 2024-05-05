package contextpkg

import (
	"fmt"
	"context"
	"testing"
)

func TestContext(t *testing.T) {

	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}