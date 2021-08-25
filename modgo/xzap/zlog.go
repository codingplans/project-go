package xzap

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zLog struct {
	log                    *zap.Logger
	outputPaths            []string
	closeOutputPathsFunc   func()
	errorOutPaths          []string
	closeErrorOutPathsFunc func()
	level                  zapcore.Level
}

var zl *zLog

// 用官方实现
func InitZLog(outputPaths []string, level zapcore.Level) (err error) {
	if zl != nil {
		return
	}
	instNo := os.Getenv("INST_NO")
	if instNo == "" {
		instNo = "1"
	}
	errorOutPaths := []string{}
	for k, v := range outputPaths {
		if v == "stderr" || v == "stdin" || v == "stdout" {
			continue
		}
		ps := strings.Split(v, ".")
		ps[0] = ps[0] + "-" + instNo
		v = strings.Join(ps, ".")

		outputPaths[k] = v
		errorOutPaths = append(errorOutPaths, v+".error")
	}

	zl = &zLog{
		outputPaths:   outputPaths,
		errorOutPaths: errorOutPaths,
		level:         level,
	}
	err = zl.init()
	if err != nil {
		return
	}
	fmt.Println("new zlog", outputPaths, errorOutPaths)
	zl.splitByTime()
	return
}

func (this *zLog) init() (err error) {
	var sink zapcore.WriteSyncer
	// 初始化日志输入文件
	sink, this.closeOutputPathsFunc, err = zap.Open(this.outputPaths...)
	if err != nil {
		return
	}
	allWriter := zapcore.AddSync(sink)
	// 初始化错误日志输入文件
	sink, this.closeErrorOutPathsFunc, err = zap.Open(this.errorOutPaths...)
	if err != nil {
		return
	}
	errorWriter := zapcore.AddSync(sink)
	// 初始化日志格式配置
	config := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     this.timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewJSONEncoder(config)

	// 一次写行为到多个输出端
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, allWriter, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= this.level
		})),
		zapcore.NewCore(encoder, errorWriter, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})),
	)

	this.log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	this.log = this.log.WithOptions(zap.AddCallerSkip(1))
	return
}

func (this *zLog) timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000000"))
}

func (this *zLog) splitByTime() {
	fmt.Println("do splitByTime")

	go func() {
		var lastSplitHour = -1
		for {
			time.Sleep(200 * time.Millisecond)

			// 每分钟写入一个测试日志
			/*if time.Now().Second() == 0 {
				this.Debug("zlog")
			}*/
			// 整点切换文件
			if time.Now().Minute() == 59 {
				currHour := time.Now().Hour()
				if currHour == lastSplitHour {
					continue
				}
				lastSplitHour = currHour

				for _, file := range this.outputPaths {
					_, err := os.Stat(file)
					if err == nil {
						newFile := file + "." + time.Now().Format("2006-01-02_15")
						err = os.Rename(file, newFile)
						if err != nil {
							fmt.Println(err)
						} else {
							fmt.Println("RenameFile", newFile)
						}
					}
				}
				if currHour == 23 {
					for _, file := range this.errorOutPaths {
						_, err := os.Stat(file)
						if err == nil {
							newFile := file + "." + time.Now().Format("2006-01-02_15")
							err = os.Rename(file, newFile)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("RenameFile", newFile)
							}
						}
					}
				}

				this.log.Sync()
				this.closeOutputPathsFunc()
				this.closeErrorOutPathsFunc()
				this.init()
			}
		}
	}()
}

func Debug(msg string, fields ...zap.Field) {
	zl.log.Debug(msg, fields...)
}

func DebugContext(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, getLogField(ctx))
	zl.log.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	zl.log.Info(msg, fields...)
}

// 批量打印用于测试使用
func Infos(cont ...interface{}) {
	arr := []zap.Field{}
	for _, v := range cont {
		arr = append(arr, zap.Any("key:", v))
	}
	Info("打印信息:", arr...)
}

func InfoContext(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, getLogField(ctx))
	zl.log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zl.log.Error(msg, fields...)
}

func ErrorContext(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, getLogField(ctx))
	zl.log.Error(msg, fields...)
}

func getLogField(ctx context.Context) (field zap.Field) {
	log_id := ""
	if ctx.Value("X-Request-Id") != nil {
		log_id = ctx.Value("X-Request-Id").(string)
	}
	field = zap.String("X-Request-Id", log_id)
	return
}
