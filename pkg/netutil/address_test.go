// Copyright 2022 TiKV Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package netutil

import (
	"crypto/tls"
	"net/http"
	"testing"

	. "github.com/pingcap/check"
)

func Test(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&testNetSuite{})

type testNetSuite struct{}

func (s *testNetSuite) TestResolveLoopBackAddr(c *C) {
	nodes := []struct {
		address     string
		backAddress string
	}{
		{address: "127.0.0.1:2379", backAddress: "192.168.130.22:10080"},
		{address: "0.0.0.0:2379", backAddress: "192.168.130.22:10080"},
		{address: "localhost:2379", backAddress: "192.168.130.22:10080"},
		{address: "192.168.130.22:2379", backAddress: "0.0.0.0:10080"},
	}

	for _, n := range nodes {
		c.Assert(ResolveLoopBackAddr(n.address, n.backAddress), Equals, "192.168.130.22:2379")
	}
}

func (s *testNetSuite) TestIsEnableHttps(c *C) {
	c.Assert(IsEnableHTTPS(http.DefaultClient), IsFalse)
	httpClient := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig:   nil,
		},
	}
	c.Assert(IsEnableHTTPS(httpClient), IsFalse)
	httpClient = &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig:   &tls.Config{},
		},
	}
	c.Assert(IsEnableHTTPS(httpClient), IsFalse)
}
