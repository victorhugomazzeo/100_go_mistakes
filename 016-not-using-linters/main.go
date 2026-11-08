// Mistake #16: Not using linters
//
// go vet and gofmt don't catch everything: linters like shadow and errcheck
// find bugs that compile fine. Run them, ideally in CI or a pre-commit hook.
//
// Tools that ship with Go:
//   - compiler:      rejects unused variables and imports (always an error)
//   - go vet ./...:  finds code that compiles but is likely wrong
//     (Printf args, unused results, self-assignment...)
//   - go tool vet help: lists every check go vet runs
//   - gofmt -l .:    lists files that are not formatted
//   - gofmt -d file: shows what gofmt would change
//   - gofmt -w file: rewrites the file formatted
//   - go fmt ./...:  runs gofmt -l -w on every package
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func mistake() {
	fmt.Printf("mistake: age: %d\n", "42") // go vet: format %d has arg "42" of wrong type string
	fmt.Println("mistake: age: %d", 42)    // go vet: Println call has possible Printf formatting directive %d

}

func fix() {
	fmt.Printf("fix: age: %s\n", "42")
	fmt.Println("fix: age:", 42)
}
