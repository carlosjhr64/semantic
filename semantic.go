// Way to do semantic versioning.
package semantic

import "os"
import "fmt"
import "github.com/carlosjhr64/to"

type Versioner interface {
  MNBC() (int, int, int, string)
}

type Version string
const VERSION Version = "0.1.1.alpha"

func (v Version) MNBC() (int, int, int, string) {
  return to.Version(v).MNBC()
}

func Cmp(x, y Versioner) int {
  mx, nx, bx, _ := x.MNBC()
  my, ny, by, _ := y.MNBC()
  if mx > my { return  3 }
  if mx < my { return -3 }
  if nx > ny { return  2 }
  if nx < ny { return -2 }
  if bx > by { return  1 }
  if bx < by { return -1 }
  return 0
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

func Like(x Versioner, i ...int) bool {
  m, n, b, _ := x.MNBC()
  j := len(i)
  if j<1 || j>3 { panic("Expected 1 to 3 arguments.") }
  if m != i[0] { return false }    // Major differences not Like eachother.
  if j > 1 {
    if n < i[1] { return false }   // Must contain minor differences.
    if j > 2 && n == i[1] {
      if b < i[2] { return false } // Must include some bug fix.
    }
  }
  return true
}

func (x Version) Like(i ...int) bool {
  return Like(x, i...)
}

func MustLike(x Versioner, i ...int) {
  if !Like(x, i...){
    msg := fmt.Sprintf("Did not Like %T %s.", x, x)
    if to.Panic {
      panic(msg)
    } else {
      // Something was found in an unconfigured or misconfigured state.
      fmt.Fprintln(os.Stderr, msg)
      os.Exit(78)
    }
  }
}

func (x Version) MustLike(i ...int) {
  MustLike(x, i...)
}
