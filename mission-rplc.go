package main
import "fmt"


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


func main() {
  fmt.Println("Welcome to Mission RPLC.\nCopyright (c) 2021 by Katin Imes under the GPL 2.0 License.\n")

  network := []*ss_Component {
	{"41011","SX81 Camera","910-770423","lab1-north","motion-detected",15,1},
	{"41012","SX81 Camera","910-770423","lab2-north","idle",19,1},
	{"41013","SX81 Camera","910-770423","manuf-north","idle",22,1},
	{"41014","SX81 Camera","910-770423","docks-1","idle",2,1},
	{"41015","SX81 Camera","910-770423","docks-2","idle",2,11},
	{"41016","SX81 Camera","910-770423","manuf-east","idle",2,30},
	{"41017","SX81 Camera","910-770423","lab1-south","motion-detected",3,12},
	{"41018","SX81 Camera","910-770423","lab2-south","idle",3,16},
	{"41019","SX81 Camera","910-770423","lab3-north","idle",4,12},
	{"41020","SX81 Camera","910-770423","lab4-north","idle",4,16},
	{"41021","SX81 Camera","910-770423","warehouse-east","idle",6,11},
	{"41022","SX81 Camera","910-770423","warehouse-south","idle",8,5},
	{"41023","SX81 Camera","910-770423","lab3-south","idle",8,12},
	{"41024","D9 Door Controller","910-770423","Lab4","auto",8,18},
	{"41025","SX81 Camera","910-770423","manuf-south-1","idle",8,20},
	{"41026","SX81 Camera","910-770423","manuf-south-2","idle",8,26},
	{"41027","SX81 Camera","910-770423","pods-4","idle",9,5},
	{"41028","SX81 Camera","910-770423","server-room-west","idle",9,15},
	{"41029","SX81 Camera","910-770423","server-room-north","idle",9,17},
	{"41030","SX81 Camera","910-770423","server-room-door","idle",9,19},
	{"41031","D9 Door Controller","910-770423","ServerRoom","auto",9,18},
	{"41032","SX81 Camera","910-770423","pods-3","idle",9,26},
	{"41033","SX81 Camera","910-770423","gallery-south","idle",11,16},
	{"41034","SX81 Camera","910-770423","cafeteria-north","idle",12,8},
	{"41035","SX81 Camera","910-770423","meeting-meet","idle",12,14},
	{"41036","SX81 Camera","910-770423","elevators-northwest","idle",12,15},
	{"41037","SX81 Camera","910-770423","elevators-northeast","idle",12,18},
	{"41038","SX81 Camera","910-770423","open-collab-north","idle",12,27},
	{"41039","SX81 Camera","910-770423","cafeteria-east","idle",14,10},
	{"41040","SX81 Camera","910-770423","open-collab-west","idle",14,23},
	{"41041","SX81 Camera","910-770423","lobby-main","idle",15,18},
	{"41042","SX81 Camera","910-770423","pods-1","idle",16,14},
	{"41042","SX81 Camera","910-770423","entryway","idle",16,18},
	{"41042","SX81 Camera","910-770423","pods-2","idle",16,19},
	{"41042","SX81 Camera","910-770423","office-block-west","idle",17,4},
	{"41042","SX81 Camera","910-770423","office-block-east","idle",17,30},
	{"41042","SX81 Camera","910-770423","parklot-front-west-corner","idle",19,1},
	{"41042","SX81 Camera","910-770423","parklot-front-entrance-west","idle",19,14},
	{"41042","SX81 Camera","910-770423","parklot-front-entrance-east","idle",19,19},
	{"41042","SX81 Camera","910-770423","parklot-front-east-corner","idle",19,31},
	{"41042","D9 Door Controller","910-770423","Lab1","auto",3,13},
	{"41042","D9 Door Controller","910-770423","Lab2","auto",3,16},
	
}


  mission_instructions()

	for idx, val := range network {
		fmt.Println(idx,val.name, val.devicetype)
	}




}



func mission_instructions() {
  fmt.Println("==== Instructions ====\nYour mission, should you choose to accept it, is to use an internal security system of cameras, sensors, and door locking controls to locate, photograph, and trap a thief that has broken into Metalistic Labs, Inc.\n")

  fmt.Println("You have access to a terminal that is connected to a security system camera in the lobby of the building. The problem is, the thief apparently knew that the security system components communicated via WiFi, and he has disabled all WiFi signals in the building by hacking the WiFi access points.\n")

  fmt.Println("This means the security system is mostly non-functional, and cannot be commanded and controlled from the central security center nor by the security system company. Video feeds currently cannot leave the building. Luckily, the security system components do have a legacy communications system still embedded called RPLC. This allows slower communications between components over the A/C power wiring of the building. However, this method only works to a maximum distance of about 100', and the Metalistic Labs building is 600' long by 350' wide. It is a single story building.\n")

  fmt.Println("Finally, you do not have a map of the building or of the security components. We believe the thief will be trying to steal prototype devices from offices and the four labs, as well as trying to break into the main server room to download data or compromise the servers there. All the lhe labs and the server room have electronically lockable doors and no windows, and so if you can communicate with the door lock controllers and command them to 'maglock' while the thief is inside, he will be locked in. This is your goal. The offices do not have such door locks, and can't be used to trap the thief.\n")

  fmt.Println("It's 4am, so no one else is in the building. There are no known pets or active robots in the building, either, so any motion detected by security system cameras or motion detectors should be the thief.\n\nUseful commands include 'help', 'info', and 'status'. Different components of the security system have different functions and commands, so you'll probably want to try the 'help' command on each different component to see what commands they support. You can also type 'mission' to see these instructions, and 'goal' to get a status of your progress. Good luck!\n") 


}



