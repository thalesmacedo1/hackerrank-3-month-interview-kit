package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'flippingBits' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER n as parameter.
 */

func flippingBits(n int64) int64 {
    // transforma o numero em binario
    binaryN := strconv.FormatInt(n, 2)
    // adiciona os zeros que faltam (32 - tamanho do numero)
    nZeroes := strings.Repeat("0", 32 - len(binaryN))
    fullBinaryN := nZeroes + binaryN
    // trocar 0 por 1 e 1 por 0
    var reversedBinaryN string 
    for i, _ := range fullBinaryN{
        if fullBinaryN[i] == '0' {
            reversedBinaryN += "1"
        } else if fullBinaryN[i] == '1' {
            reversedBinaryN += "0"
        }
    }
    // converter de binaio para inteiro
    finalReversedN, _ := strconv.ParseInt(reversedBinaryN, 2, 64)
    // retornar inteiro
    return finalReversedN
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)

        result := flippingBits(n)

        fmt.Fprintf(writer, "%d\n", result)
    }

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
