# Running pg_tileserv as a systemd service

by: Regina Obe - lr (@) pcorp.us

This example demonstrates how to run pg_tileserv as a Linux systemd service
To use:

Copy the pg_tileserv.service file to services directory
On Debian/Ubuntu this would be  `/etc/systemd/system/`
On Redhat / CentOS based probably `/usr/lib/systemd/system`

Make edits as necessary. Things you might want to change:

Create a user that will run this service, making sure that user has rights
to the working directory and executable e.g.
```
sudo useradd pgtileserv
```

The account it runs under, change to whatever account you want
```
User=pgtileserv
Group=pgtileserv
```

The path to working directory and runtime
```
WorkingDirectory=/pg_tileserv/
ExecStart=/pg_tileserv/pg_tileserv --config /pg_tileserv/config/pg_tileserv.toml
```

Once you have made the edits, do
```
sudo systemctl daemon-reload
sudo systemctl enable pg_tileserv #this will make it start on server restarts
sudo systemctl start
#confirm it's running
sudo systemctl status pg_tileserv
```

To stop and start
```
sudo systemctl stop pg_tileserv
sudo systemctl restart pg_tileserv
#confirm it's running
sudo systemctl status pg_tilese
```
