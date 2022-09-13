package main
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

/*
 * Complete the 'findMedian' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func findMedian(arr[] int32) int32 {
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })

    var medianIndex int

    if len(arr) % 2 == 0 {
        medianIndex = (len(arr) / 2) - 1
    } else {
        medianIndex = len(arr) / 2
    }

    median: = arr[medianIndex]
    return median
}

func main() {
    reader: = bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
    stdout,
    err: = os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)
    defer stdout.Close()
    writer: = bufio.NewWriterSize(stdout, 16 * 1024 * 1024)
    nTemp,
    err: = strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10,
        64)
    checkError(err)
    n: = int32(nTemp)
    arrTemp: = strings.Split(strings.TrimSpace(readLine(reader)), " ")
    var arr[] int32
    for i: = 0;i < int(n);i++{
        arrItemTemp, err: = strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem: = int32(arrItemTemp)
        arr = append(arr, arrItem)
    }
    result: = findMedian(arr)
    fmt.Fprintf(writer, "%d\n", result)
    writer.Flush()
}

func readLine(reader * bufio.Reader) string {
    str, _, err: = reader.ReadLine()
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