// Public Domain (-) 2013-2014 The GoPlivo Authors.
// See the GoPlivo UNLICENSE file for details.

/*
Package plivo is a client for the Plivo.com API. It is documented at:

http://godoc.org/github.com/micrypt/go-plivo/plivo

A short example of retrieving account details:

  func main() {
    client = plivo.NewClient(authID, authToken)
    acc, _, err := client.Account.Get()
    if err != nil {
      t.Errorf("AccountGet failed: %v", err)
    } else {
      t.Logf("Account: %v\n", acc)
    }
  }
*/
package plivo
