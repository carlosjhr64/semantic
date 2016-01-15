// Way to do semantic versioning.
package semantic

import "os"
import "fmt"
import "strings"
import "github.com/carlosjhr64/to"

type Versioner interface {
  MNBC() (int, int, int, string)
}

type Version string
const VERSION Version = "1.1.0.alpha"

func (v Version) MNBC() (int, int, int, string) {
  return to.Version(v).MNBC()
}

func cmp(mx, nx, bx, my, ny, by int) int {
  if mx > my { return  3 }
  if mx < my { return -3 }
  if nx > ny { return  2 }
  if nx < ny { return -2 }
  if bx > by { return  1 }
  if bx < by { return -1 }
  return 0
}

func Cmp(x, y Versioner) int {
  mx, nx, bx, _ := x.MNBC()
  my, ny, by, _ := y.MNBC()
  return cmp(mx, nx, bx, my, ny, by)
}

func (x Version) Cmp(y Versioner) int {
  return Cmp(x, y)
}

func Less(x, y Versioner) bool {
  return Cmp(x, y) < 0
}

func (x Version) Less(y Versioner) bool {
  return Less(x, y)
}

var upgraded = false
func Like(x Versioner, i ...int) bool {
  upgraded = false
  m, n, b, _ := x.MNBC()
  j := len(i)
  if j<1 || j>3 { panic("Expected 1 to 3 arguments.") }
  if m != i[0] { return false }    // Major differences not Like eachother.
  if j > 1 {
    if n < i[1] { return false }   // Must contain minor differences.
    if n == i[1] {
      if j > 2 {
        if b < i[2] { return false } // Must include some bug fix.
        if b > i[2] { upgraded = true }
      }
    }else{ upgraded = true}
  }
  return true
}

func (x Version) Like(i ...int) bool {
  return Like(x, i...)
}

var Warn = true
func MustLike(x Versioner, pkg string, i ...int) {
  if !Like(x, i...){
    msg := fmt.Sprintf("Did not Like %s %s.", pkg, x)
    if to.Panic {
      panic(msg)
    } else {
      // Something was found in an unconfigured or misconfigured state.
      fmt.Fprintln(os.Stderr, msg)
      os.Exit(78)
    }
  }
  // HACK!!! lol
  if Warn && upgraded { fmt.Fprintf(os.Stderr, "Warning: %s upgraded.\n", pkg) }
}

func (x Version) MustLike(pkg string, i ...int) {
  MustLike(x, pkg, i...)
}

// semantic.Likes(to.VERSION, "to-0.2.0")
func Likes(version interface{}, match string) {
  pkgreq := strings.SplitN(match, "-", 2)
  pkg, req := pkgreq[0], pkgreq[1]
  w := strings.SplitN(req, ".", 3)
  i := make([]int, len(w))
  for j, n := range(w) { i[j] = to.Int(n) }
  // Just use the genious fmt.Sprintf function to handle anything.
  v := fmt.Sprintf("%v", version)
  Version(v).MustLike(pkg, i...)
}
