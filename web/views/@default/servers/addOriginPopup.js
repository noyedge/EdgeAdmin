Tea.context(function () {
	this.addr = ""
	this.protocol = ""

	this.addrError = ""

	// 当前服务协议
	this.isHTTP = (this.serverType == "httpProxy" || this.serverType == "httpWeb")
	if (this.serverType == "httpProxy") {
		this.protocol = "http"
	} else if (this.serverType == "tcpProxy") {
		this.protocol = "tcp"
	} else if (this.serverType == "udpProxy") {
		this.protocol = "udp"
	}

	this.changeProtocol = function () {
		this.checkPort()
	}

	this.changeAddr = function () {
		if (this.serverType == "httpProxy") {
			if (this.addr.startsWith("http://")) {
				this.protocol = "http"
			} else if (this.addr.startsWith("https://")) {
				this.protocol = "https"
			}
		}

		this.checkPort()
	}

	this.checkPort = function () {
		this.addrError = ""

		// HTTP
		if (this.protocol == "http") {
			if (this.addr.endsWith(":443")) {
				this.addrError = "443通常是HTTPS协议端口，请确认源站协议选择是否正确。"
			}
		}

		// HTTPS
		if (this.protocol == "https") {
			if (this.addr.endsWith(":80")) {
				this.addrError = "80通常是HTTP协议端口，请确认源站协议选择是否正确。"
			}
		}
	}
})