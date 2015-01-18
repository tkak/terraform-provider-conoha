variable "conoha_username" {}
variable "conoha_password" {}
variable "conoha_tenant" {}

provider "conoha" {
    username = "${var.conoha_username}"
    password = "${var.conoha_password}"
    tenant_name = "${var.conoha_tenant}"
}

resource "conoha_container" "example" {
    name = "test1"
}
