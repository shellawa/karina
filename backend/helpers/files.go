package helpers

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type FileService struct {
	ctx context.Context
}

func (s *FileService) Initialize(ctx context.Context) {
	s.ctx = ctx
}

func (s *FileService) SelectFile() string {
	filePath, _ := runtime.OpenFileDialog(s.ctx, runtime.OpenDialogOptions{Title: "Select the submission file"})
	return filePath
}
