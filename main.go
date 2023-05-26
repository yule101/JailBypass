package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/frida/frida-go/frida"
)

func getBunldID() string {
	bID := flag.String("f", "", "./exp -f com.app.example")
	flag.Parse()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: ")
		flag.PrintDefaults()
	}
	if *bID == "" {
		flag.Usage()
		os.Exit(1)
	}

	return *bID
}

func getJSFile() string {
	jsFile, err := ioutil.ReadFile("jailbrekon.js")
	if err != nil {
		fmt.Println("open js file failed.")
		os.Exit(1)
	}
	jsCode := string(jsFile)
	return jsCode
}

func main() {
	appName := getBunldID()
	fmt.Println(frida.Version())
	dev := frida.USBDevice()
	fmt.Println(dev.Name())

	pid, err := dev.Spawn(appName, nil)

	if err != nil {
		fmt.Println("spawn process failed.")
		panic(err)
	}
	fmt.Println("pid: ", pid)
	session, err := dev.Attach(pid, nil)
	if err != nil {
		fmt.Println("attach failed!")
		panic(err)
	}
	s, err := session.CreateScript(getJSFile())
	if err != nil {
		panic(err)
	}

	s.On("message", func(message string, data []byte) {
		fmt.Println("Script message received:", message)
	})
	if err := s.Load(); err != nil {
		fmt.Println("script load failed! ")
	}
	dev.Resume(pid)
	r := bufio.NewReader(os.Stdin)
	r.ReadLine()
}
