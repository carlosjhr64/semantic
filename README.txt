package semantic // import "github.com/carlosjhr64/semantic"

Way to do semantic versioning.
const VERSION string = "2.0.0"
var Upgraded = false
var Warn = true
func Cmp(x, y string) int
func Less(x, y string) bool
func Like(version string, i ...int) bool
func Likes(version string, match string)
func MNBC(version string) (int, int, int, string)
func MustLike(version string, pkg string, i ...int)
