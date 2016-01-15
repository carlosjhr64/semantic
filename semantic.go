// Way to do semantic versioning.
package semantic

import "strings"
import "github.com/carlosjhr64/to"

type Versioner interface {
  MNBC() (int, int, int, string)
}

type Version string

var VERSION Version = "0.0.0.alpha"

func (v Version) MNBC() (int, int, int, string) {
  a := strings.SplitN(string(v), ".", 4)
  if len(a) < 3 { panic("Version number not in m.n.b form.") }
  major := to.Int(a[0])
  minor := to.Int(a[1])
  build := to.Int(a[2])
  note := "" // or comment
  if len(a) == 4 { note = a[3] }
  return major, minor, build, note
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

func Less(x, y Versioner) bool {
  return Cmp(x, y) < 0
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
