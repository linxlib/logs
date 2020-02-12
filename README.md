# logs

package logs is modified from logrus
修改自 [logrus](https://github.com/sirupsen/logrus)

## 特征
1. 直接用 `logs.Info()`, 不用写一堆初始化自定义之类, 每写一个程序都要拷贝一份也够烦的了
2. 写文件日志只需`logs.AddFileHook("appname")`, 内置了一个自动滚动的`RotateFileLoggerHook`, 最大保留7天, 达到5M后改名备份
3. 内置了对windows命令行的颜色输出支持
4. 没有其它依赖, 该依赖的都内置了
5. 没有json输出, 平时写写小工具根本用不上这个东西
6. 日志内容也会有颜色, 不像原版的内容没颜色

## 注意
如果你不是自己平时玩玩的项目, 最好是clone一份过去进行一顿魔改
or
如果直接用满足不了你的要求, 最好是clone一份过去进行一顿魔改