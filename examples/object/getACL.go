package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"net/http"

	"github.com/mozillazg/go-cos"
	"github.com/mozillazg/go-cos/examples"
)

func main() {
	u, _ := url.Parse("https://test-1253846586.cn-north.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &examples.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	name := "test/hello.txt"
	v, _, err := c.Object.GetACL(context.Background(), name)
	if err != nil {
		panic(err)
	}
	for _, a := range v.AccessControlList {
		fmt.Printf("%s, %s, %s\n", a.Grantee.Type, a.Grantee.UIN, a.Permission)
	}

}
