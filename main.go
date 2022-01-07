package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var ws_conn *websocket.Conn

// Devices struct which contains
// an array of devices from the config.json
type Devices struct {
	Devices []Device `json:"devicesList"`
}

// Device struct which contains device info
type Device struct {
	AppiumPort      int    `json:"appium_port"`
	DeviceName      string `json:"device_name"`
	DeviceOSVersion string `json:"device_os_version"`
	DeviceUDID      string `json:"device_udid"`
	WdaMjpegPort    int    `json:"wda_mjpeg_port"`
	WdaPort         int    `json:"wda_port"`
}

// ProjectConfig struct which contains the project configuration values
type ProjectConfig struct {
	DevicesHost             string `json:"devices_host"`
	SeleniumHubHost         string `json:"selenium_hub_host"`
	SeleniumHubPort         string `json:"selenium_hub_port"`
	SeleniumHubProtocolType string `json:"selenium_hub_protocol_type"`
	WdaBundleID             string `json:"wda_bundle_id"`
}

type ProjectConfigPageData struct {
	WebDriverAgentProvided bool
	SudoPasswordSet        bool
	UdevIOSListenerStatus  string
	ImageStatus            string
	ProjectConfigValues    ProjectConfig
}

type ContainerRow struct {
	ContainerID     string
	ImageName       string
	ContainerStatus string
	ContainerPorts  string
	ContainerName   string
	DeviceUDID      string
}

