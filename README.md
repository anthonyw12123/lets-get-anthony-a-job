# golang
This section is used for learning golang. It's massively parallel and completely clever, so it's a good one to learn.

For the initial data manipulations, I'm using [IMDB's Bulk Data](https://www.imdb.com/interfaces/).

## hello.go
This is my slightly-more-complicated hello world. It reads a file, splits every line using the csv library in the Standard Library, maps it to JSON objects, and prints it to the command line. 

### Future Goals
* Make this an API, move the data to a cloud storage. 
* Possibly make this use a local data engine. Likely SQLite.