package cmd

import (
	"github.com/spf13/cobra"
	"os/user"
	"gopkg.in/yaml.v2"
	"github.com/CoolMMediaCodec/brightFFmpeg/tools/brish/utils"
	"github.com/mitchellh/go-homedir"
	"os"
	"fmt"
	"path/filepath"
	"bytes"
)

var globalCtx = &Config{}

type Config struct {
	AccessKey        string
	SecretKey        string
	Author           string
	BrightFFmpegPath string
}

func readConfig() {
	path := configPath()
	if !utils.IsExist(path) {
		Fatalf("Please use 'brish config' to set configs first.")
	}

	raw, err := utils.Read(path)
	CheckFatalf(err, "Read config failed")

	err = yaml.Unmarshal(raw, globalCtx)
	CheckFatalf(err, "Read config failed")
}

func configPath() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfgPath := filepath.Join(home, ".brish")
	return cfgPath
}

func genConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "init brish's config",
	}

	cfg := &Config{}

	cmd.Flags().StringVar(&cfg.AccessKey, "ak", "", "The AccessKey")
	cmd.Flags().StringVar(&cfg.SecretKey, "sk", "", "The SecretKey")
	cmd.Flags().StringVarP(&cfg.Author, "author", "a", "", "The Author")
	cmd.Flags().StringVarP(&cfg.BrightFFmpegPath, "path", "p", "", "The local path of brightFFmpeg project")
	cmd.MarkFlagRequired("accessKey")
	cmd.MarkFlagRequired("secretKey")
	cmd.MarkFlagRequired("path")

	cmd.Run = func(cmd *cobra.Command, args []string) {
		if cfg.Author == "" {
			u, err := user.Current()
			if err != nil {
				cfg.Author = u.Username
			}
		}

		var err error
		cfg.BrightFFmpegPath, err = filepath.Abs(cfg.BrightFFmpegPath)
		CheckFatalf(err, "Invalid brightFFmpeg project path")

		cfgString, _ := yaml.Marshal(cfg)
		cfgPath := configPath()

		err = utils.WriteToFile(cfgPath, bytes.NewReader(cfgString))
		CheckFatalf(err, "Write configs to file failed, path %s", cfgPath)

		fmt.Println("Config Updated success!")
	}

	return cmd
}
