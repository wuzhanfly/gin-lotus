package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

var peers = []string{"f0688165", "f0127595", "f0142720", "f0123261", "f0135467", "f0442370", "f0131822", "f0134867", "f0154294", "f0128559", "f025002", "f0133509", "f0133505", "f01248", "f02770", "f01238", "f070932", "f01231", "f021479", "f0152337", "f016398", "f03176", "f0143858", "f019354", "f07830", "f02303", "f020330", "f02626", "f0133957", "/dns4/bootstrap-0.mainnet.filops.net/tcp/1347/p2p/12D3KooWCVe8MmsEMes2FzgTpt9fXtmCY7wrq91GRiaC8PHSCCBj", "/dns4/bootstrap-1.mainnet.filops.net/tcp/1347/p2p/12D3KooWCwevHg1yLCvktf2nvLu7L9894mcrJR4MsBCcm4syShVc", "/dns4/bootstrap-2.mainnet.filops.net/tcp/1347/p2p/12D3KooWEWVwHGn2yR36gKLozmb4YjDJGerotAPGxmdWZx2nxMC4", "/dns4/bootstrap-3.mainnet.filops.net/tcp/1347/p2p/12D3KooWKhgq8c7NQ9iGjbyK7v7phXvG6492HQfiDaGHLHLQjk7R", "/dns4/bootstrap-4.mainnet.filops.net/tcp/1347/p2p/12D3KooWL6PsFNPhYftrJzGgF5U18hFoaVhfGk7xwzD8yVrHJ3Uc", "/dns4/bootstrap-5.mainnet.filops.net/tcp/1347/p2p/12D3KooWLFynvDQiUpXoHroV1YxKHhPJgysQGH2k3ZGwtWzR4dFH", "/dns4/bootstrap-6.mainnet.filops.net/tcp/1347/p2p/12D3KooWP5MwCiqdMETF9ub1P3MbCvQCcfconnYHbWg6sUJcDRQQ", "/dns4/bootstrap-7.mainnet.filops.net/tcp/1347/p2p/12D3KooWRs3aY1p3juFjPy8gPN95PEQChm2QKGUCAdcDCC4EBMKf", "/dns4/bootstrap-8.mainnet.filops.net/tcp/1347/p2p/12D3KooWScFR7385LTyR4zU1bYdzSiiAb5rnNABfVahPvVSzyTkR", "/dns4/lotus-bootstrap.ipfsforce.com/tcp/41778/p2p/12D3KooWGhufNmZHF3sv48aQeS13ng5XVJZ9E6qy2Ms4VzqeUsHk", "/dns4/bootstrap-0.starpool.in/tcp/12757/p2p/12D3KooWGHpBMeZbestVEWkfdnC9u7p6uFHXL1n7m1ZBqsEmiUzz", "/dns4/bootstrap-1.starpool.in/tcp/12757/p2p/12D3KooWQZrGH1PxSNZPum99M1zNvjNFM33d1AAu5DcvdHptuU7u", "/dns4/node.glif.io/tcp/1235/p2p/12D3KooWBF8cpp65hp2u9LK5mh19x67ftAam84z9LsfaquTDSBpt", "/dns4/bootstrap-0.ipfsmain.cn/tcp/34721/p2p/12D3KooWQnwEGNqcM2nAcPtRR9rAX8Hrg4k9kJLCHoTR5chJfz6d", "/dns4/bootstrap-1.ipfsmain.cn/tcp/34723/p2p/12D3KooWMKxMkD5DMpSWsW7dBddKxKT7L2GgbNuckz9otxvkvByP"}

func main() {

	//connectPeer()
	// timeout示例，写法v1的相同：
	// 每12s运行一次："@every 12s" 或 "*/12 * * * * *"
	// 每分钟的第0s执行一次："0 */1 * * * *"
	// 每分钟的第5s、25s、45s各执行一次：5,25,45 * * * * *
	// 每12s执行一次：*/12 * * * * *
	// 每隔1分钟的第0秒执行一次：0 */1 * * * *
	// 每天23:30:00执行一次：0 30 23 * * *
	// 每天凌晨1:00:00执行一次：0 0 1 * * *
	// 每月1号早上6:00:00执行一次：0 0 6 1 * *
	// 在每小时的26分、29分、33分各执行一次：0 26,29,33 * * *
	// 每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * *
	// 每天十点到十二点每五秒执行一次：*/5 * 10-12 * * *

	var timeout string = "*/12 * * * * *" // 定时器时间区间，默认精度为30s/次：0,30 * * * * *
	var intervalId int                    // 定时器id
	fmt.Println("comming......")
	go func() {
		num := 0 // 运行次数
		// 设置时区
		local, _ := time.LoadLocation("Asia/Shanghai")
		interval := cron.New(cron.WithLocation(local), cron.WithSeconds()) // 设置时区并且精度按秒。
		_timeout := timeout
		_intervalId, err := interval.AddFunc(_timeout, func() {
			num++
			//log.Println("全局定时器已开启 >>> ", " 定时器ID=", intervalId, " 运行次数num=", num, " 时间区间=", timeout)
			// 下面调用其他函数
			TimeInterval(intervalId, num, _timeout)

		})
		if err != nil {
			log.Println("全局定时器报错：", " error=", err, " num=", num)
			os.Exit(200)
		}
		intervalId = int(_intervalId)
		log.Println("全局定时器已开启 >>>", " 定时器Id=", intervalId, " 默认精度=", _timeout)
		interval.Start()

		//关闭着计划任务, 但是不能关闭已经在执行中的任务.
		//defer interval.Stop()
		//select{} // 阻塞主线程而不退出

	}()



}

// TimeInterval 全局定时器，默认精度30s/次
func TimeInterval(intervalId int, num int, timeout string) {
	var maxLog int = 5
	if num < maxLog { // ���必全部打印，只打印前几个即可
		if intervalId == 0 && num == 0 {
			log.Println("全局定时器初始化完成 >>>", " 定时器ID=", intervalId, " 运行次数num=", num, " 时间区间=", timeout)
		} else {
			log.Println("全局定时器已开启 >>>", " 定时器ID=", intervalId, " 运行次数num=", num, " 时间区间=", timeout)
		}
	} else if num == maxLog {
		log.Println("全局定时器日志显示已关闭，定时任务会继续运行。maxLog=", maxLog)
	}
	// 其他定时任务
	connectPeer() //connect peers
}

func connectPeer() {
	fmt.Println("关闭主机")
	for _, r := range peers {
		cmt := fmt.Sprintf("/opt/raid0/lotus/bin/lotus net conect %s", r)
		fmt.Println(cmt)
		//arg := []string{"-s", "-t", "20"}
		//cmd := exec.Command(cmt, arg...)
		//d, err := cmd.CombinedOutput()
		//if err != nil {
		//	log.Println("Error:", err)
		//	return
		//}
		//fmt.Println(string(d))
		//return

	}
}
