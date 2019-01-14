package cmd

import (
	"github.com/spf13/cobra"
	"github.com/CoolMMediaCodec/brightFFmpeg/tools/brish/utils"
	"path/filepath"
	"strings"
	"time"
	"fmt"
)

type cmdContextNew struct {
	name string
}

func genNewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new example",
	}
	ctx := &cmdContextNew{}

	cmd.Flags().StringVarP(&ctx.name, "name", "n", "", "example name")
	cmd.MarkFlagRequired("name")

	cmd.Run = func(cmd *cobra.Command, args []string) {
		egPath := globalCtx.ExamplePath(ctx.name)
		if utils.IsExist(egPath) {
			Fatalf("Example %s already exist", ctx.name)
		}

		err := utils.Makedir(egPath)
		CheckFatalf(err, "Create example's directory failed")

		err = createTemplateFile(egPath, ctx)
		CheckFatalf(err, "Create example's templates failed")

		fmt.Printf("Create new examle %s success, path: %s.\n", ctx.name, egPath)
	}

	return cmd
}

const (
	templateReadme = `---
Name: {name}
Author: {author}
CreateTime: {date}
Tags: 
    - {name}
---

# {name}

## Introduction

Introduce briefly

## Description

Describe in detail


`

	templateCmd = `#!/usr/bin/env bash

cd $(dirname "$0")

`
)

func createTemplateFile(path string, ctx *cmdContextNew) error {

	readme := generateReadmeMd(ctx.name)
	readmePath := filepath.Join(path, "README.md")
	if err := utils.WriteToFile(readmePath, strings.NewReader(readme)); err != nil {
		return err
	}

	cmdsh := generateCmdSh()
	cmdshPath := filepath.Join(path, "cmd.sh")
	if err := utils.WriteToFile(cmdshPath, strings.NewReader(cmdsh)); err != nil {
		return err
	}

	return nil
}

func generateReadmeMd(name string) string {
	ret := strings.Replace(templateReadme, "{name}", name, -1)
	ret = strings.Replace(ret, "{author}", globalCtx.Author, -1)
	ret = strings.Replace(ret, "{date}", time.Now().Format("Mon, 2006-01-02"), -1)
	return ret
}

func generateCmdSh() string {
	return templateCmd
}
