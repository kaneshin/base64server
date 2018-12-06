resource "google_cloudbuild_trigger" "build_trigger" {
  trigger_template {
    branch_name = "master"
    repo_name   = "github_kaneshin_base64server"
  }
  build {
    images = ["gcr.io/$PROJECT_ID/base64server"]
    step {
      name = "gcr.io/cloud-builders/docker"
      args = "build -t gcr.io/$PROJECT_ID/base64server ."
    }
  }
}

data "terraform_remote_state" "default" {
  backend = "gcs"
  config {
    bucket  = "${var.tfstate_bucket}"
    prefix  = "${var.env}/base64server"
  }
}

provider "google" {
  version     = "~> 1.19"
  project     = "${var.project_id}"
  region      = "${var.region}"
  credentials = "${file("${var.credentials_file_path}")}"
}

variable "region" {
  default = "asia-northeast1"
}

variable "region_zone" {
  default = "asia-northeast1-b"
}

variable "project_id" {}

variable "credentials_file_path" {}

variable "tfstate_bucket" {
  default = "terraform-state"
}

variable "env" {}
