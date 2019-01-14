package cmd

import (
	"github.com/spf13/cobra"
	"github.com/CoolMMediaCodec/brightFFmpeg/tools/brish/utils"
	"net/http"
	"fmt"
	"path/filepath"
)

func genPullCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Pull the example package from the cloud and extract",
	}

	var name string

	cmd.Flags().StringVarP(&name, "name", "n", "", "example name")

	cmd.Run = func(cmd *cobra.Command, args []string) {
		dir := globalCtx.ExamplePath(name)
		if !utils.IsExist(dir) {
			Fatalf("Example %s does not exist")
		}

		f := filepath.Join(dir, utils.ZipName(name))

		err := downloadPackage(f, name)
		defer utils.DeleteFile(f)
		CheckFatalf(err, "pull package failed")

		err = utils.Unpack(f, dir)
		CheckFatalf(err, "unpack failed")
	}

	return cmd
}

func downloadUrl(name string) string {
	return "http://" + downHost + "/" + name + ".zip"
}

func downloadPackage(path, name string) error {
	u := downloadUrl(name)
	r, err := http.Get(u)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpect download status %s", r.StatusCode)
	}

	if err := utils.WriteToFile(path, r.Body); err != nil {
		return err
	}
	return nil
}
