package main
import (
  "fmt"
  "math"
  "time"
  "math/rand"
)

// define security system (ss) components

type ss_Component struct {
  ss_id string
  devicetype string
  serial_no string
  name string
  status string
  location_x int
  location_y int
}

type ss_Event struct {
  ss_delta_time int64
  ss_event_function func()
  ss_event_name string
}

var thief_location string;

  var network = []*ss_Component {
	{"41011","SX81 Camera","910-770423","lab1-north","motion-detected",15,1},
	{"41012","SX81 Camera","910-771811","lab2-north","idle",19,1},
	{"41013","SX81 Camera","910-768860","manuf-north","idle",22,1},
	{"41014","SX81 Camera","910-768891","docks-1","idle",2,1},
	{"41015","SX81 Camera","910-768312","docks-2","idle",2,11},
	{"41016","SX81 Camera","910-768456","manuf-east","idle",2,30},
	{"41017","SX81 Camera","910-771998","lab1-south","motion-detected",3,12},
	{"41018","SX81 Camera","910-771766","lab2-south","idle",3,16},
	{"41019","SX81 Camera","910-771043","lab3-north","idle",4,12},
	{"41020","SX81 Camera","910-771228","lab4-north","idle",4,16},
	{"41021","SX81 Camera","910-771229","warehouse-east","idle",6,11},
	{"41022","SX81 Camera","910-771230","warehouse-south","idle",8,5},
	{"41023","SX81 Camera","910-771231","lab3-south","idle",8,12},
	{"41024","D9 Door Controller","914-32244","Lab4","auto",8,18},
	{"41025","SX81 Camera","910-770117","manuf-south-1","idle",8,20},
	{"41026","SX81 Camera","910-770228","manuf-south-2","idle",8,26},
	{"41027","SX81 Camera","910-770231","pods-4","idle",9,5},
	{"41028","SX81 Camera","910-770338","server-room-west","idle",9,15},
	{"41029","SX81 Camera","910-770340","server-room-north","idle",9,17},
	{"41030","SX81 Camera","910-770342","server-room-door","idle",9,19},
	{"41031","D9 Door Controller","914-32810","ServerRoom","auto",9,18},
	{"41032","SX81 Camera","910-770343","pods-3","idle",9,26},
	{"41033","SX81 Camera","910-773009","gallery-south","idle",11,16},
	{"41034","SX81 Camera","910-773108","cafeteria-north","idle",12,8},
	{"41035","SX81 Camera","910-773443","meeting-meet","idle",12,14},
	{"41036","SX81 Camera","910-773816","elevators-northwest","idle",12,15},
	{"41037","SX81 Camera","910-773814","elevators-northeast","idle",12,18},
	{"41038","SX81 Camera","910-773815","open-collab-north","idle",12,27},
	{"41039","SX81 Camera","910-773826","cafeteria-east","idle",14,10},
	{"41040","SX81 Camera","910-773827","open-collab-west","idle",14,23},
	{"41041","SX81 Camera","910-770661","lobby-main","idle",15,18},
	{"41042","SX81 Camera","910-770663","pods-1","idle",16,14},
	{"41042","SX81 Camera","910-770665","entryway","idle",16,18},
	{"41042","SX81 Camera","910-770667","pods-2","idle",16,19},
	{"41042","SX81 Camera","910-770669","office-block-west","idle",17,4},
	{"41042","SX81 Camera","910-770671","office-block-east","idle",17,30},
	{"41042","SX81 Camera","910-770673","parklot-front-west-corner","idle",19,1},
	{"41042","SX81 Camera","910-770675","parklot-front-entrance-west","idle",19,14},
	{"41042","SX81 Camera","910-770677","parklot-front-entrance-east","idle",19,19},
	{"41042","SX81 Camera","910-770679","parklot-front-east-corner","idle",19,31},
	{"41042","D9 Door Controller","914-31455","Lab1","auto",3,13},
	{"41042","D9 Door Controller","914-32906","Lab2","auto",3,16},
	
    }