// Load the initial page
func GetInitialPage(w http.ResponseWriter, r *http.Request) {
	var index = template.Must(template.ParseFiles("static/index.html"))
	if err := index.Execute(w, nil); err != nil {
		log.WithFields(log.Fields{
			"event": "index_page_load",
		}).Error("Couldn't load index.html")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Load the initial page with the project configuration info
func GetProjectConfigurationPage(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("./configs/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var projectConfig ProjectConfig
	json.Unmarshal(byteValue, &projectConfig)
	var configRow = ProjectConfig{
		DevicesHost:             projectConfig.DevicesHost,
		SeleniumHubHost:         projectConfig.SeleniumHubHost,
		SeleniumHubPort:         projectConfig.SeleniumHubPort,
		SeleniumHubProtocolType: projectConfig.SeleniumHubProtocolType,
		WdaBundleID:             projectConfig.WdaBundleID}

	var index = template.Must(template.ParseFiles("static/project_config.html"))
	pageData := ProjectConfigPageData{WebDriverAgentProvided: CheckWDAProvided(), SudoPasswordSet: CheckSudoPasswordSet(), UdevIOSListenerStatus: UdevIOSListenerState(), ImageStatus: ImageExists(), ProjectConfigValues: configRow}
	if err := index.Execute(w, pageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateProjectConfigHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request_config ProjectConfig
	err := decoder.Decode(&request_config)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonFile, err := os.Open("./configs/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	if request_config.DevicesHost != "" {
		result["devices_host"] = request_config.DevicesHost
	}
	if request_config.SeleniumHubHost != "" {
		result["selenium_hub_host"] = request_config.SeleniumHubHost
	}
	if request_config.SeleniumHubPort != "" {
		result["selenium_hub_port"] = request_config.SeleniumHubPort
	}
	if request_config.SeleniumHubProtocolType != "" {
		result["selenium_hub_protocol_type"] = request_config.SeleniumHubProtocolType
	}
	if request_config.WdaBundleID != "" {
		result["wda_bundle_id"] = request_config.WdaBundleID
	}

	byteValue, err = json.Marshal(result)
	if err != nil {
		panic(err)
	}

	// Prettify the json so it looks good inside the file
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, []byte(byteValue), "", "  ")

	err = ioutil.WriteFile("./configs/config.json", []byte(prettyJSON.String()), 0644)
	if err != nil {
		panic(err)
	}
}

func UpdateProjectConfigHandler2(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	devices_host := gjson.Get(string(requestBody), "devices_host").Str
	selenium_hub_host := gjson.Get(string(requestBody), "selenium_hub_host").Str
	selenium_hub_port := gjson.Get(string(requestBody), "selenium_hub_port").Str
	selenium_hub_protocol_type := gjson.Get(string(requestBody), "selenium_hub_protocol_type").Str
	wda_bundle_id := gjson.Get(string(requestBody), "wda_bundle_id").Str
	// Open the configuration json file
	jsonFile, err := os.Open("./configs/config.json")
	if err != nil {
		JSONError(w, "config_file_error", "Could not open the config.json file.", 500)
	}
	defer jsonFile.Close()

	// Read the configuration json file into byte array
	configJson, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		JSONError(w, "config_file_error", "Could not read the config.json file.", 500)
	}

	var updatedJSON string
	updatedJSON, _ = sjson.Set(string(configJson), "devicesList.-1", devices_host)

	if devices_host != "" {
		updatedJSON, _ = sjson.Set(string(configJson), "devices_host", devices_host)
	}
	if selenium_hub_host != "" {
		updatedJSON, _ = sjson.Set(string(configJson), "selenium_hub_host", selenium_hub_host)
	}
	if selenium_hub_port != "" {
		updatedJSON, _ = sjson.Set(string(configJson), "selenium_hub_port", selenium_hub_port)
	}
	if selenium_hub_protocol_type != "" {
		updatedJSON, _ = sjson.Set(string(configJson), "selenium_hub_protocol_type", selenium_hub_protocol_type)
	}
	if wda_bundle_id != "" {
		updatedJSON, _ = sjson.Set(string(configJson), "wda_bundle_id", wda_bundle_id)
	}

	// Prettify the json so it looks good inside the file
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, []byte(updatedJSON), "", "  ")

	err = ioutil.WriteFile("./configs/config.json", []byte(prettyJSON.String()), 0644)
	if err != nil {
		panic(err)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func setLogging() {
	log.SetFormatter(&log.JSONFormatter{})
	f, err := os.OpenFile("./logs/project.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
}

func GetLogsPage(w http.ResponseWriter, r *http.Request) {
	var logs_page = template.Must(template.ParseFiles("static/project_logs.html"))
	if err := logs_page.Execute(w, nil); err != nil {
		log.WithFields(log.Fields{
			"event": "project_logs_page",
		}).Error("Couldn't load project_logs.html")
		return
	}
}

func GetLogs(w http.ResponseWriter, r *http.Request) {
	// Execute the command to restart the container by container ID
	commandString := "tail -n 1000 ./logs/project.log"
	cmd := exec.Command("bash", "-c", commandString)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.WithFields(log.Fields{
			"event": "get_project_logs",
		}).Error("Attempted to get project logs but no logs available.")
		fmt.Fprintf(w, "No logs available")
	}
	//SimpleJSONResponse(w, "get_project_logs", out.String(), 200)
	fmt.Fprintf(w, out.String())
}

func handleRequests() {
	// Create a new instance of the mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	// iOS containers endpoints
	myRouter.HandleFunc("/ios-containers/{device_udid}/create", CreateIOSContainer)
	myRouter.HandleFunc("/ios-containers/update", UpdateIOSContainers).Methods("POST")

	// Android containers endpoints

	// General containers endpoints
	myRouter.HandleFunc("/containers/{container_id}/restart", RestartContainer).Methods("POST")
	myRouter.HandleFunc("/containers/{container_id}/remove", RemoveContainer).Methods("POST")
	myRouter.HandleFunc("/containers/{container_id}/logs", GetContainerLogs).Methods("GET")

	// Configuration endpoints
	myRouter.HandleFunc("/configuration/build-image", BuildDockerImage).Methods("POST")
	myRouter.HandleFunc("/configuration/remove-image", RemoveDockerImage).Methods("POST")
	myRouter.HandleFunc("/configuration/setup-ios-listener", SetupUdevListener).Methods("POST")
	myRouter.HandleFunc("/configuration/remove-ios-listener", RemoveUdevListener).Methods("POST")
	myRouter.HandleFunc("/configuration/update-config", UpdateProjectConfigHandler2).Methods("PUT")
	myRouter.HandleFunc("/configuration/set-sudo-password", SetSudoPassword).Methods("PUT")
	myRouter.HandleFunc("/configuration/upload-wda", UploadWDA).Methods("POST")
	myRouter.HandleFunc("/configuration/upload-app", UploadApp).Methods("POST")

	// Devices endpoints
	myRouter.HandleFunc("/device/{device_udid}", ReturnDeviceInfo).Methods("GET")
	myRouter.HandleFunc("/device-logs/{log_type}/{device_udid}", GetDeviceLogs).Methods("GET")
	myRouter.HandleFunc("/ios-devices", GetConnectedIOSDevices).Methods("GET")
	myRouter.HandleFunc("/ios-devices/register", RegisterIOSDevice).Methods("POST")

	// Logs
	myRouter.HandleFunc("/project-logs", GetLogs).Methods("GET")

	// Asset endpoints
	myRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	myRouter.PathPrefix("/main/").Handler(http.StripPrefix("/main/", http.FileServer(http.Dir("./"))))

	// Page loads
	myRouter.HandleFunc("/configuration.html", GetProjectConfigurationPage)
	myRouter.HandleFunc("/android-containers.html", getAndroidContainers)
	myRouter.HandleFunc("/ios-containers.html", GetIOSContainers)
	myRouter.HandleFunc("/project-logs.html", GetLogsPage)
	myRouter.HandleFunc("/", GetInitialPage)

	// Test endpoints
	myRouter.HandleFunc("/test", CreateIOSContainer)

	//log.Fatal(http.ListenAndServeTLS(":10000", "ca-cert.pem", "ca-key.pem", myRouter))
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	setLogging()
	handleRequests()
}
