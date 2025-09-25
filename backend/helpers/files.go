package helpers

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (s *Service) SelectFile() string {
	filePath, _ := runtime.OpenFileDialog(s.ctx, runtime.OpenDialogOptions{Title: "Select the submission file"})
	return filePath
}
