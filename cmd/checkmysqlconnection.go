package cmd

import (
	"fmt"
	"log"

	"github.com/go-mysql-org/go-mysql/client"
	"github.com/spf13/cobra"
)

var mysqlConnectionCmd = &cobra.Command{
	Use:   "check",
	Short: "check mysql",
	Long:  `This command check connection to mysql database`,
	Run: func(cmd *cobra.Command, args []string) {
		dbUser, _ := cmd.Flags().GetString("dbuser")
		dbPass, _ := cmd.Flags().GetString("dbpass")
		dbIp, _ := cmd.Flags().GetString("dbip")
		dbPort, _ := cmd.Flags().GetString("dbport")
		dbCommand, _ := cmd.Flags().GetString("dbcommand")
		checkConnection(dbUser,dbPass,dbIp,dbPort)
		if dbCommand != "" {
			runCommand(dbUser,dbPass,dbIp,dbPort,dbCommand)
		}
	},
}

func init()  {
	rootCmd.AddCommand(mysqlConnectionCmd)
	rootCmd.PersistentFlags().String("dbuser", "", "dbuser.")
	rootCmd.PersistentFlags().String("dbpass", "", "dbpass")
	rootCmd.PersistentFlags().String("dbip", "", "dbip")
	rootCmd.PersistentFlags().String("dbport", "", "dbport")
	rootCmd.PersistentFlags().String("dbcommand", "", "dbcommand")
}

func checkConnection(dbuser string, dbpass string, dbip string, dbport string) {
	addr := dbip + ":" + dbport

	fmt.Printf("Connection info: %s" , addr)
	fmt.Println("")
	fmt.Printf("dbusername: %s" , dbuser)
	fmt.Println("")
	fmt.Printf("dbpassword: %s" , dbpass)
	fmt.Println("")
	
	conn, _ := client.Connect( addr, dbuser, dbpass, "danymaster" )
	err := conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func runCommand(dbuser string, dbpass string, dbip string, dbport string, dbcommand string) {
	addr := dbip + ":" + dbport

	conn, _ := client.Connect( addr, dbuser, dbpass, "danymaster" )
	err, _ := conn.Execute(dbcommand)
	if err != nil {
		log.Fatal(err)
	}
}
