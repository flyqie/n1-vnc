#ifndef N1_VNC_C_H
#define N1_VNC_C_H
#include <stdlib.h>
#include <stdint.h>
#include <pthread.h>
#include <rfb/rfb.h>

typedef struct {
    rfbScreenInfoPtr server;
    char *frontBuffer;
    char *backBuffer;
    int bufferSize;
    int stop_vnc_server_flag;
    int stopped_vnc_server_flag;
    int have_client_flag;
    pthread_rwlock_t frontBufferLock;
    pthread_mutex_t backBufferLock;
    pthread_mutex_t backBufferFuncLock;
} BufferManager;

typedef struct {
    char serverID[256];
} ServerCustomData;

/**
 * 申请修改双缓冲区
 * 注意: 该函数会锁定双缓冲区, 所以请务必调用release_vnc_buf进行切换
 *
 * @param manager
 * @return 双缓冲区内存地址
 */
char *request_back_vnc_buf(BufferManager *manager);

/**
 * 释放双缓冲区
 * 注意: 该函数会解锁双缓冲区, 所以请务必先调用request_back_vnc_buf来获取双缓冲区
 *
 * @param manager
 * @param w1 需要发送到客户端修改的起始宽度
 * @param y1 需要发送到客户端修改的起始高度
 * @param w2 需要发送到客户端修改的终止宽度
 * @param y2 需要发送到客户端修改的终止宽度
 * @return
 */
int release_vnc_buf(BufferManager *manager, int w1, int y1, int w2, int y2);

/**
 * 初始化vnc服务器, 该函数会同步创建双缓冲区
 * 注意: 请务必在不需要时调用cleanup_vnc_server释放内存, 否则会造成内存泄露!
 *
 * @param width 屏幕宽度
 * @param height 屏幕高度
 * @param bits_per_pixel BPP
 * @param port vnc端口
 * @param desktopName 桌面名称
 * @param password vnc密码, 为空为无鉴权
 * @param serverID vnc服务器实例唯一标识符, notifyServerKeyEvent等callback func将会传递该内容
 * @return
 */
BufferManager *
init_vnc_server(const int width, const int height, const int bits_per_pixel, const int port, const char *desktopName,
                const char *password, const char *serverID);

/**
 * 停止vnc服务器
 * 建议停止后等待几秒再进行清理
 *
 * @param manager
 * @return
 */
int stop_vnc_server(BufferManager *manager);

/**
 * 运行vnc服务器
 * 注意: 该函数为阻塞函数
 *
 * @param manager
 * @return
 */
int run_vnc_server(BufferManager *manager);

/**
 * 清理vnc服务器
 * 注意: 请务必在vnc服务器停止后调用
 *
 * @param manager
 * @return
 */
int cleanup_vnc_server(BufferManager *manager);

void setServerRfbLog();

#endif //N1_VNC_C_H