// Copyright 2019 by linx, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package logs

import (
	"io"
	"strings"
)

type RotateFileConfig struct {
	Filename          string
	MaxSize           int
	MaxBackups        int
	MaxAge            int
	SeparateLevelFile bool
	Level             Level
	Formatter         Formatter
}

type RotateFileHook struct {
	Config     RotateFileConfig
	logWriter  io.Writer
	logWriters map[Level]io.Writer
}

func NewRotateFileHook(config RotateFileConfig) (Hook, error) {

	hook := RotateFileHook{
		Config: config,
	}
	if config.SeparateLevelFile {
		levels := hook.Levels()
		hook.logWriters = make(map[Level]io.Writer)
		for _, level := range levels {

			hook.logWriters[level] = &RotateFileLogger{
				Filename:   strings.ReplaceAll(config.Filename, "<level>", level.String()),
				MaxSize:    config.MaxSize,
				MaxBackups: config.MaxBackups,
				MaxAge:     config.MaxAge,
				LocalTime:  true,
			}
		}

	} else {
		hook.logWriter = &RotateFileLogger{
			Filename:   config.Filename,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			LocalTime:  true,
		}
	}

	return &hook, nil
}

func (hook *RotateFileHook) Levels() []Level {
	return AllLevels[:hook.Config.Level+1]
}

func (hook *RotateFileHook) Fire(entry *Entry) (err error) {
	b, err := hook.Config.Formatter.Format(entry)
	if err != nil {
		return err
	}
	if hook.logWriter != nil {
		hook.logWriter.Write(b)
	} else {
		if w, ok := hook.logWriters[entry.Level]; ok {
			w.Write(b)
		}
	}
	return nil
}
