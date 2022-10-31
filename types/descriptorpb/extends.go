package descriptorpb

import protoimpl "github.com/infiniteloopcloud/protoc-gen-go-types/runtime/protoimpl"

// GetExtensionFields returns the extension fields of the field options
// this function is needed when we want to read the field options, passed by protoc
func (x *FieldOptions) GetExtensionFields() protoimpl.ExtensionFields {
	if x != nil {
		return x.extensionFields
	}
	return nil
}
