// cmd/artificialmasterpiecegeneratorhubultra/main.go
package main

import (
"flag"
"log"
"os"

"artificialmasterpiecegeneratorhubultra/internal/artificialmasterpiecegeneratorhubultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialmasterpiecegeneratorhubultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
