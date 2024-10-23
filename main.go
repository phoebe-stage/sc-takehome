package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	// res := folder.GetAllFolders()
	data := folder.GetSampleData("test_files/another_sample.json")
	// res := folder.GetAllFolders()


	// example usage

	// folderDriver := folder.NewDriver(res)
	folderDriver := folder.NewDriver(data)
	orgFolder := folderDriver.GetFoldersByOrgID(orgID)
	// folder.PrettyPrint(data)
	fmt.Printf("\n Folders for orgID: %s", orgID)
	folder.PrettyPrint(orgFolder)
	new_folder := folderDriver.GetAllChildFolders(orgID, "nearby-secret")
	folder.PrettyPrint(new_folder)

}
