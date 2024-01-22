package config

// Config is the configuration of the application loaded from the json file
type Config struct {
	// WordleURL is the URL to the wordle API
	WordleURL string `json:"wordle_url,required"`
	// Port is the port where the application will run
	Port int `json:"port,required"`
	// DictionaryURL is the URL to the dictionary API
	DictionaryURL string `json:"dictionary_url,required"`
	// DatabasePath is the path to the database file
	DatabasePath string `json:"database_path,required"`
	// DatabaseFile is the database file
	DatabaseFile string `json:"database_file,required"`
}
