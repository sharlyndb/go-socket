/**
 * @Time: 2022/2/17 11:33
 * @Author: yt.yin
 */

package I

// ServerI 定义服务端的接口
type ServerI interface {

	// Start 启动服务
	Start()

	// Stop 停止服务
	Stop()

	// Run 启动业务
	Run()

	// ConnManager 获取连接管理
	ConnManager() ConnManagerI

}