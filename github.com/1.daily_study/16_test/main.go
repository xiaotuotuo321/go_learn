package main

import (
	"fmt"
	"os"
	"testing"
)

// go语言-单元测试

// 不写测试的开发不是好程序员，

// 1.go test 工具：go语言中的测试依赖go test 命令。编写测试代码和编写普通的go代码过程是类似的，并不需要学习新的语法、规则或工具
// go test命令是一个按照一定约定和组织的测试代码的驱动程序。在包目录内，所有以_test.go为后缀的源代码文件都是go test测试的一部分，不会被go build
// 编译到最终的可执行文件中
// 在 *_test.go文件中有三种尅性的函数，单元测试函数、基准测试函数和示例函数
/*
类型						格式								作用
测试函数					函数名前缀为Test					测试程序的一些逻辑行为是否正确
基准函数					函数名前缀为Benchmark				测试函数的性能
示例函数					函数名前缀为Example				为文档提供示例文档

go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数，然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后
清理测试中生成的临时文件
*/

// 2.测试函数
// 2.1.测试函数的格式：每个测试函数必须导入testing包，测试函数的基本格式（签名）如下：
/*
func TestName(t *testing.T) {
	// ...
}
测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头，举几个例子
func TestAdd(t *testing.T){ ... }
func TestSum(t *testing.T){ ... }
func TestLog(t *testing.T){ ... }
其中参数t用于报告测试失败和附加的日志信息。testing.T的拥有的方法如下：
 */
//func (c *T) Error(args ...interface{})
//func (c *T) Errorf(format string, args ...interface{})
//func (c *T) Fail()
//func (c *T) FailNow()
//func (c *T) Failed() bool
//func (c *T) Fatal(args ...interface{})
//func (c *T) Fatalf(format string, args ...interface{})
//func (c *T) Log(args ...interface{})
//func (c *T) Logf(format string, args ...interface{})
//func (c *T) Name() string
//func (t *T) Parallel()
//func (t *T) Run(name string, f func(t *T)) bool
//func (c *T) Skip(args ...interface{})
//func (c *T) SkipNow()
//func (c *T) Skipf(format string, args ...interface{})
//func (c *T) Skipped() bool

// 3.测试函数示例：就像细胞是构成我们身体的基本单位，一个软件程序也是由很多单元组件构成的。单元组件可以是函数体、结构体、方法和最终用户可能依赖的
// 任意东西。总值我们要确保这些组件是能够正常运行的。单元测试是一些利用各种方法测试单元组件的程序，它会将结果与预期输出进行比较
// 见目录中的'split'文件

// 4.测试组：测试一下split函数对中文字符串的支持，这个时候可以再编写一个TestChineseSplit测试函数，但是我们也可以使用如下更友好的一种方式来
// 添加更多的测试实例

// 5.自测试：当如果测试用例较多时，没有办法一眼看出来是哪个测试用例失败了

// 6.测试覆盖率：测试覆盖率是代码被测试套件覆盖的百分比。通常使用的语句是语句的覆盖率，也就是在测试中至少被运行一次的代码占总代码的比例
/*
	go提供内置功能来检查代码覆盖率，可以使用go test -cover来查看测试覆盖率
	go还提供了一个额外的-coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件。这个命令会将覆盖率相关的信息输出到当前文件夹下面的c.out中
	然后我们执行go tool cover -html=c.out, 使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告
	HTML文件中的每个用绿色标记的语句块表示被覆盖了，而红色的表示没有被覆盖
*/

// 7.基准测试
// 7.1.基准测试函数的格式：基准测试就是在一定的工作负载之下检测性能的一种方式。
/*
	func BenchmarkName(b *testing.B){
		// ...
	}

	基准测试以BenchMark为前缀，需要一个*testing.B类型的参数b,基准测试必须要执行b.N次，这样的测试才有对照性，b.N的值是系统根据实际情况去调整的，
	从而保证测试的稳定性。testing.B拥有的方法如下
*/

//func (c *B) Error(args ...interface{})
//func (c *B) Errorf(format string, args ...interface{})
//func (c *B) Fail()
//func (c *B) FailNow()
//func (c *B) Failed() bool
//func (c *B) Fatal(args ...interface{})
//func (c *B) Fatalf(format string, args ...interface{})
//func (c *B) Log(args ...interface{})
//func (c *B) Logf(format string, args ...interface{})
//func (c *B) Name() string
//func (b *B) ReportAllocs()
//func (b *B) ResetTimer()
//func (b *B) Run(name string, f func(b *B)) bool
//func (b *B) RunParallel(body func(*PB))
//func (b *B) SetBytes(n int64)
//func (b *B) SetParallelism(p int)
//func (c *B) Skip(args ...interface{})
//func (c *B) SkipNow()
//func (c *B) Skipf(format string, args ...interface{})
//func (c *B) Skipped() bool
//func (b *B) StartTimer()
//func (b *B) StopTimer()

