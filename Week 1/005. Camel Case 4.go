package main
import (
  "fmt"
  "bufio"
  "os"
  "unicode"
  "strings"
  )

func main() {
  reader := bufio.NewReader(os.Stdin)

  for {
    input, _, err := reader.ReadLine()
    if err != nil {
      return
    }
    text := string(input)
    
    operation, outType := text[0], text[2]
    var formattedWord string = text[4:]
    var outputString string
    
    if operation == 'S' {
      if outType == 'M' {
        formattedWord = formattedWord[:len(text)-6]
        for i, letter := range formattedWord {
            if unicode.IsUpper(letter) {
                outputString += " " + strings.ToLower(string(formattedWord[i]))
                continue
            }
            outputString += string(letter)
        }
      } else if outType == 'C' {
        for i, letter := range formattedWord {
            if unicode.IsUpper(letter) && i != 0 {
                outputString += " " + string(formattedWord[i])
                continue
            }
            outputString += string(letter)
        }
        outputString = strings.ToLower(outputString)
      } else if outType == 'V' {
        for i, letter := range formattedWord {
            if unicode.IsUpper(letter) {
                outputString += " " + strings.ToLower(string(formattedWord[i]))
                continue
            }
            outputString += string(letter)
        }
      }
      
    } else if operation == 'C' {
      if outType == 'M' {
        formattedWord = strings.Title(formattedWord)
        for _, letter := range formattedWord {
            if letter == ' ' {
                continue
            }
            outputString += string(letter)
        }
        outputRune := []rune(outputString)
        outputRune[0] = unicode.ToLower(outputRune[0])
        outputString = string(outputRune)

        methodBtn := "()"
        outputString += methodBtn
      } else if outType == 'C' {
        formattedWord = strings.Title(formattedWord)
        for _, letter := range formattedWord {
            if letter == ' ' {
                continue
            }
            outputString += string(letter)
        }
      } else if outType == 'V' {
        formattedWord = strings.Title(formattedWord)
        for _, letter := range formattedWord {
            if letter == ' ' {
                continue
            }
            outputString += string(letter)
        }
        outputRune := []rune(outputString)
        outputRune[0] = unicode.ToLower(outputRune[0])
        outputString = string(outputRune)
      }
    }
    
    fmt.Println(outputString)
  }
}
