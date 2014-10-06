variable "conoha_user" {}
variable "conoha_password" {}
variable "conoha_tenant" {}

provider "conoha" {
    user = "${var.conoha_user}"
    password = "${var.conoha_password}"
    tenant = "${var.conoha_tenant}"
}

resource "conoha_container" "example" {
    name = "test1"
}
