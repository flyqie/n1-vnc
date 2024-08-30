package main

import (
	"bytes"
	"encoding/json"
	"github.com/flyqie/n1-vnc/libvnc"
	"github.com/spf13/pflag"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var url string
var password string
var port int
var debug bool

func init() {
	pflag.IntVar(&port, "port", 5900, "vnc port")
	pflag.StringVar(&url, "url", "http://192.168.1.82:8080", "n1 url")
	pflag.StringVar(&password, "password", "", "vnc password")
	pflag.BoolVar(&debug, "debug", false, "debug mode")
}

func main() {
	log.Printf("[MAIN] Hello~\n")
	pflag.Parse()
	kbCH := make(chan uint32, 255)
	if strings.HasSuffix(url, "/") {
		url = url[:len(url)-1]
	}
	w, h := getN1ScreenWH(url)
	if w <= 0 || h <= 0 {
		log.Fatalf("[MAIN] N1 screen width and height must be positive\n")
	}

	vs := libvnc.NewVNCServer(w, h, 32, port, "N1", password, "n1")
	vs.SetKeyEventCallback(func(down bool, key uint32) {
		if down {
			kbCH <- key
		}
	})
	vs.SetHaveClientStatusCallback(func(status bool) {
		// TODO: 在无连接时不获取图像
		if debug {
			log.Printf("[HAVE_CLIENT_STATUS] %t\n", status)
		}
	})

	go n1KeyBoard(kbCH, url)
	go n1Screen(vs, url, w, h)
	go sigListen(vs)
	log.Printf("[MAIN] forward %s to :%d\n", url, port)
	vs.Run()
	log.Printf("[MAIN] Bye~\n")
}

func sigListen(vs *libvnc.VNCServer) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals
	log.Printf("[MAIN] Shutting down...\n")
	vs.Stop()
	time.Sleep(3 * time.Second)
	vs.Clean()
}

func getN1ScreenWH(url string) (int, int) {
	client := &http.Client{}
	response, err := client.Get(url + "/v1/screenshot")
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

func n1Screen(vs *libvnc.VNCServer, url string, w, h int) {
	client := &http.Client{}
	var oldRGBAImg *image.RGBA
	for {
		response, err := client.Get(url + "/v1/screenshot")
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
		vncBlackBuf, _ := vs.RequestBackBuf()
		copy(vncBlackBuf, rgbaImg.Pix)
		vs.ReleaseBackBuf(w1, y1, w2, y2)
	}
}

func n1KeyBoard(ch chan uint32, url string) {
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
		req, err := http.NewRequest("POST", url+"/v1/keyevent", bytes.NewBuffer(jsonData))
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
