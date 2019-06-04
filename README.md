<div align="center">
  <img src="logo.png" width="50%"/>
</div>

Generate Swagger Documentation from Insomnia REST Client. This is a fork of swaggymnia. That repo is now dead.

## Download

Below are the available downloads for the latest version of Swaggomnia  (1.0.0-beta). Please download the proper package for your operating system and architecture.

### Linux:

```
wget https://s3.amazonaws.com/swaggymnia/1.0.0-beta/linux/swaggymnia
```

### Windows:

```
wget https://s3.amazonaws.com/swaggymnia/1.0.0-beta/windows/swaggymnia
```

### Mac OS X:

```
wget https://s3.amazonaws.com/swaggymnia/1.0.0-beta/osx/swaggymnia
```

### FreeBSD:

```
wget https://s3.amazonaws.com/swaggymnia/1.0.0-beta/freebsd/swaggymnia
```

### OpenBSD:

```
wget https://s3.amazonaws.com/swaggymnia/1.0.0-beta/openbsd/swaggymnia
```

## How to use it

See usage with:

```
$ swaggymnia --help
```

Generate Swagger documentation:

```
$ swaggymnia generate -insomnia INSOMNIA_EXPORTED_FILE -config CONFIG_FILE -output FORMAT
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
$ swaggymnia generate -i examples/watchnow.json -c examples/config.json -o json
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

- Mohamed Labouardy - mohamed@labouardy.com - Original Author
- Nick Wallace - nwallace@fyberstudios.com - New Author

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
