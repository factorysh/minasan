#!/usr/bin/env python
import smtplib
import sys

server = smtplib.SMTP(host='localhost', port=2525)
server.set_debuglevel(1)
if len(sys.argv) > 1:
    mail = sys.argv[1]
else:
    mail = "factory.minasan@example.com"
server.sendmail("debug@example.com", mail, """From: robert
Subject: super mail

This is fine""")
server.quit()