package uuid

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func New() string {
	s := fmt.Sprintf("guid string: %v", rand.Int63())
	u, _ := uuid.NewV5(uuid.NamespaceURL, []byte(s))

	return u.String()
}
