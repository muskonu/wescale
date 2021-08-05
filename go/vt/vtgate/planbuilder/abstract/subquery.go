/*
Copyright 2021 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package abstract

import (
	"vitess.io/vitess/go/vt/sqlparser"
	"vitess.io/vitess/go/vt/vtgate/engine"
	"vitess.io/vitess/go/vt/vtgate/semantics"
)

// SubQuery stores the information about subquery
type SubQuery struct {
	Inner []*SubQueryInner
	Outer Operator
	Type  engine.PulloutOpcode
}

var _ Operator = (*SubQuery)(nil)

type SubQueryInner struct {
	Inner   Operator
	Type    engine.PulloutOpcode
	ArgName string
}

// TableID implements the Operator interface
func (s *SubQuery) TableID() semantics.TableSet {
	ts := s.Outer.TableID()
	for _, inner := range s.Inner {
		ts = ts.Merge(inner.Inner.TableID())
	}
	return ts
}

// PushPredicate implements the Operator interface
func (s *SubQuery) PushPredicate(expr sqlparser.Expr, semTable *semantics.SemTable) error {
	panic("implement me")
}
