<div align="center">
  <img src="logo.png"/>
</div>

<br />
<br />

## Description

Generate Swagger Documentation from Insomnia REST Client.
<br />
*This is a fork of [swaggymnia](https://github.com/mlabouardy/swaggymnia) which is now dead.*

## Releases
[![Build Status](https://travis-ci.com/Fyb3roptik/swaggomnia.svg?branch=master)](https://travis-ci.com/Fyb3roptik/swaggomnia) | [Latest Release](https://github.com/Fyb3roptik/swaggomnia/releases/latest)

## How to use it

See usage with:

```
$ swaggomnia --help
```

Generate Swagger documentation:

```
$ swaggomnia generate -insomnia INSOMNIA_EXPORTED_FILE -config CONFIG_FILE -output FORMAT
```

| Option | Description |
| ------ | ----------- |
| -insomnia | Insomnia exported file |
| -config | API Global Configuration file (see [Configuration Format](#configuration-format))|
| -output | Insomnia output format (json or yaml, default json)  |


## Example

Let's convert the following Insomnia API documentation to Swagger:

<div align="center">
  <img src="insomnia.png"/>
</div>

Issue the following command:

```
$ swaggomnia generate -i examples/watchnow.json -c examples/config.json -o json
```

<div align="center">
  <img src="swagger.png"/>
</div>

## Configuration Format

```
{
  "title" : "API Name",
  "version" : "API version",
  "host" : "API URL",
  "basePath" : "Base URL",
  "schemes" : "HTTP protocol",
  "description" : "API description"
}
```

## Maintainers
- Nick Wallace - nwallace@fyberstudios.com - New Author

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
