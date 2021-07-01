package cmd

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var createSiteCmd = &cobra.Command{
	Use:   "site",
	Short: "Create site in roma site ",
	Long:  `This command Containers`,
	Run: func(cmd *cobra.Command, args []string) {
		createSite()
	},
}

func init()  {
	rootCmd.AddCommand(createSiteCmd)
}

func createSite() {

  url := "https://100.112.22.230:4444/api/v1/sites"
  method := "POST"

  payload := &bytes.Buffer{}
  writer := multipart.NewWriter(payload)
  file, errFile1 := os.Open("./5vm_20210331.json")
  defer file.Close()
  part1,
         errFile1 := writer.CreateFormFile("file",filepath.Base("/root/5vm_20210331.json"))
  _, errFile1 = io.Copy(part1, file)
  if errFile1 != nil {
    fmt.Println(errFile1)
    return
  }
  _ = writer.WriteField("name", "erbay")
  _ = writer.WriteField("remark", "ercan_test")
  _ = writer.WriteField("nic", "eth0")
  _ = writer.WriteField("dbType", "mysql")
  _ = writer.WriteField("port", "3308")
  _ = writer.WriteField("user", "root")
  _ = writer.WriteField("password", "password")
  _ = writer.WriteField("siteTypeId", "1")
  _ = writer.WriteField("privateNetworkAddress", "10.10.10.10")
  _ = writer.WriteField("managementIpAddress", "11.11.11.11")
  _ = writer.WriteField("databaseIpAddress", "10.216.72.32")
  err := writer.Close()
  if err != nil {
    fmt.Println(err)
    return
  }

  tr := &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }
  client := &http.Client {
	  Transport: tr,
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJhZG1pbiIsImF1dGhvcml0eSI6IlJPTEVfQURNSU4iLCJpYXQiOjE2MjUxMjMxNTUsImV4cCI6MTYyNTEyNDk1NX0.YbkI1x6IcN6q6lqWr2KhFZTu_peXa58Zequ5AitqPUIhXRwjCNMty9ITFmGyaFy2JQty9GJqfCef9CPd_lmbfg")
  
  req.Header.Set("Content-Type", writer.FormDataContentType())
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}