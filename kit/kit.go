package kit

const (
	DevDirectory    = "dev"
	KitMetadataFile = "kit.yml"
)

type Kit struct {
	Name    string
	Version string
	IsDev   bool

	Summary  string
	Author   string
	Homepage string
	Github   string

	Vault map[string]interface{}
}
