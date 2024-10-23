package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
// func Test_folder_GetFoldersByOrgID(t *testing.T) {
// 	t.Parallel()
// 	tests := [...]struct {
// 		name    string
// 		orgID   uuid.UUID
// 		folders []folder.Folder
// 		want    []folder.Folder
// 	}{
// 		// TODO: your tests here
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// f := folder.NewDriver(tt.folders)
// 			// get := f.GetFoldersByOrgID(tt.orgID)

// 		})
// 	}
// }

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	READ_ME_sample_folders := folder.GetSampleData("test_files/another_sample.json")
	tests := [...]struct {
		name    string		`json:"name"`
		orgID   uuid.UUID	`json:"org_id"`
		nodeName string		`json:"paths"`
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name: "README 1",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			nodeName: "alpha",
			folders: READ_ME_sample_folders,
			want: folder.GetSampleData("test_files/out/1.json"),
		},
		{
			name: "README 2",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			nodeName: "bravo",
			folders: READ_ME_sample_folders,
			want: folder.GetSampleData("test_files/out/2.json"),
		},
		{
			name: "README 3",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			nodeName: "charlie",
			folders: READ_ME_sample_folders,
			want: folder.GetSampleData("test_files/out/3.json"),
		},
		{
			name: "README 4",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			nodeName: "echo",
			folders: READ_ME_sample_folders,
			want: folder.GetSampleData("test_files/out/4.json"),
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetAllChildFolders(tt.orgID, tt.nodeName)
			if !equal(get, tt.want) {
				t.Errorf("got\n %v\nwant\n %v", get, tt.want)
			} 
		})
	}
}
func equal(a, b []folder.Folder) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i].Name != b[i].Name || 
           a[i].Paths != b[i].Paths || 
           a[i].OrgId != b[i].OrgId {
            return false
        }
    }
    return true
}

