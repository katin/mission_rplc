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
  mode string
  status string
  location_x int
  location_y int
  doorzone string
}

type ss_Event struct {
  ss_delta_time int64
  ss_event_function func()
  ss_event_name string
}

var thief_location string

const (  // iota is an int incrementing with each appearance, reset by const
  REMOTE_CONNECTION = iota
  SERVER_ROOM = iota
  DOOR_CONTROLS = iota
  PHOTOMODE_CAMERA = iota
  SPOTTED_THIEF = iota
  FIVE_PIVOTS_CHAIN = iota
  LOCKED_IN_THIEF = iota
  PHOTOGRAPHED_THIEF = iota
)

var point_value = map[int64]int64 { 
  REMOTE_CONNECTION: 25,
  SERVER_ROOM: 25,
  DOOR_CONTROLS: 25,
  PHOTOMODE_CAMERA: 25,
  SPOTTED_THIEF: 400,
  FIVE_PIVOTS_CHAIN: 500,
  LOCKED_IN_THIEF: 750,
  PHOTOGRAPHED_THIEF: 750,
}

var score_description = map[int64]string {
  REMOTE_CONNECTION: "Successful remote connection",
  SERVER_ROOM: "Connection to Server Room",
  DOOR_CONTROLS: "Operated Door Contols",
  PHOTOMODE_CAMERA: "Activated Motion-Photo Mode",
  SPOTTED_THIEF: "Spotted the Thief",
  FIVE_PIVOTS_CHAIN: "Connection-pivot across five hops or more",
  LOCKED_IN_THIEF: "Locked in the thief",
  PHOTOGRAPHED_THIEF: "Photographed the thief",
}

