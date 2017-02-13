package kit

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func (k *Kit) load(in io.Reader) error {
	var meta = struct {
		Name     string                 `yaml:"name"`
		Author   string                 `yaml:"author"`
		Homepage string                 `yaml:"homepage"`
		Github   string                 `yaml:"github"`
		Vault    map[string]interface{} `yaml:"vault"`
	}{}

	b, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, &meta)
	if err != nil {
		return err
	}

	k.Summary = meta.Name
	k.Author = meta.Author
	k.Homepage = meta.Author
	k.Github = meta.Github
	return nil
}
