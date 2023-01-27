package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func getTerraformFileName(resourceNameLine string) (string, error) {
	newFileName := ""
	fileNameParts := strings.Split(resourceNameLine, " ")
	if len(fileNameParts) == 4 {
		newFileName = fileNameParts[0] + "." + strings.Replace(fileNameParts[1], "\"", "", -1) + "." + strings.Replace(fileNameParts[2], "\"", "", -1) + ".tf"
	} else {
		err := errors.New("Unable to parse file name from \"" + resourceNameLine + "\". Skipping...")
		return newFileName, err
	}
	return newFileName, nil
}

func writeStringToTerraformFile(terrformBlock string) error {

	// Create a file
	resourceNameLine := strings.Split(terrformBlock, "\n")[0]
	terraformFileName, err := getTerraformFileName(resourceNameLine)
	if err != nil {
		return err
	}
	file, err := os.Create(terraformFileName)
	if err != nil {
		return err
	}
	_, err = file.WriteString(terrformBlock)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func getTerraformFileNamesInCurrentDirectory() ([]string, error) {
	// Get a list of files in the current working directory
	currentdirectory, err := os.Open(".")
	if err != nil {
		return nil, err
	}

	defer currentdirectory.Close()
	files, err := currentdirectory.Readdir(-1)
	if err != nil {
		return nil, err
	}

	// iterate through the files and find the terraform file
	terraformFiles := []string{}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".tf") && !file.IsDir() {
			terraformFiles = append(terraformFiles, file.Name())
		}
	}

	return terraformFiles, nil
}

func checkError(err error) {
	if err != nil {
		// To-do: print this error in red
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func getTerraformBlocksFromFile(sourceTerraformFile *os.File) ([]string, error) {
	terraformBlocks := []string{}

	// Prepare the scanner and the string to hold the resource block
	fileScanner := bufio.NewScanner(sourceTerraformFile)
	fileScanner.Split(bufio.ScanLines)
	resourceBlock := ""

	// Iterate through the file line by line and find the resource block(s)
	for fileScanner.Scan() {
		if strings.HasPrefix(fileScanner.Text(), "resource") {
			// Start of the block
			resourceBlock = fileScanner.Text() + "\n"
			for fileScanner.Scan() {
				resourceBlock = resourceBlock + fileScanner.Text() + "\n"
				if fileScanner.Text() == "}" {
					// End of the blcok
					terraformBlocks = append(terraformBlocks, resourceBlock)
					break
				}
			}
		}
	}
	return terraformBlocks, nil
}

func main() {
	// Get a list of files in the current working directory
	sourceTerraformFileNames, err := getTerraformFileNamesInCurrentDirectory()
	checkError(err)

	// If there are no Terraform files in the current directory, exit
	if len(sourceTerraformFileNames) == 0 {
		fmt.Println("No terraform files found in the current directory. Exiting...")
		os.Exit(0)
	}

	// Iterate through the files and split Terraform files into individual resource blocks
	for _, sourceTerraformFileName := range sourceTerraformFileNames {
		// Open the file
		sourceTerraformFile, err := os.Open(sourceTerraformFileName)
		checkError(err)
		defer sourceTerraformFile.Close()

		// Get the resource blocks from the file
		terraformBlocks, err := getTerraformBlocksFromFile(sourceTerraformFile)
		checkError(err)

		// Write the resource blocks to individual files
		for _, terraformBlock := range terraformBlocks {
			err = writeStringToTerraformFile(terraformBlock)
			checkError(err)
		}

		sourceTerraformFile.Close()
	}
}