func main() {

var (
   home_base_node	*ss_Component
   remote_chain         []*ss_Component
   user_command		string
)


  fmt.Println("Welcome to Mission RPLC.\nCopyright (c) 2021 by Katin Imes under the GPL 2.0 License.\n")


  thief_location = "pods-2"

  mission_instructions()

  move_thief()


//  for idx, val := range network {
//    fmt.Println(idx,val.name, val.devicetype)
//  }

  home_base_node = get_node_by_name("lobby-main",network)
  remote_chain = append(remote_chain,home_base_node)

//  fmt.Println( home_base_node )

fmt.Println("Current time in the Metalistic Labs Building is 4:09am Pacific Standard Time.")
fmt.Println("Value of Second is:",time.Second)


// main game loop
  for {
    print_prompt( home_base_node, remote_chain )
    fmt.Scanln(&user_command)
    process_cmd(&remote_chain, user_command, network)



  }



}

/////////////////////////////////////////////////
// functions
/////////////////////////////////////////////////


// move_thief()
//
// Choose randomly from the eight locations the thief will pillage
// and move the thief to that spot.
// Then, set a timer (randomized) until his next move.

func move_thief() {

const (
  min_time	= 30
  variable_time	= 60
)

  thief_spots := []string {
    "lab1-south",
    "pods-3",
    "pods-4",
    "manuf-south-1",
    "lab2-south",
    "lab4-south",
    "lab3-south",
    "ServerRoom",
  }    

  var thief_next int
  for {
    thief_next= int(math.Abs( rand.Float64() * 8 ))
    if thief_spots[thief_next] != thief_location {
      break
    }
  }
    thief_location = thief_spots[thief_next]
fmt.Println("<TEST> Thief moved to "+thief_location)

  time_till_move := min_time + math.Abs( rand.Float64() * variable_time )
  timer_seconds := time.Duration(time_till_move) * time.Second
  f := func() {
    move_thief()
  }   

fmt.Println("<TEST> Timer set for",timer_seconds)

  time.AfterFunc( timer_seconds, f)

}



func print_prompt( home_base_node *ss_Component, remote_chain []*ss_Component ) {
  fmt.Print( "(Admin) " + home_base_node.name )
  for idx, val := range remote_chain {
//fmt.Println( val )
    if idx > 0 {
      fmt.Print( "->"+val.name )
    }
  }
  fmt.Print( "> " )
}


func display_error( msg string ) {
  fmt.Println( msg )
}


func get_node_by_name( name string, network []*ss_Component ) *ss_Component {

fmt.Println("looking for:",name)

  for idx, val := range network {
    if val.name == name { 
//      fmt.Println("FOUND IT",idx)
      return network[idx] 
    } 
//fmt.Println(idx,val.name, val.devicetype)
  }
  display_error( "node name "+name+" not found in network." )
  return nil
}


