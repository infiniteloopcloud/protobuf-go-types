## Fork of [Protobuf Golang](https://github.com/protocolbuffers/protobuf-go) specialized for type generation

This package supposed to focus on type generation based on proto file.

### Features compared to original repository

- Flags to REMOVE protobuf specific field generation
- Make difference between optional and non-optional in case of structs. So it's generating non-pointer in case of non-optional

### Installation

```shell
go install github.com/infiniteloopcloud/protoc-gen-go-types@latest
```

### Flags

These are actually environment variables.

- `SKIP_PROTOBUF_SPECIFIC=false` - Skip protobuf specific code
- `TYPE_OVERRIDE` - Enable type override

### Type override

Supported overwrite:

- `TimeTime` -> `time.Time`
- `RepeatedString` -> `[]string`

#### Example

```protobuf
syntax = "proto3";

message SomeStruct {
  TimeTime created_at = 1; // This will be time.Time
}

message TimeTime {}
```

#### Important

Currently the overwritten FieldOptions (go_type, go_import, go_import_alias, zero_override) must be paired with these numbers:
go_type = 1001
go_import = 1002
go_import_alias = 1003
zero_override = 1004
Because `protoc` can't process the extended options, so we can't find the by name, just by place.
