package main

import(
    // "net"
    "crud-echo/application"
)
func main(){
//     config.ConnectDB()
//     e := routes.Init()
//     l, err := net.Listen("tcp", ":9094")
//     if err != nil {
//       e.Logger.Fatal(err)
//     }
//     e.Listener = l
//     e.Logger.Fatal(e.Start(""))
application.StartApp()
}

