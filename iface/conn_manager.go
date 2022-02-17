/**
 * @Time: 2022/2/17 14:15
 * @Author: yt.yin
 */

package iface

// ConnManagerI TCP 连接管理，添加、删除、通过以恶搞连接ID获得连接对象，当前连接数、清空全部连接等方法
type ConnManagerI interface{

	// Add 添加链接
	Add(conn ConnI)

	// Remove 删除连接
	Remove(conn ConnI)

	// Conn 根据ConnID获取Tcp连接
	Conn(connID uint32) (ConnI, error)

	// Size 获取当前连接数
	Size() int

	// Vacuum 清空连接
	Vacuum()
}