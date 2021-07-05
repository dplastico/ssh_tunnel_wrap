package main
import (
	"fmt"
	"os/exec"
)
func main(){
	
	var conn string
	conn = "binario_test@tunnel.us-2.checkpoint.security" //CHANGE connection information
	var port string
	port = "7777" //CHANGE port to be redirected (localhost)

	c := exec.Command("ssh", "-i", "cp.pem", "-NL", port+":tunnel:1", conn)
	if err := c.Run(); err != nil {
		fmt.Println("Eroorrrrr ... ", err)
	}
}
