package main
import (
  "fmt"
  "os"
  "strconv"
  "regexp"
  "net"
)

func main() {
  ports := []int{21, 22, 23, 80, 443}

  //check args
  if len(os.Args) != 2 {
    fmt.Println("Usage: ip")
    os.Exit(0)
  }
  //verify os.Arg[1] is an ip addr
  ipreg := "([0-9]{1,3}\\.){3}[0-9]{1,3}"
  if(boolRegex(ipreg, os.Args[1]) != true) {
    fmt.Println("Argument 1 is the IP to scan...")
    os.Exit(0)
  }

  //scan ports
  var results []string
  fmt.Println("Starting scan...")
  for _, port := range ports {
    if (tcpConn(os.Args[1] + ":" + strconv.Itoa(port))) {
      results = append(results, strconv.Itoa(port))
    }
  }

  //Display open ports
  fmt.Println("Open Ports", results)
}

//Bump err return from regex.MatchString
func boolRegex(reg string, content string) bool {
  match, _ := regexp.MatchString(reg, content)
  return match
}

func checkError(err error) bool {
  if err != nil { return true } else { return false }
}

//Bump err, and return bool
func tcpConn(tcpString string) bool {
  fmt.Println(tcpString)
  tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpString)
  if (checkError(err)) {return false} // tcp setup err
  _, err = net.DialTCP("tcp", nil, tcpAddr)
  if (checkError(err)) {return false} else {
    return true
  }
}

