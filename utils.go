package golem

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var cachePath = fmt.Sprintf("%s/.golem", homeDir())
var dataInterface interface{}

// SetCachePath sets the cache path that is used to cache information
// from the pokeapi.
func SetCachePath(path string) {
	// store the old path temporary to be used to copy old cache files to new directory
	oldPath := fixPath(cachePath)
	cachePath = fixPath(path)
	created := createCache()
	if !created {
		fmt.Println("Failed to create cache path!")
	}
	result := moveCache(oldPath, cachePath)
	if result != "" {
		fmt.Println(result)
		fmt.Println("Setting cachepath back to old")
		cachePath = oldPath
	} else {
		fmt.Println("Cache has been set to ", cachePath)
	}

}

func fixPath(path string) string {
	// removes the suffix slash from the path
	return strings.TrimSuffix(path, "/")
}

func createCache() bool {
	// Check to see if the path does not exist
	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		// Create the path
		err = os.MkdirAll(cachePath, os.ModePerm)
		if err != nil {
			return false
		}
		return true
	}
	return true
}

func moveCache(old, new string) string {
	files, err := ioutil.ReadDir(fmt.Sprintf("%s/", old))
	if err != nil {
		return fmt.Sprintf("There was an error trying to read the old cache directory!\nError details: %s", err.Error())
	}
	// check to see if the length is greater than 0
	if len(files) > 0 {
		if runtime.GOOS == "windows" {
			// handle windows copy command here
			err = exec.Command("cmd", "/C", fmt.Sprintf("xcopy %s %s /E /H /K /O /Y", old, new)).Run()
			if err != nil {
				return fmt.Sprintf("There was an error trying to move the old cache files!\nError details: %s", err.Error())
			}

		} else {
			// cp command /should/ work on all other platforms.
			// if not, I'll write a fix in the future as needed.
			err = exec.Command("/bin/sh", "-c", fmt.Sprintf("cp -n -r %s/* %s/", old, new)).Run()
			if err != nil {
				return fmt.Sprintf("There was an error trying to move the old cache files!\nError details: %s", err.Error())
			}
		}

	}
	return ""
}

// GetCachePath returns the cachepath
func GetCachePath() string {
	return cachePath
}

// getDataCache gets data from the cache
func getDataCache(endpoint, searchQuery, searchType string, data interface{}) interface{} {
	path := fmt.Sprintf("%s/%s/json/", cachePath, endpoint)
	f := ""
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false
	}
	for _, file := range files {
		if searchType == "ID" {
			if strings.HasPrefix(file.Name(), fmt.Sprintf("%s_", searchQuery)) {
				f = fmt.Sprintf("%s/%s/json/%s", cachePath, endpoint, file.Name())
			}
		} else if searchType == "Name" {
			if strings.HasSuffix(file.Name(), fmt.Sprintf("%s.json", searchQuery)) {
				f = fmt.Sprintf("%s/%s/json/%s", cachePath, endpoint, file.Name())
			}
		}
	}
	if f == "" {
		return nil
	}
	jsonData, err := os.Open(f)
	if err != nil {
		return nil
	}
	bytes, err := ioutil.ReadAll(jsonData)
	if err != nil {
		return nil
	}
	if json.Unmarshal(bytes, &data) != nil {
		return nil
	}
	return data
}

// getDataAPI gets data from the api
// If there's an error, it will return nil
func getDataAPI(endpoint, searchParameter string, data interface{}) interface{} {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", fmt.Sprintf("https://pokeapi.co/api/v2/%s/%s", endpoint, searchParameter), nil)

	req.Header.Set("User-Agent", "PokeAPIGo/0.1")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to fetch data!")
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		fmt.Println("No information for your query found!")
		return nil
	}
	if json.NewDecoder(resp.Body).Decode(&data) != nil {
		return nil
	}
	return data
}

func storeData(idOnly bool, data interface{}, endpoint string, ID int, objName string) {
	filename := ""
	if idOnly {
		filename = fmt.Sprintf("%s/json/%d.json", endpoint, ID)
	} else {
		filename = fmt.Sprintf("%s/json/%d_%s.json", endpoint, ID, objName)
	}

	// check to see if the dir we want exists.
	if _, err := os.Stat(fmt.Sprintf("%s/%s/json/", cachePath, endpoint)); os.IsNotExist(err) {
		// if it doesn't exist, create it
		os.MkdirAll(fmt.Sprintf("%s/%s/json/", cachePath, endpoint), os.ModePerm)
	}

	jsonData, _ := json.Marshal(data)
	err := ioutil.WriteFile(fmt.Sprintf("%s/%s", cachePath, filename), jsonData, 0644)
	if err != nil {
		fmt.Println("Failed to write cache file!")
	}
}

