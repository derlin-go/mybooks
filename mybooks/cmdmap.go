package main

import (
    "fmt"
)

/**
 * A command that can be executed from the prompt
 * @param books The map of books (map[string]Book)
 * @param index The current indexes ([]string), i.e. mapping between an integer and 
 *    a title key
 * @param args the arguments from the user, if any
 * @return bool : false upon error, true upon success; 
 *   Index: the new index (if it was modified), nil otherwise
 */ 
type CommandFunc func (books Books, index Index, args ... string) (bool, Index)

/**
 * Simplify the help command management: each command has a function F,
 * a description of the arguments (if any) and a detailed description
 */
type Command struct {
    F CommandFunc  // the function to call
    Args string    // a description of the expected arguments 
    Details string // the detailed description of this command
}

/**
 * the list of available commands. The key is the keyword used by the user
 * in the command prompt. Multiple keywords can refer to the same command.
 */
var cmdMap map[string]Command = map[string]Command{
    "list" : { list, "", "list all the books"},
    "add" : {addBook, "", "add a book interactively. Use the save command afterwards to perenise your changes" },
    "search" : {search, "[author|title|date] word [word,]", "search for specified word. Use one of the keywords author, title or date as first argument to limit your search to this field." },
    "find" : {search, "[author|title|date] word [word,]", "search for specified word. Use one of the keywords author, title or date as first argument to limit your search to this field." },
    "details" : { showDetails, "nbr", "show the details of the book at the specified index number" },
    "show" : { showDetails, "nbr", "show the details of the book at the specified index number." },
    "delete" : { deleteBook, "nbr", "delete the book at the specified index" },
    "save" : { saveFile, "", "save the changes to dropbox. Must explicitely be called !" },
}


/**
 * execute one command, returns false if the command 
 * does not exist.
 */
func runCommand(cmd string, args []string) bool {
    
    if cmd == "help" {
        printHelp(args)
        return true
    }

    if c, ok := cmdMap[cmd]; ok {
        // execute command and update index if needed
        res, idx := c.F(books, index, args...); 
        if !res {
            // print usage ??
        }else if idx != nil {
            // success + index modified, update it
            index  = idx
        }
        return true
    }

    return false
}

/**
 * print a detailed usage if the first argument matches a command. 
 * Otherwise, print the list of commands with their arguments.
 */
func printHelp(args []string) {

    if len(args) == 0 {
        // no extra arg: print the list of available commands
        for k, v := range cmdMap {
            fmt.Printf("  %s %s\n", k, v.Args)
        }
        return
    }
        
    // one command specified: print details if exists.
    var arg0 = args[0]
    if c, ok := cmdMap[arg0]; ok {
        fmt.Printf(" %s %s", arg0,  c.Args);
        fmt.Println("   " + c.Details);

    }else{
        fmt.Printf("command '%s' does not exist. Try 'help'\n", arg0)
    } 

}