// this array is used a boolean chart of scores achieved
//
var player_score = map[int64]int64 {
  REMOTE_CONNECTION: 0,
  SERVER_ROOM: 0,
  DOOR_CONTROLS: 0,
  PHOTOMODE_CAMERA: 0,
  SPOTTED_THIEF: 0,
  FIVE_PIVOTS_CHAIN: 0,
  LOCKED_IN_THIEF: 0,
  PHOTOGRAPHED_THIEF: 0,
}



  var network = []*ss_Component {
	{"41011","SX81 Camera","910-770423","lab1-north","streaming","idle",15,1,"Lab1"},
	{"41012","SX81 Camera","910-771811","lab2-north","streaming","idle",19,1,"Lab2"},
	{"41013","SX81 Camera","910-768860","manuf-north","streaming","idle",22,1,""},
	{"41014","SX81 Camera","910-768891","docks-1","streaming","idle",2,1,""},
	{"41015","SX81 Camera","910-768312","docks-2","streaming","idle",2,11,""},
	{"41016","SX81 Camera","910-768456","manuf-east","streaming","idle",2,30,""},
	{"41017","SX81 Camera","910-771998","lab1-south","streaming","idel",3,12,"Lab1"},
	{"41018","SX81 Camera","910-771766","lab2-south","streaming","idle",3,16,"Lab2"},
	{"41019","SX81 Camera","910-771043","lab3-north","streaming","idle",4,12,"Lab3"},
	{"41020","SX81 Camera","910-771228","lab4-north","streaming","idle",4,16,"Lab4"},
	{"41021","SX81 Camera","910-771229","warehouse-east","streaming","idle",6,11,""},
	{"41022","SX81 Camera","910-771230","warehouse-south","streaming","idle",8,5,""},
	{"41023","SX81 Camera","910-771231","lab3-south","streaming","idle",8,12,"Lab3"},
	{"41024","D9 Door Controller","914-32244","Lab4","","auto",8,18,""},
	{"41025","SX81 Camera","910-770117","manuf-south-1","streaming","idle",8,20,""},
	{"41026","SX81 Camera","910-770228","manuf-south-2","streaming","idle",8,26,""},
	{"41027","SX81 Camera","910-770231","pods-4","streaming","idle",9,5,""},
	{"41028","SX81 Camera","910-770338","server-room-west","streaming","idle",9,15,"ServerRoom"},
	{"41029","SX81 Camera","910-770340","server-room-north","streaming","idle",9,17,"ServerRoom"},
	{"41030","SX81 Camera","910-770342","server-room-door","streaming","idle",9,19,"ServerRoom"},
	{"41031","D9 Door Controller","914-32810","ServerRoom","","auto",9,18,""},
	{"41032","SX81 Camera","910-770343","pods-3","streaming","idle",9,26,""},
	{"41033","SX81 Camera","910-773009","gallery-south","streaming","idle",11,16,""},
	{"41034","SX81 Camera","910-773108","cafeteria-north","streaming","idle",12,8,""},
	{"41035","SX81 Camera","910-773443","meeting-meet","streaming","idle",12,14,""},
	{"41036","SX81 Camera","910-773816","elevators-northwest","streaming","idle",12,15,""},
	{"41037","SX81 Camera","910-773814","elevators-northeast","streaming","idle",12,18,""},
	{"41038","SX81 Camera","910-773815","open-collab-north","streaming","idle",12,27,""},
	{"41039","SX81 Camera","910-773826","cafeteria-east","streaming","idle",14,10,""},
	{"41040","SX81 Camera","910-773827","open-collab-west","streaming","idle",14,23,""},
	{"41041","SX81 Camera","910-770661","lobby-main","streaming","idle",15,18,""},
	{"41042","SX81 Camera","910-770663","pods-1","streaming","idle",16,14,""},
	{"41043","SX81 Camera","910-770665","entryway","streaming","idle",16,18,""},
	{"41044","SX81 Camera","910-770667","pods-2","streaming","idle",16,19,""},
	{"41045","SX81 Camera","910-770669","office-block-west","streaming","idle",17,4,""},
	{"41046","SX81 Camera","910-770671","office-block-east","streaming","idle",17,30,""},
	{"41047","SX81 Camera","910-770673","parklot-front-west-corner","streaming","idle",19,1,""},
	{"41048","SX81 Camera","910-770675","parklot-front-entrance-west","streaming","idle",19,14,""},
	{"41049","SX81 Camera","910-770677","parklot-front-entrance-east","streaming","idle",19,19,""},
	{"41050","SX81 Camera","910-770679","parklot-front-east-corner","streaming","idle",19,31,""},
	{"41051","D9 Door Controller","914-31455","Lab1","","auto",3,13,""},
	{"41052","D9 Door Controller","914-32906","Lab2","","auto",3,16,""},
	{"41053","D9 Door Controller","914-32919","Lab3","","auto",3,16,""},
	
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

    // check for win condition
    undone := false
    for i, x := range player_score {
      if player_score[i] == 0 {
        undone = true;
      }
      _=x	// for golang's wonkiness!
    }
    if !undone {
      fmt.Println("=========================================================================")
      fmt.Println("=========================================================================")
      fmt.Println("=========================================================================")
      fmt.Println("     You have won! The thief has been locked in and photographed!        ")
      fmt.Println("=========================================================================")
      fmt.Println("    The police arrive and aarrest the thief.  Well done!")
      fmt.Println("=========================================================================")
      fmt.Println("=========================================================================")
      fmt.Println("=========================================================================")
      break;
    }

  }



}

/////////////////////////////////////////////////
// functions
/////////////////////////////////////////////////


// door_unlocked
// given a nodename
// returns a bool true of the door is auto or unlocked
// works for door nodes and for camera nodes (scans door zone)

func door_unlocked ( nodename string ) bool {

  // if no nodename given then we assume it is an accessible area.
  if nodename == "" {
    return true
  }  

  check_node := get_node_by_name(nodename, network)
  if check_node.devicetype == "D9 Door Controller" {
    if check_node.status == "auto" || check_node.status == "unlocked" {
      return true
    }
  }
  if check_node.devicetype == "SX81 Camera" {
    if check_node.doorzone != "" {
      door_node := get_node_by_name( check_node.doorzone, network )
      if door_node != nil {
        if door_node.status == "auto" || door_node.status == "unlocked" {
          return true
        }
      }      
    }
  }

  return false
}

