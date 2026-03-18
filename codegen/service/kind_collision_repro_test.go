package service

import (
	"testing"

	"github.com/stretchr/testify/require"

	"goa.design/goa/v3/codegen"
)

func TestBuildUnionTypeDataKindConstsCollision(t *testing.T) {
	// Union "Foo" with variant "KindBar" -> KindName "FooKind", KindConst "FooKindKindBar".
	union1 := makeUnionForOrderTest("Foo", "KindBar")
	// Union "FooKind" with variant "Bar" -> KindName "FooKindKind", KindConst "FooKindKindBar" without uniquing.
	union2 := makeUnionForOrderTest("FooKind", "Bar")

	loc := &codegen.Location{RelImportPath: "gen/service"}
	scope := codegen.NewNameScope()

	data1 := buildUnionTypeData(union1, scope, loc)
	data2 := buildUnionTypeData(union2, scope, loc)

	require.Equal(t, "FooKindKindBar", data1.Fields[0].KindConst)
	require.NotEqual(t, data1.Fields[0].KindConst, data2.Fields[0].KindConst, "KindConst names must not collide")
	// Union "FooKind" gets renamed to "FooKind2" because "FooKind" was already reserved as a kind name.
	require.Equal(t, "FooKind2KindBar", data2.Fields[0].KindConst)
}
