package command

import (
	"testing"
)

func TestBuildOptionsValidate(t *testing.T) {
	bf := new(BuildOptions)

	err := bf.Validate()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	// Both set
	bf.Except = make([]string, 1)
	bf.Only = make([]string, 1)
	err = bf.Validate()
	if err == nil {
		t.Fatal("should error")
	}

	// One set
	bf.Except = make([]string, 1)
	bf.Only = make([]string, 0)
	err = bf.Validate()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	bf.Except = make([]string, 0)
	bf.Only = make([]string, 1)
	err = bf.Validate()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestBuildOptionsValidate_userVarFiles(t *testing.T) {
	bf := new(BuildOptions)

	err := bf.Validate()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	// Non-existent file
	bf.UserVarFiles = []string{"ireallyshouldntexistanywhere"}
	err = bf.Validate()
	if err == nil {
		t.Fatal("should error")
	}
}
