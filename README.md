# param-validator

go-zero httpx的validator插件

param-validator 集成了 github.com/go-playground/validator/v10，可以使用其所有的校验方法

同时 param-validator 内置了一部分自定义校验方法

## 注意事项

httpx会优先go-zero自带的校验。请注意，举个简单的例子

```api
type (
	demo {
        Name string `form:"name,optional" validate:"xPhone"` // 姓名
    }
)
```

这会先通过 httpx自带校验 optional
再通过param-validator的xPhone校验，看起来是没问题的

```api
type (
	demo {
        Name string `form:"name" validate:"omitempty,xPhone"` // 姓名
    }
)
```

这种写法代表在 go-zero自带校验中，是不允许为空的，但是在param-validator中是允许为空，值不为空的时候，进行xPhone校验，
会有冲突。

建议 go-zero 中都写上 optional


## 内置自定义校验

### xPhone

校验手机号

规则位1开头，11位纯数字

```api
type (
	demo {
        Name string `form:"name,optional" validate:"xPhone"` // 姓名
    }
)
```

返回错误为

```text
{0}必须为手机号，1开头，长度为11位
```
{0}为参数名称

### xPassword
校验密码

长度自定义，使用方法为 validate:"xPassword=8-15" 代表 字符串长度为 8到15位。左右都是闭区间

为密码校验规则, 校验其长度位，需由字母（同时要大小和写）、数字、特殊字符串三种组成，不能使用空格、中文

```api
type (
	demo {
        Name string `form:"name,optional" validate:"xPassword=8-15"` // 姓名
    }
)
```

返回错误为

```text
{0}长度{1}，需由字母（区分大小写）、数字、特殊字符串三种组成，不能使用空格、中文
```
{0}为参数名称

{1}为param即8-15

### xStr
首尾不能有空格的字符串校验

长度自定义，使用方法为 validate:"xStr=1-300" 代表字符串长度为 1-300位，左右都为闭区间

首尾不能有空格，中间可以有空格

允许中文，特殊字符串，英文，数字

```api
type (
	demo {
        Name string `form:"name,optional" validate:"xStr=8-15"` // 姓名
    }
)
```

返回错误为

```text
{0}长度{1}，首尾不能有空格
```
{0}为参数名称

{1}为param即8-15

### xStrWithoutZh
首尾不能有空格，不能包含中文的字符串校验

为字符串校验方法，使用方法为 validate:"xStrWithoutZh=1-300" 代表字符串长度为 1-300位，左右都为闭区间

长度自定义

首尾不能有空格，中间可以有空格

不能包含中文

允许特殊字符串，英文，数字
```api
type (
	demo {
        Name string `form:"name,optional" validate:"xStrWithoutZh=8-15"` // 姓名
    }
)
```

返回错误为

```text
{0}长度{1}，首尾不能有空格，不能包含中文
```
{0}为参数名称

{1}为param即8-15

### xStrWithoutZhAndSpec
首尾不能有空格，不能包含中文，不能包含特殊字符串的字符串校验

使用方法为 validate:"xStrWithoutZhAndSpec=1-300" 代表字符串长度为 1-300位，左右都为闭区间

长度自定义

首尾不能有空格，中间可以有空格

不能包含中文

不能包含特殊字符 !@#$%^&*()_+\-=[]{};':"\\|,.<>/?等

允许字母和数字

```api
type (
	demo {
        Name string `form:"name,optional" validate:"xStrWithoutZhAndSpec=8-15"` // 姓名
    }
)
```

返回错误为

```text
{0}长度{1}，首尾不能有空格，不能包含中文和特殊字符串
```
{0}为参数名称

{1}为param即8-15

### xStrWithoutSpec
首尾不能有空格，不能包含特殊字符串的字符串校验

使用方法为 validate:"xStrWithoutSpec=1-300" 代表字符串长度为 1-300位，左右都为闭区间

长度自定义

首尾不能有空格，中间可以有空格

不能包含特殊字符 !@#$%^&*()_+\-=[]{};':"\\|,.<>/?等

可以包含其他字符（如字母、数字、空格、中文等）

```api
type (
	demo {
        Name string `form:"name,optional" validate:"xStrWithoutSpec=8-15"` // 姓名
    }
)
```

返回错误为

```text
{0}长度{1}，首尾不能有空格，不能包含特殊字符串
```
{0}为参数名称

{1}为param即8-15

### xStrWithoutSpec
不能有空格，不能包含特殊字符串的字符串校验

使用方法为 validate:"xStrWithoutSpecAndSpace=1-300" 代表字符串长度为 1-300位，左右都为闭区间
长度自定义
不能有空格
不能包含特殊字符 !@#$%^&*()_+\-=[]{};':"\\|,.<>/?等
可以包含其他字符（如字母、数字、中文等）。

```api
type (
	demo {
        Name string `form:"name,optional" validate:"xStrWithoutSpec=8-15"` // 姓名
    }
)
```

返回错误为

```text
{0}长度{1}，不能有空格，不能包含特殊字符串
```
{0}为参数名称

{1}为param即8-15