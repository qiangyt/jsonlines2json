package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/qiangyt/jsonlines2json/util"
)

// JSONLine ...
type JSONLine = map[string]interface{}

// ParseRawLine ...
func ParseRawLine(lineNo int, rawLine string) JSONLine {
	r := make(JSONLine)

	line := strings.TrimSpace(rawLine)
	if len(line) == 0 {
		log.Printf("line %d is blank\n", lineNo)
		return nil
	}

	posOfLeftBracket := strings.IndexByte(line, '{')
	if posOfLeftBracket < 0 {
		log.Printf("line %d is not JSON line: <%s>\n", lineNo, rawLine)
		return nil
	}
	if posOfLeftBracket > 0 {
		line = line[posOfLeftBracket:]
	}

	if err := json.Unmarshal([]byte(line), &r); err != nil {
		line = strings.ReplaceAll(line, "\\\"", "\"")
		if err := json.Unmarshal([]byte(line), &r); err != nil {
			log.Printf("failed to parse line %d: <%s>\n\treason %v\n", lineNo, rawLine, errors.Wrap(err, ""))
			return r
		}
	}

	return r
}

// ProcessRawLine ...
func ProcessRawLine(first bool, lineNo int, rawLine string) bool {
	jl := ParseRawLine(lineNo, rawLine)
	if jl == nil {
		return false
	}

	if len(jl) == 0 {
		return false
	}

	output, err := json.Marshal(jl)
	if err != nil {
		log.Printf("failed to marshal line %d: <%s>\n\treason %v\n", lineNo, rawLine, errors.Wrap(err, ""))
		return false
	}

	if !first {
		fmt.Print(", ")
	}

	fmt.Println(string(output))
	return true
}

// ProcessLocalFile ...
func ProcessLocalFile(localFilePath string) {
	f, err := os.Open(localFilePath)
	if err != nil {
		panic(errors.Wrap(err, ""))
	}
	log.Printf("file is opened: %s\n", localFilePath)
	defer f.Close()

	processReader(f)
}

// processReader ...
func processReader(reader io.Reader) {

	buf := bufio.NewReader(reader)

	fmt.Print("[")

	first := true

	for lineNo := 1; true; lineNo++ {
		rawLine, err := buf.ReadString('\n')
		len := len(rawLine)

		if len != 0 {
			// trim the tail \n
			if rawLine[len-1] == '\n' {
				rawLine = rawLine[:len-1]
			}
		}

		if err != nil {
			if err == io.EOF {
				log.Printf("got EOF, line %d\n", lineNo)
				ProcessRawLine(first, lineNo, rawLine)
				fmt.Print("]")
				return
			}
			panic(errors.Wrapf(err, "failed to read line %d", lineNo))
		}

		ok := ProcessRawLine(first, lineNo, rawLine)
		if ok {
			first = false
		}
	}
}

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[]")
			os.Exit(1)
			return
		}
	}()

	ok, cmdLine := ParseCommandLine()
	if !ok {
		return
	}

	logFile := util.InitLogger()
	defer logFile.Close()

	ProcessLocalFile(cmdLine.LogFilePath)
}
