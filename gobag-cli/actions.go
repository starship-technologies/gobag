package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/klauspost/compress/gzip"

	"github.com/starship-technologies/gobag/rosbag"
)

func dumpChunkInfo(filename string) error {
	log.Printf("Working with bag file %v.", filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Printf("Unable to open input file, error %s", err)
	}
	defer f.Close()
	rb := rosbag.NewRosBag()
	err = rb.Read(f)
	if err != nil {
		log.Printf("Unable to create ros bag, error %s", err)
		return err
	}
	rb.DumpChunkInfo(filename)
	log.Printf("Done with bag file %v.", filename)
	return nil
}

func dumpChunks(filename string) error {
	log.Printf("Working with bag file %v.", filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Printf("Unable to open input file, error %s", err)
	}
	defer f.Close()
	rb := rosbag.NewRosBag()
	err = rb.Read(f)
	if err != nil {
		log.Printf("Unable to create ros bag, error %s", err)
		return err
	}
	rb.DumpChunks(filename)
	log.Printf("Done with bag file %v.", filename)
	return nil
}

func dumpMessageDefinitions(filename string) error {
	log.Printf("Working with bag file %v.", filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Printf("Unable to open input file, error %s", err)
	}
	defer f.Close()
	rb := rosbag.NewRosBag()
	err = rb.Read(f)
	if err != nil {
		log.Printf("Unable to create ros bag, error %s", err)
		return err
	}
	rb.DumpMessageDefinitions(filename)
	log.Printf("Done with bag file %v.", filename)
	return nil
}

func dumpTableDefinitions(filename string) error {
	log.Printf("Working with bag file %v.", filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Printf("Unable to open input file, error %s", err)
	}
	defer f.Close()
	rb := rosbag.NewRosBag()
	err = rb.Read(f)
	if err != nil {
		log.Printf("Unable to create ros bag, error %s", err)
		return err
	}
	rosbag.DumpTableDefinitions(".")
	log.Printf("Done with bag file %v.", filename)
	return nil
}

func dumpJSON(inputfilename string, outputfilename string) error {
	log.Printf("Working with bag file %v.", inputfilename)
	f, err := os.Open(inputfilename)
	if err != nil {
		log.Printf("Unable to open input file, error %s", err)
	}
	defer f.Close()
	rb := rosbag.NewRosBag()
	err = rb.Read(f)
	if err != nil {
		log.Printf("Unable to create ros bag, error %s", err)
		return err
	}
	of, err := os.Create(outputfilename)
	defer of.Close()
	if err != nil {
		log.Printf("Unable to open output file, err %s", err)
		return err
	}
	gzof, _ := gzip.NewWriterLevel(of, gzip.BestSpeed)
	defer gzof.Close()
	err = rb.WriteJSON(gzof)
	gzof.Flush()
	if err != nil {
		log.Printf("Error on generating or writing JSON file, error %s", err)
		return err
	}
	log.Printf("Done with bag file %v.", inputfilename)
	return nil
}

func dumpTopicsJSON(inputfilename string, outputdirname string, startTime int64, endTime int64, topicsFilter []string) error {
	log.Printf("Working with bag file %v.", inputfilename)
	f, err := os.Open(inputfilename)
	if err != nil {
		log.Printf("Unable to open input file, error %s", err)
	}
	defer f.Close()
	outputPath, err := filepath.Abs(outputdirname)
	if err != nil {
		log.Printf("Error getting output path, error %s", err)
		return err
	}
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		log.Printf("Output path does not exist, error %s", err)
		return err
	}
	rb := rosbag.NewRosBag()
	err = rb.Read(f)
	if err != nil {
		log.Printf("Unable to create ros bag, error %s", err)
		return err
	}
	err = rb.WriteTopicsJSON(outputPath, startTime, endTime, topicsFilter)
	if err != nil {
		log.Printf("Error on generating or writing JSON file, error %s", err)
		return err
	}
	log.Printf("Done with bag file %v.", inputfilename)
	return nil
}
