package common

import (
	"github.com/ljg-cqu/core/esignbox_swagin/template/models/models"
	"github.com/ljg-cqu/core/postgres_job"
)

const (
	TemplateUploadStatusSyncQueue = "template_upload_status_sync_queue"
	TemplateUploadStatusSyncTask  = "template_upload_status_sync_task"
)

var JobQueueClient *postgres_job.Client

var JobQueueWorker *postgres_job.WorkerPool

type TemplateUploadStatusSyncArgs struct {
	TemplateId       string
	StopSyncAtStatus models.TemplateFileStatus
}
