package commands

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/commands"
	"github.com/goal-web/supports/logs"
	"io/ioutil"
	"os"
)

type runner struct {
	commands.Command
	app contracts.Application
}

func Runner(app contracts.Application) contracts.Command {
	return &runner{
		Command: commands.Base("run", "启动 goal"),
		app:     app,
	}
}

func (this *runner) Handle() interface{} {
	path := this.app.Get("path").(string)

	pidPath := path + "/goal.pid"
	// 写入 pid 文件
	_ = ioutil.WriteFile(pidPath, []byte(fmt.Sprintf("%d", os.Getpid())), os.ModePerm)

	if errors := this.app.Start(); len(errors) > 0 {
		logs.WithField("errors", errors).Fatal("goal 启动异常!")
	} else {
		_ = os.Remove(pidPath)
		logs.Default().Info("goal 已关闭")
	}
	return nil
}
