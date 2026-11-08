# 100 Go Mistakes

Estudo do livro *100 Go Mistakes and How to Avoid Them* (Teiva Harsanyi).
Cada mistake vira uma pasta com um `main.go` executável que mostra o erro e a correção.

## Como rodar

```sh
go run ./001-unintended-variable-shadowing
```

## Progresso

| #   | Mistake                        | Pasta                                                                 |
| --- | ------------------------------ | --------------------------------------------------------------------- |
| 1   | Unintended variable shadowing  | [001-unintended-variable-shadowing](001-unintended-variable-shadowing) |
| 2   | Unnecessary nested code        | [002-unnecessary-nested-code](002-unnecessary-nested-code)             |
| 3   | Misusing init functions        | [003-misusing-init-functions](003-misusing-init-functions)             |
| 4   | Overusing getters and setters  | [004-overusing-getters-and-setters](004-overusing-getters-and-setters) |
| 5   | Interface pollution            | [005-interface-pollution](005-interface-pollution)                     |
| 6   | Interface on the producer side | [006-interface-on-the-producer-side](006-interface-on-the-producer-side) |
| 7   | Returning interfaces           | [007-returning-interfaces](007-returning-interfaces)                   |
| 8   | any says nothing               | [008-any-says-nothing](008-any-says-nothing)                           |
| 9   | Being confused about when to use generics | [009-being-confused-about-when-to-use-generics](009-being-confused-about-when-to-use-generics) |
| 10  | Not being aware of the possible problems with type embedding | [010-not-being-aware-of-the-possible-problems-with-type-embedding](010-not-being-aware-of-the-possible-problems-with-type-embedding) |
| 11  | Not using the functional options pattern | [011-not-using-the-functional-options-pattern](011-not-using-the-functional-options-pattern) |
| 12  | Project misorganization | [012-project-misorganization](012-project-misorganization) |
| 13  | Creating utility packages | [013-creating-utility-packages](013-creating-utility-packages) |
| 14  | Ignoring package name collisions | [014-ignoring-package-name-collisions](014-ignoring-package-name-collisions) |
| 15  | Missing code documentation | [015-missing-code-documentation](015-missing-code-documentation) |
| 16  | Not using linters | [016-not-using-linters](016-not-using-linters) |
| 17  | Creating confusion with octal literals | [017-creating-confusion-with-octal-literals](017-creating-confusion-with-octal-literals) |
| 18  | Neglecting integer overflows | [018-neglecting-integer-overflows](018-neglecting-integer-overflows) |
| 19  | Not understanding floating-points | [019-not-understanding-floating-points](019-not-understanding-floating-points) |
| 20  | Not understanding slice length and capacity | [020-not-understanding-slice-length-and-capacity](020-not-understanding-slice-length-and-capacity) |
| 21  | Inefficient slice initialization | [021-inefficient-slice-initialization](021-inefficient-slice-initialization) |
| 22  | Being confused about nil vs. empty slices | [022-being-confused-about-nil-vs-empty-slices](022-being-confused-about-nil-vs-empty-slices) |
| 23  | Not properly checking if a slice is empty | [023-not-properly-checking-if-a-slice-is-empty](023-not-properly-checking-if-a-slice-is-empty) |
| 24  | Not making slice copies correctly | [024-not-making-slice-copies-correctly](024-not-making-slice-copies-correctly) |
| 25  | Unexpected side effects using slice append | [025-unexpected-side-effects-using-slice-append](025-unexpected-side-effects-using-slice-append) |
| 26  | Slices and memory leaks | [026-slices-and-memory-leaks](026-slices-and-memory-leaks) |
| 27  | Inefficient map initialization | [027-inefficient-map-initialization](027-inefficient-map-initialization) |
| 28  | Maps and memory leaks | [028-maps-and-memory-leaks](028-maps-and-memory-leaks) |
| 29  | Comparing values incorrectly | [029-comparing-values-incorrectly](029-comparing-values-incorrectly) |
| 30  | Ignoring that elements are copied in range loops | [030-ignoring-that-elements-are-copied-in-range-loops](030-ignoring-that-elements-are-copied-in-range-loops) |
| 31  | Ignoring how arguments are evaluated in range loops | [031-ignoring-how-arguments-are-evaluated-in-range-loops](031-ignoring-how-arguments-are-evaluated-in-range-loops) |
| 32  | Ignoring the impacts of using pointer elements in range loops (obsoleto desde o Go 1.22) | [032-ignoring-the-impacts-of-using-pointer-elements-in-range-loops](032-ignoring-the-impacts-of-using-pointer-elements-in-range-loops) |
| 33  | Making wrong assumptions during map iterations | [033-making-wrong-assumptions-during-map-iterations](033-making-wrong-assumptions-during-map-iterations) |
| 34  | Ignoring how the break statement works | [034-ignoring-how-the-break-statement-works](034-ignoring-how-the-break-statement-works) |
| 35  | Using defer inside a loop | [035-using-defer-inside-a-loop](035-using-defer-inside-a-loop) |
| 36  | Not understanding the concept of rune | [036-not-understanding-the-concept-of-rune](036-not-understanding-the-concept-of-rune) |
