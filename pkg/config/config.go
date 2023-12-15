package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/TrevorEdris/retropie-utils/pkg/storage"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		// TODO: Re-evaluate a good UX for this
		// CustomLocations []CustomLocation
		Emulators []Emulator
		Storage   Storage
	}

	// CustomLocation struct {
	// 	Name         string
	// 	AbsolutePath string
	// 	Extensions   []string
	// }

	Emulator struct {
		Name                 string
		InGameSaves          bool
		SaveStates           bool
		Roms                 bool
		AdditionalExtensions []string
	}

	Storage struct {
		GoogleDrive storage.GDriveConfig
		S3          storage.S3Config
		SFTP        storage.SFTPConfig
	}
)

var example = Config{
	Emulators: []Emulator{
		{
			Name:        "gb",
			InGameSaves: true,
			SaveStates:  false,
			Roms:        false,
		},
		{
			Name:        "gba",
			InGameSaves: true,
			SaveStates:  false,
			Roms:        false,
			AdditionalExtensions: []string{
				"mycustomext", "backupthisextensiontoo",
			},
		},
		{
			Name:        "gbc",
			InGameSaves: true,
			SaveStates:  false,
			Roms:        false,
		},
		{
			Name:        "n64",
			InGameSaves: true,
			SaveStates:  false,
			Roms:        false,
		},
		{
			Name:        "nes",
			InGameSaves: true,
			SaveStates:  false,
			Roms:        false,
		},
		{
			Name:        "snes",
			InGameSaves: true,
			SaveStates:  false,
			Roms:        false,
		},
	},
	Storage: Storage{
		GoogleDrive: storage.GDriveConfig{
			Enabled: false,
		},
		S3: storage.S3Config{
			Enabled: true,
			Bucket:  "retropie-sync",
		},
		SFTP: storage.SFTPConfig{
			Enabled: false,
		},
	},
	// CustomLocations: []CustomLocation{
	// 	{
	// 		Name:         "boot config",
	// 		AbsolutePath: "/boot/config.txt",
	// 	},
	// },
}

var validate *validator.Validate

func CreateExample(outputDir string) error {
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return err
	}
	filename := filepath.Join(outputDir, "config.example.yaml")
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	yamlData, err := yaml.Marshal(&example)
	if err != nil {
		return err
	}
	_, err = f.Write(yamlData)
	if err != nil {
		return err
	}
	fmt.Printf("Created %s\n", filename)
	return nil
}

func ValidateConfig(configFile string) error {
	validate = validator.New()

	bytes, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	config := &Config{}
	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		return err
	}

	err = validate.Struct(config)
	if err != nil {
		return err
	}
	return nil
}
