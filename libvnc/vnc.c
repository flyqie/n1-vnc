#include "vnc.h"
#include "_cgo_export.h"

/**
 * 申请修改双缓冲区
 * 注意: 该函数会锁定双缓冲区, 所以请务必调用release_vnc_buf进行切换
 *
 * @param manager
 * @return 双缓冲区内存地址
 */
char *request_back_vnc_buf(BufferManager *manager) {
    pthread_mutex_lock(&manager->backBufferLock);
    pthread_mutex_lock(&manager->backBufferFuncLock);
    char *buffer = manager->backBuffer;
    pthread_mutex_unlock(&manager->backBufferFuncLock);
    return buffer;
}

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
int release_vnc_buf(BufferManager *manager, int w1, int y1, int w2, int y2) {
    pthread_rwlock_wrlock(&manager->frontBufferLock);
    pthread_mutex_lock(&manager->backBufferFuncLock);
    char *temp = manager->frontBuffer;
    manager->frontBuffer = manager->backBuffer;
    manager->backBuffer = temp;
    manager->server->frameBuffer = manager->frontBuffer;
    pthread_mutex_unlock(&manager->backBufferLock);
    pthread_rwlock_unlock(&manager->frontBufferLock);
    rfbMarkRectAsModified(manager->server, w1, y1, w2, y2);
    pthread_mutex_unlock(&manager->backBufferFuncLock);
    return 0;
}

void key_event(rfbBool down, rfbKeySym key, rfbClientPtr cl) {
    ServerCustomData *server_custom_data = (ServerCustomData *) cl->screen->screenData;
    notifyServerKeyEvent(down, key, server_custom_data->serverID);
}

void ptr_event(int buttonMask, int x, int y, rfbClientPtr cl) {
    ServerCustomData *server_custom_data = (ServerCustomData *) cl->screen->screenData;
    notifyServerPtrEvent(buttonMask, x, y, server_custom_data->serverID);
}

void have_client_status(BufferManager *manager) {
    ServerCustomData *server_custom_data = (ServerCustomData *) manager->server->screenData;
    notifyServerHaveClientStatus(manager->have_client_flag, server_custom_data->serverID);
}

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
                const char *password, const char *serverID) {
    BufferManager *manager = (BufferManager *) calloc(1, sizeof(BufferManager));
    manager->bufferSize = width * height * (bits_per_pixel / 8);
    manager->frontBuffer = (char *) calloc(1, manager->bufferSize);
    manager->backBuffer = (char *) calloc(1, manager->bufferSize);
    manager->server = rfbGetScreen(NULL, NULL, width, height, 8, 4, (bits_per_pixel / 8));
    manager->server->frameBuffer = manager->frontBuffer;
    manager->server->desktopName = strdup(desktopName);
    manager->server->alwaysShared = TRUE;
    manager->server->httpDir = NULL;
    manager->server->port = port;
    manager->server->kbdAddEvent = key_event;
    manager->server->ptrAddEvent = ptr_event;
    // 用来识别多个屏幕
    ServerCustomData *customData = (ServerCustomData *) malloc(sizeof(ServerCustomData));
    strcpy(customData->serverID, serverID);
    manager->server->screenData = (void *) customData;
    // 密码设置
    if (*password != '\0') {
        char **passwordList = malloc(sizeof(char **) * 2);
        passwordList[0] = strdup(password);
        passwordList[1] = NULL;
        manager->server->authPasswdData = (void *) passwordList;
        manager->server->passwordCheck = rfbCheckPasswordByList;
    }

    rfbInitServer(manager->server);
    /* Mark as dirty since we haven't sent any updates at all yet. */
    rfbMarkRectAsModified(manager->server, 0, 0, width, height);
    pthread_rwlock_init(&manager->frontBufferLock, NULL);
    pthread_mutex_init(&manager->backBufferLock, NULL);
    pthread_mutex_init(&manager->backBufferFuncLock, NULL);

    return manager;
}

/**
 * 停止vnc服务器
 * 建议停止后等待几秒再进行清理
 *
 * @param manager
 * @return
 */
int stop_vnc_server(BufferManager *manager) {
    manager->stop_vnc_server_flag = 1;
    return 0;
}

/**
 * 运行vnc服务器
 * 注意: 该函数为阻塞函数
 *
 * @param manager
 * @return
 */
int run_vnc_server(BufferManager *manager) {
    int hasRLock;
    manager->stop_vnc_server_flag = 0;
    manager->stopped_vnc_server_flag = 0;
    while (!manager->stop_vnc_server_flag) {
        if (manager->server->clientHead != NULL) {
            if(manager->have_client_flag != 1) {
                manager->have_client_flag = 1;
                have_client_status(manager);
            }
            hasRLock = 1;
        } else {
            if(manager->have_client_flag != 0) {
                manager->have_client_flag = 0;
                have_client_status(manager);
            }
            hasRLock = 0;
        }
        if (hasRLock) {
            pthread_rwlock_rdlock(&manager->frontBufferLock);
        }
        rfbProcessEvents(manager->server, -1);
        if (hasRLock) {
            pthread_rwlock_unlock(&manager->frontBufferLock);
        }
    }
    rfbScreenCleanup(manager->server);
    manager->stopped_vnc_server_flag = 1;
    return 0;
}

/**
 * 清理vnc服务器
 * 注意: 请务必在vnc服务器停止后调用
 *
 * @param manager
 * @return
 */
int cleanup_vnc_server(BufferManager *manager) {
    if (manager->stop_vnc_server_flag != 1 || manager->stopped_vnc_server_flag != 1) {
        return -1;
    }
    pthread_rwlock_destroy(&manager->frontBufferLock);
    pthread_mutex_destroy(&manager->backBufferLock);
    pthread_mutex_destroy(&manager->backBufferFuncLock);
    free(manager->frontBuffer);
    free(manager->backBuffer);
    free(manager->server->screenData);
    free(manager);
    return 0;
}

void rfbServerLogInfoToString(const char *format, ...) {
    int bufferSize = 4096;
    char buffer[bufferSize];
    va_list argptr;
    va_start(argptr, format);
    int n = vsnprintf(buffer, bufferSize, format, argptr);
    va_end(argptr);
    notifyServerLogInfo(buffer, n);
}

void rfbServerLogErrToString(const char *format, ...) {
    int bufferSize = 4096;
    char buffer[bufferSize];
    va_list argptr;
    va_start(argptr, format);
    int n = vsnprintf(buffer, bufferSize, format, argptr);
    va_end(argptr);
    notifyServerLogErr(buffer, n);
}

void setServerRfbLog() {
    rfbLog = rfbServerLogInfoToString;
    rfbErr = rfbServerLogErrToString;
}
