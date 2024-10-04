package main

// import (
// 	// "fmt"
// 	"fmt"
// 	"os"
// 	"os/exec"
// )




// func main() {
// 	db_url := os.Getenv("DATABASE_URL")
// 	cmd := exec.Command("goose -dir storage/migrations postgres \"" +  fmt.Sprint(db_url) +"\" up")
// 	stdout, err := cmd.Output()

//     if err != nil {
//         fmt.Println(err.Error())
//         return
//     }

//     // Print the output
//     fmt.Println(string(stdout))
// }