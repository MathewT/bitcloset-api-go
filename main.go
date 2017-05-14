package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

var argDebug = flag.Bool("debug", false, "run in debug mode")
var argLogJSON = flag.Bool("logjson", false, "output log in JSON format")

// Info is the info logger
var Info *log.Logger

// Error is the error logger
var Error *log.Logger

const logo = `

bbbbbbbb                                                                                                                                                                                                          
b::::::b              iiii          tttt                                   lllllll                                                                tttt                                                      iiii  
b::::::b             i::::i      ttt:::t                                   l:::::l                                                             ttt:::t                                                     i::::i 
b::::::b              iiii       t:::::t                                   l:::::l                                                             t:::::t                                                      iiii  
 b:::::b                         t:::::t                                   l:::::l                                                             t:::::t                                                            
 b:::::bbbbbbbbb    iiiiiiittttttt:::::ttttttt             cccccccccccccccc l::::l    ooooooooooo       ssssssssss       eeeeeeeeeeee    ttttttt:::::ttttttt           aaaaaaaaaaaaa  ppppp   ppppppppp   iiiiiii 
 b::::::::::::::bb  i:::::it:::::::::::::::::t           cc:::::::::::::::c l::::l  oo:::::::::::oo   ss::::::::::s    ee::::::::::::ee  t:::::::::::::::::t           a::::::::::::a p::::ppp:::::::::p  i:::::i 
 b::::::::::::::::b  i::::it:::::::::::::::::t          c:::::::::::::::::c l::::l o:::::::::::::::oss:::::::::::::s  e::::::eeeee:::::eet:::::::::::::::::t           aaaaaaaaa:::::ap:::::::::::::::::p  i::::i 
 b:::::bbbbb:::::::b i::::itttttt:::::::tttttt         c:::::::cccccc:::::c l::::l o:::::ooooo:::::os::::::ssss:::::se::::::e     e:::::etttttt:::::::tttttt                    a::::app::::::ppppp::::::p i::::i 
 b:::::b    b::::::b i::::i      t:::::t               c::::::c     ccccccc l::::l o::::o     o::::o s:::::s  ssssss e:::::::eeeee::::::e      t:::::t                   aaaaaaa:::::a p:::::p     p:::::p i::::i 
 b:::::b     b:::::b i::::i      t:::::t               c:::::c              l::::l o::::o     o::::o   s::::::s      e:::::::::::::::::e       t:::::t                 aa::::::::::::a p:::::p     p:::::p i::::i 
 b:::::b     b:::::b i::::i      t:::::t               c:::::c              l::::l o::::o     o::::o      s::::::s   e::::::eeeeeeeeeee        t:::::t                a::::aaaa::::::a p:::::p     p:::::p i::::i 
 b:::::b     b:::::b i::::i      t:::::t    tttttt     c::::::c     ccccccc l::::l o::::o     o::::ossssss   s:::::s e:::::::e                 t:::::t    tttttt     a::::a    a:::::a p:::::p    p::::::p i::::i 
 b:::::bbbbbb::::::bi::::::i     t::::::tttt:::::t     c:::::::cccccc:::::cl::::::lo:::::ooooo:::::os:::::ssss::::::se::::::::e                t::::::tttt:::::t     a::::a    a:::::a p:::::ppppp:::::::pi::::::i
 b::::::::::::::::b i::::::i     tt::::::::::::::t      c:::::::::::::::::cl::::::lo:::::::::::::::os::::::::::::::s  e::::::::eeeeeeee        tt::::::::::::::t     a:::::aaaa::::::a p::::::::::::::::p i::::::i
 b:::::::::::::::b  i::::::i       tt:::::::::::tt       cc:::::::::::::::cl::::::l oo:::::::::::oo  s:::::::::::ss    ee:::::::::::::e          tt:::::::::::tt      a::::::::::aa:::ap::::::::::::::pp  i::::::i
 bbbbbbbbbbbbbbbb   iiiiiiii         ttttttttttt           ccccccccccccccccllllllll   ooooooooooo     sssssssssss        eeeeeeeeeeeeee            ttttttttttt         aaaaaaaaaa  aaaap::::::pppppppp    iiiiiiii
                                                                                                                                                                                       p:::::p                    
                                                                                                                                                                                       p:::::p                    
                                                                                                                                                                                      p:::::::p                   
                                                                                                                                                                                      p:::::::p                   
                                                                                                                                                                                      p:::::::p                   
                                                                                                                                                                                      ppppppppp

`

// BitClosetIndex returns the index for the bitcloset resource
func BitClosetIndex(c *gin.Context) {

}

func main() {
	flag.Parse()

	// Print logo.
	fmt.Print(logo)

	// Initialize the Loggers.  First Info Logger
	Info = log.New()
	Info.Out = os.Stdout

	// Initialize Error Logger
	Error = log.New()
	Error.Out = os.Stderr

	if *argDebug {
		log.SetLevel(log.DebugLevel)
	}

	if *argLogJSON {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	}

	router := gin.Default()
	v1 := router.Group("api/v1/bitcloset")
	{
		v1.GET("/", BitClosetIndex)
	}

	router.Run()

}
