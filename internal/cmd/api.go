package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/needl3/goreact-template/internal/api"
	"github.com/needl3/goreact-template/internal/cmdutil"
	"github.com/spf13/cobra"
)

func APICmd(ctx context.Context) *cobra.Command {
	apicmd := &cobra.Command{
		Use:   "api",
		Short: "Runs the restful api for goreact app",
		RunE: func(cmd *cobra.Command, args []string) error {
			port := 3000
			if os.Getenv("PORT") != "" {
				port, _ = strconv.Atoi(os.Getenv("PORT"))
			}

			db, err := cmdutil.NewDatabasePool(ctx, 10)
			if err != nil {
				return err
			}
			defer db.Close()

			newApi := api.New(ctx, db)
			srv := newApi.Server(ctx, port)
			go srv.ListenAndServe()

			fmt.Println("Running on http://localhost:" + strconv.Itoa(port))

			<-ctx.Done()
			return nil
		},
	}

	return apicmd
}
