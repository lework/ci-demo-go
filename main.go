package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// 用于接收构建时注入的版本信息
var (
        branch  = "unknown"
        commit  = "unknown"
        appName = "go-app"
)

// Config 应用配置结构
type Config struct {
        Server struct {
                Port string `yaml:"port"`
                Mode string `yaml:"mode"`
                Env  string `yaml:"env"`
        } `yaml:"server"`
        App struct {
                Name string `yaml:"name"`
        } `yaml:"app"`
        Log struct {
                Level  string `yaml:"level"`
                Format string `yaml:"format"`
        } `yaml:"log"`
}

// 定义全局日志对象
var logger *logrus.Logger

// 初始化日志系统
func initLogger(config *Config) {
        logger = logrus.New()

        // 设置日志格式
        if config.Log.Format == "json" {
                logger.SetFormatter(&logrus.JSONFormatter{})
        } else {
                logger.SetFormatter(&logrus.TextFormatter{
                        FullTimestamp: true,
                })
        }

        // 设置日志级别
        switch strings.ToLower(config.Log.Level) {
        case "debug":
                logger.SetLevel(logrus.DebugLevel)
        case "info":
                logger.SetLevel(logrus.InfoLevel)
        case "warn":
                logger.SetLevel(logrus.WarnLevel)
        case "error":
                logger.SetLevel(logrus.ErrorLevel)
        default:
                logger.SetLevel(logrus.InfoLevel)
        }

        // 设置输出
        logger.SetOutput(os.Stdout)

        logger.Info("日志系统初始化完成")
}

// 加载配置文件
func loadConfig(configPath string) (*Config, error) {
        config := &Config{}

        data, err := os.ReadFile(configPath)
        if err != nil {
                return nil, err
        }

        err = yaml.Unmarshal(data, config)
        if err != nil {
                return nil, err
        }

        return config, nil
}

func main() {
        // 定义命令行参数
        configPath := flag.String("f", "etc/app_dev.yaml", "Environment config file")
        flag.Parse()

        // 根据环境选择配置文件
        config, err := loadConfig(*configPath)
        if err != nil {
                log.Fatalf("Error loading config: %v", err)
        }

        // 初始化日志
        initLogger(config)

        // 设置Gin模式
        if config.Server.Mode == "release" {
                gin.SetMode(gin.ReleaseMode)
        } else if config.Server.Mode == "debug" {
                gin.SetMode(gin.DebugMode)
        } else {
                gin.SetMode(gin.ReleaseMode)
        }

        // 创建Gin路由
        r := gin.Default()

        // 版本信息路由
        r.GET("/version", func(c *gin.Context) {
                message := "Version: " + branch + " Build: " + commit
                logger.Info(message)
                c.String(200, message)
        })


        // 健康检查路由
        r.GET("/health", func(c *gin.Context) {
                c.String(200, "ok")
        })


        // 欢迎路由
        r.GET("/:message", func(c *gin.Context) {
                message := c.Param("message")
                responseMsg := "Hello, you got the message: " + message
                logger.WithFields(logrus.Fields{
                        "message": message,
                }).Info("收到请求")
                c.String(200, responseMsg)
        })

        // 使用配置的端口或环境变量
        port := config.Server.Port
        if fromEnv := os.Getenv("PORT"); fromEnv != "" {
                port = fromEnv
        }

        // 启动服务
        logger.WithFields(logrus.Fields{
                "port": port,
                "mode": config.Server.Mode,
                "env":  config.Server.Env,
        }).Info("服务器启动")

        if err := r.Run(":" + port); err != nil {
                logger.Fatal(err)
        }
}