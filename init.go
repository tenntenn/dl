package dl

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var _ Cmd = (*Init)(nil)

// Init structs for $ dl init.
type Init struct{}

// NewInit for running $ dl init.
func NewInit() *Init {
	return &Init{}
}

// Run prepares an environment for dl commands.
func (i *Init) Run(ctx context.Context, baseDir string) error {
	// check if hooks directory exists or not
	if _, err := os.Stat(filepath.Join(baseDir, ".git", "hooks")); os.IsNotExist(err) {
		return err
	}

	// Initialization
	if err := i.addGitPreHookScript(ctx, baseDir); err != nil {
		return err
	}
	if err := i.addGitPostHookScript(ctx, baseDir); err != nil {
		return err
	}
	if err := i.createDlDirIfNotExist(ctx, baseDir); err != nil {
		return err
	}
	if err := i.addDlIntoGitIgnore(ctx, baseDir); err != nil {
		return err
	}

	return nil
}

func (i *Init) addGitPreHookScript(ctx context.Context, baseDir string) error {
	return i.insertCodesIfNotExist(ctx, filepath.Join(baseDir, ".git", "hooks", "pre-commit"), cleanCmd, preCommitScript)
}

func (i *Init) addGitPostHookScript(ctx context.Context, baseDir string) error {
	return i.insertCodesIfNotExist(ctx, filepath.Join(baseDir, ".git", "hooks", "post-commit"), restoreCmd, postCommitScript)
}

func (i *Init) addDlIntoGitIgnore(ctx context.Context, baseDir string) error {
	return i.insertCodesIfNotExist(ctx, filepath.Join(baseDir, ".gitignore"), dlDir, fmt.Sprintf("\n%s\n", dlDir))
}

func (*Init) insertCodesIfNotExist(ctx context.Context, targetFilePath string, checkedCodesIfExists string, addedCodes string) error {
	f, err := os.OpenFile(targetFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	// It checks if `checkedCodesIfExists` has been installed or not.
	// If so, not inserting codes.
	buf, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	if bytes.Contains(buf, []byte(checkedCodesIfExists)) {
		return nil
	}

	if _, err := fmt.Fprint(f, addedCodes); err != nil {
		return err
	}

	return os.Chmod(targetFilePath, 0755)
}

func (*Init) createDlDirIfNotExist(ctx context.Context, baseDir string) error {
	path := filepath.Join(baseDir, dlDir)
	if stat, err := os.Stat(path); err == nil {
		if stat.IsDir() {
			return nil
		}
		return fmt.Errorf("%s has been already existed as file. Please rename or delete it.", path)
	}

	return os.MkdirAll(path, 0755)
}
