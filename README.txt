package semantic // import "github.com/carlosjhr64/semantic"

Way to do semantic versioning.
const VERSION Version = "0.1.0.alpha"
func Cmp(x, y Versioner) int
func Less(x, y Versioner) bool
func Like(x Versioner, i ...int) bool
func MNBC(v string) (int, int, int, string)
type Version string
type Versioner interface { ... }
