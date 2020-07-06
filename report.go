package main

import (
	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/c3m-common/log"
	rpch "github.com/tidusant/chadmin-repo/cuahang"
	"github.com/tidusant/chadmin-repo/models"

	"encoding/json"
	"time"

	"flag"
	"net"
	"net/rpc"
	"strconv"
	"strings"
)

const (
	defaultcampaigncode string = "XVsdAZGVmY"
)

type Arith int

func (t *Arith) Run(data string, result *string) error {
	log.Debugf("Call RPC orders args:" + data)
	*result = ""
	//parse args
	args := strings.Split(data, "|")

	if len(args) < 3 {
		return nil
	}
	var usex models.UserSession
	usex.Session = args[0]
	usex.Action = args[2]
	info := strings.Split(args[1], "[+]")
	usex.UserID = info[0]
	ShopID := info[1]
	usex.Params = ""
	if len(args) > 3 {
		usex.Params = args[3]
	}

	//check shop permission
	shop := rpch.GetShopById(usex.UserID, ShopID)
	if shop.Status == 0 {
		*result = c3mcommon.ReturnJsonMessage("-4", "Shop is disabled.", "", "")
		return nil
	}
	usex.Shop = shop

	if usex.Action == "la" {
		*result = LoadAll(usex)
	} else if usex.Action == "l3" {
		*result = Load3Month(usex)
	} else if usex.Action == "l6" {
		*result = Load6Month(usex)
	} else if usex.Action == "l9" {
		*result = Load9Month(usex)
	} else if usex.Action == "l12" {
		*result = Load12Month(usex)
	} else { //default
		*result = c3mcommon.ReturnJsonMessage("-5", "Action not found.", "", "")
	}

	return nil
}

func LoadAll(usex models.UserSession) string {

	//default current month
	var camps []models.Campaign
	byship := false
	if usex.Params == "true" {
		byship = true
	}
	t := time.Now()
	d, _ := time.ParseDuration(strconv.Itoa(23-t.Hour()) + "h" + strconv.Itoa(60-t.Minute()) + "m")

	camp := rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, 0, -t.Day()).Add(d), t.AddDate(0, 1, -t.Day()).Add(d), byship)
	camp.Name = t.Month().String()
	camps = append(camps, camp)

	info, _ := json.Marshal(camps)
	strrt := string(info)
	return c3mcommon.ReturnJsonMessage("1", "", "success", strrt)
}
func Load3Month(usex models.UserSession) string {

	//default current month
	var camps []models.Campaign
	t := time.Now()
	byship := false
	if usex.Params == "true" {
		byship = true
	}
	d, _ := time.ParseDuration(strconv.Itoa(23-t.Hour()) + "h" + strconv.Itoa(60-t.Minute()) + "m")
	log.Debugf("d %s, checked: %v", d, usex.Params)
	camp := rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, 0, -t.Day()).Add(d), t.AddDate(0, 1, -t.Day()).Add(d), byship)
	camp.Name = t.Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -1, -t.Day()).Add(d), t.AddDate(0, 0, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -1, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -2, -t.Day()).Add(d), t.AddDate(0, -1, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -2, 0).Month().String()
	camps = append(camps, camp)

	info, _ := json.Marshal(camps)
	strrt := string(info)
	return c3mcommon.ReturnJsonMessage("1", "", "success", strrt)
}

func Load6Month(usex models.UserSession) string {

	//default current month
	var camps []models.Campaign
	t := time.Now()
	d, _ := time.ParseDuration(strconv.Itoa(23-t.Hour()) + "h" + strconv.Itoa(60-t.Minute()) + "m")
	byship := false
	if usex.Params == "true" {
		byship = true
	}
	camp := rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, 0, -t.Day()).Add(d), t.AddDate(0, 1, -t.Day()).Add(d), byship)
	camp.Name = t.Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -1, -t.Day()).Add(d), t.AddDate(0, 0, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -1, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -2, -t.Day()).Add(d), t.AddDate(0, -1, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -2, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -3, -t.Day()).Add(d), t.AddDate(0, -2, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -3, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -4, -t.Day()).Add(d), t.AddDate(0, -3, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -4, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -5, -t.Day()).Add(d), t.AddDate(0, -4, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -5, 0).Month().String()
	camps = append(camps, camp)

	info, _ := json.Marshal(camps)
	strrt := string(info)
	return c3mcommon.ReturnJsonMessage("1", "", "success", strrt)
}

