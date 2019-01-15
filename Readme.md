Just a little cli wrapper for helpful docker and other commands I use often


```
Usage:
   [flags]
   [command]

Available Commands:
  cntr        Docker Container utils
  help        Help about any command
  port        port command utils

Flags:
  -h, --help   help for this command
  
#CNTR
  Usage:
     cntr [flags]
     cntr [command]
  
  Available Commands:
    ka          Kill all running docker containers
    kf          Kill first running docker container
    
    
#Port
    Usage:
       port [flags]
       port [command]
    
    Available Commands:
      kill        Kills all applications with the specified port
      show        Show applications blocking the given port
    
    Flags:
      -h, --help   help for port
    
```