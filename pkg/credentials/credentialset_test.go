package credentials

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredentialSet(t *testing.T) {
	is := assert.New(t)
	if err := os.Setenv("TEST_USE_VAR", "kakapu"); err != nil {
		t.Fatal("could not setup env")
	}
	defer os.Unsetenv("TEST_USE_VAR")

	credset, err := Load("testdata/staging.yaml")
	if err != nil {
		//t.Fatal(err)
		t.Error(err)
	}

	results, err := credset.Resolve()
	if err != nil {
		t.Fatal(err)
	}
	count := 5
	is.Len(results, count, "Expected %d credentials", count)

	for _, tt := range []struct {
		name   string
		key    string
		expect string
		path   string
	}{
		{name: "run_program", key: "TEST_RUN_PROGRAM", expect: "wildebeest\n"},
		{name: "use_var", key: "TEST_USE_VAR", expect: "kakapu"},
		{name: "read_file", key: "TEST_READ_FILE", expect: "serval"},
		{name: "fallthrough", key: "TEST_FALLTHROUGH", expect: "quokka", path: "animals/quokka.txt"},
		{name: "plain_value", key: "TEST_PLAIN_VALUE", expect: "cassowary"},
	} {
		dest := results[tt.name]
		is.Equal(tt.key, dest.EnvVar)
		is.Equal(tt.expect, dest.Value)
		is.Equal(tt.path, dest.Path)
	}
}