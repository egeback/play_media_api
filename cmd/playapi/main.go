package main

import (
	"fmt"
	"sort"
	"time"

	_ "github.com/egeback/playapi/internal/docs"
	"github.com/egeback/playapi/internal/models"
	"github.com/egeback/playapi/internal/parsers"
	"github.com/egeback/playapi/internal/parsers/svtplay"
	"github.com/egeback/playapi/internal/parsers/tv4play"
	"github.com/egeback/playapi/internal/version"
	"github.com/egeback/playapi/pkg/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	jsonp "github.com/tomwei7/gin-jsonp"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Play API
// @version 1.0.5
// @description API including SVT and TV4 Play

// @contact.name API Support
// @contact.url https://github.com/egeback/playapi/issues
// @contact.email support@egeback.se

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @BasePath /api/v1/
func main() {
	fmt.Printf("%s Running Play Media API version: %s (%s)\n", time.Now().Format("2006-01-02 15:04:05"), version.BuildVersion, version.BuildTime)
	parsers.Set([]parsers.ParserInterface{new(svtplay.Parser), new(tv4play.Parser)})
	//parsers.Set([]parsers.ParserInterface{new(tv4play.Parser)})

	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(jsonp.JsonP())
	c := controllers.NewController()
	v1 := r.Group("/api/v1")
	{
		shows := v1.Group("/shows")
		{
			shows.GET("/", c.ListShows)
			shows.GET("/:slug", c.ShowShow)
			shows.GET("/:slug/seasons", c.ShowShow)
		}
		common := v1.Group("/")
		{
			common.GET("/ping", c.Ping)
			common.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}

	go func() {
		gocron.Every(15).Minutes().Do(updateShows)
		<-gocron.Start()
	}()
	updateShows()

	r.Run(":8080")
}

func updateShows() {
	t1 := time.Now()
	shows := make([]models.Show, 0)
	for _, parser := range parsers.All("") {
		s := parser.GetShowsWithSeasons()
		shows = append(shows, parser.PostCheckShows(s)...)
	}

	showsWithSeasons := 0
	showsWithNoSeasons := 0
	notProcessed := 0

	for _, show := range shows {
		if show.Name == "" {
			fmt.Println(show)
		}
		if len(show.Seasons) == 0 {
			showsWithNoSeasons++
		} else {
			showsWithSeasons++
		}
		if !show.Prossesed {
			notProcessed++
		}
	}
	sort.SliceStable(shows, func(i, j int) bool {
		return shows[i].Slug < shows[j].Slug
	})
	models.ShowsSet(shows)
	diff := time.Now().Sub(t1).Seconds()
	fmt.Printf("%s [shows with-shows]/[total]: %d/%d, this took: %fs\n", time.Now().Format("2006-01-02 15:04:05"), showsWithSeasons, len(shows), diff)
	fmt.Printf("%s Shows not processed: %d\n", time.Now().Format("2006-01-02 15:04:05"), notProcessed)
}