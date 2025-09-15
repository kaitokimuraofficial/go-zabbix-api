package zabbix_test

import (
	"testing"

	zapi "github.com/claranet/go-zabbix-api"
)

func testCreateGlobalMacro(t *testing.T) *zapi.GlobalMacro {
	globalmacros := zapi.GlobalMacros{{
		MacroName:   "{$TEST_MACRO}",
		Value:       "test_value",
		Description: "Test macro",
	}}
	err := testGetAPI(t).GlobalMacrosCreate(globalmacros)
	if err != nil {
		t.Fatal(err)
	}
	return &globalmacros[0]
}

func testDeleteGlobalMacro(globalmacro *zapi.GlobalMacro, t *testing.T) {
	err := testGetAPI(t).GlobalMacrosDelete(zapi.GlobalMacros{*globalmacro})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGlobalMacros(t *testing.T) {
	api := testGetAPI(t)

	globalmacro := testCreateGlobalMacro(t)

	_, err := api.GlobalMacroGetByID(globalmacro.MacroID)
	if err != nil {
		t.Fatal(err)
	}

	globalmacro.Value = "another_value"
	err = api.GlobalMacrosUpdate(zapi.GlobalMacros{*globalmacro})
	if err != nil {
		t.Error(err)
	}

	testDeleteGlobalMacro(globalmacro, t)
}

