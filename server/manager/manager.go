/**
 * @Time: 2022/2/18 14:39
 * @Author: yt.yin
 */

package manager

import "go-socket/server/connect"

// 定义连接管理器的接口

// ManagerI TCP 连接管理，添加、删除、通过以恶搞连接ID获得连接对象，当前连接数、清空全部连接等方法
type ManagerI interface{

	// Add 添加链接
	Add(conn connect.ConnI) error

	// Remove 删除连接
	Remove(conn connect.ConnI) error

	// Conn 根据ConnID获取Tcp连接
	Conn(connID int) (connect.ConnI, error)

	// Size 获取当前连接数
	Size() int

	// Vacuum 清空连接
	Vacuum()
}
