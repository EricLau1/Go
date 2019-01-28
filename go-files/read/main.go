package main

import(
  "bufio"
  "fmt"
  "log"
  "os"
)

func main() {
  txt, err := read("levels.txt")
  if err != nil {
    log.Fatal(err)
  }
  var content []string
  content = append(content, "package main\n")
  
  content = append(content, "import \"fmt\"")

  content = append(content, "type Level struct {")
  content = append(content, "\tId int")
  content = append(content, "\tValue int")
  content = append(content, "\tExp int")
  content = append(content, "\tTitle string")
  content = append(content, "}\n")

  content = append(content, "func main() {")
  var values string = "\tvalues := []int{"
  i := 0
  for i < len(txt) {
    if i < len(txt) - 1  {
      values += txt[i] + ", "
    } else {
      values += txt[i]
    }
     i++
  }
  values += "}"
  content = append(content, values)
  content = append(content, "\tfmt.Println(values)")
  content = append(content, "}")
  for _, line := range content {
    fmt.Println(line)
  }

  err = write(content, "levels.go")
  if err != nil {
    log.Fatalf("Error", err)
  }
  fmt.Println("Arquivo criado com sucesso!")
}

func read(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()
  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func write(lines []string, path string) error {
  file , err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()
  reader := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(reader, line)
  }
  return reader.Flush()
}
