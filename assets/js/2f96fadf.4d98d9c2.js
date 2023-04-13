"use strict";(self.webpackChunkdocumentation=self.webpackChunkdocumentation||[]).push([[3160],{41837:e=>{e.exports=JSON.parse('{"url":"api/swagger.json","themeId":"theme-redoc","isSpecFile":false,"spec":{"openapi":"3.0.0","info":{"title":"The ipsw API.","version":"v1.0"},"paths":{"/_ping":{"get":{"description":"This will return \\"OK\\" if the daemon is running.","tags":["Daemon"],"summary":"Ping","operationId":"getDaemonPing","responses":{"default":{"description":"Default response"}}},"head":{"description":"This will return if 200 the daemon is running.","tags":["Daemon"],"summary":"Ping","operationId":"headDaemonPing","responses":{"default":{"description":"Default response"}}}},"/device_list":{"get":{"description":"This will return JSON of all XCode devices.","tags":["DeviceList"],"summary":"List XCode Devices.","operationId":"getDeviceList","responses":{"default":{"description":"Default response"}}}},"/download/ipsw/ios/latest/build":{"get":{"description":"Get latest iOS build.","tags":["Download"],"summary":"Latest iOS Build","operationId":"getDownloadLatestIPSWsBuild","responses":{"default":{"description":"Default response"}}}},"/download/ipsw/ios/latest/version":{"get":{"description":"Get latest iOS version.","tags":["Download"],"summary":"Latest iOS Version","operationId":"getDownloadLatestIPSWsVersion","responses":{"default":{"description":"Default response"}}}},"/dsc/imports":{"get":{"description":"Get list of dylibs that import a given dylib.","tags":["DSC"],"summary":"Imports","operationId":"getDscImports","parameters":[{"description":"path to IPSW","name":"path","in":"query","required":true,"schema":{"type":"string"}},{"description":"dylib to search for","name":"dylib","in":"query","required":true,"schema":{"type":"string"}}],"responses":{"default":{"description":"Default response"}}}},"/dsc/info":{"get":{"description":"Get info about a given DSC","tags":["DSC"],"summary":"Info","operationId":"getDscInfo","parameters":[{"description":"path to IPSW","name":"path","in":"query","required":true,"schema":{"type":"string"}}],"responses":{"default":{"description":"Default response"}}}},"/dsc/macho":{"get":{"description":"Get MachO info for a given dylib in the DSC.","tags":["DSC"],"summary":"MachO","operationId":"getDscMacho","parameters":[{"description":"path to IPSW","name":"path","in":"query","required":true,"schema":{"type":"string"}},{"description":"dylib to search for","name":"dylib","in":"query","required":true,"schema":{"type":"string"}}],"responses":{"default":{"description":"Default response"}}}},"/dsc/str":{"get":{"description":"Get strings in the DSC that match a given pattern.","tags":["DSC"],"summary":"Strings","operationId":"getDscStrings","parameters":[{"description":"path to IPSW","name":"path","in":"query","required":true,"schema":{"type":"string"}},{"description":"regex to search for","name":"pattern","in":"query","required":true,"schema":{"type":"string"}}],"responses":{"default":{"description":"Default response"}}}},"/dsc/symaddr":{"get":{"description":"Get symbols addresses in the DSC that match a given lookup JSON payload.","tags":["DSC"],"summary":"Symbols","operationId":"getDscSymbols","parameters":[{"description":"path to IPSW","name":"path","in":"query","required":true,"schema":{"type":"string"}}],"responses":{"default":{"description":"Default response"}}}},"/extract/dmg":{"get":{"description":"Extract DMGs from an IPSW.","tags":["Extract"],"summary":"DMG","operationId":"getExtractDmg","responses":{"default":{"description":"Default response"}}}},"/extract/dsc":{"get":{"description":"Extract dyld_shared_caches from an IPSW.","tags":["Extract"],"summary":"DSC","operationId":"getExtractDsc","responses":{"default":{"description":"Default response"}}}},"/extract/kbag":{"get":{"description":"Extract KBAGs from an IPSW.","tags":["Extract"],"summary":"KBAG","operationId":"getExtractKbags","responses":{"default":{"description":"Default response"}}}},"/extract/kernel":{"get":{"description":"Extract kernelcaches from an IPSW.","tags":["Extract"],"summary":"Kernel","operationId":"getExtractKernel","responses":{"default":{"description":"Default response"}}}},"/extract/pattern":{"get":{"description":"Extract files from an IPSW that match a given pattern.","tags":["Extract"],"summary":"Pattern","operationId":"getExtractPattern","responses":{"default":{"description":"Default response"}}}},"/idev/info":{"get":{"description":"Get info about USB connected devices.","tags":["USB"],"summary":"Info","operationId":"getIdevInfo","responses":{"default":{"description":"Default response"}}}},"/info/ipsw":{"get":{"description":"Get IPSW info.","tags":["Info"],"summary":"IPSW","operationId":"getIpswInfo","parameters":[{"description":"path to IPSW","name":"path","in":"query","required":true,"schema":{"type":"string"}}],"responses":{"default":{"description":"Default response"}}}},"/info/ipsw/remote":{"get":{"description":"Get remote IPSW info.","tags":["Info"],"summary":"Remote IPSW","operationId":"getRemoteIpswInfo","parameters":[{"description":"url to IPSW","name":"url","in":"query","required":true,"schema":{"type":"string"}},{"description":"http proxy to use","name":"proxy","in":"query","schema":{"type":"string"}},{"description":"ignore TLS errors","name":"insecure","in":"query","schema":{"type":"boolean"}}],"responses":{"default":{"description":"Default response"}}}},"/info/ota":{"get":{"description":"Get OTA info.","tags":["Info"],"summary":"OTA","operationId":"getOtaInfo","parameters":[{"description":"path to OTA","name":"path","in":"query","required":true,"schema":{"type":"string"}}],"responses":{"default":{"description":"Default response"}}}},"/info/ota/remote":{"get":{"description":"Get remote OTA info.","tags":["Info"],"summary":"Remote OTA","operationId":"getRemoteOtaInfo","parameters":[{"description":"url to OTA","name":"url","in":"query","required":true,"schema":{"type":"string"}},{"description":"http proxy to use","name":"proxy","in":"query","schema":{"type":"string"}},{"description":"ignore TLS errors","name":"insecure","in":"query","schema":{"type":"boolean"}}],"responses":{"default":{"description":"Default response"}}}},"/ipsw/fs/ents":{"get":{"description":"Get IPSW Filesystem DMG MachO entitlements.","tags":["IPSW"],"summary":"Entitlements","operationId":"getIpswFsEntitlements","parameters":[{"description":"path to IPSW","name":"path","in":"query","required":true,"schema":{"type":"string"}}],"responses":{"default":{"description":"Default response"}}}},"/ipsw/fs/files":{"get":{"description":"Get IPSW Filesystem DMG file listing.","tags":["IPSW"],"summary":"Files","operationId":"getIpswFsFiles","parameters":[{"description":"path to IPSW","name":"path","in":"query","required":true,"schema":{"type":"string"}}],"responses":{"default":{"description":"Default response"}}}},"/kernel/kexts":{"get":{"description":"Get kernelcache KEXTs info.","tags":["Kernel"],"summary":"Kexts","operationId":"getKernelKexts","parameters":[{"description":"path to kernelcache","name":"path","in":"query","required":true,"schema":{"type":"string"}}],"responses":{"default":{"description":"Default response"}}}},"/kernel/syscall":{"get":{"description":"Get kernelcache syscalls info.","tags":["Kernel"],"summary":"Syscalls","operationId":"getKernelSyscalls","parameters":[{"description":"path to kernelcache","name":"path","in":"query","required":true,"schema":{"type":"string"}}],"responses":{"default":{"description":"Default response"}}}},"/kernel/version":{"get":{"description":"Get kernelcache version.","tags":["Kernel"],"summary":"Version","operationId":"getKernelVersion","parameters":[{"description":"path to kernelcache","name":"path","in":"query","required":true,"schema":{"type":"string"}}],"responses":{"default":{"description":"Default response"}}}},"/macho/info":{"get":{"description":"Get MachO info.","tags":["MachO"],"summary":"Info","operationId":"getMachoInfo","parameters":[{"description":"path to MachO","name":"path","in":"query","required":true,"schema":{"type":"string"}},{"description":"architecture to get info for in universal MachO","name":"arch","in":"query","schema":{"type":"string"}}],"responses":{"default":{"description":"Default response"}}}},"/mount/{type}":{"post":{"description":"Mount a DMG inside a given IPSW.","tags":["Mount"],"summary":"Mount","operationId":"postMount","parameters":[{"description":"type of DMG to mount","name":"type","in":"path","required":true,"schema":{"type":"string"}},{"description":"path to IPSW","name":"path","in":"query","required":true,"schema":{"type":"string"}}],"responses":{"default":{"description":"Default response"}}}},"/unmount":{"post":{"description":"Unmount a previously mounted DMG.","tags":["Mount"],"summary":"Unmount","operationId":"postUnmount","parameters":[{"description":"mount point of DMG","name":"mount_point","in":"path","required":true,"schema":{"type":"string"}},{"description":"path to DMG","name":"dmg_path","in":"query","schema":{"type":"string"}}],"responses":{"default":{"description":"Default response"}}}},"/version":{"get":{"description":"This will return the daemon version info.","tags":["Daemon"],"summary":"Version","operationId":"getDaemonVersion","responses":{"default":{"description":"Default response"}}}}},"servers":[{"url":"http://localhost:8080/v1"}]}}')}}]);