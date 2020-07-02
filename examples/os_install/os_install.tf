resource "intersight_os_install" "os1" {
  name = "InstallTemplatee165"
  server {
    object_type = "compute.RackUnit"
    moid = "5db6c1366176752d30b9bfd7"
  }
  image {
    object_type = "softwarerepository.OperatingSystemFile"
    moid = intersight_softwarerepository_operating_system_file.osf1.moid
  }
  osdu_image {
    moid = intersight_firmware_server_configuration_utility_distributable.scu1.moid
    object_type = "firmware.ServerConfigurationUtilityDistributable"
  }
  answers {
    answer_file = "vmaccepteula\nrootpw Nbv@12345\n\ninstall --overwritevmfs --disk=<DISK> \n#network --bootproto=static --vlanid=1024\nnetwork --bootproto=static  --ip=10.106.233.228 --netmask=255.255.255.128 --gateway=10.106.233.129 --nameserver=64.104.128.236\n\n%pre --interpreter=busybox\nhwclock -d %LIVE_VAR_DATE_1% -t %LIVE_VAR_TIME_UTC_1%\ndate -s %LIVE_VAR_DATE_TIME_UTC_1%\ncd /tmp\n\n%firstboot --interpreter=busybox\ncd /tmp\nesxcfg-vswitch -A 'VM Network' vSwitch0\n\n###############################\n\n# enable & start remote ESXi Shell  (SSH)\n\n###############################\nvim-cmd hostsvc/enable_ssh\nvim-cmd hostsvc/start_ssh\n\n###############################\n# enable & start ESXi Shell (TSM)\n###############################\nvim-cmd hostsvc/enable_esx_shell\nvim-cmd hostsvc/start_esx_shell\nesxcli system hostname set --host=VM_NAME\n#%post --interpreter=busybox --ignorefailure=true\n%post --interpreter=busybox --ignorefailure=true\nESXI_INSTALL_LOG=/var/log/esxi_install.log\necho \\\"OS INSTALL COMPLETED\\\" >>   /var/log/Xinstall.log\n/opt/ucs_tool_esxi/ucs_ipmitool write_file  /var/log/Xinstall.log osProgress.log\n#cd /tmp\\r\\nlocalcli network firewall set --default-action true\nlocalcli network firewall set --enabled false\n\nlocalcli network firewall set --default-action false\nlocalcli network firewall set --enabled true\n\nreboot\n"
    source = "File"
  }
  description = "Install Template 5"
  install_method = "vMedia"
}

/*
SAMPLE PAYLOAD
-----------------
{
  "Answers": {
    "AnswerFile": "vmaccepteula \nrootpw changeMe \n \ninstall --overwritevmfs --disk=<DISK>  \n \n#network --bootproto=static --vlanid=1024 \nnetwork --bootproto=static  --ip=10.105.219.159 --netmask=255.255.255.0 --gateway=10.105.219.1 --nameserver=64.104.128.236 --hostname=ESX65ii \n \n \n \n%pre --interpreter=busybox \nhwclock -d %LIVE_VAR_DATE_1% -t %LIVE_VAR_TIME_UTC_1% \ndate -s %LIVE_VAR_DATE_TIME_UTC_1% \ncd /tmp \n \n%firstboot --interpreter=busybox \ncd /tmp \nesxcfg-vswitch -A 'VM Network' vSwitch0 \n \n \n############################### \n# enable & start remote ESXi Shell  (SSH) \n############################### \nvim-cmd hostsvc/enable_ssh \nvim-cmd hostsvc/start_ssh \n  \n############################### \n# enable & start ESXi Shell (TSM) \n############################### \nvim-cmd hostsvc/enable_esx_shell \nvim-cmd hostsvc/start_esx_shell \n \nesxcli system hostname set --host=VM_NAME \n \n#%post --interpreter=busybox --ignorefailure=true \n \n%post --interpreter=busybox --ignorefailure=true \nESXI_INSTALL_LOG=/var/log/esxi_install.log \necho \"OS INSTALL COMPLETED\" >>   /var/log/Xinstall.log \n/opt/ucs_tool_esxi/ucs_ipmitool write_file  /var/log/Xinstall.log osProgress.log \n \n#cd /tmp \nlocalcli network firewall set --default-action true \nlocalcli network firewall set --enabled false \n \nlocalcli network firewall set --default-action false \nlocalcli network firewall set --enabled true \n \nreboot \n",
    "Source": "File"
  },
  "Description": "Install TRemplate 5",
  "InstallMethod": "vMedia",
  "Name": "InstallTemplatee165",
  "Server": {
    "ObjectType": "compute.RackUnit",
    "Moid": ""
  },
  "Image": {
    "ObjectType": "softwarerepository.OperatingSystemFile",
    "Moid": ""
  },
  "OsduImage": {
    "Moid": "",
    "ObjectType": "firmware.ServerConfigurationUtilityDistributable"
  }
}
*/