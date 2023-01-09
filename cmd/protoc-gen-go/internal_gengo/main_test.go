package internal_gengo

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/infiniteloopcloud/protoc-gen-go-types/compiler/protogen"
	"github.com/infiniteloopcloud/protoc-gen-go-types/parser"
)

func TestGenerateFile(t *testing.T) {
	t.Setenv("TYPE_OVERRIDE", "true")
	gen, err := parser.Parse("google/protobuf/descriptor.proto", "./test_data/config.proto", "./test_data/test.proto")
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range gen.Files {
		if f.Generate {
			content, err := GenerateFile(gen, f).Content()
			if err != nil {
				t.Fatal(err)
			}
			f, err := os.Create("./test_data/" + f.GeneratedFilenamePrefix + ".pb.go")
			if err != nil {
				t.Fatal(err)
			}
			io.Copy(f, bytes.NewReader(content))
		}
	}
}

func TestBuildTags(t *testing.T) {
	tags := buildTags("json=something,omitempty;validate=date")
	if len(tags) != 2 {
		t.Fatal("invalid tag count, must be 2")
	}
	if tags[0][0] != "json" {
		t.Errorf("Invalid tag name, should be 'json', instead of %q", tags[0][0])
	}
	if tags[0][1] != "something,omitempty" {
		t.Errorf("Invalid tag name, should be 'something,omitempty', instead of %q", tags[0][1])
	}
	if tags[1][0] != "validate" {
		t.Errorf("Invalid tag name, should be 'validate', instead of %q", tags[1][0])
	}
	if tags[1][1] != "date" {
		t.Errorf("Invalid tag name, should be 'date', instead of %q", tags[1][1])
	}
}

func TestGetAdditionalTags(t *testing.T) {
	TypeOverride = true
	overrideFields = map[string]map[string]overrideParams{
		"TestStruct": {
			"TestField": overrideParams{
				goStructTags: `json=id_dont_know,omitempty;boil=donno;validate=true`,
			},
		},
	}
	tags := getAdditionalTags(&messageInfo{
		Message: &protogen.Message{
			GoIdent: protogen.GoIdent{
				GoName: "TestStruct",
			},
		},
	}, &protogen.Field{
		GoName: "TestField",
	})

	if len(tags) != 3 {
		t.Fatal("Tags length must be 3")
	}
	if tags[0][0] != "json" {
		t.Errorf("First tag should be `json`, instead of %s", tags[0][0])
	}
	if tags[0][1] != "id_dont_know,omitempty" {
		t.Errorf("First tag should be `id_dont_know,omitempty`, instead of %s", tags[0][1])
	}
	if tags[1][0] != "boil" {
		t.Errorf("First tag should be `boil`, instead of %s", tags[1][0])
	}
	if tags[1][1] != "donno" {
		t.Errorf("First tag should be `donno`, instead of %s", tags[1][1])
	}
	if tags[2][0] != "validate" {
		t.Errorf("First tag should be `validate`, instead of %s", tags[2][0])
	}
	if tags[2][1] != "true" {
		t.Errorf("First tag should be `true`, instead of %s", tags[2][1])
	}
}
