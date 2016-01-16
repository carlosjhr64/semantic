// Way to do semantic versioning.
package semantic

import "os"
import "fmt"
import "strings"
import "github.com/carlosjhr64/to"

const VERSION string = "2.0.0"

func mnbc(version *string) (int, int, int, string) {
  a := strings.SplitN(*version, ".", 4)
  if len(a) < 3 { to.Oops(version, "MNBC Version") }
  major := to.Int(&a[0])
  minor := to.Int(&a[1])
  build := to.Int(&a[2])
  comment := ""
  if len(a) == 4 { comment = a[3] }
  return major, minor, build, comment
}

func MNBC(version string) (int, int, int, string) {
  return mnbc(&version)
}

func cmp(x, y *string) int {
  mx, nx, bx, _ := mnbc(x)
  my, ny, by, _ := mnbc(y)
  if mx > my { return  3 }
  if mx < my { return -3 }
  if nx > ny { return  2 }
  if nx < ny { return -2 }
  if bx > by { return  1 }
  if bx < by { return -1 }
  return 0
}

func Cmp(x, y string) int {
  return cmp(&x, &y)
}

func Less(x, y string) bool {
  return cmp(&x, &y) < 0
}

var Upgraded = false
func like(version *string, i ...int) bool {
  Upgraded = false
  m, n, b, _ := mnbc(version)
  j := len(i)
  if j<1 || j>3 { panic("Expected 1 to 3 int.") }
  if m != i[0] { return false }    // Major differences not Like eachother.
  if j > 1 {
    if n < i[1] { return false }   // Must contain minor differences.
    if n == i[1] {
      if j > 2 {
        if b < i[2] { return false } // Must include some bug fix.
        if b > i[2] { Upgraded = true }
      }
    }else{ Upgraded = true}
  }
  return true
}
func Like(version string, i ...int) bool {
  return like(&version, i...)
}

var Warn = true
func mustLike(version *string, pkg *string, i ...int) {
  if !like(version, i...){
    msg := fmt.Sprintf("Did not Like %s %s.", *pkg, *version)
    if to.Panic {
      panic(msg)
    } else {
      // Something was found in an unconfigured or misconfigured state.
      fmt.Fprintln(os.Stderr, msg)
      os.Exit(78)
    }
  }
  // HACK!!! lol
  if Warn && Upgraded { fmt.Fprintf(os.Stderr, "Warning: %s upgraded.\n", *pkg) }
}
func MustLike(version string, pkg string, i ...int) {
  mustLike(&version, &pkg, i...)
}

// semantic.Likes(to.VERSION, "to-0.2.0")
func Likes(version string, match string) {
  pkgreq := strings.SplitN(match, "-", 2)
  pkg, req := pkgreq[0], pkgreq[1]
  w := strings.SplitN(req, ".", 3)
  i := make([]int, len(w))
  for j, n := range(w) { i[j] = to.Int(&n) }
  mustLike(&version, &pkg, i...)
}
