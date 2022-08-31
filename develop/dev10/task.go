package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

var (
	timeOut time.Duration
)

type Client struct {
	address string
	timeout time.Duration
	conn    net.Conn
}

func InitClient(addr string, timeOut time.Duration) Client {
	return Client{
		address: addr,
		timeout: timeOut,
	}
}

func (c *Client) Open() error {
	if c.conn == nil {
		conn, err := net.DialTimeout("tcp", c.address, c.timeout)
		if err != nil {
			return err
		}
		c.conn = conn
	} else {
		fmt.Println("already connected to: ", c.address)
	}
	return nil
}

func (c *Client) Close() error {
	if c.conn != nil {
		err := c.conn.Close()
		if err != nil {
			return err
		}
	} else {
		fmt.Println("no any connection!")
	}
	return nil
}

func (c *Client) Write(b []byte) (int, error) {
	if c.conn == nil {
		fmt.Println("no any connection!")
		return 0, nil
	}
	return c.conn.Write(b)
}

func (c *Client) Read(b []byte) (int, error) {
	if c.conn == nil {
		fmt.Println("no any connection!")
		return 0, nil
	}
	return c.conn.Read(b)
}

func startRead(errChan chan error, telnet *Client) {
	_, err := io.Copy(os.Stdout, telnet)
	if err != nil {
		errChan <- err
		return
	}
	fmt.Println("received", os.Stdin)
}

func startWrite(errChan chan error, telnet *Client) {
	x, err := io.Copy(telnet, os.Stdin)
	if err != nil {
		errChan <- err
		return
	}
	fmt.Println("send", x)
}

func main() {
	flag.DurationVar(&timeOut, "timeout", 10*time.Second, "Connection timeout")
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("error: need host & port!")
	}
	addr := args[0] + ":" + args[1]
	telnet := InitClient(addr, timeOut)
	err := telnet.Open()
	if err != nil {
		log.Fatal(err)
	}
	errChan := make(chan error, 1)
	go startRead(errChan, &telnet)
	go startWrite(errChan, &telnet)
	ctrlD := make(chan os.Signal, 1)
	signal.Notify(ctrlD, os.Interrupt)
	select {
	case err := <-errChan:
		_ = telnet.Close()
		log.Fatal(err)
	case <-ctrlD:
		_ = telnet.Close()
		fmt.Println("Ctrl+D")
	}
}
