# APUE 3

## 文件 `I/O`

`open`, `read`, `write`, `lseek` & `close`

`dup`, `fcntl`, `sync`, `fsync` & `ioctl`

不带缓冲的 `I/O` 不 buffered `I/O` ：每个 `read` 和 `write` 都调用内核中的一个系统调用。不是 `ISO C` 的组成部分。是 `POSIX.1` 和 `Single UNIX Specification` 的组成部分。

### 文件描述符

对于内核而言，所有打开的文件都通过文件描述符引用。
文件描述符是一个非负整数。

`UNIX` 系统 `shell` 将文件描述符 0 与进程的标准输入相关，1 与标准输出关联，2 与标准错误关联。

常量：`STDIN_FILENO`, `STDOUT_FILENO` & `STDERR_FILENO，` `<unistd.h>` 中。

文件描述符变化范围：0 ~ `OPEN_MAX` - 1

### `open` & `openat`

```C
#include <fcntl.h>

int open(const char *path, int oflag, ... /* mode_t mode */);

int openat(int fd, const char *path, int oflag, ... /* mode_t mode */);
```

最后一个参数 `...` ，`ISO C` 用这种方法表明余下的参数数量及类型是可变的。

`open` 函数，仅当创建新文件时才使用最后的参数。

`path` 参数：要打开或创建的文件名。

`oflag` 参数：用于说明此函数的多个选项。用下列一个或多个常量 或 运算构成 `oflag`，定义于 `<fcntl.h>` 头文件中。

这 5 个常量必须且只能指定一个：

`O_RDONLY` 只读打开

`O_WRONLY` 只写打开

`O_RDWR` 读、写打开

`O_EXEC` 只执行打开

`O_SEARCH` 只搜索打开（用于目录中）

下列可选：

`O_APPEND`, `O_CLOSEXEC`, `O_CREAT`, `O_DIRECTORY`, `O_EXCL`, `O_NOCTTY`, `O_NOFOLLOW`, `O_NONBLOCK`, `O_SYNC`, `O_TRUNC`, `O_TTY_INIT`

下面两个标志也是可选的：

`O_DSYNC`, `O_RSYNC`

由 `open` 和 `openat` 返回的文件描述符一定是最小的未用描述符数值。

`fd` 参数区分 `open` 和 `openat`，3 种可能性：

1. `path` 参数指定的是绝对路径名， `fd` 被忽略

2. `path` 参数指定的是相对路径名， `fd` 指出相对路径名在文件系统中的开始地址， `fd` 参数通过打开相对路径名所在的目录获得。

3. `path` 参数指定的是相对路径名， `fd` 具有特殊值 `AT_FDCWD` 。路径名在当前工作目录获得。

`openat` 是 `POSIX.1` 新增的。让线程可以通过相对路径名打开目录中的文件。避免 `time-of-check-to-time-of-use` `TOCTTOU` 错误。
