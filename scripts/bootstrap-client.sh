#!/bin/bash

export IP_ADDRESS=$(curl -s -H "Metadata-Flavor: Google" \
  http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/ip)

apt-get update
# Install unzip and dnsmasq
apt-get install -y unzip dnsmasq

# Download and install Nomad
wget https://releases.hashicorp.com/nomad/0.6.2/nomad_0.6.2_linux_amd64.zip
unzip nomad_0.6.2_linux_amd64.zip
mv nomad /usr/local/bin/

mkdir -p /var/lib/nomad
mkdir -p /etc/nomad

# clean up Nomad zipfile
rm nomad_0.6.2_linux_amd64.zip

cat > client.hcl <<EOF
addresses {
    rpc  = "ADVERTISE_ADDR"
    http = "ADVERTISE_ADDR"
}
advertise {
    http = "ADVERTISE_ADDR:4646"
    rpc  = "ADVERTISE_ADDR:4647"
}
data_dir  = "/var/lib/nomad"
log_level = "DEBUG"
client {
    enabled = true
    servers = [
      "nomad-1", "nomad-2", "nomad-3"
    ]
    options {
        "driver.raw_exec.enable" = "1"
    }
}
EOF
# Replace ADVERTISE_ADDR with IP address
sed -i "s/ADVERTISE_ADDR/${IP_ADDRESS}/" client.hcl
mv client.hcl /etc/nomad/client.hcl

cat > nomad.service <<'EOF'
[Unit]
Description=Nomad
Documentation=https://nomadproject.io/docs/
[Service]
ExecStart=/usr/local/bin/nomad agent -config /etc/nomad
ExecReload=/bin/kill -HUP $MAINPID
LimitNOFILE=65536
[Install]
WantedBy=multi-user.target
EOF
# Install systemd file
mv nomad.service /etc/systemd/system/nomad.service

systemctl enable nomad
systemctl start nomad

# Configure dnsmasq
mkdir -p /etc/dnsmasq.d
cat > /etc/dnsmasq.d/10-consul <<'EOF'
server=/consul/127.0.0.1#8600
EOF

systemctl enable dnsmasq
systemctl start dnsmasq

# Install docker
apt-get install -y docker.io