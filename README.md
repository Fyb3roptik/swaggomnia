[![Build Status](https://travis-ci.com/Fyb3roptik/swaggomnia.svg?branch=master)](https://travis-ci.com/Fyb3roptik/swaggomnia)

<div align="center">
  <img src="logo.png" width="50%"/>
</div>


Generate Swagger Documentation from Insomnia REST Client. This is a fork of [swaggymnia](https://github.com/mlabouardy/swaggymnia) which is now dead.

## Download

Below are the available downloads for the latest version of swaggomnia  (1.0.1). Please download the proper package for your operating system and architecture.

### Linux:

```
wget https://jrpublic.s3-us-west-2.amazonaws.com/linux/1.0.1/swaggomnia
```

### Windows:

```
wget https://jrpublic.s3-us-west-2.amazonaws.com/windows/1.0.1/swaggomnia.exe
```

### Mac:

```
wget https://jrpublic.s3-us-west-2.amazonaws.com/mac/1.0.1/swaggomnia
```

### FreeBSD:

```
wget https://jrpublic.s3-us-west-2.amazonaws.com/freebsd/1.0.1/swaggomnia
```

### OpenBSD:

```
wget https://jrpublic.s3-us-west-2.amazonaws.com/openbsd/1.0.1/swaggomnia
```

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
