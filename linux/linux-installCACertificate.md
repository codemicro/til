# Install a root CA certificate to an Ubuntu (or derivatives) system

```bash
mkdir -p /usr/share/ca-certificates/extra
cd /usr/share/ca-certificates/extra
# wget <certificate url>
sudo dpkg-reconfigure ca-certificates

# Certificates to install must be in .crt format.

# to convert .pem to .crt
# openssl x509 -in foo.pem -inform PEM -out bar.crt

# to convert .cer to .crt
# openssl x509 -in foo.cer -inform DER -out bar.crt
```
