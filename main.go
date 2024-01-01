package main

import "fmt"

func main() {
    inp := "";
    for {
        fmt.Print("$ ");
        fmt.Scan(&inp);
        if inp == "quit" || inp == "exit" {
            break;
        } else if inp == "help" {
            fmt.Println("Useless-SH by Moltony");
            fmt.Println("help - display this message");
            fmt.Println("quit or exit - quit the shell");
        } else {
            fmt.Printf("Useless-SH: %s: command not found\n", inp);
        }
    }
}
