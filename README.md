# ups-agent

Monitors an uninterruptible power supply (UPS) via SNMP and automatically shuts
down the system in the event of a power failure.

The program:

- periodically polls the UPS SNMP card (polling);
- receives SNMP traps from the UPS;
- writes an event log with severity levels;
- initiates a graceful shutdown of the host (and, if necessary, the UPS itself)
  on low battery, when the remaining runtime expires, or on critical traps;
- can display the current UPS status and manage the SNMP card configuration.

## Features

- Retrieving UPS state over SNMP (model, output status, input voltage/frequency,
  output voltage/frequency, load, battery capacity and temperature, remaining
  runtime, firmware, etc.) — the `status` command.
- Reading, writing, and resetting to factory defaults the SNMP card
  configuration — the `cardconfig` command.
- Receiving SNMP traps with source address and community validation.
- Automatic system shutdown based on several independent conditions:
  - low battery (according to the UPS data);
  - battery capacity less than or equal to `powerctl.batterylevel`;
  - remaining runtime less than or equal to `powerctl.minutes`;
  - runtime on battery greater than or equal to `powerctl.timeout`;
  - critical traps requiring an immediate shutdown.
- Optional shutdown of the UPS itself before powering off the host
  (`powerctl.upspoweroff`).

## Requirements

- Linux (uses `syscall.Sync` and `seteuid(0)` when shutting down the host);
- Go 1.27.1 or newer;
- A UPS with an SNMP card (for example, one supporting the `Upsmate.mib` MIB).

## Building

```sh
go build -o ups-agent .
```

The resulting binary and the configuration file are placed in the same
directory, for example `/opt/ups-agent/`.

## Configuration

The program uses the `ups-agent.yml` configuration file, which is looked up next
to the executable by default. Its path can be changed with the `-config` flag.

An example configuration is provided in `ups-agent.yml`. Main sections:

| Section | Parameter | Description |
| ------ | -------- | -------- |
| — | `upsname` | UPS name (optional). |
| — | `model` | UPS model (optional). |
| `logger` | `level` | Log level: `emerg`, `alert`, `crit`, `error`, `warn`, `notice`, `info`, `debug`. |
| `logger` | `logdir` | Log directory (default `logs`). |
| `logger` | `filename` | Log file name (default `ups-agent.log`). |
| `powerctl` | `batterylevel` | Battery capacity threshold in percent for shutdown (default `10`). |
| `powerctl` | `minutes` | Remaining runtime threshold for shutdown (default `10m`). |
| `powerctl` | `timeout` | Maximum runtime on battery (default `0s` — disabled). |
| `powerctl` | `poweroff` | Whether to shut down the host (default `true`). |
| `powerctl` | `poweroffcmd` | Shutdown command (default `/sbin/poweroff`). |
| `powerctl` | `upspoweroff` | Whether to power off the UPS on shutdown (default `false`). |
| `snmp` | `addr` | IP address of the UPS SNMP card (required). |
| `snmp` | `community` | SNMP community for reading (default `public`). |
| `snmp` | `polltime` | Polling interval (default `60s`). |
| `snmp` | `trapcommunity` | SNMP community for receiving traps (default `public`). |
| `snmp` | `traplistener` | Address for receiving traps, e.g. `0.0.0.0:162`. |
| `cardconfig` | … | SNMP card configuration parameters (see `ups-agent.yml`). |

Durations are specified in Go format (`30s`, `10m`, `2h`, etc.).

### Checking the configuration

```sh
./ups-agent -configtest
./ups-agent -configdump   # print the resulting configuration
```

## Usage

### Running the daemon

```sh
./ups-agent
```

The daemon handles the following signals:

- `SIGHUP` — reopen the log file (for rotation);
- `SIGINT`, `SIGTERM`, `SIGQUIT` — graceful shutdown.

### UPS status

```sh
./ups-agent -status
```

Prints the current UPS state: model, output status, voltages, frequency, load,
battery state and capacity, remaining runtime, etc.

### Managing the SNMP card configuration

```sh
./ups-agent -cardconfig get      # read the card settings
./ups-agent -cardconfig set      # write the settings from the cardconfig section
./ups-agent -cardconfig default  # reset the card to factory defaults
```

Values to write are set in the `cardconfig` section of the configuration file;
the allowed values for each parameter are listed in the `ups-agent.yml`
comments.

## Shutdown logic

When any of the following conditions is detected, the agent initiates a system
shutdown:

- the UPS battery status is "low battery";
- battery capacity `<= powerctl.batterylevel`;
- remaining runtime `<= powerctl.minutes`;
- runtime on battery `>= powerctl.timeout` (if `timeout > 0`);
- a critical trap flagged with the `shutdown` flag arrives (for example,
  `upsLowBatteryShutdown`, `upsEmergencyStop`, `upsOverLoadShutdown`).

If `powerctl.upspoweroff` is set, before shutting down the host the agent sends
the UPS a power-off command, then waits 5 seconds for the SNMP packet to be
delivered. The host shutdown is performed only once (`sync.Once` protection).

## Installing as a systemd service

The `contrib` directory contains helper files:

- `contrib/ups-agent.service` — systemd unit file;
- `contrib/ups-agent.logrotate` — log rotation configuration;
- `contrib/prepare-traps.pl` — a utility for generating a trap map from the
  `card/Upsmate.mib` MIB.

Installing the service:

```sh
# place the binary and the configuration
sudo mkdir -p /opt/ups-agent
sudo cp ups-agent ups-agent.yml /opt/ups-agent/

# create the user
sudo useradd --system --no-create-home --shell /usr/sbin/nologin ups-agent

# install the unit file and start the service
sudo cp contrib/ups-agent.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now ups-agent
```

For proper host shutdown, the process needs the `CAP_SETUID` capability (for
`seteuid(0)`) and `CAP_NET_BIND_SERVICE` (for receiving traps on port 162),
which is already handled in the unit file via `AmbientCapabilities`.

If trap receiving is configured on port `162` and the service runs under an
unprivileged user, the `CAP_NET_BIND_SERVICE` capability allows listening on
that port.

To prevent systemd from periodically stopping the daemon for a user without an
active session, run:

```sh
loginctl enable-linger ups-agent
```

### Log rotation

Install the rotation configuration and adjust the path if necessary:

```sh
sudo cp contrib/ups-agent.logrotate /etc/logrotate.d/ups-agent
```

Rotation sends the `HUP` signal to the daemon (`pkill -HUP ups-agent`), on which
the agent reopens the log file.

## Project structure

| Directory | Purpose |
| ------- | ---------- |
| `card` | Working with the UPS SNMP card: OIDs, polling, traps, status, shutdown. |
| `cardconfig` | Managing the SNMP card configuration (`get`/`set`/`default`). |
| `config` | Loading and validating the configuration, parsing command-line arguments. |
| `log` | Logging with severity levels. |
| `monitor` | The monitoring loop and the shutdown decision. |
| `poweroff` | Shutting down the host and the UPS. |
| `status` | Printing the UPS status. |
| `contrib` | systemd and logrotate files and the trap map generation utility. |
