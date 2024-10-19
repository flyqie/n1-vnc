package libvnc

/*
#cgo pkg-config: libvncserver
#include "vnc.h"
*/
import "C"

import (
	"io"
	"unsafe"
)

var RfbInfoLogger io.Writer
var RfbErrLogger io.Writer
var VncServers map[string]*VNCServer

type KeyEventCallback func(down bool, key uint32)
type PtrEventCallback func(buttonMask int, x int, y int)
type HaveClientStatusCallback func(have bool)

type VNCServer struct {
	manager                  *C.BufferManager
	serverID                 string
	keyEventCallback         KeyEventCallback
	ptrEventCallback         PtrEventCallback
	haveClientStatusCallback HaveClientStatusCallback
}

func init() {
	VncServers = make(map[string]*VNCServer)
	C.setServerRfbLog()
}

// NewVNCServer 创建vnc server
func NewVNCServer(width, height, bitsPerPixel, port int, desktopName, password, serverID string) *VNCServer {
	if _, ok := VncServers[serverID]; ok {
		return nil
	}

	cDesktopName := C.CString(desktopName)
	cPassword := C.CString(password)
	cServerID := C.CString(serverID)
	defer C.free(unsafe.Pointer(cDesktopName))
	defer C.free(unsafe.Pointer(cPassword))
	defer C.free(unsafe.Pointer(cServerID))

	manager := C.init_vnc_server(C.int(width), C.int(height), C.int(bitsPerPixel), C.int(port), cDesktopName, cPassword, cServerID)
	vncServer := &VNCServer{
		manager:          manager,
		serverID:         serverID,
		keyEventCallback: nil,
	}
	VncServers[serverID] = vncServer
	return vncServer
}

// Run 阻塞形式运行vnc server
func (v *VNCServer) Run() int {
	return int(C.run_vnc_server(v.manager))
}

// Stop 停止vnc server, 需要几秒
func (v *VNCServer) Stop() int {
	return int(C.stop_vnc_server(v.manager))
}

// Clean 清理vnc server
func (v *VNCServer) Clean() int {
	code := int(C.cleanup_vnc_server(v.manager))
	if code == 0 {
		delete(VncServers, v.serverID)
	}
	return code
}

// SetKeyEventCallback 设置vnc server的key event callback
func (v *VNCServer) SetKeyEventCallback(cb KeyEventCallback) {
	v.keyEventCallback = cb
}

// SetPtrEventCallback 设置vnc server的ptr event callback
func (v *VNCServer) SetPtrEventCallback(cb PtrEventCallback) {
	v.ptrEventCallback = cb
}

// SetHaveClientStatusCallback 设置vnc server的have client status callback
func (v *VNCServer) SetHaveClientStatusCallback(cb HaveClientStatusCallback) {
	v.haveClientStatusCallback = cb
}

// RequestBackBuf 申请写入双缓冲区, 操作后请务必使用ReleaseBackBuf来释放锁并切换!
func (v *VNCServer) RequestBackBuf() ([]byte, int) {
	backBuffer := C.request_back_vnc_buf(v.manager)
	size := int(v.manager.bufferSize)
	// 使用 Go 切片来封装缓冲区
	return (*[1 << 30]byte)(unsafe.Pointer(backBuffer))[:size:size], size
}

// ReleaseBackBuf 释放锁并切换双缓冲区, RequestBackBuf操作后务必调用此方法
func (v *VNCServer) ReleaseBackBuf(w1, y1, w2, y2 int) int {
	return int(C.release_vnc_buf(v.manager, C.int(w1), C.int(y1), C.int(w2), C.int(y2)))
}

//export notifyServerLogInfo
func notifyServerLogInfo(str *C.char, n C.int) {
	if RfbInfoLogger != nil {
		goStr := C.GoStringN(str, n)
		_, _ = RfbInfoLogger.Write([]byte(goStr))
	}
}

//export notifyServerLogErr
func notifyServerLogErr(str *C.char, n C.int) {
	if RfbErrLogger != nil {
		goStr := C.GoStringN(str, n)
		_, _ = RfbErrLogger.Write([]byte(goStr))
	}
}

//export notifyServerKeyEvent
func notifyServerKeyEvent(down C.int8_t, key C.uint32_t, serverID *C.char) {
	goServerID := C.GoString(serverID)
	goDown := down != 0

	if vncServer, ok := VncServers[goServerID]; ok {
		if vncServer.keyEventCallback != nil {
			vncServer.keyEventCallback(goDown, uint32(key))
		}
	}
}

//export notifyServerPtrEvent
func notifyServerPtrEvent(buttonMask C.int, x C.int, y C.int, serverID *C.char) {
	goServerID := C.GoString(serverID)

	if vncServer, ok := VncServers[goServerID]; ok {
		if vncServer.ptrEventCallback != nil {
			vncServer.ptrEventCallback(int(buttonMask), int(x), int(y))
		}
	}
}

//export notifyServerHaveClientStatus
func notifyServerHaveClientStatus(flag C.int, serverID *C.char) {
	goServerID := C.GoString(serverID)
	goFlag := flag != 0

	if vncServer, ok := VncServers[goServerID]; ok {
		if vncServer.haveClientStatusCallback != nil {
			vncServer.haveClientStatusCallback(goFlag)
		}
	}
}
