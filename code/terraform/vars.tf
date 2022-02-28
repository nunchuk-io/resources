variable "env" {
    default = "dev"
    type = string
    description = "Environment to deploy to"
}

variable "ses_resource_arn" {
    type = string
    description = "SES Resource ARN"
}

variable "hcaptcha_secret_key" {
    type = string
    description = "Recapcha Secret Key"
}
