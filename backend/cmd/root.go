package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/marjamis/UtilityBelt/internal/kubernetes"
	"github.com/marjamis/UtilityBelt/internal/redis"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "backend",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/", handler)
		http.HandleFunc("/delay/", delay)
		http.Handle("/metrics", promhttp.Handler())
		http.HandleFunc("/redis", redis.RedisHandler)
		http.HandleFunc("/kubernetes", kubernetes.Handler)
		fmt.Println("Listening on port 8081...")
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	var filename string

	if len(title) == 0 {
		filename = "./static/templates/index.html"
	} else {
		filename = "./" + title
		w.Header().Set("Cache-Control", "max-age=86400")
	}

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		http.Error(w, "404 Not Found", http.StatusNotFound)
	} else {
		fmt.Fprintf(w, "%s", body)
	}
	fmt.Printf("page: %s - IP: %s\n", filename, r.RemoteAddr)
	return
}

func delay(w http.ResponseWriter, r *http.Request) {
	di, _ := strconv.Atoi(r.URL.Path[len("/delay/"):])
	time.Sleep(time.Duration(di) * time.Second)
	output := "Delayed " + r.URL.Path[len("/delay/"):] + " seconds and continued."
	fmt.Println(output)
	fmt.Fprintf(w, "%s", output)
	return
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
