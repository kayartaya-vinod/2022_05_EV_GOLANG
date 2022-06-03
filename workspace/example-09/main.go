package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var filename = "../example-08/main.go"

func main1() {
	data, err := os.ReadFile(filename)
	check(err)
	fmt.Println(string(data))
}

func main2() {
	file, err := os.Open(filename)
	check(err)

	loopCount := 0
	for {

		var bytes = make([]byte, 70)
		_, err := file.Read(bytes)
		if err != nil {
			break
		}
		loopCount++
		// fmt.Printf("\nRead %v bytes from the file\n", n)
		fmt.Print(string(bytes))
	}
	fmt.Println("No. of iterations =", loopCount)
	file.Close()
}

func main3() {
	html := `<html>
		<head>
			<title>Go is coll</title>
		</head>
		<body>
			<h1>Hello, world!</h1>
		</body>
		</html>`

	err := os.WriteFile("hello.html", []byte(html), 0640)
	check(err)
	fmt.Println("File hello.html created!")
}

func main4() {
	f, err := os.Create("names.txt")
	check(err)

	for {
		var name string
		fmt.Print(`Enter a firstname or "quit" to stop: `)
		fmt.Scanf("%s", &name)
		if name == "quit" {
			break
		}

		f.WriteString(name + "\n")
	}
	f.Close()
	fmt.Println("File names.txt created with few names")
}

func main5() {
	f, err := os.OpenFile("names.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)

	for {
		var name string
		fmt.Print(`Enter a firstname or "quit" to stop: `)
		fmt.Scanf("%s", &name)
		if name == "quit" {
			break
		}

		_, err := f.WriteString(name + "\n")

		check(err)
	}
	f.Close()
}

func main() {
	f, err := os.OpenFile("names.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	w := bufio.NewWriter(f)
	for {
		var name string
		fmt.Print(`Enter a firstname or "quit" to stop: `)
		fmt.Scanf("%s", &name)
		if name == "quit" {
			break
		}

		_, err := w.WriteString(name + "\n")
		w.Flush() // writes any buffered data to the underlying io.Writer.
		check(err)
	}

	f.Close()
}
