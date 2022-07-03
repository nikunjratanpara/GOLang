package apis

import (
	"fmt"

	"plex.com/ingest"
)

func main() {
	message := ingest.Hello("Test")
	fmt.Println(message)
}
