package main

import (
	"fmt"
	"os/exec"
	"os"
	"log"
	"bytes"
	"io"
	"math/rand"
	"time"
	"bufio"
	"strings"
	
	"github.com/claudiu/gocron"
	"github.com/bregydoc/gtranslate"
)

func main() {
	gocron.Start()
	s := gocron.NewScheduler()
	gocron.Every(10).Minute().Do(sendWord)
	<-s.Start()
}

func sendWord() {
	inputBox()
	word := read()
	translated := translate("en", "tr", word)
	notification(word, translated)
}

func inputBox() {
	out, _ := exec.Command("zenity", "--entry", "--text", "Please Enter a English Word").Output()
	write(string(out))
}

func notification(word, translate string) {
	cmd := exec.Command("notify-send", "-i", "stock_lock-broken ", word, translate)
	cmd.Run()
}

func translate(source, target, text string) (translated string){
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: source,
			To:   target,
		},
	)
	if err != nil {
		panic(err)
	}
	return translated
}

func read() string{
	lastLine := fileEndOf() + 1 
	rand.Seed(time.Now().UnixNano())
	lineNumber := rand.Intn( lastLine - 2) + 1

	file, err := os.Open("./words")
	if err != nil {
        log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
        lineNumber++
        if lastLine == lineNumber {
            return sc.Text()
        }
	}
	return ""
}

func write(data string) {
    f, err := os.OpenFile("./words", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return
	}
    _, err = fmt.Fprintln(f, strings.TrimSuffix(data, "\n"))
    if err != nil {
        fmt.Println(err)
                f.Close()
        return
    }
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}

func fileEndOf() int {
	file, err := os.Open("./words")
	if err != nil {
        log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 32*1024)
    count := 0
    lineSep := []byte{'\n'}

    for {
        c, err := file.Read(buf)
        count += bytes.Count(buf[:c], lineSep)

        switch {
		case err == io.EOF:
            return count
    	}

	}
}
