package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"
)

func ValidationRollbackImage(values []string) {
	for i := 0; i < len(values); i++ {
		err := os.Remove(values[i])
		if err != nil {
			fmt.Println(err)
		}
	}
}

func FileImgValidation(files *multipart.FileHeader, part []string) error {
	if files.Size >= 2003000 {
		ValidationRollbackImage(part)
		return fmt.Errorf("file size cannot exceed 2MB") 
	}

	ekstension1 := strings.HasSuffix(files.Filename, ".jpg")
	ekstension2 := strings.HasSuffix(files.Filename, ".jpeg")
	ekstension3 := strings.HasSuffix(files.Filename, ".png")

	if ekstension1 != true && ekstension2 != true && ekstension3 != true {
		ValidationRollbackImage(part)
		return fmt.Errorf("Unrecognized file format")
	} 
	return nil
}