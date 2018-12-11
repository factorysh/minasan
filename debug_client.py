#!/usr/bin/env python
import smtplib
import sys

server = smtplib.SMTP(host='localhost', port=2525)
server.set_debuglevel(1)
if len(sys.argv) > 1:
    mail = sys.argv[1]
else:
    mail = "factory.furonto@example.com"
server.sendmail("debug@example.com", mail, """From: robert@example.com
Subject: super mail

This is fine""")
server.quit()