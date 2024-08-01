# GitLab CI/CD Pipeline Configuration

This repository contains the GitLab CI/CD pipeline configuration for automating the deployment and management of Terraform infrastructure. The `.gitlab-ci.yml` file is used to define the CI/CD process, leveraging GitLab CI/CD features to ensure efficient infrastructure provisioning and management.

## Overview

The pipeline is designed to automate the process of managing Terraform infrastructure using GitLab CI/CD. It includes various stages to authenticate, initialize Terraform, and manage infrastructure resources, including tainting and untainting Terraform resources.

### Stages

The pipeline consists of the following stages:

1. **Login**: Authenticates to Vault and retrieves necessary credentials.
2. **Extract**: Initializes Terraform, creates a child pipeline configuration, and prepares resources for subsequent stages.
3. **Global**: Handles Terraform operations such as refresh and destroy.
4. **Trigger**: Manages manual interventions like tainting and untainting resources.

## Variables

The pipeline uses several environment variables to configure Terraform, Vault, and other tools. Below is a brief overview of the key variables:

- **VAULT_ADDR**: The address of the Vault server used for retrieving secrets.
- **VAULT_NAMESPACE**: Specifies the Vault namespace.
- **VAULT_AUTH_NAME**: Authentication method for accessing Vault.
- **VAULT_ROLE**: Role used to authenticate with Vault.
- **TF_ROOT**: Path to the Terraform root directory.
- **TF_ADDRESS**: Address of the Terraform state.
- **ALLOW_REPLACE_ON_APPLY**: Flag indicating if replacements are allowed on Terraform apply.
- **IBM_ACCOUNT_ID**: IBM account ID used for authentication.
- **REALM_ID**: IBM Cloud realm identifier.
- **CI_REGISTRY**: The Docker image registry for pulling the necessary images.
- **CI_API_V4_URL**: GitLab API URL used for configuring Terraform state.

## Cache

The pipeline utilizes caching to speed up repeated Terraform operations by storing Terraform modules and plugins.

```yaml
cache:
  key: "${TF_ROOT}"
  paths:
    - ${TF_ROOT}/.terraform/
