/**
 * @Time: 2022/2/18 14:44
 * @Author: yt.yin
 */

package server

import "github.com/goworkeryyt/go-socket/server/manager"

// ServerI 定义服务端的接口
type ServerI interface {

	// Start 启动服务
	Start()

	// Stop 停止服务
	Stop()

	// Run 启动业务
	Run()

	// ConnManager 获取连接管理
	ConnManager() manager.ManagerI

}
