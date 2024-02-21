# README

`yt-kiosk` is a simple tool to display a list of Youtube videos in a kiosk mode.
Every video is displayed in full screen and the next video is played automatically.
Each video is displayed for a fixed amount of time, then the next video is played.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Dependencies

Currently, `yt-kiosk` needs the following tools installed:

- `vlc`
- `playerctl`

## Building

To build the project you need to have go 1.22 installed. Then you can run:

```bash
make build
```

This will create a binary in the `bin` directory.

## Multi-arch builds

To build the project for multiple architectures you can use the following command:

```bash
make all-arch-build
```

This will create a binary for each supported architecture in the `bin` directory.
Supported architectures are: `amd64`, `arm64` and `arm`.

Alternatively, you can use the following command to build for a specific architecture:

```bash
make build-<arch>
```

## Running

To run the project can be run by using one of the binaries created above or you can use the following
command to run it directly from the `Makefile`:

```bash
make run
```

## Usage

To get a list of supported flags and options simply run the binary with the `-h` flag:

```bash
./bin/yt-kiosk -h
NAME:
   yt-kiosk - Run the yt-kiosk

USAGE:
   yt-kiosk [global options] command [command options] 

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --playlist value, -p value  The playlist to play (default: Embedded Playlist)
   --interval value, -i value  The time interval to switch videos (default: 30 seconds)
   --help, -h                  show help
```

### Playlist

The playlist is the list of videos to play. By default, the playlist is set to `Embedded Playlist`.
A custom playlist can be set by using the `--playlist` flag.
The format contains links to youtube videos, each link on a new line.

The following playlist is bundled with the project:

```txt
https://www.youtube.com/watch?v=Iv0FYtEqykw
https://www.youtube.com/watch?v=JWm-ZMXqm7c
https://www.youtube.com/watch?v=rpT8zknUPTI
https://www.youtube.com/watch?v=-CFAcRLZqnM
https://www.youtube.com/watch?v=kHwmzef842g
https://www.youtube.com/watch?v=coVy-NeAd0U
https://www.youtube.com/watch?v=f9ALdmDLPGU
```

## Roadmap

- [ ] Use libvlc to play videos, to get rid of the dependency on `vlc` and `playerctl`
- [ ] Windows support (requires libvlc to be used)