// thief_lockedin
// Given a network node pointer, check to see if
//  a) we are in a room with a lockable door, and
//  b) if that door is locked.
// Return true if yes to both.

func thief_lockedin( check_node *ss_Component ) bool {
  if check_node.doorzone == "" {
    return false
  }
  if door_unlocked( check_node.doorzone ) {
    return false
  }
  return true
}


// move_thief()
//
// Choose randomly from the eight locations the thief will pillage
// and move the thief to that spot.
// Then, set a timer (randomized) until his next move.

func move_thief() {

  thief_spots := []string {
    // these must be matching, valid names of SX81 cameras in the network
    "lab1-south",
    "pods-3",
    "pods-4",
    "manuf-south-1",
    "lab2-south",
    "lab4-north",
    "lab3-south",
    "server-room-north",
  }    

  current_node := get_node_by_name(thief_location, network)

  // check to see if we are locked in and can't move
  if thief_lockedin(current_node) {
fmt.Println("<TEST> Thief can't move! Door is locked!")
    // just to be all-logic-possible safe, log the locked in event
    score_event(LOCKED_IN_THIEF)
    set_move_timer()
    return
  }

  var thief_next int
  for {
    thief_next= int(math.Abs( rand.Float64() * float64(len(thief_spots)) ))

    // check that this isn't the same as current location
    if thief_spots[thief_next] != thief_location {

      // check for locked doors; they block entry
      if door_unlocked(thief_spots[thief_next]) {
          break; 
      }
    }
  }

  // remove any motion detection status on the current location
  if current_node != nil {
    if current_node.devicetype == "SX81 Camera" {
      current_node.status = "idle"
    }
  }

  thief_location = thief_spots[thief_next]
  check_node := get_node_by_name( thief_location, network )
  if check_node != nil {
    if check_node.devicetype == "SX81 Camera" {
      check_node.status = "motion detected"
    }
  }

fmt.Println("<TEST> Thief moved to "+thief_location)

  set_move_timer()

}


func set_move_timer() {

const (
  min_time	= 30
  variable_time	= 60
)

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
  display_error( "node name '"+name+"' not found in network." )
  return nil
}

// score_event
// Add the points to the player's score for an achievement
// and display a congratulatory message

