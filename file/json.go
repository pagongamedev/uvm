package file

import (
	"encoding/json"
	"os"
)

func ReadJSONFile(fileName string) map[string]interface{} {
	file, _ := os.Open(fileName)
	defer file.Close()

	data := map[string]interface{}{}
	json.NewDecoder(file).Decode(&data)
	return data
}

func WriteJSONFile(fileName string, data map[string]interface{}) bool {
	file, _ := os.OpenFile(fileName, os.O_CREATE, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(data)
	return true
}
