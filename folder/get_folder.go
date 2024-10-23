package folder

import "github.com/gofrs/uuid"
import "strings"

func GetAllFolders() []Folder {
	return GetSampleData("sample.json")
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Your code here...
	folders := f.GetFoldersByOrgID(orgID)
	res := []Folder{}
	children :=[]string{}
	for _, f := range folders{
		path_parts := strings.Split(f.Paths, ".")
		for i := 0; i < len(path_parts); i++  {
			part := path_parts[i]
			if part == name && i!=(len(path_parts)-1){
				children = append(children, path_parts[len(path_parts)-1])
			}
		}

	}
	for _, f := range folders{
		for _, child := range children{
			if f.Name == child {
				res = append(res, f)
			}
		}
	}
	return res
}
