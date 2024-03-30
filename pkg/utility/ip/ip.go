package utils

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gcharset"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

// IpInBlackListRange 判断IP是否在黑名单
func IpInBlackListRange(ip string, ipList []string) (result bool) {
	result = false
	cidr := strings.Join(ipList[:], ",")
	cidrs, err := ParseIpRange(cidr)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	return isInRangeList(ip, cidrs)
}

// isInRange 判断IP是否在指定的范围
// 支持单个IP，支持多个IP，多IP时需要用“,”隔开
// 支持IP段，如192.168.0.1/24
// 支持IP范围，格式如：192.168.1.xx-192.168.1.xx
func isInRange(ip, cidr string) bool {
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}
	return ipnet.Contains(net.ParseIP(ip))
}

// isInRangeList 判断一个 IP 地址是否在多个 CIDR 范围内
func isInRangeList(ip string, cidrs []string) bool {
	for _, cidr := range cidrs {
		if isInRange(ip, cidr) {
			return true
		}
	}
	return false
}

// ParseIpRange 用于将以 "," 分割的多个 CIDR 范围字符串解析为 CIDR 数组。
// 如果一个 CIDR 范围是以 "-" 分割的两个 IP 地址，那么我们会使用 binaryToInt 和 intToIP 函数将它们转换为整数并再次转换为 CIDR 字符串，
// 并将其加入到 CIDR 数组中。
func ParseIpRange(cidr string) ([]string, error) {
	var cidrs []string

	parts := strings.Split(cidr, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if strings.Contains(part, "-") {
			ips := strings.Split(part, "-")
			startIP := net.ParseIP(ips[0]).To4()
			endIP := net.ParseIP(ips[1]).To4()
			if startIP == nil || endIP == nil {
				return nil, fmt.Errorf("无效IP范围: %s", part)
			}
			start := binaryToInt(startIP)
			end := binaryToInt(endIP)
			for i := start; i <= end; i++ {
				startIP := net.ParseIP(intToIP(i).String())
				if startIP == nil {
					return nil, fmt.Errorf("无效IP地址: %s", intToIP(i).String())
				}

				cidrs = append(cidrs, intToIP(i).String()+"/32")
			}
		} else {
			startIP := net.ParseIP(part)
			if startIP == nil {
				return nil, fmt.Errorf("无效IP地址: %s", part)
			}
			cidrs = append(cidrs, part)
		}
	}
	return cidrs, nil
}

func binaryToInt(ip net.IP) int {
	return int(ip[0])<<24 | int(ip[1])<<16 | int(ip[2])<<8 | int(ip[3])
}

func intToIP(n int) net.IP {
	b := make([]byte, 4)
	b[0] = byte(n >> 24)
	b[1] = byte(n >> 16)
	b[2] = byte(n >> 8)
	b[3] = byte(n)
	return net.IPv4(b[0], b[1], b[2], b[3])
}

// 获取客户端IP
func GetClientIp(r *ghttp.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.GetClientIp()
	}
	return ip
}

// 服务端ip
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

// 获取ip所属城市
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "[::1]" || ip == "127.0.0.1" {
		return "内网IP"
	}
	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bytes := g.Client().GetBytes(context.Background(), url)
	src := string(bytes)
	srcCharset := "GBK"
	tmp, _ := gcharset.ToUTF8(srcCharset, src)
	json, err := gjson.DecodeToJson(tmp)
	if err != nil {
		return ""
	}
	if json.Get("code").Int() == 0 {
		city := fmt.Sprintf("%s %s", json.Get("pro").String(), json.Get("city").String())
		return city
	} else {
		return ""
	}
}

// GetPublicIP 获取公网IP
func GetPublicIP() (ip string, err error) {
	resp, err := http.Get("https://ifconfig.co/ip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	ip = string(body)
	// 去除空格
	ip = strings.Replace(ip, " ", "", -1)
	// 去除换行符
	ip = strings.Replace(ip, "\n", "", -1)

	return
}