func score_event( event int64 ) {
  repeat := player_score[event]
  player_score[event] = point_value[event]
  if repeat == 0 {
    fmt.Println("==================================================================================")
    fmt.Println("ACHIEVEMENT",score_description[event]," --",point_value[event],"points awarded.")
    fmt.Println("==================================================================================")
  }
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
        fmt.Println( "  setmode ............... set video capture mode")
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
       if current_node.devicetype == "SX81 Camera" {
         fmt.Println( "          Mode:  " + current_node.mode )
       }
       fmt.Println( "        Status:  " + current_node.status )
       if current_node.status == "motion detected" {
         score_event( SPOTTED_THIEF )
       }


     case "status":
       fmt.Println( "   WiFi Network:   ***searching...")
       fmt.Println( "     IP Address:   ***lost connection")
       fmt.Println( "           RPLC:   active")
       fmt.Println( "   RPLC Address:   " + current_node.ss_id )
       fmt.Println( "         Status:   " + current_node.status )
       if current_node.status == "motion detected" {
         score_event( SPOTTED_THIEF )
       }


     case "score":
       fmt.Println( "Your Current Score: " )
       var total int64 = 0
       for x := range player_score {
         fmt.Println( "    "+prefix_pad_string(score_description[x]+":",55), player_score[x] )
         total = total + player_score[x]
       }
       fmt.Println("  Total score:", total, "points")

     case "nodes":
       display_visible_nodes( current_node, network )

     case "connect":
       if current_node.devicetype == "D9 Door Controller" {
          fmt.Println("Unrecognized command.")
          break
       }
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
	   if len(*lmr_chain) > 5 {
             score_event(FIVE_PIVOTS_CHAIN)
           }
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
      // we're using the status field in the device array for storing the door mode
        var network_node *ss_Component
        var mode string
 
        network_node = get_node_by_name(current_node.name,network)
        if network_node != nil {
          fmt.Print("Enter mode [auto, lock, unlock]:")
          fmt.Scanln(&mode)
          switch mode {
            case "auto":
              network_node.status = "auto"

            case "lock":
              network_node.status = "locked"
              if thief_in_zone( network_node ) {
                score_event( LOCKED_IN_THIEF )
              }

            case "unlock":
              network_node.status = "unlocked"
 
            default:
              fmt.Println("Unrecognized mode.")
          }
        }
      }
      if current_node.devicetype == "SX81 Camera" {
        var network_node *ss_Component
        var mode string
 
        network_node = get_node_by_name(current_node.name,network)

        if network_node != nil {
          fmt.Println("Video modes available are: streaming video, motion-detection video, timed photos every 5 seconds, and photos upon motion detection.")
          fmt.Print("Enter mode [streaming, motdet-video, 5sec-photo, motdet-photo]:")
          fmt.Scanln(&mode)
          switch mode {
            case "streaming":
              network_node.mode = "streaming"

            case "motdet-video":
              network_node.mode = "motdet-video"

            case "5sec-photo":
              network_node.mode = "5sec-photo"
 
            case "motdet-photo":
              network_node.mode = "motdet-photo"

            default:
              fmt.Println("Unrecognized mode.")
          }
        }
      }

 
    default:
      fmt.Println( "Unrecognized command." )
  }


}

// thief_in_zone
// Given a door controller node, check for motion detection on all cameras
// in the room that the door controls.

func thief_in_zone ( doornode *ss_Component ) bool {
  if doornode.devicetype != "D9 Door Controller" {
    display_error(" Device "+doornode.name+" passed to thief_in_zone() check isn't a door.")
    return false;
  }
  for i, node := range network {
    if node.doorzone == doornode.name {
      if node.status == "motion detected" {
fmt.Println("<TEST> THIEF IN ZONE!")
        _=i
        return true
      }
    }
  }
  return false
}


func prefix_pad_string( msg string, desired_size int ) string {
  var padding = desired_size - len(msg)
  var suffix string
  if padding > 0 {
    for x:=1; x<padding; x++ {
      suffix = suffix + " "
      if ( x < 0 ) {}
    }
    return suffix + msg; 
  }
  return msg 
}

func pad_string( msg string, desired_size int ) string {
  var padding = desired_size - len(msg)
  var suffix string
  if padding > 0 {
    for x:=1; x<padding; x++ {
      suffix = suffix + " "
      if ( x < 0 ) {}
    }
    return msg + suffix; 
  }
  return msg 
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

  fmt.Println("Finally, you do not have a map of the building or of the security components. We believe the thief will be trying to steal prototype devices from offices and the four labs, as well as trying to break into the main server room to download data or compromise the servers there. All the lhe labs and the server room have electronically lockable doors and no windows, and so if you can communicate with the door lock controllers and command them to 'maglock' while the thief is inside, he will be locked in. This is your goal. The offices do not have such door locks, and thus can't be used to trap the thief.\n")

  fmt.Println("It's 4am, so no one else is in the building. There are no known pets or active robots in the building, either, so any motion detected by security system cameras or motion detectors should be the thief.\n\nUseful commands include 'help', 'info', and 'status'. Different components of the security system have different functions and commands, so you'll probably want to try the 'help' command on each different component to see what commands they support. You can also type 'mission' to see these instructions, and 'goal' to get a status of your progress. Good luck!\n") 


}



