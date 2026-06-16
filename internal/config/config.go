package config

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Version   string
	Project   Project
	Migration Migration
	Logs      Logs
	AI        AI
}

type Project struct {
	Source string
	Module string
}

type Migration struct {
	Out    string
	Report string
	DryRun bool
}

type Logs struct {
	File string
}

type AI struct {
	Provider string
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	return Parse(data)
}

func Parse(data []byte) (Config, error) {
	var cfg Config
	scanner := bufio.NewScanner(bytes.NewReader(data))
	section := ""
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		raw := scanner.Text()
		withoutComment := stripComment(raw)
		if strings.TrimSpace(withoutComment) == "" {
			continue
		}
		indent := len(withoutComment) - len(strings.TrimLeft(withoutComment, " "))
		trimmed := strings.TrimSpace(withoutComment)
		if indent == 0 && strings.HasSuffix(trimmed, ":") {
			section = strings.TrimSuffix(trimmed, ":")
			if !knownSection(section) {
				return Config{}, fmt.Errorf("config line %d: unknown section %q", lineNumber, section)
			}
			continue
		}
		key, value, ok := strings.Cut(trimmed, ":")
		if !ok {
			return Config{}, fmt.Errorf("config line %d: expected key: value", lineNumber)
		}
		if err := assign(&cfg, section, strings.TrimSpace(key), cleanValue(value), lineNumber); err != nil {
			return Config{}, err
		}
	}
	if err := scanner.Err(); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func knownSection(section string) bool {
	switch section {
	case "project", "migration", "logs", "ai":
		return true
	default:
		return false
	}
}

func assign(cfg *Config, section string, key string, value string, lineNumber int) error {
	switch section {
	case "":
		if key != "version" {
			return fmt.Errorf("config line %d: unknown root key %q", lineNumber, key)
		}
		cfg.Version = value
	case "project":
		switch key {
		case "source":
			cfg.Project.Source = value
		case "module":
			cfg.Project.Module = value
		default:
			return fmt.Errorf("config line %d: unknown project key %q", lineNumber, key)
		}
	case "migration":
		switch key {
		case "out":
			cfg.Migration.Out = value
		case "report":
			cfg.Migration.Report = value
		case "dryRun":
			parsed, err := strconv.ParseBool(value)
			if err != nil {
				return fmt.Errorf("config line %d: dryRun must be true or false", lineNumber)
			}
			cfg.Migration.DryRun = parsed
		default:
			return fmt.Errorf("config line %d: unknown migration key %q", lineNumber, key)
		}
	case "logs":
		if key != "file" {
			return fmt.Errorf("config line %d: unknown logs key %q", lineNumber, key)
		}
		cfg.Logs.File = value
	case "ai":
		if key != "provider" {
			return fmt.Errorf("config line %d: unknown ai key %q", lineNumber, key)
		}
		cfg.AI.Provider = value
	default:
		return fmt.Errorf("config line %d: key %q is outside a supported section", lineNumber, key)
	}
	return nil
}

func stripComment(line string) string {
	if index := strings.Index(line, "#"); index >= 0 {
		return line[:index]
	}
	return line
}

func cleanValue(value string) string {
	value = strings.TrimSpace(value)
	value = strings.Trim(value, "\"'")
	return value
}
