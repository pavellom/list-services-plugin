# Cloud Foundry List App Services plugin
Cloud Foundry plugin enables you to list services which are bound to application

* Easy to monitor applications bound services instead of traditional list all services view
* Service Instance URL for each listed service instance

## Installation
### From CF-Community
```$ cf install-plugin -r CF-Community "ListServices"```

### Manual Installation
Download the binary file for your target OS from the [latest release](https://github.com/pavellom/list-services-plugin/releases/latest).

Install the plugin with `cf install-plugin [list-services-plugin]` (replace `[list-services-plugin]` with the actual binary name you will use, which depends on the OS you are running).

<pre>
Installing plugin ListServices...
OK

Plugin ListServices 1.0.0 successfully installed.
</pre>

## Usage

<pre>
NAME:
   list-services - List services which are bound to tha application

USAGE:
   list-services APP_NAME - Output to Console the list of services which are bound to the application.
</pre>   

## Uninstall

<pre>
$cf uninstall-plugin ListServices
Uninstalling plugin ListServices...
OK

Plugin ListServices 1.0.0 successfully uninstalled.
</pre> 

## License
This project is licensed under the Apache Software License, v. 2 except as noted otherwise in the [LICENSE](https://github.com/pavellom/list-services-plugin/blob/master/LICENSE) file.
