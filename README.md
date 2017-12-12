# AdjAdjAnimal

A super-simple package to generate GfyCat-style slugs.

No other features are provided right now, as the goal is to just have an efficient thread-safe generator.

## Usage

```golang
import aaa github.com/serjlee/adjadjanimal
...

id := aaa.New()

fmt.Println(id) // prints something like "ArdentCuriousCrocodile"
```

