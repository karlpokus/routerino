package routerino

import "testing"

func TestRoutes(t *testing.T) {
  Test(t, []Table{
    {
      "hi",
      nil,
      hi(),
      200,
      []byte("hi"),
    },
  })
}
