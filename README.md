# Find binary and dynamic so library dependence belongs to which RPM

## Usage

find all RPM deps with missing info

	# finddeprpm /opt/f8n/nginx 
	krb5-libs	1.12.2
	libcom_err	1.42.9
	xz-libs	5.1.2
	libjpeg-turbo	1.2.90
	keyutils-libs	1.5.8
	libXau	1.0.8
	pcre	8.32
	libxslt	1.1.28
	...

just find RPM deps existed without missing info

	# finddeprpm /opt/f8n/nginx 2> /dev/null

find missing library

	# finddeprpm /opt/f8n/nginx 1> /dev/null