func homeDir() string {
	dir, _ := os.UserHomeDir()
	return dir
}

// checkCache checks to see if an object is cached or not.
// endpoint is the endpoint of the api, so berry, berry-firmness, etc.
// searchQuery is either an ID or a string, e.g 1 or poke-ball.
// searchType can either be ID or Name.
func checkCache(endpoint, searchQuery, searchType string) bool {
	path := fmt.Sprintf("%s/%s/json/", cachePath, endpoint)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false
	}
	for _, file := range files {
		if searchType == "ID" {
			if strings.HasPrefix(file.Name(), fmt.Sprintf("%s_", searchQuery)) {
				return true
			}
		} else if searchType == "Name" {
			if strings.HasSuffix(file.Name(), fmt.Sprintf("%s.json", searchQuery)) {
				return true
			}
		}
	}

	return false
}

// ClearCache will clear the cache
// WARNING: Only do this if absolutely necessary, you don't want to clear your cache everytime as it will hurt the pokeapi!
// Also, I shouldn't need to say this, but DO NOT set your cache folder to anything important and then run this command as
// it will wipe out anything in that directory, and it will unlikely be recoverable.
// You have been warned!
func ClearCache() {
	if os.RemoveAll(fmt.Sprintf("%s/", cachePath)) != nil {
		fmt.Println("Failed to delete cache!")
	}
}

func idOrString(data string) string {

	if _, err := strconv.Atoi(data); err != nil {
		return "Name"
	}
	return "ID"
}

func getID(url string) int {
	r, _ := regexp.Compile("/[0-9]{1,100}/")
	match := strings.Trim(r.FindString(url), "/")
	id, err := strconv.Atoi(match)
	if err != nil {
		panic(err)
	}
	return id

}

func checkSpriteCache(path string) bool {
	if _, err := os.Stat(fmt.Sprintf("%s%s", cachePath, path)); os.IsNotExist(err) {
		return false
	}
	return true
}

func createSpriteFolders() {
	os.MkdirAll(fmt.Sprintf("%s/sprites/items", cachePath), os.ModePerm)
	os.MkdirAll(fmt.Sprintf("%s/sprites/pokemon/back", cachePath), os.ModePerm)
	os.MkdirAll(fmt.Sprintf("%s/sprites/pokemon/back/female", cachePath), os.ModePerm)
	os.MkdirAll(fmt.Sprintf("%s/sprites/pokemon/back/shiny", cachePath), os.ModePerm)
	os.MkdirAll(fmt.Sprintf("%s/sprites/pokemon/back/shiny/female", cachePath), os.ModePerm)
	os.MkdirAll(fmt.Sprintf("%s/sprites/pokemon/female", cachePath), os.ModePerm)
	os.MkdirAll(fmt.Sprintf("%s/sprites/pokemon/shiny", cachePath), os.ModePerm)
	os.MkdirAll(fmt.Sprintf("%s/sprites/pokemon/shiny/female", cachePath), os.ModePerm)
}

func getSprite(url string) string {
	// regex stuff
	pkmnRegex, _ := regexp.Compile("/sprites/pokemon.{1,100}.[(png | jpg | jpeg | gif)]")
	itemRegex, _ := regexp.Compile("/sprites/items/.{1,100}.[(png | jpg | jpeg | gif)]")
	path := ""
	name := ""
	if pkmnRegex.FindString(url) != "" {
		path = pkmnRegex.FindString(url)
	} else if itemRegex.FindString(url) != "" {
		path = itemRegex.FindString(url)
	}
	// first check to see if the sprite exists
	if checkSpriteCache(path) {
		// if it does, then return the sprite
		file, err := os.Open(fmt.Sprintf("%s%s", cachePath, path))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		name = file.Name()
	} else {
		// download the image
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// create the new file
		file, err := os.Create(fmt.Sprintf("%s%s", cachePath, path))

		if err != nil {
			if os.IsNotExist(err) {
				// run createSpriteFolders
				createSpriteFolders()
				// try again
				file, err = os.Create(fmt.Sprintf("%s%s", cachePath, path))
				if err != nil {
					panic(err)
				}
				defer file.Close()
			} else {
				panic(err)
			}
		}
		defer file.Close()
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			panic(err)
		}
		name = file.Name()
	}
	return name
}
