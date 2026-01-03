# go-toolkit

一个功能丰富的 Go 语言工具库，提供了常用的工具函数和实用功能。

## 功能特性

- 🔧 **类型转换** - 提供各种类型之间的转换工具
- 📁 **文件操作** - CSV 读写、文件操作等实用功能
- 🔤 **字符串处理** - 丰富的字符串操作函数
- ⏰ **时间处理** - 时间格式化、计算等工具
- 📊 **切片操作** - 切片过滤、映射、去重等功能
- 🗺️ **Map 操作** - Map 的常用操作工具
- 🔐 **加密工具** - MD5、SHA1、SHA256、SHA512 等哈希函数
- 🔄 **并发控制** - 信号量等并发控制工具
- 📄 **JSON 操作** - JSON 路径查找、值转换、路径操作等工具

## 安装

```bash
go get github.com/cx-luo/go-toolkit
```

## 使用示例

### 类型转换 (convert)

```go
import "github.com/cx-luo/go-toolkit/convert"

// 转换为字符串
str := convert.ToString(123)        // "123"
str = convert.ToString(3.14)        // "3.14"

// 转换为整数
num := convert.ToInt("123")         // 123
num = convert.ToInt(3.14)           // 3

// 转换为 int64
num64 := convert.ToInt64("123456")  // 123456

// 转换为 float64
f := convert.ToFloat64("3.14")      // 3.14

// 转换为 bool
b := convert.ToBool("true")         // true
```

### 字符串处理 (stringutil)

```go
import "github.com/cx-luo/go-toolkit/stringutil"

// 检查是否为空
isEmpty := stringutil.IsEmpty("")   // true

// 字符串反转
reversed := stringutil.Reverse("hello")  // "olleh"

// 驼峰转蛇形
snake := stringutil.CamelToSnake("HelloWorld")  // "hello_world"

// 蛇形转驼峰
camel := stringutil.SnakeToCamel("hello_world")  // "HelloWorld"

// 生成随机字符串
random, _ := stringutil.RandomString(16)

// 截断字符串
truncated := stringutil.TruncateWithEllipsis("very long string", 10)  // "very lo..."

// 检查是否为数字
isNum := stringutil.IsNumeric("123")  // true
```

### 时间处理 (timeutil)

```go
import "github.com/cx-luo/go-toolkit/timeutil"

// 获取当前时间
now := timeutil.Now()

// 格式化时间
formatted := timeutil.Format(now, timeutil.FormatDateTime)  // "2006-01-02 15:04:05"

// 解析时间
t, _ := timeutil.Parse("2006-01-02 15:04:05", timeutil.FormatDateTime)

// 获取今天的开始时间
today := timeutil.Today()

// 获取本周的开始和结束
weekStart := timeutil.StartOfWeek(now)
weekEnd := timeutil.EndOfWeek(now)

// 获取本月的开始和结束
monthStart := timeutil.StartOfMonth(now)
monthEnd := timeutil.EndOfMonth(now)

// 计算时间差
days := timeutil.DiffDays(t1, t2)
hours := timeutil.DiffHours(t1, t2)

// 判断是否同一天
sameDay := timeutil.IsSameDay(t1, t2)
```

### 切片操作 (slice)

```go
import "github.com/cx-luo/go-toolkit/slice"

// 检查是否包含
contains := slice.Contains([]int{1, 2, 3}, 2)  // true

// 查找索引
index := slice.IndexOf([]string{"a", "b", "c"}, "b")  // 1

// 去重
unique := slice.Unique([]int{1, 2, 2, 3, 3})  // [1, 2, 3]

// 过滤
filtered := slice.Filter([]int{1, 2, 3, 4, 5}, func(x int) bool {
    return x > 3
})  // [4, 5]

// 映射
mapped := slice.Map([]int{1, 2, 3}, func(x int) int {
    return x * 2
})  // [2, 4, 6]

// 归约
sum := slice.Reduce([]int{1, 2, 3}, 0, func(acc, x int) int {
    return acc + x
})  // 6

// 分块
chunks := slice.Chunk([]int{1, 2, 3, 4, 5}, 2)  // [[1, 2], [3, 4], [5]]

// 交集
intersection := slice.Intersect([]int{1, 2, 3}, []int{2, 3, 4})  // [2, 3]

// 并集
union := slice.Union([]int{1, 2}, []int{2, 3})  // [1, 2, 3]

// 差集
diff := slice.Difference([]int{1, 2, 3}, []int{2, 3})  // [1]
```

### Map 操作 (maputil)

```go
import "github.com/cx-luo/go-toolkit/maputil"

m := map[string]int{"a": 1, "b": 2, "c": 3}

// 获取所有键
keys := maputil.Keys(m)  // ["a", "b", "c"]

// 获取所有值
values := maputil.Values(m)  // [1, 2, 3]

// 检查键是否存在
exists := maputil.ContainsKey(m, "a")  // true

// 获取值或默认值
value := maputil.GetOrDefault(m, "d", 0)  // 0

// 合并多个 map
merged := maputil.Merge(map1, map2, map3)

// 过滤
filtered := maputil.Filter(m, func(k string, v int) bool {
    return v > 1
})

// 映射
mapped := maputil.Map(m, func(k string, v int) string {
    return fmt.Sprintf("%s:%d", k, v)
})
```

### 文件操作 (file)

