//Created by Daniel Z. Dias de Moraes - 19 Jul 2019
//Open to be used, modified, changed, copied...It is just as PoC afterall.
//Made just to test a functionality that were intented to be added on a much
//bigger project. The purpose was to check a Temp folder and delete only the
//non persistent files, it was needed to determine those, and delete of course.
//The deletion part isn't included.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Pay attention on the json tag, that will be used by the unmarshal
//to automatically fit the right field of the json on the right field
//of the struct
type fileList struct {
	Files    []string `json:"files"`
	TargPath string   `json:"targetPath"`
}

//Main
func main() {

	//Open the json file and reads it
	jsonFile, err := os.Open("filelist.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValueJson, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()

	//Declares a struct of type fileList then unmarshal the
	//json file inside it
	var configfile fileList
	json.Unmarshal(byteValueJson, &configfile)

	//Gives to the exclusionList the declared fileList and get back results
	toBeDeleted, inFolderFileList := exclusionList(configfile)

	//Print results
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

//This function compares the files inside the target folder with those in the list
//then is assembles another list with results, which is returned.
func exclusionList(f fileList) (finalList []string, inTargetFiles []os.FileInfo) {
	target := f.TargPath
	tFolder, _ := os.Open(target)
	defer tFolder.Close()
	inTargetFiles, _ = ioutil.ReadDir(string(tFolder.Name()))

	finalList = []string{}
	for _, file := range inTargetFiles {
		if !(checkFile(string(file.Name()), f.Files)) {
			finalList = append(finalList, string(file.Name()))
		}
	}

	return finalList, inTargetFiles
}

//Check if the name of the file is on the list, probably not the most efficient way
//to do this, but works fast on a short list. Returns true if the file is on the list,
//and false if it's not.
func checkFile(fileName string, fileList []string) bool {
	for _, f := range fileList {
		if f == fileName {
			return true
		}
	}
	return false
}
