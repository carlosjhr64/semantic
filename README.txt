package semantic // import "github.com/carlosjhr64/semantic"

Way to do semantic versioning.
const VERSION Version = "1.0.0.alpha"
var Warn = true
func Cmp(x, y Versioner) int
func Less(x, y Versioner) bool
func Like(x Versioner, i ...int) bool
func MustLike(x Versioner, pkg string, i ...int)
type Version string
type Versioner interface { ... }
