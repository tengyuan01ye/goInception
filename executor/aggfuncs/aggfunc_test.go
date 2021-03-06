// Copyright 2018 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package aggfuncs_test

import (
	"testing"
	"time"

	"github.com/hanchuanchuan/goInception/parser"
	"github.com/hanchuanchuan/goInception/sessionctx"
	"github.com/hanchuanchuan/goInception/util/mock"
	"github.com/hanchuanchuan/goInception/util/testleak"
	. "github.com/pingcap/check"
)

var _ = Suite(&testSuite{})

func TestT(t *testing.T) {
	CustomVerboseFlag = true
	TestingT(t)
}

type testSuite struct {
	*parser.Parser
	ctx sessionctx.Context
}

func (s *testSuite) SetUpSuite(c *C) {
	s.Parser = parser.New()
	s.ctx = mock.NewContext()
	s.ctx.GetSessionVars().StmtCtx.TimeZone = time.Local
}

func (s *testSuite) TearDownSuite(c *C) {
}

func (s *testSuite) SetUpTest(c *C) {
	s.ctx.GetSessionVars().PlanColumnID = 0
	testleak.BeforeTest()
}

func (s *testSuite) TearDownTest(c *C) {
	s.ctx.GetSessionVars().StmtCtx.SetWarnings(nil)
	testleak.AfterTest(c)()
}
