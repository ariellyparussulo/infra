## Necessary packages
* azure client
* ansible-playbook setup.yaml --ask-become-pass

ansible-playbook setup-harbor.yml

* change children inside inventory to be the azure machine IP
* create a registry in you route 53
* change the domain name inside inventory list
* run ansible-playbook harbor-playbook.yml