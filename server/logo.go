package server

import "fmt"

func LinuxLogoMsg() string {
	msg := fmt.Sprintln("Welcome to TCP-Chat!")
	msg += fmt.Sprintln("         _nnnn_")
	msg += fmt.Sprintln("	dGGGGMMb")
	msg += fmt.Sprintln("       @p~qp~~qMb")
	msg += fmt.Sprintln("       M|@||@) M|")
	msg += fmt.Sprintln("       @,----.JM|")
	msg += fmt.Sprintln("      JS^\\__/  qKL")
	msg += fmt.Sprintln("     dZP        qKRb")
	msg += fmt.Sprintln("    dZP          qKKb")
	msg += fmt.Sprintln("   fZP            SMMb")
	msg += fmt.Sprintln("   HZM            MMMM")
	msg += fmt.Sprintln("   FqM            MMMM")
	msg += fmt.Sprintln(" |    `.       | `' \\Zq")
	msg += fmt.Sprintln("_)      \\.___.,|     .'")
	msg += fmt.Sprintln("\\____   )MMMMMP|   .'")
	msg += fmt.Sprintln("     `-'       `--'")
	return msg
}
