package pkg

import (
	"GoPool/util"
	"fmt"
	"testing"
)

var domains = []string{
	"www.google.com",
	"www.4399.com",
	"www.7k7k.com",
	"www.qq.com",
	"www.youku.com",
	"www.tudou.com",
	"www.youdao.com",
	"www.weixin.com",
	"www.csdn.com",
	"www.bilibili.com",
}

type TestTask struct {
	CMD string
}

func (task *TestTask) Run() {
	util.Exec(task.CMD)
}

func TestNewExecutor(t *testing.T) {
	t.Log(t.Name())
	ex := NewExecutor(4)

	for _, domain := range domains {
		ex.Execute(&TestTask{
			fmt.Sprintf("ping %s -c 10", domain),
		})
	}
	ex.Shutdown()
}
