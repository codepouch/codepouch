package main

import "os"
import "bufio"
import "strconv"
import "strings"
import "fmt"

func main() {
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    if err != nil {
        return
    }

    cases, err := strconv.Atoi(strings.TrimSpace(line))
    if err != nil {
        return
    }

    for i := 0; i < cases; i++ {
        line, err = reader.ReadString('\n')
        if err != nil {
            return
        }

        count, err := strconv.Atoi(strings.TrimSpace(line))
        if err != nil {
            return
        }

        line, err = reader.ReadString('\n')
        if err != nil {
            return
        }

        integers := strings.Fields(strings.TrimSpace(line))

        A := make([]int, count)
        for k := 0; k < count; k++ {
            A[k], err = strconv.Atoi(integers[k])
        }

        output := SingleNumber(A)
        fmt.Printf("%d\n", output)
    }
}