// 7.2.基准测试示例
/*
$ go test -bench=Split
	goos: darwin
	goarch: amd64
	pkg: go_learn/github.com/1.daily_study/16_test/split
	cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
	BenchmarkSplit-4         5332932               214.6 ns/op
	PASS
	ok      go_learn/github.com/1.daily_study/16_test/split   1.848s

	其中BenchmarkSplit-4 表示对Split函数进行基准测试, 数字4表示 GOMAXPROCES的值，这个对于并发基准测试很重要。5332932 和 214.6 ns/op
	表示每次调用Split函数耗时203ns,这个结果是5332932次调用的平均值

还可以为基准测试添加-benchmem参数，来获得内存分配的统计数据
$ go test -bench=Split -benchmem
	goos: darwin
	goarch: amd64
	pkg: go_learn/github.com/1.daily_study/16_test/split
	cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
	BenchmarkSplit-4         5228787               222.3 ns/op           112 B/op          3 allocs/op
	PASS
	ok      go_learn/github.com/1.daily_study/16_test/split   1.914s

	112 B/op 表示每次操作内存分配了112字节， 3 allocs/op 则表示每次操作进行了3次内存分配。
优化之后的内存分配
$ go test -bench=Split -benchmem
	goos: darwin
	goarch: amd64
	pkg: go_learn/github.com/1.daily_study/16_test/split
	cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
	BenchmarkSplit-4         9046653               133.8 ns/op            48 B/op          1 allocs/op
	PASS
	ok      go_learn/github.com/1.daily_study/16_test/split   1.880s

	使用make函数提前分配内存的改动，减少2/3的内存分配次数，并减少一半的内存分配
*/

// 8.性能比较函数：上面的基准测试只能得到给定操作的绝对耗时，但是在很多性能问题是发生在两个不同操作之间的相对耗时，比如同一个函数处理1000个元素的耗时
// 与处理1万个甚至100万个元素的耗时的差别是多少？再或者对于同一个任务究竟使用哪种算法性能最佳？我们通常需要对两个算法的实现使用相同的输入来进行基准比较测试

// 性能比较函数通常是一个带有参数的函数，被多个笔筒的Benchmark函数传入不同的值来调用

func benchmark(b *testing.B, size int){/* ... */}
func Benchmark10(b *testing.B){ benchmark(b, 10) }
func Benchmark100(b *testing.B){ benchmark(b, 100) }
func Benchmark1000(b *testing.B){ benchmark(b, 1000) }

// 这里需要注意的是，默认情况下，每个基准测试至少运行1秒，如果在Benchmark函数返回时没有到1秒，则b.N的值会按照1，2，5，10，20，50增加，并且函数再次运行
// 最终的BenchmarkFib40只运行了两次，每次运行的平均值只有不到一秒。像这种情况下我们应该可以使用-benchtime标志增加最小基准时间，以产生更准确的结果

// 9.重置时间：b.ResetTimer 之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作

// 10.并行测试：
/*
	func (b *B) RunParallel(body func(*PB)) 会以并行的方式执行给定的基准测试
	RunParallel 会创建出多个goroutine，并将b.N分配给这些goroutine执行，其中goroutine数量的默认值为GOMAXPROCS。用户如果想要增加非CPU受限
	(non-CPU-Bound)基准测试的并行性，那么可以在RunParallel之前调用SetParallelism。RunParallel通常会与-CPU标志一同使用

$ go test -bench=.
	goos: darwin
	goarch: amd64
	pkg: go_learn/github.com/1.daily_study/16_test/split
	cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
	BenchmarkSplit-4                 9273751               128.2 ns/op
	BenchmarkSplitParallel-4        14585432                79.40 ns/op
	PASS
	ok      go_learn/github.com/1.daily_study/16_test/split   28.127s

	还可以通过在测试命令后添加-cpu参数如go test -bench=. -cpu 1来指定使用的CPU数量。
*/

// 11.setup 和 teardown：测试程序时需要在测试前进行额外的设置（setup）或在测试之后进行拆卸（teardown）
// 11.1.testMain：通过*_test.go文件中定义TestMain函数来可以在测试之前进行额外的设置(setup)或在测试之后进行拆卸(teardown)操作
// 如果测试文件包含函数：func TestMain(m *testing.M)那么生成的测试会先调用TestMain(m)，然后再运行具体测试。TestMain运行在主goroutine中
// 可以在调用m.Run前后做任何设置(setup)和拆卸(teardown)。退出测试的时候应该使用m.Run的返回值作为参数调用os.Exit.

func TestMain(m *testing.M){
	fmt.Println("write setup code here ...")// 测试之前的做一些设置
	// 如果 TestMain 使用了flags, 这里应该加上flag.Parse()
	retCode := m.Run()  // 执行测试
	fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
	os.Exit(retCode)	// 退出测试
}
// 需要注意的是：在调用TestMain时，flag.Parse并没有被调用，所以如果TestMan依赖于command-line标志（包括testing包的标记），则应该显示的调用flag.Parse

// 11.2.子测试的setup 与 Teardown: 有时候我们需要为每个测试集设置setup与Teardown，也有可能要为每个子测试设置Setup与Teardown.

// 12.示例函数：
// 12.1.示例函数的格式：被go test特殊对待的第三种函数就是示例函数，他们的函数名以example为前缀，他们既没有参数也没有返回值。
/*
func ExampleName() {
	// ....
}
*/

// 12.2.示例函数示例
/*
为代码编写示例代码有如下三个用处：
	1.示例函数能够作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联
	2.示例函数只要包含了 // Output: 也可以通过go test 运行的可执行测试
	3.示例函数提供了可以直接运行的示例代码，可以直接在golang.org的godoc文档服务器上使用GoPlayground 运行示例代码，下图为strings.ToUpper
 	  函数在Playground的示例函数效果。
*/

