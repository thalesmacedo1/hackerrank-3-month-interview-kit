package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/*
 * Complete the 'marsExploration' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func marsExploration(s string) int32 {
    changedLetters := 0
    
    for i := 0; i < len(s); i++ {
        if  (string(s[i]) == "S" && i % 3 == 0) ||
            (string(s[i]) == "O" && i % 3 == 1) ||
            (string(s[i]) == "S" && i % 3 == 2) {
                continue
        }
        changedLetters++
    }
    
    return int32(changedLetters)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := marsExploration(s)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
