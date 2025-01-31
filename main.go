package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/flyqie/n1-vnc/libvnc"
	"github.com/spf13/pflag"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var ip string
var password string
var port int
var debug bool

func init() {
	pflag.IntVar(&port, "port", 5900, "vnc port")
	pflag.StringVar(&ip, "ip", "192.168.1.89", "n1 ip")
	pflag.StringVar(&password, "password", "", "vnc password")
	pflag.BoolVar(&debug, "debug", false, "debug mode")
}

type ChannelWriter struct {
	C chan []byte
}

func (cw *ChannelWriter) Write(p []byte) (n int, err error) {
	cp := make([]byte, len(p))
	copy(cp, p)
	cw.C <- cp
	return len(p), nil
}

func main() {
	log.Printf("[MAIN] Hello~\n")
	pflag.Parse()
	kbCh := make(chan uint32, 255)
	csrCh := make(chan bool)
	cssCh := make(chan struct{})
	w, h := getN1ScreenWH(ip)
	if w <= 0 || h <= 0 {
		log.Fatalf("[MAIN] N1 screen width and height must be positive\n")
	}
	if debug {
		log.Printf("[MAIN] Debug mode enabled\n")
		rfbErrCh := make(chan []byte, 255)
		rfbInfoCh := make(chan []byte, 255)
		rfbErrChWriter := ChannelWriter{C: rfbErrCh}
		rfbInfoChWriter := ChannelWriter{C: rfbInfoCh}
		go func() {
			for {
				select {
				case err := <-rfbErrCh:
					log.Printf("[LIBVNCSERVER_ERR] %s", err)
				case info := <-rfbInfoCh:
					log.Printf("[LIBVNCSERVER_INFO] %s", info)
				}
			}
		}()
		libvnc.RfbInfoLogger = &rfbInfoChWriter
		libvnc.RfbErrLogger = &rfbErrChWriter
	}

	vs := libvnc.NewVNCServer(w, h, 32, port, "N1", password, "n1")
	vs.SetKeyEventCallback(func(down bool, key uint32) {
		if down {
			kbCh <- key
		}
	})
	vs.SetPtrEventCallback(func(buttonMask int, x int, y int) {
		if debug {
			log.Printf("buttonMask %d x %d y %d\n", buttonMask, x, y)
		}
	})
	vs.SetHaveClientStatusCallback(func(status bool) {
		if debug {
			log.Printf("[HAVE_CLIENT_STATUS] %t\n", status)
		}
		csrCh <- status
		<-cssCh
	})

	go n1KeyBoard(kbCh, ip)
	go n1ScreenMgr(vs, ip, w, h, csrCh, cssCh)
	go sigListen(vs)
	log.Printf("[MAIN] forward %s to :%d\n", ip, port)
	vs.Run()
	vs.Clean()
	log.Printf("[MAIN] Bye~\n")
}

func sigListen(vs *libvnc.VNCServer) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals
	log.Printf("[MAIN] Shutting down...\n")
	vs.Stop()
	time.Sleep(3 * time.Second)
}

func getN1ScreenWH(ip string) (int, int) {
	client := &http.Client{}
	response, err := client.Get(fmt.Sprintf("http://%s:8080/v1/screenshot", ip))
	if err != nil {
		log.Printf("[TEST] %s\n", err.Error())
		return 0, 0
	}
	body, err := ioutil.ReadAll(response.Body)
	_ = response.Body.Close()
	if err != nil {
		log.Printf("[TEST] %s\n", err.Error())
		return 0, 0
	}
	img, err := jpeg.Decode(bytes.NewReader(body))
	if err != nil {
		log.Printf("[TEST] %s\n", err.Error())
		return 0, 0
	}
	return img.Bounds().Dx(), img.Bounds().Dy()
}

func compareRGBAImages(img1 *image.RGBA, img2 *image.RGBA) image.Rectangle {
	var rect image.Rectangle
	isFirstDiff := true
	for y := img1.Rect.Min.Y; y < img1.Rect.Max.Y; y++ {
		for x := img1.Rect.Min.X; x < img1.Rect.Max.X; x++ {
			if img1.At(x, y) != img2.At(x, y) {
				if isFirstDiff {
					rect = image.Rect(x, y, x+1, y+1)
					isFirstDiff = false
				} else {
					rect = rect.Union(image.Rect(x, y, x+1, y+1))
				}
			}
		}
	}
	return rect
}

