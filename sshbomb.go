package main

import (
	"flag"
	"log"
	"net"
	"os"
	"time"
)

var (
	advertFlag = flag.Bool("advertise", true, "advertise sshd version by displaying banner")
	bannerFlag = flag.String("banner", "SSH-2.0-OpenSSH_5.3", "sshd banner to present")
	listenFlag = flag.String("listen", ":2222", "address to listen on")
	sizeFlag   = flag.Int("size", 1024*1024, "size in bytes of data to send")
)

func handle(c net.Conn) {
	defer c.Close()

	log.Println("connection from", c.RemoteAddr())

	c.SetDeadline(time.Now().Add(10 * time.Second))

	f, err := os.Open("/dev/urandom")
	if err != nil {
		log.Println("can't open:", err)
		return
	}

	data := make([]byte, *sizeFlag)
	rsize, err := f.Read(data)
	if err != nil {
		log.Println("can't read:", err)
		return
	}

	if *advertFlag {
		_, err = c.Write([]byte(*bannerFlag + "\n"))
		if err != nil {
			log.Printf("can't write banner:", err)
			return
		}
	}

	wsize, err := c.Write(data)
	if err != nil {
		log.Printf("write error (%d of %d bytes were written): %s", wsize, rsize, err.Error())
		return
	}
}

func main() {
	flag.Parse()

	log.Println("*** listening on", *listenFlag)

	l, err := net.Listen("tcp", *listenFlag)
	if err != nil {
		log.Fatalln("can't listen:", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("can't accept:", err)
			continue
		}

		go handle(conn)
	}
}
