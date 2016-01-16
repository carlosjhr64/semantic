package semantic

import "testing"
import "github.com/carlosjhr64/to"

func TestMNBC(test *testing.T) {
  bad := test.Error
  var m, n, b int
  var c string

  m, n, b, c = MNBC("1.2.3.a.b.c")

  if m != 1 { bad("Major") }
  if n != 2 { bad("Minor") }
  if b != 3 { bad("Build") }
  if c != "a.b.c" { bad("Note or Comment") }

  m, n, b, c = MNBC("1.2.3")

  if m != 1 { bad("Major") }
  if n != 2 { bad("Minor") }
  if b != 3 { bad("Build") }
  if c != "" { bad("Note or Comment") }
}

func TestCmp(test *testing.T) {
  bad := test.Error
  var i int

  i = Cmp("1.2.3", "1.2.3")
  if i!= 0 { bad("1.2.3<=>1.2.3") }

  // Major

  i = Cmp("1.2.3", "0.2.3")
  if i!= 3 { bad("1.2.3<=>0.2.3") }

  i = Cmp("1.2.3", "2.2.3")
  if i!= -3 { bad("1.2.3<=>2.2.3") }

  // Minor

  i = Cmp("1.2.3", "1.0.3")
  if i!= 2 { bad("1.2.3<=>1.0.3") }

  i = Cmp("1.2.3", "1.5.3")
  if i!= -2 { bad("1.2.3<=>1.5.3") }

  // Build

  i = Cmp("1.2.3", "1.2.2")
  if i!= 1 { bad("1.2.3<=>1.2.2") }

  i = Cmp("1.2.3", "1.2.7")
  if i!= -1 { bad("1.2.3<=>1.2.7") }
}

func TestLess(test *testing.T) {
  bad := test.Error
  a := "1.2.3.alpha"
  b := "1.2.3.beta"

  if Less(a,b) { bad("A. 1.2.3 == 1.2.3") }
  if Less(b,a) { bad("B. 1.2.3 == 1.2.3") }

  b = "2.2.3.beta"
  if !Less(a,b) { bad("C. a < b") }
  if Less(b,a) { bad("D. b > a") }

  b = "1.3.3.beta"
  if !Less(a,b) { bad("E. a < b") }
  if Less(b,a) { bad("F. b > a") }

  b = "1.2.4.beta"
  if !Less(a,b) { bad("G. a < b") }
  if Less(b,a) { bad("H. b > a") }

  b = "0.99.9.zzz"
  if Less(a,b) { bad("I. a > b") }
  if !Less(b,a) { bad("J. b < a") }

  b = "9.0.0"
  if !Less(a,b) { bad("K. a < b") }
  if Less(b,a) { bad("L. b > a") }
}

func TestLike(test *testing.T) {
  bad := test.Error
  a := "1.2.3.beta"

  if !Like(a,1) { bad("But a is 1!") }
  if !Like(a, 1, 2) { bad("But a is 1.2!") }
  if !Like(a, 1, 2, 3) { bad("But a is 1.2.3!") }

  if !Like(a, 1, 2, 2) { bad("But 1.2.3 > 1.2.2.") }
  if Like(a, 1, 2, 4) { bad("But 1.2.3 < 1.2.4.") }
  if !Like(a, 1, 1, 4) { bad("But 1.2.3 > 1.1.4.") }

  if Like(a, 1, 3) { bad("But 1.2 < 1.3.") }
  if !Like(a, 1, 1) { bad("But 1.2 > 1.1.") }
}

func TestVERSION(test *testing.T) {
  bad := test.Error
  if !Like(to.VERSION, 1, 0, 0) { bad("Unexpected to.VERSION") }
  if Cmp(VERSION, "2.0.0") != 0 { bad("Expected to be version 2.0.0") }
}

func TestLikes(test *testing.T){
  Likes("1.2.3.20160115", "pkg-1.2.3")
}
