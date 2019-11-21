# Examples for using Terraform Provider
This directory has examples for server configuration in Intersight and OS Installation using Terraform provider.
Each directory has ```.tf``` files with the file name corresponding to the resource that is described in the file.

## Handling secrets
Secrets like API Key Id, passwords and other keys should stored in a file named ```terraform.tfvars```.
Variables that need to be initialized before ```terraform apply``` are described in the variables.tf file.
Steps to define variable values:
* Create a new file with the name ```terraform.tfvars``` parallel to other ```.tf``` files.
* Copy contents of ```terraform.tfvars.example``` into the new file created above.
* Assigns values to these variables. Below is a sample ```terraform.vars``` file

```hcl-terraform
// provider intersight
api_key = "ABCDEFGHIJKLMNOPQRST1234567890"
// ipmi over lay policy props
encryption_key = "123456"
// snmp policy props
auth_password = "cisco1234"
privacy_password = "cisco1234"
```
***NOTE:*** These steps will hide secret from directly being visible in the tf files ONLY. They will still be visible as 
plain text in the state files.  
## Defining Variables
Variables can be defined in ```variables.tf``` file. The already existing variable definitions can be used as a reference.
Values to these variables must be assigned in ```terraform.tfvars``` or any other ```.tfvars``` file.