```go
import "github.com/cx-luo/go-toolkit/file"

// 读取 CSV
records, err := file.ReadCSV("data.csv")

// 写入 CSV
err = file.WriteCSV("output.csv", records)

// 读取文件所有行
lines, err := file.ReadLines("file.txt")

// 写入文件所有行
err = file.WriteLines("file.txt", lines)

// 读取整个文件
data, err := file.ReadFile("file.txt")

// 写入文件
err = file.WriteFile("file.txt", data)

// 检查文件是否存在
exists := file.Exists("file.txt")

// 检查是否为目录
isDir := file.IsDir("path")

// 复制文件
err = file.CopyFile("source.txt", "dest.txt")

// 获取文件扩展名
ext := file.GetExt("file.txt")  // ".txt"

// ===== 大型文件处理 =====

// 逐行流式读取（使用回调函数，内存高效）
err = file.ReadLinesStream("large_file.txt", func(line string, lineNum int) error {
    // 处理每一行
    fmt.Println(line)
    // 可以返回错误来停止读取
    return nil
})

// 逐行流式读取（使用通道，适合并发处理）
lines, errChan := file.ReadLinesChannel("large_file.txt", 100)
for line := range lines {
    // 处理每一行
    fmt.Println(line)
}
if err := <-errChan; err != nil {
    // 处理错误
}

// 分块读取（指定偏移量和大小）
chunk, n, err := file.ReadChunk("large_file.bin", 0, 1024*1024)  // 读取1MB

// 分块流式读取（使用回调函数）
err = file.ReadChunksStream("large_file.bin", 1024*1024, func(chunk []byte, offset int64) error {
    // 处理每个块（1MB）
    // offset 是当前块在文件中的位置
    return nil
})

// 分块流式读取（使用通道）
chunks, errChan := file.ReadChunksChannel("large_file.bin", 1024*1024, 10)
for chunk := range chunks {
    // chunk.Data 是块数据
    // chunk.Offset 是块在文件中的偏移量
    // chunk.Size 是块的大小
    fmt.Printf("Offset: %d, Size: %d\n", chunk.Offset, chunk.Size)
}

// 读取指定偏移量的块
chunk, err := file.ReadChunkWithOffset("large_file.bin", 1024*1024, 512*1024)  // 从1MB位置读取512KB

// 读取前N行（适合只处理文件开头的情况）
lines, err := file.ReadLinesWithLimit("large_file.txt", 100)  // 只读取前100行
```

### 加密工具 (crypto)

```go
import "github.com/cx-luo/go-toolkit/crypto"

// MD5 哈希
hash := crypto.MD5("hello")  // "5d41402abc4b2a76b9719d911017c592"

// SHA1 哈希
hash = crypto.SHA1("hello")

// SHA256 哈希
hash = crypto.SHA256("hello")

// SHA512 哈希
hash = crypto.SHA512("hello")

// 使用指定算法
hash, err := crypto.HashString("hello", "sha256")
```

### 并发控制 (concurrency)

```go
import "github.com/cx-luo/go-toolkit/concurrency"

// 创建信号量（最大并发数为 10）
sem := concurrency.NewSemaphore(10)

// 获取许可
sem.Acquire(1)

// 执行任务
go func() {
    defer sem.Release()
    // 执行你的任务
}()

// 使用函数方式
sem.AcquireWithFunc(func(args ...interface{}) {
    // 执行任务
}, arg1, arg2)

// 等待所有任务完成
sem.Wait()
```

### JSON 操作 (jsonutil)

```go
import "github.com/cx-luo/go-toolkit/jsonutil"

jsonStr := `{
    "user": {
        "name": "John",
        "age": 30,
        "items": [
            {"id": 1, "name": "item1"},
            {"id": 2, "name": "item2"}
        ]
    }
}`

var data interface{}
json.Unmarshal([]byte(jsonStr), &data)

// 将所有值转换为字符串
converted, err := jsonutil.ConvertValuesToString(data)
// 结果: {"user":{"name":"John","age":"30","items":[{"id":"1","name":"item1"},...]}}

// 从JSON字符串转换所有值为字符串
convertedStr, err := jsonutil.ConvertJSONStringValuesToString(jsonStr)

// 根据路径获取值
name, err := jsonutil.GetValueByPath(data, "user.name")           // "John"
age, err := jsonutil.GetIntByPath(data, "user.age")               // 30
itemName, err := jsonutil.GetStringByPath(data, "user.items[0].name")  // "item1"

// 检查路径是否存在
exists := jsonutil.HasPath(data, "user.name")  // true
exists = jsonutil.HasPath(data, "user.email")  // false

// 设置路径的值
err = jsonutil.SetValueByPath(data, "user.name", "Jane")

// 获取所有路径
allPaths := jsonutil.GetAllPaths(data)
// 结果: ["user", "user.name", "user.age", "user.items", "user.items[0]", ...]

// 查找路径（根据条件）
options := &jsonutil.FindOptions{
    KeyPattern: "name",  // 查找所有包含"name"的键
}
paths, err := jsonutil.FindPaths(data, options)
// 结果: ["user.name", "user.items[0].name", "user.items[1].name"]

// 查找特定值的路径
options = &jsonutil.FindOptions{
    ExactValue: "John",  // 查找值为"John"的路径
}
paths, err = jsonutil.FindPaths(data, options)
// 结果: ["user.name"]

// 查找特定类型的路径
options = &jsonutil.FindOptions{
    ValueType: "number",  // 查找所有数字类型的路径
}
paths, err = jsonutil.FindPaths(data, options)
// 结果: ["user.age", "user.items[0].id", "user.items[1].id"]
```

## 模块说明

- `convert` - 类型转换工具
- `stringutil` - 字符串处理工具
- `timeutil` - 时间处理工具
- `slice` - 切片操作工具
- `maputil` - Map 操作工具
- `file` - 文件操作工具
- `crypto` - 加密工具
- `concurrency` - 并发控制工具
- `jsonutil` - JSON 操作工具

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

## 作者

chengxiang.luo
