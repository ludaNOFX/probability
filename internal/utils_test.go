package internal

import (
	"path/filepath"
	"strings"
	"testing"
)

func stringBuilder(sep byte, strs ...string) string {
	var builder strings.Builder
	buffer := 0
	for _, str := range strs {
		buffer += len(str)
		buffer += 1
	}
	builder.Grow(buffer)
	for idx, str := range strs {
		builder.WriteString(str)
		if idx != len(strs)-1 {
			builder.WriteByte(sep)
		}
	}
	return builder.String()
}

func Test_findRootDir(t *testing.T) {
	assertRootDir := "mat_stat"
	rootDir, err := FindRootDir(1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	prj_dir := filepath.Base(rootDir)
	if assertRootDir != prj_dir {
		t.Errorf("Expected %s, got %s", assertRootDir, prj_dir)
	}
}

func Test_BuildStoragePathOne(t *testing.T) {
	rootDir := "some_root_dir"
	prjName := "some_prj_name"
	storageType := "some_storage_type"
	assertRes := stringBuilder(
		'/',
		rootDir,
		BaseStorage,
		prjName,
		storageType,
	)
	res := BuildStoragePathOne(rootDir, prjName, storageType)
	if assertRes != res {
		t.Errorf("Expected %s, got %s", assertRes, res)
	}
}

func Test_BuildStoragePathMany(t *testing.T) {
	const defaultTestRootDir string = "back_door"
	tests := []struct {
		name         string
		prjMap       *PrjMap
		assertResult PathMap
	}{
		{
			name: "1",
			prjMap: NewPrjMap(
				"some_prj1",
				[]string{
					"csv",
					"plot",
				},
			),
			assertResult: PathMap{
				"csv": stringBuilder(
					'/',
					defaultTestRootDir,
					BaseStorage,
					"some_prj1",
					"csv",
				),
				"plot": stringBuilder(
					'/',
					defaultTestRootDir,
					BaseStorage,
					"some_prj1",
					"plot",
				),
			},
		},
		{
			name: "2",
			prjMap: NewPrjMap(
				"some_prj2",
				[]string{
					"csv",
					"plot",
					"logs",
					"jsons",
					"images",
				},
			),
			assertResult: PathMap{
				"csv": stringBuilder(
					'/',
					defaultTestRootDir,
					BaseStorage,
					"some_prj1",
					"csv",
				),
				"plot": stringBuilder(
					'/',
					defaultTestRootDir,
					BaseStorage,
					"some_prj1",
					"plot",
				),
				"logs": stringBuilder(
					'/',
					defaultTestRootDir,
					BaseStorage,
					"some_prj2",
					"logs",
				),
				"jsons": stringBuilder(
					'/',
					defaultTestRootDir,
					BaseStorage,
					"some_prj2",
					"jsons",
				),
				"images": stringBuilder(
					'/',
					defaultTestRootDir,
					BaseStorage,
					"some_prj2",
					"images",
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := BuildStoragePathMany(defaultTestRootDir, tt.prjMap)
			keys := make(map[string]struct{})
			values := make(map[string]struct{})
			for k, v := range res {
				keys[k] = struct{}{}
				values[v] = struct{}{}
			}
			for k := range tt.assertResult {
				if _, ok := keys[k]; !ok {
					t.Error("Key not found in path map")
				}
			}
		})
	}
}
