package files

import (
	"fmt"
	"io"
	"net/http"
)

const Url = "http://services.explorecalifornia.org/json/tours.php"

func RequestFunc() {
	fmt.Println("Network requests")
	client := http.Client{}
	req, err := http.NewRequest("GET", Url, nil)
	CheckError(err)

	req.Header.Set("User-Agent", "")

	resp, err := client.Do(req)
	CheckError(err)
	defer resp.Body.Close()

	fmt.Printf("Response type:%T\n", resp)

	bytes, err := io.ReadAll(resp.Body)
	CheckError(err)
	content := string(bytes)
	fmt.Print(content)
}
