package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var todos []string

var rootCmd = &cobra.Command{
	Use: "app",
	Short: "This is my ToDo application",
}

var todoCmd = &cobra.Command{
	Use: "todo",
	Short: "Add a new todo",
	Run: func(cmd *cobra.Command, args []string) {
		f, _ := os.OpenFile("todos.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		f.WriteString(args[0] + "\n")
		fmt.Println("Adding a new todo:", args[0])
	},
}

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		f, _ := os.Open("todos.txt")
		defer f.Close()
		scanner := bufio.NewScanner(f)
		fmt.Println("Here are all your todos:")
		i := 1
		for scanner.Scan() {
			fmt.Printf("%d: %s\n", i, strings.TrimSpace(scanner.Text()))
			i++
		}
	},
}

func main() {
	rootCmd.AddCommand(todoCmd)
	rootCmd.AddCommand(listCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}