func n1ScreenMgr(vs *libvnc.VNCServer, ip string, w, h int, recvCh chan bool, sendCh chan struct{}) {
	var ctx context.Context
	var cancel context.CancelFunc
	for {
		status := <-recvCh
		if status == false {
			// 没有客户端连接, 关闭自动获取
			if cancel != nil {
				cancel()
			}
		} else {
			ctx, cancel = context.WithCancel(context.Background())
			go n1Screen(vs, ip, w, h, ctx)
			// 保证屏幕线程不会出现用户连接时未获取到的问题
			time.Sleep(1 * time.Second)
		}
		sendCh <- struct{}{}
	}
}

func n1Screen(vs *libvnc.VNCServer, ip string, w, h int, ctx context.Context) {
	client := &http.Client{}
	var oldRGBAImg *image.RGBA
ScreenLoop:
	for {
		select {
		case <-ctx.Done():
			break ScreenLoop
		default:
			response, err := client.Get(fmt.Sprintf("http://%s:8080/v1/screenshot", ip))
			if err != nil {
				log.Printf("[SCREEN] %s\n", err.Error())
				continue
			}
			body, err := ioutil.ReadAll(response.Body)
			_ = response.Body.Close()
			if err != nil {
				log.Printf("[SCREEN] %s\n", err.Error())
				continue
			}
			img, err := jpeg.Decode(bytes.NewReader(body))
			if err != nil {
				log.Printf("[SCREEN] %s\n", err.Error())
				continue
			}
			bounds := img.Bounds()
			rgbaImg := image.NewRGBA(bounds)
			for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					c := img.At(x, y)
					rgbaImg.Set(x, y, c)
				}
			}
			w1 := 0
			y1 := 0
			w2 := w
			y2 := h
			if oldRGBAImg != nil {
				diffRect := compareRGBAImages(oldRGBAImg, rgbaImg)
				w1 = diffRect.Min.X
				y1 = diffRect.Min.Y
				w2 = diffRect.Max.X
				y2 = diffRect.Max.Y
			}
			oldRGBAImg = rgbaImg
			if debug {
				log.Printf("[SCREEN] %d %d %d %d", w1, y1, w2, y2)
			}
			if w1 == 0 && y1 == 0 && w2 == 0 && y2 == 0 {
				continue
			}
			vncBlackBuf, _ := vs.RequestBackBuf()
			copy(vncBlackBuf, rgbaImg.Pix)
			vs.ReleaseBackBuf(w1, y1, w2, y2)
		}
	}
}

func n1KeyBoard(ch chan uint32, ip string) {
	client := &http.Client{}
	for {
		var keyCode int
		key := <-ch
		switch key {
		case libvnc.XK_Left:
			keyCode = 21
		case libvnc.XK_Right:
			keyCode = 22
		case libvnc.XK_Up:
			keyCode = 19
		case libvnc.XK_Down:
			keyCode = 20
		case libvnc.XK_Return:
			// 确认键
			keyCode = 23
		case libvnc.XK_Home:
			// 主页键
			keyCode = 3
		case libvnc.XK_Escape:
			// 返回键
			keyCode = 4
		default:
			if debug {
				log.Printf("[KEYBOARD] Unknown Key 0x%X\n", key)
			}
			continue
		}

		jsonData, err := json.Marshal(map[string]interface{}{
			"keycode":   keyCode,
			"longclick": false,
		})
		if err != nil {
			log.Printf("[KEYBOARD] %s\n", err.Error())
			continue
		}
		req, err := http.NewRequest("POST", fmt.Sprintf("http://%s:8080/v1/keyevent", ip), bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("[KEYBOARD] %s\n", err.Error())
			continue
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("[KEYBOARD] %s\n", err.Error())
			continue
		}
		_, err = ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			log.Printf("[KEYBOARD] %s\n", err.Error())
		}
	}
}
