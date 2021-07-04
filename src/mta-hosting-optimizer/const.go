package hostingoptimizer

import "os"

const (
	StInMemory = "inmemory"
	StMongo    = "mongo"

	StMysql = "mysql"

	IpHostnameCollection = "ipconfig"

	IpHostNameTable = "ipconfig"

	//Defaults to use
	DEFAULT_EXTENSION        = "jpeg"
	DEFAULT_RESOURCE         = "default"
	NOT_FOUND_RESOURCE       = "404notfound"
	DEFAULT_VARIATION_WIDTH  = 250
	DEFAULT_VARIATION_HEIGHT = 200

	//Define possible Image storage adapters
	ADAPTER_TYPE_LOCAL = "Local"
	ADAPTER_TYPE_S3    = "S3"

	//Misc
	ORIGINAL_IMAGE               = "original"
	IMAGE_ENCODE_FORMAT          = "base64"
	IMAGE_ENCODE_FORMAT_PLAIN    = "plain"
	MONGO_NOT_FOUND_MSG          = "not found"
	DISTANCE_IN_KILOMETER        = "k"
	VALIDATE_KILOMETER_ADD_SIGHT = 5
)

//Directory Separator
var DS string = string(os.PathSeparator)