func process_cmd( lmr_chain *[]*ss_Component, user_command string, network []*ss_Component ) {
//  fmt.Println("Command entered: "+user_command)

  remote_chain := *lmr_chain
  var current_node = remote_chain[len(remote_chain)-1]
  switch user_command {
    case "":

    case "help":
      fmt.Println( current_node.devicetype + " commands available:" )
      fmt.Println( "  help .................. displays this screen")
      fmt.Println( "  info .................. displays device information")
      fmt.Println( "  status ................ displays device status information")
      fmt.Println( "  nodes ................. lists visible network components")
      if current_node.devicetype == "SX81 Camera" {
        fmt.Println( "  connect ............... remotely connect to node via RPLC or WiFi")
        fmt.Println( "  exit .................. disconnect from a remote node")
      }
      if current_node.devicetype == "D9 Door Controller" {
        fmt.Println( "  setmode ............... set door mode to auto, lock, or unlock")
      }

     case "info":
       fmt.Println( "     Device ID:  " + current_node.ss_id )
       fmt.Println( "   Device Type:  " + current_node.devicetype )
       fmt.Println( "    Serial No.:  " + current_node.serial_no )
       fmt.Println( "          Name:  " + current_node.name )
       fmt.Println( "        Status:  " + current_node.status )

     case "status":
       fmt.Println( "   WiFi Network:   ***searching...")
       fmt.Println( "     IP Address:   ***lost connection")
       fmt.Println( "           RPLC:   active")
       fmt.Println( "   RPLC Address:   " + current_node.ss_id )
       fmt.Println( "         Status:   " + current_node.status )

     case "nodes":
       display_visible_nodes( current_node, network )

     case "connect":
       var target_name string
       var distance float64
       fmt.Print("Enter node name: ")
       fmt.Scanln(&target_name)
       var remote_node = get_node_by_name(target_name, network) 
       if remote_node != nil {
         distance = math.Abs((float64)(current_node.location_x - remote_node.location_x)) + math.Abs((float64)(current_node.location_y - remote_node.location_y))
         if distance < 7 {
           *lmr_chain = append(remote_chain,remote_node)
           time.Sleep(1)
           fmt.Println("*** Connection to "+remote_node.name+" successful. ***")
         }
       }

     case "exit":
       if len(*lmr_chain) > 1 {
         var my_chain = *lmr_chain
         my_chain = my_chain[:len(my_chain)-1]
         *lmr_chain = my_chain
         fmt.Println("--- disconnected --")
       } else {
         fmt.Println("Cannot exit. No connection to remote node is active.")
       }


     case "setmode":
      if current_node.devicetype == "D9 Door Controller" {
        var network_node *ss_Component
        var mode string
 
        network_node = get_node_by_name(current_node.name,network)

        fmt.Print("Enter mode [auto, lock, unlock]:")
        fmt.Scanln(&mode)
        switch mode {
          case "auto":
            network_node.status = "auto"

          case "lock":
            network_node.status = "locked"

          case "unlock":
            network_node.status = "unlocked"
 
          default:
            fmt.Println("Unrecognized mode.")
        }
      }

 
    default:
      fmt.Println( "Unrecognized command." )
  }


}

func display_visible_nodes( current_node *ss_Component, network []*ss_Component ) {

var distance float64

  for idx, val := range network {
    distance = math.Abs((float64)(current_node.location_x - val.location_x)) + math.Abs((float64)(current_node.location_y - val.location_y))
    if distance < 7 && val != current_node {
      fmt.Println("    "+val.name, "   ["+val.devicetype+"]")
      if (idx < 0 ) {
        }
    }
  }
}


func mission_instructions() {
  fmt.Println("==== Instructions ====\nYour mission, should you choose to accept it, is to use an internal security system of cameras, sensors, and door locking controls to locate, photograph, and trap a thief that has broken into Metalistic Labs, Inc.\n")

  fmt.Println("You have access to a terminal that is connected to a security system camera in the lobby of the building. The problem is, the thief apparently knew that the security system components communicated via WiFi, and he has disabled all WiFi signals in the building by hacking the WiFi access points.\n")

  fmt.Println("This means the security system is mostly non-functional, and cannot be commanded and controlled from the central security center nor by the security system company. Video feeds currently cannot leave the building. Luckily, the security system components do have a legacy communications system still embedded called RPLC. This allows slower communications between components over the A/C power wiring of the building. However, this method only works to a maximum distance of about 100', and the Metalistic Labs building is 600' long by 350' wide. It is a single story building.\n")

  fmt.Println("Finally, you do not have a map of the building or of the security components. We believe the thief will be trying to steal prototype devices from offices and the four labs, as well as trying to break into the main server room to download data or compromise the servers there. All the lhe labs and the server room have electronically lockable doors and no windows, and so if you can communicate with the door lock controllers and command them to 'maglock' while the thief is inside, he will be locked in. This is your goal. The offices do not have such door locks, and can't be used to trap the thief.\n")

  fmt.Println("It's 4am, so no one else is in the building. There are no known pets or active robots in the building, either, so any motion detected by security system cameras or motion detectors should be the thief.\n\nUseful commands include 'help', 'info', and 'status'. Different components of the security system have different functions and commands, so you'll probably want to try the 'help' command on each different component to see what commands they support. You can also type 'mission' to see these instructions, and 'goal' to get a status of your progress. Good luck!\n") 


}



