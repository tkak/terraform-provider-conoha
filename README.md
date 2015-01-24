terraform-provider-conoha
=========================

[![Build Status](https://travis-ci.org/tkak/terraform-provider-conoha.svg?branch=master)](https://travis-ci.org/tkak/terraform-provider-conoha)


This project is a [terraform](http://www.terraform.io/) provider for [ConoHa](https://www.conoha.jp/en) Object Storage.

This current version only supports a container management function.

## Usage

Terraform file

```
variable "conoha_username" {}
variable "conoha_password" {}
variable "conoha_tenant" {}

provider "conoha" {
    username = "${var.conoha_username}"
    password = "${var.conoha_password}"
    tenant_name = "${var.conoha_tenant}"
}

resource "conoha_container" "example" {
    name = "foo"
}
```

Run terraform.

```
$ terraform apply \
-var "conoha_username=${CONOHA_USERNAME}" \
-var "conoha_password=${CONOHA_PASSWORD}" \
-var "conoha_tenant=${CONOHA_TENANT}"
```

## References

* [gophercloud](http://gophercloud.io/)
* [API Reference](https://www.conoha.jp/guide/guide.php?g=52)
* [Writing Custom Terraform Providers](https://www.hashicorp.com/blog/terraform-custom-providers.html)
