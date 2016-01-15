package semantic

import "testing"
import "fmt"
import "strings"

func TestMNBC(test *testing.T) {
  bad := test.Error
  var version Version
  var m, n, b int
  var c string

  version = Version("1.2.3.a.b.c")
  m, n, b, c = version.MNBC()

  if m != 1 { bad("Major") }
  if n != 2 { bad("Minor") }
  if b != 3 { bad("Build") }
  if c != "a.b.c" { bad("Note or Comment") }

  version = Version("1.2.3")
  m, n, b, c = version.MNBC()

  if m != 1 { bad("Major") }
  if n != 2 { bad("Minor") }
  if b != 3 { bad("Build") }
  if c != "" { bad("Note or Comment") }
}

func TestCMP(test *testing.T) {
  bad := test.Error
  var a, b Version
  var i int

  a = Version("1.2.3")

  b = Version("1.2.3")
  i = Cmp(a,b)
  if i!= 0 { bad("Cmp(a,b), 1.2.3<=>1.2.3") }
  i = a.Cmp(b) // and the method version works
  if i!= 0 { bad("Cmp(a,b), 1.2.3<=>1.2.3") }

  // Major

  b = Version("0.2.3")
  i = Cmp(a,b)
  if i!= 3 { bad("Cmp(a,b), 1.2.3<=>0.2.3") }

  b = Version("2.2.3")
  i = Cmp(a,b)
  if i!= -3 { bad("Cmp(a,b), 1.2.3<=>2.2.3") }

  // Minor

  b = Version("1.0.3")
  i = Cmp(a,b)
  if i!= 2 { bad("Cmp(a,b), 1.2.3<=>1.0.3") }

  b = Version("1.5.3")
  i = Cmp(a,b)
  if i!= -2 { bad("Cmp(a,b), 1.2.3<=>1.5.3") }

  // Build

  b = Version("1.2.2")
  i = Cmp(a,b)
  if i!= 1 { bad("Cmp(a,b), 1.2.3<=>1.2.2") }

  b = Version("1.2.7")
  i = Cmp(a,b)
  if i!= -1 { bad("Cmp(a,b), 1.2.3<=>1.2.7") }
}

func TestLess(test *testing.T) {
  bad := test.Error
  var a, b Version

  a = Version("1.2.3.alpha")
  b = Version("1.2.3.beta")

  if Less(a,b) { bad("A. 1.2.3 == 1.2.3") }
  if Less(b,a) { bad("B. 1.2.3 == 1.2.3") }
  if a.Less(b) { bad("B. 1.2.3 == 1.2.3") } // method version

  b = Version("2.2.3.beta")
  if !Less(a,b) { bad("C. a < b") }
  if Less(b,a) { bad("D. b > a") }

  b = Version("1.3.3.beta")
  if !Less(a,b) { bad("E. a < b") }
  if Less(b,a) { bad("F. b > a") }

  b = Version("1.2.4.beta")
  if !Less(a,b) { bad("G. a < b") }
  if Less(b,a) { bad("H. b > a") }

  b = Version("0.99.9.zzz")
  if Less(a,b) { bad("I. a > b") }
  if !Less(b,a) { bad("J. b < a") }

  b = Version("9.0.0")
  if !Less(a,b) { bad("K. a < b") }
  if Less(b,a) { bad("L. b > a") }
}

func TestLike(test *testing.T) {
  bad := test.Error
  a := Version("1.2.3.beta")

  if !Like(a,1) { bad("But a is 1!") }
  if !Like(a, 1, 2) { bad("But a is 1.2!") }
  if !Like(a, 1, 2, 3) { bad("But a is 1.2.3!") }
  if !a.Like(1, 2, 3) { bad("But a is 1.2.3!") } // method version

  if !Like(a, 1, 2, 2) { bad("But 1.2.3 > 1.2.2.") }
  if Like(a, 1, 2, 4) { bad("But 1.2.3 < 1.2.4.") }
  if !Like(a, 1, 1, 4) { bad("But 1.2.3 > 1.1.4.") }

  if Like(a, 1, 3) { bad("But 1.2 < 1.3.") }
  if !Like(a, 1, 1) { bad("But 1.2 > 1.1.") }
}

func TestPrint(test *testing.T) {
  version := Version("1.2.3")
  fmt.Println(version)
  fmt.Println(strings.Split(string(version), "."))
  fmt.Printf("VERSION: %s\n", version)
}
