package GoStub_demo

import "io/ioutil"

// app.go

var (
	configFile = "config.json"
	maxNum     = 10
)

func GetConfig() ([]byte, error) {
	return ioutil.ReadFile(configFile)
}

func ShowNumber() int {
	// ...
	return maxNum
}
