# README

## 简介

该软件用于通过具有风险的公共信道（如各类聊天软件、论坛），软件使用AES-256算法，端到端地安全传输文字信息。加密时，软件将文字信息加密并转换为二维码图片。解密时，软件自动读取二维码图片，并解密为文字。

功能说明：

(0) 设置密钥(KEY)

打开设置菜单，设置一个KEY，软件之后将使用该KEY进行AES加解密。您可以把您的KEY以及该软件分享给您的朋友。

(1) 加密

点击加密按钮，输入的文本将会按照您设置的KEY加密并转换为二维码，并自动复制到剪贴板。（请勿超过200字，二维码太复杂会难以识别）

(2) 解密

实时识别屏幕上显示的二维码，使用您设置的KEY解码，并转换为消息。（请让二维码完整地显示在屏幕上）

## 构建说明

该项目的图形界面使用 Wails Vue 构建。
项目依赖于gozxing(https://github.com/makiuchi-d/gozxing )，编译前需将其源代码放置于3rd/gozxing-master中。

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

### Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

### Building

To build a redistributable, production mode package, use `wails build`.
