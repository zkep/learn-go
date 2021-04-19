package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

/*
	gopacket到底是什么呢？是个抓取网络数据包的库，这么说可能还有点抽象，但是抓包工具大家可能都使用过。
	Windows平台下有Wireshark抓包工具，其底层抓包库是npcap（以前是winpcap）；
	Linux平台下有Tcpdump，其抓包库是libpcap；
	而gopacket库可以说是libpcap和npcap的go封装，提供了更方便的go语言操作接口。
	对于抓包库来说，常规功能就是抓包，而网络抓包有以下几个步骤：
	1、枚举主机上网络设备的接口
	2、针对某一网口进行抓包
	3、解析数据包的mac层、ip层、tcp/udp层字段等
	4、ip分片重组，或tcp分段重组成上层协议如http协议的数据
	5、对上层协议进行头部解析和负载部分解析
*/

var (
	device       string = "en0"
	snapshot_len int32  = 1024
	promiscuous  bool   = false
	err          error
	timeout      time.Duration = 30 * time.Second
	handle       *pcap.Handle
)

func main() {
	// 得到所有的(网络)设备
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	// 打印设备信息
	fmt.Println("Devices found:")
	for _, device := range devices {
		//fmt.Println("\nName: ", device.Name)
		//fmt.Println("Description: ", device.Description)
		//fmt.Println("Devices addresses: ", device.Description)
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(string(packet.Data()))
	}
}
