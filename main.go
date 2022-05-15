package main

import (
   "bufio"
   "fmt"
   "net"
   "net/http"
   "os"
   "strconv"
   "strings"
   "time"
)

type data_struct struct { // 结构体
   Channel chan string //数据管道
   Domain string //域名
   Dict string //字典
   Port string //端口
   Num int //子域名爆破总数合
   Taskend int //线程结束总数合
   Consoleid int //输出ID
   Portchanel chan string
} //整体结构体

func (this *data_struct) Sondomain() (resurl string) { //域名生成器
   var buffer=make([]byte,1024*9999)
   file,err:=os.Open(this.Dict)
   if err!=nil{
      fmt.Println("文件不存在")
      os.Exit(0)
   }
   f:=bufio.NewReader(file)
   n,_:=f.Read(buffer)
   url:=string(buffer[:n])
   urlok:=strings.Split(url,"\r\n")
   this.Num++
   if this.Num>=len(urlok){
      return "END"
   }
   if strings.Contains(this.Domain,"http"){
      this.Domain=strings.Replace(this.Domain,"http://","",-1)
   }
   if strings.Contains(this.Domain,"https"){
      this.Domain=strings.Replace(this.Domain,"https://","",-1)
   }
   resurl="http://" +urlok[this.Num]+"."+this.Domain
   return
}// 该方法控制域名的生成

func (this *data_struct)scanport(ip string,port string) {
   _,err:=net.Dial("tcp",ip+":"+port)
   <-time.After(time.Second*1)
   if err!=nil{
   }else{
      fmt.Println(ip+" "+port+" open")
      f, _ := os.OpenFile("port.txt", os.O_APPEND|os.O_CREATE, 0666)
      defer f.Close()
      file := bufio.NewWriter(f)
      file.WriteString(ip+" "+port+" open"+"\r\n")
      file.Flush()
      this.Portchanel<-port+" open"
   }
}

func (this *data_struct) Geturl() { //中枢处理器
Loop:
   for {
      select {
      case value, ok := <-this.Channel:
         if value!="END"{
            if ok {
               c := &http.Client{}
               req, _ := http.NewRequest("GET", value, nil)
               res, err := c.Do(req)
               if err != nil {

               } else {
                  code := res.Status
                  if code == "200 OK" {
                     var url string
                     if strings.Contains(value,"http"){
                        url=strings.Replace(value,"http://","",-1)
                     }
                     if strings.Contains(value,"https"){
                        url=strings.Replace(value,"https://","",-1)
                     }
                     conn,_:=net.Dial("ip:icmp",url)
                     address:=conn.RemoteAddr()
                     this.Consoleid++
                     fmt.Printf("#%d domain:%s ip: %s SUCCESS\n",this.Consoleid,value,address)
                     log:=fmt.Sprintf("#%d domain: %s ip:%s SUCCESS\n",this.Consoleid,value,address)
                     _, err := os.Open("ok.txt")
                     if err != nil {
                        f, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE, 0666)
                        defer f.Close()
                        file := bufio.NewWriter(f)
                        file.WriteString(log+"\r\n")
                        file.Flush()
                     } else {
                        f, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_APPEND, 0666)
                        defer f.Close()
                        file := bufio.NewWriter(f)
                        file.WriteString(log + "\r\n")
                        file.Flush()
                     }

                     ipaddr:=fmt.Sprintf("%s",address)
                     portlist:=strings.Split(this.Port,",")
                     for _,v:=range portlist{
                        go this.scanport(ipaddr,v)
                        for{
                           select {
                           case <-time.After(time.Second * 2):
                              goto end
                           case catch:=<-this.Portchanel:
                              fmt.Println(catch)
                           }
                        }
                        end:
                     }

                  } else {

                  }
                  defer res.Body.Close()
               }
            }
         }else{
            this.Taskend++
            break Loop
         }

      case <-time.After(time.Second*3):

      }

   }

}
func (this *data_struct) Putdomin(thread int)(end int){ //数据投入器
   for{
      this.Channel<-this.Sondomain()
      if this.Taskend==thread{
         end=889
      }
      return

   }
}//投入channel管道里面
func opt()(dict string,url string,thread int,port string){
   defer func() {
      err:=recover()
      if err!=nil{
         fmt.Println("参数不能为空值，错误。BY：wineme")
         os.Exit(0)
      }
   }()
   if len(os.Args)==1{
      fmt.Println(`子域名爆破 护网专用加强版[+++++]
作者:BY WINEME - ANONYMOUSE
BEAUTIFULE ON PRETTY DOG~
[程序.exe] -u URL -f DICTFILE.TXT -t THREAD
博客:https://www.cnblogs.com/wineme/`)
      os.Exit(0)
   }
   for k,v:=range os.Args{
      if v=="-u"{
         url=os.Args[k+1]
      }
      if v=="-t"{
         a,_:=strconv.Atoi(os.Args[k+1])
         thread=a
      }
      if v=="-f"{
         dict=os.Args[k+1]
      }
      if v=="-p"{
         port=os.Args[k+1]
      }
   }
   return
} //该方法收集用户输入的参数
func main() {
   var Channel=make(chan string,50)
   dict,url,thread,port:=opt()
   if dict==""||url==""||thread==0{
      fmt.Println("参数不全，重新填写 BY：wineme")
      os.Exit(0)
   }
   menu:=&data_struct{
      Channel:Channel,
      Domain:url,
      Dict:dict,
      Port:port,
   }
   typeword:="A,N,O,N,Y,M,O,U,S,E,—,—,—,—,—,—,—,—,W,I,N,E,M,E\n"
   start:=strings.Split(typeword,",")
   for i:=0;i<len(start);i++{
      fmt.Printf("%s",start[i])
      <-time.After(time.Millisecond*50)
   }
   typeword2:="2,0,1,3,年,/,/,/,/,/,/,/,/,|,\\,\\,\\,\\,\\,\\,\\,2,0,2,0,年\n"
   start2:=strings.Split(typeword2,",")
   for i:=0;i<len(start2);i++{
      fmt.Printf("%s",start2[i])
      <-time.After(time.Millisecond*50)
   }
   for i:=0;i<thread;i++{
      go menu.Geturl()
   }

   for{
      endok:=menu.Putdomin(thread)
      if endok==889{
         fmt.Println("爆破完毕，携程退出 BY：wineme")
         typeword3:="I+N+ +Y+O+U+N+G+ +I+ +H+A+V+E+ +A+ +D+R+E+A+M+,+I+S+ +H+A+C+K+,+B+E+ +I+N+T+E+R+N+E+T+ +D+E+A+B+E+W+,+F+O+R+ +T+H+I+S+,+I+ +F+O+R+G+E+T+ +T+O+O+ +M+U+C+H+,+L+O+V+E+ +F+R+I+E+N+D+ +A+N+D+ +S+T+U+D+Y+,+B+U+T+ +I+ +D+O+N+'+T+ +R+E+G+R+E+T+,+B+E+C+A+U+S+E+ +I+ +T+H+I+N+K+ +T+H+I+S+ +I+ +P+A+Y+ +I+S+ +B+E+ +W+O+R+T+H+.+.+.+.+.+.+.+.+.+.+.+.+.+."
         start3:=strings.Split(typeword3,"+")
         for i:=0;i<len(start3);i++{
            fmt.Printf("%s",start3[i])
            <-time.After(time.Millisecond*50)
         }
         break
      }
   }
}