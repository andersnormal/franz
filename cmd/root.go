package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sync"
	"time"

	"github.com/andersnormal/franz/config"
	"github.com/andersnormal/franz/pool"
	"github.com/andersnormal/franz/runner"
	"github.com/andersnormal/franz/version"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cfg *config.Config
)

var root = &cobra.Command{
	Use:     "franz",
	Version: version.Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		// create context
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// create pool
		pool, err := chromedp.NewPool(chromedp.PoolLog(log.Printf, log.Printf, log.Printf))
		if err != nil {
			log.Fatal(err)
		}

		r := runner.New()

		// // create chrome instance
		// c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// // run task list
		// var str string
		// err = c.Run(ctxt, getUrl("https://digitalocean.com", &str))
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// loop over the URLs
		var wg sync.WaitGroup
		for i, urlstr := range []string{
			"https://katallaxie.me",
			"https://pixelmilk.com",
			"https://katallaxie.me",
		} {
			wg.Add(1)
			go takeScreenshot(ctx, &wg, pool, i, urlstr)
		}

		// wait for to finish
		wg.Wait()

		// shutdown chrome
		err = pool.Shutdown()
		if err != nil {
			log.Fatal(err)
		}

		// wait for chrome to finish
		// err = c.Wait()
		// if err != nil {
		// 	log.Fatal(err)
		// }

		return nil
	},
}

func init() {
	// seed
	rand.Seed(time.Now().UnixNano())

	// init config
	// create config
	cfg = config.New()

	// add flags
	cfg.AddFlags(root)

	// set default formatter
	log.SetFormatter(&log.TextFormatter{})

	// silence on the root cmd
	root.SilenceErrors = true
	root.SilenceUsage = true

	// initialize upon running commands
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// setup logger
	cfg.SetupLogger()
}

func Execute() {
	if err := root.Execute(); err != nil {
		log.Error(err)
	}
}

func takeScreenshot(ctxt context.Context, wg *sync.WaitGroup, pool *chromedp.Pool, id int, urlstr string) {
	defer wg.Done()

	// allocate
	c, err := pool.Allocate(ctxt)
	if err != nil {
		log.Printf("url (%d) `%s` error: %v", id, urlstr, err)
		return
	}
	defer c.Release()

	// run tasks
	var buf []byte
	err = c.Run(ctxt, screenshot(urlstr, &buf))
	if err != nil {
		log.Printf("url (%d) `%s` error: %v", id, urlstr, err)
		return
	}

	// write to disk
	err = ioutil.WriteFile(fmt.Sprintf("%d.png", id), buf, 0644)
	if err != nil {
		log.Printf("url (%d) `%s` error: %v", id, urlstr, err)
		return
	}
}

func screenshot(urlstr string, picbuf *[]byte) chromedp.Action {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Sleep(2 * time.Second),
		chromedp.WaitVisible(`#app`),
		chromedp.ActionFunc(func(ctxt context.Context, h cdp.Executor) error {
			buf, err := page.CaptureScreenshot().Do(ctxt, h)
			if err != nil {
				return err
			}
			*picbuf = buf
			return nil
		}),
	}
}
