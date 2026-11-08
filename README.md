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
