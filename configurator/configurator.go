package configurator

import (
	"github.com/Appliscale/perun/configuration"
	"github.com/Appliscale/perun/context"
	"github.com/Appliscale/perun/myuser"
	"os"
	"strconv"
)

var resourceSpecificationURL = map[string]string{
	"us-east-2":      "https://dnwj8swjjbsbt.cloudfront.net",
	"us-east-1":      "https://d1uauaxba7bl26.cloudfront.net",
	"us-west-1":      "https://d68hl49wbnanq.cloudfront.net",
	"us-west-2":      "https://d201a2mn26r7lk.cloudfront.net",
	"ap-south-1":     "https://d2senuesg1djtx.cloudfront.net",
	"ap-northeast-2": "https://d1ane3fvebulky.cloudfront.net",
	"ap-southeast-1": "https://doigdx0kgq9el.cloudfront.net",
	"ap-southeast-2": "https://d2stg8d246z9di.cloudfront.net",
	"ap-northeast-1": "https://d33vqc0rt9ld30.cloudfront.net",
	"ca-central-1":   "https://d2s8ygphhesbe7.cloudfront.net",
	"eu-central-1":   "https://d1mta8qj7i28i2.cloudfront.net",
	"eu-west-1":      "https://d3teyb21fexa9r.cloudfront.net",
	"eu-west-2":      "https://d1742qcu2c1ncx.cloudfront.net",
	"sa-east-1":      "https://d3c9jyj3w509b0.cloudfront.net",
}

func FileName(context *context.Context) {
	homePath, pathError := myuser.GetUserHomeDir()
	if pathError != nil {
		context.Logger.Error(pathError.Error())
	}
	homePath += "/.config/perun"
	context.Logger.Always("Configure file could be in \n  " + homePath + "\n  /etc/perun")
	var yourPath string
	var yourName string
	context.Logger.GetInput("Your path ", &yourPath)
	context.Logger.GetInput("Filename ", &yourName)
	findFile(yourPath+"/"+yourName, context)
}

func findFile(path string, context *context.Context) {
	context.Logger.Always("File will be created in " + path)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		showRegions(context)
		con := createConfig(context)
		configuration.SaveToFile(con, path, *context.Logger)
	} else {
		context.Logger.Always("File already exists in this path")
	}
}

func showRegions(context *context.Context) {
	regions := makeArrayRegions()
	context.Logger.Always("Regions:")
	for i := 0; i < len(regions); i++ {
		pom := strconv.Itoa(i)
		context.Logger.Always("Number " + pom + " region " + regions[i])
	}
}

func setRegions(context *context.Context) (region string, err bool) {
	var numberRegion int
	context.Logger.GetInput("Choose region", &numberRegion)
	regions := makeArrayRegions()
	if numberRegion >= 0 && numberRegion < 14 {
		region = regions[numberRegion]
		context.Logger.Always("Your region is: " + region)
		err = true
	} else {
		context.Logger.Error("Invalid region")
		err = false
	}
	return
}

func setProfile(context *context.Context) (profile string, err bool) {
	context.Logger.GetInput("Input name of profile", &profile)
	if profile != "" {
		context.Logger.Always("Your profile is: " + profile)
		err = true
	} else {
		context.Logger.Error("Invalid profile")
		err = false
	}
	return
}

func createConfig(context *context.Context) configuration.Configuration {
	myRegion, err := setRegions(context)
	for !err {
		context.Logger.Always("Try again, invalid region")
		myRegion, err = setRegions(context)
	}
	myProfile, err1 := setProfile(context)
	for !err1 {
		context.Logger.Always("Try again, invalid profile")
		myProfile, err1 = setProfile(context)
	}
	myResourceSpecificationURL := resourceSpecificationURL

	myConfig := configuration.Configuration{
		DefaultProfile:        myProfile,
		DefaultRegion:         myRegion,
		SpecificationURL:      myResourceSpecificationURL,
		DefaultDecisionForMFA: false,
		DefaultDurationForMFA: 3600,
		DefaultVerbosity:      "INFO"}

	return myConfig
}

func makeArrayRegions() [14]string {
	var regions [14]string
	regions[0] = "us-east-1"
	regions[1] = "us-east-2"
	regions[2] = "us-west-1"
	regions[3] = "us-west-2"
	regions[4] = "ca-central-1"
	regions[5] = "ca-central-1"
	regions[6] = "eu-west-1"
	regions[7] = "eu-west-2"
	regions[8] = "ap-northeast-1"
	regions[9] = "ap-northeast-2"
	regions[10] = "ap-southeast-1"
	regions[11] = "ap-southeast-2"
	regions[12] = "ap-south-1"
	regions[13] = "sa-east-1"

	return regions
}
