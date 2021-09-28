package Task
// 硬件参数

// #include <unistd.h>
import "C"

import (
	"fmt"
	"ginvel.com/app/Common"
	"github.com/shirou/gopsutil/v3/cpu"
	_ "github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load" // CPU负载
	"github.com/shirou/gopsutil/v3/mem"  // 内存占用
	_ "github.com/shirou/gopsutil/v3/net"
	"log"
	"os/exec"
	"runtime"
	"time"
)



func Hardware()  {

	// 内存信息
	// total // 内存大小
	// available // 闲置可用内存
	// used // 已使用内存
	// usedPercent // 已使用百分比
	memVirtual, _ := mem.VirtualMemory()
	fmt.Println("虚拟内存free=", memVirtual.Free/1024/1024)
	fmt.Println("虚拟内存used=", memVirtual.Used/1024/1024)
	fmt.Println("虚拟内存UsedPercent=", memVirtual.UsedPercent/1024/1024)

	// 逻辑CPU数量
	cpuNum := runtime.NumCPU()
	fmt.Println("逻辑CPU数量=", cpuNum)
	Common.SetGlobalData("cpu_num", int64(cpuNum))

	// CPU使用率（此插件此功能比较耗时）
	cpuPercent, _ := cpu.Percent(time.Second, false)
	fmt.Println("cpuPercent=", cpuPercent)
	Common.SetGlobalData("cpu_percent", cpuPercent[0])

	// C语言
	println(C.sysconf(C._SC_PHYS_PAGES)*C.sysconf(C._SC_PAGE_SIZE), " bytes")

	// CPU负载（不耗时）
	cpuLoad, _ := load.Avg()
	// {"load1":3.62109375,"load5":2.93408203125,"load15":2.58251953125}
	// load表示每1分钟、5分钟、15分钟的平均队列（平均负载）,值为进程或线程数
	// 具体示意请参考load average：https://blog.csdn.net/bd_zengxinxin/article/details/51781630
	fmt.Printf("CPU负载：%v ", cpuLoad)
	Common.SetGlobalData("cpu_load1", cpuLoad.Load1)
	Common.SetGlobalData("cpu_load5", cpuLoad.Load1)
	Common.SetGlobalData("cpu_load15", cpuLoad.Load1)

	//diskPart, _ := disk.Partitions(true)
	//fmt.Println("diskPart=", diskPart)

	// net IO
	//io, _ := net.IOCounters(true)
	//for index, v := range io {
	//	fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	//}
}




var peers = []string{"f0688165", "f0127595", "f0142720", "f0123261", "f0135467", "f0442370", "f0131822", "f0134867", "f0154294", "f0128559", "f025002", "f0133509", "f0133505", "f01248", "f02770", "f01238", "f070932", "f01231", "f021479", "f0152337", "f016398", "f03176", "f0143858", "f019354", "f07830", "f02303", "f020330", "f02626", "f0133957", "/dns4/bootstrap-0.mainnet.filops.net/tcp/1347/p2p/12D3KooWCVe8MmsEMes2FzgTpt9fXtmCY7wrq91GRiaC8PHSCCBj", "/dns4/bootstrap-1.mainnet.filops.net/tcp/1347/p2p/12D3KooWCwevHg1yLCvktf2nvLu7L9894mcrJR4MsBCcm4syShVc", "/dns4/bootstrap-2.mainnet.filops.net/tcp/1347/p2p/12D3KooWEWVwHGn2yR36gKLozmb4YjDJGerotAPGxmdWZx2nxMC4", "/dns4/bootstrap-3.mainnet.filops.net/tcp/1347/p2p/12D3KooWKhgq8c7NQ9iGjbyK7v7phXvG6492HQfiDaGHLHLQjk7R", "/dns4/bootstrap-4.mainnet.filops.net/tcp/1347/p2p/12D3KooWL6PsFNPhYftrJzGgF5U18hFoaVhfGk7xwzD8yVrHJ3Uc", "/dns4/bootstrap-5.mainnet.filops.net/tcp/1347/p2p/12D3KooWLFynvDQiUpXoHroV1YxKHhPJgysQGH2k3ZGwtWzR4dFH", "/dns4/bootstrap-6.mainnet.filops.net/tcp/1347/p2p/12D3KooWP5MwCiqdMETF9ub1P3MbCvQCcfconnYHbWg6sUJcDRQQ", "/dns4/bootstrap-7.mainnet.filops.net/tcp/1347/p2p/12D3KooWRs3aY1p3juFjPy8gPN95PEQChm2QKGUCAdcDCC4EBMKf", "/dns4/bootstrap-8.mainnet.filops.net/tcp/1347/p2p/12D3KooWScFR7385LTyR4zU1bYdzSiiAb5rnNABfVahPvVSzyTkR", "/dns4/lotus-bootstrap.ipfsforce.com/tcp/41778/p2p/12D3KooWGhufNmZHF3sv48aQeS13ng5XVJZ9E6qy2Ms4VzqeUsHk", "/dns4/bootstrap-0.starpool.in/tcp/12757/p2p/12D3KooWGHpBMeZbestVEWkfdnC9u7p6uFHXL1n7m1ZBqsEmiUzz", "/dns4/bootstrap-1.starpool.in/tcp/12757/p2p/12D3KooWQZrGH1PxSNZPum99M1zNvjNFM33d1AAu5DcvdHptuU7u", "/dns4/node.glif.io/tcp/1235/p2p/12D3KooWBF8cpp65hp2u9LK5mh19x67ftAam84z9LsfaquTDSBpt", "/dns4/bootstrap-0.ipfsmain.cn/tcp/34721/p2p/12D3KooWQnwEGNqcM2nAcPtRR9rAX8Hrg4k9kJLCHoTR5chJfz6d", "/dns4/bootstrap-1.ipfsmain.cn/tcp/34723/p2p/12D3KooWMKxMkD5DMpSWsW7dBddKxKT7L2GgbNuckz9otxvkvByP"}



func ConnectPeer() {
	fmt.Println("关闭主机")
	for _, r := range peers {
		cmt := fmt.Sprintf("/opt/raid0/lotus/bin/lotus net conect %s", r)
		fmt.Println(cmt)
		arg := []string{"-s", "-t", "20"}
		cmd := exec.Command(cmt, arg...)
		d, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("Error:", err)
		}
		fmt.Println(string(d))
	}
}

