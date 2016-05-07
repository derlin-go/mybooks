# mybooks
Command line app to keep track of the books I read. The data are saved in a json file synchronised in Dropbox.
To get more information about the project, have a look at the [Android application](https://github.com/derlin/mybooks-android/) README.

### Prerequisites

- go 1.5+ 
- Dropbox, with default settings (the Dropbox folder in in `${user.dir}/Dropbox`)

### How to use

 - get the source: `go get github.com/derlin-go/mybooks`
 - build the program : `go build -o mybooks src/github.com/derlin/mybooks/*.go`
 - run it: `./mybooks`
 
Once the program is started, you can use `help` to get a list of all the available commands. Don't forget
that the changes are not automatically saved to dropbox. For that, you have to use the command `save`. 

The json file will be kept in `Dropbox/Apps/MyBooks/mybooks.json`. It can be modified manually, as long as the structure is valid.
