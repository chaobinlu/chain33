/*
扫描chain33项目下plugin中所有的插件，根据扫描到的结果重新更新共识、执行器和数据操作的初始化文件 init.go
*/
package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.33.cn/chain33/chain33/cmd/tools/strategy"
	"gitlab.33.cn/chain33/chain33/cmd/tools/types"
)

func UpdateInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "updateinit",
		Short: "Update chain33 plugin consensus、dapp、store init.go file",
		Run:   updateInit,
	}
	cmd.Flags().StringP("path", "p", "plugin", "path of plugin")
	cmd.Flags().StringP("packname", "", "", "project package name")
	return cmd
}

func updateInit(cmd *cobra.Command, args []string) {
	path, _ := cmd.Flags().GetString("path")
	packname, _ := cmd.Flags().GetString("packname")
	s := strategy.New(types.KeyUpdateInit)
	if s == nil {
		fmt.Println(types.KeyUpdateInit, "Not support")
		return
	}
	s.SetParam("path", path)
	s.SetParam("packname", packname)
	s.Run()
}