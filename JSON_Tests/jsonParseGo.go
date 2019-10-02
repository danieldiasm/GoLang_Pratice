package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Pay attention on the json tag, that will be used by the unmarshall
//to automatically fit the right field of the json on the right field
//of the struct
type fileList struct {
	Files    []string `json:"files"`
	TargPath string   `json:"targetPath"`
}

func main() {
	jsonFile, err := os.Open("filelist.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValueJson, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()

	var configfile fileList

	json.Unmarshal(byteValueJson, &configfile)

	toBeDeleted, inFolderFileList := exclusionList(configfile)

	fmt.Println("There are those files inside the Target Folder:")
	for _, inFile := range inFolderFileList {
		fmt.Println(inFile.Name())
	}
	fmt.Println("--------------------")
	fmt.Println("There was some files listed on JSON to be saved from damnation:")
	for _, File := range configfile.Files {
		fmt.Println(File)
	}
	fmt.Println("--------------------")
	fmt.Println("There are the file list of which are going to be destoyed for EVER and EVER:")
	for _, delFile := range toBeDeleted {
		fmt.Println(delFile)
	}

	fmt.Println("\nAnd...That is it.")
}

func exclusionList(f fileList) (finalList []string, inTargetFiles []os.FileInfo) {
	target := f.TargPath
	tFolder, _ := os.Open(target)
	defer tFolder.Close()
	inTargetFiles, _ = ioutil.ReadDir(string(tFolder.Name()))
	//FIX THIS CRAP
	finalList = []string{}
	for _, file := range inTargetFiles {
		if !(checkFile(string(file.Name()), f.Files)) {
			finalList = append(finalList, string(file.Name()))
		}
	}

	return finalList, inTargetFiles
}

//Check if the name of the file is on the list, not the most efficient way, yet.
//Just for a PoC
//Returns true if the file is on the list, and false if it's not.
func checkFile(fileName string, fileList []string) bool {
	for _, f := range fileList {
		if f == fileName {
			return true
		}
	}
	return false
}
