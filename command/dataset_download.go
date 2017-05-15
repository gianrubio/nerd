package command

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/mitchellh/cli"
	"github.com/nerdalize/nerd/nerd/aws"
	v1payload "github.com/nerdalize/nerd/nerd/client/batch/v1/payload"
	v1datatransfer "github.com/nerdalize/nerd/nerd/service/datatransfer/v1"
	"github.com/pkg/errors"
)

const (
	//OutputDirPermissions are the output directory's permissions.
	OutputDirPermissions = 0755
	//DownloadConcurrency is the amount of concurrent download threads.
	DownloadConcurrency = 64
	//DatasetPrefix is the prefix of each dataset ID.
	DatasetPrefix = "d-"
	//TagPrefix is the prefix of a tag identifier
	TagPrefix = "tag-"
)

//DownloadOpts describes command options
type DownloadOpts struct {
	NerdOpts
	AlwaysOverwrite bool `long:"always-overwrite" default-mask:"false" description:"always overwrite files when they already exist"`
}

//Download command
type Download struct {
	*command
}

//DatasetDownloadFactory returns a factory method for the join command
func DatasetDownloadFactory() (cli.Command, error) {
	comm, err := newCommand("nerd dataset download <dataset> <output-dir>", "download data from the cloud to a local directory", "", nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create command")
	}
	cmd := &Download{
		command: comm,
	}
	cmd.runFunc = cmd.DoRun
	return cmd, nil
}

//DoRun is called by run and allows an error to be returned
func (cmd *Download) DoRun(args []string) (err error) {
	if len(args) < 2 {
		return fmt.Errorf("not enough arguments, see --help")
	}

	config, err := cmd.conf.Read()
	if err != nil {
		HandleError(err)
	}

	downloadObject := args[0]
	outputDir := args[1]

	// Folder create and check
	fi, err := os.Stat(outputDir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(outputDir, OutputDirPermissions)
		if err != nil {
			HandleError(errors.Errorf("The provided path '%s' does not exist and could not be created.", outputDir))
		}
		fi, err = os.Stat(outputDir)
	}
	if err != nil {
		HandleError(err)
	} else if !fi.IsDir() {
		HandleError(errors.Errorf("The provided path '%s' is not a directory", outputDir))
	}

	// Clients
	batchclient, err := NewClient(cmd.ui, cmd.conf)
	if err != nil {
		HandleError(err)
	}
	dataOps, err := aws.NewDataClient(
		aws.NewNerdalizeCredentials(batchclient, config.CurrentProject.Name),
		config.CurrentProject.AWSRegion,
	)
	if err != nil {
		HandleError(errors.Wrap(err, "could not create aws dataops client"))
	}

	// Gather dataset IDs
	var datasetIDs []string
	if !strings.HasPrefix(downloadObject, TagPrefix) {
		datasetIDs = append(datasetIDs, downloadObject)
	} else {
		var datasets *v1payload.ListDatasetsOutput
		datasets, err = batchclient.ListDatasets(config.CurrentProject.Name, downloadObject)
		if err != nil {
			HandleError(err)
		}
		datasetIDs = make([]string, len(datasets.Datasets))
		for i, dataset := range datasets.Datasets {
			datasetIDs[i] = dataset.DatasetID
		}
	}

	for _, datasetID := range datasetIDs {
		logrus.Infof("Downloading dataset with ID '%v'", datasetID)
		downloadConf := v1datatransfer.DownloadConfig{
			BatchClient: batchclient,
			DataOps:     dataOps,
			LocalDir:    outputDir,
			ProjectID:   config.CurrentProject.Name,
			DatasetID:   datasetID,
			Concurrency: 64,
		}
		if !cmd.confOpts.JSONOutput { // show progress bar
			progressCh := make(chan int64)
			progressBarDoneCh := make(chan struct{})
			var size int64
			size, err = v1datatransfer.GetRemoteDatasetSize(context.Background(), batchclient, dataOps, config.CurrentProject.Name, datasetID)
			if err != nil {
				HandleError(err)
			}
			go ProgressBar(size, progressCh, progressBarDoneCh)
			downloadConf.ProgressCh = progressCh
			err = v1datatransfer.Download(context.Background(), downloadConf)
			if err != nil {
				HandleError(errors.Wrapf(err, "failed to download dataset '%v'", datasetID))
			}
			<-progressBarDoneCh
		} else { //do not show progress bar
			err = v1datatransfer.Download(context.Background(), downloadConf)
			if err != nil {
				HandleError(errors.Wrapf(err, "failed to download dataset '%v'", datasetID))
			}
		}
	}

	return nil
}