func Load9Month(usex models.UserSession) string {

	//default current month
	var camps []models.Campaign
	t := time.Now()
	d, _ := time.ParseDuration(strconv.Itoa(23-t.Hour()) + "h" + strconv.Itoa(60-t.Minute()) + "m")
	byship := false
	if usex.Params == "true" {
		byship = true
	}
	camp := rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, 0, -t.Day()).Add(d), t.AddDate(0, 1, -t.Day()).Add(d), byship)
	camp.Name = t.Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -1, -t.Day()).Add(d), t.AddDate(0, 0, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -1, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -2, -t.Day()).Add(d), t.AddDate(0, -1, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -2, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -3, -t.Day()).Add(d), t.AddDate(0, -2, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -3, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -4, -t.Day()).Add(d), t.AddDate(0, -3, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -4, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -5, -t.Day()).Add(d), t.AddDate(0, -4, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -5, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -6, -t.Day()).Add(d), t.AddDate(0, -5, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -6, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -7, -t.Day()).Add(d), t.AddDate(0, -6, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -7, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -8, -t.Day()).Add(d), t.AddDate(0, -7, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -8, 0).Month().String()
	camps = append(camps, camp)

	info, _ := json.Marshal(camps)
	strrt := string(info)
	return c3mcommon.ReturnJsonMessage("1", "", "success", strrt)
}

func Load12Month(usex models.UserSession) string {

	//default current month
	var camps []models.Campaign
	t := time.Now()
	d, _ := time.ParseDuration(strconv.Itoa(23-t.Hour()) + "h" + strconv.Itoa(60-t.Minute()) + "m")
	byship := false
	if usex.Params == "true" {
		byship = true
	}
	camp := rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, 0, -t.Day()).Add(d), t.AddDate(0, 1, -t.Day()).Add(d), byship)
	camp.Name = t.Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -1, -t.Day()).Add(d), t.AddDate(0, 0, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -1, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -2, -t.Day()).Add(d), t.AddDate(0, -1, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -2, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -3, -t.Day()).Add(d), t.AddDate(0, -2, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -3, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -4, -t.Day()).Add(d), t.AddDate(0, -3, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -4, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -5, -t.Day()).Add(d), t.AddDate(0, -4, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -5, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -6, -t.Day()).Add(d), t.AddDate(0, -5, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -6, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -7, -t.Day()).Add(d), t.AddDate(0, -6, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -7, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -8, -t.Day()).Add(d), t.AddDate(0, -7, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -8, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -9, -t.Day()).Add(d), t.AddDate(0, -8, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -9, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -10, -t.Day()).Add(d), t.AddDate(0, -9, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -10, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), t.AddDate(0, -11, -t.Day()).Add(d), t.AddDate(0, -10, -t.Day()).Add(d), byship)
	camp.Name = t.AddDate(0, -11, 0).Month().String()
	camps = append(camps, camp)

	info, _ := json.Marshal(camps)
	strrt := string(info)
	return c3mcommon.ReturnJsonMessage("1", "", "success", strrt)
}

func main() {
	var port int
	var debug bool
	flag.IntVar(&port, "port", 9888, "help message for flagname")
	flag.BoolVar(&debug, "debug", false, "Indicates if debug messages should be printed in log files")
	flag.Parse()

	// logLevel := log.DebugLevel
	// if !debug {
	// 	logLevel = log.InfoLevel

	// }

	// log.SetOutputFile(fmt.Sprintf("adminReport-"+strconv.Itoa(port)), logLevel)
	// defer log.CloseOutputFile()
	// log.RedirectStdOut()

	//init db
	arith := new(Arith)
	rpc.Register(arith)
	log.Infof("running with port:" + strconv.Itoa(port))

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(port))
	c3mcommon.CheckError("rpc dail:", err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	c3mcommon.CheckError("rpc init listen", err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
