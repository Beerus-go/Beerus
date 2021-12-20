package data_type

const (
	Bool   = "bool"
	Int    = "int"
	Uint   = "uint"
	Float  = "float"
	String = "string"
	Slice  = "slice"
	Struct = "struct"

	// BeeFile Beerus provides a struct to store files, if the request Content-Type is formData, then you can use this type to receive files from the client
	BeeFile = "beefile"

	BeeRequest  = "beerequest"
	BeeResponse = "beeresponse"
)
