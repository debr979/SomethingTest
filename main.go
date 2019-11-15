package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/robbert229/jwt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
)

type Data struct {
	AccountID string `form:"account_id" json:"account_id" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

const (
	secretKey = `DougService`
)

func init() {

}
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Origin"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"X-AUTH-TOKEN"},
	}))
	r.GET(`/ip`, func(c *gin.Context) {
		ip := net.ParseIP(c.ClientIP())
		IPScan(ip)
		log.Print(ip)
	})

	r.POST(`/test/`, func(c *gin.Context) {
		names := c.QueryMap(`name`)
		fmt.Printf("%v \n",names["one"])
	})

	r.POST(`/login`, func(c *gin.Context) {
		var data Data
		if err := c.ShouldBind(&data); err != nil {
			log.Print(err)
		}
		if data.AccountID == `debr979` && data.Password == `sd958969` {
			jwToken := JwtGenerator(data.AccountID)
			c.Header(`X-AUTH-TOKEN`, jwToken)
			c.JSON(200, gin.H{"result": "success"})
		}
	})
	r.GET(`/getUser`, func(c *gin.Context) {
		userName := JwtValidate(c.GetHeader(`X-AUTH-TOKEN`))
		c.JSON(200, gin.H{"username": userName})
	})
	r.POST(`/char`, func(c *gin.Context) {
		var data Data
		if err := c.ShouldBind(&data); err != nil {
			log.Print(err)
		}
		if data.AccountID != "" && data.Password != "" {
			if CharacterCheck(data.AccountID) && CharacterCheck(data.Password) {
				c.String(200, "OK")
			}
		}
	})
	r.POST(`/Pointer`, func(c *gin.Context) {
		var data Data
		if err := c.ShouldBind(&data); err != nil {
			log.Print(err)
		}
		acc := []byte(data.AccountID)
		pwd := []byte(data.Password)
		fmt.Print(pwd)
		fmt.Printf("%#v\n", acc)
		c.String(200, string(acc))
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"result": "Not_Found"})
	})
	if err := r.Run(":8081"); err != nil {
		log.Print(err)
	}

}

func IPScan(ip net.IP) {
	ipDB, err := geoip2.Open(`tools/GeoLite2-Country.mmdb`)
	if err != nil {
		log.Print(err)
	}

	defer func() {
		if err := ipDB.Close(); err != nil {
			log.Print(err)
		}
	}()

	record, err := ipDB.Country(ip)
	if err != nil {
		log.Print(err)
	}

	log.Print(record.Country.Names["en"])
}

func CharacterCheck(wannaCheck string) bool {
	alpha := "abcdefghijklmnopqrstuvwxyz!@#$%,."
	for _, char := range wannaCheck {
		if !strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}

func JwtGenerator(userName string) string {
	algorithm := jwt.HmacSha256(StringTo256(secretKey))
	log.Print(StringTo256(secretKey))
	claims := jwt.NewClaim()
	claims.Set(`username`, userName)
	claims.SetTime(`exp`, time.Now().Add(time.Hour))
	token, err := algorithm.Encode(claims)
	if err != nil {
		log.Print(err)
	}
	return token
}

func JwtValidate(jwToken string) (userName string) {
	algorithm := jwt.HmacSha256(StringTo256(secretKey))
	loadedClaims, err := algorithm.DecodeAndValidate(jwToken)
	if err != nil {
		userName = `Token has expired`
	} else {
		role, err := loadedClaims.Get(`username`)
		if err != nil {
			log.Print(err)
		} else {
			roleString, ok := role.(string)
			if !ok {
				log.Print(err)
			}
			userName = roleString
		}
	}
	log.Printf("username:%s", userName)
	return userName
}

func StringTo256(secret string) string {
	hash := sha256.Sum256([]byte(secret))
	return hex.EncodeToString(hash[:])
}
