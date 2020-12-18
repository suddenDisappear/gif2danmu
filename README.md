# gif转换高级弹幕文本
## 环境要求
+ golang >= 1.13
## 安装
```
git clone https://github.com/suddenDisappear/gif2danmu.git
cd gif2danmu
go env GO111MODULE=on
go env GOPROXY=https://goproxy.cn,direct
go mod download
go build
```
## 导出文件描述
导出文件存放在导出目录/文件名文件夹下：
+ delay.txt:gif帧延时信息，每行分别对应每帧图像在屏幕上停留时间，单位:10ms
+ *.txt:每帧图片分解成像素标识情况，文件内按照rgba颜色分组
+ *.png:每帧图片还原成像素点情况，仅作为还原效果供参考，实际效果不同
## 建议
+ gif图像使用透明背景
