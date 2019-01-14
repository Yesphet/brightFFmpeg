package cmd

import (
	"github.com/spf13/cobra"
	"github.com/CoolMMediaCodec/brightFFmpeg/tools/brish/utils"
	"fmt"
	"gitlab.meitu.com/pl_storage_sdk/qiniu-sdk/api.v7/storage"
	"context"
	"gitlab.meitu.com/pl_storage_sdk/qiniu-sdk/api.v7/auth/qbox"
)

const (
	bucket   = "brightFFmpeg"
	upHost   = "upmt.meitudata.com"
	downHost = "brightFFmpeg.zone1.meitudata.com"
)

func genPushCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "push",
		Short: "Compress the example asserts and push to the cloud",
	}

	var name string

	cmd.Flags().StringVarP(&name, "name", "n", "", "example name")
	cmd.MarkFlagRequired("name")

	cmd.Run = func(cmd *cobra.Command, args []string) {
		path := globalCtx.ExamplePath(name)
		if !utils.IsExist(path) {
			Fatalf("Example %s does not exist")
		}

		zipPath, err := utils.Package(name, path)
		CheckFatalf(err, "Compress example asserts failed")

		_, err = uploadPackage(name, zipPath)
		CheckFatalf(err, "Upload example package failed")

		fmt.Printf("Push example %s success. Download url %s \n", name, downloadUrl(name))

	}

	return cmd
}

func uploadPackage(name, path string) (*storage.PutRet, error) {
	key := name + ".zip"

	storageConf := &storage.Config{
		Zone: &storage.Zone{
			SrcUpHosts: []string{upHost},
		},
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putPolicy := &storage.PutPolicy{}

	mac := qbox.NewMac(globalCtx.AccessKey, globalCtx.SecretKey)
	putPolicy.Scope = fmt.Sprintf("%s:%s", bucket, key)
	upToken := putPolicy.UploadToken(mac)

	resumeUploader := storage.NewResumeUploader(storageConf)
	ret := storage.PutRet{}
	err := resumeUploader.PutFile(context.Background(), &ret, upToken, key, path, nil)
	return &ret, err
}
