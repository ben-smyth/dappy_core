package yaml



type ApplicatonIndex struct {
	Applications []struct {
		Name   string `yaml:"name"`
		URL    string `yaml:"url"`
		Folder string `yaml:"folder"`
	} `yaml:"Applications"`
}


type Appliciation struct {
	Name string `yaml:"name"`
	URL string `yaml:"url"`
	Desc string `yaml:"desc"`
}
