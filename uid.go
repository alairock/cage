package cage

import (
	"fmt"
	"go.jetpack.io/typeid"
)

func GenUID(prefix string) string {
	tid := typeid.Must(typeid.New(prefix))
	return fmt.Sprintf("%s", tid)
}
