package main

import(
	"errors"
	"fmt" 
	// "encoding/json"
	// "io/ioutil"
	"exporter"
	"exporter/cmd/facebook"
	"exporter/cmd/twitter"
	"exporter/cmd/linkedin"
	"os"
	
)

func main(){
	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	lnkdin := new(linkedin.Linkedin)
	err := export(fb, "fbdata.json")
	err = export(twtr, "twtr.json")
	err = export(lnkdin, "lnkdin.json")

	if err != nil{
		panic(err)
	}

	// twtr := new(twitter.Twitter)
	// lnkdin := new(linkedin.Linkedin)
	
	// ScrollFeeds(fb, twtr, lnkdin)
	}
	
	func ScrollFeeds(platforms ...exporter.SocialMedia){
		for _, sm := range platforms{
			for _, fd := range sm.Feed() {
				fmt.Println(fd)
			}
			fmt.Println("==============================")
		} 
	}


	func export(u exporter.SocialMedia, filename string) error {
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return errors.New("an error occured opening the file: " + err.Error())
		}
		for _, fd := range u.Feed(){
			n, err := f.Write([]byte(fd+"\n"))
			if err != nil{
				return errors.New("an error occured writing to file: " + err.Error())
			}
			fmt.Printf("wrote %d bytes\n", n)
		}
		return nil
	}

